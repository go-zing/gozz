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
	"errors"
	"strings"

	zcore "github.com/go-zing/gozz-core"

	_ "github.com/go-zing/gozz/internal/ormdrivers"
)

//go:embed orm.go.tmpl
var ormTemplateText string

func init() {
	zcore.RegisterPlugin(Orm{})
}

type (
	Orm struct {
		Tables []zcore.OrmTable
	}
)

func (o Orm) Name() string { return "orm" }

func (o Orm) Args() ([]string, map[string]string) {
	return []string{
			"schema:specify databases schema to load tables",
		},
		map[string]string{
			"filename": "specify which file to generate orm types and template files. default: ./zzgen.orm.go",
			"driver":   "specify databases schema driver. default: mysql.",
			"type":     `specify database schema datatype binding to golang typing. example: varchar=string. add "*" prefix for nullable type. example: [ *timestamp=*time.Time ]`,
			"table":    `specify table names to load. default: * (load all tables).use "," to split if multi. example: [ table=user,book,order ]`,
			"dsn":      "specify sql dsn to load driver schema.",
			"password": "specify password for default dsn from driver provided.",
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

	if typeStr := entity.Options.Get("type", ""); len(typeStr) > 0 {
		zcore.SplitKVSlice2Map(strings.Split(typeStr, ","), "=", extTypes)
	}

	for key, value := range extTypes {
		types[key] = value
	}

	dsn := opt.Get("dsn", "")

	if len(dsn) == 0 {
		if dsn = driver.Dsn(opt.Get("password", "")); len(dsn) == 0 {
			err = errors.New("invalid schema driver dsn")
			return
		}
	}

	schemas := entity.Args[0]
	zcore.TryExecuteTemplate(entity, schemas, &schemas)

	var tmp []zcore.OrmTable
	// parse dsn and get tables
	for _, schema := range strings.Split(schemas, ",") {
		if schema = strings.TrimSpace(schema); len(schema) == 0 {
			continue
		}
		if tmp, err = driver.Parse(dsn, schema, opt.Get("table", "*"), types, opt); err != nil {
			return
		}
		tables = append(tables, tmp...)
	}
	return
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
