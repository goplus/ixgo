package ixgo

import (
	"go/constant"
	"go/token"
	"go/types"
	"math"
	"os"
	"reflect"
	"strings"
	"sync"
	"testing"

	pkg "github.com/goplus/ixgo/testdata/namedpkg"
	"golang.org/x/tools/go/ssa"
)

func TestTypedConstantBoundaries(t *testing.T) {
	for _, tc := range []struct {
		value int64
		ok    bool
	}{
		{-1 << 63, true}, {1<<63 - 1, true},
	} {
		got, ok := typedConstantValue[int64](constant.MakeInt64(tc.value))
		if ok != tc.ok || got != tc.value {
			t.Fatalf("int64(%d) = %d, %v", tc.value, got, ok)
		}
	}
	if _, ok := typedConstantValue[int8](constant.MakeInt64(128)); ok {
		t.Fatal("int8 overflow accepted")
	}
	if _, ok := typedConstantValue[int64](constant.MakeUint64(1 << 63)); ok {
		t.Fatal("int64 overflow accepted")
	}
	for _, n := range []int64{-1, 256} {
		if _, ok := typedConstantValue[uint8](constant.MakeInt64(n)); ok {
			t.Fatalf("unrepresentable uint8 constant %d accepted", n)
		}
	}
	if got, ok := typedConstantValue[uint64](constant.MakeUint64(^uint64(0))); !ok || got != ^uint64(0) {
		t.Fatal("MaxUint64 rejected")
	}
	pkg, err := NewContext(0).LoadFile("max.go", `package main
func maximum(n int64) int64 {return n+(1<<63-1)}
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, build := range []func(*ssa.Function) TypedRunner{makeTypedScalarRunnerOf[int64], makeTypedArgumentsRunnerOf[int64]} {
		runner := build(pkg.Func("maximum"))
		if runner == nil || runner.Run(int64(0)) != int64(1<<63-1) {
			t.Fatal("MaxInt64 function did not compile or execute correctly")
		}
	}
}

func TestTypedArgumentsFailedGraphCache(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("cache.go", `package main
func root(n int) int {if n>0 {other(n-1)};return len("unsupported")+int(float64(n))}
func other(n int) int {return root(n)}
func valid(n int) int {return n+1}
`)
	if err != nil {
		t.Fatal(err)
	}
	c := typedArgsCompiler[int]{compiled: make(map[*ssa.Function]*typedArgsRunner[int])}
	for _, name := range []string{"root", "root", "other", "other", "root"} {
		if c.build(pkg.Func(name)) != nil {
			t.Fatalf("failed graph %s was returned from cache", name)
		}
		for fn, runner := range c.compiled {
			if runner != nil {
				t.Fatalf("failed graph retained partial runner %s", fn.Name())
			}
		}
	}
	if r := c.build(pkg.Func("valid")); r == nil || r.Run(4) != 5 {
		t.Fatal("failed graph prevented independent compilation")
	}
}

func TestTypedRunnerPanicRecover(t *testing.T) {
	saved := typedRunnerFactory
	t.Cleanup(func() { RegisterTypedRunnerFactory(saved) })
	const source = `package main
func quotient(n int) int {return 10/n}
func remainder(n int) int {return 10%n}
func load(p *int) int {return *p}
func store(p *int) {*p=1}
func bounds(n int) int {a:=[2]int{1,2};return a[n]}
func direct(which,n int) (ok bool) {
 defer func(){ok=recover()!=nil}()
 p:=&n;if n==0 {p=nil}
 switch which {case 0:quotient(n);case 1:remainder(n);case 2:load(p);case 3:store(p);case 4:if n==0 {bounds(2)} else {bounds(1)}}
 return
}

func indirect(fn func(int) int) (ok bool) {defer func(){ok=recover()!=nil}();fn(0);return}
func deferred(fn func(int) int) (ok bool) {defer func(){ok=recover()!=nil}();defer fn(0);return}
func indirectEntry() bool {return indirect(quotient)}
func deferredEntry() bool {return deferred(quotient)}
func callback(invoke func(func(int)int)) (ok bool) {
 defer func(){ok=recover()=="outer"}()
 defer func(){invoke(quotient)}()
 panic("outer")
}
`
	for _, typed := range []bool{true, false} {
		RegisterTypedRunnerFactory(nil)
		if typed {
			RegisterTypedRunnerFactory(MakeTypedRunnerOf[int])
		}
		interp, err := NewContext(0).LoadInterp("recover.go", source)
		if err != nil {
			t.Fatal(err)
		}
		t.Cleanup(interp.UnsafeRelease)
		for _, name := range []string{"quotient", "remainder", "load", "store", "bounds"} {
			if (interp.funcs[interp.mainpkg.Func(name)].typedRunner != nil) != typed {
				t.Fatalf("%s typed runner selection differs from mode %v", name, typed)
			}
		}
		for which := 0; which < 5; which++ {
			for _, n := range []int{0, 1} {
				if got, err := interp.RunFunc("direct", which, n); err != nil || got != (n == 0) {
					t.Fatalf("typed=%v direct(%d,%d): got=%v err=%v", typed, which, n, got, err)
				}
			}
		}
		for _, name := range []string{"indirect", "deferred"} {
			got, err := interp.RunFunc(name + "Entry")
			if err != nil || got != true {
				t.Fatalf("typed=%v %s: got=%v err=%v", typed, name, got, err)
			}
		}
		invoke := func(fn func(int) int) {
			id := goroutineID()
			before, ok := interp.deferMap.Load(id)
			if !ok {
				t.Fatal("missing deferred caller")
			}
			func() {
				defer func() {
					if recover() == nil {
						t.Error("callback did not panic")
					}
				}()
				fn(0)
			}()
			if after, ok := interp.deferMap.Load(id); !ok || after != before {
				t.Error("callback panic changed deferred caller")
			}
		}
		if got, err := interp.RunFunc("callback", invoke); err != nil || got != true {
			t.Fatalf("typed=%v callback: got=%v err=%v", typed, got, err)
		}
		if _, ok := interp.deferMap.Load(goroutineID()); ok {
			t.Fatal("deferred caller retained after return")
		}
	}
}

func TestTypedRunnerCallResults(t *testing.T) {
	saved := typedRunnerFactory
	RegisterTypedRunnerFactory(MakeTypedRunnerOf[int])
	t.Cleanup(func() { RegisterTypedRunnerFactory(saved) })
	interp, err := NewContext(0).LoadInterp("results.go", `package main
func many(a,b,c,d,e,f,g,h,i,j int) (int,bool) {return a+b+c+d+e+f+g+h+i+j,j>0}
func pair(n int) (int,int) {return n+1,n-1}
func void(n int,p *int) {*p=n}
func predicate(n int) bool {return n>0}
func scalarVoid(n int) {}
func entry(tick func()) (int,bool,int,int) {
 tick();scalarVoid(3);p:=0;void(4,&p);a,b:=pair(3);sum,ok:=many(1,2,3,4,5,6,7,8,9,10)
 if predicate(p) {sum+=p};return sum,ok,a,b
}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, name := range []string{"many", "pair", "void", "predicate"} {
		if interp.funcs[interp.mainpkg.Func(name)].typedRunner == nil {
			t.Fatalf("%s missing runner", name)
		}
	}
	for n := 0; n < 3; n++ {
		called := false
		got, err := interp.RunFunc("entry", func() { called = true })
		if err != nil || !called || !reflect.DeepEqual(got, Tuple{59, true, 4, 2}) {
			t.Fatalf("got=%v called=%v err=%v", got, called, err)
		}
	}
}

func init() {
	RegisterTypedRunnerFactory(testTypedRunnerFactory)
}

func testTypedRunnerFactory(fn *ssa.Function) TypedRunner {
	for _, block := range fn.Blocks {
		for _, instr := range block.Instrs {
			if _, ok := instr.(*ssa.Phi); ok {
				return MakeTypedRunnerOf[int](fn)
			}
			if call, ok := instr.(*ssa.Call); ok && call.Call.Value == fn {
				return MakeTypedRunnerOf[int](fn)
			}
		}
	}
	return nil
}

func fibishTypedRunner(tb testing.TB) *typedRunner[int] {
	tb.Helper()
	pkg, err := NewContext(0).LoadFile("fibish.go", `package main
func f(x int) (int,int) {
 if x<3 {return 0,x}
 a,b:=f(x-2);c,d:=f(x-1);return a+d,b+c
}`)
	if err != nil {
		tb.Fatal(err)
	}
	runner, ok := makeTypedScalarRunnerOf[int](pkg.Func("f")).(*typedRunner[int])
	if !ok {
		tb.Fatal("fibish scalar runner missing")
	}
	return runner
}

func pairResult[T TypedScalar](runner *typedRunner[T], input T) [2]T {
	result := runner.runSelectedTyped(input)
	return [2]T{result.values[0], result.values[1]}
}

func TestTypedRunnerFibish(t *testing.T) {
	runner := fibishTypedRunner(t)
	for n, want := range []Tuple{{0, 0}, {0, 1}, {0, 2}, {2, 1}, {1, 4}, {6, 2}, {3, 10}} {
		got := pairResult(runner, n)
		if got[0] != want[0] || got[1] != want[1] {
			t.Fatalf("f(%d) = (%d, %d), want %v", n, got[0], got[1], want)
		}
	}
	// Keep exponential workloads small under race/coverage instrumentation.
	// Input 40 remains in BenchmarkTypedRunnerFibish.
	got := pairResult(runner, 20)
	if got[0] != 2584 || got[1] != 8362 {
		t.Fatalf("f(20) = (%d, %d)", got[0], got[1])
	}
}

func TestTypedRunnerGenericScalar(t *testing.T) {
	runner := &typedRunner[int64]{code: []typedOp[int64]{
		{kind: typedConst, dst: 1, value: 2},
		{kind: typedAdd, dst: 2, x: 0, y: 1},
		{kind: typedSub, dst: 3, x: 0, y: 1},
		{kind: typedReturn, nres: 2, results: []uint8{2, 3}},
	}}
	if got := pairResult(runner, 40); got != [2]int64{42, 38} {
		t.Fatalf("got (%d, %d)", got[0], got[1])
	}
}

func TestTypedScalarLoopKeepsLocalResults(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("loops.go", typedLoopSource)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"sum", "swap", "fibloop", "branch", "nested"} {
		runner, ok := MakeTypedRunnerOf[int](pkg.Func(name)).(*typedRunner[int])
		if !ok || runner.recursive {
			t.Fatalf("%s must use the ordinary scalar path", name)
		}
	}
}

func TestTypedScalarRecursiveResultRegisters(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("results.go", `package main
func zero(n int) {if n>0 {zero(n-1)}}
func one(n int) int {if n==0 {return 1};return one(n-1)+1}
func eight(n int)(int,int,int,int,int,int,int,int) {
 if n==0 {return 1,2,3,4,5,6,7,8}
 a,b,c,d,e,f,g,h:=eight(n-1)
 return b,c,d,e,f,g,h,a
}
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"zero", "one", "eight"} {
		runner, ok := MakeTypedRunnerOf[int](pkg.Func(name)).(*typedRunner[int])
		if !ok {
			t.Fatalf("%s did not use the scalar runner", name)
		}
		if !runner.recursive {
			t.Fatalf("%s did not enable recursive result registers", name)
		}
		check := func(n int) {
			var want Value
			switch name {
			case "one":
				want = n + 1
			case "eight":
				values := make(Tuple, 8)
				for i := range values {
					values[i] = (i+n)%8 + 1
				}
				want = values
			}
			if got := runner.Run(n); !reflect.DeepEqual(got, want) {
				t.Errorf("%s(%d): got=%v want=%v", name, n, got, want)
			}
		}
		// Force stack growth while every caller's output registers are live.
		for _, n := range []int{0, 1, 2, 8, 1025} {
			check(n)
		}
		var wg sync.WaitGroup
		for worker := 0; worker < 8; worker++ {
			wg.Add(1)
			go func() { defer wg.Done(); check(100) }()
		}
		wg.Wait()
	}
}

func TestTypedRunnerIntegration(t *testing.T) {
	ctx := NewContext(0)
	interp, err := ctx.LoadInterp("typed_runner_integration.go", `package main
func fibish(x int) (int, int) {
		if x < 3 { return 0, x }
		a, b := fibish(x - 2)
		c, d := fibish(x - 1)
		return a + d, b + c
}

func caller(x int) (int, int) { return fibish(x) }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	if pfn := interp.funcs[interp.mainpkg.Func("fibish")]; pfn.typedRunner == nil {
		t.Fatal("fibish was not compiled to a typed runner")
	}
	got, err := interp.RunFunc("caller", 20)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, Tuple{2584, 8362}) {
		t.Fatalf("got %v", got)
	}
}

func BenchmarkTypedRunnerFibish(b *testing.B) {
	runner := fibishTypedRunner(b)
	for _, tc := range []struct {
		name string
		n    int
		want [2]int
	}{
		{"n20", 20, [2]int{2584, 8362}},
		{"n40", 40, [2]int{39088169, 126491972}},
	} {
		b.Run(tc.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				got := pairResult(runner, tc.n)
				if got != tc.want {
					b.Fatal(got)
				}
			}
		})
	}
}

func TestTypedRunnerForwardedParameters(t *testing.T) {
	const source = `package main
type Recur func(int, Recur) (int, int)
func f(x int, unused Recur) (int, int) {
	if x < 3 { return 0, x }
	a,b := f(x-2, unused)
	c,d := f(x-1, unused)
	return a+d,b+c
}

func caller(x int) (int,int) { return f(x,f) }
func changed(x int, unused Recur) (int,int) {
	if x < 3 { return 0,x }
	return changed(x-1, nil)
}
func invoked(x int, fn Recur) (int,int) {
	if x < 3 { return fn(x,fn) }
	return invoked(x-1,fn)
}
`
	interp, err := NewContext(0).LoadInterp("forwarded.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	pfn := interp.funcs[interp.mainpkg.Func("f")]
	if pfn.typedRunner == nil {
		t.Fatal("forwarded parameter prevented typed compilation")
	}
	for _, name := range []string{"changed", "invoked"} {
		if interp.funcs[interp.mainpkg.Func(name)].typedRunner != nil {
			t.Fatalf("%s must use the normal runner", name)
		}
	}
	for _, args := range []struct {
		name string
		args []Value
	}{
		{"caller", []Value{20}},
		{"f", []Value{20, nil}},
	} {
		got, err := interp.RunFunc(args.name, args.args...)
		if err != nil || !reflect.DeepEqual(got, Tuple{2584, 8362}) {
			t.Fatalf("%s: got %v, %v", args.name, got, err)
		}
	}
}

func TestTypedRunnerResultShapes(t *testing.T) {
	const source = `package main
func one(x int) int {
	if x < 1 { return x }
	return one(x-1) + 1
}
func zero(x int) {
	if x < 1 { return }
	zero(x-1)
}
func three(x int) (int, int, int) {
	if x < 1 { return 0, 1, 2 }
	a,b,c := three(x-1)
	return a+1,b+1,c+1
}
func ops(x int) (int, int) {
	if x <= 0 { return 7, 3 }
	a,b := ops(x-1)
	y := ((a*2 + b/2) % 32) &^ 1
	z := ((a | b) ^ 3) << 1
	if y >= z { return y, z }
	return z, y
}
`
	interp, err := NewContext(0).LoadInterp("result_shapes.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, name := range []string{"one", "zero", "three", "ops"} {
		if interp.funcs[interp.mainpkg.Func(name)].typedRunner == nil {
			t.Fatalf("%s was not compiled to a typed runner", name)
		}
	}
	if got, err := interp.RunFunc("one", 20); err != nil || got != 20 {
		t.Fatalf("one: got %v, %v", got, err)
	}
	if got, err := interp.RunFunc("zero", 20); err != nil || got != nil {
		t.Fatalf("zero: got %v, %v", got, err)
	}
	if got, err := interp.RunFunc("three", 20); err != nil || !reflect.DeepEqual(got, Tuple{20, 21, 22}) {
		t.Fatalf("three: got %v, %v", got, err)
	}
	fast := interp.funcs[interp.mainpkg.Func("ops")]
	got, err := interp.RunFunc("ops", 10)
	if err != nil {
		t.Fatal(err)
	}
	fast.typedRunner = nil
	want, err := interp.RunFunc("ops", 10)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("ops: typed=%v boxed=%v err=%v", got, want, err)
	}
}

const typedLoopSource = `package main
func sum(n int) int {
 s := 0
 for i := 0; i < n; i++ { s += i }
 return s
}
func swap(n int) (int,int) {
 a,b := 1,2
 for i := 0; i < n; i++ { a,b = b,a }
 return a,b
}
func fibloop(n int) (int,int) {
 a,b := 0,1
 for i := 0; i < n; i++ { a,b = b,a+b }
 return a,b
}
func branch(n int) int {
 s := 0
 for i := 0; i < n; i++ {
  if i & 1 == 0 { s += i } else { s += 3 }
 }
 return s
}
func exits(n int) int {
 s := 0
 for i := 0; i < n; i++ {
  if i > 12 { break }
  if i & 1 == 0 { continue }
  s += i
 }
 return s
}
func nested(n int) int {
 s := 0
 for i := 0; i < n; i++ {
  for j := 0; j < i; j++ { s += j }
 }
 return s
}
func boolphi(n int) int {
 flag := true
 s := 0
 for i := 0; i < n; i++ {
  if flag { s += i }
  flag = i < 3
 }
 return s
}
func recursiveLoop(n int) int {
 if n < 1 { return 0 }
 s := recursiveLoop(n-1)
 for i := 0; i < n; i++ { s += i }
 return s
}
func join(n int) int {
 x := 7
 if n > 0 { x = n+3 }
 return x*2
}
func boolcompare(n int) int {
 a := n > 2
 b := n > 3
 if a == b { return 1 }
 return 0
}
func pointerphi(n int) int {
 a,b := 3,5
 p := &a
 if n > 0 { p = &b }
 return *p
}
`

func TestTypedRunnerPhi(t *testing.T) {
	interp, err := NewContext(0).LoadInterp("phi.go", typedLoopSource)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, name := range []string{"sum", "swap", "fibloop", "branch", "exits", "nested", "boolphi", "recursiveLoop", "join"} {
		t.Run(name, func(t *testing.T) {
			fn := interp.mainpkg.Func(name)
			pfn := interp.funcs[fn]
			runner := pfn.typedRunner
			if runner == nil {
				t.Fatal("not compiled to typed runner")
			}
			for _, n := range []int{-1, 0, 1, 2, 3, 4, 7, 13, 20} {
				pfn.typedRunner = runner
				got, err := interp.RunFunc(name, n)
				if err != nil {
					t.Fatal(err)
				}
				pfn.typedRunner = nil
				want, err := interp.RunFunc(name, n)
				if err != nil || !reflect.DeepEqual(got, want) {
					t.Fatalf("n=%d: typed=%v boxed=%v err=%v", n, got, want, err)
				}
				if name == "swap" {
					expected := Tuple{1, 2}
					if n > 0 && n%2 == 1 {
						expected = Tuple{2, 1}
					}
					if !reflect.DeepEqual(got, expected) {
						t.Fatalf("parallel copies: n=%d got=%v", n, got)
					}
				}
			}
			pfn.typedRunner = runner
		})
	}
	for _, name := range []string{"pointerphi"} {
		if makeTypedScalarRunnerOf[int](interp.mainpkg.Func(name)) != nil {
			t.Fatalf("unsupported %s compiled", name)
		}
	}
}

func BenchmarkTypedRunnerLoops(b *testing.B) {
	for _, name := range []string{"sum", "swap", "fibloop", "branch", "nested"} {
		for _, typed := range []bool{true, false} {
			mode := "boxed"
			if typed {
				mode = "typed"
			}
			b.Run(name+"/"+mode, func(b *testing.B) {
				interp, err := NewContext(0).LoadInterp("loops.go", typedLoopSource)
				if err != nil {
					b.Fatal(err)
				}
				defer interp.UnsafeRelease()
				pfn := interp.funcs[interp.mainpkg.Func(name)]
				if pfn.typedRunner == nil {
					b.Fatal("typed runner missing")
				}
				n := 1000
				if name == "nested" {
					n = 50
				}
				if !typed {
					pfn.typedRunner = nil
				}
				want, err := interp.RunFunc(name, n)
				if err != nil {
					b.Fatal(err)
				}
				b.ReportAllocs()
				for b.Loop() {
					got, err := interp.RunFunc(name, n)
					if err != nil || !reflect.DeepEqual(got, want) {
						b.Fatalf("got=%v want=%v err=%v", got, want, err)
					}
				}
			})
		}
	}
}

const typedArgsSource = `package main
func pair(a,b int) (int,int) {return a+b,a-b}
func calls(a,b int) int {x,y:=pair(a,b);return x*y}
func memory(n int,p,q *int) int {
 a:=[3]int{n,2,3}; b:=a; a[0]=99
 *p+=b[0];*q+=b[1]
 return b[0]+b[2]+*p+*q
}
func pointerPhi(n int,p,q *int) int {for i:=0;i<n;i++ {p,q=q,p};return *p-*q}
func unary(n int,p *int) int {if !(n<0) {*p= -n ^ 3};return *p}
func index(n int,p *int) int {a:=[3]int{1,2,3};*p=a[n];return *p}
func alias(p,q *int) int {if p==q {return 1};return 0}
func deep(n int,p *int) {if n==0 {return};a:=n; other(n-1,&a,p);*p+=a}
func other(n int,p,q *int) {deep(n,q);*p+=1}
func escape(n int) *int {return &n}
func badcallee(n int) int {return external(n)}
func external(n int) int {defer func(){}();return n}
func cycle(n int) int {s:=0;for i:=0;i<n;i++ {p:=new(int);*p=i;s+=*p};return s}
`

func TestTypedArgumentsRunner(t *testing.T) {
	interp, err := NewContext(0).LoadInterp("args.go", typedArgsSource)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, name := range []string{"calls", "memory", "pointerPhi", "unary", "index", "alias", "deep"} {
		t.Run(name, func(t *testing.T) {
			runner, ok := MakeTypedRunnerOf[int](interp.mainpkg.Func(name)).(TypedArgumentsRunner)
			if !ok {
				t.Fatal("not compiled")
			}
			for _, n := range []int{0, 1, 2, 3, 9, 100} {
				var a, b int = 7, 11
				makeArgs := func() []Value {
					switch name {
					case "calls":
						return []Value{n, 5}
					case "memory", "pointerPhi":
						return []Value{n, &a, &b}
					case "alias":
						return []Value{&a, &a}
					case "index":
						return []Value{n % 3, &a}
					default:
						return []Value{n, &a}
					}
				}
				got := runner.RunArgs(makeArgs())
				typedA, typedB := a, b
				a, b = 7, 11
				for _, fn := range interp.funcs {
					fn.typedRunner = nil
				}
				want, err := interp.RunFunc(name, makeArgs()...)
				if err != nil || !reflect.DeepEqual(got, want) || a != typedA || b != typedB {
					t.Fatalf("n=%d: typed=%v (%d,%d) boxed=%v (%d,%d) err=%v", n, got, typedA, typedB, want, a, b, err)
				}
			}
			if name == "memory" {
				a := 7
				got := runner.RunArgs([]Value{3, &a, &a})
				if a != 12 || got != 30 {
					t.Fatalf("aliasing: %v %d", got, a)
				}
			}
		})
	}

	for _, name := range []string{"escape", "badcallee"} {
		if MakeTypedRunnerOf[int](interp.mainpkg.Func(name)) != nil {
			t.Fatalf("%s must fall back", name)
		}
	}
	// A single immutable runner is safe to use concurrently, including arena growth.
	runner := MakeTypedRunnerOf[int](interp.mainpkg.Func("deep")).(TypedArgumentsRunner)
	var wg sync.WaitGroup
	for j := 0; j < 8; j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			p := 0
			runner.RunArgs([]Value{100, &p})
			if p != 5150 {
				t.Errorf("deep result=%d", p)
			}
		}()
	}
	wg.Wait()
}

func TestTypedArgumentsRunnerPanics(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("args.go", typedArgsSource)
	if err != nil {
		t.Fatal(err)
	}
	for _, tc := range []struct {
		name string
		args []Value
	}{
		{"index", []Value{-1, new(int)}}, {"index", []Value{3, new(int)}}, {"unary", []Value{1, (*int)(nil)}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			runner := MakeTypedRunnerOf[int](pkg.Func(tc.name)).(TypedArgumentsRunner)
			defer func() {
				if recover() == nil {
					t.Error("expected panic")
				}
			}()
			runner.RunArgs(tc.args)
		})
	}
}

func TestTypedArgumentsRunnerUglyFib(t *testing.T) {
	source, err := os.ReadFile("testdata/fibish/uglyfib.go")
	if err != nil {
		t.Fatal(err)
	}
	interp, err := NewContext(0).LoadInterp("ugly.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, name := range []string{"f", "g", "h", "k"} {
		runner, ok := MakeTypedRunnerOf[int](interp.mainpkg.Func(name)).(TypedArgumentsRunner)
		if !ok {
			t.Fatalf("%s not compiled", name)
		}
		for n := 0; n <= 16; n++ {
			a, b, p := n-1, n-2, 0
			runner.RunArgs([]Value{n, &a, &b, &p})
			got := p
			a, b, p = n-1, n-2, 0
			for _, fn := range interp.funcs {
				fn.typedRunner = nil
			}
			_, err := interp.RunFunc(name, n, &a, &b, &p)
			if err != nil || got != p {
				t.Fatalf("%s(%d): typed=%d boxed=%d err=%v", name, n, got, p, err)
			}
		}
	}
}

func BenchmarkTypedArgumentsUglyFib(b *testing.B) {
	source, err := os.ReadFile("testdata/fibish/uglyfib.go")
	if err != nil {
		b.Fatal(err)
	}
	for _, typed := range []bool{false, true} {
		name := "boxed"
		if typed {
			name = "typed"
		}
		b.Run(name, func(b *testing.B) {
			interp, err := NewContext(0).LoadInterp("ugly.go", source)
			if err != nil {
				b.Fatal(err)
			}
			defer interp.UnsafeRelease()
			for _, fn := range interp.funcs {
				fn.typedRunner = nil
			}
			if typed {
				interp.funcs[interp.mainpkg.Func("f")].typedRunner = MakeTypedRunnerOf[int](interp.mainpkg.Func("f"))
			}
			b.ReportAllocs()
			for b.Loop() {
				a, c, p := 19, 18, 0
				_, err := interp.RunFunc("f", 20, &a, &c, &p)
				if err != nil || p != 6765 {
					b.Fatalf("got %d, err=%v", p, err)
				}
			}
		})
	}
}

func TestTypedArgumentsGroupedCallRegisters(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("grouped.go", `package main
func mixed(b bool, p *int, n int, q *int, c bool) (int,bool) {
 *p += n
 *q += 2
 return *p+*q,b!=c
}
func run(n int,p *int)(int,bool) {
 a,b:=mixed(true,p,n,p,false)
 if n>0 {x,c:=run(n-1,p);return a+x,b!=c}
 return a,b
}
`)
	if err != nil {
		t.Fatal(err)
	}
	runner, ok := MakeTypedRunnerOf[int](pkg.Func("run")).(TypedArgumentsRunner)
	if !ok {
		t.Fatal("argument runner missing")
	}
	for _, n := range []int{0, 1, 5, 20} {
		value, wantValue, wantSum := 1, 1, 0
		for i := n; i >= 0; i-- {
			wantValue += i + 2
			wantSum += 2 * wantValue
		}
		want := Tuple{wantSum, n%2 == 0}
		got := runner.RunArgs([]Value{n, &value})
		if !reflect.DeepEqual(got, want) || value != wantValue {
			t.Fatalf("n=%d: got=%v value=%d want=%v value=%d", n, got, value, want, wantValue)
		}
	}
}

func BenchmarkTypedArgumentsRunArgs(b *testing.B) {
	pkg, err := NewContext(0).LoadFile("arena.go", `package main
func shallow(n int,p *int) int {*p=n+1;return *p}
func mixed(n int,p *int) (int,bool) {*p=n+1;return *p,n>0}
func deep(n int,p *int) int {if n==0 {return *p};return deep(n-1,p)+1}
`)
	if err != nil {
		b.Fatal(err)
	}
	for _, name := range []string{"shallow", "mixed", "deep"} {
		b.Run(name, func(b *testing.B) {
			runner := MakeTypedRunnerOf[int](pkg.Func(name)).(TypedArgumentsRunner)
			p := 1
			args := []Value{100, &p}
			var want Value = 101
			if name == "mixed" {
				want = Tuple{101, true}
			}
			b.ReportAllocs()
			for b.Loop() {
				if got := runner.RunArgs(args); !reflect.DeepEqual(got, want) {
					b.Fatalf("got=%v want=%v", got, want)
				}
			}
		})
	}
}

func TestTypedRunnerUnaryAndBoolean(t *testing.T) {
	source := `package main
func unary(n int) int {s:=0; for i:=0;i<n;i++ {if !(i<2) {s+= -i + ^i}};return s}
func neq(n int) int {a:=n>2;b:=n>3;if a!=b {return 1};return 0}
func bits(n uint8) uint8 {return -n + ^n}
func wideShift(n uint8) uint8 {return n << 256}
func mixedShift(n int) int {return n << uint64(1)}
func uintConstant(n uint64) uint64 {return n ^ 18446744073709551615}
func arrayIndex(n uint8,p *uint8) uint8 {a:=[2]uint8{3,5};*p=a[int(n)+256];return *p}
`
	interp, err := NewContext(0).LoadInterp("unary.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, name := range []string{"unary", "neq"} {
		fn := interp.mainpkg.Func(name)
		runner := MakeTypedRunnerOf[int](fn)
		if _, ok := runner.(*typedRunner[int]); !ok {
			t.Fatalf("%s did not use scalar runner", name)
		}
		for _, n := range []int{-3, 0, 1, 2, 3, 4, 10} {
			got := runner.Run(n)
			interp.funcs[fn].typedRunner = nil
			want, err := interp.RunFunc(name, n)
			if err != nil || !reflect.DeepEqual(got, want) {
				t.Fatalf("%s(%d): typed=%v boxed=%v err=%v", name, n, got, want, err)
			}
		}
	}
	bits := MakeTypedRunnerOf[uint8](interp.mainpkg.Func("bits"))
	if _, ok := bits.(*typedRunner[uint8]); !ok {
		t.Fatal("uint8 unary not scalar")
	}
	for _, n := range []uint8{0, 1, 127, 128, 255} {
		if got := bits.Run(n); got != -n+^n {
			t.Fatalf("bits(%d)=%v", n, got)
		}
	}
	if MakeTypedRunnerOf[uint8](interp.mainpkg.Func("arrayIndex")) != nil {
		t.Fatal("wide array index must fall back")
	}
	big := MakeTypedRunnerOf[uint64](interp.mainpkg.Func("uintConstant"))
	if big == nil || big.Run(uint64(7)) != ^uint64(7) {
		t.Fatal("full width uint64 constant failed")
	}
	// Bool parameters and ! also run in the multiargument memory path.
	pkg, err := NewContext(0).LoadFile("boolargs.go", `package main
func f(n int,a,b bool) int {if !a == b {return n};return -n}
`)
	if err != nil {
		t.Fatal(err)
	}
	runner := MakeTypedRunnerOf[int](pkg.Func("f")).(TypedArgumentsRunner)
	for _, a := range []bool{false, true} {
		for _, b := range []bool{false, true} {
			want := -7
			if !a == b {
				want = 7
			}
			if got := runner.RunArgs([]Value{7, a, b}); got != want {
				t.Fatalf("bool args: %v %v got=%v want=%v", a, b, got, want)
			}
		}
	}
}

func TestTypedRunnerNamedIntegers(t *testing.T) {
	ctx := NewContext(0)
	spkg, err := ctx.LoadFile("named.go", `package main
import "github.com/goplus/ixgo/testdata/namedpkg"
func count(n pkg.Count) pkg.Count { if n < 1 { return n }; return count(n-1)+1 }
func bits(n pkg.Bits) pkg.Bits { return (n ^ 3) + 1 }
func add(a, b pkg.Count) pkg.Count { return a + b }
`)
	if err != nil {
		t.Fatal(err)
	}
	runner, ok := MakeTypedRunnerOf[pkg.Count](spkg.Func("count")).(*typedRunner[pkg.Count])
	if !ok {
		t.Fatal("named ~int did not select the scalar typed runner")
	}
	for _, input := range []pkg.Count{-2, 0, 1, 9} {
		if got := runner.Run(input); got != input {
			t.Fatalf("count(%d) = %d, want %d", input, got, input)
		}
	}
	bits, ok := MakeTypedRunnerOf[pkg.Bits](spkg.Func("bits")).(*typedRunner[pkg.Bits])
	if !ok {
		t.Fatal("named ~uint8 did not select the scalar typed runner")
	}
	for input, want := range map[pkg.Bits]pkg.Bits{0: 4, 1: 3, 5: 7, 255: 253} {
		if got := bits.Run(input); got != want {
			t.Fatalf("bits(%d) = %d, want %d", input, got, want)
		}
	}
	args, ok := MakeTypedRunnerOf[pkg.Count](spkg.Func("add")).(*typedArgsRunner[pkg.Count])
	if !ok {
		t.Fatal("named ~int did not select the argument typed runner")
	}
	for _, tc := range []struct {
		a, b, want pkg.Count
	}{{2, 3, 5}, {-4, 9, 5}, {0, -7, -7}} {
		if got := args.RunArgs([]Value{tc.a, tc.b}); got != tc.want {
			t.Fatalf("add(%d,%d) = %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
	saved := typedRunnerFactory
	RegisterTypedRunnerFactory(MakeTypedRunnerOf[pkg.Count])
	t.Cleanup(func() { RegisterTypedRunnerFactory(saved) })
	interp, err := ctx.LoadInterp("named.go", `package main
import "github.com/goplus/ixgo/testdata/namedpkg"
func count(n pkg.Count) pkg.Count { if n < 1 { return n }; return count(n-1)+1 }
func add(a, b pkg.Count) pkg.Count { return a + b }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, tc := range []struct {
		name string
		args []Value
		want Value
	}{
		{"count", []Value{pkg.Count(9)}, pkg.Count(9)},
		{"add", []Value{pkg.Count(-4), pkg.Count(9)}, pkg.Count(5)},
	} {
		got, err := interp.RunFunc(tc.name, tc.args...)
		if err != nil || got != tc.want {
			t.Fatalf("interpreted %s = %v, want %v (err=%v)", tc.name, got, tc.want, err)
		}
	}
}

func TestTypedRunnerFloats(t *testing.T) {
	t.Run("float32", testTypedRunnerFloat[float32])
	t.Run("float64", testTypedRunnerFloat[float64])
}

func testTypedRunnerFloat[T interface{ ~float32 | ~float64 }](t *testing.T) {
	var zero T
	typ := reflect.TypeOf(zero).Name()
	source := strings.ReplaceAll(`package main
func arithmetic(x FLOAT) FLOAT {return -((x+0.1)*1.25-0.5)/2}
func neg(x FLOAT) FLOAT {return -x}
func loop(n FLOAT) FLOAT {s:=FLOAT(0); for i:=FLOAT(0);i<n;i++ {s+=i/2};return s}
func quotient(x,y FLOAT) FLOAT {return x/y}
func compare(x,y FLOAT) (bool,bool,bool,bool,bool,bool) {return x<y,x<=y,x>y,x>=y,x==y,x!=y}
func memory(x FLOAT,p *FLOAT) FLOAT {a:=[3]FLOAT{x,0.1,0.25};b:=a;a[0]=99;*p+=b[0];return b[1]+b[2]+*p}
func mixed(x FLOAT,n int) FLOAT {for i:=0;i<n;i++ {x+=0.5};return x}
func convert(x FLOAT) int {return int(x)}
type Named FLOAT
func named(x Named) Named {return -x}
`, "FLOAT", typ)
	interp, err := NewContext(0).LoadInterp("float.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	values := []T{0, T(math.Copysign(0, -1)), 0.1, -2.5, T(math.Inf(1)), T(math.Inf(-1)), T(math.NaN())}
	for _, name := range []string{"arithmetic", "neg", "loop"} {
		fn := interp.mainpkg.Func(name)
		for _, build := range []func(*ssa.Function) TypedRunner{makeTypedScalarRunnerOf[T], makeTypedArgumentsRunnerOf[T]} {
			runner := build(fn)
			if runner == nil {
				t.Fatalf("%s did not compile", name)
			}
			inputs := values
			if name == "loop" {
				inputs = []T{-1, 0, 1, 5, 20}
			}
			for _, x := range inputs {
				got := runner.Run(x).(T)
				want, err := interp.RunFunc(name, x)
				if err != nil {
					t.Fatal(err)
				}
				checkTypedFloat(t, got, want.(T))
			}
		}
	}
	for _, name := range []string{"quotient", "compare"} {
		runner, ok := MakeTypedRunnerOf[T](interp.mainpkg.Func(name)).(TypedArgumentsRunner)
		if !ok {
			t.Fatalf("%s did not compile", name)
		}
		for _, x := range values {
			for _, y := range values {
				got := runner.RunArgs([]Value{x, y})
				want, err := interp.RunFunc(name, x, y)
				if err != nil {
					t.Fatal(err)
				}
				if name == "quotient" {
					checkTypedFloat(t, got.(T), want.(T))
				} else if !reflect.DeepEqual(got, want) {
					t.Fatalf("compare(%v,%v): %v != %v", x, y, got, want)
				}
			}
		}
	}
	runner, ok := MakeTypedRunnerOf[T](interp.mainpkg.Func("memory")).(TypedArgumentsRunner)
	if !ok {
		t.Fatal("float memory did not compile")
	}
	p, q := T(2), T(2)
	got := runner.RunArgs([]Value{T(3.5), &p})
	want, err := interp.RunFunc("memory", T(3.5), &q)
	if err != nil {
		t.Fatal(err)
	}
	checkTypedFloat(t, got.(T), want.(T))
	checkTypedFloat(t, p, q)
	for _, name := range []string{"mixed", "convert", "named"} {
		if MakeTypedRunnerOf[T](interp.mainpkg.Func(name)) != nil {
			t.Fatalf("%s must fall back", name)
		}
	}
	// Test the actual standalone fixture, including recursive pair results.
	data, err := os.ReadFile("testdata/fibish/fibish_" + typ + ".go")
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := NewContext(0).LoadFile("fibish.go", string(data))
	if err != nil {
		t.Fatal(err)
	}
	for _, build := range []func(*ssa.Function) TypedRunner{makeTypedScalarRunnerOf[T], makeTypedArgumentsRunnerOf[T]} {
		r := build(pkg.Func("f"))
		if r == nil {
			t.Fatal("float fibish did not compile")
		}
		if got := r.Run(T(20)); !reflect.DeepEqual(got, tuple{T(2584), T(8362)}) {
			t.Fatal(got)
		}
	}
}

func checkTypedFloat[T TypedScalar](t *testing.T, got, want T) {
	t.Helper()
	g, w := float64(got), float64(want)
	if math.IsNaN(g) && math.IsNaN(w) {
		return
	}
	if math.Float64bits(g) != math.Float64bits(w) {
		t.Fatalf("got %v (%x), want %v (%x)", got, math.Float64bits(g), want, math.Float64bits(w))
	}
}

func TestTypedRunnerFloatConstants(t *testing.T) {
	for _, literal := range []string{"0.1", "1e-45", "3.4028234663852886e38"} {
		v := constant.MakeFromLiteral(literal, token.FLOAT, 0)
		got, ok := typedConstantValue[float32](v)
		want, _ := constant.Float32Val(v)
		if !ok || got != want {
			t.Fatalf("float32(%s): %v, %v", literal, got, ok)
		}
	}
	if _, ok := typedConstantValue[float32](constant.MakeFromLiteral("1e39", token.FLOAT, 0)); ok {
		t.Fatal("float32 overflow accepted")
	}
	if _, ok := typedConstantValue[float64](constant.MakeFromLiteral("1e309", token.FLOAT, 0)); ok {
		t.Fatal("float64 overflow accepted")
	}
	if _, ok := typedConstantValue[float32](constant.MakeInt64(1<<24 + 1)); ok {
		t.Fatal("rounded integer index accepted")
	}
	type MyFloat float32
	pkg, err := NewContext(0).LoadFile("named.go", `package main;type F float32;func f(x F) F {return -x+0.5}`)
	if err != nil {
		t.Fatal(err)
	}
	if r := MakeTypedRunnerOf[MyFloat](pkg.Func("f")); r == nil || r.Run(MyFloat(2)) != MyFloat(-1.5) {
		t.Fatal("named float failed")
	}
}

func BenchmarkTypedRunnerFloatFibish(b *testing.B) {
	b.Run("float32", benchmarkTypedRunnerFloatFibish[float32])
	b.Run("float64", benchmarkTypedRunnerFloatFibish[float64])
}

func benchmarkTypedRunnerFloatFibish[T interface{ ~float32 | ~float64 }](b *testing.B) {
	var zero T
	data, err := os.ReadFile("testdata/fibish/fibish_" + reflect.TypeOf(zero).Name() + ".go")
	if err != nil {
		b.Fatal(err)
	}
	pkg, err := NewContext(0).LoadFile("fibish.go", string(data))
	if err != nil {
		b.Fatal(err)
	}
	runner, ok := MakeTypedRunnerOf[T](pkg.Func("f")).(*typedRunner[T])
	if !ok {
		b.Fatal("scalar float runner missing")
	}
	b.ReportAllocs()
	for b.Loop() {
		got := runner.runTyped(T(20))
		if got.n != 2 || got.values[0] != T(2584) || got.values[1] != T(8362) {
			b.Fatal(got)
		}
	}
}

func TestTypedIntegerInstructionsAfterFloatSupport(t *testing.T) {
	t.Run("int8", func(t *testing.T) { testTypedIntegerInstructions(t, []int8{-128, -7, -1, 0, 1, 7, 127}) })
	t.Run("int64", func(t *testing.T) { testTypedIntegerInstructions(t, []int64{-1 << 63, -7, -1, 0, 1, 7, 1<<63 - 1}) })
	t.Run("uint8", func(t *testing.T) { testTypedIntegerInstructions(t, []uint8{0, 1, 7, 127, 128, 255}) })
	t.Run("uint64", func(t *testing.T) { testTypedIntegerInstructions(t, []uint64{0, 1, 7, 1 << 63, 1<<64 - 1}) })
}

func testTypedIntegerInstructions[T TypedScalar](t *testing.T, values []T) {
	var zero T
	typ := reflect.TypeOf(zero).Name()
	source := strings.ReplaceAll(`package main
func bits(x,y INT) (INT,INT,INT,INT,INT) {return x&y,x|y,x^y,x&^y,^x}
func rem(x,y INT) INT {return x%y}
func shl(x,y INT) INT {return x<<y}
func shr(x,y INT) INT {return x>>y}
`, "INT", typ)
	interp, err := NewContext(0).LoadInterp("ints.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, name := range []string{"bits", "rem", "shl", "shr"} {
		r, ok := MakeTypedRunnerOf[T](interp.mainpkg.Func(name)).(TypedArgumentsRunner)
		if !ok {
			t.Fatalf("%s did not compile", name)
		}
		for _, x := range values {
			for _, y := range values {
				if name == "rem" && y == 0 || (name == "shl" || name == "shr") && y < 0 {
					continue
				}
				got := r.RunArgs([]Value{x, y})
				want, err := interp.RunFunc(name, x, y)
				if err != nil || !reflect.DeepEqual(got, want) {
					t.Fatalf("%s(%v,%v): typed=%v boxed=%v err=%v", name, x, y, got, want, err)
				}
			}
		}
		if name == "rem" || name == "shl" || name == "shr" {
			y := zero
			if name != "rem" {
				y = values[0]
				if y >= 0 {
					continue
				}
			}
			func() {
				defer func() {
					if recover() == nil {
						t.Errorf("%s did not panic", name)
					}
				}()
				r.RunArgs([]Value{T(1), y})
			}()
		}
	}
}

func BenchmarkTypedIntegerOperators(b *testing.B) {
	b.Run("int", benchmarkTypedIntegerOperators[int])
	b.Run("uint64", benchmarkTypedIntegerOperators[uint64])
}

func benchmarkTypedIntegerOperators[T TypedScalar](b *testing.B) {
	var zero T
	source := strings.ReplaceAll(`package main
func rem(n INT) INT {s:=INT(0);for i:=INT(1);i<n;i++ {s+=i%7};return s}
func bits(n INT) INT {s:=INT(0);for i:=INT(1);i<n;i++ {s+=(i&31)^(i|7)};return s}
func shl(n INT) INT {s:=INT(0);for i:=INT(1);i<n;i++ {s+=i<<3};return s}
func shr(n INT) INT {s:=INT(0);for i:=INT(1);i<n;i++ {s+=i>>3};return s}
`, "INT", reflect.TypeOf(zero).Name())
	pkg, err := NewContext(0).LoadFile("operators.go", source)
	if err != nil {
		b.Fatal(err)
	}
	for _, name := range []string{"rem", "bits", "shl", "shr"} {
		b.Run(name, func(b *testing.B) {
			runner, ok := MakeTypedRunnerOf[T](pkg.Func(name)).(*typedRunner[T])
			if !ok {
				b.Fatal("scalar integer runner missing")
			}
			n := 1000
			want := runner.runTyped(T(n))
			b.ReportAllocs()
			for b.Loop() {
				if got := runner.runTyped(T(n)); got != want {
					b.Fatal(got)
				}
			}
		})
	}
}

func TestTypedScalarBooleanAndShiftPaths(t *testing.T) {
	interp, err := NewContext(0).LoadInterp("scalar.go", `package main
func equal(n int) int {a,b:=n>2,n>3;if !a==b {return n>>1};return -n}
func unequal(n int) int {a,b:=n>2,n>3;if a!=b {return n>>2};return n}
func numericUnequal(n int) int {if n!=2 {return n};return -n}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, name := range []string{"equal", "unequal", "numericUnequal"} {
		fn := interp.mainpkg.Func(name)
		r, ok := makeTypedScalarRunnerOf[int](fn).(*typedRunner[int])
		if !ok {
			t.Fatalf("%s did not compile to a scalar runner", name)
		}
		if r.Run("invalid argument") != nil {
			t.Fatal("invalid scalar argument was accepted")
		}
		interp.funcs[fn].typedRunner = nil
		for _, n := range []int{-8, 0, 1, 2, 3, 4, 8} {
			got := r.Run(n)
			want, err := interp.RunFunc(name, n)
			if err != nil || !reflect.DeepEqual(got, want) {
				t.Fatalf("%s(%d): typed=%v boxed=%v err=%v", name, n, got, want, err)
			}
		}
	}
}

func TestTypedArgumentsBooleanAndNilPaths(t *testing.T) {
	interp, err := NewContext(0).LoadInterp("boolargs.go", `package main
func toggle(n int,b bool) bool {if n==0 {return b};return toggle(n-1,!b)}
func entry(n int) bool {return toggle(n,true)}
func different(a,b bool) bool {return a!=b}
func pointers(p,q *int) (bool,bool) {return p!=nil,p!=q}
func load(p *int) int {return *p}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	// Compare against the ordinary interpreter, including nested bool arguments
	// and results transferred between typed frames.
	for _, fn := range interp.funcs {
		fn.typedRunner = nil
	}
	p := 7
	for _, tc := range []struct {
		name string
		args []Value
	}{
		{"entry", []Value{0}}, {"entry", []Value{1}}, {"entry", []Value{20}},
		{"different", []Value{true, false}}, {"different", []Value{false, false}},
		{"pointers", []Value{&p, &p}}, {"pointers", []Value{&p, (*int)(nil)}},
		{"pointers", []Value{(*int)(nil), (*int)(nil)}},
	} {
		r, ok := makeTypedArgumentsRunnerOf[int](interp.mainpkg.Func(tc.name)).(TypedArgumentsRunner)
		if !ok {
			t.Fatalf("%s did not compile", tc.name)
		}
		got := r.RunArgs(tc.args)
		want, err := interp.RunFunc(tc.name, tc.args...)
		if err != nil || !reflect.DeepEqual(got, want) {
			t.Fatalf("%s: typed=%v boxed=%v err=%v", tc.name, got, want, err)
		}
	}
	r := makeTypedArgumentsRunnerOf[int](interp.mainpkg.Func("load")).(TypedArgumentsRunner)
	for _, tc := range []struct {
		name string
		args []Value
	}{
		{"nil dereference", []Value{(*int)(nil)}}, {"missing argument", nil},
		{"extra argument", []Value{&p, &p}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatal("invalid call did not panic")
				}
			}()
			r.RunArgs(tc.args)
		})
	}
}

func TestTypedRunnerFallbackLimits(t *testing.T) {
	source := `package main
type Count int
func namedPointer(p *Count) int {return int(*p)}
func slice(p []int) int {return len(p)}
func large(n int) int {a:=[65]int{n};return a[0]}
func copies(n int) int {a:=[40]int{n};b:=a;return b[0]}
func booleanAllocation(n int) int {p:=new(bool);*p=n>0;return booleanLoad(p)}
func booleanLoad(p *bool) int {if *p {return 1};return 0}
func pointerReturn(n int) *int {return &n}
func exhausted(n int) int {
` + strings.Repeat("n=n*3+1\n", 40) + `return n}
`
	pkg, err := NewContext(0).LoadFile("limits.go", source)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"namedPointer", "slice", "large", "copies", "booleanAllocation", "exhausted"} {
		t.Run(name, func(t *testing.T) {
			if makeTypedScalarRunnerOf[int](pkg.Func(name)) != nil {
				t.Fatal("scalar runner accepted unsupported function")
			}
			if makeTypedArgumentsRunnerOf[int](pkg.Func(name)) != nil {
				t.Fatal("argument runner accepted unsupported function")
			}
		})
	}
	// Pointer results are rejected by the scalar path even when its input is T.
	if makeTypedScalarRunnerOf[int](pkg.Func("pointerReturn")) != nil {
		t.Fatal("scalar pointer return was accepted")
	}
}

func TestTypedRejectedConstantAndOperators(t *testing.T) {
	for _, value := range []constant.Value{
		constant.MakeInt64(-129),
		constant.MakeFromLiteral("18446744073709551616", token.INT, 0),
		constant.MakeFromLiteral("0.5", token.FLOAT, 0),
	} {
		if _, ok := typedConstantValue[int8](value); ok {
			t.Fatalf("int8 accepted %v", value)
		}
	}
	if _, ok := typedConstantValue[uint8](constant.MakeUint64(1 << 63)); ok {
		t.Fatal("uint8 accepted an overflowing unsigned constant")
	}
	if _, ok := typedConstantValue[float64](constant.MakeBool(true)); ok {
		t.Fatal("float64 accepted a bool constant")
	}
	for _, c := range []*ssa.Const{
		ssa.NewConst(nil, types.NewPointer(types.Typ[types.Int])),
		ssa.NewConst(nil, types.Typ[types.Invalid]),
		ssa.NewConst(constant.MakeString("text"), types.Typ[types.String]),
	} {
		if _, ok := typedConstOp[int](c, 0, false); ok {
			t.Fatalf("scalar runner accepted %v", c)
		}
	}
	if isTypedInteger(types.NewSlice(types.Typ[types.Int])) {
		t.Fatal("slice accepted as an integer type")
	}
	if _, ok := typedBinaryOpKind(token.LAND); ok {
		t.Fatal("logical AND accepted as a scalar binary operation")
	}
	if _, ok := typedUnaryOpKind[float64](token.XOR, types.Typ[types.Float64]); ok {
		t.Fatal("float complement accepted")
	}
	if _, ok := typedValueBinaryOpKind[float64](token.REM, types.Typ[types.Float64], types.Typ[types.Float64]); ok {
		t.Fatal("float remainder accepted")
	}
}

// These are valid Go functions whose intermediate values cannot be represented
// by the selected typed frame, even when their results have a supported type.
func TestTypedIntermediateFallbacks(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("intermediates.go", `package main
func stringPhi(n int) int {s:="a";if n>0 {s="bb"};return len(s)}
func pointerConstant(n int,p *int) int {if p==nil {return n};return n+1}
func unusedCondition(n int,b bool) int {if b {return n};return -n}
func extraLeft(n,x int) int {return x+n}
func extraRight(n,x int) int {return n+x}
func extraUnary(n,x int) int {return -x}
func extraReturn(n,x int) int {return x}
func shift(n int,x uint) int {return n<<x}
func constantShift(n int,x int) int {return (n+x)<<uint(1)}
func badArray(n int) int {a:=[2]bool{true};if a[n] {return 1};return 0}
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"stringPhi", "pointerConstant", "unusedCondition", "extraLeft", "extraRight", "extraUnary", "extraReturn", "shift"} {
		if makeTypedScalarRunnerOf[int](pkg.Func(name)) != nil {
			t.Fatalf("%s unexpectedly compiled to a scalar runner", name)
		}
	}
	for _, name := range []string{"stringPhi", "shift", "badArray"} {
		if makeTypedArgumentsRunnerOf[int](pkg.Func(name)) != nil {
			t.Fatalf("%s unexpectedly compiled to an argument runner", name)
		}
	}
	r := makeTypedArgumentsRunnerOf[int](pkg.Func("constantShift"))
	if r == nil || r.(interface{ RunArgs([]Value) Value }).RunArgs([]Value{3, 4}) != 14 {
		t.Fatal("representable constant shift count was rejected")
	}
}

func TestTypedFrameCapacityFallbacks(t *testing.T) {
	source := "package main\n"
	source += "func params(" + strings.Repeat("_ int,", 65) + ") int {return 0}\n"
	source += "func constants(n int) int {" + strings.Repeat("n=n+1;", 35) + "return n}\n"
	source += "func phiConstants(n int) int {for i:=0;i<n;i++ {" + strings.Repeat("n=n+1;", 35) + "};return n}\n"
	source += "func scratch(n int) int {a,b,c,d,e,f,g:=n,n,n,n,n,n,n;for n>0 {n--;a++;b++;c++;d++;e++;f++;g++};return a+b+c+d+e+f+g}\n"
	source += "func many(n int)(int,int,int,int,int,int,int,int){return n,n,n,n,n,n,n,n}\n"
	source += "func tuples(n int) int {" + strings.Repeat("n,_,_,_,_,_,_,_=many(n);", 9) + "return n}\n"
	source += "func recursive(n int)(int,int,int,int,int,int,int,int){if n>0 {" + strings.Repeat("n=-n;", 25) + "return recursive(n-1)};return n,n,n,n,n,n,n,n}\n"
	pkg, err := NewContext(0).LoadFile("capacity.go", source)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"params", "constants", "phiConstants", "tuples"} {
		if makeTypedArgumentsRunnerOf[int](pkg.Func(name)) != nil {
			t.Fatalf("%s exceeded frame capacity without falling back", name)
		}
	}
	for _, name := range []string{"constants", "phiConstants", "recursive", "scratch"} {
		if makeTypedScalarRunnerOf[int](pkg.Func(name)) != nil {
			t.Fatalf("%s exceeded scalar capacity without falling back", name)
		}
	}
}

// A third-party factory can implement TypedRunner directly. Report malformed
// tuple results at the call boundary instead of silently writing bad registers.
type invalidTypedResultRunner struct{}

func (invalidTypedResultRunner) Run(Value) Value { return nil }

func TestTypedRunnerInvalidCallResult(t *testing.T) {
	fn := &function{narg: 1, nres: 2, typedRunner: invalidTypedResultRunner{}}
	if makeCallSpecializerInstr(nil, fn, []register{0}, []register{1}) != nil {
		t.Fatal("mismatched result destinations were accepted")
	}
	if makeCallSpecializerInstr(nil, fn, []register{0}, []register{1, -1}) != nil {
		t.Fatal("negative result destination was accepted")
	}
	call := makeCallSpecializerInstr(nil, fn, []register{0}, []register{1, 2})
	if call == nil {
		t.Fatal("missing typed call instruction")
	}
	defer func() {
		if got := recover(); got != "ixgo: typed runner returned an invalid result" {
			t.Fatalf("invalid result panic = %v", got)
		}
	}()
	call(&frame{stack: []Value{1, nil, nil}})
}

func TestTypedArgumentsFusionUglyFib(t *testing.T) {
	source, err := os.ReadFile("testdata/fibish/uglyfib.go")
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := NewContext(0).LoadFile("uglyfib.go", source)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"f", "g", "h", "k"} {
		before := makeTypedArgumentsRunnerOf[int](pkg.Func(name)).(*typedArgsRunner[int])
		after := makeTypedArgumentsRunnerOf[int](pkg.Func(name)).(*typedArgsRunner[int])
		if n := fuseTypedArguments(after); n != 8 {
			t.Fatalf("%s: got %d fusions, want 8", name, n)
		}
		if n := batchTypedConstants(after); n != 4 {
			t.Fatalf("%s: got %d constant batches, want 4", name, n)
		}
		for n := 0; n <= 20; n++ {
			for _, alias := range []bool{false, true} {
				run := func(r *typedArgsRunner[int]) int {
					a, b, result := n-1, n-2, 0
					pa, pb := &a, &b
					if alias {
						pa, pb = &result, &result
					}
					r.RunArgs([]Value{n, pa, pb, &result})
					return result
				}
				if want, got := run(before), run(after); want != got {
					t.Fatalf("%s(%d) alias=%v: fused=%d original=%d", name, n, alias, got, want)
				}
			}
		}
	}
}

func TestTypedArgumentsConstantBatches(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("args.go", typedArgsSource)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"calls", "memory", "pointerPhi", "unary", "index", "alias", "deep"} {
		before := makeTypedArgumentsRunnerOf[int](pkg.Func(name)).(*typedArgsRunner[int])
		after := makeTypedArgumentsRunnerOf[int](pkg.Func(name)).(*typedArgsRunner[int])
		fuseTypedArguments(after)
		batchTypedConstants(after)
		for _, n := range []int{0, 1, 2, 7, 100} {
			run := func(r *typedArgsRunner[int]) (Value, int, int) {
				a, b := 7, 11
				var args []Value
				switch name {
				case "calls":
					args = []Value{n, 5}
				case "memory", "pointerPhi":
					args = []Value{n, &a, &b}
				case "alias":
					args = []Value{&a, &a}
				case "index":
					args = []Value{n % 3, &a}
				default:
					args = []Value{n, &a}
				}
				return r.RunArgs(args), a, b
			}
			want, wa, wb := run(before)
			got, ga, gb := run(after)
			if !reflect.DeepEqual(got, want) || ga != wa || gb != wb {
				t.Fatalf("%s(%d): optimized=%v (%d,%d), original=%v (%d,%d)", name, n, got, ga, gb, want, wa, wb)
			}
		}
	}
	pkg, err = NewContext(0).LoadFile("constants.go", `package main
func floats(n float64)float64 {return n*2.5+1.25}
func narrow(n uint8)uint8 {return (n*3+7)^13}
func flags(n int)(int,bool) {b:=false;s:=2;for i:=0;i<n;i++ {b=!b;s+=3};return s+5,b}
`)
	if err != nil {
		t.Fatal(err)
	}
	floatRunner := makeTypedArgumentsRunnerOf[float64](pkg.Func("floats")).(*typedArgsRunner[float64])
	if batchTypedConstants(floatRunner) != 1 {
		t.Fatal("float constants not batched")
	}
	for _, n := range []float64{-2.25, 0, 1.5} {
		if got := floatRunner.RunArgs([]Value{n}); got != n*2.5+1.25 {
			t.Fatal("incorrect float result", got)
		}
	}
	narrow := makeTypedArgumentsRunnerOf[uint8](pkg.Func("narrow")).(*typedArgsRunner[uint8])
	if batchTypedConstants(narrow) != 1 {
		t.Fatal("narrow constants not batched")
	}
	for _, n := range []uint8{0, 1, 250, 255} {
		if got := narrow.RunArgs([]Value{n}); got != (n*3+7)^13 {
			t.Fatal("incorrect overflow result", got)
		}
	}
	flags := makeTypedArgumentsRunnerOf[int](pkg.Func("flags")).(*typedArgsRunner[int])
	batchTypedConstants(flags)
	for _, n := range []int{0, 1, 2, 7, 100} {
		want := Tuple{7 + 3*n, n%2 != 0}
		if got := flags.RunArgs([]Value{n}); !reflect.DeepEqual(got, want) {
			t.Fatalf("numeric/bool constants: got=%v want=%v", got, want)
		}
	}
}

func BenchmarkTypedArgumentsFusionUglyFib(b *testing.B) {
	source, err := os.ReadFile("testdata/fibish/uglyfib.go")
	if err != nil {
		b.Fatal(err)
	}
	pkg, err := NewContext(0).LoadFile("uglyfib.go", source)
	if err != nil {
		b.Fatal(err)
	}
	input, expected := 20, 6765
	if os.Getenv("IXGO_FIB40") == "1" {
		input, expected = 40, 102334155
	}
	for _, mode := range []string{"original", "fused", "constants", "default"} {
		b.Run(mode, func(b *testing.B) {
			r := makeTypedArgumentsRunnerOf[int](pkg.Func("f")).(*typedArgsRunner[int])
			if mode == "default" {
				r = MakeTypedRunnerOf[int](pkg.Func("f")).(*typedArgsRunner[int])
			}
			if (mode == "fused" || mode == "fused_constants") && fuseTypedArguments(r) != 8 {
				b.Fatal("expected 8 fusions across f/g/h/k")
			}
			if (mode == "constants" || mode == "fused_constants") && batchTypedConstants(r) != 4 {
				b.Fatal("expected 4 batched constant prologues")
			}
			b.ReportAllocs()
			b.ResetTimer()
			for j := 0; j < b.N; j++ {
				a, c, result := input-1, input-2, 0
				r.RunArgs([]Value{input, &a, &c, &result})
				if result != expected {
					b.Fatal(result)
				}
			}
		})
	}
}

func TestTypedArgumentsFusionValues(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("fusion.go", `package main
func update(n int,p *int)(int,int) {a:=*p;b:=a+n;*p=b;return a,b}
func reversed(n int,p *int)(int,int) {a:=*p;b:=n+a;*p=b;return a,b}
func twice(p *int)(int,int) {a:=*p;b:=a+a;*p=b;return a,b}
func compare(n int)(int,bool) {b:=n<5;if b{return n+1,b};return n-1,b}
func floats(n float64,p *float64)(float64,float64) {a:=*p;b:=a+n;*p=b;return a,b}
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"update", "reversed", "twice", "compare"} {
		r := makeTypedArgumentsRunnerOf[int](pkg.Func(name)).(*typedArgsRunner[int])
		if fuseTypedArguments(r) != 1 {
			t.Fatalf("%s: expected one fusion", name)
		}
		for _, n := range []int{-3, 0, 9} {
			p := 7
			args, want := []Value{n, &p}, Tuple{7, 7 + n}
			if name == "twice" {
				args, want = []Value{&p}, Tuple{7, 14}
			} else if name == "compare" {
				args, want = []Value{n}, Tuple{n - 1, false}
				if n < 5 {
					want = Tuple{n + 1, true}
				}
			}
			if got := r.RunArgs(args); !reflect.DeepEqual(got, want) {
				t.Fatalf("%s: got=%v want=%v", name, got, want)
			}
			if name != "compare" && p != want[1] {
				t.Fatalf("%s: pointer write got=%v want=%v", name, p, want[1])
			}
		}
		if name != "compare" {
			func() {
				defer func() {
					if recover() == nil {
						t.Fatalf("%s: nil pointer did not panic", name)
					}
				}()
				args := []Value{1, (*int)(nil)}
				if name == "twice" {
					args = args[1:]
				}
				r.RunArgs(args)
			}()
		}
	}
	r := makeTypedArgumentsRunnerOf[float64](pkg.Func("floats")).(*typedArgsRunner[float64])
	if fuseTypedArguments(r) != 1 {
		t.Fatal("float fusion missing")
	}
	p := 1.5
	if got := r.RunArgs([]Value{2.25, &p}); !reflect.DeepEqual(got, Tuple{1.5, 3.75}) || p != 3.75 {
		t.Fatalf("float result=%v pointer=%v", got, p)
	}
}

func TestTypedRunnerDefaultOptimizations(t *testing.T) {
	source, err := os.ReadFile("testdata/fibish/uglyfib.go")
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := NewContext(0).LoadFile("uglyfib.go", source)
	if err != nil {
		t.Fatal(err)
	}
	runner := MakeTypedRunnerOf[int](pkg.Func("f")).(*typedArgsRunner[int])
	seen := make(map[*typedArgsRunner[int]]bool)
	var visit func(*typedArgsRunner[int])
	visit = func(r *typedArgsRunner[int]) {
		if seen[r] {
			return
		}
		seen[r] = true
		kinds := make(map[typedOpKind]bool)
		for _, op := range r.code {
			kinds[op.kind] = true
			if op.kind == typedCall {
				visit(op.callee)
			}
		}
		for _, kind := range []typedOpKind{typedConstBlock, typedLessBranch, typedLoadAddStore} {
			if !kinds[kind] {
				t.Errorf("reachable runner missing optimized instruction %v", kind)
			}
		}
	}
	visit(runner)
	if len(seen) != 4 {
		t.Fatalf("optimized %d functions, want 4", len(seen))
	}
	a, b, result := 19, 18, 0
	runner.RunArgs([]Value{20, &a, &b, &result})
	if result != 6765 {
		t.Fatal("incorrect optimized uglyfib result", result)
	}
}

func TestTypedOptimizationsOrderAndConcurrency(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("args.go", typedArgsSource)
	if err != nil {
		t.Fatal(err)
	}
	for _, constantsFirst := range []bool{false, true} {
		r := makeTypedArgumentsRunnerOf[int](pkg.Func("deep")).(*typedArgsRunner[int])
		if constantsFirst {
			batchTypedConstants(r)
			fuseTypedArguments(r)
		} else {
			fuseTypedArguments(r)
			batchTypedConstants(r)
		}
		if fuseTypedArguments(r) != 0 || batchTypedConstants(r) != 0 {
			t.Fatal("reapplying optimization changed already optimized code")
		}
		var wg sync.WaitGroup
		for worker := 0; worker < 8; worker++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				for _, n := range []int{0, 1, 20, 100} {
					p := 7
					r.RunArgs([]Value{n, &p})
					if want := 7 + n*(n+3)/2; p != want {
						t.Errorf("deep(%d): got=%d want=%d", n, p, want)
					}
				}
			}()
		}
		wg.Wait()
	}
}

func TestTypedFusionFloatBoundaries(t *testing.T) {
	pkg, err := NewContext(0).LoadFile("float_fusion.go", `package main
func compare(n float64)float64 {if n<0 {return -1};return 1}
func add(n float64,p *float64)(float64,float64) {a:=*p;b:=a+n;*p=b;return a,b}
func reverse(n float64,p *float64)(float64,float64) {a:=*p;b:=n+a;*p=b;return a,b}
`)
	if err != nil {
		t.Fatal(err)
	}
	values := []float64{0, math.Copysign(0, -1), 1, -1, math.Inf(1), math.Inf(-1), math.Float64frombits(0x7ff8000000000001)}
	for _, name := range []string{"compare", "add", "reverse"} {
		before := makeTypedArgumentsRunnerOf[float64](pkg.Func(name))
		after := makeTypedArgumentsRunnerOf[float64](pkg.Func(name))
		if fuseTypedArguments(after.(*typedArgsRunner[float64])) != 1 {
			t.Fatal("float instruction sequence not fused")
		}
		for _, x := range values {
			for _, y := range values {
				run := func(r TypedRunner) []uint64 {
					if name == "compare" {
						return []uint64{math.Float64bits(r.Run(x).(float64))}
					}
					p := y
					result := r.(TypedArgumentsRunner).RunArgs([]Value{x, &p}).(Tuple)
					return []uint64{math.Float64bits(result[0].(float64)), math.Float64bits(result[1].(float64)), math.Float64bits(p)}
				}
				if want, got := run(before), run(after); !reflect.DeepEqual(got, want) {
					t.Fatalf("%s(%v,%v): fused bits=%x original bits=%x", name, x, y, got, want)
				}
			}
		}
	}
}
