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
	"errors"
	"fmt"
	"strings"

	"github.com/Just-maple/gozz/zcore"
	"github.com/Just-maple/gozz/zorm"
	"github.com/Just-maple/gozz/zutils"

	_ "github.com/Just-maple/gozz/zorm/drivers"
)

func init() {
	zcore.RegisterPlugin(Orm{})
}

type (
	Orm struct {
		Tables []zorm.Table
	}
)

const (
	ormTemplateText = `import (
	"context"
	"time"
	"encoding/json"
	"database/sql"
)

var (
	_ = (*context.Context)(nil)
    _ = (*json.RawMessage)(nil)
    _ = (*time.Time)(nil)
    _ = (*sql.NullString)(nil)
)

var tables = []interface{}{
{{ range .Tables }} {{ .Name }}{},
{{ end }} }

{{ range .Tables }}
/*
{{ .Schema }}.{{ .Table }} {{ if .Comment }} 

{{ .Comment }} {{ end }}
*/
const Table{{ .Name }} = "{{ .Table }}"

type {{ .Name }} struct{ {{ range .Columns }} 
	// {{ .Column }} {{ if .Comment }} 
    // {{ .Comment }} {{ end }}
    {{ .Name }} {{ .Type }} {{ end }} 
}

func ({{ .Name }}) TableName() string { return Table{{ .Name }} } 

func (m *{{ .Name }}) FieldMapping() (map[string]interface{}){
    return map[string]interface{}{ 
		{{ range .Columns }} {{ quote .Column }}: &m.{{ .Name }},
		{{ end }} }
}

type Slice{{ .Name }} []{{ .Name }}

func (s *Slice{{ .Name }}) Range(f func(interface{}, bool) bool) {
	for i := 0; ; i++ {
		if c := i >= len(*s); !c {
			if !f(&(*s)[i], c) {
				return
			}
		} else if n := append(*s, {{ .Name }}{}); f(&n[i], c) {
			*s = n
		} else {
			*s = n[:i]
			return
		}
	}
}

{{ end }}
`
)

func (o Orm) Name() string { return "orm" }

func (o Orm) Args() ([]string, map[string]string) {
	return []string{"schema", "filename"}, nil
}

func (o Orm) Description() string { return "" }

func (o Orm) Run(entities zcore.DeclEntities) (err error) {
	group, err := o.group(entities)
	if err != nil {
		return
	}

	eg := zutils.ErrGroup{}
	for key := range group {
		filename := key
		pkg := zutils.GetImportName(filename)
		eg.Go(func() error {
			return zcore.RenderWithDefaultTemplate(Orm{Tables: group[filename]}, ormTemplateText, filename, pkg, false)
		})
	}
	return eg.Wait()
}

func (o Orm) parseTables(entity zcore.DeclEntity) (tables []zorm.Table, err error) {
	opt := entity.Options

	// driver
	driverName := opt.Get("driver", "mysql")
	driver := zorm.Get(driverName)
	if driver == nil {
		return nil, errors.New("unregister driver: " + driverName)
	}

	types := zorm.DefaultTypes()

	zutils.SplitKVSlice2Map(strings.Split(entity.Options["types"], ","), "=", types)

	return driver.Parse(fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		opt.Get("user", "root"),
		opt.Get("password", ""),
		opt.Get("host", "localhost"),
		opt.Get("port", "3306"),
	), entity.Args[0], opt.Get("table", "*"), types)
}

func (o Orm) group(entities zcore.DeclEntities) (group map[string][]zorm.Table, err error) {
	group = make(map[string][]zorm.Table)

	for _, entity := range entities {
		var tables []zorm.Table
		if tables, err = o.parseTables(entity); err != nil {
			return
		}
		filename := entity.RelFilename(entity.Args[1], "zzgen.orm.go")

		for i := range tables {
			table := &tables[i]
			for ci := range table.Columns {
				column := &table.Columns[ci]
				column.Comment = strings.Replace(column.Comment, "\n", "\n// ", -1)
			}
		}
		group[filename] = append(group[filename], tables...)
	}
	return
}
