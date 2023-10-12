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
	"reflect"
	"sort"
	"strconv"
	"strings"

	zcore "github.com/go-zing/gozz-core"
)

func init() {
	zcore.RegisterPlugin(&Tag{})
}

type (
	Tag struct {
		ModifySet *zcore.ModifySet
		Tags      map[string]string
		FieldTags map[string]string
		Keys      []string
		Decl      *zcore.AnnotatedDecl
	}
)

func (t *Tag) Name() string { return "tag" }

func (t *Tag) Args() ([]string, map[string]string) {
	return []string{
		`tag:specify tag key on field. use "," to separate multiple keys. use "+" prefix for additional prefix on field.`,
		"format:template format of generate tag value. using golang template syntax. example: [ json:{{ .FieldName }} ]",
	}, nil
}

func (t *Tag) Description() string {
	return "add or modify struct field tags from format templates."
}

func (t *Tag) Run(entities zcore.DeclEntities) (err error) {
	group := make(map[*zcore.AnnotatedDecl]zcore.DeclEntities)

	for _, entity := range entities {
		if entity.TypeSpec == nil {
			continue
		}
		group[entity.AnnotatedDecl] = append(group[entity.AnnotatedDecl], entity)
	}

	t.ModifySet = &zcore.ModifySet{}
	t.FieldTags = make(map[string]string)
	t.Tags = make(map[string]string)

	for t.Decl, entities = range group {
		for _, entity := range entities {
			for _, tag := range strings.Split(entity.Args[0], ",") {
				t.Tags[tag] = entity.Args[1]
			}
		}
		ast.Walk(t, t.Decl.TypeSpec.Type)
		for k := range t.Tags {
			delete(t.Tags, k)
		}
	}

	return t.ModifySet.Apply()
}

func (t *Tag) reset() {
	for k, v := range t.Tags {
		t.FieldTags[k] = v
	}
	for k := range t.FieldTags {
		if _, exist := t.Tags[k]; !exist {
			delete(t.FieldTags, k)
		}
	}
	t.Keys = t.Keys[:0]
}

func tagKV(key, v string) string { return key + ":" + strconv.Quote(v) }

func (t *Tag) modifyNode(tag *ast.BasicLit, data []byte) {
	t.ModifySet.Add(t.Decl.File.Path).Nodes[tag] = data
}

func (t *Tag) modifyField(field *ast.Field, name string) {
	t.reset()

	// parse fields tags extra annotations
	docs, annotations := zcore.ParseCommentGroup(zcore.AnnotationPrefix, field.Doc, field.Comment)
	for _, entity := range (&zcore.AnnotatedField{Annotations: annotations}).Parse(t.Name(), 2, nil) {
		for _, key := range strings.Split(entity.Args[0], ",") {
			if tag, ok := zcore.TrimPrefix(key, "+"); ok {
				// if field annotation starts with "+"
				// $field_tag = $struct_tag + $field_tag
				//
				// example:
				// // +zz:tag:json:{{ snake .FieldName }}
				// type User struct{
				//     // field without annotations
				//     Name string `json:"name"
				//     // field with extra annotations
				//     // +zz:tag:+json:,omitempty
				//     Address string `json:"address,omitempty"
				// }
				t.FieldTags[tag] += entity.Args[1]
			} else {
				t.FieldTags[tag] = entity.Args[1]
			}
		}
	}

	// render fields
	for key, value := range t.FieldTags {
		if len(key) == 0 {
			continue
		}

		if strings.Contains(value, "{{") && strings.Contains(value, "}}") {
			// render tag template string
			if str := (&strings.Builder{}); zcore.ExecuteTemplate(struct {
				FieldName string
				Docs      string
			}{
				FieldName: name,
				Docs:      zcore.JoinDocs(docs),
			}, value, str) == nil {
				value = str.String()
				t.FieldTags[key] = value
			}
		}

		if len(value) > 0 {
			t.Keys = append(t.Keys, key)
		}
	}

	if len(t.Keys) == 0 {
		return
	}

	sort.Strings(t.Keys)

	// check field tag exists
	if field.Tag == nil {
		bf := zcore.BuffPool.Get().(*bytes.Buffer)
		bf.Reset()

		// write tags
		bf.WriteString(" `")
		for i, key := range t.Keys {
			if i > 0 {
				bf.WriteRune(' ')
			}
			bf.WriteString(tagKV(key, t.FieldTags[key]))
		}
		bf.WriteRune('`')
		t.modifyNode(&ast.BasicLit{ValuePos: field.Type.End()}, bf.Bytes())
		return
	}

	// parse existing tag
	str, err := strconv.Unquote(field.Tag.Value)
	if err != nil {
		return
	}

	st := reflect.StructTag(str)
	updated := false

	// replace existing keys or append
	for _, key := range t.Keys {
		value := t.FieldTags[key]
		if exist, ok := st.Lookup(key); !ok {
			// append
			str = strings.TrimSpace(str) + " " + tagKV(key, value)
			updated = true
		} else if exist != value {
			// replace
			str = strings.Replace(str, tagKV(key, exist), tagKV(key, value), 1)
			updated = true
		}
	}

	// add tag node modify
	if updated {
		bf := zcore.BuffPool.Get().(*bytes.Buffer)
		bf.Reset()
		bf.WriteRune('`')
		bf.WriteString(strings.TrimSpace(str))
		bf.WriteRune('`')
		t.modifyNode(field.Tag, bf.Bytes())
	}
}

func (t *Tag) visitFieldList(fl *ast.FieldList) {
	if fl != nil {
		for _, r := range fl.List {
			ast.Walk(t, r.Type)
		}
	}
}

func (t *Tag) Visit(n ast.Node) ast.Visitor {
	switch node := n.(type) {
	default:
		return t
	case *ast.InterfaceType:
		t.visitFieldList(node.Methods)
	case *ast.FuncType:
		t.visitFieldList(node.Params)
		t.visitFieldList(node.Results)
	case *ast.Field:
		if len(node.Names) != 1 {
			ast.Walk(t, node.Type)
		} else if name := node.Names[0]; name.IsExported() {
			t.modifyField(node, name.Name)
			return t
		}
	}
	return nil
}
