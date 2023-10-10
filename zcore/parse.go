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

package zcore

import (
	"bytes"
	"go/ast"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/Just-maple/gozz/internal/ztore"
	"github.com/Just-maple/gozz/zutils"
)

var (
	// SkipDirs contains some directory name would skip in walk
	SkipDirs = map[string]struct{}{
		"vendor":       {},
		"node_modules": {},
		"testdata":     {},
	}

	// declParsedStore to cached parsed AnnotatedDecls from *ast.File
	// same *ast.File always has same parsed results
	declParsedStore = new(ztore.VersionStore)
)

// Types of annotated declaration
const (
	DeclTypeInterface = iota + 1 // type T interface{}
	DeclTypeStruct               // type T struct{}
	DeclTypeMap                  // type T map[string]string
	DeclTypeArray                // type T []string
	DeclTypeFunc                 // type T func()
	DeclTypeRefer                // type T T2  or  type T = T2
	DeclFunc                     // func Fn()
	DeclValue                    // var variable = 1 or var v Type  or  const constant = 1
)

type (
	AnnotatedDecl struct {
		File *zutils.File

		Type      int
		FuncDecl  *ast.FuncDecl
		TypeSpec  *ast.TypeSpec
		ValueSpec *ast.ValueSpec

		Docs        []string
		Annotations []string
		Fields      []*AnnotatedField
	}

	AnnotatedField struct {
		Decl        *AnnotatedDecl
		Field       *ast.Field
		Docs        []string
		Annotations []string
	}

	AnnotatedDecls []*AnnotatedDecl
)

// ParseFileOrDirectory try parse provided path annotated declarations with annotations prefix
// if directory provided. walks file tree from provided path as root and returns all parsed
func ParseFileOrDirectory(path string, prefix string) (decls AnnotatedDecls, err error) {
	stat, err := os.Stat(path)
	if err != nil {
		return
	}

	if !stat.IsDir() {
		// single file
		return ParseFileDecls(path, prefix)
	}

	// directory
	// walk all child directories and files

	// use error group and pre alloc slots to collect parsed results
	slots := make([]*AnnotatedDecls, 0)
	eg := new(zutils.ErrGroup)

	if err = filepath.Walk(path, func(filename string, info fs.FileInfo, e error) (err error) {
		if e != nil {
			return e
		}

		if info.IsDir() {
			name := info.Name()
			// some specific skip name or dirs starts with .
			if _, skip := SkipDirs[name]; skip || strings.HasPrefix(name, ".") {
				return filepath.SkipDir
			}
			return
		}

		// parse files with goroutine error group
		// results would be placed in slot
		index := len(slots)
		slots = append(slots, new(AnnotatedDecls))
		eg.Go(func() (err error) {
			*slots[index], err = ParseFileDecls(filename, prefix)
			return
		})
		return
	}); err != nil {
		return
	}

	// wait for all results or error
	if err = eg.Wait(); err != nil {
		return
	}

	// expand results from slots
	for _, slot := range slots {
		decls = append(decls, *slot...)
	}
	return
}

// ParseFileDecls parse provided file into ast and analysis declarations annotations
// return annotated declarations list or error while reading file or parsing ast
func ParseFileDecls(filename string, prefix string) (decls AnnotatedDecls, err error) {
	filename, err = filepath.Abs(filename)
	if err != nil {
		return
	}

	if !zutils.IsGoFile(filename) {
		return
	}

	// read file data
	data, version, err := zutils.ReadFile(filename)
	if err != nil {
		return
	}

	// check data contains annotations prefix or return
	if !bytes.Contains(data, []byte(prefix)) {
		return
	}

	// parse file ast
	f, err := zutils.ParseFile(filename)
	if err != nil {
		return
	}

	// parse annotated decls
	ret, _ := declParsedStore.Load(f.Ast, version, func() (interface{}, error) {
		return parseFileDecls(f, prefix), nil
	})

	decls = ret.(AnnotatedDecls)
	return
}

func parseFileDecls(file *zutils.File, prefix string) (decls AnnotatedDecls) {
	for _, astDecl := range file.Ast.Decls {
		for _, decl := range ParseDecls(astDecl, prefix) {
			decl.File = file
			decls = append(decls, decl)
		}
	}
	return
}

// ParseGenericDecl parse generic declaration to match annotation prefix
func ParseGenericDecl(gen *ast.GenDecl, prefix string) (decls AnnotatedDecls) {
	genDocs, genAnnotations := ParseCommentGroup(prefix, gen.Doc)

	single := !gen.Lparen.IsValid() || len(gen.Specs) == 1

	switch gen.Tok {
	case token.CONST, token.VAR:
		/*
			merged type declaration for variable or constant

			// +zz:annotation:args:key=value
			var (
			    variableA = 1
			    variableB = 2
			)

			// +zz:annotation:args:key=value
			var variableC = 4

			// +zz:annotation:args:key=value
			const (
			    constantA = 3
			    constantB = 4
			)

			// +zz:annotation:args:key=value
			const constantC = 4
		*/
		for _, spec := range gen.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}

			docs, annotations := ParseCommentGroup(prefix, vs.Doc, vs.Comment)
			// generic annotations would be appended to each element in merged declaration

			if annotations = append(genAnnotations, annotations...); len(annotations) == 0 {
				continue
			}

			if single {
				docs = append(genDocs, docs...)
			}

			decls = append(decls, &AnnotatedDecl{
				ValueSpec:   vs,
				Docs:        docs,
				Annotations: annotations,
				Type:        DeclValue,
			})
		}

	case token.TYPE:
		/*
			separated struct or interface type declaration

			// +zz:annotation:args:key=value
			type structA struct{
				Field0 int
				Field1 int
			}

			// +zz:annotation:args:key=value
			type structB struct{
				Field0 int
				Field1 int
			}

			// +zz:annotation:args:key=value
			type interfaceC interface{
				Foo()
			}

			same annotation for grouped types can use
			merged type declaration
			would be same effect as upper

			// +zz:annotation:args:key=value
			type (
			    structA struct{
					Field0 int
					Field1 int
				}

				structB struct{
					Field0 int
					Field1 int
				}

				interfaceC interface{
					Foo()
				}
			)
		*/
		for _, s := range gen.Specs {
			spec, ok := s.(*ast.TypeSpec)
			if !ok {
				continue
			}

			docs, annotations := ParseCommentGroup(prefix, spec.Doc, spec.Comment)

			// generic annotations would be appended to each element in merged declaration
			if annotations = append(annotations, genAnnotations...); len(annotations) == 0 {
				continue
			}

			if single {
				docs = append(genDocs, docs...)
			}

			decl := &AnnotatedDecl{
				TypeSpec:    spec,
				Docs:        docs,
				Annotations: annotations,
			}

			// check type spec type
			switch typ := spec.Type.(type) {
			case *ast.InterfaceType:
				decl.Type = DeclTypeInterface
				decl.parseAnnotatedFields(typ.Methods, prefix)
			case *ast.StructType:
				decl.Type = DeclTypeStruct
				decl.parseAnnotatedFields(typ.Fields, prefix)
			case *ast.MapType:
				decl.Type = DeclTypeMap
			case *ast.ArrayType:
				decl.Type = DeclTypeArray
			case *ast.FuncType:
				decl.Type = DeclTypeFunc
			case *ast.Ident, *ast.SelectorExpr, *ast.StarExpr:
				decl.Type = DeclTypeRefer
			default:
				continue
			}

			decls = append(decls, decl)
		}
	}
	return
}

// parseAnnotatedFields parse fields docs and comments to match annotations prefix
// fields match annotations will be collect as AnnotatedField
func (decl *AnnotatedDecl) parseAnnotatedFields(fl *ast.FieldList, prefix string) {
	for _, field := range fl.List {
		if len(field.Names) == 0 {
			continue
		}
		if docs, annotations := ParseCommentGroup(prefix, field.Doc, field.Comment); len(annotations) > 0 {
			decl.Fields = append(decl.Fields, &AnnotatedField{
				Docs:        docs,
				Annotations: annotations,
				Field:       field,
				Decl:        decl,
			})
		}
	}
}

// ParseFuncDecl parse function declaration docs to match annotations prefix
//
// Example:
//
// +zz:annotation:args:key=value
// func Foo() {
// }
func ParseFuncDecl(decl *ast.FuncDecl, prefix string) (d *AnnotatedDecl) {
	docs, annotations := ParseCommentGroup(prefix, decl.Doc)
	if len(annotations) == 0 {
		return nil
	}
	return &AnnotatedDecl{
		FuncDecl:    decl,
		Docs:        docs,
		Annotations: annotations,
		Type:        DeclFunc,
	}
}

// ParseDecls check declaration type
// parse generic declaration or function declaration and get annotated declarations
func ParseDecls(d ast.Decl, prefix string) (items AnnotatedDecls) {
	switch decl := d.(type) {
	case *ast.GenDecl:
		items = append(items, ParseGenericDecl(decl, prefix)...)
	case *ast.FuncDecl:
		if item := ParseFuncDecl(decl, prefix); item != nil {
			items = append(items, item)
		}
	}
	return
}

// ParseCommentGroup extract comment group text and split by lines
// if line match annotation prefix then append line to annotations
// else append line to docs
func ParseCommentGroup(prefix string, cg ...*ast.CommentGroup) (docs, annotations []string) {
	for _, g := range cg {
		if g == nil {
			continue
		}
		docs = append(docs, strings.Split(strings.TrimSpace(g.Text()), "\n")...)
	}

	// no prefix provided. return all comment lines as doc
	if len(prefix) == 0 {
		return docs, nil
	}

	// comments matched annotation prefix would be appended as annotations
	// or appended as docs in same slice memory
	offset := 0
	for _, doc := range docs {
		if annotation, exist := zutils.TrimPrefix(strings.TrimSpace(doc), prefix); exist {
			annotations = append(annotations, annotation)
		} else {
			docs[offset] = doc
			offset++
		}
	}
	docs = docs[:offset]
	return
}
