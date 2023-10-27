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
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseFunc(t *testing.T, str string) *funcType {
	t.Helper()
	v, err := parser.ParseFile(token.NewFileSet(), "", fmt.Sprintf(`package t;func t%s{}`, str), parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	ft, ok := v.Scope.Lookup("t").Decl.(*ast.FuncDecl)
	if !ok {
		t.Fatalf("%T", ft)
	}
	return (*funcType)(ft.Type)
}

func TestFuncTypeReturns(t *testing.T) {
	cases := [][4]interface{}{
		{"()(int,error)", true, "return", ""},
		{"()(int,int)", false, "", ""},
		{"()(int)", true, "return", ",nil"},
		{"()(error)", true, "return nil,", ""},
		{"()()", true, "", ";return nil,nil"},
		{"()(int,int,int)", false, "", ""},
	}

	for _, c := range cases {
		ft := parseFunc(t, c[0].(string))
		r, r2, valid := ft.returns()
		if valid != c[1] || r != c[2] || r2 != c[3] {
			t.Fatal(c, valid, r, r2)
		}
	}
}

func TestFuncTypeParams(t *testing.T) {
	cases := [][4]interface{}{
		{"(context.Context,int)", true, "ctx,in", paramDecode},
		{"(context.Context,*int)", true, "ctx,&in", paramDecode},
		{"(int)", true, "in", paramDecode},
		{"(*int)", true, "&in", paramDecode},
		{"(context.Context)", true, "ctx", ""},
		{"(*context.Context)", true, "&in", paramDecode},
		{"()", true, "", ""},
		{"(int,int)", false, "", ""},
		{"(context.Context,int,int)", false, "", ""},
	}

	for _, c := range cases {
		ft := parseFunc(t, c[0].(string))
		param, decode, _, valid := ft.params()
		if valid != c[1] || param != c[2] || decode != c[3] {
			t.Fatal(c, valid, param, decode)
		}
	}
}
