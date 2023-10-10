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

package doc

import (
	"io"
	"net/http"
)

//go:generate gozz run -p "doc" ./

// +zz:doc
// this is a struct type
type T struct {
	// this is a struct field
	Field string
}

// +zz:doc
type (
	// this is another struct type declared in group
	T2 struct {
		// this is a struct field
		Field string
		// this is another struct field
		Field2 string
	}

	// this is an interface type
	T3 interface {
		// this is an interface method
		Method()
		// this is a refer anonymous interface
		io.Closer
	}

	// this is a refer type
	T4 http.Client

	// this is a map type
	T5 map[string][]string

	// this is an array type
	T6 []string
)
