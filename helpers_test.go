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

package main

import (
	"os"
	"testing"

	zcore "github.com/go-zing/gozz-core"
)

func TestGetCoreVersion(t *testing.T) {
	version := getCoreVersion()
	v, err := zcore.ExecCommand("go list -f '{{ .Module.Version }}' "+coreDepPath, "")
	if err != nil {
		t.Fatal(err)
	}
	if version != v {
		t.Fatal()
	}
}

func TestGetPluginDir(t *testing.T) {
	if err := os.Setenv("GOZZ_PLUGINS_DIR", "test"); err != nil {
		t.Fatal(err)
	}
	if getPluginDir() != "test" {
		t.Fatal()
	}
}
