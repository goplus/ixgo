package mixedrunner

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/goplus/ixgo"
	"golang.org/x/tools/go/ssa"
)

func TestMixedSelection(t *testing.T) {
	const source = `package main
func sum(n int,step float64) float64 {s:=0.0;for i:=0;i<n;i++ {s+=step+float64(i)};return s}
func call(n int,step float64) float64 {return sum(n,step)}
func unsupported(n int,step float64) float64 {a:=[2]float64{step,2};return a[n]}
`
	ctx := ixgo.NewContext(0)
	pkg, err := ctx.LoadFile("mixed.go", source)
	if err != nil {
		t.Fatal(err)
	}
	r, ok := MakeRunner(pkg.Func("sum")).(ixgo.TypedArgumentsRunner)
	if !ok || r.RunArgs([]ixgo.Value{5, 0.5}) != 12.5 {
		t.Fatal("mixed fallback not selected")
	}
	r, ok = MakeRunner(pkg.Func("call")).(ixgo.TypedArgumentsRunner)
	if !ok || r.RunArgs([]ixgo.Value{5, 0.5}) != 12.5 {
		t.Fatal("mixed call graph not selected")
	}
	if MakeRunner(pkg.Func("unsupported")) != nil {
		t.Fatal("unsupported memory must fall back")
	}
	interp, err := ixgo.NewContext(0).LoadInterp("mixed.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, tc := range []struct {
		name string
		n    int
		want float64
	}{{"sum", 5, 12.5}, {"call", 5, 12.5}, {"unsupported", 0, 0.5}} {
		got, err := interp.RunFunc(tc.name, tc.n, 0.5)
		if err != nil || got != tc.want {
			t.Fatalf("%s: %v %v", tc.name, got, err)
		}
	}
}

func TestDivmodMixedSelection(t *testing.T) {
	path := filepath.Join(runtime.GOROOT(), "test", "divmod.go")
	source, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		t.Skip("Go distribution test sources unavailable")
	}
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := ixgo.NewContext(0).LoadFile(path, source)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"udiv", "idiv"} {
		fn := pkg.Func(name)
		if fn == nil {
			t.Fatalf("%s: expected mixed graph", name)
		}
		r := MakeRunner(fn)
		if r == nil || reflect.TypeOf(r) != reflect.TypeOf(MakeRunner(fn)) {
			t.Fatalf("%s: mixed runner not selected", name)
		}
	}
	r := MakeRunner(pkg.Func("idiv")).(ixgo.TypedArgumentsRunner)
	if got := r.RunArgs([]ixgo.Value{int64(-17), int64(3)}); !reflect.DeepEqual(got, ixgo.Tuple{int64(-5), int64(-2)}) {
		t.Fatalf("idiv: %v", got)
	}
}

func TestStandaloneInstallation(t *testing.T) {
	selected := make(map[string]bool)
	ixgo.RegisterMixedRunner(func(fn *ssa.Function) ixgo.TypedRunner {
		selected[fn.Name()] = true
		return MakeRunner(fn)
	})
	t.Cleanup(func() { ixgo.RegisterMixedRunner(MakeRunner) })
	i, err := ixgo.NewContext(0).LoadInterp("standalone.go", `package main
func plain(n int) int {return n+1}
func integer(n int,u uint64) int {return n+int(u)}
func floating(x float32,y float64) float64 {return float64(x)+y}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	for _, tc := range []struct {
		name string
		args []ixgo.Value
		want ixgo.Value
	}{
		{"plain", []ixgo.Value{3}, 4},
		{"integer", []ixgo.Value{3, uint64(4)}, 7},
		{"floating", []ixgo.Value{float32(1.5), 2.5}, 4.0},
	} {
		if !selected[tc.name] {
			t.Fatalf("%s did not reach standalone mixed factory", tc.name)
		}
		if got, err := i.RunFunc(tc.name, tc.args...); err != nil || got != tc.want {
			t.Fatalf("%s: %v %v", tc.name, got, err)
		}
	}
}

func TestMixedRuntimeParity(t *testing.T) {
	const source = `package main
func wide(x uint64) uint64 { return x ^ 0xffffffffffffffff }
func pair(n int, x float64) (bool, float64) { return n>0, float64(n)+x }
func empty() {}
func recurse(n int) int { if n==0{return 0};return recurse(n-1)+1 }
func even(n int) bool {if n==0{return true};return odd(n-1)}
func odd(n int) bool {if n==0{return false};return even(n-1)}
func div(x,y int) int {return x/y}
func rem(x,y int) int {return x%y}
func shift(x uint64,y int8) uint64 {return x<<y}
func discard(x int) {div(1,x)}
func recoverCall(x int) (ok bool) {defer func(){ok=recover()!=nil}();discard(x);return}
func closure(x int) int {return func(y int) int {return x+y}(2)}
`
	pkg, err := ixgo.NewContext(0).LoadFile("parity.go", source)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name string
		args []ixgo.Value
		want ixgo.Value
	}{
		{"wide", []ixgo.Value{uint64(1)}, ^uint64(1)},
		{"pair", []ixgo.Value{2, 0.5}, ixgo.Tuple{true, 2.5}},
		{"empty", nil, nil},
		{"recurse", []ixgo.Value{10000}, 10000},
		{"even", []ixgo.Value{10001}, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			r, ok := MakeRunner(pkg.Func(tc.name)).(ixgo.TypedArgumentsRunner)
			if !ok {
				t.Fatal("runner rejected")
			}
			if got := r.RunArgs(tc.args); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v want %v", got, tc.want)
			}
		})
	}
	capture := func(f func()) (v interface{}) { defer func() { v = recover() }(); f(); return }
	zero, negative := 0, int8(-1)
	for _, tc := range []struct {
		name   string
		args   []ixgo.Value
		native func()
	}{
		{"div", []ixgo.Value{1, 0}, func() { _ = 1 / zero }},
		{"rem", []ixgo.Value{1, 0}, func() { _ = 1 % zero }},
		{"shift", []ixgo.Value{uint64(1), negative}, func() { _ = uint64(1) << negative }},
	} {
		r := MakeRunner(pkg.Func(tc.name)).(ixgo.TypedArgumentsRunner)
		got := capture(func() { r.RunArgs(tc.args) })
		want := capture(tc.native)
		if got == nil || !reflect.DeepEqual(got, want) {
			t.Fatalf("%s panic: %T %v; want %T %v", tc.name, got, got, want, want)
		}
		if _, ok := got.(runtime.Error); !ok {
			t.Fatalf("%s: not runtime.Error", tc.name)
		}
	}
	r := MakeRunner(pkg.Func("wide")).(ixgo.TypedArgumentsRunner)
	type named uint64
	for _, args := range [][]ixgo.Value{nil, {uint64(0), uint64(0)}, {int64(0)}, {named(0)}} {
		if capture(func() { r.RunArgs(args) }) == nil {
			t.Fatalf("accepted invalid args %v", args)
		}
	}
	i, err := ixgo.NewContext(0).LoadInterp("parity.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	fn, ok := i.GetFunc("pair")
	if !ok {
		t.Fatal("missing pair")
	}
	b, x := fn.(func(int, float64) (bool, float64))(2, 0.5)
	if !b || x != 2.5 {
		t.Fatalf("reflect tuple: %v %v", b, x)
	}
	fn, ok = i.GetFunc("empty")
	if !ok {
		t.Fatal("missing empty")
	}
	fn.(func())()
	for _, tc := range []struct {
		name string
		want ixgo.Value
	}{{"recoverCall", true}, {"closure", 2}} {
		if got, err := i.RunFunc(tc.name, 0); err != nil || got != tc.want {
			t.Fatalf("%s: %v %v", tc.name, got, err)
		}
	}
}

func TestMixedNumericDifferential(t *testing.T) {
	const source = `package main
func i8(x, y int8) (int8, bool) { return x+y, x < y }
func i16(x, y int16) int16 { return x &^ y }
func i32(x, y int32) int32 { return x ^ y }
func i64(x, y int64) int64 { return x / y }
func u(x, y uint) uint { return x >> y }
func u8(x, y uint8) uint8 { return x % y }
func u16(x, y uint16) uint16 { return x | y }
func u32(x, y uint32) uint32 { return x << y }
func u64(x, y uint64) uint64 { return x - y }
func up(x, y uintptr) uintptr { return x * y }
func f32(x, y float32) float32 { return x / y }
func f64(x, y float64) float64 { return x*y + y }
func conversions(x int64) (uint8, float32, float64) { return uint8(x), float32(x), float64(x) }
func constants() (uint64, int64, float64, bool) { return 0xffffffffffffffff, -1, 1.25, true }
func cInt(x int) int { return x + 1 }
func cInt8() int8 { return 7 }
func cInt16() int16 { return 7 }
func cInt32() int32 { return 7 }
func cUint() uint { return 7 }
func cUint8() uint8 { return 7 }
func cUint16() uint16 { return 7 }
func cUint32() uint32 { return 7 }
func cFloat32() float32 { return 1.25 }
func shifts(x int64, y uint8) (int64, uint64) { return x >> y, uint64(x) << y }
func bools(x, y bool) (bool, bool, bool) { return !x, x == y, x != y }
func neg(x int32) (int32, int32) { return -x, ^x }
func multi(x int, y float64) (int, float64) { return x+1, y/2 }
func unsupported(x int) int { return []int{x}[0] }
`
	ctx := ixgo.NewContext(0)
	pkg, err := ctx.LoadFile("numeric.go", source)
	if err != nil {
		t.Fatal(err)
	}
	interp, err := ctx.LoadInterp("numeric.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	tests := []struct {
		name string
		args []ixgo.Value
	}{
		{"i8", []ixgo.Value{int8(5), int8(7)}}, {"i16", []ixgo.Value{int16(0x55), int16(0x0f)}},
		{"i32", []ixgo.Value{int32(0x55), int32(0x0f)}}, {"i64", []ixgo.Value{int64(-17), int64(3)}},
		{"u", []ixgo.Value{uint(0x80), uint(3)}}, {"u8", []ixgo.Value{uint8(17), uint8(5)}},
		{"u16", []ixgo.Value{uint16(0x55), uint16(0x0f)}}, {"u32", []ixgo.Value{uint32(3), uint32(4)}},
		{"u64", []ixgo.Value{uint64(9), uint64(4)}}, {"up", []ixgo.Value{uintptr(3), uintptr(4)}},
		{"f32", []ixgo.Value{float32(7.5), float32(2)}}, {"f64", []ixgo.Value{2.5, 4.0}},
		{"conversions", []ixgo.Value{int64(-17)}}, {"constants", nil}, {"cInt", []ixgo.Value{6}}, {"cInt8", nil}, {"cInt16", nil}, {"cInt32", nil}, {"cUint", nil}, {"cUint8", nil}, {"cUint16", nil}, {"cUint32", nil}, {"cFloat32", nil},
		{"shifts", []ixgo.Value{int64(-16), uint8(2)}}, {"bools", []ixgo.Value{true, false}},
		{"neg", []ixgo.Value{int32(0x55)}}, {"multi", []ixgo.Value{3, 2.5}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			fn := pkg.Func(tc.name)
			r, ok := MakeRunner(fn).(ixgo.TypedArgumentsRunner)
			if !ok {
				t.Fatalf("no mixed runner for %s", tc.name)
			}
			want, err := interp.RunFunc(tc.name, tc.args...)
			if err != nil {
				t.Fatal(err)
			}
			if got := r.RunArgs(tc.args); !reflect.DeepEqual(got, want) {
				t.Fatalf("got %#v want %#v", got, want)
			}
		})
	}
	if MakeRunner(pkg.Func("unsupported")) != nil {
		t.Fatal("unsupported function selected")
	}
	if got := MakeRunner(pkg.Func("cInt")).Run(6); got != 7 {
		t.Fatalf("Run: %v", got)
	}
}

func TestMixedFallbackAndInvalidInputs(t *testing.T) {
	const source = `package main
func named(x MyInt) MyInt { return x }
type MyInt int
func badcall(x int) int { return len([]int{x}) }
func badresult(x int) []int { return []int{x} }
func zero() {}
func div(x, y int) int { return x/y }
`
	ctx := ixgo.NewContext(0)
	pkg, err := ctx.LoadFile("fallback.go", source)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"named", "badcall", "badresult"} {
		if MakeRunner(pkg.Func(name)) != nil {
			t.Fatalf("%s unexpectedly selected", name)
		}
	}
	zero := MakeRunner(pkg.Func("zero")).(ixgo.TypedArgumentsRunner)
	if got := zero.RunArgs(nil); got != nil {
		t.Fatalf("zero: %#v", got)
	}
	r := MakeRunner(pkg.Func("div")).(ixgo.TypedArgumentsRunner)
	for _, args := range [][]ixgo.Value{{}, {int8(1), int(1)}, {1}} {
		func() {
			defer func() {
				if recover() == nil {
					t.Errorf("invalid args %v did not panic", args)
				}
			}()
			r.RunArgs(args)
		}()
	}
}
