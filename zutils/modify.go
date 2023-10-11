/*
 * Copyright (c) 2023 Maple Wu <justmaplewu@gmail.com>
 *   National Electronics and Computer Technology Center, Thailand
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package zutils

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path"
	"sort"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/tools/go/ast/astutil"
)

type (
	// File store parsed *ast.File and data bytes
	File struct {
		Path string
		Data []byte
		Ast  *ast.File
	}

	// ModifySet store inited *Modify with filename as key
	ModifySet struct {
		set map[string]*Modify
		mu  sync.Mutex
	}

	// Modify store changes of target filename golang ast modify
	Modify struct {
		Filename string
		Imports  Imports
		Nodes    map[ast.Node][]byte
		Appends  [][]byte
	}

	// Imports represents a key-value store with import path as key and import name as value
	// it helps to deduplicated import path and rotate import name if duplicated
	Imports map[string]string

	// Import contains import name and import path
	// this struct is used to render templates
	Import struct {
		Name string
		Path string
	}
)

// Node return data bytes of ast.Node
func (f *File) Node(node ast.Node) []byte {
	return f.Data[node.Pos()-1 : node.End()-1]
}

// Lookup is a shortcut of ast scope lookup
func (f *File) Lookup(name string) *ast.Object {
	return f.Ast.Scope.Lookup(name)
}

func (f *File) nodeReplacer(node ast.Node) bytesReplacer {
	return bytesReplacer{origin: f.Node(node), offset: int(node.Pos() - 1)}
}

// Imports dumps file ast imports list and return Imports map
func (f *File) Imports() Imports {
	return LoadImports(f.Ast)
}

// ReplacePackages try replaces type package selector to provide node according to dst filename
// return modified node data bytes
func (f *File) ReplacePackages(node ast.Node, dstFilename string, dstImports Imports) (data []byte) {
	if node == nil {
		return
	}
	pr := &packagesReplacer{
		bytesReplacer: f.nodeReplacer(node),
		srcImports:    LoadImports(f.Ast),
		dstImports:    dstImports,
		srcImportPath: GetImportPath(f.Path),
		dstImportPath: GetImportPath(dstFilename),
	}
	if len(pr.srcImportPath) == 0 || len(pr.dstImportPath) == 0 {
		return nil
	}
	ast.Walk(pr, node)
	return pr.Bytes()
}

// packagesReplacer implements ast.Visitor to modify visited ast node type packages
type packagesReplacer struct {
	bytesReplacer
	srcImports    Imports
	dstImports    Imports
	srcImportPath string
	dstImportPath string
}

// Visit implements ast.Visitor to walk each ast.Node and replace type packages
func (pr *packagesReplacer) Visit(node ast.Node) (w ast.Visitor) {
	switch n := node.(type) {
	default:
		return pr
	case *ast.Field:
		// only walk field type
		ast.Walk(pr, n.Type)
	case *ast.SelectorExpr:
		// exist package selector
		// check package import path
		if p := pr.srcImports.Which(fmt.Sprintf("%s", n.X)); len(p) > 0 {
			if p == pr.dstImportPath {
				// import path equals to dst import path
				// remove package selector
				pr.ReplaceAstNode(n, []byte(n.Sel.Name))
			} else if name := pr.dstImports.Add(p); name == "." {
				// dst imports as .
				// remove package selector
				pr.ReplaceAstNode(n, []byte(n.Sel.Name))
			} else {
				// update package selector as dst imports
				pr.ReplaceAstNode(n.X, []byte(name))
			}
		}
	case *ast.Ident:
		// not exist package selector
		if pr.srcImportPath != pr.dstImportPath && ast.IsExported(n.Name) {
			// try adds package selector
			if name := pr.dstImports.Add(pr.srcImportPath); name != "." {
				pr.ReplaceAstNode(n, []byte(name+"."+n.Name))
			}
		}
	}

	return nil
}

// LoadImports parse *ast.File import spec list as Imports map
func LoadImports(f *ast.File) (imps Imports) {
	imps = make(Imports, len(f.Imports))
	for _, imp := range f.Imports {
		name := ""
		if imp.Name != nil {
			// specified import name
			name = imp.Name.Name
		}

		if name == "_" {
			continue
		}

		p, e := strconv.Unquote(imp.Path.Value)
		if e != nil {
			continue
		}

		if len(name) == 0 {
			if IsStandardImportPath(p) {
				name = path.Base(p)
			} else if name = GetImportName(p); len(name) == 0 {
				// no import name specified
				// default import name from GetImportName
				// or use base path
				name = path.Base(p)
			}
		}
		imps[p] = name
	}
	return
}

func (imps Imports) add(path, as string) string {
	if len(imps.Which(as)) == 0 {
		imps[path] = as
		return as
	}
	return imps.add(path, as+"2")
}

// Add check import path exist or adds into imports and return imports name
func (imps Imports) Add(p string) (name string) {
	if n, exist := imps[p]; exist {
		return n
	}
	return imps.add(p, importNameReplacer.Replace(path.Base(p)))
}

var importNameReplacer = strings.NewReplacer("-", "", ".", "")

// Which check import name exist and return import path
func (imps Imports) Which(name string) (path string) {
	for p, n := range imps {
		if n == name {
			return p
		}
	}
	return ""
}

// List convert Imports map into sorted Import slice
func (imps Imports) List() []Import {
	list := make([]Import, 0, len(imps))
	for p, name := range imps {
		list = append(list, Import{
			Name: name,
			Path: p,
		})
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].Path < list[j].Path
	})
	return list
}

// Apply handles all filenames in ModifySet and apply all Modify
func (set *ModifySet) Apply() (err error) {
	set.mu.Lock()
	defer set.mu.Unlock()

	var filenames []string
	for filename := range set.set {
		filenames = append(filenames, filename)
	}

	sort.Strings(filenames)

	for _, filename := range filenames {
		if err = set.set[filename].Apply(); err != nil {
			return
		}
	}
	return
}

// Add try adds filename and alloc *Modify object into ModifySet
func (set *ModifySet) Add(filename string) *Modify {
	set.mu.Lock()
	defer set.mu.Unlock()
	m, ok := set.set[filename]
	if !ok {
		m = &Modify{
			Filename: filename,
			Imports:  make(Imports),
			Nodes:    make(map[ast.Node][]byte),
		}
		if set.set == nil {
			set.set = make(map[string]*Modify)
		}
		set.set[filename] = m
	}
	return m
}

// applyImports to update provide data (golang file) imports and return updated data
func (m *Modify) applyImports(data []byte) (ret []byte, err error) {
	if len(m.Imports) == 0 {
		return data, nil
	}

	// parse provided data
	fileSet := token.NewFileSet()
	fileAst, err := parser.ParseFile(fileSet, "", data, parser.ParseComments)
	if err != nil {
		return
	}

	// parse exists import spec list
	exists := LoadImports(fileAst)

	for _, imp := range m.Imports.List() {
		if _, exist := exists[imp.Path]; exist {
			continue
		}
		// add import if not exist
		if IsStandardImportPath(imp.Path) && path.Base(imp.Path) == imp.Name {
			astutil.AddImport(fileSet, fileAst, imp.Path)
		} else {
			astutil.AddNamedImport(fileSet, fileAst, imp.Name, imp.Path)
		}
	}

	bf := BuffPool.Get().(*bytes.Buffer)
	bf.Reset()
	// format as bytes
	if err = format.Node(bf, fileSet, fileAst); err != nil {
		return
	}
	return bf.Bytes(), nil
}

func (m *Modify) applyAppends(data []byte) []byte {
	for _, appendData := range m.Appends {
		data = append(data, '\n')
		data = append(data, appendData...)
		data = append(data, '\n')
	}
	return data
}

// applyNodes replace provided data and replace registered ast node position data to new data
func (m *Modify) applyNodes(data []byte) []byte {
	if len(m.Nodes) == 0 {
		return data
	}

	nodes := make([]ast.Node, 0, len(m.Nodes))
	for node := range m.Nodes {
		if node != nil {
			nodes = append(nodes, node)
		}
	}

	sort.Slice(nodes, func(i, j int) bool { return nodes[i].Pos() < nodes[j].Pos() })
	replacer := &bytesReplacer{origin: data}
	for _, node := range nodes {
		replacer.ReplaceAstNode(node, m.Nodes[node])
	}
	return replacer.Bytes()
}

// Apply to updates all modify changes and write data to target filename
func (m *Modify) Apply() (err error) {
	if len(m.Appends)+len(m.Imports)+len(m.Nodes) == 0 {
		return
	}
	data, err := m.apply()
	if err != nil {
		return
	}
	_, err = WriteFile(m.Filename, data, 0o664)
	return
}

// apply to load filename data or create new file with filename
// then handle all modify changes. return formatted modified data bytes
func (m *Modify) apply() (data []byte, err error) {
	// read file data
	if data, _, err = ReadFile(m.Filename); err != nil {
		if !os.IsNotExist(err) {
			return
		}
		data = []byte("package " + GetImportName(m.Filename))
	}

	// nodes
	data = m.applyNodes(data)
	// appends
	data = m.applyAppends(data)
	// imports
	data2, err := m.applyImports(data)
	if err != nil {
		return
	}

	// may format in imports
	if bytes.Equal(data2, data) {
		return format.Source(data2)
	}

	return data2, nil
}

// bytesReplacer to do multiple bytes replacement by position offsets
// all replacements should according to origin position and replacer would adjust replace offsets
type bytesReplacer struct {
	origin  []byte
	updated *bytes.Buffer
	offset  int
}

// Bytes return replaced data
func (r *bytesReplacer) Bytes() []byte {
	if r.updated == nil {
		return r.origin
	}
	return r.updated.Bytes()
}

// ReplaceAstNode replace bytes by ast.Node
func (r *bytesReplacer) ReplaceAstNode(node ast.Node, target []byte) {
	r.Replace(int(node.Pos())-1, int(node.End())-1, target)
}

// Replace bytes to dst bytes by start,end position offset
func (r *bytesReplacer) Replace(start, end int, dst []byte) {
	if r.updated == nil {
		if bytes.Equal(r.origin[start-r.offset:end-r.offset], dst) {
			return
		}
		// alloc buffer and copy origin bytes
		r.updated = BuffPool.Get().(*bytes.Buffer)
		r.updated.Reset()
		r.updated.Write(r.origin)
	}

	updated := r.updated.Bytes()

	// compute replaced offset
	offset := len(updated) - len(r.origin) - r.offset
	start += offset
	end += offset

	// skip if src equal to dst
	if bytes.Equal(updated[start:end], dst) {
		return
	}

	// alloc new buffer to store replaced
	bf := BuffPool.Get().(*bytes.Buffer)
	bf.Reset()

	// do replace
	bf.Write(updated[:start])
	bf.Write(dst)
	bf.Write(updated[end:])

	// release previous buffer and update buffer
	BuffPool.Put(r.updated)
	r.updated = bf
}
