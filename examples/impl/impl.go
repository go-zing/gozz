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

package impl

import "context"

var (
	_ Interface = (*T)(nil)
)

type T struct{}

func (t *T) Api() {
	panic("not implemented")
}

func (t *T) Api1(ctx context.Context, param Param) Result {
	panic("not implemented")
}

func (t *T) Api2(ctx context.Context, param Param) []Result {
	panic("not implemented")
}

func (t *T) Api3(ctx context.Context, param Param) (r []Result, err error) {
	panic("not implemented")
}

func (t *T) Api4(ctx context.Context, param Param) (r map[*context.Context]Result, err error) {
	panic("not implemented")
}

func (t *T) Api5(ctx context.Context) (fn func(ctx context.Context) interface {
	Foo(param Param) (Result, error)
}) {
	panic("not implemented")
}
