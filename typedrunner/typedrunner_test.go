package typedrunner

import (
	"math"
	"reflect"
	"testing"

	"github.com/goplus/ixgo"
	"golang.org/x/tools/go/ssa"
)

func TestLoopSelection(t *testing.T) {
	ctx := ixgo.NewContext(0)
	pkg, err := ctx.LoadFile("loops.go", `package main
func ints(n int) int { s:=0; for i:=0;i<n;i++ { s+=i };return s }
func uints(n uint) uint { s:=uint(0);for i:=uint(0);i<n;i++ { s+=i };return s }
func bytes(n uint8) uint8 { s:=uint8(0);for i:=uint8(0);i<n;i++ { s+=i };return s }
func longs(n int64) int64 { s:=int64(0);for i:=int64(0);i<n;i++ { s+=i };return s }
func plain(n int) int { return n+1 }
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name        string
		input, want any
	}{
		{"ints", 20, 190}, {"uints", uint(20), uint(190)},
		{"bytes", uint8(30), uint8(179)}, {"longs", int64(20), int64(190)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner := makeTypedRunner(pkg.Func(tc.name))
			if runner == nil {
				t.Fatal("nonrecursive loop was not selected")
			}
			if got := runner.Run(tc.input); got != tc.want {
				t.Fatalf("got %v, want %v", got, tc.want)
			}
		})
	}
	if runner := makeTypedRunner(pkg.Func("plain")); runner == nil || runner.Run(4) != 5 {
		t.Fatal("plain scalar function was not selected")
	}
}

func TestArgumentsSelection(t *testing.T) {
	ctx := ixgo.NewContext(0)
	pkg, err := ctx.LoadFile("arguments.go", `package main
func add(a,b int) int {return a+b}
func memory(n uint,p *uint) uint {a:=[2]uint{n,3};*p+=a[0];return *p+a[1]}
func recur(n int,p *int) {if n>0 {other(n-1,p);*p+=n}}
func other(n int,p *int) {recur(n,p)}
func entry(n int) int {p:=0;recur(n,&p);return p}
`)
	if err != nil {
		t.Fatal(err)
	}
	runner := makeTypedRunner(pkg.Func("add"))
	argsRunner, ok := runner.(ixgo.TypedArgumentsRunner)
	if !ok {
		t.Fatal("multiargument runner was not selected")
	}
	if got := argsRunner.RunArgs([]ixgo.Value{5, 9}); got != 14 {
		t.Fatalf("add: %v", got)
	}
	p := uint(7)
	runner = makeTypedRunner(pkg.Func("memory"))
	argsRunner, ok = runner.(ixgo.TypedArgumentsRunner)
	if !ok {
		t.Fatal("uint pointer runner was not selected")
	}
	if got := argsRunner.RunArgs([]ixgo.Value{uint(4), &p}); got != uint(14) || p != 11 {
		t.Fatalf("memory: %v %v", got, p)
	}
	interp, err := ctx.LoadInterp("entry.go", `package main
func recur(n int,p *int) {if n>0 {other(n-1,p);*p+=n}}
func other(n int,p *int) {recur(n,p)}
func entry(n int) int {p:=0;recur(n,&p);return p}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	got, err := interp.RunFunc("entry", 100)
	if err != nil || got != 5050 {
		t.Fatalf("entry: %v %v", got, err)
	}
}

func TestBroaderSelection(t *testing.T) {
	pkg, err := ixgo.NewContext(0).LoadFile("broader.go", `package main
func unary(n int) int {return -n}
func flag(b bool,n uint8) uint8 {if b {return n+1};return n}
func constant() uint64 {return 42}
func empty() {}
func mixed(a int,b uint64) int {return a+int(b)}
func pointer(n int) *int {return &n}
func deferred(n int) int {defer empty();return n}
func strings(n string) string {return n}
`)
	if err != nil {
		t.Fatal(err)
	}
	if r := makeTypedRunner(pkg.Func("unary")); r == nil || r.Run(3) != -3 {
		t.Fatal("unary selection")
	}
	for _, tc := range []struct {
		name string
		args []ixgo.Value
		want ixgo.Value
	}{
		{"flag", []ixgo.Value{true, uint8(4)}, uint8(5)},
		{"constant", nil, uint64(42)}, {"empty", nil, nil},
	} {
		r, ok := makeTypedRunner(pkg.Func(tc.name)).(ixgo.TypedArgumentsRunner)
		if !ok {
			t.Fatalf("%s not selected", tc.name)
		}
		if got := r.RunArgs(tc.args); got != tc.want {
			t.Fatalf("%s got=%v want=%v", tc.name, got, tc.want)
		}
	}
	for _, name := range []string{"mixed", "pointer", "deferred", "strings"} {
		if makeTypedRunner(pkg.Func(name)) != nil {
			t.Fatalf("%s must fall back", name)
		}
	}
}

const boolSource = `package main
func predicate(n int) bool {return n>3}
func lookup(n int) (int,bool) {return n*2,predicate(n)}
func use(n int) (bool,int) {v,ok:=lookup(n);return !ok,v}
func pure(a,b bool) bool {return !a==b}
func toggle(n int,b bool) bool {if n==0 {return b};return toggle(n-1,!b)}
func loop(n int) bool {b:=false;for i:=0;i<n;i++ {b=!b};return b}
func yes() bool {return true}
func caller() bool {return pure(yes(),false)}
func unsupported(n int) *bool {b:=predicate(n);return &b}
`

func TestBoolResults(t *testing.T) {
	interp, err := ixgo.NewContext(0).LoadInterp("bool.go", boolSource)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	pkg, err := ixgo.NewContext(0).LoadFile("bool.go", boolSource)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"predicate", "lookup", "use", "pure", "toggle", "loop", "yes", "caller"} {
		if makeTypedRunner(pkg.Func(name)) == nil {
			t.Fatalf("%s not compiled", name)
		}
	}
	for _, n := range []int{0, 1, 2, 3, 4, 9, 100} {
		for _, tc := range []struct {
			name string
			args []ixgo.Value
			want ixgo.Value
		}{
			{"predicate", []ixgo.Value{n}, n > 3},
			{"lookup", []ixgo.Value{n}, ixgo.Tuple{n * 2, n > 3}},
			{"use", []ixgo.Value{n}, ixgo.Tuple{n <= 3, n * 2}},
			{"toggle", []ixgo.Value{n, true}, n%2 == 0},
			{"loop", []ixgo.Value{n}, n%2 != 0},
			{"yes", nil, true}, {"caller", nil, true},
		} {
			got, err := interp.RunFunc(tc.name, tc.args...)
			if err != nil || !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("%s n=%d got=%v want=%v err=%v", tc.name, n, got, tc.want, err)
			}
		}
	}
	for _, a := range []bool{false, true} {
		for _, b := range []bool{false, true} {
			got, err := interp.RunFunc("pure", a, b)
			if err != nil || got != (!a == b) {
				t.Fatalf("pure: %v %v", got, err)
			}
		}
	}
	if makeTypedRunner(pkg.Func("unsupported")) != nil {
		t.Fatal("pointer bool must fall back")
	}
}

func TestNamedBoundaryFallback(t *testing.T) {
	source := `package main
type Count int
func consume(n Count) {}
func consumePointer(n *Count) {}
func result(n int) Count {return 1}
func entry() int {consume(1);n:=Count(2);consumePointer(&n);return int(result(3))}
`
	pkg, err := ixgo.NewContext(0).LoadFile("named.go", source)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"consume", "consumePointer", "result"} {
		if makeTypedRunner(pkg.Func(name)) != nil {
			t.Fatalf("%s named boundary must fall back", name)
		}
	}
	interp, err := ixgo.NewContext(0).LoadInterp("named.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	got, err := interp.RunFunc("entry")
	if err != nil || got != 1 {
		t.Fatalf("got=%v err=%v", got, err)
	}
}

func TestFloatSelection(t *testing.T) {
	ctx := ixgo.NewContext(0)
	pkg, err := ctx.LoadFile("floats.go", `package main
func single(x float32) float32 {return -x+0.5}
func double(x float64) float64 {s:=0.0;for i:=0.0;i<x;i++ {s+=i/2};return s}
func memory(x float64,p *float64) (float64,bool) {*p+=x;return *p,*p>0}
func constant() float32 {return 0.1}
func mixed(x float32,y float64) float64 {return float64(x)+y}
type F float64
func named(x F) F {return x+1}
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name        string
		input, want ixgo.Value
	}{
		{"single", float32(2), float32(-1.5)}, {"double", float64(5), float64(5)},
	} {
		r := makeTypedRunner(pkg.Func(tc.name))
		if r == nil || r.Run(tc.input) != tc.want {
			t.Fatalf("%s selection or result failed", tc.name)
		}
	}
	r, ok := makeTypedRunner(pkg.Func("constant")).(ixgo.TypedArgumentsRunner)
	if !ok || r.RunArgs(nil) != float32(0.1) {
		t.Fatal("float constant selection failed")
	}
	r, ok = makeTypedRunner(pkg.Func("memory")).(ixgo.TypedArgumentsRunner)
	if !ok {
		t.Fatal("float memory selection failed")
	}
	p := 2.0
	r.RunArgs([]ixgo.Value{3.5, &p})
	if p != 5.5 {
		t.Fatal("float pointer update failed")
	}
	for _, name := range []string{"mixed", "named"} {
		if makeTypedRunner(pkg.Func(name)) != nil {
			t.Fatalf("%s must fall back", name)
		}
	}
	interp, err := ctx.LoadInterp("entry.go", `package main
func f(x float32) (float32,float32) {if x<3 {return 0,x};a,b:=f(x-2);c,d:=f(x-1);return a+d,b+c}
func call(x float32) (float32,float32) {return f(x)}
func divide(x,y float64) float64 {return x/y}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	got, err := interp.RunFunc("call", float32(20))
	result := reflect.ValueOf(got)
	if err != nil || result.Kind() != reflect.Slice || result.Len() != 2 || result.Index(0).Interface() != float32(2584) || result.Index(1).Interface() != float32(8362) {
		t.Fatalf("float integration: %v, %v", got, err)
	}
	got, err = interp.RunFunc("divide", 1.0, 0.0)
	if err != nil || !math.IsInf(got.(float64), 1) {
		t.Fatalf("float division by zero: %v, %v", got, err)
	}
}

func TestScalarWidths(t *testing.T) {
	for _, tc := range []struct{ input, want ixgo.Value }{
		{int(7), int(22)}, {int8(7), int8(22)}, {int16(7), int16(22)},
		{int32(7), int32(22)}, {int64(7), int64(22)},
		{uint(7), uint(22)}, {uint8(7), uint8(22)}, {uint16(7), uint16(22)},
		{uint32(7), uint32(22)}, {uint64(7), uint64(22)}, {uintptr(7), uintptr(22)},
		{float32(7), float32(22)}, {float64(7), float64(22)},
	} {
		typ := reflect.TypeOf(tc.input).Name()
		t.Run(typ, func(t *testing.T) {
			pkg, err := ixgo.NewContext(0).LoadFile("width.go", "package main;func f(x "+typ+") "+typ+" {return x*3+1}")
			if err != nil {
				t.Fatal(err)
			}
			r := makeTypedRunner(pkg.Func("f"))
			if r == nil {
				t.Fatal("scalar type was not selected")
			}
			if got := r.Run(tc.input); got != tc.want {
				t.Fatalf("got %v (%T), want %v (%T)", got, got, tc.want, tc.want)
			}
		})
	}
}

func TestInvalidAndAggregateSignatures(t *testing.T) {
	for _, fn := range []*ssa.Function{nil, {}} {
		if makeTypedRunner(fn) != nil {
			t.Fatal("invalid function was selected")
		}
	}
	pkg, err := ixgo.NewContext(0).LoadFile("aggregate.go", `package main
func slice(x []int) {}
func array(x [2]int) {}
func object(x struct{N int}) {}
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"slice", "array", "object"} {
		if makeTypedRunner(pkg.Func(name)) != nil {
			t.Fatalf("%s must fall back", name)
		}
	}
}
