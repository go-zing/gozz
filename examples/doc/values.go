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

// +zz:doc:label=variable
var (
	// this is a string value
	ValueString = ""
	// this is an int value
	ValueInt = 0
	// this is a struct value
	ValueStruct = struct{}{}
	// this is a pointer value
	ValuePointer = &struct{}{}
	// they are inline declaration value
	ValueInlineA, ValueInlineB = "a", "b"
	// they are inline type declaration value
	ValueInlineC, ValueInlineD int
)

// +zz:doc:label=const
const (
	// this is a constant string
	ConstantString = ""
	// this is a constant int
	ConstantInt1 = 1
	// this is another constant int
	ConstantInt2 = 2
)
