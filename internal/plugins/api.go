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
	"sort"
	"strings"

	zcore "github.com/go-zing/gozz-core"
)

func init() {
	zcore.RegisterPlugin(Api{})
}

type (
	Api struct {
		Imports    []zcore.Import
		Interfaces []ApiInterface
	}

	ApiHandler struct {
		Name     string
		Method   string
		Resource string
		Options  map[string]string
		Invoke   string
	}

	ApiInterface struct {
		Package  string
		Type     string
		Handlers []ApiHandler
	}
)

func (i ApiInterface) FieldName() string {
	s := i.Type
	if len(i.Package) > 0 {
		s = i.Package + "_" + s
	}
	return strings.Title(s)
}

func (i ApiInterface) TypeName() string {
	if len(i.Package) > 0 {
		return i.Package + "." + i.Type
	}
	return i.Type
}

const (
	apiDefaultFilename = "zzgen.api.go"

	apiTemplateText = `import (
	{{ range .Imports }} {{ .Name }} "{{ .Path }}" 
	{{ end }}
)

var _  = context.Context(nil)

type Apis struct {
    {{ range .Interfaces }} {{ .FieldName }} {{ .TypeName }}
    {{ end }}
}

func (s Apis) Range(fn func(interface{},[]map[string]interface{})){
	for _,f:=range []func()(interface{},[]map[string]interface{}){
		{{ range .Interfaces }} s._{{ .FieldName }},
		{{ end }}
	}{
		fn(f())
	}
}

{{ range .Interfaces }} 
func (s Apis) _{{ .FieldName }}() (interface{},[]map[string]interface{}){
	t := s.{{ .FieldName }}
	return &t,[]map[string]interface{}{	{{ range  .Handlers }}
		{
			"name": "{{ .Name }}",
			"method": "{{ .Method }}",
			"resource": "{{ .Resource }}",
			"options": map[string]string{ {{ range $key,$value := .Options }}
			{{ quote $key }} : {{ quote $value }}, {{ end }}
			},  {{ if .Invoke }} 
			"invoke": {{ .Invoke }}, {{ end }}
		},	{{ end }}
	}
}
{{ end }}
`

	paramDecode        = "var in {{ .Param }};if err:=dec(&in);err!=nil{return nil,err};"
	invokeBaseTemplate = `func(ctx context.Context, dec func(interface{}) error) (interface{},error) {%s%s t.{{ .Name }}(%s)%s}`
)

func (a Api) Name() string { return "api" }

func (a Api) Args() ([]string, map[string]string) {
	return []string{"filename:specify filename to generate api stubs and template files"}, nil
}

func (a Api) Description() string {
	return "generate api specification stubs from interface declarations with template."
}

func (a Api) Run(entities zcore.DeclEntities) (err error) {
	group, err := a.group(entities)
	if err != nil {
		return
	}
	eg := new(zcore.ErrGroup)
	for key := range group {
		filename := key
		eg.Go(func() error { return a.generateApi(filename, group[filename]) })
	}
	return eg.Wait()
}

func (a Api) group(entities zcore.DeclEntities) (group map[string]map[*ast.TypeSpec]zcore.FieldEntities, err error) {
	group = make(map[string]map[*ast.TypeSpec]zcore.FieldEntities)

	for _, entity := range entities {
		if entity.Type != zcore.DeclTypeInterface {
			continue
		}

		filename := entity.RelFilename(entity.Args[0], apiDefaultFilename)

		if group[filename] == nil {
			group[filename] = make(map[*ast.TypeSpec]zcore.FieldEntities)
		}

		group[filename][entity.TypeSpec] = append(group[filename][entity.TypeSpec], entity.ParseFields(2, entity.Options)...)
	}
	return
}

func (Api) generateApi(filename string, typeMap map[*ast.TypeSpec]zcore.FieldEntities) (err error) {
	var (
		imports       = zcore.Imports{"context": "context"}
		dstImportPath = zcore.GetImportPath(filename)
		interfaces    = make([]ApiInterface, 0)
	)

	for typ, fields := range typeMap {
		api := ApiInterface{Type: typ.Name.Name}

		for i, field := range fields {
			if i == 0 {
				if srcImportPath := zcore.GetImportPath(field.Decl.File.Path); srcImportPath != dstImportPath {
					api.Package = imports.Add(srcImportPath)
				}
			}

			funcName, ft, ok := zcore.AssertFuncType(field.Field)
			if !ok {
				continue
			}

			handler := ApiHandler{
				Name:     funcName,
				Method:   field.Args[0],
				Resource: field.Args[1],
				Options:  field.Options,
			}

			// try parse method invoke function
			if pt, tmpl := (*funcType)(ft).InvokeTemplate(); len(tmpl) > 0 {
				// render invoke template
				if str := (&strings.Builder{}); zcore.ExecuteTemplate(struct{ Name, Param string }{
					Name:  funcName,
					Param: zcore.UnsafeBytes2String(field.Decl.File.ReplacePackages(pt, filename, imports)),
				}, tmpl, str) == nil {
					handler.Invoke = str.String()
				}
			}

			api.Handlers = append(api.Handlers, handler)
		}

		if len(api.Handlers) > 0 {
			interfaces = append(interfaces, api)
		}
	}

	if len(interfaces) == 0 {
		return
	}

	sort.Slice(interfaces, func(i, j int) bool {
		return interfaces[i].Package+"."+interfaces[i].Type < interfaces[j].Package+"."+interfaces[j].Type
	})

	return zcore.RenderWithDefaultTemplate(&Api{
		Imports:    imports.List(),
		Interfaces: interfaces,
	}, apiTemplateText, filename, zcore.GetImportName(filename), false)
}

type funcType ast.FuncType

func (ft *funcType) InvokeTemplate() (paramType ast.Node, template string) {
	// params/results values count must equal to params/results types count
	// because valid multi params/results must in different types
	// example: func(context.Context, types.Query) (types.Response, error)
	if (ft.Params != nil && ft.Params.NumFields() != len(ft.Params.List)) ||
		(ft.Results != nil && ft.Results.NumFields() != len(ft.Results.List)) {
		return
	}

	// alloc param / call decode / invoke method with params
	param, decode, paramType, ok := ft.params()
	if !ok {
		return
	}

	// assign result values / return values
	ret, ret2, ok := ft.returns()
	if !ok {
		return
	}

	template = fmt.Sprintf(invokeBaseTemplate, decode, ret, param, ret2)
	return
}

func (ft *funcType) returns() (ret, ret2 string, valid bool) {
	isError := func(e ast.Expr) bool { i, ok := e.(*ast.Ident); return ok && i.Name == "error" }

	// no results value
	// func (...)
	if ft.Results == nil || ft.Results.NumFields() == 0 {
		ret2 = ";return nil,nil"
		valid = true
		return
	}

	switch len(ft.Results.List) {
	case 1:
		if isError(ft.Results.List[0].Type) {
			// func (...) (error)
			ret = "return nil,"
		} else {
			// func (...) (types.Response)
			ret = "return"
			ret2 = ",nil"
		}
	case 2:
		if !isError(ft.Results.List[1].Type) {
			// invalid second results type
			// func (...) (types.Response, types.Other)
			return
		}

		// func (...) (types.Response, error)
		ret = "return"
	default:
		// more than 2 results
		// func (...) (types.Response, types.Response2, error)
		return
	}

	valid = true
	return
}

func (ft *funcType) params() (param, decode string, paramType ast.Node, valid bool) {
	isContext := func(e ast.Expr) bool {
		s, ok := e.(*ast.SelectorExpr)
		return ok && fmt.Sprintf("%s.%s", s.X, s.Sel) == "context.Context"
	}

	// no params
	// func () ...
	if ft.Params == nil || ft.Params.NumFields() == 0 {
		valid = true
		return
	}

	switch len(ft.Params.List) {
	case 1:
		if isContext(ft.Params.List[0].Type) {
			// func (context.Context) ...
			param = "ctx"
		} else {
			// func (types.Param) ...
			decode = paramDecode
			paramType = ft.Params.List[0].Type
			if se, ok := paramType.(*ast.StarExpr); ok {
				paramType = se.X
				param = "&in"
			} else {
				param = "in"
			}
		}
	case 2:
		if !isContext(ft.Params.List[0].Type) {
			// invalid first params type
			// func (types.Other, types.Query) ...
			return
		}

		// func (context.Context, types.Param) ...
		decode = paramDecode
		paramType = ft.Params.List[1].Type
		if se, ok := paramType.(*ast.StarExpr); ok {
			paramType = se.X
			param = "ctx,&in"
		} else {
			param = "ctx,in"
		}
	default:
		// more than 2 params
		// func (context.Context, types.Param, types.Param2) ...
		return
	}

	valid = true
	return
}
