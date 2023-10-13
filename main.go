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

	"github.com/spf13/cobra"

	zcore "github.com/go-zing/gozz-core"

	_ "github.com/go-zing/gozz/internal/plugins"
)

var (
	extensions []string

	cmd = cobra.Command{
		Use: zcore.ExecName,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) (err error) {
			for _, name := range extensions {
				if err = zcore.LoadExtension(name); err != nil {
					return
				}
			}
			// load extension in ~/.gozz/extensions
			if homeDir, _ := os.UserHomeDir(); len(homeDir) > 0 {
				_ = zcore.WalkDir(filepath.Join(homeDir, ".gozz", "extensions"), func(name string) error {
					_ = zcore.LoadExtension(name)
					return nil
				})
			}
			return
		},
	}
)

func main() {
	cmd.AddCommand(run, list)
	cmd.PersistentFlags().StringArrayVarP(&extensions, "extension", "x", nil, "extra .so extensions plugin to load")
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
