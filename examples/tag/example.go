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

package tag

//go:generate gozz run -p "tag" ./

// +zz:tag:json:{{ snake .FieldName }}
type (
	T struct {
		FieldA string `json:"field_a"`
		FieldB string `json:"field_b"`
		FieldC string `json:"field_c"`
	}

	// +zz:tag:json:{{ camel .FieldName }}
	T2 struct {
		FieldA string `json:"fieldA"`
		FieldB string `json:"fieldB"`
		FieldC string `json:"fieldC"`
	}

	// +zz:tag:json:{{ kebab .FieldName }}
	T3 struct {
		// +zz:tag:json:{{ camel .FieldName }}
		FieldA     string `json:"fieldA"`
		FieldB     string `json:"field-b"`
		FieldInner []map[struct {
			// +zz:tag:+json:,omitempty
			FieldA string `json:"field-a,omitempty"`
			FieldB string `json:"field-b"`
		}]func() struct {
			FieldC string `json:"field-c"`
			FieldD string `json:"field-d"`
		} `json:"field-inner"`
	}
)

// +zz:tag:json:{{ snake .FieldName }}
type T4 map[string]struct {
	FieldA string `json:"field_a"`
	FieldB string `json:"field_b"`
	FieldC string `json:"field_c"`
}

// +zz:tag:json:{{ snake .FieldName }}
type T5 func() map[string]struct {
	// +zz:tag:json:{{ camel .FieldName }}
	FieldA string `json:"fieldA"`
	FieldB string `json:"field_b"`
	FieldC string `json:"field_c"`
}
