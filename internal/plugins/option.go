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
	"path/filepath"
	"strconv"
	"strings"

	zcore "github.com/go-zing/gozz-core"
)

func init() {
	zcore.RegisterPlugin(Option{})
}

type (
	Option struct {
		Imports []zcore.Import
		Types   []optionType
	}

	optionType struct {
		Typename string
		Name     string
		Fields   []optionField
	}

	optionField struct {
		Func string
		Name string
		Type string
		Doc  string
	}
)

func (ot optionType) Type() string {
	if len(ot.Typename) > 0 {
		return ot.Typename
	} else {
		return fmt.Sprintf(`func(*%s)`, ot.Name)
	}
}

func (o Option) Name() string { return "option" }

func (o Option) Args() ([]string, map[string]string) {
	return nil, map[string]string{"type": "define a function type to replace func(*{{ .Name }})"}
}

func (o Option) Description() string {
	return "generate functional options from option struct."
}

const optionTemplate = `import  (
	{{ range .Imports }} {{ .Name }} "{{ .Path }}"
	{{ end }}
)

{{ range .Types }} {{ $n  := .Name }} {{ $t := .Type }} {{ if .Typename }} 
// functional options type for *{{ $n }}
type {{ .Typename }} func(*{{ $n }})
{{ end }}
// apply functional options for *{{ $n }}
func (o *{{ $n }}) applyOptions(opts ...{{ $t }}){ for _,opt :=range opts{	opt(o) } }

{{ range .Fields }} {{ if .Doc }}
{{ comment .Doc }} {{ end }}
func {{ .Func }}(v {{ .Type }}) {{ $t }} { return func(o *{{ $n }}){ o.{{ .Name }} = v	} }
{{ end }} {{ end }}
`

func (o Option) Run(es zcore.DeclEntities) (err error) {
	for dir, entities := range es.GroupByDir() {
		if err = o.run(dir, entities); err != nil {
			return
		}
	}
	return
}

func (o Option) run(dir string, entities zcore.DeclEntities) (err error) {
	types := make([]optionType, 0)
	filename := filepath.Join(dir, "zzgen.option.go")
	imports := make(zcore.Imports)
	group := make(map[*zcore.AnnotatedDecl][]int)
	decls := make([]*zcore.AnnotatedDecl, 0, len(group))
	named := make(map[string]int)
	pkg := ""

	for i, entity := range entities {
		if len(pkg) == 0 {
			pkg = entity.File.Ast.Name.Name
		}
		if entity.Type != zcore.DeclTypeStruct {
			continue
		}
		if entity.TypeSpec.Type.(*ast.StructType).Fields == nil {
			continue
		}
		if len(group[entity.AnnotatedDecl]) == 0 {
			decls = append(decls, entity.AnnotatedDecl)
		}
		group[entity.AnnotatedDecl] = append(group[entity.AnnotatedDecl], i)
	}

	for _, decl := range decls {
		option := optionType{Name: decl.Name()}
		for _, index := range group[decl] {
			entity := entities[index]
			if typename := entities[index].Options.Get("type", ""); len(typename) > 0 {
				if str := (&strings.Builder{}); zcore.ExecuteTemplate(entity, typename, str) == nil {
					option.Typename = str.String()
				}
			}
		}

		for _, field := range decl.TypeSpec.Type.(*ast.StructType).Fields.List {
			docs, _ := zcore.ParseCommentGroup(zcore.AnnotationPrefix, field.Doc, field.Comment)
			doc := zcore.JoinDocs(docs)
			typ := zcore.UnsafeBytes2String(decl.File.ReplacePackages(field.Type, filename, imports))
			for _, name := range field.Names {
				of := optionField{
					Doc:  doc,
					Name: name.Name,
					Func: "With" + name.Name,
					Type: typ,
				}
				named[of.Func] += 1
				if dup := named[of.Func]; dup > 1 {
					of.Func += strconv.Itoa(dup)
				}
				option.Fields = append(option.Fields, of)
			}
		}

		if len(option.Fields) > 0 {
			types = append(types, option)
		}
	}

	if len(types) == 0 {
		return
	}

	return zcore.RenderWithDefaultTemplate(Option{Types: types, Imports: imports.List()}, optionTemplate, filename, pkg, false)
}
