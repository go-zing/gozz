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
	"os"
	"path/filepath"
	"testing"

	zcore "github.com/go-zing/gozz-core"
)

const testWireData = `
package x

import (
	"bytes"
	"database/sql"
	"io"

	"github.com/google/wire"
)

// provide value and interface value
// +zz:wire:bind=io.Writer:aop
// +zz:wire
var Buffer = &bytes.Buffer{}

// provide referenced type
// +zz:wire
type NullString nullString

type nullString sql.NullString

// use provider function to provide referenced type alias
// +zz:wire
type String = string

func ProvideString() String {
	return ""
}

// provide value from implicit type
// +zz:wire
var Bool = false

// +zz:wire:inject=./
type Target struct {
	Buffer     *bytes.Buffer
	Writer     io.Writer
	NullString NullString
	Int        int
	Float      Float
	Bytes      *Bytes
}

type Float interface{}

// +zz:wire:bind=Float
func InitFloat() float64 {
	return 0
}

type Bytes []byte

// +zz:wire
func InitBytes() *Bytes {
	return nil
}

// origin wire set
// +zz:wire
var Set = wire.NewSet(wire.Value(Int))

var Int = 0

// mock set injector
// +zz:wire:inject=./inner:set=mock
type MockString sql.NullString

// mock set string
// provide type from function
// +zz:wire:set=mock
func InitMockString() String {
	return "mock"
}

// mock set struct type provide fields
// +zz:wire:set=mock:field=*
type MockConfig struct{ Bool bool }

// mock set value
// +zz:wire:set=mock
var MockConfigValue = &MockConfig{Bool: true}
`

func TestWire(t *testing.T) {
	_ = os.MkdirAll("test", 0o775)
	defer os.RemoveAll("test")
	if err := os.WriteFile(filepath.Join("test", "types.go"), []byte(testWireData), 0o664); err != nil {
		t.Fatal(err)
	}
	decls, err := zcore.ParseFileOrDirectory(filepath.Join("test", "types.go"), zcore.AnnotationPrefix)
	if err != nil {
		return
	}
	plugin := Wire{}
	if err = plugin.Run(decls.Parse(plugin, nil)); err != nil {
		t.Fatal(err)
	}
}
