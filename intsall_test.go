//go:build !windows
// +build !windows

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
	"runtime"
	"testing"
)

func TestGetGoEnv(t *testing.T) {
	env, err := getGoenv("./")
	if err != nil {
		t.Fatal(err)
	}
	if env["GOARCH"] != runtime.GOARCH {
		t.Fatalf("get value unexpected: %v != %v ", env["GOARCH"], runtime.GOARCH)
	}
	if env["GOOS"] != runtime.GOOS {
		t.Fatalf("get value unexpected: %v != %v ", env["GOOS"], runtime.GOOS)
	}
}

func TestInstall(t *testing.T) {
	defer os.Remove("tmp.so")
	installBuildOutput = "tmp.so"
	installBuildTarget = "./ormdrivers/mysql"
	if err := Install("https://github.com/go-zing/gozz-plugins"); err != nil {
		t.Fatal(err)
	}
}
