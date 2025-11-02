// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ixgo_test

// This test runs the SSA interpreter over sample Go programs.
// Because the interpreter requires intrinsics for assembly
// functions and many low-level runtime routines, it is inherently
// not robust to evolutionary change in the standard library.
// Therefore the test cases are restricted to programs that
// use a fake standard library in testdata/src containing a tiny
// subset of simple functions useful for writing assertions.
//
// We no longer attempt to interpret any real standard packages such as
// fmt or testing, as it proved too fragile.

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/testdata/info"
	"github.com/goplus/ixgo/testdata/pkg"

	_ "github.com/goplus/ixgo/pkg/bytes"
	_ "github.com/goplus/ixgo/pkg/errors"
	_ "github.com/goplus/ixgo/pkg/fmt"
	_ "github.com/goplus/ixgo/pkg/math"
	_ "github.com/goplus/ixgo/pkg/os"
	_ "github.com/goplus/ixgo/pkg/path/filepath"
	_ "github.com/goplus/ixgo/pkg/reflect"
	_ "github.com/goplus/ixgo/pkg/runtime"
	_ "github.com/goplus/ixgo/pkg/strings"
	_ "github.com/goplus/ixgo/pkg/sync"
	_ "github.com/goplus/ixgo/pkg/time"
)

// These are files in github.com/goplus/ixgo/testdata/.
var testdataTests = []string{
	"boundmeth.go",
	"complit.go",
	"coverage.go",
	"defer.go",
	"fieldprom.go",
	"ifaceconv.go",
	"ifaceprom.go",
	"initorder.go",
	"methprom.go",
	"mrvchain.go",
	"range.go",
	"recover.go",
	"reflect.go",
	"static.go",
	"recover2.go",
	"static.go",
	"issue23536.go",
	"tinyfin.go",
	"issue5963.go",
}

func runInput(t *testing.T, input string) bool {
	fmt.Println("Input:", input)
	start := time.Now()
	_, err := ixgo.Run(input, nil, 0)
	sec := time.Since(start).Seconds()
	if err != nil {
		t.Error(err)
		fmt.Printf("FAIL %0.3fs\n", sec)
		return false
	}
	fmt.Printf("PASS %0.3fs\n", sec)
	return true
}

// TestTestdataFiles runs the interpreter on testdata/*.go.
func TestTestdataFiles(t *testing.T) {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var failures []string
	for _, input := range testdataTests {
		if !runInput(t, filepath.Join(cwd, "testdata", input)) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
}

func printFailures(failures []string) {
	if failures != nil {
		fmt.Println("The following tests failed:")
		for _, f := range failures {
			fmt.Printf("\t%s\n", f)
		}
	}
}

func TestUntypedNil(t *testing.T) {
	src := `package main

type T func()

func main() {
	if T(nil) != nil {
		panic("error")
	}
	if !(T(nil) == nil) {
		panic("error")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterExternal(t *testing.T) {
	ctx := ixgo.NewContext(0)
	ctx.RegisterExternal("main.call", func(i, j int) int {
		return i * j
	})
	var v int = 200
	ctx.RegisterExternal("main.v", &v)
	src := `package main

func call(i, j int) int {
	return i+j
}

var v int

func main() {
	if n := call(10,20); n != 200 {
		panic(n)
	}
	if v != 200 {
		panic(v)
	}
}
`
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}

	// reset override func
	ctx.RegisterExternal("main.call", nil)
	_, err = ctx.RunFile("main.go", src, nil)
	if err == nil || err.Error() != "30" {
		t.Fatal("must panic 30")
	}
}

func TestOsExit(t *testing.T) {
	src := `package main

import "os"

func main() {
	os.Exit(-2)
}
`
	code, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	if code != -2 {
		t.Fatalf("exit code %v, must -2", code)
	}
}

func TestOpAlloc(t *testing.T) {
	src := `package main

type T struct {
	n1 int
	n2 int
}

func (t T) call() int {
	return t.n1 * t.n2
}

func main() {
	var n int
	for i := 0; i < 3; i++ {
		n += T{i,3}.call()
	}
	if n != 9 {
		panic(n)
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpBin(t *testing.T) {
	src := `package main

func main() {
	a := 10 // 1010
	b := 12 // 1100
	n1 := a + b
	n2 := a - b
	n3 := a * b
	n4 := a / b
	n5 := a % b
	n6 := a & b
	n7 := a | b
	n8 := a ^ b
	n9 := a &^ b
	n10 := a << 3
	n11 := a >> 3
	v1 := a > b
	v2 := a < b
	v3 := a >= b
	v4 := a <= b
	v5 := a == b
	v6 := a != b
	if n1 != 22 || n2 != -2 || n3 != 120 || n4 != 0 || n5 != 10 || n6 != 8 ||
		n7 != 14 || n8 != 6 || n9 != 2 || n10 != 80 || n11 != 1 ||
		v1 != false || v2 != true || v3 != false || v4 != true || v5 != false || v6 != true {
		println(n1, n2, n3, n4, n5, n6, n7, n8, n9, n10, n11)
		println(v1, v2, v3, v4, v5, v6)
		panic("error")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpChangeType(t *testing.T) {
	src := `package main

type T func(int, int) int

func add(i, j int) int {
	return i + j
}

type I interface{}

func main() {
	// static change type
	fn := T(add)
	if n := fn(10, 20); n != 30 {
		panic(n)
	}
	var i interface{} = I(nil)
	if i != nil {
		panic("must nil")
	}
	// dynamic change type
	var k interface{}
	i = I(k)
	if i != nil {
		panic("must nil")
	}
	i = I(100)
	if i == nil {
		panic("must not nil")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpIndex(t *testing.T) {
	src := `package main

func main() {
	var n int
	for _, v := range [3]int{1, 2, 4} {
		n += v
	}
	if n != 7 {
		panic(n)
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnderscopeMap(t *testing.T) {
	src := `package main

type key struct {
	_ int
	x int
}

var (
	s1 = "s1"
	s2 = "s2"
	s3 = "s3"
	s4 = "s4"
)

func main() {
	m := make(map[key]string)
	m[key{0, 1}] = s1
	m[key{0, 2}] = s2
	m[key{0, 3}] = s3
	m[key{1, 1}] = s4
	m[key{1, 2}] = "s22"
	if n := len(m); n != 3 {
		panic(n)
	}
	if v := m[key{2, 2}]; v != "s22" {
		panic(v)
	}
	if v, ok := m[key{1,4}]; ok {
		panic(v)
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpSend(t *testing.T) {
	src := `package main

import (
	"fmt"
)

type T struct{}

func main() {
	ch := make(chan *T)
	go func() {
		ch <- nil
	}()
	v := <-ch
	if v != nil {
		panic("must nil")
	}
	if s := fmt.Sprintf("%T", v); s != "*main.T" {
		panic(s)
	}
	go func() {
		ch <- &T{}
	}()
	v = <-ch
	if v == nil {
		panic("must not nil")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestShadowedMethod(t *testing.T) {
	src := `// run

// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// When computing method sets with shadowed methods, make sure we
// compute whether a method promotion involved a pointer traversal
// based on the promoted method, not the shadowed method.

package main

import (
	"bytes"
	"fmt"
)

type mystruct struct {
	f int
}

func (t mystruct) String() string {
	return "FAIL"
}

func main() {
	type deep struct {
		mystruct
	}
	s := struct {
		deep
		*bytes.Buffer
	}{
		deep{},
		bytes.NewBufferString("ok"),
	}

	if got := s.String(); got != "ok" {
		panic(got)
	}

	var i fmt.Stringer = s
	if got := i.String(); got != "ok" {
		panic(got)
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestConvertUnsafePointer(t *testing.T) {
	src := `package main

import (
	"unsafe"
)

func main() {
	a := [4]int{0, 1, 2, 3}
	p := unsafe.Pointer(&a)
	p2 := unsafe.Pointer(uintptr(p) + 2*unsafe.Sizeof(a[0]))
	*(*int)(p2) = 4
	if a != [4]int{0, 1, 4, 3} {
		panic("error")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuiltinPrintln(t *testing.T) {
	src := `// run

// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test internal print routines that are generated
// by the print builtin.  This test is not exhaustive,
// we're just checking that the formatting is correct.

package main

func main() {
	println((interface{})(nil)) // printeface
	println((interface {        // printiface
		f()
	})(nil))
	println((map[int]int)(nil)) // printpointer
	println(([]int)(nil))       // printslice
	println(int64(-7))          // printint
	println(uint64(7))          // printuint
	println(uint32(7))          // printuint
	println(uint16(7))          // printuint
	println(uint8(7))           // printuint
	println(uint(7))            // printuint
	println(uintptr(7))         // printuint
	println(8.0)                // printfloat
	println(complex(9.0, 10.0)) // printcomplex
	println(true)               // printbool
	println(false)              // printbool
	println("hello")            // printstring
	println("one", "two")       // printsp

	// test goprintf
	defer println((interface{})(nil))
	defer println((interface {
		f()
	})(nil))
	defer println((map[int]int)(nil))
	defer println(([]int)(nil))
	defer println(int64(-11))
	defer println(uint64(12))
	defer println(uint32(12))
	defer println(uint16(12))
	defer println(uint8(12))
	defer println(uint(12))
	defer println(uintptr(12))
	defer println(13.0)
	defer println(complex(14.0, 15.0))
	defer println(true)
	defer println(false)
	defer println("hello")
	defer println("one", "two")
}
`
	out := `(0x0,0x0)
(0x0,0x0)
0x0
[0/0]0x0
-7
7
7
7
7
7
7
+8.000000e+000
(+9.000000e+000+1.000000e+001i)
true
false
hello
one two
one two
hello
false
true
(+1.400000e+001+1.500000e+001i)
+1.300000e+001
12
12
12
12
12
12
-11
[0/0]0x0
0x0
(0x0,0x0)
(0x0,0x0)
`
	ctx := ixgo.NewContext(0)
	var buf bytes.Buffer
	ctx.SetPrintOutput(&buf)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() != out {
		t.Fatal("error")
	}
}

type panicinfo struct {
	src string
	err string
}

func TestPanicInfo(t *testing.T) {
	infos := []panicinfo{
		{`panic(100)`, `100`},
		{`panic(100.0)`, `+1.000000e+002`},
		{`panic("hello")`, `hello`},
		{`panic((*interface{})(nil))`, `(*interface {}) 0x0`},
		{`type T int; panic(T(100))`, `main.T(100)`},
		{`type T float64; panic(T(100.0))`, `main.T(+1.000000e+002)`},
		{`type T string; panic(T("hello"))`, `main.T("hello")`},
		{`type T struct{}; panic((*T)(nil))`, `(*main.T) 0x0`},
	}
	ctx := ixgo.NewContext(0)
	for _, info := range infos {
		src := fmt.Sprintf("package main;func main(){%v}", info.src)
		code, err := ctx.RunFile("main.go", src, nil)
		if code != 2 || err == nil {
			t.Fatalf("%v must panic", info.src)
		}
		if s := err.Error(); s != info.err {
			t.Fatalf("%v err is %v, want %v", info.src, s, info.err)
		}
	}
}

func TestPanicError(t *testing.T) {
	src := `package main
import "errors"

func test() {
	panic(errors.New("error info"))
}

func main() {
	test()
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	if s := err.Error(); s != "error info" {
		t.Fatalf("error %q", s)
	}
	pe, ok := err.(ixgo.PanicError)
	if !ok {
		t.Fatal("error type must ixgo.PanicError")
	}
	ve, ok := pe.Value.(error)
	if !ok {
		t.Fatal("e.Value type must error")
	}
	if s := ve.Error(); s != "error info" {
		t.Fatalf("e.Value %q", s)
	}
	t.Log("dump panic stack\n", string(pe.Stack()))
}

func TestPanicErrorRecover(t *testing.T) {
	src := `package main
import "errors"
func main() {
	defer func() {
		err := recover()
		if err == nil {
			panic("must error")
		}
		if s := err.(error).Error(); s != "error info" {
			panic(s)
		}
	}()
	panic(errors.New("error info"))
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEnablePrintAny(t *testing.T) {
	src := `package main

type T struct {
	X int
	Y int
}
func main() {
	println([2]int{100,200})
	println(T{100,200})
}
`
	ctx := ixgo.NewContext(0)
	var buf bytes.Buffer
	ctx.SetPrintOutput(&buf)
	_, err := ctx.RunFile("main.go", src, nil)
	if err == nil {
		t.Fatal("must panic")
	}
	if s := err.Error(); s != "illegal types for operand: print\n\t[2]int" {
		t.Fatal(s)
	}
	ctx.Mode |= ixgo.EnablePrintAny
	_, err = ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
	if s := buf.String(); s != "[100 200]\n{100 200}\n" {
		t.Fatal(s)
	}
}

func TestFib(t *testing.T) {
	src := `package main

import (
	"sync"
)

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func main() {
	var wg sync.WaitGroup
	const N = 10
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			if fib(20) != 6765 {
				panic("error")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFibv(t *testing.T) {
	src := `package main

import "sync"

func main() {
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n-2) + fib(n-1)
	}

	var wg sync.WaitGroup
	const N = 10
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			if fib(20) != 6765 {
				panic("error")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFibi(t *testing.T) {
	src := `package main

import "sync"

type I interface {
	fib(i I, n int) int
}

type T struct {
}

func (t *T) fib(i I, n int) int {
	if n < 2 {
		return n
	}
	return i.fib(i, n-2) + i.fib(i, n-1)
}

func fib(n int) int {
	t := &T{}
	return t.fib(t, n)
}

func main() {
	var wg sync.WaitGroup
	const N = 10
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			if fib(20) != 6765 {
				panic("error")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBinOpInt(t *testing.T) {
	// + - * / % & ^ &^ < <= > >=
	tsrc := `package main

import "fmt"

type T $int

func main() {
	// 0101 0110
	test(5, 6)
	testConst()
	testDivideZero()
}

func testConst() {
	var a T = 4
	var b T = 2
	var c T = 4
	check(a+b,6)
	check(a-b,2)
	check(a*b,8)
	check(a/b,2)
	check(a%b,0)
	check(a&b,0)
	check(a|b,6)
	check(a^b,6)
	check(a&^b,4)
	assert(a>b)
	assert(a>=b)
	assert(b<a)
	assert(a<=a)
	assert(a==c)
	assert(a!=b)
}

func testDivideZero() {
	defer func() {
		if r := recover(); r == nil {
			panic("must panic")
		}
	}()
	var a T
	var b T
	_ = a/b
}

func test(a, b T) {
	check(a+1, 6)
	check(1+a, 6)
	check(a+b, 11)
	check(a-1, 4)
	check(6-a, 1)
	check(b-a, 1)
	check(a*2, 10)
	check(2*a, 10)
	check(a*b, 30)
	check(a/2, 2)
	check(T(10)/a, 2)
	check(a/b, 0)
	check(a%2, 1)
	check(10%a, 0)
	check(b%a, 1)
	check(a&4, 4)
	check(1&a, 1)
	check(a&b, 4)
	check(a|4, 5)
	check(2|a, 7)
	check(a|b, 7)
	check(a^2, 7)
	check(7^a, 2)
	check(a^b, 3)
	check(a&^3, 4)
	check(3&^a, 2)
	check(a&^b, 1)
	assert(a < 6)
	assert(4 < a)
	assert(a < b)
	assert(a <= 5)
	assert(4 <= a)
	assert(a <= b)
	assert(a > 4)
	assert(6 > a)
	assert(b > a)
	assert(a >= 5)
	assert(6 >= a)
	assert(b >= a)
	assert(a == 5)
	assert(5 == a)
	assert(a != 6)
	assert(6 != a)
	assert(a != b)
	assert(a+1 == b)
}

func assert(t bool) {
	if !t {
		panic("error")
	}
}

func check(a, b T) {
	if a != b {
		panic(fmt.Errorf("error %v != %v", a, b))
	}
}
`
	types := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}
	for _, s := range types {
		t.Log("test binop xtype", s)
		src := strings.Replace(tsrc, "$int", "="+s, 1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range types {
		t.Log("test binop named", s)
		src := strings.Replace(tsrc, "$int", s, 1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestBinOpFloat(t *testing.T) {
	tsrc := `package main

import "fmt"

type T $float

func main() {
	test(5.0, 6.5)
	testConst()
}

func testConst() {
	var a T = 4.0
	var b T = 2.0
	var c T = 4.0
	check(a+b,6)
	check(a-b,2)
	check(a*b,8)
	check(a/b,2)
	assert(a>b)
	assert(a>=b)
	assert(b<a)
	assert(a<=a)
	assert(a==c)
	assert(a!=b)
}

func test(a, b T) {
	check(a+1, 6)
	check(1+a, 6)
	check(a+b, 11.5)
	check(a-1, 4)
	check(6-a, 1)
	check(b-a, 1.5)
	check(a*2, 10)
	check(2*a, 10)
	check(a*b, 32.5)
	check(a/2, 2.5)
	check(10/a, 2)
	check(a/b, 5.0/6.5)
	assert(a < 6)
	assert(4 < a)
	assert(a < b)
	assert(a <= 5)
	assert(4 <= a)
	assert(a <= b)
	assert(a > 4)
	assert(6 > a)
	assert(b > a)
	assert(a >= 5)
	assert(6 >= a)
	assert(b >= a)
	assert(a == 5)
	assert(5 == a)
	assert(a != 6)
	assert(6 != a)
	assert(a != b)
	assert(a+1.5 == b)
}

func assert(t bool) {
	if !t {
		panic("error")
	}
}

func check(a, b T) {
	if a != b {
		panic(fmt.Errorf("error %v != %v", a, b))
	}
}
`
	types := []string{
		"float32",
		"float64",
	}
	for _, s := range types {
		t.Log("test binop xtype", s)
		src := strings.Replace(tsrc, "$float", "="+s, 1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range types {
		t.Log("test binop named", s)
		src := strings.Replace(tsrc, "$float", s, 1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestBinOpComplex(t *testing.T) {
	tsrc := `package main

import "fmt"

type T $complex

func main() {
	test(1+2i, 3+4i)
	testConst()
}

func testConst() {
	var a T = 4i
	var b T = 2i
	var c T = 4i
	check(a+b,6i)
	check(a-b,2i)
	check(a*b,-8)
	check(a/b,2)
	assert(a == c)
	assert(a != b)
}

func test(a, b T) {
	check(a+1+2i, 2+4i)
	check(1+a, 2+2i)
	check(a+b, 4+6i)
	check(a-1, 2i)
	check(6-a, 5-2i)
	check(b-a, 2+2i)
	check(a*(2+3i), -4+7i)
	check((2-3i)*a, 8+1i)
	check(a*b, -5+10i)
	check(a/2i, 1-0.5i)
	check(10/a, 2-4i)
	check(a/b, 0.44+0.08i)
	assert(a == 1+2i)
	assert(1+2i == a)
	assert(a != 2+2i)
	assert(2+2i != a)
	assert(a != b)
	assert(a+2+2i == b)
}

func assert(t bool) {
	if !t {
		panic("error")
	}
}

func check(a, b T) {
	if a != b {
		panic(fmt.Errorf("error %v != %v", a, b))
	}
}

`
	types := []string{
		"complex64",
		"complex128",
	}
	for _, s := range types {
		t.Log("test binop xtype", s)
		src := strings.Replace(tsrc, "$complex", "="+s, 1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range types {
		t.Log("test binop named", s)
		src := strings.Replace(tsrc, "$complex", s, 1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestBinOpString(t *testing.T) {
	tsrc := `package main

import "fmt"

type T = string

func main() {
	test("go", "ssa", "go")
	testConst()
}

func testConst() {
	var a T = "hello"
	var b T = "world"
	var c T = "hello"
	check(a+b,"helloworld")
	assert(a < b)
	assert(a <= b)
	assert(b > a)
	assert(b >= a)
	assert(a == c)
	assert(a != b)
}

func test(a, b, c T) {
	check(a+"run", "gorun")
	check("run"+a, "rungo")
	check(a+b, "gossa")
	assert(a < "go1")
	assert("ao" < a)
	assert(a < b)
	assert(a <= "go")
	assert("ao" <= a)
	assert(a <= b)
	assert(a > "ao")
	assert("go1" > a)
	assert(b > a)
	assert(a >= "go")
	assert("go1" >= a)
	assert(b >= a)
	assert(a == "go")
	assert("go" == a)
	assert(a != "go1")
	assert("go1" != a)
	assert(a != b)
	assert(a == c)
}

func assert(t bool) {
	if !t {
		panic("error")
	}
}

func check(a, b T) {
	if a != b {
		panic(fmt.Errorf("error %v != %v", a, b))
	}
}
`
	t.Log("test binop string")
	_, err := ixgo.RunFile("main.go", tsrc, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("test binop named string")
	src := strings.Replace(tsrc, "= string", "string", 1)
	_, err = ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBinOpShift(t *testing.T) {
	tsrc := `package main
type T1 $T1
type T2 $T2

func main() {
	test(4, 2)
}
func test(a T1, b T2) {
	assert(a << b == 16)
	assert(a >> b == 1)
	assert(a << T2(2) == 16)
	assert(a >> T2(2) == 1)
	assert(T1(4) << b == 16)
	assert(T1(4) >> b == 1)
	var c T1 = 4
	var d T2 = 2
	assert(c << d == 16)
	assert(c >> d == 1)
}
func assert(t bool) {
	if !t {
		panic("error")
	}
}
`
	types := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}
	for _, t1 := range types {
		t.Log("test binop shift", t1)
		for _, t2 := range types {
			r := strings.NewReplacer("$T1", "="+t1, "$T2", "="+t2)
			src := r.Replace(tsrc)
			_, err := ixgo.RunFile("main.go", src, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
		}
		for _, t2 := range types {
			r := strings.NewReplacer("$T1", "="+t1, "$T2", t2)
			src := r.Replace(tsrc)
			_, err := ixgo.RunFile("main.go", src, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
		}
		for _, t2 := range types {
			r := strings.NewReplacer("$T1", t1, "$T2", "="+t2)
			src := r.Replace(tsrc)
			_, err := ixgo.RunFile("main.go", src, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
		}
		for _, t2 := range types {
			r := strings.NewReplacer("$T1", t1, "$T2", t2)
			src := r.Replace(tsrc)
			_, err := ixgo.RunFile("main.go", src, nil, 0)
			if err != nil {
				t.Fatal(err)
			}
		}
	}
}

func TestUnOpNot(t *testing.T) {
	src := `package main
type T bool
func main() {
	test(false,false)
	testConst()
}
func test(b1 bool, b2 T) {
	if v := !b1; v != true {
		panic("must true")
	}
	if v := !b2; v != true {
		panic("must true")
	}
}
func testConst() {
	var b1 bool
	var b2 T
	if v := !b1; v != true {
		panic("must true")
	}
	if v := !b2; v != true {
		panic("must true")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnOpSubInt(t *testing.T) {
	tsrc := `package main
type T $int

func main() {
	test(10,10)
	testConst()
}

func test(a $int, b T) {
	if r := -a; r != -10 {
		panic("must -10")
	}
	if r := -b; r != -10 {
		panic("must -10")
	}
}
func testConst() {
	var a $int = 10
	var b T = 10
	if r := -a; r != -10 {
		panic("must -10")
	}
	if r := -b; r != -10 {
		panic("must -10")
	}
}
`
	types := []string{
		"int", "int8", "int16", "int32", "int64",
	}

	for _, s := range types {
		t.Log("test unop sub", s)
		src := strings.Replace(tsrc, "$int", s, -1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpSubUint(t *testing.T) {
	tsrc := `package main

type T $uint

func main() {
	test(0b1010, 0b1011)
	testConst()
}

func test(a $uint, b T) {
	if r := -a; r&0xff != 0b11110110 {
		panic(r)
	}
	if r := -b; r&0xff != 0b11110101 {
		panic(r)
	}
}
func testConst() {
	var a $uint = 0b1010
	var b T = 0b1011
	if r := -a; r&0xff != 0b11110110 {
		panic(r)
	}
	if r := -b; r&0xff != 0b11110101 {
		panic(r)
	}
}
`
	types := []string{
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}
	for _, s := range types {
		src := strings.Replace(tsrc, "$uint", s, -1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpSubFloat(t *testing.T) {
	tsrc := `package main

type T $float

func main() {
	test(5.0, 6.5)
	testConst()
}

func test(a $float, b T) {
	if r := -a; r != -5.0 {
		panic(r)
	}
	if r := -b; r != -6.5 {
		panic(r)
	}
}
func testConst() {
	var a $float = 5.0
	var b T = 6.5
	if r := -a; r != -5.0 {
		panic(r)
	}
	if r := -b; r != -6.5 {
		panic(r)
	}
}
`
	types := []string{
		"float32", "float64",
	}
	for _, s := range types {
		t.Log("test unop sub", s)
		src := strings.Replace(tsrc, "$float", s, -1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpSubComplex(t *testing.T) {
	tsrc := `package main

type T $complex

func main() {
	test(1+2i, 3+4i)
	testConst()
}

func test(a $complex, b T) {
	if r := -a; r != -1-2i {
		panic(r)
	}
	if r := -b; r != -3-4i {
		panic(r)
	}
}
func testConst() {
	var a $complex = 1+2i
	var b T = 3+4i
	if r := -a; r != -1-2i {
		panic(r)
	}
	if r := -b; r != -3-4i {
		panic(r)
	}
}
`
	types := []string{
		"complex64", "complex128",
	}
	for _, s := range types {
		t.Log("test unop sub", s)
		src := strings.Replace(tsrc, "$complex", s, -1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpXorInt(t *testing.T) {
	tsrc := `package main

type T $int

func main() {
	test(10, 11)
	testConst()
}

func test(a $int, b T) {
	if r := ^a; r != -11 {
		panic(r)
	}
	if r := ^b; r != -12 {
		panic(r)
	}
}
func testConst() {
	var a $int = 10
	var b T = 11
	if r := ^a; r != -11 {
		panic(r)
	}
	if r := ^b; r != -12 {
		panic(r)
	}
}
`
	types := []string{
		"int", "int8", "int16", "int32", "int64",
	}

	for _, s := range types {
		t.Log("test unop xor", s)
		src := strings.Replace(tsrc, "$int", s, -1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestUnOpXorUint(t *testing.T) {
	tsrc := `package main

type T $uint

func main() {
	test(0b1010, 0b1011)
	testConst()
}

func test(a $uint, b T) {
	if r := ^a; r&0xff != 0b11110101 {
		panic(r)
	}
	if r := ^b; r&0xff != 0b11110100 {
		panic(r)
	}
}
func testConst() {
	var a $uint = 0b1010
	var b T = 0b1011
	if r := ^a; r&0xff != 0b11110101 {
		panic(r)
	}
	if r := ^b; r&0xff != 0b11110100 {
		panic(r)
	}
}
`
	types := []string{
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}

	for _, s := range types {
		t.Log("test unop xor", s)
		src := strings.Replace(tsrc, "$uint", s, -1)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestBinOpEQ(t *testing.T) {
	src := `package main

import "unsafe"

type T struct {
	a int
	b int
	_ int
	_ int
}

func main() {
	// array
	ar1 := [2]int{10,20}
	ar2 := [2]int{10,20}
	if ar1 != ar2 {
		panic("error array")
	}
	ar1[0] = 1
	if ar1 == ar2 {
		panic("error array")
	}
	// struct & interface{}
	t1 := T{1,2,3,4}
	t2 := T{1,2,7,8}
	if t1 != t2 {
		panic("error struct")
	}
	if (interface{})(t1) != (interface{})(t2) {
		panic("error interface")
	}
	t1.a = 10
	if t1 == t2 {
		panic("error struct")
	}
	if (interface{})(t1) == (interface{})(t2) {
		panic("error interface")
	}
	// ptr
	ptr1 := &t1
	ptr2 := &t2
	if ptr1 == ptr2 {
		panic("error ptr")
	}
	ptr2 = &t1
	if ptr1 != ptr2 {
		panic("error ptr")
	}
	// unsafe pointer
	p1 := unsafe.Pointer(&t1)
	p2 := unsafe.Pointer(&t2)
	if p1 == p2 {
		panic("error unsafe pointer")
	}
	p2 = unsafe.Pointer(ptr2)
	if p1 != p2 {
		panic("error unsafe pointer")
	}
	// chan
	ch1 := make(chan int)
	ch2 := make(chan int)
	if ch1 == ch2 {
		panic("error chan")
	}
	ch3 := ch1
	if ch1 != ch3 {
		panic("error chan")
	}
	ch4 := (<-chan int)(ch1)
	ch5 := (chan<- int)(ch1)
	if ch1 != ch4 {
		panic("error chan")
	}
	if ch5 != ch1 {
		panic("error chan")
	}
	ch6 := (<-chan int)(ch2)
	ch7 := (chan<- int)(ch2)
	if ch6 == ch1 {
		panic("error chan")
	}
	if ch1 == ch7 {
		panic("error chan")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpTypeChangeBasic(t *testing.T) {
	src := `package main

import (
	"bytes"
	"fmt"
)

func main() {
	testNil()
	testConst()
	testFunc()
	testPtr()
	testSlice()
	testMap()
	testStruct()
	testArray()
	testChan()
	testInterface()
}

func testNil() {
	var p *int
	type Ptr *int
	v := Ptr(p)
	if s := fmt.Sprintf("%T %v", v, v); s != "main.Ptr <nil>" {
		panic(s)
	}
}

func testConst() {
	var n int = 10
	type T int
	v := T(n)
	if s := fmt.Sprintf("%T %v", v, v); s != "main.T 10" {
		panic(s)
	}
}

func myfunc(int) {}

type Func func(int)

func testFunc() {
	v := Func(myfunc)
	if s := fmt.Sprintf("%T %v", v, v == nil); s != "main.Func false" {
		panic(s)
	}
}

func testPtr() {
	i := 10
	p := &i
	type T *int
	v := T(p)
	i = 20
	if s := fmt.Sprintf("%T %v", v, *v); s != "main.T 20" {
		panic(s)
	}
	p2 := (*int)(v)
	i = 30
	if s := fmt.Sprintf("%T %v", p2, *p2); s != "*int 30" {
		panic(s)
	}
}

func testSlice() {
	i := []int{100, 200, 300}
	type T []int
	v := T(i)
	if s := fmt.Sprintf("%T %v", v, v); s != "main.T [100 200 300]" {
		panic(s)
	}
	p := []int(v)
	v[0] = 1
	if s := fmt.Sprintf("%T %v", p, p); s != "[]int [1 200 300]" {
		panic(s)
	}
}

func testMap() {
	i := map[int]string{1: "hello", 2: "world"}
	type T map[int]string
	v := T(i)
	if s := fmt.Sprintf("%T %v", v, v); s != "main.T map[1:hello 2:world]" {
		panic(s)
	}
	type T2 map[int]string
	v2 := T(v)
	i[1] = "hello2"
	if s := fmt.Sprintf("%T %v", v2, v2); s != "main.T map[1:hello2 2:world]" {
		panic(s)
	}
}

func testStruct() {
	type Pt struct {
		x int
		y int
	}
	type T Pt
	i := Pt{10, 20}
	v := T(i)
	i.x = 1
	if s := fmt.Sprintf("%T %v", v, v); s != "main.T {10 20}" {
		panic(s)
	}
}

func testArray() {
	ar := [3]int{100, 200, 300}
	type T [3]int
	v := T(ar)
	ar[0] = 1
	if s := fmt.Sprintf("%T %v", v, v); s != "main.T [100 200 300]" {
		panic(s)
	}
}

func testChan() {
	ch := make(chan int)
	v := (chan<- int)(ch)
	if s := fmt.Sprintf("%T", v); s != "chan<- int" {
		panic(s)
	}
	type T1 chan int
	v1 := T1(ch)
	if s := fmt.Sprintf("%T", v1); s != "main.T1" {
		panic(s)
	}
	type T2 <-chan int
	v2 := T2(ch)
	if s := fmt.Sprintf("%T", v2); s != "main.T2" {
		panic(s)
	}
	type T3 chan<- int
	v3 := T3(ch)
	if s := fmt.Sprintf("%T", v3); s != "main.T3" {
		panic(s)
	}
}

func testInterface() {
	var buf fmt.Stringer = bytes.NewBufferString("hello")
	type T fmt.Stringer
	v := T(buf)
	if s := fmt.Sprintf("%T %v", v, v.String()); s != "*bytes.Buffer hello" {
		panic(s)
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpTypeChange3(t *testing.T) {
	tsrc := `package main
import "fmt"
type T $int
type T2 $int
func main() {
	test($value)
}
func test(v $int) {
	r := T(v)
	if s := fmt.Sprintf("%T", r); s != "main.T" {
		panic(s)
	}
	r2 := T2(r)
	if s := fmt.Sprintf("%T", r2); s != "main.T2" {
		panic(s)
	}
	n := $int(r2)
	if s := fmt.Sprintf("%T", n); s != "$int" {
		panic(s)
	}
	if n != v {
		panic("error")
	}
}
`
	ints := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}
	for _, s := range ints {
		t.Log("test changetype xtype", s)
		r := strings.NewReplacer("$int", s, "$value", "10")
		src := r.Replace(tsrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	floats := []string{"float32", "float64"}
	for _, s := range floats {
		t.Log("test changetype xtype", s)
		r := strings.NewReplacer("$int", s, "$value", "100.5")
		src := r.Replace(tsrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	comps := []string{"complex64", "complex128"}
	for _, s := range comps {
		t.Log("test changetype xtype", s)
		r := strings.NewReplacer("$int", s, "$value", "1+2i")
		src := r.Replace(tsrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		t.Log("test changetype xtype", "bool")
		r := strings.NewReplacer("$int", "bool", "$value", "true")
		src := r.Replace(tsrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	{
		t.Log("test changetype xtype", "string")
		r := strings.NewReplacer("$int", "string", "$value", `"hello"`)
		src := r.Replace(tsrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestOpConvert(t *testing.T) {
	tsrc := `package main

type T = $int

const (
	V = $V
	N = $N
	F = $F
)

func main() {
	test1(V)
	test2(V)
}

func test1(n T) {
	if int(n) != N {
		panic("error")
	}
	if int8(n) != N {
		panic("error")
	}
	if int16(n) != N {
		panic("error")
	}
	if int32(n) != N {
		panic("error")
	}
	if int64(n) != N {
		panic("error")
	}
	if uint(n) != N {
		panic("error")
	}
	if uint8(n) != N {
		panic("error")
	}
	if uint16(n) != N {
		panic("error")
	}
	if uint32(n) != N {
		panic("error")
	}
	if uint64(n) != N {
		panic("error")
	}
	if uintptr(n) != N {
		panic("error")
	}
	if float32(n) != F {
		panic("error")
	}
	if float64(n) != F {
		panic("error")
	}
}

type Int int
type Int8 int8
type Int16 int16
type Int32 int32
type Int64 int64
type Uint uint
type Uint8 uint8
type Uint16 uint16
type Uint32 uint32
type Uint64 uint64
type Uintptr uintptr
type Float32 float32
type Float64 float64

func test2(n T) {
	if Int(n) != N {
		panic("error")
	}
	if Int8(n) != N {
		panic("error")
	}
	if Int16(n) != N {
		panic("error")
	}
	if Int32(n) != N {
		panic("error")
	}
	if Int64(n) != N {
		panic("error")
	}
	if Uint(n) != N {
		panic("error")
	}
	if Uint8(n) != N {
		panic("error")
	}
	if Uint16(n) != N {
		panic("error")
	}
	if Uint32(n) != N {
		panic("error")
	}
	if Uint64(n) != N {
		panic("error")
	}
	if Uintptr(n) != N {
		panic("error")
	}
	if Float32(n) != F {
		panic("error")
	}
	if Float64(n) != F {
		panic("error")
	}
}
`
	csrc := `package main

type T = $complex

const (
	N = $N
)

func main() {
	test1(N)
	test2(N)
}

func test1(n T) {
	if complex64(n) != N {
		panic("error")
	}
	if complex128(n) != N {
		panic("error")
	}
}

type Complex64 complex64
type Complex128 complex128


func test2(n T) {
	if Complex64(n) != N {
		panic("error")
	}
	if Complex128(n) != N {
		panic("error")
	}
}
`
	ints := []string{
		"int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
	}
	floats := []string{"float32", "float64"}
	comps := []string{"complex64", "complex128"}
	for _, s := range ints {
		t.Log("test convert xtype", s)
		r := strings.NewReplacer("$int", s, "$V", "100", "$N", "100", "$F", "100")
		src := r.Replace(tsrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range ints {
		t.Log("test convert typed", s)
		r := strings.NewReplacer("= $int", s, "$V", "100", "$N", "100", "$F", "100")
		src := r.Replace(tsrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range floats {
		t.Log("test convert xtype", s)
		r := strings.NewReplacer("$int", s, "$V", "100.5", "$N", "100", "$F", "100.5")
		src := r.Replace(tsrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range floats {
		t.Log("test convert typed", s)
		r := strings.NewReplacer("= $int", s, "$V", "100.5", "$N", "100", "$F", "100.5")
		src := r.Replace(tsrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range comps {
		t.Log("test convert xtype", s)
		r := strings.NewReplacer("$complex", s, "$N", "1+2i")
		src := r.Replace(csrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
	for _, s := range comps {
		t.Log("test convert typed", s)
		r := strings.NewReplacer("= $complex", s, "$N", "1+2i")
		src := r.Replace(csrc)
		_, err := ixgo.RunFile("main.go", src, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestOpIf(t *testing.T) {
	src := `package main

type Bool bool

func main() {
	test1(true)
	test1(false)
	test2(true)
	test2(false)
	testConst1()
	testConst2()
}

func test1(b bool) {
	if b {
		println(true)
	} else {
		println(false)
	}
}
func test2(b Bool) {
	if b {
		println(true)
	} else {
		println(false)
	}
}
func testConst1() {
	var b bool
	if !b {
		println(false)
	}
	b = true
	if b {
		println(true)
	}
}
func testConst2() {
	var b Bool
	if !b {
		println(false)
	}
	b = true
	if b {
		println(true)
	}
}
`
	out := `true
false
true
false
false
true
false
true
`
	ctx := ixgo.NewContext(0)
	var buf bytes.Buffer
	ctx.SetPrintOutput(&buf)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
	if buf.String() != out {
		t.Fatal("error")
	}
}

func TestGoexitDeadlock(t *testing.T) {
	src := `package main
import (
	"os"
	"runtime"
)

func init() {
	runtime.Goexit()
	os.Exit(-1)
}

func main() {
	os.Exit(-2)
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	if err.Error() != ixgo.ErrGoexitDeadlock.Error() {
		t.Fatal(err)
	}
}

func TestGlobalExtFunc(t *testing.T) {
	src := `package main
import "math"
var (
	max = math.Max
)
func main() {
	if max(1,3) != 3 {
		panic("error")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterBuiltin(t *testing.T) {
	src := `package main

func main() {
	dump_info(typeof("hello"))
	dump_ints(20,30,50)
}
`
	ixgo.RegisterCustomBuiltin("typeof", reflect.TypeOf)
	var info interface{}
	ixgo.RegisterCustomBuiltin("dump_info", func(v interface{}) {
		info = v
	})
	var all int
	ixgo.RegisterCustomBuiltin("dump_ints", func(a ...int) {
		for _, v := range a {
			all += v
		}
	})
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		panic(err)
	}
	if info != reflect.TypeOf("hello") {
		panic("error")
	}
	if all != 100 {
		panic(all)
	}
}

func TestVariadicFunc(t *testing.T) {
	src := `package main
type Context struct{}
type ClientConn struct{}
type CallOption struct{}

type UnaryInvoker func(ctx Context, method string, req, reply interface{}, cc *ClientConn, opts ...CallOption) error

func main() {
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUnsafePointer(t *testing.T) {
	src := `package main

import (
	"unsafe"
)

func main() {
	type T1 unsafe.Pointer
	type T2 unsafe.Pointer
	i := 10
	p := unsafe.Pointer(&i)
	p1 := T1(uintptr(unsafe.Pointer(&i)))
	p2 := T2(&i)
	if T1(p) != p1 {
		panic("error T1")
	}
	if T2(p1) != (p2) {
		panic("error T2")
	}
	if unsafe.Pointer(p1) != unsafe.Pointer(p2) {
		panic("error unsafe")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddImport(t *testing.T) {
	pkg := `package pkg
import "fmt"

func Add(i, j int) int {
	return i+j
}

func Println(a ...interface{}) {
	fmt.Println(a...)
}
`
	src := `package main
import "pkg"

func main() {
	if pkg.Add(100, 200) != 300 {
		panic("error pkg.Add")
	}
	pkg.Println("Hello")
}
`
	ctx := ixgo.NewContext(0)
	err := ctx.AddImportFile("pkg", "pkg.go", pkg)
	if err != nil {
		t.Fatal(err)
	}
	_, err = ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAddImportDir(t *testing.T) {
	src := `package main
import "ixgo/pkg"

func main() {
	if pkg.Add(100, 200) != 300 {
		panic("error pkg.Add")
	}
	pkg.Println("Hello")
}
`
	ctx := ixgo.NewContext(0)
	err := ctx.AddImport("ixgo/pkg", "./testdata/pkg")
	if err != nil {
		t.Fatal(err)
	}
	_, err = ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLoadImport(t *testing.T) {
	if runtime.GOOS == "darwin" {
		switch runtime.Version()[:6] {
		case "go1.14", "go1.15":
			t.Skip("skip for workflows")
		}
	}
	src := `package main
import "github.com/goplus/ixgo/testdata/pkg"

func main() {
	if pkg.Add(100, 200) != 300 {
		panic("error pkg.Add")
	}
	pkg.Println("Hello")
}
`
	ctx := ixgo.NewContext(0)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestAppend(t *testing.T) {
	src := `package main

type T string

func main() {
	a := []uint8("hello")
	b := "world"
	c := T("world")
	_ = append(a,b...)
	_ = append(a,c...)
}
	`
	ctx := ixgo.NewContext(0)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestStructEmbed(t *testing.T) {
	p := &ixgo.Package{
		Name: "info",
		Path: "github.com/goplus/ixgo/testdata/info",
		Interfaces: map[string]reflect.Type{
			"Info": reflect.TypeOf((*info.Info)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"MyInfo": reflect.TypeOf((*info.MyInfo)(nil)).Elem(),
		},
		Funcs: map[string]reflect.Value{
			"NewInfo": reflect.ValueOf(info.NewInfo),
		},
	}
	ixgo.RegisterPackage(p)

	src := `package main

import (
	"github.com/goplus/ixgo/testdata/info"
)

type MyInfo struct {
	*info.MyInfo
}

type MyInfo2 struct {
	info.Info
}

func main() {
	info := info.NewInfo(1)
	m := &MyInfo{info}
	check(m)
	m2 := &MyInfo2{info}
	check(m2)
}

func check(info info.Info) {
	info.SetMode(-1)
	if info.Mode() != -1 {
		panic("mode")
	}
	if info.Count(1, 2, 3) != 3 {
		panic("count")
	}
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestBuildTags(t *testing.T) {
	if runtime.GOOS == "darwin" {
		switch runtime.Version()[:6] {
		case "go1.14", "go1.15":
			t.Skip("skip for workflows")
		}
	}
	src1 := `package main

import (
	"github.com/goplus/ixgo/testdata/msg"
)

func main() {
	if msg.Info != "nomsg" {
		panic(msg.Info)
	}
}
`
	ctx1 := ixgo.NewContext(0)
	_, err := ctx1.RunFile("main.go", src1, nil)
	if err != nil {
		t.Fatal(err)
	}

	src2 := `package main

import (
	"github.com/goplus/ixgo/testdata/msg"
)

func main() {
	if msg.Info != "msg" {
		panic(msg.Info)
	}
}
`
	ctx2 := ixgo.NewContext(0)
	ctx2.BuildContext.BuildTags = []string{"msg"}
	_, err = ctx2.RunFile("main.go", src2, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEqualChan(t *testing.T) {
	src := `package main

func main() {
	ch := make(chan bool)
	ch1 := (<-chan bool)(ch)
	ch2 := (chan<- bool)(ch)
	var i interface{} = ch
	var i1 interface{} = ch1
	var i2 interface{} = ch2

	check(ch == ch1, true)
	check(ch == ch2, true)
	check(i == i1, false)
	check(i == i2, false)
	check(i1 == i2, false)

	check(ch != ch1, false)
	check(ch != ch2, false)
	check(i != i1, true)
	check(i != i2, true)
	check(i1 != i2, true)
}

func check(b1 bool, b2 bool) {
	if b1 != b2 {
		panic("error")
	}
}
`
	ctx := ixgo.NewContext(0)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIssue23545(t *testing.T) {
	src := `package main

func main() {
	if a := Get(); a != dummyID(1234) {
		panic("FAIL")
	}
	var i1 interface{} = Get()
	var i2 interface{} = dummyID(1234)
	if i1 == i2 {
		panic("FAIL")
	}
}

func dummyID(x int) [Size]interface{} {
	var out [Size]interface{}
	out[0] = x
	return out
}

const Size = 32

type OutputID [Size]interface{}

//go:noinline
func Get() OutputID {
	return dummyID(1234)
}
`
	ctx := ixgo.NewContext(0)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestMakeChan(t *testing.T) {
	src := `package main

type Send chan<- int
type Recv <-chan int

func main() {
	_ = make(chan int)
	_ = make(<-chan int)
	_ = make(chan<- int)
	_ = make(Send)
	_ = make(Recv)
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestNoStrictMode(t *testing.T) {
	src := `package main
import "fmt"

func f() int

func main() {
	var a int
	println(100)
}
`
	ctx := ixgo.NewContext(ixgo.EnableNoStrict)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestUserFuncConvert(t *testing.T) {
	src := `package main

import "fmt"

type myint int
type point struct {
	x int
	y int
}

func mytest1(n myint, pt point) myint
func mytest2(n myint, pt *point) myint

func main() {
	n := mytest1(100, point{100,200})
	if n != 400 {
		panic(fmt.Errorf("error mytest1, must 400, have %v",n))
	}
	n = mytest2(100, &point{100,200})
	if n != 30000 {
		panic(fmt.Errorf("error mytest2, must 30000, have %v",n))
	}
}
`
	ctx := ixgo.NewContext(0)
	ixgo.RegisterExternal("main.mytest1", func(n int, pt struct {
		x int
		y int
	}) int {
		return n + pt.x + pt.y
	})
	ctx.RegisterExternal("main.mytest2", func(n int, pt *struct {
		x int
		y int
	}) int {
		return n * (pt.x + pt.y)
	})
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRunContext(t *testing.T) {
	var src string = `
package main

import "time"

func main() {
	println("hello world")
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			timeout := time.After(time.Second * 1)
			select {
			case <-timeout:
				println("work")
			}
		}
		ch <- 1
	}()
	<-ch
	println("exit")
}
`
	ctx := ixgo.NewContext(0)
	ctx.RunContext, _ = context.WithTimeout(context.Background(), 1e9)
	code, err := ctx.RunFile("main.go", src, nil)
	if code != 2 {
		t.Fatalf("exit code must 2")
	}
	t.Log("cancel context:", err)
}

func TestReflectArray(t *testing.T) {
	var src = `package main

import (
	"reflect"
	"fmt"
)

type IntArray [2]int

func (i IntArray) String() string {
	return fmt.Sprintf("(%v,%v)", i[0], i[1])
}

func (i *IntArray) Set(x, y int) {
	*(*int)(&i[0]), *(*int)(&i[1]) = x, y
}

func (i IntArray) Get() (int, int) {
	return i[0], i[1]
}

func (i IntArray) Scale(v int) IntArray {
	return IntArray{i[0] * v, i[1] * v}
}

func main() {
	var a IntArray
	a.Set(100,200)
	b := a.Scale(5)
	if b[0] != 500 || b[1] != 1000 {
		panic("error")
	}
	typ := reflect.TypeOf((*IntArray)(nil)).Elem()
	v := reflect.New(typ).Elem()
	v.Addr().MethodByName("Set").Call([]reflect.Value{reflect.ValueOf(100),reflect.ValueOf(200)})
	x := v.MethodByName("Scale").Call([]reflect.Value{reflect.ValueOf(5)})
	if x[0].Index(0).Int() != 500 || x[0].Index(1).Int() != 1000 {
		panic("error reflect")
	}
}
`
	ctx := ixgo.NewContext(0)
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

// A FileInfo describes a file and is returned by Stat.
type fileInfo struct {
	name    string
	size    int64
	mode    fs.FileMode
	modTime time.Time
	isDir   bool
}

func (f fileInfo) Name() string {
	return f.name
}
func (f fileInfo) Size() int64 {
	return f.size
}
func (f fileInfo) Mode() fs.FileMode {
	return f.mode
}
func (f fileInfo) ModTime() time.Time {
	return f.modTime
}
func (f fileInfo) IsDir() bool {
	return f.isDir
}
func (f fileInfo) Sys() interface{} {
	return nil
}

func TestBuildContext(t *testing.T) {
	ctx := ixgo.NewContext(0)
	const pkgDir = "mypkg"
	const mypkg = `package mypkg
import "fmt"

func Demo(n int) {
	fmt.Println(n)
}
`
	ctx.Lookup = func(root, path string) (dir string, found bool) {
		if path == "mypkg" {
			return pkgDir, true
		}
		return
	}
	ctx.BuildContext.IsDir = func(path string) bool {
		if path == pkgDir {
			return true
		}
		return false
	}
	ctx.BuildContext.ReadDir = func(dir string) ([]fs.FileInfo, error) {
		if dir == pkgDir {
			return []fs.FileInfo{&fileInfo{name: "mypkg.go", size: int64(len(mypkg)), mode: 0644}}, nil
		}
		return nil, nil
	}
	ctx.BuildContext.OpenFile = func(path string) (io.ReadCloser, error) {
		var buf bytes.Buffer
		buf.WriteString(mypkg)
		return io.NopCloser(&buf), nil
	}
	_, err := ctx.RunFile("main.go", `package main
import "mypkg"

func main() {
	mypkg.Demo(100)
}
`, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRuntimeRefelct(t *testing.T) {
	ctx := ixgo.NewContext(ixgo.OptionLoadRutimeImethod)
	ctx.RegisterExternal("main.run", func(v interface{}) {
		if e, ok := v.(interface{ MainEntry() }); ok {
			e.MainEntry()
		}
	})
	_, err := ctx.RunFile("main.go", `package main

type Game struct {
	run bool
}

func (p *Game) MainEntry() {
	println("main entry")
	p.run = true
}

type T = Game

func run(interface{})

func main() {
	g := &T{}
	run(g)
	if !g.run {
		panic("error")
	}
}
`, nil)
	if err != nil {
		panic(err)
	}
}

func TestEmbedImethod(t *testing.T) {
	p := &ixgo.Package{
		Name: "pkg",
		Path: "github.com/goplus/ixgo/testdata/pkg",
		NamedTypes: map[string]reflect.Type{
			"App": reflect.TypeOf((*pkg.App)(nil)).Elem(),
		},
		Funcs: map[string]reflect.Value{
			"RunApp":  reflect.ValueOf(pkg.RunApp),
			"RunTest": reflect.ValueOf(pkg.RunTest),
		},
	}
	ixgo.RegisterPackage(p)
	src := `
package main

import "github.com/goplus/ixgo/testdata/pkg"

type MyApp struct {
	pkg.App
}

func (p *MyApp) MainEntry() {
	if !p.IsInit() {
		panic("init app error")
	}
	println("MainEntry: app")
}

type TestApp struct {
	pkg.App
}

func (p *TestApp) MainEntry() {
	if !p.IsTest() {
		panic("init test error")
	}
	println("MainEntry: test")
}

func main() {
	pkg.RunApp(&MyApp{})
	pkg.RunTest(&TestApp{})
}
`
	for _, mode := range []ixgo.Mode{ixgo.OptionLoadDefaultImethod, ixgo.OptionLoadRutimeImethod, ixgo.OptionLoadAllImethod} {
		_, err := ixgo.RunFile("main.go", src, nil, mode)
		if err != nil {
			t.Errorf("error: %v mode: %d", err, mode)
		}
	}
}

func TestExternFunc(t *testing.T) {
	type Stub struct {
		Add func(int, int) int
		Sub func(int, int) int
	}
	src := `
package main

type Stub struct {
	Add func(int, int) int
	Sub func(int, int) int
}

func GetStub() *Stub

func main() {
	stub := GetStub()
	v1 := stub.Add(100, 200)
	if v1 != 300 {
		panic("error add")
	}
	v2 := stub.Sub(100, 200)
	if v2 != -100 {
		panic("error sub")
	}
}
`
	ctx := ixgo.NewContext(0)
	ctx.RegisterExternal("main.GetStub", func() *Stub {
		return &Stub{
			Add: func(a, b int) int {
				return a + b
			},
			Sub: reflect.MakeFunc(reflect.TypeOf(func(int, int) int { return 0 }), func(args []reflect.Value) []reflect.Value {
				return []reflect.Value{reflect.ValueOf(int(args[0].Int()) - int(args[1].Int()))}
			}).Interface().(func(int, int) int),
		}
	})
	_, err := ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRegisterPatch(t *testing.T) {
	ctx := ixgo.NewContext(0)
	err := ctx.RegisterPatch("fmt", `
package fmt

import (
	"fmt"
	"math"
)

const Pi = math.Pi

var (
	Patch = "v1.0"
)

func Println1(v ...interface{}) {
	fmt.Println(v...)
}
`)
	if err != nil {
		t.Fatal(err)
	}
	err = ctx.RegisterPatch("fmt", `
package fmt
import "fmt"

func Println2(v ...interface{}) {
	fmt.Println(v...)
}
`)
	if err != nil {
		t.Fatal(err)
	}
	src := `
package main

import "fmt"

func main() {
	fmt.Println1(fmt.Pi,fmt.Patch)
	fmt.Println2("hello world")
}
`
	_, err = ctx.RunFile("main.go", src, nil)
	if err != nil {
		t.Fatal(err)
	}
}

func TestIcallCached(t *testing.T) {
	src := `
package main

import (
	"bytes"
	"fmt"
)

type T struct {
	bytes.Buffer
}

func (t *T) Text() string {
	return t.String()
}

func main() {
	var t T
	t.WriteString("hello world")
	fmt.Println(&t)
}
`
	ctx := ixgo.NewContext(ixgo.SupportMultipleInterp)

	pkg1, err := ctx.LoadFile("main.go", src)
	if err != nil {
		t.Fatal(err)
	}
	pkg2, err := ctx.LoadFile("main.go", src)
	if err != nil {
		t.Fatal(err)
	}

	_, orgAllocate, _ := ixgo.IcallStat()
	orgCached := ixgo.IcallCached()
	defer func() {
		_, allocate, _ := ixgo.IcallStat()
		cached := ixgo.IcallCached()
		if orgAllocate != allocate {
			t.Fatalf("allocate release error: %v", allocate-orgAllocate)
		}
		if orgCached != cached {
			t.Fatalf("cached release error: %v", cached-orgCached)
		}
	}()

	interp1, err := ctx.NewInterp(pkg1)
	if err != nil {
		t.Fatal(err)
	}
	interp2, err := ctx.NewInterp(pkg2)
	if err != nil {
		t.Fatal(err)
	}

	checkStats := func(step string) {
		_, allocate, _ := ixgo.IcallStat()
		cached := ixgo.IcallCached()
		alloc1 := interp1.IcallAlloc()
		alloc2 := interp2.IcallAlloc()
		if cached == 0 || allocate != cached+alloc1+alloc2 {
			t.Fatalf("%s: cached error: all %v, cached %v, alloc1 %v, alloc2 %v",
				step, allocate, cached, alloc1, alloc2)
		}
	}

	checkStats("after creating interpreters")

	_, err = ctx.RunInterp(interp1, "main.go", nil)
	if err != nil {
		t.Fatal(err)
	}
	interp1.UnsafeRelease()

	checkStats("after releasing interp1")

	_, err = ctx.RunInterp(interp2, "main.go", nil)
	if err != nil {
		t.Fatal(err)
	}
	interp2.UnsafeRelease()
}
