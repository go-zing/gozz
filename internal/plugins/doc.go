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
	_ "embed"
	"go/ast"
	"path/filepath"

	zcore "github.com/go-zing/gozz-core"
)

func init() {
	zcore.RegisterPlugin(Doc{})
}

//go:embed doc.go.tmpl
var docTemplate string

const (
	docDefaultFilename = "zzgen.doc.go"
)

type (
	DocType struct {
		Name   string
		Fields []DocField
		Data   bool
	}

	DocField struct {
		Name string
		Docs string
	}

	Doc struct {
		Types  []DocType
		Values []DocType
	}
)

func (d Doc) Name() string { return "doc" }

func (d Doc) Args() ([]string, map[string]string) { return nil, nil }

func (d Doc) Description() string {
	return "generate types and type fields (struct fields or interface methods) description text map from comments with template."
}

func (d Doc) Run(entities zcore.DeclEntities) (err error) {
	group := entities.GroupByDir()
	eg := new(zcore.ErrGroup)
	for key := range group {
		dir := key
		eg.Go(func() error { return d.GenDoc(dir, group[dir]) })
	}
	return eg.Wait()
}

func (d Doc) GenDoc(dir string, entities zcore.DeclEntities) (err error) {
	var (
		types  []DocType
		values []DocType
	)

	valuesMap := make(map[string]int)

	for _, entity := range entities {
		fields := make([]DocField, 0)

		if entity.TypeSpec != nil {
			if docs := zcore.JoinDocs(entity.Docs); len(docs) > 0 {
				fields = append(fields, DocField{Docs: docs})
			}

			data := true

			switch typ := entity.TypeSpec.Type.(type) {
			case *ast.InterfaceType:
				data = false
				fields = append(fields, parseFieldsDocs(typ.Methods)...)
			case *ast.StructType:
				fields = append(fields, parseFieldsDocs(typ.Fields)...)
			}

			if len(fields) > 0 {
				types = append(types, DocType{Name: entity.Name(), Fields: fields, Data: data})
			}
		}

		if entity.ValueSpec != nil {
			for _, name := range entity.ValueSpec.Names {
				fields = append(fields, DocField{
					Name: name.Name,
					Docs: zcore.JoinDocs(entity.Docs),
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

	if len(types)+len(values) == 0 {
		return
	}

	filename := filepath.Join(dir, docDefaultFilename)

	return zcore.RenderWithDefaultTemplate(&Doc{
		Types:  types,
		Values: values,
	}, docTemplate, filename, entities[0].File.Ast.Name.Name, false)
}

func parseFieldsDocs(fields *ast.FieldList) (fs []DocField) {
	for _, field := range fields.List {
		docs, _ := zcore.ParseCommentGroup(zcore.AnnotationPrefix, field.Doc, field.Comment)
		content := zcore.JoinDocs(docs)
		if len(content) == 0 {
			continue
		}
		for _, name := range field.Names {
			fs = append(fs, DocField{Name: name.String(), Docs: content})
		}
	}
	return
}
