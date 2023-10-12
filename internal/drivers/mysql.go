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

package drivers

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stoewer/go-strcase"

	zcore "github.com/go-zing/gozz-core"
)

func init() { zcore.RegisterSchemaDriver(Mysql{}) }

type (
	Mysql struct{}

	Column struct {
		zcore.SqlColumn

		ColumnType    string
		ColumnComment string
		ColumnKey     *string
	}

	SliceMysqlColumn []Column
)

func (s *SliceMysqlColumn) Range(f func(interface{}, bool) bool) {
	for i := 0; ; i++ {
		if c := i >= len(*s); !c {
			if !f(&(*s)[i], c) {
				return
			}
		} else if n := append(*s, Column{}); f(&n[i], c) {
			*s = n
		} else {
			*s = n[:i]
			return
		}
	}
}

func (column *Column) FieldMapping() map[string]interface{} {
	m := column.SqlColumn.FieldMapping()
	m["column_type"] = &column.ColumnType
	m["column_key"] = &column.ColumnKey
	m["column_comment"] = &column.ColumnComment
	return m
}

func (m Mysql) Name() string { return "mysql" }

func (m Mysql) Parse(dsn, schema, table string, types map[string]string) (tables []zcore.OrmTable, err error) {
	db, err := sql.Open(m.Name(), dsn)
	if err != nil {
		return
	}
	defer db.Close()

	columns, err := m.queryColumns(db, schema, table)
	if err != nil {
		return
	}

	return m.parseTables(db, columns, types), nil
}

const (
	sqlStatementSelectColumns      = "SELECT `%s` FROM `information_schema`.`COLUMNS` WHERE `TABLE_SCHEMA` = ?"
	sqlStatementSelectTableComment = "SELECT `table_comment` FROM `information_schema`.`TABLES` WHERE `TABLE_NAME` = ? AND `TABLE_SCHEMA` = ?"
)

func (m Mysql) queryColumns(db *sql.DB, schema string, table string) (columns SliceMysqlColumn, err error) {
	fields := zcore.FieldsOf(&columns)
	args := []interface{}{schema}
	statement := &strings.Builder{}
	_, _ = fmt.Fprintf(statement, sqlStatementSelectColumns, strings.Join(fields, "`,`"))
	if tables := strings.Split(table, ","); table != "*" {
		statement.WriteString(" AND `TABLE_NAME` in (")
		for i, tb := range tables {
			if statement.WriteRune('?'); len(tables)-1 == i {
				statement.WriteRune(')')
			} else {
				statement.WriteRune(',')
			}
			args = append(args, tb)
		}
	}
	statement.WriteString(" ORDER BY `ordinal_position` ASC")

	rows, err := db.Query(statement.String(), args...)
	if err != nil {
		return
	}
	defer rows.Close()
	err = zcore.Scan(rows, fields, &columns)
	return
}

func (m Mysql) parseTables(db *sql.DB, columns []Column, types map[string]string) (tables []zcore.OrmTable) {
	stmt, _ := db.Prepare(sqlStatementSelectTableComment)
	if stmt != nil {
		defer stmt.Close()
	}

	tbs := make(map[string]int)

	for _, column := range columns {
		// init table
		index, ok := tbs[column.TableName]
		if !ok {
			tables = append(tables, zcore.OrmTable{
				Name:   strcase.UpperCamelCase(column.TableName),
				Table:  column.TableName,
				Schema: column.TableSchema,
			})
			index = len(tables) - 1
			tbs[column.TableName] = index

			// get table comment
			if stmt != nil {
				_ = stmt.QueryRow(column.TableName, column.TableSchema).Scan(&tables[index].Comment)
			}
		}

		table := &tables[index]

		// table primary key
		if column.ColumnKey != nil && *column.ColumnKey == "PRI" {
			table.Primary = column.ColumnName
		}

		c := zcore.OrmColumn{
			Name:    strcase.UpperCamelCase(column.ColumnName),
			Column:  column.ColumnName,
			Type:    column.ColumnType,
			Comment: column.ColumnComment,
			Ext:     column,
		}

		// max length
		if column.CharacterMaximumLength != nil {
			c.MaximumLength = *column.CharacterMaximumLength
		}

		// nullable
		c.Nullable = column.IsNullable == "YES"

		// golang types match
		matches := []string{column.ColumnType, column.DataType}

		if c.Nullable {
			matches = append([]string{"*" + column.ColumnType, "*" + column.DataType}, matches...)
		}

		// match first
		for _, key := range matches {
			if c.Type = types[key]; len(c.Type) > 0 {
				break
			}
		}

		table.Columns = append(table.Columns, c)
	}

	sort.Slice(tables, func(i, j int) bool { return tables[i].Name < tables[j].Name })
	return
}
