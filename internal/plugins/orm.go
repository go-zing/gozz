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

	zcore "github.com/go-zing/gozz-core"

	_ "github.com/go-zing/gozz/internal/ormdrivers"
)

func init() {
	zcore.RegisterPlugin(Orm{})
}

type (
	Orm struct {
		Tables []zcore.OrmTable
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

{{ range .Tables }} // {{ .Schema }}.{{ .Table }} {{ if .Comment }} 
{{ comment .Comment }} {{ end }}
const Table{{ .Name }} = "{{ .Table }}"

type {{ .Name }} struct{ {{ range .Columns }} 
	// {{ .Column }} : {{ if .Nullable }}NULLABLE {{ end }}{{ if .Ext }}{{ .Ext.ColumnType }} {{ end }}{{ if .Comment }}
    {{ comment .Comment }} {{ end }}
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
	return []string{
			"schema:specify databases schema to load tables",
		},
		map[string]string{
			"filename": "specify which file to generate orm types and template files. default: ./zzgen.orm.go",
			"driver": fmt.Sprintf("specify databases schema driver. default: mysql. available drivers: [ %s ]",
				strings.Join(zcore.GetOrmSchemaDrivers(), ",")),
			"type":     `specify database schema datatype binding to golang typing. example: varchar=string. add "*" prefix for nullable type. example: [ *timestamp=*time.Time ]`,
			"table":    `specify table names to load. default: * (load all tables).use "," to split if multi. example: [ table=user,book,order ]`,
			"user":     "user in sql default format dsn. default: root",
			"host":     "host in sql default format dsn. default: localhost",
			"port":     "port in sql default format dns. default: 3306",
			"password": "password in default format sql dsn",
			"dsn":      "specify sql dsn to load schema. other options to format dsn would be ignored if provide. default: [ ${user}:${password}@tpc(${host}:${port})/ ]",
		}
}

func (o Orm) Description() string {
	return "generate type struct field relation mapping from databases schema."
}

func (o Orm) Run(entities zcore.DeclEntities) (err error) {
	group, err := o.group(entities)
	if err != nil {
		return
	}

	for filename, tables := range group {
		for i := range tables {
			table := &tables[i]
			for ci := range table.Columns {
				column := &table.Columns[ci]
				if len(column.Type) == 0 {
					column.Type = "interface{}"
				}
			}
		}
		if err = zcore.RenderWithDefaultTemplate(Orm{Tables: tables},
			ormTemplateText, filename, zcore.GetImportName(filename), false); err != nil {
			return
		}
	}
	return
}

func (o Orm) parseTables(entity zcore.DeclEntity) (tables []zcore.OrmTable, err error) {
	opt := entity.Options

	// get driver. default mysql
	driverName := opt.Get("driver", "mysql")
	driver := zcore.GetOrmSchemaDriver(driverName)
	if driver == nil {
		return nil, errors.New("unregister driver: " + driverName)
	}

	// default types
	types := zcore.OrmTypeMapping()

	// commands or annotations defined types
	// extract types from options
	extTypes := make(map[string]string)
	zcore.SplitKVSlice2Map(strings.Split(entity.Options["type"], ","), "=", extTypes)
	for key, value := range extTypes {
		types[key] = value
	}

	dsn := opt.Get("dsn", "")
	if len(dsn) == 0 {
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/",
			opt.Get("user", "root"),
			opt.Get("password", ""),
			opt.Get("host", "localhost"),
			opt.Get("port", "3306"),
		)
	}

	// parse dsn and get tables
	return driver.Parse(dsn, entity.Args[1], opt.Get("table", "*"), types)
}

func (o Orm) group(entities zcore.DeclEntities) (map[string][]zcore.OrmTable, error) {
	group := make(map[string][]zcore.OrmTable)
	for _, entity := range entities {
		tables, e := o.parseTables(entity)
		if e != nil {
			return nil, e
		}
		filename := entity.RelFilename(entity.Options.Get("filename", "./"), "zzgen.orm.go")
		group[filename] = append(group[filename], tables...)
	}
	return group, nil
}
