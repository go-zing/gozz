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
	"fmt"
	"os"
	"runtime"
	"runtime/debug"

	"github.com/spf13/cobra"

	zcore "github.com/go-zing/gozz-core"

	_ "github.com/go-zing/gozz/internal/plugins"
)

var (
	extensions []string

	pluginDir = getPluginDir()

	cmd = cobra.Command{
		Use: zcore.ExecName,
	}

	version = &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			if info, ok := debug.ReadBuildInfo(); ok {
				fmt.Printf("%s\nruntime: %s %s/%s\ncore: %s@%s\n",
					info.Main.Path,
					runtime.Version(), runtime.GOOS, runtime.GOARCH, coreDepPath, getCoreVersion())
			}
		},
	}
)

func main() {
	cmd.AddCommand(run, list, install, version)
	cmd.PersistentFlags().StringArrayVarP(&extensions, "extension", "x", nil, "extra .so extensions plugin to load")
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
