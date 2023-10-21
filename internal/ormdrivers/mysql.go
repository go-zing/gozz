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

package ormdrivers

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	zcore "github.com/go-zing/gozz-core"
)

func init() { zcore.RegisterOrmSchemaDriver(Mysql{}) }

// +zz:orm:information_schema:table=COLUMNS
type Mysql struct{}

func (m Mysql) Name() string { return "mysql" }

func (m Mysql) Dsn(password string) (dsn string) {
	return fmt.Sprintf("root:%s@tcp(localhost:3306)/", password)
}

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

func getColumnKeys(columns SliceColumns) (fields []string) {
	mapping := make(map[string]interface{})
	zcore.IterateOrmFieldMapper(&columns, func(m zcore.OrmFieldMapper, b bool) bool {
		m.FieldMapping(mapping)
		fields = make([]string, 0, len(mapping))
		for key := range mapping {
			fields = append(fields, key)
		}
		return false
	})
	sort.Strings(fields)
	return
}

func (m Mysql) queryColumns(db *sql.DB, schema string, table string) (columns SliceColumns, err error) {
	fields := getColumnKeys(columns)
	args := []interface{}{schema}
	statement := &strings.Builder{}

	// build query sql
	_, _ = fmt.Fprintf(statement, sqlStatementSelectColumns, strings.Join(fields, "`,`"))
	if tables := strings.Split(table, ","); table != "*" {
		statement.WriteString(" AND `TABLE_NAME` in (")
		for i, tb := range tables {
			statement.WriteRune('?')
			if len(tables)-1 != i {
				statement.WriteRune(',')
			} else {
				statement.WriteRune(')')
			}
			args = append(args, tb)
		}
	}
	statement.WriteString(" ORDER BY `ordinal_position` ASC")

	// do query
	rows, err := db.Query(statement.String(), args...)
	if err != nil {
		return
	}
	defer rows.Close()

	// scan rows values
	err = zcore.ScanSqlRows(rows, fields, &columns)
	return
}

func (m Mysql) parseTables(db *sql.DB, columns []Columns, types map[string]string) (tables []zcore.OrmTable) {
	stmt, _ := db.Prepare(sqlStatementSelectTableComment)
	if stmt != nil {
		defer stmt.Close()
	}

	tbs := make(map[string]int)

	for _, column := range columns {
		if !column.TableName.Valid || !column.ColumnName.Valid {
			continue
		}

		tableName := column.TableName.String
		tableSchema := column.TableSchema.String
		columnName := column.ColumnName.String

		// init table
		index, ok := tbs[tableName]
		if !ok {
			table := zcore.OrmTable{
				Name:   zcore.UpperCamelCase(tableName),
				Table:  tableName,
				Schema: tableSchema,
			}

			// get table comment
			if stmt != nil {
				_ = stmt.QueryRow(tableName, tableSchema).Scan(&table.Comment)
			}

			tables = append(tables, table)
			index = len(tables) - 1
			tbs[tableName] = index
		}

		table := &tables[index]

		// table primary key
		if column.ColumnKey == "PRI" {
			table.Primary = columnName
		}

		c := zcore.OrmColumn{
			Name:    zcore.UpperCamelCase(columnName),
			Column:  columnName,
			Type:    column.ColumnType,
			Comment: column.ColumnComment,
			Ext:     column,
		}

		// max length
		c.MaximumLength = column.CharacterMaximumLength.Int64

		// nullable
		c.Nullable = column.IsNullable == "YES"

		// golang types match
		matches := []string{column.ColumnType, column.DataType.String}

		if c.Nullable {
			matches = append([]string{"*" + column.ColumnType, "*" + column.DataType.String}, matches...)
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
