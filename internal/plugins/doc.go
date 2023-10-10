package plugins

import (
	"go/ast"
	"path/filepath"

	"github.com/Just-maple/gozz/zcore"
	"github.com/Just-maple/gozz/zutils"
)

func init() {
	zcore.RegisterPlugin(Doc{})
}

type (
	DocType struct {
		Name   string
		Fields []DocField
	}

	DocField struct{ Name, Docs string }

	Doc struct{ DataTypes, Interfaces []DocType }
)

const (
	docDefaultFilename = "zzgen.doc.go"

	docTemplate = `var OpenapiDocMap = map[interface{}]map[string]string{
     {{ range .DataTypes }} (*{{ .Name }})(nil) : _map_{{ .Name }},
     {{ end }} {{ range .Interfaces }} (*{{ .Name }})(nil) : _map_{{ .Name }},
     {{ end }}
}

{{ range .DataTypes }}
var _map_{{ .Name }} = map[string]string{
	{{ range .Fields }} "{{ .Name }}" : {{ quote .Docs }},
	{{ end }}
}

func ({{ .Name }}) OpenapiDoc(f string) string { return _map_{{ .Name }}[f] }
{{ end }}


{{ range .Interfaces }}
var _map_{{ .Name }} = map[string]string{
	{{ range .Fields }} "{{ .Name }}" : {{ quote .Docs }},
	{{ end }}
}
{{ end }}
`
)

func (d Doc) Name() string { return "doc" }

func (d Doc) Args() ([]string, map[string]string) { return nil, nil }

func (d Doc) Description() string { return "" }

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
	)

	for _, entity := range entities {
		fields := make([]DocField, 0)
		types := &dataTypes

		if docs := zutils.JoinDocs(entity.Docs); len(docs) > 0 {
			fields = append(fields, DocField{Docs: docs})
		}

		switch typ := entity.TypeSpec.Type.(type) {
		case *ast.InterfaceType:
			types = &interfaces
			fields = append(fields, parseFieldsDocs(typ.Methods)...)
		case *ast.StructType:
			fields = append(fields, parseFieldsDocs(typ.Fields)...)
		}

		if len(fields) > 0 {
			*types = append(*types, DocType{Name: entity.Typename(), Fields: fields})
		}
	}

	if len(dataTypes)+len(interfaces) == 0 {
		return
	}

	filename := filepath.Join(dir, docDefaultFilename)

	return zcore.RenderWithDefaultTemplate(&Doc{
		DataTypes:  dataTypes,
		Interfaces: interfaces,
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
