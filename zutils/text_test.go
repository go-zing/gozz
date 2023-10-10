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

package zutils

import (
	"testing"
)

func TestSplitKV(t *testing.T) {
	for _, c := range [][4]string{
		{"k=v=3", "=", "k", "v=3"},
		{"k=v", "=", "k", "v"},
		{"k", "=", "k", ""},
		{"", "=", "", ""},
	} {
		k, v := SplitKV(c[0], c[1])
		if k != c[2] || v != c[3] {
			t.Fatalf("want (%v,%v) got (%v,%v)", c[2], c[3], k, v)
		}
	}
}
