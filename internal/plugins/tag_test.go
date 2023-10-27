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

package plugins

import (
	"bytes"
	"encoding/base64"
	"io/ioutil"
	"os"
	"testing"

	zcore "github.com/go-zing/gozz-core"
)

const testTagData = `package tag05

import (
	"time"
)

//go:generate gozz run -p "tag" ./

// +zz:tag:json,bson:{{ snake .FieldName }}
type (
	UserStruct struct {
		Id        string    
		Name      string    
		Address   string    
		CreatedAt time.Time 
		UpdatedAt time.Time 
		Friends   []struct {
			Id        string    
			Name      string    
			Address   string    
			CreatedAt time.Time 
			UpdatedAt time.Time 
		} 
	}

	UserMap map[string]struct {
		Id        string    
		Name      string    
		Address   string    
		CreatedAt time.Time 
		UpdatedAt time.Time 
	}

	UserSlice []struct {
		Id        string    
		Name      string    
		Address   string    
		CreatedAt time.Time 
		UpdatedAt time.Time 
	}

	UserInterface interface {
		User() struct {
			Id        string    
			Name      string    
			Address   string    
			CreatedAt time.Time 
			UpdatedAt time.Time 
		}
	}

	UserFunc func(struct {
		Id        string    
		Name      string    
		Address   string    
		CreatedAt time.Time 
		UpdatedAt time.Time 
	})
)
`

var testTagRetData, _ = base64.StdEncoding.DecodeString(`cGFja2FnZSB0YWcwNQoKaW1wb3J0ICgKCSJ0aW1lIgopCgovL2dvOmdlbmVyYXRlIGdvenogcnVuIC1wICJ0YWciIC4vCgovLyAreno6dGFnOmpzb24sYnNvbjp7eyBzbmFrZSAuRmllbGROYW1lIH19CnR5cGUgKAoJVXNlclN0cnVjdCBzdHJ1Y3QgewoJCUlkICAgICAgICBzdHJpbmcgICAgYGJzb246ImlkIiBqc29uOiJpZCJgCgkJTmFtZSAgICAgIHN0cmluZyAgICBgYnNvbjoibmFtZSIganNvbjoibmFtZSJgCgkJQWRkcmVzcyAgIHN0cmluZyAgICBgYnNvbjoiYWRkcmVzcyIganNvbjoiYWRkcmVzcyJgCgkJQ3JlYXRlZEF0IHRpbWUuVGltZSBgYnNvbjoiY3JlYXRlZF9hdCIganNvbjoiY3JlYXRlZF9hdCJgCgkJVXBkYXRlZEF0IHRpbWUuVGltZSBgYnNvbjoidXBkYXRlZF9hdCIganNvbjoidXBkYXRlZF9hdCJgCgkJRnJpZW5kcyAgIFtdc3RydWN0IHsKCQkJSWQgICAgICAgIHN0cmluZyAgICBgYnNvbjoiaWQiIGpzb246ImlkImAKCQkJTmFtZSAgICAgIHN0cmluZyAgICBgYnNvbjoibmFtZSIganNvbjoibmFtZSJgCgkJCUFkZHJlc3MgICBzdHJpbmcgICAgYGJzb246ImFkZHJlc3MiIGpzb246ImFkZHJlc3MiYAoJCQlDcmVhdGVkQXQgdGltZS5UaW1lIGBic29uOiJjcmVhdGVkX2F0IiBqc29uOiJjcmVhdGVkX2F0ImAKCQkJVXBkYXRlZEF0IHRpbWUuVGltZSBgYnNvbjoidXBkYXRlZF9hdCIganNvbjoidXBkYXRlZF9hdCJgCgkJfSBgYnNvbjoiZnJpZW5kcyIganNvbjoiZnJpZW5kcyJgCgl9CgoJVXNlck1hcCBtYXBbc3RyaW5nXXN0cnVjdCB7CgkJSWQgICAgICAgIHN0cmluZyAgICBgYnNvbjoiaWQiIGpzb246ImlkImAKCQlOYW1lICAgICAgc3RyaW5nICAgIGBic29uOiJuYW1lIiBqc29uOiJuYW1lImAKCQlBZGRyZXNzICAgc3RyaW5nICAgIGBic29uOiJhZGRyZXNzIiBqc29uOiJhZGRyZXNzImAKCQlDcmVhdGVkQXQgdGltZS5UaW1lIGBic29uOiJjcmVhdGVkX2F0IiBqc29uOiJjcmVhdGVkX2F0ImAKCQlVcGRhdGVkQXQgdGltZS5UaW1lIGBic29uOiJ1cGRhdGVkX2F0IiBqc29uOiJ1cGRhdGVkX2F0ImAKCX0KCglVc2VyU2xpY2UgW11zdHJ1Y3QgewoJCUlkICAgICAgICBzdHJpbmcgICAgYGJzb246ImlkIiBqc29uOiJpZCJgCgkJTmFtZSAgICAgIHN0cmluZyAgICBgYnNvbjoibmFtZSIganNvbjoibmFtZSJgCgkJQWRkcmVzcyAgIHN0cmluZyAgICBgYnNvbjoiYWRkcmVzcyIganNvbjoiYWRkcmVzcyJgCgkJQ3JlYXRlZEF0IHRpbWUuVGltZSBgYnNvbjoiY3JlYXRlZF9hdCIganNvbjoiY3JlYXRlZF9hdCJgCgkJVXBkYXRlZEF0IHRpbWUuVGltZSBgYnNvbjoidXBkYXRlZF9hdCIganNvbjoidXBkYXRlZF9hdCJgCgl9CgoJVXNlckludGVyZmFjZSBpbnRlcmZhY2UgewoJCVVzZXIoKSBzdHJ1Y3QgewoJCQlJZCAgICAgICAgc3RyaW5nICAgIGBic29uOiJpZCIganNvbjoiaWQiYAoJCQlOYW1lICAgICAgc3RyaW5nICAgIGBic29uOiJuYW1lIiBqc29uOiJuYW1lImAKCQkJQWRkcmVzcyAgIHN0cmluZyAgICBgYnNvbjoiYWRkcmVzcyIganNvbjoiYWRkcmVzcyJgCgkJCUNyZWF0ZWRBdCB0aW1lLlRpbWUgYGJzb246ImNyZWF0ZWRfYXQiIGpzb246ImNyZWF0ZWRfYXQiYAoJCQlVcGRhdGVkQXQgdGltZS5UaW1lIGBic29uOiJ1cGRhdGVkX2F0IiBqc29uOiJ1cGRhdGVkX2F0ImAKCQl9Cgl9CgoJVXNlckZ1bmMgZnVuYyhzdHJ1Y3QgewoJCUlkICAgICAgICBzdHJpbmcgICAgYGJzb246ImlkIiBqc29uOiJpZCJgCgkJTmFtZSAgICAgIHN0cmluZyAgICBgYnNvbjoibmFtZSIganNvbjoibmFtZSJgCgkJQWRkcmVzcyAgIHN0cmluZyAgICBgYnNvbjoiYWRkcmVzcyIganNvbjoiYWRkcmVzcyJgCgkJQ3JlYXRlZEF0IHRpbWUuVGltZSBgYnNvbjoiY3JlYXRlZF9hdCIganNvbjoiY3JlYXRlZF9hdCJgCgkJVXBkYXRlZEF0IHRpbWUuVGltZSBgYnNvbjoidXBkYXRlZF9hdCIganNvbjoidXBkYXRlZF9hdCJgCgl9KQopCg==`)

func TestTag(t *testing.T) {
	_ = os.MkdirAll("test", 0o775)
	defer os.RemoveAll("test")

	if err := os.WriteFile("test/types.go", []byte(testTagData), 0o664); err != nil {
		t.Fatal(err)
	}

	decls, err := zcore.ParseFileOrDirectory("test/types.go", zcore.AnnotationPrefix)
	if err != nil {
		return
	}

	plugin := &Tag{}
	if err = plugin.Run(decls.Parse(plugin, nil)); err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadFile("test/types.go")
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(data, testTagRetData) {
		t.Fatal(err)
	}
}
