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
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"

	zcore "github.com/go-zing/gozz-core"
)

var (
	run = &cobra.Command{
		Use:     "run",
		Short:   "run annotations analysis and use plugins to do awesome things",
		Example: zcore.ExecName + ` run -p "api" [ -p "plugin:options" ...] ./`,
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if err := Run(args); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
				os.Exit(2)
			}
		},
	}

	runPlugins = make([]string, 0)
)

func init() {
	flags := run.Flags()
	flags.StringArrayVarP(&runPlugins, "plugin", "p", nil, "plugins to run")
}

func Run(args []string) (err error) {
	if err = loadPlugins(); err != nil {
		return
	}

	//  get analysis path absolute
	filename, err := filepath.Abs(args[0])
	if err != nil {
		return errors.New("get annotation analysis path absolute error: " + err.Error())
	}

	// validate plugins
	if len(runPlugins) == 0 {
		return errors.New("invalid plugins list. use -p to specify plugins")
	}

	// parse plugin entity with key-value options
	plugins := make(zcore.PluginEntities, 0, len(runPlugins))
	registry := zcore.PluginRegistry()

	for i, plugin := range runPlugins {
		// split plugin name and options string
		// options would add to each comments annotation options
		// Example: name:option1=value1:option2=value2

		commands := strings.Split(zcore.EscapeAnnotation(plugin), ":")
		name := commands[0]

		// get registry plugin entity
		entity, ok := registry[name]
		if !ok {
			return fmt.Errorf(`unregistered plugin name: %s. use "%s list" to get registered plugins`, name, zcore.ExecName)
		}

		// append entities
		plugins = append(plugins, zcore.PluginEntity{
			Plugin:  entity,
			Options: make(map[string]string, len(commands)-1),
		})

		option := plugins[i].Options

		// parse entity options
		zcore.SplitKVSlice2Map(commands[1:], "=", option)

		for k, v := range option {
			option[k] = zcore.UnescapeAnnotation(v)
		}
	}

	return plugins.Run(filename)
}
