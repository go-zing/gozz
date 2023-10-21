// Code generated by gozz:orm github.com/go-zing/gozz. DO NOT EDIT.

package ormdrivers

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

var (
	_ = (*context.Context)(nil)
	_ = (*json.RawMessage)(nil)
	_ = (*time.Time)(nil)
	_ = (*sql.NullString)(nil)
)

// information_schema.COLUMNS
const TableColumns = "COLUMNS"

type Columns struct {
	// TABLE_CATALOG : NULLABLE varchar(64)
	TableCatalog sql.NullString
	// TABLE_SCHEMA : NULLABLE varchar(64)
	TableSchema sql.NullString
	// TABLE_NAME : NULLABLE varchar(64)
	TableName sql.NullString
	// COLUMN_NAME : NULLABLE varchar(64)
	ColumnName sql.NullString
	// ORDINAL_POSITION : int unsigned
	OrdinalPosition int
	// COLUMN_DEFAULT : NULLABLE text
	ColumnDefault sql.NullString
	// IS_NULLABLE : varchar(3)
	IsNullable string
	// DATA_TYPE : NULLABLE longtext
	DataType sql.NullString
	// CHARACTER_MAXIMUM_LENGTH : NULLABLE bigint
	CharacterMaximumLength sql.NullInt64
	// CHARACTER_OCTET_LENGTH : NULLABLE bigint
	CharacterOctetLength sql.NullInt64
	// NUMERIC_PRECISION : NULLABLE bigint unsigned
	NumericPrecision sql.NullInt64
	// NUMERIC_SCALE : NULLABLE bigint unsigned
	NumericScale sql.NullInt64
	// DATETIME_PRECISION : NULLABLE int unsigned
	DatetimePrecision sql.NullInt32
	// CHARACTER_SET_NAME : NULLABLE varchar(64)
	CharacterSetName sql.NullString
	// COLLATION_NAME : NULLABLE varchar(64)
	CollationName sql.NullString
	// COLUMN_TYPE : mediumtext
	ColumnType string
	// COLUMN_KEY : enum('','PRI','UNI','MUL')
	ColumnKey string
	// EXTRA : NULLABLE varchar(256)
	Extra sql.NullString
	// PRIVILEGES : NULLABLE varchar(154)
	Privileges sql.NullString
	// COLUMN_COMMENT : text
	ColumnComment string
	// GENERATION_EXPRESSION : longtext
	GenerationExpression string
}

func (m *Columns) FieldMapping(dst map[string]interface{}) {
	dst["TABLE_CATALOG"] = &m.TableCatalog
	dst["TABLE_SCHEMA"] = &m.TableSchema
	dst["TABLE_NAME"] = &m.TableName
	dst["COLUMN_NAME"] = &m.ColumnName
	dst["ORDINAL_POSITION"] = &m.OrdinalPosition
	dst["COLUMN_DEFAULT"] = &m.ColumnDefault
	dst["IS_NULLABLE"] = &m.IsNullable
	dst["DATA_TYPE"] = &m.DataType
	dst["CHARACTER_MAXIMUM_LENGTH"] = &m.CharacterMaximumLength
	dst["CHARACTER_OCTET_LENGTH"] = &m.CharacterOctetLength
	dst["NUMERIC_PRECISION"] = &m.NumericPrecision
	dst["NUMERIC_SCALE"] = &m.NumericScale
	dst["DATETIME_PRECISION"] = &m.DatetimePrecision
	dst["CHARACTER_SET_NAME"] = &m.CharacterSetName
	dst["COLLATION_NAME"] = &m.CollationName
	dst["COLUMN_TYPE"] = &m.ColumnType
	dst["COLUMN_KEY"] = &m.ColumnKey
	dst["EXTRA"] = &m.Extra
	dst["PRIVILEGES"] = &m.Privileges
	dst["COLUMN_COMMENT"] = &m.ColumnComment
	dst["GENERATION_EXPRESSION"] = &m.GenerationExpression
}

type SliceColumns []Columns

func (s *SliceColumns) Range(f func(interface{}, bool) bool) {
	for i := 0; ; i++ {
		if c := i >= len(*s); !c {
			if !f(&(*s)[i], c) {
				return
			}
		} else if n := append(*s, Columns{}); f(&n[i], c) {
			*s = n
		} else {
			*s = n[:i]
			return
		}
	}
}
