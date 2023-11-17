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

	zcore "github.com/go-zing/gozz-core"
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
		RecvName string
		Methods  map[string]*ImplMethod
		Files    map[string]*zcore.File
	}

	ImplMethod struct {
		Ptr      bool
		Filename string
		SrcFile  *zcore.File
		SrcType  *ast.FuncType
		DstFile  *zcore.File
		DstType  *ast.FuncType
		Order    int
	}
)

func (i Impl) Name() string { return "impl" }

func (i Impl) Args() ([]string, map[string]string) {
	return []string{"filename:specify which file to generate implements type and method if not exist"},
		map[string]string{
			"aop":  "add aop wire options when generate implements type declaration. bool flag option",
			"wire": "add wire annotation when generate implements type declaration. bool flag option",
			"type": "specify implement typename. add * as prefix if use pointer type receiver. example: [ type=*T ]",
		}
}

func (i Impl) Description() string {
	return "generate and sync interface functions type signature to target implement type."
}

func (dst *implDstType) init(modifySet *zcore.ModifySet, key implDstKey) {
	var (
		wires   = make([]string, 0)
		impls   = make([]string, 0)
		added   = make(map[*ast.TypeSpec]struct{})
		modify  = modifySet.Add(dst.Filename)
		dstPath = zcore.GetImportPath(dst.Filename)
	)

	for _, entity := range dst.Entities {
		if _, exist := added[entity.TypeSpec]; exist {
			continue
		}
		added[entity.TypeSpec] = struct{}{}

		name := entity.Name()

		// import interface package if implements in different package
		if srcPath := zcore.GetImportPath(entity.File.Path); dstPath != srcPath {
			name = modify.Imports.Add(srcPath) + "." + name
		}

		// add type implement
		zcore.Appendf(&impls, implTypeAssert, name, key.Typename)

		// add wire annotations
		if entity.Options.Exist("wire") {
			aop := ""
			if entity.Options.Exist("aop") {
				aop = ":aop"
			}
			zcore.Appendf(&wires, implWireAnnotation, zcore.AnnotationPrefix, wireName, name, aop)
		}
	}

	// append type defines
	modify.Appends = append(modify.Appends,
		zcore.Bytesf(implTypeDeclaration, strings.Join(impls, ";"), strings.Join(wires, ""), key.Typename))
}

func (dst *implDstType) prepare(set *zcore.ModifySet, key implDstKey) (err error) {
	typeDecl := false

	// parse package type methods
	dst.Files, err = zcore.WalkPackage(key.Package, func(file *zcore.File) (err error) {
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
				dst.RecvName = rec.Names[0].Name
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
	if len(dst.RecvName) == 0 {
		sp := strings.Split(zcore.SnakeCase(key.Typename), "_")
		dst.RecvName = strings.ToLower(sp[len(sp)-1])
	}
	return
}

func (dst *implDstType) apply(set *zcore.ModifySet, key implDstKey) (err error) {
	// sort by order
	names := make([]string, 0, len(dst.Methods))
	for name := range dst.Methods {
		names = append(names, name)
	}
	sort.Slice(names, func(i, j int) bool { return dst.Methods[names[i]].Order < dst.Methods[names[j]].Order })

	// implement methods
	for _, name := range names {
		method := dst.Methods[name]
		if f, ok := dst.Files[method.Filename]; ok && method.DstFile == nil {
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
			file.Appends = append(file.Appends, zcore.Bytesf(implMethodTemplate, dst.RecvName, typename, name, sign))
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
	set := new(zcore.ModifySet)
	keys := make([]implDstKey, 0)
	for key := range group {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Package+keys[i].Typename < keys[j].Package+keys[j].Typename
	})

	for _, key := range keys {
		if dst := group[key]; len(dst.Methods) > 0 {
			if err = dst.prepare(set, key); err != nil {
				return
			}
			zcore.Logger.Printf("implement %s:%s\n", key.Package, key.Typename)
			if err = dst.apply(set, key); err != nil {
				return
			}
		}
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
		typename, ptr := zcore.TrimPrefix(entity.Options.Get("type", "*"+entity.Name()+"Impl"), "*")

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
			if name, _, ok := zcore.AssertFuncType(field.Field); ok {
				fieldEntities[name] = field
			}
		}

		// register method types
		for _, method := range it.Methods.List {
			name, ft, ok := zcore.AssertFuncType(method)
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
