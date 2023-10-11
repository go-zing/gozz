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

package wire

import (
	"context"
	"database/sql"
)

//go:generate gozz run -p "wire" ./

// provide type and generate inject entry

// +zz:wire:inject=./
type Application struct {
	Service
}

// provide type

// +zz:wire
type Database struct {
	*sql.DB
}

// function to provide types

// +zz:wire
func ProvideSql() (*sql.DB, error) {
	return sql.Open("mysql", "dsn")
}

// interface and struct types bind and inject with aop

type Service interface {
	Query(ctx context.Context, id int64) (name string, err error)
}

// +zz:wire:bind=Service:aop
type ServiceImpl struct {
	Database Database
}

func (impl *ServiceImpl) Query(ctx context.Context, id int64) (name string, err error) {
	err = impl.Database.QueryRowContext(ctx, "SELECT `name` FROM `user` WHERE `id` = ?", id).Scan(&name)
	return
}
