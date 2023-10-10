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

type mapperSlice interface {
	Range(f func(interface{}, bool) bool)
}

type mapper interface {
	FieldMapping() map[string]interface{}
}

func fieldsOf(ms mapperSlice) (fields []string) {
	rangeSlice(ms, func(m mapper, b bool) bool {
		for key := range m.FieldMapping() {
			fields = append(fields, key)
		}
		return false
	})
	sort.Strings(fields)
	return
}

func rangeSlice(ms mapperSlice, f func(m mapper, b bool) bool) {
	ms.Range(func(v interface{}, b bool) bool { m, ok := v.(mapper); return ok && f(m, b) })
}

func scan(rows *sql.Rows, fields []string, ms mapperSlice) (err error) {
	if rows.Next() {
		values := make([]interface{}, len(fields))
		rangeSlice(ms, func(m mapper, b bool) bool {
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
