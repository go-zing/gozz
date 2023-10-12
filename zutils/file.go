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
	"bytes"
	"crypto/md5"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	// file store cached opened file bytes with version key consists of size and modify time
	fileStore = new(VersionStore)
	// ast store cached parsed file *ast.File with version key consists of size and modify time
	astStore = new(VersionStore)
)

// fileVersion return file version key consists of size and modify time
func fileVersion(info os.FileInfo) string {
	return fmt.Sprintf("%d-%s", info.Size(), info.ModTime())
}

// ReadFile try read filename and return data bytes.
// data bytes would be cached by version key
func ReadFile(filename string) (data []byte, version string, err error) {
	info, err := os.Stat(filename)
	if err != nil {
		return
	}
	version = fileVersion(info)
	r, err := fileStore.Load(filename, version, func() (interface{}, error) {
		return ioutil.ReadFile(filename)
	})
	if err != nil {
		return
	}
	data = r.([]byte)
	return
}

// ParseFile try read file data and parse ast file.
// return values would be cached by version key
func ParseFile(filename string) (file *File, err error) {
	data, version, err := ReadFile(filename)
	if err != nil {
		return
	}
	r, err := astStore.Load(filename, version, func() (interface{}, error) {
		return parser.ParseFile(token.NewFileSet(), filename, data, parser.ParseComments)
	})
	if err != nil {
		return
	}
	return &File{Path: filename, Data: data, Ast: r.(*ast.File)}, nil
}

// WriteFile checks data and exists filename md5 sum
// and update data if file not exists or md5 sum not matched
func WriteFile(filename string, data []byte, perm fs.FileMode) (updated bool, err error) {
	if err = os.MkdirAll(filepath.Dir(filename), 0o775); err != nil {
		return
	}

	// check file exist
	exist, _, err := ReadFile(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return
		}
		return true, ioutil.WriteFile(filename, data, perm)
	}

	// compare exist and new data by md5 sum
	oldSum := md5.Sum(exist)
	if newSum := md5.Sum(data); bytes.Equal(newSum[:], oldSum[:]) {
		return false, nil
	}

	if info, e := os.Stat(filename); e == nil {
		perm = info.Mode()
	}

	if err = ioutil.WriteFile(filename, data, perm); err != nil {
		return
	}

	if info, e := os.Stat(filename); e == nil {
		fileStore.Update(filename, fileVersion(info), data)
	}
	return true, nil
}

// WalkPackage walk package directory and parse file as *File. return *File map with filename
func WalkPackage(dir string, fn func(file *File) (err error)) (files map[string]*File, err error) {
	files = make(map[string]*File)
	err = filepath.Walk(dir, func(filename string, info fs.FileInfo, err error) error {
		if os.IsNotExist(err) {
			return filepath.SkipDir
		} else if err != nil {
			return err
		} else if info.IsDir() && filename != dir {
			return filepath.SkipDir
		} else if !IsGoFile(filename) {
			return nil
		}

		f, err := ParseFile(filename)
		if err != nil {
			return err
		}
		files[filename] = f
		return fn(f)
	})
	return
}

// IsGoFile check path is valid golang file and ignore test file
func IsGoFile(path string) bool {
	return strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go")
}
