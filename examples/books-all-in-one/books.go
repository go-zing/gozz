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

package books_all_in_one

import (
	"context"
	"time"
)

//go:generate gozz run -p "api" -p "tag" -p "doc" -p "impl" -p "wire" ./

// +zz:wire:inject=./inject
type Application struct {
	Apis
}

// +zz:api:./:prefix=books
// +zz:doc
// +zz:impl:./impls:*BookImpl:wire
// BookService provide book management services
type BookService interface {
	// +zz:api:get:
	// List all books. return ListBook
	List(ctx context.Context, query QueryBook) (ret ListBook, err error)
	// +zz:api:get:{id}
	// Get book by book id. returns DataBook
	Get(ctx context.Context, book QueryBook) (data DataBook, err error)
	// +zz:api:post:
	// Create new book from FormBook. returns DataBook created
	Create(ctx context.Context, book FormBook) (data DataBook, err error)
	// +zz:api:put:{id}
	// Edit book by id from FormBook. returns DataBook edited
	Edit(ctx context.Context, book FormBook) (data DataBook, err error)
}

// +zz:doc
// +zz:tag:json:{{ snake .FieldName }}
type (
	// QueryBook represents struct for querying book list or book item
	// +zz:tag:form:{{ snake .FieldName }}
	QueryBook struct {
		// specify query id
		Id int `json:"id" form:"id"`
		// specify query title keywords
		Title string `json:"title" form:"title"`
		// specify query pagination page no. default: 1
		PageNo int `json:"page_no" form:"page_no"`
		// specify query pagination page count. default: 20
		PageCount int `json:"page_count" form:"page_count"`
	}

	// FormBook represents struct for creating or editing book
	FormBook struct {
		// +zz:tag:+json:,omitempty
		// +zz:tag:uri:{{ snake .FieldName }}
		// specify editing id. only works for editing
		Id int `json:"id,omitempty" uri:"id"`
		// book title
		Title string `json:"title"`
		// book type
		Type string `json:"type"`
		// book description
		Description string `json:"description"`
		// book author
		Author string `json:"author"`
	}

	// DataBook represents struct for book item
	DataBook struct {
		FormBook
		// book created time
		CreatedAt time.Time `json:"created_at"`
		// book created username
		CreatedBy string `json:"created_by"`
		// book last updated time
		UpdatedAt time.Time `json:"updated_at"`
		// book last updated username
		UpdatedBy string `json:"updated_by"`
	}

	// ListBook represents struct form querying book list and total count of books in datastore
	ListBook struct {
		// query book results
		List []DataBook `json:"list"`
		// total count of books in datastore
		Total int64 `json:"total"`
	}
)

// +zz:doc:label=book_type
const (
	BookTypeAdventureStories  = "Adventure stories"
	BookTypeClassics          = "Classics"
	BookTypeCrime             = "Crime"
	BookTypeFantasy           = "Fantasy"
	BookTypeHistoricalFiction = "Historical fiction"
	BookTypeHorror            = "Horror"
	BookTypeHumourAndSatire   = "Humour and satire"
	BookTypeLiteraryFiction   = "Literary fiction"
	BookTypeMystery           = "Mystery"
	BookTypePoetry            = "Poetry"
	BookTypePlays             = "Plays"
	BookTypeRomance           = "Romance"
	BookTypeScienceFiction    = "Science fiction"
	BookTypeShortStories      = "Short stories"
	BookTypeThrillers         = "Thrillers"
	BookTypeWar               = "War"
	BookTypeWomenFiction      = "Women’s fiction"
	BookTypeYoungAdult        = "Young adult"
)
