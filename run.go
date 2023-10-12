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

	"github.com/go-zing/gozz/zcore"
	"github.com/go-zing/gozz/zutils"
)

var (
	run = &cobra.Command{
		Use:     "run",
		Short:   "run annotations analysis and use plugins to do awesome things",
		Long:    "",
		Example: zcore.ExecName + ` run -p "api" -p "enum" -p "sql:dsn=${SQL_DSN}" ./...`,
		Args:    cobra.ExactArgs(1),
		RunE:    Run,
	}

	plugins = make([]string, 0)
)

func init() {
	flags := run.Flags()
	flags.StringArrayVarP(&plugins, "plugin", "p", nil, "plugins to run")
}

func Run(_ *cobra.Command, args []string) (err error) {
	//  get analysis path absolute
	filename, err := filepath.Abs(args[0])
	if err != nil {
		return errors.New("get annotation analysis path absolute error: " + err.Error())
	}

	// validate plugins
	if len(plugins) == 0 {
		return errors.New("invalid plugins list. use -p to specify plugins")
	}

	// parse plugin entity with key-value options
	entities := make(zcore.PluginEntities, 0, len(plugins))
	registry := zcore.PluginRegistry()

	for i, plugin := range plugins {
		// split plugin name and options string
		// options would add to each comments annotation options
		// Example: name:option1=value1:option2=value2
		commands := strings.Split(plugin, ":")
		name := commands[0]

		// get registry plugin entity
		entity, ok := registry[name]
		if !ok {
			return errors.New("unregistered plugin name: " + name)
		}

		// append entities
		entities = append(entities, zcore.PluginEntity{
			Plugin:  entity,
			Options: make(map[string]string, len(commands)-1),
		})

		// parse entity options
		zutils.SplitKVSlice2Map(commands[1:], "=", entities[i].Options)
	}

	if err = entities.Run(filename); err != nil {
		_, _ = fmt.Fprint(os.Stderr, err.Error()+"\n")
		os.Exit(2)
	}
	return nil
}
