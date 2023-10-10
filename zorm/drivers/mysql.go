package drivers

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Just-maple/gozz/internal/strcase"
	"github.com/Just-maple/gozz/zorm"
)

func init() { zorm.Register(Mysql{}) }

type (
	Mysql struct{}

	SqlColumn struct {
		TableSchema            string
		TableName              string
		ColumnName             string
		OrdinalPosition        int
		IsNullable             string
		DataType               string
		CharacterSetName       *string
		CollationName          *string
		NumericPrecision       *int64
		CharacterMaximumLength *int64
	}

	MysqlColumn struct {
		SqlColumn
		ColumnType    string
		ColumnComment string
		ColumnKey     *string
	}

	SliceMysqlColumn []MysqlColumn
)

func (s *SliceMysqlColumn) Range(f func(interface{}, bool) bool) {
	for i := 0; ; i++ {
		if c := i >= len(*s); !c {
			if !f(&(*s)[i], c) {
				return
			}
		} else if n := append(*s, MysqlColumn{}); f(&n[i], c) {
			*s = n
		} else {
			*s = n[:i]
			return
		}
	}
}

func (column *MysqlColumn) FieldMapping() map[string]interface{} {
	m := column.SqlColumn.FieldMapping()
	m["column_type"] = &column.ColumnType
	m["column_key"] = &column.ColumnKey
	m["column_comment"] = &column.ColumnComment
	return m
}

func (column *SqlColumn) FieldMapping() map[string]interface{} {
	return map[string]interface{}{
		"table_schema":             &column.TableSchema,
		"table_name":               &column.TableName,
		"column_name":              &column.ColumnName,
		"ordinal_position":         &column.OrdinalPosition,
		"is_nullable":              &column.IsNullable,
		"data_type":                &column.DataType,
		"character_set_Name":       &column.CharacterSetName,
		"collation_name":           &column.CollationName,
		"numeric_precision":        &column.NumericPrecision,
		"character_maximum_length": &column.CharacterMaximumLength,
	}
}

func (m Mysql) Name() string { return "mysql" }

func (m Mysql) Parse(dsn, schema, table string, types map[string]string) (tables []zorm.Table, err error) {
	db, err := sql.Open(m.Name(), dsn)
	if err != nil {
		return
	}
	defer db.Close()

	columns, err := m.queryColumns(db, schema, table)
	if err != nil {
		return
	}

	tables = m.parseTables(columns, types)
	return
}

const sqlStatementSelectColumns = "SELECT `%s` FROM `information_schema`.`COLUMNS` WHERE `TABLE_SCHEMA` = ?"

func (m Mysql) queryColumns(db *sql.DB, schema string, table string) (columns SliceMysqlColumn, err error) {
	fields := fieldsOf(&columns)
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
	err = scan(rows, fields, &columns)
	return
}

func (m Mysql) parseTables(columns []MysqlColumn, types map[string]string) (tables []zorm.Table) {
	tbs := make(map[string]*zorm.Table)

	for _, column := range columns {
		// init table
		tb, ok := tbs[column.TableName]
		if !ok {
			tb = &zorm.Table{
				Name:   strcase.UpperCamelCase(column.TableName),
				Table:  column.TableName,
				Schema: column.TableSchema,
			}
			tbs[column.TableName] = tb
		}

		// table primary key
		if column.ColumnKey != nil && *column.ColumnKey == "PRI" {
			tb.Primary = column.ColumnName
		}

		c := zorm.Column{
			Name:    strcase.UpperCamelCase(column.ColumnName),
			Column:  column.ColumnName,
			Type:    column.ColumnType,
			Comment: column.ColumnComment,
		}

		// max length
		if column.CharacterMaximumLength != nil {
			c.MaximumLength = *column.CharacterMaximumLength
		}

		// nullable
		c.Nullable = column.IsNullable == "YES"

		// golang types match
		matches := make([]string, 0)
		if c.Nullable {
			matches = append(matches, "*"+column.ColumnType, "*"+column.DataType)
		}
		matches = append(matches, column.ColumnType, column.DataType)

		// match first
		for _, key := range matches {
			if c.Type = types[key]; len(c.Type) > 0 {
				break
			}
		}

		tb.Columns = append(tb.Columns, c)
	}

	for _, tb := range tbs {
		tables = append(tables, *tb)
	}

	sort.Slice(tables, func(i, j int) bool { return tables[i].Name < tables[j].Name })
	return
}
