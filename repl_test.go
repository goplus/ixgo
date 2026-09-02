/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ixgo_test

import (
	"fmt"
	"go/token"
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/fmt"
	_ "github.com/goplus/ixgo/pkg/log"
	_ "github.com/goplus/ixgo/pkg/runtime"
)

func TestReplExpr(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`a+b`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[3 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplRuntimeVersionThenLogPrintln(t *testing.T) {
	r := ixgo.NewRepl(ixgo.NewContext(0))
	for _, expr := range []string{
		`a := 100`,
		`import "runtime"`,
		`import ("fmt";"runtime")`,
		`fmt.Println(runtime.Version())`,
		`import "log"`,
		`log.Println("hello")`,
		`import stdlog "log"`,
		`stdlog.Println("world")`,
	} {
		if _, _, err := r.Eval(expr); err != nil {
			t.Fatalf("Eval(%q): %v", expr, err)
		}
	}
}

func TestReplImportPreservesMainGlobals(t *testing.T) {
	repl := ixgo.NewRepl(ixgo.NewContext(0))
	list := []string{
		`var a = 1`,
		`import "runtime"`,
		`var b = 2`,
		`import "log"`,
		`var c = 3`,
		`a`,
		`b`,
		`c`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[1 int]`,
		`[2 int]`,
		`[3 int]`,
	}
	for i, expr := range list {
		_, got, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(got) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, got, repl.Source())
		}
	}
}

func TestReplImports(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`import ("fmt"; "runtime")`,
		`c := fmt.Sprintf("%v-%v",a,b)`,
		`c`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[1-2 string]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplBlankImportThenNamedImport(t *testing.T) {
	repl := ixgo.NewRepl(ixgo.NewContext(0))
	list := []string{
		`import _ "fmt"`,
		`import fmt "fmt"`,
		`fmt.Sprintf("%v", 100)`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[100 string]`,
	}
	for i, expr := range list {
		_, got, err := repl.Eval(expr)
		if err != nil {
			t.Fatalf("Eval(%q): %v", expr, err)
		}
		if fmt.Sprint(got) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, got, repl.Source())
		}
	}
}

func TestReplImportRunsPackageInit(t *testing.T) {
	ixgo.RegisterPackage(&ixgo.Package{
		Name:   "replinit",
		Path:   "test/replinit",
		Source: "package replinit\nvar Value = 0\nfunc init() { Value = 42 }\n",
	})
	repl := ixgo.NewRepl(ixgo.NewContext(0))
	for _, expr := range []string{
		`import "test/replinit"`,
		`replinit.Value`,
	} {
		_, got, err := repl.Eval(expr)
		if err != nil {
			t.Fatalf("Eval(%q): %v", expr, err)
		}
		if expr != `replinit.Value` && fmt.Sprint(got) != `[]` {
			t.Fatalf("expr:%v dump:%v", expr, got)
		}
		if expr == `replinit.Value` && fmt.Sprint(got) != `[42 int]` {
			t.Fatalf("expr:%v dump:%v", expr, got)
		}
	}
}

func TestReplClosure(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`fn := func() int { return a }`,
		`d := fn()+b`,
		`d`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[3 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplVar(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`a := 1`,
		`b := 2`,
		`var c = 100`,
		`fn := func() int { return a+c }`,
		`d := fn()+b`,
		`d`,
		`c`,
	}
	result := []string{
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[]`,
		`[103 int]`,
		`[100 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplType(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`type T struct {
	X int
	Y int
}`,
		`v1 := &T{10,20}`,
		`import "fmt"`,
		`v1`,
		`func (t *T) String() string {
	return fmt.Sprintf("%v-%v",t.X,t.Y)
}`,
		`v2 := &T{10,20}`,
		`v2`,
	}
	result := []string{
		`-`,
		`-`,
		`-`,
		`[&{10 20} *main.T]`,
		`-`,
		`-`,
		`[10-20 *main.T]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplFunc(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`a := "hello"`,
		`import "fmt"`,
		`a`,
		`fmt.Println(a)`,
		`s := fmt.Sprint(a)`,
		`s`,
		`fmt.Println(s)`,
	}
	result := []string{
		`[]`,
		`-`,
		`[hello string]`,
		`[6 int <nil> error]`,
		`-`,
		`[hello string]`,
		`[6 int <nil> error]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplTok(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`a := "hello"`,
		`import "fmt"`,
		`var c int`,
		`const d = 100`,
		`func test(v int) {}`,
		`fmt.Println(a)`,
		`type T int`,
	}
	toks := []token.Token{
		token.IDENT,
		token.IMPORT,
		token.VAR,
		token.CONST,
		token.FUNC,
		token.IDENT,
		token.TYPE,
	}
	for i, expr := range list {
		tok, _, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if tok != toks[i] {
			t.Fatalf("expr:%v token:%v", expr, tok)
		}
	}
}

func TestReplInit(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`var i int = 10`,
		`var j int = 20`,
		`func init() { i++ }`,
		`func init() { j++ }`,
		`var c int`,
		`i`,
		`j`,
	}
	result := []string{
		`-`,
		`-`,
		`-`,
		`-`,
		`-`,
		`[11 int]`,
		`[21 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplFunLit(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`var a int`,
		`func(x int){ a = x }(100)`,
		`a`,
	}
	result := []string{
		`-`,
		`-`,
		`[100 int]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if result[i] == "-" {
			continue
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}

func TestReplFor(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`import "fmt"`,
		`for i := 0; i < 3; i++ { fmt.Println(i) }`,
		`fmt.Println(100)`,
	}
	for _, expr := range list {
		_, _, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestReplIf(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`import "fmt"`,
		`a := 2`,
		`if a == 1 { fmt.Println(1) } else { fmt.Println(a) }`,
		`fmt.Println(100)`,
	}
	for _, expr := range list {
		_, _, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestReplSwitch(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`import "fmt"`,
		`a := 2`,
		`switch a {
		case 1: 
			if a == 1 { fmt.Println(1) }
		default: 
			fmt.Println(a)
		}`,
		`fmt.Println(100)`,
	}
	for _, expr := range list {
		_, _, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestReplSelect(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`import "fmt"`,
		`a := 1`,
		`ch := make(chan int)`,
		`select {
		case ch <- 1:
			if a == 1 { fmt.Println(1) }
			fmt.Println(1)
		default:
			fmt.Println(ch)
		}`,
		`fmt.Println(100)`,
	}
	for _, expr := range list {
		_, _, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestReplLiteral(t *testing.T) {
	ctx := ixgo.NewContext(0)
	repl := ixgo.NewRepl(ctx)
	list := []string{
		`100`,
		`100.0*200`,
		`(100*200)`,
		`map[int]string{1:"100"}`,
		`1e100<<2`,
		`+100*2`,
		`-100/2`,
		`^100`,
		`"hello"+"world"`,
	}
	result := []string{
		`[100 int]`,
		`[20000 float64]`,
		`[20000 int]`,
		`[map[1:100] map[int]string]`,
		`[40000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000 untyped int]`,
		`[200 int]`,
		`[-50 int]`,
		`[-101 int]`,
		`[helloworld string]`,
	}
	for i, expr := range list {
		_, v, err := repl.Eval(expr)
		if err != nil {
			t.Fatal(err)
		}
		if fmt.Sprint(v) != result[i] {
			t.Fatalf("expr:%v dump:%v src:%v", expr, v, repl.Source())
		}
	}
}
