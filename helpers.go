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
	"path/filepath"
	"runtime/debug"
	"strconv"
	"strings"

	zcore "github.com/go-zing/gozz-core"
)

func loadPlugins() (err error) {
	for _, name := range extensions {
		if _, err = zcore.LoadExtension(name); err != nil {
			return
		}
	}
	if len(pluginDir) > 0 {
		_ = zcore.WalkDir(pluginDir, func(name string) error {
			_, _ = zcore.LoadExtension(name)
			return nil
		})
	}
	return
}

func getGoenv(dir string) (env map[string]string, err error) {
	goenv, err := zcore.ExecCommand("go env", dir)
	if err != nil {
		return
	}
	env = make(map[string]string)
	for _, line := range strings.Split(goenv, "\n") {
		if line = strings.TrimSpace(line); len(line) == 0 {
			continue
		}
		if kv := strings.SplitN(line, "=", 2); len(kv) >= 2 {
			env[kv[0]], _ = strconv.Unquote(strings.Trim(kv[1], "'"))
		}
	}
	return
}

func getCoreVersion() string {
	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return ""
	}
	for _, m := range bi.Deps {
		if m.Path == coreDepPath {
			return m.Version
		}
	}
	return ""
}

func getPluginDir() string {
	if dir := os.Getenv("GOZZ_PLUGINS_DIR"); len(dir) > 0 {
		return dir
	} else if homeDir, _ := os.UserHomeDir(); len(homeDir) > 0 {
		return filepath.Join(homeDir, ".gozz", "extensions")
	}
	return ""
}
