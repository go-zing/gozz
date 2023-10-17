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

import (
	"time"
)

//go:generate gozz run -p "tag" ./

// +zz:tag:json,bson:{{ snake .FieldName }}
type (
	UserStruct struct {
		Id        string    `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Address   string    `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
		Friends   []struct {
			Id        string    `bson:"id" json:"id"`
			Name      string    `bson:"name" json:"name"`
			Address   string    `bson:"address" json:"address"`
			CreatedAt time.Time `bson:"created_at" json:"created_at"`
			UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
		} `bson:"friends" json:"friends"`
	}

	UserMap map[string]struct {
		Id        string    `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Address   string    `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	}

	UserSlice []struct {
		Id        string    `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Address   string    `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	}

	UserInterface interface {
		User() struct {
			Id        string    `bson:"id" json:"id"`
			Name      string    `bson:"name" json:"name"`
			Address   string    `bson:"address" json:"address"`
			CreatedAt time.Time `bson:"created_at" json:"created_at"`
			UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
		}
	}

	UserFunc func(struct {
		Id        string    `bson:"id" json:"id"`
		Name      string    `bson:"name" json:"name"`
		Address   string    `bson:"address" json:"address"`
		CreatedAt time.Time `bson:"created_at" json:"created_at"`
		UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	})
)
