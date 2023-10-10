package plugins

import (
	"bytes"
	"go/ast"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/Just-maple/gozz/zcore"
	"github.com/Just-maple/gozz/zutils"
)

func init() {
	zcore.RegisterPlugin(&Tag{})
}

type (
	Tag struct {
		ModifySet *zutils.ModifySet
		Tags      map[string]string
		FieldTags map[string]string
		Keys      []string
		Entity    zcore.DeclEntity
		Fields    zcore.FieldEntities
	}
)

func (t *Tag) Name() string { return "tag" }

func (t *Tag) Args() ([]string, map[string]string) { return []string{"tag", "template"}, nil }

func (t *Tag) Description() string { return "" }

func (t *Tag) Run(entities zcore.DeclEntities) (err error) {
	t.ModifySet = &zutils.ModifySet{}
	t.Tags = make(map[string]string)
	t.FieldTags = make(map[string]string)

	for _, t.Entity = range entities {
		if t.Entity.TypeSpec == nil {
			continue
		}

		t.Fields = t.Entity.ParseFields(2, nil)

		for k := range t.Tags {
			delete(t.Tags, k)
		}

		for _, tag := range strings.Split(t.Entity.Args[0], ",") {
			t.Tags[tag] = t.Entity.Args[1]
		}

		ast.Walk(t, t.Entity.TypeSpec.Type)
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
	t.ModifySet.Add(t.Entity.File.Path).Nodes[tag] = data
}

func (t *Tag) modifyField(field *ast.Field, name string) {
	t.reset()

	// parse fields tags extra annotations
	docs, annotations := zcore.ParseCommentGroup(zcore.AnnotationPrefix, field.Doc, field.Comment)
	for _, entity := range (&zcore.AnnotatedField{Annotations: annotations}).Parse(t.Entity.Plugin, 2, nil) {
		for _, key := range strings.Split(entity.Args[0], ",") {
			if tag, ok := zutils.TrimPrefix(key, "+"); ok {
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
	for key := range t.FieldTags {
		if len(key) == 0 {
			continue
		}

		// render tag string
		if str := (&strings.Builder{}); zcore.ExecuteTemplate(struct {
			FieldName string
			Docs      string
		}{FieldName: name, Docs: zutils.JoinDocs(docs)}, t.FieldTags[key], str) == nil && str.Len() > 0 {
			t.FieldTags[key] = str.String()
			t.Keys = append(t.Keys, key)
		}
	}

	if len(t.Keys) == 0 {
		return
	}

	sort.Strings(t.Keys)

	// check field tag exists
	if field.Tag == nil {
		bf := zutils.BuffPool.Get().(*bytes.Buffer)
		bf.Reset()
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
		if exist, ok := st.Lookup(key); !ok {
			// append
			str = strings.TrimSpace(str) + " " + tagKV(key, t.FieldTags[key])
			updated = true
		} else if exist != t.FieldTags[key] {
			// replace
			str = strings.Replace(str, tagKV(key, exist), tagKV(key, t.FieldTags[key]), 1)
			updated = true
		}
	}

	// add tag node modify
	if updated {
		bf := zutils.BuffPool.Get().(*bytes.Buffer)
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
