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
	"sort"
	"strings"

	"github.com/spf13/cobra"

	"github.com/go-zing/gozz/zcore"
	"github.com/go-zing/gozz/zutils"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "list all registered plugins",
	Run: func(cmd *cobra.Command, args []string) {
		registry := zcore.PluginRegistry()
		names := make([]string, 0)
		for name := range registry {
			names = append(names, name)
		}
		sort.Strings(names)
		fmt.Printf("total %d plugins avaiable:\n", len(names))

		for _, name := range names {
			p := registry[name]
			name := p.Name()
			desc := p.Description()
			args, options := p.Args()

			usage := zcore.AnnotationPrefix + name
			argsHelp := ""

			for i, arg := range args {
				arg, help := zutils.SplitKV(arg, ":")
				if i == 0 {
					argsHelp += "\n\targs:"
				}
				usage += ":[" + arg + "]"
				argsHelp += "\n\t\t" + arg + ": " + help
			}

			var keys []string
			for k := range options {
				keys = append(keys, k)
			}
			sort.Strings(keys)

			for i, key := range keys {
				if i == 0 {
					argsHelp += "\n\toptions:"
					usage += ":[options...]"
				}
				argsHelp += "\n\t\t" + key + ": " + options[key]
			}

			str := fmt.Sprintf("\n%s: %s\n\t%s%s\n", name, usage, desc, argsHelp)
			fmt.Print(strings.Replace(str, "\t", "    ", -1))
		}
	},
}
