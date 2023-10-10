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
	"sort"
)

type (
	// ranger provide range method for slice elements range and alloc
	ranger interface {
		Range(f func(element interface{}, alloc bool) (next bool))
	}

	// mapper return mapping of struct field and column name
	// keys represents column names
	// values represents pointers to struct field
	mapper interface {
		FieldMapping() map[string]interface{}
	}
)

// fieldsOf extract fields from ranger slice with mapper items
func fieldsOf(ms ranger) (fields []string) {
	rangeMapper(ms, func(m mapper, b bool) bool {
		for key := range m.FieldMapping() {
			fields = append(fields, key)
		}
		return false
	})
	sort.Strings(fields)
	return
}

// rangeMapper range slice and apply function receive mapper
func rangeMapper(ms ranger, f func(m mapper, b bool) bool) {
	ms.Range(func(v interface{}, b bool) bool { m, ok := v.(mapper); return ok && f(m, b) })
}

// scan range mapper slice and scan sql.Rows values into ranger elements
func scan(rows *sql.Rows, fields []string, ms ranger) (err error) {
	if rows.Next() {
		values := make([]interface{}, len(fields))
		rangeMapper(ms, func(m mapper, b bool) bool {
			mapping := m.FieldMapping()
			for i, field := range fields {
				values[i] = mapping[field]
			}
			err = rows.Scan(values...)
			return err == nil && rows.Next()
		})
	}
	return
}
