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
	"go/ast"
	"path/filepath"

	"github.com/go-zing/gozz/zcore"
	"github.com/go-zing/gozz/zutils"
)

func init() {
	zcore.RegisterPlugin(Doc{})
}

type (
	DocType struct {
		Name   string
		Fields []DocField
	}

	DocField struct {
		Name string
		Docs string
	}

	Doc struct {
		DataTypes  []DocType
		Interfaces []DocType
		Values     []DocType
	}
)

const (
	docDefaultFilename = "zzgen.doc.go"

	docTemplate = `var (
	_types_doc = map[interface{}]map[string]string{
	 	{{ range .Interfaces }} (*{{ .Name }})(nil) : _doc_{{ .Name }},
     	{{ end }} 
     	{{ range .DataTypes }} (*{{ .Name }})(nil) : _doc_{{ .Name }},
     	{{ end }} 
	}
	
	_values_doc = map[string]map[interface{}]string{
		{{ range .Values }} {{ quote .Name }} : map[interface{}]string{
			{{ range .Fields }} {{ .Name }} : {{ quote .Docs }},
			{{ end }}
		},
		{{ end }}
	}
)

{{ range .Interfaces }}
var _doc_{{ .Name }} = map[string]string{
	{{ range .Fields }} "{{ .Name }}" : {{ quote .Docs }},
	{{ end }}
}
{{ end }}

{{ range .DataTypes }}
var _doc_{{ .Name }} = map[string]string{
	{{ range .Fields }} "{{ .Name }}" : {{ quote .Docs }},
	{{ end }}
}

func ({{ .Name }}) FieldDoc(f string) string { return _doc_{{ .Name }}[f] }
{{ end }}

`
)

func (d Doc) Name() string { return "doc" }

func (d Doc) Args() ([]string, map[string]string) { return nil, nil }

func (d Doc) Description() string {
	return "generate types and type fields (struct fields or interface methods) description text map from comments with template."
}

func (d Doc) Run(entities zcore.DeclEntities) (err error) {
	group := entities.GroupByDir()
	eg := new(zutils.ErrGroup)
	for key := range group {
		dir := key
		eg.Go(func() error { return d.GenDoc(dir, group[dir]) })
	}
	return eg.Wait()
}

func (d Doc) GenDoc(dir string, entities zcore.DeclEntities) (err error) {
	var (
		dataTypes  []DocType
		interfaces []DocType
		values     []DocType
	)

	valuesMap := make(map[string]int)

	for _, entity := range entities {
		fields := make([]DocField, 0)

		if entity.TypeSpec != nil {
			if docs := zutils.JoinDocs(entity.Docs); len(docs) > 0 {
				fields = append(fields, DocField{Docs: docs})
			}

			types := &dataTypes
			switch typ := entity.TypeSpec.Type.(type) {
			case *ast.InterfaceType:
				types = &interfaces
				fields = append(fields, parseFieldsDocs(typ.Methods)...)
			case *ast.StructType:
				fields = append(fields, parseFieldsDocs(typ.Fields)...)
			}

			if len(fields) > 0 {
				*types = append(*types, DocType{Name: entity.Name(), Fields: fields})
			}
		}

		if entity.ValueSpec != nil {
			for _, name := range entity.ValueSpec.Names {
				fields = append(fields, DocField{
					Name: name.Name,
					Docs: zutils.JoinDocs(entity.Docs),
				})
			}

			label := entity.Options.Get("label", "")
			if _, ok := valuesMap[label]; !ok {
				values = append(values, DocType{
					Name: label,
				})
				valuesMap[label] = len(values) - 1
			}

			values[valuesMap[label]].Fields = append(values[valuesMap[label]].Fields, fields...)
		}
	}

	if len(dataTypes)+len(interfaces)+len(values) == 0 {
		return
	}

	filename := filepath.Join(dir, docDefaultFilename)

	return zcore.RenderWithDefaultTemplate(&Doc{
		DataTypes:  dataTypes,
		Interfaces: interfaces,
		Values:     values,
	}, docTemplate, filename, entities[0].File.Ast.Name.Name, false)
}

func parseFieldsDocs(fields *ast.FieldList) (fs []DocField) {
	for _, field := range fields.List {
		docs, _ := zcore.ParseCommentGroup(zcore.AnnotationPrefix, field.Doc, field.Comment)
		content := zutils.JoinDocs(docs)
		if len(content) == 0 {
			continue
		}
		for _, name := range field.Names {
			fs = append(fs, DocField{Name: name.String(), Docs: content})
		}
	}
	return
}
