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

package impls

import (
	"context"

	booksallinone "github.com/go-zing/gozz/examples/books-all-in-one"
)

var (
	_ booksallinone.BookService = (*BookServiceImpl)(nil)
)

// +zz:wire:bind=booksallinone.BookService:aop
type BookServiceImpl struct{}

func (impl *BookServiceImpl) List(ctx context.Context, query booksallinone.QueryBook) (ret booksallinone.ListBook, err error) {
	panic("not implemented")
}

func (impl *BookServiceImpl) Get(ctx context.Context, book booksallinone.QueryBook) (data booksallinone.DataBook, err error) {
	panic("not implemented")
}

func (impl *BookServiceImpl) Create(ctx context.Context, book booksallinone.FormBook) (data booksallinone.DataBook, err error) {
	panic("not implemented")
}

func (impl *BookServiceImpl) Edit(ctx context.Context, book booksallinone.FormBook) (data booksallinone.DataBook, err error) {
	panic("not implemented")
}
