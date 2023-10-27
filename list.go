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
	"sort"
	"strings"

	zcore "github.com/go-zing/gozz-core"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "list all registered plugins",
	Run: func(_ *cobra.Command, _ []string) {
		if err := loadPlugins(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(2)
		}
		runList()
	},
}

func runList() {
	registry := zcore.PluginRegistry()
	names := make([]string, 0)
	for name := range registry {
		names = append(names, name)
	}
	sort.Strings(names)
	fmt.Printf("totally %d plugins avaiable:\n", len(names))

	for _, name := range names {
		p := registry[name]
		desc := p.Description()
		args, options := p.Args()

		usage := zcore.AnnotationPrefix + name
		argsHelp := &strings.Builder{}

		for i, arg := range args {
			arg, help := zcore.SplitKV(arg, ":")
			if i == 0 {
				argsHelp.WriteString("\n\targs:")
			}
			usage += ":[" + arg + "]"
			argsHelp.WriteString("\n\t\t" + arg + ": " + help)
		}

		var keys []string
		for k := range options {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for i, key := range keys {
			if i == 0 {
				argsHelp.WriteString("\n\toptions:")
				usage += ":[options...]"
			}
			argsHelp.WriteString("\n\t\t" + key + ": " + options[key])
		}

		str := fmt.Sprintf("\n%s: %s\n\t%s%s\n", name, usage, desc, argsHelp)
		fmt.Print(strings.Replace(str, "\t", "    ", -1))
	}
}
