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

package api

import (
	"context"
)

//go:generate gozz run -p "api" ./

type (
	QueryBook struct{}
	FormBook  struct{}
	DataBook  struct{}
	ListBook  struct{}
)

// +zz:api:./:prefix=books
type BookService interface {
	// +zz:api:get:
	List(ctx context.Context, query QueryBook) (ret ListBook, err error)
	// +zz:api:get:{id}
	Get(ctx context.Context, book QueryBook) (data DataBook, err error)
	// +zz:api:post:
	Create(ctx context.Context, book FormBook) (data DataBook, err error)
	// +zz:api:put:{id}
	Edit(ctx context.Context, book FormBook) (data DataBook, err error)
}