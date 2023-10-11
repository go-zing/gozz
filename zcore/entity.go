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

import (
	"path/filepath"
	"strings"

	"github.com/Just-maple/gozz/zutils"
)

type (
	DeclEntity struct {
		*AnnotatedDecl

		Plugin  string
		Args    []string
		Options Options
	}

	DeclEntities []DeclEntity

	FieldEntity struct {
		*AnnotatedField

		Args    []string
		Options Options
	}

	FieldEntities []FieldEntity

	Options map[string]string
)

func (opt Options) Get(key string, def string) string {
	if v, ok := opt[key]; ok && len(v) > 0 {
		return v
	}
	return def
}

// parseAnnotation parse annotation string
// annotation strings would split by ":" and check first matches provided name
// if not matched then return ok=false
// on matched rests items would be divided into args and options according to args count
// args is strings slice and options is key-value pairs split by "="
// extOptions would fill options if extOptions key not in parsed options
//
// annotation format  $name:$args1:$args2:...$argsN:$key1=$value1:$key2=$value2:...
//
// for example
//
// params:
// 	 annotation  foo:args1:args2:key1=value1:key2=value2
//   name        foo
//   argsCount   2
//   extOptions  [key3:value3 key4:value4]
//
// returns:
// 	 args        [args1 args2]
//   options     [key1:value1 key2:value2 key3:value3 key4:value4]
//   ok          true
func parseAnnotation(annotation, name string, argsCount int, extOptions map[string]string) (args []string, options map[string]string, ok bool) {
	sp := strings.Split(annotation, ":")
	if sp[0] != name || len(sp)-1 < argsCount {
		return
	}
	options = make(map[string]string)
	zutils.SplitKVSlice2Map(sp[1+argsCount:], "=", options)
	for k, v := range extOptions {
		if _, exist := options[k]; exist {
			continue
		}
		options[k] = v
	}
	return sp[1 : 1+argsCount], options, true
}

func (entities DeclEntities) GroupByDir() (m map[string]DeclEntities) {
	return entities.GroupBy(func(entity DeclEntity) string {
		return filepath.Dir(entity.File.Path)
	})
}

func (entities DeclEntities) GroupBy(fn func(entity DeclEntity) string) (m map[string]DeclEntities) {
	m = make(map[string]DeclEntities)
	for _, entity := range entities {
		if key := fn(entity); len(key) > 0 {
			m[key] = append(m[key], entity)
		}
	}
	return
}

func (entity *DeclEntity) ParseFields(argsCount int, options map[string]string) (fields FieldEntities) {
	for _, field := range entity.Fields {
		fields = append(fields, field.Parse(entity.Plugin, argsCount, options)...)
	}
	return
}

func (decl *AnnotatedDecl) Name() string {
	if decl.TypeSpec != nil && decl.TypeSpec.Name != nil {
		return decl.TypeSpec.Name.Name
	}
	if decl.FuncDecl != nil && decl.FuncDecl.Name != nil {
		return decl.FuncDecl.Name.Name
	}
	if decl.ValueSpec != nil && len(decl.ValueSpec.Names) == 1 {
		return decl.ValueSpec.Names[0].Name
	}
	return ""
}

func (decl *AnnotatedDecl) RelFilename(filename string, defaultName string) (ret string) {
	if !strings.HasSuffix(filename, ".go") {
		filename = filepath.Join(filename, defaultName)
	}

	if dir := filepath.Dir(decl.File.Path); filepath.IsAbs(filename) {
		ret = filepath.Join(filepath.Dir(zutils.GetModFile(dir)), filename)
	} else {
		ret = filepath.Join(dir, filename)
	}

	if strings.Contains(ret, "{{") && strings.Contains(ret, "}}") {
		if str := (&strings.Builder{}); ExecuteTemplate(decl, ret, str) == nil {
			ret = str.String()
		}
	}
	return
}

func (decls AnnotatedDecls) Parse(plugin Plugin, extOptions map[string]string) (entities DeclEntities) {
	name := plugin.Name()
	args, _ := plugin.Args()
	for _, decl := range decls {
		entities = append(entities, decl.parse(name, len(args), extOptions)...)
	}
	return
}

func (decl *AnnotatedDecl) parse(name string, argsCount int, extOptions map[string]string) (entities DeclEntities) {
	for _, annotation := range decl.Annotations {
		args, opts, ok := parseAnnotation(annotation, name, argsCount, extOptions)
		if !ok {
			continue
		}
		entities = append(entities, DeclEntity{
			AnnotatedDecl: decl,
			Plugin:        name,
			Args:          args,
			Options:       opts,
		})
	}
	return
}

func (field *AnnotatedField) Parse(name string, argsCount int, extOptions map[string]string) (entities FieldEntities) {
	for _, annotation := range field.Annotations {
		args, opts, ok := parseAnnotation(annotation, name, argsCount, extOptions)
		if !ok {
			continue
		}
		entities = append(entities, FieldEntity{
			AnnotatedField: field,
			Args:           args,
			Options:        opts,
		})
	}
	return
}
