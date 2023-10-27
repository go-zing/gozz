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

package plugins

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	zcore "github.com/go-zing/gozz-core"
)

const testImplData = `package x

import "context"

type (
	// +zz:impl:./types.go:type=*Impls
	T0 interface {
		Int() int
	}
	// +zz:impl:./types.go:type=Impls
	T2 interface {
		Int2(context.Context) map[context.Context]int
	}
	// +zz:impl:./types.go
	T3 interface {
		Int2(context.Context) map[context.Context]int
	}
	// +zz:impl:./types.go:aop
	T4 interface {
		Int2(context.Context) map[context.Context]int
	}

	Impls string
)

func (impls *T3Impl) Int2() {
	panic("not implemented")
}
`

const testImplRetData = `package x

import "context"

type (
	// +zz:impl:./types.go:type=*Impls
	T0 interface {
		Int() int
	}
	// +zz:impl:./types.go:type=Impls
	T2 interface {
		Int2(context.Context) map[context.Context]int
	}
	// +zz:impl:./types.go
	T3 interface {
		Int2(context.Context) map[context.Context]int
	}
	// +zz:impl:./types.go:aop
	T4 interface {
		Int2(context.Context) map[context.Context]int
	}

	Impls string
)

func (impls *T3Impl) Int2(context.Context) map[context.Context]int {
	panic("not implemented")
}

func (impls *Impls) Int() int {
	panic("not implemented")
}

func (impls Impls) Int2(context.Context) map[context.Context]int {
	panic("not implemented")
}

var (
	_ T3 = (*T3Impl)(nil)
)

type T3Impl struct{}

var (
	_ T4 = (*T4Impl)(nil)
)

type T4Impl struct{}

func (t4impl *T4Impl) Int2(context.Context) map[context.Context]int {
	panic("not implemented")
}
`

func TestImpl(t *testing.T) {
	_ = os.MkdirAll("test", 0o775)
	defer os.RemoveAll("test")
	if err := os.WriteFile(filepath.Join("test", "types.go"), []byte(testImplData), 0o664); err != nil {
		t.Fatal(err)
	}
	decls, err := zcore.ParseFileOrDirectory(filepath.Join("test", "types.go"), zcore.AnnotationPrefix)
	if err != nil {
		return
	}
	plugin := &Impl{}
	if err = plugin.Run(decls.Parse(plugin, nil)); err != nil {
		t.Fatal(err)
	}
	data, err := ioutil.ReadFile(filepath.Join("test", "types.go"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(data, []byte(testImplRetData)) {
		t.Fatal(err)
	}
}
