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
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

var (
	importNameCache = new(sync.Map)
	importPathCache = new(sync.Map)
	modFileCache    = new(sync.Map)
)

// loadWithStore try loads key from sync.Map or execute provided fn to store valid results
func loadWithStore(key string, m *sync.Map, fn func() string) (r string) {
	if v, ok := m.Load(key); ok {
		return v.(string)
	} else if r = fn(); len(r) > 0 {
		m.Store(key, r)
	}
	return
}

// ExecCommand execute command in provide directory and get stdout,stderr as string,error
func ExecCommand(command, dir string) (output string, err error) {
	stderr := &bytes.Buffer{}
	cmd := exec.Command("/bin/sh", "-c", command)
	cmd.Dir = dir
	cmd.Stderr = stderr
	r, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("%s:\n%s", err.Error(), stderr.String())
	}
	return UnsafeBytes2String(bytes.TrimSpace(r)), nil
}

// GetModFile get directory direct mod file by execute "go env GOMOD"
func GetModFile(dir string) string {
	return loadWithStore(dir, modFileCache, func() string {
		modFile, _ := ExecCommand("go env GOMOD", dir)
		return modFile
	})
}

// GetImportName get filename or directory module import name
// if file is not exist then return a relative calculated result from module environments
func GetImportName(filename string) string {
	return loadWithStore(filename, importNameCache, func() (name string) {
		name, dir := executeWithDir(filename, "go list -f {{.Name}}")
		if len(dir) == 0 || len(name) > 0 {
			return
		}
		// use import path base
		if p := GetImportPath(dir); len(p) > 0 {
			return path.Base(p)
		}
		// use directory base
		return filepath.Base(dir)
	})
}

// GetImportName get filename or directory module import path
// if file is not exist then return a relative calculated result from module environments
func GetImportPath(filename string) string {
	return loadWithStore(filename, importPathCache, func() (p string) {
		p, dir := executeWithDir(filename, "go list -f {{.ImportPath}}")
		if len(dir) == 0 || len(p) > 0 {
			return
		}

		// get exist directory
		tmp := dir
		for {
			if _, e := os.Stat(tmp); e == nil {
				break
			}
			tmp = filepath.Dir(tmp)
		}

		// get nearest module path
		modDir := filepath.Dir(GetModFile(tmp))
		modName, err := ExecCommand("go list -m", modDir)
		if err != nil {
			return
		}

		// computed module package import path
		rel, err := filepath.Rel(modDir, dir)
		if err != nil {
			return
		}
		return path.Join(modName, strings.Replace(rel, string(filepath.Separator), "/", -1))
	})
}

// executeInDir try executes command in provided directory or parent if filename is not directory
// return execute output and directory
func executeWithDir(filename string, command string) (ret, dir string) {
	filename, err := filepath.Abs(filename)
	if err != nil {
		return
	}
	dir = filepath.Dir(filename)
	// check file exist and is directory
	if st, e := os.Stat(filename); e == nil {
		if st.IsDir() {
			dir = filename
		}
		ret, _ = ExecCommand(command+" "+dir, dir)
	} else if !strings.HasSuffix(filename, ".go") {
		dir = filename
	}
	return
}

// FixPackage modify or add selector package to provide name according to src and dst import module info
func FixPackage(name, srcImportPath, dstImportPath string, srcImports, dstImports Imports) string {
	sp := strings.Split(name, ".")
	if len(sp) == 1 {
		if srcImportPath != dstImportPath {
			return dstImports.Add(srcImportPath) + "." + name
		}
		return name
	}
	pkgImportPath := srcImports.Which(sp[0])
	if pkgImportPath == dstImportPath {
		return sp[1]
	}
	return dstImports.Add(pkgImportPath) + "." + sp[1]
}
