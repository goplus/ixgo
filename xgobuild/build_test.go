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

package xgobuild

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/bytes"
	_ "github.com/goplus/ixgo/xgobuild/pkg/dql"
	_ "github.com/goplus/ixgo/xgobuild/pkg/encoding"
	_ "github.com/goplus/ixgo/xgobuild/pkg/gsh"
	_ "github.com/goplus/ixgo/xgobuild/pkg/tpl"
)

func gopClTest(t *testing.T, gopcode, expected string) {
	gopClTestEx(t, "main.xgo", gopcode, expected)
}

func gopClTestEx(t *testing.T, filename string, gopcode, expected string) {
	gopClTestWith(t, ixgo.NewContext(0), filename, gopcode, expected)
}

func gopClTestWith(t *testing.T, ctx *ixgo.Context, filename string, gopcode, expected string) {
	data, err := BuildFile(ctx, filename, gopcode)
	if err != nil {
		t.Fatalf("build xgo error: %v", err)
	}
	if string(data) != expected {
		fmt.Println("build xgo error:")
		fmt.Println(string(data))
		t.Fail()
	}
}

func TestXGo(t *testing.T) {
	gopClTest(t, `
echo "XGo"
`, `package main

import "fmt"
//line main.xgo:2
func main() {
//line main.xgo:2:1
	fmt.Println("XGo")
}
`)
}

func TestXGoInit(t *testing.T) {
	gopClTestEx(t, `Rect.gox`, `
var (
	w,h = 10,20
	color int
	border float64 = 1.2
)

echo *this
`, `package main

import "fmt"

type Rect struct {
	w      int
	h      int
	color  int
	border float64
}
//line Rect.gox:8
func (this *Rect) Main() {
	this.XGo_Init()
//line Rect.gox:8:1
	fmt.Println(*this)
}
func (this *Rect) XGo_Init() *Rect {
	this.w, this.h = 10, 20
	this.border = 1.2
	return this
}
func main() {
	new(Rect).Main()
}
`)
}

func TestGox(t *testing.T) {
	gopClTestEx(t, "Rect.gox", `
println "Go+"
`, `package main

import "fmt"

type Rect struct {
}
//line Rect.gox:2
func (this *Rect) Main() {
//line Rect.gox:2:1
	fmt.Println("Go+")
}
func main() {
	new(Rect).Main()
}
`)
	gopClTestEx(t, "Rect.gox", `
var (
	Buffer
	v int
)
type Buffer struct {
	buf []byte
}
println "Go+"
`, `package main

import "fmt"

type Buffer struct {
	buf []byte
}
type Rect struct {
	Buffer
	v int
}
//line Rect.gox:9
func (this *Rect) Main() {
//line Rect.gox:9:1
	fmt.Println("Go+")
}
func main() {
	new(Rect).Main()
}
`)
	gopClTestEx(t, "Rect.gox", `
var (
	*Buffer
	v int
)
type Buffer struct {
	buf []byte
}
println "Go+"
`, `package main

import "fmt"

type Buffer struct {
	buf []byte
}
type Rect struct {
	*Buffer
	v int
}
//line Rect.gox:9
func (this *Rect) Main() {
//line Rect.gox:9:1
	fmt.Println("Go+")
}
func main() {
	new(Rect).Main()
}
`)
	gopClTestEx(t, "Rect.gox", `
import "bytes"
var (
	*bytes.Buffer
	v int
)
println "Go+"
`, `package main

import (
	"bytes"
	"fmt"
)

type Rect struct {
	*bytes.Buffer
	v int
}
//line Rect.gox:7
func (this *Rect) Main() {
//line Rect.gox:7:1
	fmt.Println("Go+")
}
func main() {
	new(Rect).Main()
}
`)
	gopClTestEx(t, "Rect.gox", `
import "bytes"
var (
	bytes.Buffer
	v int
)
println "Go+"
`, `package main

import (
	"bytes"
	"fmt"
)

type Rect struct {
	bytes.Buffer
	v int
}
//line Rect.gox:7
func (this *Rect) Main() {
//line Rect.gox:7:1
	fmt.Println("Go+")
}
func main() {
	new(Rect).Main()
}
`)
}

func TestBig(t *testing.T) {
	gopClTest(t, `
a := 1/2r
echo a+1/2r
`, `package main

import (
	"fmt"
	"github.com/qiniu/x/xgo/ng"
	"math/big"
)
//line main.xgo:2
func main() {
//line main.xgo:2:1
	a := ng.Bigrat_Init__2(big.NewRat(1, 2))
//line main.xgo:3:1
	fmt.Println((ng.Bigrat).XGo_Add(a, ng.Bigrat_Init__2(big.NewRat(1, 2))))
}
`)
}

func TestBuiltin(t *testing.T) {
	ixgo.RegisterCustomBuiltin("typeof", reflect.TypeOf)
	gopClTest(t, `
v := typeof(100)
echo(v)
`, `package main

import "fmt"
//line main.xgo:2
func main() {
//line main.xgo:2:1
	v := typeof(100)
//line main.xgo:3:1
	fmt.Println(v)
}
`)
}

func TestIoxLines(t *testing.T) {
	gopClTest(t, `
import "io"

var r io.Reader

for line <- lines(r) {
	println line
}
`, `package main

import (
	"fmt"
	"github.com/qiniu/x/osx"
	"io"
)

var r io.Reader
//line main.xgo:6
func main() {
//line main.xgo:6:1
	for _xgo_it := osx.Lines(r).XGo_Enum(); ; {
		var _xgo_ok bool
		line, _xgo_ok := _xgo_it.Next()
		if !_xgo_ok {
			break
		}
//line main.xgo:7:1
		fmt.Println(line)
	}
}
`)
}

func TestErrorWrap(t *testing.T) {
	gopClTest(t, `
import (
    "strconv"
)

func add(x, y string) (int, error) {
    return strconv.Atoi(x)? + strconv.Atoi(y)?, nil
}

func addSafe(x, y string) int {
    return strconv.Atoi(x)?:0 + strconv.Atoi(y)?:0
}

echo add("100", "23")!

sum, err := add("10", "abc")
echo sum, err

echo addSafe("10", "abc")
`, `package main

import (
	"fmt"
	"github.com/qiniu/x/errors"
	"strconv"
)
//line main.xgo:6:1
func add(x string, y string) (int, error) {
//line main.xgo:7:1
	var _autoGo_1 int
//line main.xgo:7:1
	{
//line main.xgo:7:1
		var _xgo_err error
//line main.xgo:7:1
		_autoGo_1, _xgo_err = strconv.Atoi(x)
//line main.xgo:7:1
		if _xgo_err != nil {
//line main.xgo:7:1
			_xgo_err = errors.NewFrame(_xgo_err, "strconv.Atoi(x)", "main.xgo", 7, "main.add")
//line main.xgo:7:1
			return 0, _xgo_err
		}
//line main.xgo:7:1
		goto _autoGo_2
	_autoGo_2:
//line main.xgo:7:1
	}
//line main.xgo:7:1
	var _autoGo_3 int
//line main.xgo:7:1
	{
//line main.xgo:7:1
		var _xgo_err error
//line main.xgo:7:1
		_autoGo_3, _xgo_err = strconv.Atoi(y)
//line main.xgo:7:1
		if _xgo_err != nil {
//line main.xgo:7:1
			_xgo_err = errors.NewFrame(_xgo_err, "strconv.Atoi(y)", "main.xgo", 7, "main.add")
//line main.xgo:7:1
			return 0, _xgo_err
		}
//line main.xgo:7:1
		goto _autoGo_4
	_autoGo_4:
//line main.xgo:7:1
	}
//line main.xgo:7:1
	return _autoGo_1 + _autoGo_3, nil
}
//line main.xgo:10:1
func addSafe(x string, y string) int {
//line main.xgo:11:1
	return func() (_xgo_ret int) {
//line main.xgo:11:1
		var _xgo_err error
//line main.xgo:11:1
		_xgo_ret, _xgo_err = strconv.Atoi(x)
//line main.xgo:11:1
		if _xgo_err != nil {
//line main.xgo:11:1
			return 0
		}
//line main.xgo:11:1
		return
	}() + func() (_xgo_ret int) {
//line main.xgo:11:1
		var _xgo_err error
//line main.xgo:11:1
		_xgo_ret, _xgo_err = strconv.Atoi(y)
//line main.xgo:11:1
		if _xgo_err != nil {
//line main.xgo:11:1
			return 0
		}
//line main.xgo:11:1
		return
	}()
}
//line main.xgo:14
func main() {
//line main.xgo:14:1
	fmt.Println(func() (_xgo_ret int) {
//line main.xgo:14:1
		var _xgo_err error
//line main.xgo:14:1
		_xgo_ret, _xgo_err = add("100", "23")
//line main.xgo:14:1
		if _xgo_err != nil {
//line main.xgo:14:1
			_xgo_err = errors.NewFrame(_xgo_err, "add(\"100\", \"23\")", "main.xgo", 14, "main.main")
//line main.xgo:14:1
			panic(_xgo_err)
		}
//line main.xgo:14:1
		return
	}())
//line main.xgo:16:1
	sum, err := add("10", "abc")
//line main.xgo:17:1
	fmt.Println(sum, err)
//line main.xgo:19:1
	fmt.Println(addSafe("10", "abc"))
}
`)
}

func TestGsh(t *testing.T) {
	gopClTestEx(t, "exec.gsh", `
gop "run", "./foo"
exec "gop run ./foo"
exec "FOO=100 gop run ./foo"
exec {"FOO": "101"}, "gop", "run", "./foo"
exec "gop", "run", "./foo"
exec "ls $HOME"
ls "${HOME}"

`, `package main

import "github.com/qiniu/x/gsh"

type exec struct {
	gsh.App
}
//line exec.gsh:2
func (this *exec) MainEntry() {
//line exec.gsh:2:1
	this.XGo_Exec("gop", "run", "./foo")
//line exec.gsh:3:1
	this.Exec__1("gop run ./foo")
//line exec.gsh:4:1
	this.Exec__1("FOO=100 gop run ./foo")
//line exec.gsh:5:1
	this.Exec__0(map[string]string{"FOO": "101"}, "gop", "run", "./foo")
//line exec.gsh:6:1
	this.Exec__2("gop", "run", "./foo")
//line exec.gsh:7:1
	this.Exec__1("ls $HOME")
//line exec.gsh:8:1
	this.XGo_Exec("ls", this.XGo_Env("HOME"))
}
func (this *exec) Main() {
	gsh.XGot_App_Main(this)
}
func main() {
	new(exec).Main()
}
`)
}

var patch_data = `package gsh

import (
	"fmt"
	"github.com/qiniu/x/gsh"
)

const (
	Version = "ver@patch"
)

var Zero = &Point{0,0}

type MyPoint = Point

type Point struct {
	X int
	Y int
}

func (p *Point) Info() {
	println(p.X, p.Y)
}

type Info interface {
	Info()
}

func Dump(i Info) {
	i.Info()
}

func XGot_App_XGox_GetWidget[T any](app any, name string) {
	var _ gsh.App
	fmt.Printf("%T, %T, %v\n", app, (*T)(nil), name)
}
`

func TestPackagePatch(t *testing.T) {
	ctx := ixgo.NewContext(0)
	err := ctx.RegisterPatch("github.com/qiniu/x/gsh", patch_data)
	if err != nil {
		t.Fatal(err)
	}
	gopClTestWith(t, ctx, "main.gsh", `
getWidget(int,"info")
pt := &Point{100,200}
pt.Info()
println(pt.X)
dump(pt)

// var
dump(Zero)
// const
println(Version)
// alias
pt2 := &MyPoint{1,2}
dump(pt2)
`, `package main

import (
	"fmt"
	"github.com/qiniu/x/gsh"
	gsh1 "github.com/qiniu/x/gsh@patch"
)

type App struct {
	gsh.App
}
//line main.gsh:2
func (this *App) MainEntry() {
//line main.gsh:2:1
	gsh1.XGot_App_XGox_GetWidget[int](this, "info")
//line main.gsh:3:1
	pt := &gsh1.Point{100, 200}
//line main.gsh:4:1
	pt.Info()
//line main.gsh:5:1
	fmt.Println(pt.X)
//line main.gsh:6:1
	gsh1.Dump(pt)
//line main.gsh:9:1
	gsh1.Dump(gsh1.Zero)
//line main.gsh:11:1
	fmt.Println(gsh1.Version)
//line main.gsh:13:1
	pt2 := &gsh1.MyPoint{1, 2}
//line main.gsh:14:1
	gsh1.Dump(pt2)
}
func (this *App) Main() {
	gsh.XGot_App_Main(this)
}
func main() {
	new(App).Main()
}
`)
}

func TestPackagePatchNormalize(t *testing.T) {
	ctx := ixgo.NewContext(NormalizeExport)
	err := ctx.RegisterPatch("github.com/qiniu/x/gsh", patch_data)
	if err != nil {
		t.Fatal(err)
	}
	gopClTestWith(t, ctx, "main.gsh", `
getWidget(int,"info")
pt := &Point{100,200}
pt.Info()
println(pt.X)
dump(pt)

// var
dump(Zero)
// const
println(Version)
// alias
pt2 := &MyPoint{1,2}
dump(pt2)
`, `package main

import (
	"fmt"
	"github.com/qiniu/x/gsh"
)

type App struct {
	gsh.App
}
//line main.gsh:2
func (this *App) MainEntry() {
//line main.gsh:2:1
	gsh.XGot_App_XGox_GetWidget[int](this, "info")
//line main.gsh:3:1
	pt := &gsh.Point{100, 200}
//line main.gsh:4:1
	pt.Info()
//line main.gsh:5:1
	fmt.Println(pt.X)
//line main.gsh:6:1
	gsh.Dump(pt)
//line main.gsh:9:1
	gsh.Dump(gsh.Zero)
//line main.gsh:11:1
	fmt.Println(gsh.Version)
//line main.gsh:13:1
	pt2 := &gsh.MyPoint{1, 2}
//line main.gsh:14:1
	gsh.Dump(pt2)
}
func (this *App) Main() {
	gsh.XGot_App_Main(this)
}
func main() {
	new(App).Main()
}
`)
}

func TestPackagePatchNoUse(t *testing.T) {
	ctx := ixgo.NewContext(0)
	err := ctx.RegisterPatch("github.com/qiniu/x/gsh", patch_data)
	if err != nil {
		t.Fatal(err)
	}
	gopClTestWith(t, ctx, "main.gsh", `
ls "${HOME}"
`, `package main

import "github.com/qiniu/x/gsh"

type App struct {
	gsh.App
}
//line main.gsh:2
func (this *App) MainEntry() {
//line main.gsh:2:1
	this.XGo_Exec("ls", this.XGo_Env("HOME"))
}
func (this *App) Main() {
	gsh.XGot_App_Main(this)
}
func main() {
	new(App).Main()
}
`)
}

var (
	gshDemo = `
getWidget(int,"info")
pt := &Point{100,200}
pt.Info()
println(pt.X)
dump(pt)

// var
dump(Zero)
// const
println(Version)
// alias
pt2 := &MyPoint{1,2}
dump(pt2)
`
)

func TestPackagePatchRun(t *testing.T) {
	ctx := ixgo.NewContext(0)
	err := ctx.RegisterPatch("github.com/qiniu/x/gsh", patch_data)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ctx.RunFile("main.gsh", gshDemo, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPackagePatchNormalizeRun(t *testing.T) {
	ctx := ixgo.NewContext(NormalizeExport)
	err := ctx.RegisterPatch("github.com/qiniu/x/gsh", patch_data)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ctx.RunFile("main.gsh", gshDemo, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStringSlice(t *testing.T) {
	gopClTest(t, `
println "a".repeat(10)
println "hello".toTitle
`, `package main

import (
	"fmt"
	"strings"
)
//line main.xgo:2
func main() {
//line main.xgo:2:1
	fmt.Println(strings.Repeat("a", 10))
//line main.xgo:3:1
	fmt.Println(strings.ToTitle("hello"))
}
`)
}

func TestEcho(t *testing.T) {
	gopClTest(t, `
echo type(100)
`, `package main

import (
	"fmt"
	"reflect"
)
//line main.xgo:2
func main() {
//line main.xgo:2:1
	fmt.Println(reflect.TypeOf(100))
}
`)
}

var (
	tplDemo = `import "gop/tpl"

cl := tpl` + "`" + `
expr = INT % "," => {
    return tpl.ListOp[int](self, v => {
        return v.(*tpl.Token).Lit.int!
    })
}
` + "`" + `!

echo cl.parseExpr("1, 2, 3", nil)!
`
	tplGo = `package main

import (
	"fmt"
	"github.com/goplus/xgo/tpl"
	"github.com/goplus/xgo/tpl/types"
	"github.com/qiniu/x/errors"
	"strconv"
)
//line main.xgo:3
func main() {
//line main.xgo:3:1
	cl := func() (_xgo_ret tpl.Compiler) {
//line main.xgo:3:1
		var _xgo_err error
//line main.xgo:3:1
		_xgo_ret, _xgo_err = tpl.NewEx(` + "`" + `
expr = INT % "," => {
    return tpl.ListOp[int](self, v => {
        return v.(*tpl.Token).Lit.int!
    })
}
` + "`" + `, "main.xgo", 3, 10, "expr", func(self []interface{}) interface{} {
//line main.xgo:5:1
			return tpl.ListOp[int](self, func(v any) int {
//line main.xgo:6:1
				return func() (_xgo_ret int) {
//line main.xgo:6:1
					var _xgo_err error
//line main.xgo:6:1
					_xgo_ret, _xgo_err = strconv.Atoi(v.(*types.Token).Lit)
//line main.xgo:6:1
					if _xgo_err != nil {
//line main.xgo:6:1
						_xgo_err = errors.NewFrame(_xgo_err, "v.(*tpl.Token).Lit.int", "main.xgo", 6, "main.main")
//line main.xgo:6:1
						panic(_xgo_err)
					}
//line main.xgo:6:1
					return
				}()
			})
		})
//line main.xgo:3:1
		if _xgo_err != nil {
//line main.xgo:3:1
			_xgo_err = errors.NewFrame(_xgo_err, "tpl` + "`" + `\nexpr = INT % \",\" => {\n    return tpl.ListOp[int](self, v => {\n        return v.(*tpl.Token).Lit.int!\n    })\n}\n` + "`" + `", "main.xgo", 3, "main.main")
//line main.xgo:3:1
			panic(_xgo_err)
		}
//line main.xgo:3:1
		return
	}()
//line main.xgo:11:1
	fmt.Println(func() (_xgo_ret interface{}) {
//line main.xgo:11:1
		var _xgo_err error
//line main.xgo:11:1
		_xgo_ret, _xgo_err = cl.ParseExpr("1, 2, 3", nil)
//line main.xgo:11:1
		if _xgo_err != nil {
//line main.xgo:11:1
			_xgo_err = errors.NewFrame(_xgo_err, "cl.parseExpr(\"1, 2, 3\", nil)", "main.xgo", 11, "main.main")
//line main.xgo:11:1
			panic(_xgo_err)
		}
//line main.xgo:11:1
		return
	}())
}
`
)

func TestTplPatch(t *testing.T) {
	ctx := ixgo.NewContext(NormalizeExport)
	gopClTestWith(t, ctx, "main.xgo", tplDemo, tplGo)
}

func TestTplPatchRun(t *testing.T) {
	ctx := ixgo.NewContext(0)
	_, err := ctx.RunFile("main.xgo", tplDemo, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestTplPatchStaticLoad(t *testing.T) {
	ctx := ixgo.NewContext(StaticLoad | NormalizeExport)
	bctx := NewContext(ctx)
	if bctx.Importer != nil {
		t.Fatalf("bad static load")
	}
	for i := 0; i < 10; i++ {
		pkg, err := bctx.ParseFile("main.xgo", tplDemo)
		if err != nil {
			t.Fatal(err)
		}
		data, err := pkg.ToSource()
		if err != nil {
			t.Fatal(err)
		}
		if string(data) != tplGo {
			t.Fatalf("bad result:\n%v", tplGo)
		}
		_, err = ctx.RunFile("main.go", data, nil)
		if err != nil {
			t.Fatal(err)
		}
	}
}

var dqlXGo = `
doc := xgo` + "`" + `
x, y := "Hi", 123
echo x
print y
` + "`" + `!

stmts := doc.shadowEntry.body.list.*@(self.class == "ExprStmt")
for fn in stmts.x@(self.class == "CallExpr").fun@(self.class == "Ident") {
	echo fn.$name
}
`

func TestDqlXGo(t *testing.T) {
	ctx := ixgo.NewContext(0)
	_, err := ctx.RunFile("main.xgo", dqlXGo, nil)
	if err != nil {
		t.Fatal(err)
	}
}
