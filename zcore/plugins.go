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

package zcore

const (
	ExecSuffix       = "zz"
	ExecName         = "go" + ExecSuffix
	AnnotationIdent  = "+"
	AnnotationPrefix = AnnotationIdent + ExecSuffix + ":"
)

type (
	Plugin interface {
		Name() string
		Args() ([]string, map[string]string)
		Description() string
		Run(entities DeclEntities) (err error)
	}

	PluginEntity struct {
		Plugin
		Options map[string]string
	}

	PluginEntities []PluginEntity
)

var pluginRegistry = map[string]Plugin{}

func PluginRegistry() map[string]Plugin { return pluginRegistry }

func RegisterPlugin(plugin Plugin) {
	pluginRegistry[plugin.Name()] = plugin
}

func (entities PluginEntities) Run(filename string) (err error) {
	for _, entity := range entities {
		if err = entity.run(filename); err != nil {
			return
		}
	}
	return
}

func (entity PluginEntity) run(filename string) (err error) {
	decls, err := ParseFileOrDirectory(filename, AnnotationPrefix)
	if err != nil {
		return
	}
	return entity.Plugin.Run(decls.Parse(entity, entity.Options))
}
