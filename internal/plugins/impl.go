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

package plugins

import (
	"bytes"
	"go/ast"
	"go/token"
	"path/filepath"
	"sort"
	"strings"

	"github.com/stoewer/go-strcase"

	"github.com/go-zing/gozz/zcore"
	"github.com/go-zing/gozz/zutils"
)

func init() {
	zcore.RegisterPlugin(Impl{})
}

const (
	implMethodTemplate  = "func(%s %s)%s%s{\npanic(\"not implemented\")}"
	implTypeAssert      = "_ %s = (*%s)(nil)"
	implWireAnnotation  = "// %s%s:bind=%s%s\n"
	implTypeDeclaration = "var ( %s )\n\n%stype %s struct{}"
)

type (
	Impl struct{}

	implDstKey struct {
		Package  string
		Typename string
	}

	implDstType struct {
		Entities zcore.DeclEntities
		Filename string
		Methods  map[string]*ImplMethod
	}

	ImplMethod struct {
		Ptr      bool
		Filename string
		SrcFile  *zutils.File
		SrcType  *ast.FuncType
		DstFile  *zutils.File
		DstType  *ast.FuncType
		Order    int
	}
)

func (i Impl) Name() string { return "impl" }

func (i Impl) Args() ([]string, map[string]string) {
	return []string{"filename"}, nil
}

func (i Impl) Description() string { return "" }

func (dst *implDstType) init(modifySet *zutils.ModifySet, key implDstKey) {
	var (
		wires   = make([]string, 0)
		impls   = make([]string, 0)
		added   = make(map[*ast.TypeSpec]struct{})
		modify  = modifySet.Add(dst.Filename)
		dstPath = zutils.GetImportPath(dst.Filename)
	)

	for _, entity := range dst.Entities {
		if _, exist := added[entity.TypeSpec]; exist {
			continue
		}
		added[entity.TypeSpec] = struct{}{}

		name := entity.Name()

		// import interface package if implements in different package
		if srcPath := zutils.GetImportPath(entity.File.Path); dstPath != srcPath {
			name = modify.Imports.Add(srcPath) + "." + name
		}

		// add type implement
		zutils.Appendf(&impls, implTypeAssert, name, key.Typename)

		// add wire annotations
		if entity.Options.Exist("wire") {
			aop := ""
			if entity.Options.Exist("aop") {
				aop = ":aop"
			}
			zutils.Appendf(&wires, implWireAnnotation, zcore.AnnotationPrefix, wireName, name, aop)
		}
	}

	// append type defines
	modify.Appends = append(modify.Appends,
		zutils.Bytesf(implTypeDeclaration, strings.Join(impls, ";"), strings.Join(wires, ""), key.Typename))
}

func (dst *implDstType) apply(set *zutils.ModifySet, key implDstKey) (err error) {
	recName := ""
	typeDecl := false

	// parse package type methods
	files, err := zutils.WalkPackage(key.Package, func(file *zutils.File) (err error) {
		// check implement type declared
		typeDecl = typeDecl || file.Lookup(key.Typename) != nil

		// for each declaration
		for _, decl := range file.Ast.Decls {
			// function must have one receiver
			fd, ok := decl.(*ast.FuncDecl)
			if !ok || fd.Recv == nil || len(fd.Recv.List) != 1 || fd.Name == nil {
				continue
			}

			rec := fd.Recv.List[0]

			// check function receiver type
			if !bytes.Equal(bytes.TrimPrefix(file.Node(rec.Type), []byte{'*'}), []byte(key.Typename)) {
				continue
			}

			// use exist receiver name
			if len(rec.Names) > 0 && len(rec.Names[0].Name) > 0 {
				recName = rec.Names[0].Name
			}

			// exist method
			if method, match := dst.Methods[fd.Name.Name]; match {
				method.DstType = fd.Type
				method.DstFile = file
			}
		}
		return
	})
	if err != nil {
		return
	}

	// init type declaration
	if !typeDecl {
		dst.init(set, key)
	}

	// no exist receiver name found use lower case typename
	if len(recName) == 0 {
		sp := strings.Split(strcase.SnakeCase(key.Typename), "_")
		recName = strings.ToLower(sp[len(sp)-1])
	}

	// sort by order
	names := make([]string, 0, len(dst.Methods))
	for name := range dst.Methods {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool { return dst.Methods[names[i]].Order < dst.Methods[names[j]].Order })

	// implement methods
	for _, name := range names {
		method := dst.Methods[name]
		if f, ok := files[method.Filename]; ok && method.DstFile == nil {
			method.DstFile = f
		}

		// add modify file
		file := set.Add(method.Filename)
		if method.DstFile != nil {
			// use exist method dst file
			if file = set.Add(method.DstFile.Path); len(file.Imports) == 0 {
				// init modify imports
				file.Imports = method.DstFile.Imports()
			}
		}

		if sign := method.SrcFile.ReplacePackages(method.SrcType, method.Filename, file.Imports); method.DstType == nil {
			// method not implements in package
			typename := key.Typename
			if method.Ptr {
				// pointer receiver
				typename = "*" + typename
			}
			file.Appends = append(file.Appends, zutils.Bytesf(implMethodTemplate, recName, typename, name, sign))
		} else if ft := (&ast.FuncType{
			Func:    token.NoPos,            // must unset func position to adjust func type offset
			Params:  method.DstType.Params,  // params types
			Results: method.DstType.Results, // results types
		}); !bytes.Equal(method.DstFile.Node(ft), sign) {
			// add dst method node modify if bytes not equal
			file.Nodes[ft] = sign
		}
	}
	return
}

func (i Impl) Run(entities zcore.DeclEntities) (err error) {
	group := i.group(entities)
	set := new(zutils.ModifySet)
	eg := new(zutils.ErrGroup)

	for k := range group {
		key := k
		if dst := group[key]; len(dst.Methods) > 0 {
			eg.Go(func() error { return dst.apply(set, key) })
		}
	}

	if err = eg.Wait(); err != nil {
		return
	}

	return set.Apply()
}

func (i Impl) group(entities zcore.DeclEntities) map[implDstKey]*implDstType {
	types := make(map[implDstKey]*implDstType)

	for _, entity := range entities {
		// must be valid interface type declaration
		it, exist := entity.TypeSpec.Type.(*ast.InterfaceType)
		if !exist || it.Methods == nil || len(it.Methods.List) == 0 {
			continue
		}

		filename := entity.RelFilename(entity.Args[0], "impl.go")
		typename, ptr := zutils.TrimPrefix(entity.Options.Get("type", "*"+entity.Name()+"Impl"), "*")

		// group by package directory and type
		key := implDstKey{Typename: typename, Package: filepath.Dir(filename)}

		// init dst implement type
		dst, exist := types[key]
		if !exist {
			dst = &implDstType{Methods: make(map[string]*ImplMethod), Filename: filename}
			types[key] = dst
		}

		dst.Entities = append(dst.Entities, entity)

		// parse annotated fields entities
		fieldEntities := make(map[string]zcore.FieldEntity)
		for _, field := range entity.ParseFields(1, nil) {
			if name, _, ok := zutils.AssertFuncType(field.Field); ok {
				fieldEntities[name] = field
			}
		}

		// register method types
		for _, method := range it.Methods.List {
			name, ft, ok := zutils.AssertFuncType(method)
			if !ok {
				continue
			}

			// check field extra annotated
			if field, ext := fieldEntities[name]; ext {
				if field.Args[0] == "-" {
					continue
				}
			}

			// assign method to implements
			dst.Methods[name] = &ImplMethod{
				Ptr:      ptr,
				Filename: filename,
				SrcFile:  entity.File,
				SrcType:  ft,
				Order:    len(dst.Methods),
			}
		}
	}

	return types
}
