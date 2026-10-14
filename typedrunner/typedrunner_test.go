package typedrunner

import (
	"github.com/goplus/ixgo"
	"reflect"
	"testing"
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
