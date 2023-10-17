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
	"fmt"
	"go/ast"
	"os"
	"path/filepath"
	"sort"
	"strings"

	zcore "github.com/go-zing/gozz-core"
)

func init() {
	zcore.RegisterPlugin(&Wire{})
}

type (
	Wire struct {
		Imports []zcore.Import
		Sets    interface{}
		Injects interface{}
		Aops    interface{}
	}

	WireAopMethod struct {
		Name    string
		Params  []string
		Results []string
	}

	WireAop struct {
		Name      string
		Interface string
		Implement string
		Methods   []WireAopMethod
	}

	WireSetElement struct {
		Path  string
		Name  string
		Decls []string
	}

	WireSet struct {
		Name     string
		Elements []WireSetElement
	}

	wireDecl struct {
		Entities    zcore.DeclEntities
		Params      zcore.KeySet
		Binds       zcore.KeySet
		Aops        zcore.KeySet
		Fields      zcore.KeySet
		Injects     zcore.KeySet
		ReferStruct bool
		Provider    string
	}

	wireDeclSet struct {
		m    map[*zcore.AnnotatedDecl]*wireDecl
		keys []*zcore.AnnotatedDecl
	}
)

const (
	wireName               = "wire"
	wireInjectFile         = "wire_zinject.go"
	wireSetFile            = "wire_zset.go"
	wireAopFile            = "wire_zzaop.go"
	wireImportPath         = "github.com/google/wire"
	wireBuildFlag          = "//go:build wireinject\n// +build wireinject"
	wireStructBindTemplate = "wire.Bind(new(%s), new(*%s))"

	wireImportTemplate = `import  (
	{{ range .Imports }} {{ .Name }} "{{ .Path }}"
	{{ end }}
)`

	wireSetTemplate = wireImportTemplate + `
var (
	{{ range .Sets }} _{{ .Name }}Set = wire.NewSet(
		{{ range .Elements }} // {{ .Path }}.{{ .Name }}
		{{ range $decl := .Decls }} {{ $decl }}, 
		{{ end }}
		{{ end }}
	)

	{{end}}
)`

	wireInjectTemplate = wireImportTemplate + `
{{ range .Injects }} // {{ .Path }}.{{ .Name }}
func {{ .Function }}({{ .Params }}) ({{ .Object }},func(),error) {
	panic(wire.Build(_{{.Set}}Set))
}
{{end}}`

	wireAopTemplate = wireImportTemplate + `

type _aop_interceptor interface{ Intercept(v interface{},name string,params,results []interface{}) (func(),bool) }

{{ range .Aops }} {{ $n := .Name }} // {{ .Interface }} 
type (
	{{ $n }} {{ .Interface }}
	_impl{{ $n }} struct{ {{ $n }} }
)

{{ range .Methods }} {{ $p := .Params }} {{ $r := .Results }}
func(i _impl{{ $n }}){{ .Name }}({{ range $i,$v := $p }}p{{ $i }} {{ $v }},{{ end }})({{ range $i,$v := $r }}r{{ $i }} {{ $v }},{{ end }}){
	if t,x:= i.{{ $n }}.(_aop_interceptor);x{
		if up,ok:=t.Intercept( i.{{ $n }},{{ quote .Name }},
		{{ if $p }}[]interface{}{ {{ range $i,$v := $p }}&p{{ $i }},{{ end }} }{{ else }}nil{{ end }},
		{{ if $r }}[]interface{}{ {{ range $i,$v := $r }}&r{{ $i }},{{ end }} }{{ else }}nil{{ end }},
		);up!=nil{
			defer up()
		}else if !ok{
			return
		}
	}
	{{ if $r }}return{{ end }} i.{{ $n }}.{{ .Name }}({{ range $i,$v := $p }}p{{ $i }},{{ end }}) 
}

{{ end }} {{ end }}`
)

func (w Wire) Name() string { return wireName }

func (w Wire) Args() ([]string, map[string]string) {
	return nil, map[string]string{
		"aop":    "generate aop proxy type wrapper for interface binding. bool flag option",
		"struct": "refer type as struct type. only works for type reference declaration. bool flag option",
		"bind":   "bind type with interface. refer to wire.Bind. example: [ bind=io.Writer ]",
		"field":  `provide fields of struct for injects. use "*" to provide all fields for injects. refer to wire.FieldsOf. example: [ field=* ]`,
		"param":  "specify param types for injectors function. example: [ param=context.Context ]",
		"inject": "generate type initialization function in provide filename as injector entries. example: [ inject=./ ]",
		"set":    "specify wire set to provide feature of multi inject namespace. example: [ set=mock ]",
	}
}

func (w Wire) Description() string {
	return "collect and generate wire sets / injectors / aop proxy stubs files."
}

func (w Wire) parseEntitiesDeclSet(entities zcore.DeclEntities) (set *wireDeclSet) {
	set = new(wireDeclSet)

	for _, entity := range entities {
		var (
			decl  = set.init(entity)
			aop   bool
			binds []string
		)

		for key, value := range entity.Options {
			values := strings.Split(value, ",")
			switch key {
			case "bind":
				binds = values
			case "param":
				decl.Params.Add(values)
			case "inject":
				if entity.TypeSpec != nil {
					decl.Injects.Add(values)
				}
			case "field":
				if entity.Type != zcore.DeclTypeStruct {
					continue
				} else if value == "*" {
					decl.Fields.Add(zcore.ExtractStructFieldsNames(entity.TypeSpec.Type.(*ast.StructType)))
				} else {
					decl.Fields.Add(values)
				}
			case "aop":
				aop = true
			case "struct":
				if entity.Type == zcore.DeclTypeRefer {
					decl.ReferStruct = true
				}
			case "set":
				continue
			}
		}

		// add binds
		decl.Binds.Add(binds)

		// add aop
		if aop {
			decl.Aops.Add(binds)
		}
	}
	return
}

func (set *wireDeclSet) init(entity zcore.DeclEntity) *wireDecl {
	decl, ok := set.m[entity.AnnotatedDecl]
	if !ok {
		decl = &wireDecl{
			Params:  make(zcore.KeySet),
			Binds:   make(zcore.KeySet),
			Fields:  make(zcore.KeySet),
			Injects: make(zcore.KeySet),
			Aops:    make(zcore.KeySet),
		}
		if set.m == nil {
			set.m = make(map[*zcore.AnnotatedDecl]*wireDecl)
		}
		set.m[entity.AnnotatedDecl] = decl
		set.keys = append(set.keys, entity.AnnotatedDecl)

		if entity.TypeSpec != nil {
			typename := entity.Name()
			// try lookup provider function named ProvideXXX and return first args is declaration type
			if obj := entity.File.Lookup("Provide" + typename); obj != nil && obj.Decl != nil {
				if p, o := obj.Decl.(*ast.FuncDecl); o && p.Recv == nil && len(p.Type.Results.List) > 0 {
					retType := ""
					switch rt := p.Type.Results.List[0].Type.(type) {
					case *ast.StarExpr:
						retType = string(entity.File.Node(rt.X))
					case *ast.Ident:
						retType = rt.Name
					}
					if typename == retType {
						decl.Provider = p.Name.Name
					}
				}
			}
		}
	}

	decl.Entities = append(decl.Entities, entity)
	return decl
}

func (w Wire) Run(es zcore.DeclEntities) (err error) {
	injectFiles := make(map[string]map[string]zcore.AnnotatedDecls)
	setFiles := make(map[string]map[string]*wireDeclSet)

	// parse sets entities
	for set, decls := range w.parseEntities(es) {
		for _, obj := range decls.keys {
			for _, inject := range decls.m[obj].Injects.Keys() {
				// group by inject file
				inject = obj.RelFilename(inject, wireInjectFile)
				if injectFiles[inject] == nil {
					injectFiles[inject] = make(map[string]zcore.AnnotatedDecls)
				}
				injectFiles[inject][set] = append(injectFiles[inject][set], obj)

				// group by inject set directory
				setFile := filepath.Dir(inject)
				if setFiles[setFile] == nil {
					setFiles[setFile] = make(map[string]*wireDeclSet)
				}
				setFiles[setFile][set] = decls
			}
		}
	}

	// generate sets files
	if err = w.generateSets(setFiles); err != nil {
		return
	}

	// generate injects files
	if err = w.generateInjects(setFiles, injectFiles); err != nil {
		return
	}

	for dir := range setFiles {
		// remove exists wire file
		_ = os.Remove(filepath.Join(dir, "wire_gen.go"))

		// check wire imported or try go get wire
		if _, err = zcore.ExecCommand(fmt.Sprintf("go list -m %s || go get %s", wireImportPath, wireImportPath), dir); err != nil {
			return
		}

		// run wire
		if _, err = zcore.ExecCommand("wire", dir); err != nil {
			return
		}
	}
	return
}

func (w Wire) generateInjects(setFiles map[string]map[string]*wireDeclSet, injectFiles map[string]map[string]zcore.AnnotatedDecls) (err error) {
	eg := new(zcore.ErrGroup)
	for key := range injectFiles {
		filename := key
		eg.Go(func() error {
			return w.generateInject(setFiles[filepath.Dir(filename)], filename, injectFiles[filename])
		})
	}
	return eg.Wait()
}

func (w Wire) generateInject(dirSetFiles map[string]*wireDeclSet, filename string, injects map[string]zcore.AnnotatedDecls) (err error) {
	type WireInject struct {
		Set      string
		Path     string
		Function string
		Name     string
		Params   string
		Object   string
	}

	var (
		dstImportPath = zcore.GetImportPath(filename)
		dstImportName = zcore.GetImportName(filename)
		dstImports    = zcore.Imports{"github.com/google/wire": "wire"}

		wireInjects []*WireInject
	)

	for set, decls := range injects {
		for _, decl := range decls {
			srcImports := decl.File.Imports()
			srcImportPath := zcore.GetImportPath(decl.File.Path)

			fp := func(name string) string {
				return zcore.FixPackage(name, srcImportPath, dstImportPath, srcImports, dstImports)
			}

			wd, ok := dirSetFiles[set].m[decl]
			if !ok {
				continue
			}

			name := decl.Name()
			inject := &WireInject{
				Set:    set,
				Path:   srcImportPath,
				Name:   name,
				Object: fp(name),
			}
			wireInjects = append(wireInjects, inject)

			// build function name
			str := &strings.Builder{}
			str.WriteString("Initialize")
			if len(set) > 0 {
				str.WriteString("_")
				str.WriteString(set)
			}
			str.WriteString("_")
			str.WriteString(strings.Replace(inject.Object, ".", "_", -1))
			inject.Function = str.String()

			// use pointer type if is struct type object
			if decl.Type == zcore.DeclTypeStruct {
				inject.Object = "*" + inject.Object
			}

			// params
			params := wd.Params.Keys()
			for i, param := range params {
				params[i] = fp(param)
			}
			inject.Params = strings.Join(params, ",")
		}
	}

	if len(wireInjects) == 0 {
		return
	}

	sort.Slice(wireInjects, func(i, j int) bool {
		ei := wireInjects[i]
		ej := wireInjects[j]
		return ei.Set < ej.Set && (ei.Path+"."+ei.Object) < (ej.Path+"."+ej.Object)
	})

	return zcore.RenderWrite(Wire{
		Imports: dstImports.List(),
		Injects: wireInjects,
	}, wireInjectTemplate, filename, dstImportName, false, wireBuildFlag)
}

func (w Wire) generateSets(setFiles map[string]map[string]*wireDeclSet) (err error) {
	eg := new(zcore.ErrGroup)
	for key := range setFiles {
		dir := key
		eg.Go(func() error { return w.generateSet(dir, setFiles[dir]) })
	}
	return eg.Wait()
}

func (w Wire) parseInterfaceMethods(name string, dir string, imports zcore.Imports) (methods []WireAopMethod) {
	pkgPath := ""
	dstImportPath := zcore.GetImportPath(dir)

	if strings.Contains(name, ".") {
		sp := strings.Split(name, ".")
		pkgPath, name = imports.Which(sp[0]), sp[1]
	} else {
		pkgPath = dstImportPath
	}

	fl, srcFile := getInterfaceFields(name, dir, pkgPath)
	if fl == nil {
		return
	}

	appendField := func(fl *ast.FieldList, dst *[]string) {
		if fl != nil {
			for _, params := range fl.List {
				l := len(params.Names)
				for i := 0; i < l || (i == 0 && i == l); i++ {
					*dst = append(*dst, zcore.UnsafeBytes2String(srcFile.ReplacePackages(params.Type, dir, imports)))
				}
			}

			l := *dst
			for i := range l {
				if i < len(l)-1 && l[i] == l[i+1] {
					l[i] = ""
				}
			}
		}
	}

	for _, f := range fl.List {
		funcName, ft, ok := zcore.AssertFuncType(f)
		if !ok {
			continue
		}
		m := WireAopMethod{Name: funcName}
		appendField(ft.Params, &m.Params)
		appendField(ft.Results, &m.Results)
		methods = append(methods, m)
	}
	return
}

func getInterfaceFields(name, dir, pkgPath string) (fl *ast.FieldList, srcFile *zcore.File) {
	pkgDir, err := zcore.ExecCommand(`go list -f "{{ .Dir }} " `+pkgPath, dir)
	if err != nil {
		return
	}

	_, _ = zcore.WalkPackage(pkgDir, func(file *zcore.File) (err error) {
		object := file.Lookup(name)
		if object == nil || object.Kind != ast.Typ || object.Decl == nil {
			return
		}

		spec, ok := object.Decl.(*ast.TypeSpec)
		if !ok {
			return
		}

		switch typ := spec.Type.(type) {
		case *ast.InterfaceType:
			fl = typ.Methods
			srcFile = file
		case *ast.SelectorExpr:
			if pkgPath = file.Imports().Which(zcore.UnsafeBytes2String(file.Node(typ.X))); len(pkgPath) > 0 {
				fl, srcFile = getInterfaceFields(typ.Sel.Name, dir, pkgPath)
			}
			return filepath.SkipDir
		case *ast.Ident:
			fl, srcFile = getInterfaceFields(typ.Name, dir, pkgPath)
			return filepath.SkipDir
		}
		return
	})
	return
}

func (w Wire) generateSet(dir string, sets map[string]*wireDeclSet) (err error) {
	var (
		wireSets      = make([]WireSet, 0, len(sets))
		dstImports    = zcore.Imports{"github.com/google/wire": "wire"}
		dstImportPath = zcore.GetImportPath(dir)

		aopImports = make(zcore.Imports)
		aopSets    = make(map[string]*WireAop)
	)

	for set, decls := range sets {
		ws := WireSet{Name: set}

		for _, decl := range decls.keys {
			wd := decls.m[decl]
			el := WireSetElement{Path: zcore.GetImportPath(decl.File.Path), Name: decl.Name()}
			if len(el.Name) == 0 {
				continue
			}

			srcImports := decl.File.Imports()
			// fix name import package selector
			fp := func(name string) string {
				return zcore.FixPackage(name, el.Path, dstImportPath, srcImports, dstImports)
			}

			// has provider
			if len(wd.Provider) > 0 {
				el.Decls = append(el.Decls, fp(wd.Provider))
			}

			name := fp(el.Name)

			// add binds with aop
			for _, bind := range wd.Binds.Keys() {
				bindType := fp(bind)
				bindTemplate := wireStructBindTemplate

				switch decl.Type {
				case zcore.DeclTypeInterface:
					bindTemplate = `wire.Bind(new(%s), new(%s))`
				case zcore.DeclValue:
					bindTemplate = `wire.InterfaceValue(new(%s), %s)`
				}

				if _, aop := wd.Aops[bind]; !aop {
					// direct interface type binding
					zcore.Appendf(&el.Decls, bindTemplate, bindType, name)
					continue
				}

				// aop interface name
				aopTypename := "_aop_" + strings.Replace(bindType, ".", "_", -1)

				// register aop interface type
				aopType, ok := aopSets[aopTypename]
				if !ok {
					// aop interface type
					interfaceName := zcore.FixPackage(bind, el.Path, dstImportPath, srcImports, aopImports)

					// add aop generate entry
					aopType = &WireAop{
						Name:      aopTypename,
						Interface: interfaceName,
						Implement: "_impl" + aopTypename,
					}
					aopSets[aopTypename] = aopType
				}

				// aop type bindings
				zcore.Appendf(&el.Decls, bindTemplate, aopTypename, name)
				zcore.Appendf(&el.Decls, `wire.Struct(new(%s), "*")`, aopType.Implement)
				zcore.Appendf(&el.Decls, wireStructBindTemplate, bindType, aopType.Implement)
			}

			switch decl.Type {
			case zcore.DeclValue:
				if len(wd.Binds) == 0 {
					zcore.Appendf(&el.Decls, `wire.Value(%s)`, name)
				}

			case zcore.DeclFunc:
				el.Decls = append(el.Decls, name)

			case zcore.DeclTypeRefer:
				// struct refer type
				if wd.ReferStruct && len(wd.Provider) == 0 {
					zcore.Appendf(&el.Decls, `wire.Struct(new(%s), "*")`, name)
				}

			case zcore.DeclTypeStruct:
				// struct type
				if len(wd.Provider) == 0 {
					zcore.Appendf(&el.Decls, `wire.Struct(new(%s), "*")`, name)
				}

				// add fields of
				if fields := strings.Join(wd.Fields.Keys(), `","`); len(fields) > 0 {
					zcore.Appendf(&el.Decls, `wire.FieldsOf(new(%s), "%s")`, name, fields)
				}
			}

			if len(el.Decls) > 0 {
				ws.Elements = append(ws.Elements, el)
			}
		}

		if len(ws.Elements) > 0 {
			wireSets = append(wireSets, ws)
		}
	}

	if len(wireSets) == 0 {
		return
	}

	// package name
	pkg := zcore.GetImportName(dir)

	// sort by sets name
	sort.Slice(wireSets, func(i, j int) bool { return wireSets[i].Name < wireSets[j].Name })

	_ = os.Remove(filepath.Join(dir, wireAopFile))

	// render wire sets
	if err = zcore.RenderWrite(Wire{
		Imports: dstImports.List(),
		Sets:    wireSets,
	}, wireSetTemplate, filepath.Join(dir, wireSetFile), pkg, false, wireBuildFlag); err != nil {
		return
	}

	// write aop file
	return w.generateAops(dir, pkg, aopSets, aopImports)
}

func (w Wire) generateAops(dir, pkg string, sets map[string]*WireAop, aopImports zcore.Imports) (err error) {
	if len(sets) == 0 {
		return
	}

	wireAops := make([]WireAop, 0, len(sets))

	for _, e := range sets {
		e.Methods = w.parseInterfaceMethods(e.Interface, dir, aopImports)
		wireAops = append(wireAops, *e)
	}

	sort.Slice(wireAops, func(i, j int) bool { return wireAops[i].Name < wireAops[j].Name })

	return zcore.RenderWithDefaultTemplate(Wire{
		Imports: aopImports.List(),
		Aops:    wireAops,
	}, wireAopTemplate, filepath.Join(dir, wireAopFile), pkg, false)
}

func (w Wire) parseEntities(entities zcore.DeclEntities) map[string]*wireDeclSet {
	// parse entities set
	includes := make(map[string]map[int]struct{})
	excludes := make(map[int]map[string]struct{})

	for index, entity := range entities {
		for _, set := range strings.Split(entity.Options.Get("set", ""), ",") {
			if exclude, has := zcore.TrimPrefix(set, "!"); has {
				if excludes[index] == nil {
					excludes[index] = make(map[string]struct{})
				}
				excludes[index][exclude] = struct{}{}
			} else {
				if includes[set] == nil {
					includes[set] = make(map[int]struct{})
				}
				includes[set][index] = struct{}{}
			}
		}
	}

	for index, sets := range excludes {
		for set := range includes {
			if _, excluded := sets[set]; !excluded {
				includes[set][index] = struct{}{}
			}
		}
	}

	// group entities by set name
	groups := make(map[string]zcore.DeclEntities)
	for set, values := range includes {
		orders := make([]int, 0, len(values))
		for v := range values {
			orders = append(orders, v)
		}
		sort.Ints(orders)
		for _, index := range orders {
			groups[set] = append(groups[set], entities[index])
		}
	}

	// parse grouped entities to wire declaration set
	decls := make(map[string]*wireDeclSet, len(groups))
	for set, es := range groups {
		decls[set] = w.parseEntitiesDeclSet(es)
	}
	return decls
}
