package ixgo

import (
	"go/constant"
	"os"
	"reflect"
	"sync"
	"testing"

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
func entry(tick func()) (int,bool,int,int) {
 tick();p:=0;void(4,&p);a,b:=pair(3);sum,ok:=many(1,2,3,4,5,6,7,8,9,10)
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
	result := runner.runTyped(input)
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
	got := pairResult(runner, 40)
	if got[0] != 39088169 || got[1] != 126491972 {
		t.Fatalf("f(40) = (%d, %d)", got[0], got[1])
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
			runner := makeTypedArgumentsRunnerOf[int](pkg.Func(name)).(TypedArgumentsRunner)
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
	pkg, err := ctx.LoadFile("named.go", `package main
type Count int
type Bits uint8
func count(n Count) Count { if n < 1 { return n }; return count(n-1)+1 }
func bits(n Bits) Bits { return (n ^ 3) + 1 }
`)
	if err != nil {
		t.Fatal(err)
	}
	type MyInt int
	type MyBits uint8
	if runner := MakeTypedRunnerOf[MyInt](pkg.Func("count")); runner == nil || runner.Run(MyInt(9)) != MyInt(9) {
		t.Fatal("named ~int type was not compiled")
	}
	if runner := MakeTypedRunnerOf[MyBits](pkg.Func("bits")); runner == nil || runner.Run(MyBits(5)) != MyBits(7) {
		t.Fatal("named ~uint8 type was not compiled")
	}
}
