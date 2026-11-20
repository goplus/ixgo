package mixedrunner

import (
	"fmt"
	"go/constant"
	"go/token"
	"go/types"
	"math"
	"reflect"
	"strings"
	"testing"

	"github.com/goplus/ixgo"
	"golang.org/x/tools/go/ssa"
)

func TestMixedNumericOperations(t *testing.T) {
	t.Run("int", func(t *testing.T) { testMixedNumericOperations(t, int(-17), int(3)) })
	t.Run("int8", func(t *testing.T) { testMixedNumericOperations(t, int8(-128), int8(-1)) })
	t.Run("int16", func(t *testing.T) { testMixedNumericOperations(t, int16(-32768), int16(-1)) })
	t.Run("int32", func(t *testing.T) { testMixedNumericOperations(t, int32(-17), int32(3)) })
	t.Run("int64", func(t *testing.T) { testMixedNumericOperations(t, int64(math.MinInt64), int64(-1)) })
	t.Run("uint", func(t *testing.T) { testMixedNumericOperations(t, ^uint(0), uint(3)) })
	t.Run("uint8", func(t *testing.T) { testMixedNumericOperations(t, uint8(255), uint8(3)) })
	t.Run("uint16", func(t *testing.T) { testMixedNumericOperations(t, uint16(65535), uint16(3)) })
	t.Run("uint32", func(t *testing.T) { testMixedNumericOperations(t, ^uint32(0), uint32(3)) })
	t.Run("uint64", func(t *testing.T) { testMixedNumericOperations(t, ^uint64(0), uint64(3)) })
	t.Run("uintptr", func(t *testing.T) { testMixedNumericOperations(t, ^uintptr(0), uintptr(3)) })
	t.Run("float32", func(t *testing.T) { testMixedNumericOperations(t, float32(-7.5), float32(2)) })
	t.Run("float64", func(t *testing.T) { testMixedNumericOperations(t, -7.5, 2.0) })
}

func testMixedNumericOperations[T mixedScalar](t *testing.T, x, y T) {
	t.Helper()
	typ := reflect.TypeOf(x).String()
	ops := []struct {
		name, expr string
		want       func(T, T) ixgo.Value
	}{
		{"add", "x+y", func(x, y T) ixgo.Value { return x + y }},
		{"sub", "x-y", func(x, y T) ixgo.Value { return x - y }},
		{"mul", "x*y", func(x, y T) ixgo.Value { return x * y }},
		{"quo", "x/y", func(x, y T) ixgo.Value { return x / y }},
		{"neg", "-x", func(x, y T) ixgo.Value { return -x }},
		{"lt", "x<y", func(x, y T) ixgo.Value { return x < y }},
		{"le", "x<=y", func(x, y T) ixgo.Value { return x <= y }},
		{"gt", "x>y", func(x, y T) ixgo.Value { return x > y }},
		{"ge", "x>=y", func(x, y T) ixgo.Value { return x >= y }},
		{"eq", "x==y", func(x, y T) ixgo.Value { return x == y }},
		{"ne", "x!=y", func(x, y T) ixgo.Value { return x != y }},
	}
	var source strings.Builder
	source.WriteString("package main\n")
	for _, op := range ops {
		result := reflect.TypeOf(op.want(x, y)).String()
		fmt.Fprintf(&source, "func %s(x,y %s) %s {return %s}\n", op.name, typ, result, op.expr)
	}
	targets := []ixgo.Value{int(0), int8(0), int16(0), int32(0), int64(0), uint(0), uint8(0), uint16(0), uint32(0), uint64(0), uintptr(0), float32(0), float64(0)}
	for i, target := range targets {
		to := reflect.TypeOf(target).String()
		fmt.Fprintf(&source, "func conv%d(x %s) %s {return %s(x)}\n", i, typ, to, to)
	}
	pkg, err := ixgo.NewContext(0).LoadFile("operations.go", source.String())
	if err != nil {
		t.Fatal(err)
	}
	for _, op := range ops {
		t.Run(op.name, func(t *testing.T) {
			r, ok := MakeRunner(pkg.Func(op.name)).(ixgo.TypedArgumentsRunner)
			if !ok {
				t.Fatal("numeric operation rejected")
			}
			for _, pair := range [][2]T{{x, y}, {y, x}, {y, y}} {
				got := r.RunArgs([]ixgo.Value{pair[0], pair[1]})
				if want := op.want(pair[0], pair[1]); !reflect.DeepEqual(got, want) {
					t.Fatalf("(%v, %v): got %#v, want %#v", pair[0], pair[1], got, want)
				}
			}
		})
	}
	for i, target := range targets {
		// Keep float-to-integer conversions in range for every destination.
		value := T(7)
		want := reflect.ValueOf(value).Convert(reflect.TypeOf(target)).Interface()
		r, ok := MakeRunner(pkg.Func(fmt.Sprintf("conv%d", i))).(ixgo.TypedArgumentsRunner)
		if !ok {
			t.Fatalf("conversion to %T rejected", target)
		}
		if got := r.RunArgs([]ixgo.Value{value}); !reflect.DeepEqual(got, want) {
			t.Fatalf("conversion to %T: got %#v, want %#v", target, got, want)
		}
	}
}

type mixedTestInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

func TestMixedIntegerOperations(t *testing.T) {
	testMixedIntegerOperations(t, int(-17), int(3))
	testMixedIntegerOperations(t, int8(-128), int8(-1))
	testMixedIntegerOperations(t, int16(-32768), int16(-1))
	testMixedIntegerOperations(t, int32(-17), int32(3))
	testMixedIntegerOperations(t, int64(math.MinInt64), int64(-1))
	testMixedIntegerOperations(t, ^uint(0), uint(3))
	testMixedIntegerOperations(t, uint8(255), uint8(3))
	testMixedIntegerOperations(t, uint16(65535), uint16(3))
	testMixedIntegerOperations(t, ^uint32(0), uint32(3))
	testMixedIntegerOperations(t, ^uint64(0), uint64(3))
	testMixedIntegerOperations(t, ^uintptr(0), uintptr(3))
}

func testMixedIntegerOperations[T mixedTestInteger](t *testing.T, x, y T) {
	t.Helper()
	t.Run(reflect.TypeOf(x).String(), func(t *testing.T) {
		ops := []struct {
			name, expr string
			want       func(T, T, uint64) T
		}{
			{"rem", "x%y", func(x, y T, _ uint64) T { return x % y }},
			{"and", "x&y", func(x, y T, _ uint64) T { return x & y }},
			{"or", "x|y", func(x, y T, _ uint64) T { return x | y }},
			{"xor", "x^y", func(x, y T, _ uint64) T { return x ^ y }},
			{"andnot", "x&^y", func(x, y T, _ uint64) T { return x &^ y }},
			{"complement", "^x", func(x, y T, _ uint64) T { return ^x }},
			{"left", "x<<n", func(x, y T, n uint64) T { return x << n }},
			{"right", "x>>n", func(x, y T, n uint64) T { return x >> n }},
		}
		var source strings.Builder
		source.WriteString("package main\n")
		for _, op := range ops {
			fmt.Fprintf(&source, "func %s(x,y %T,n uint64) %T {return %s}\n", op.name, x, x, op.expr)
		}
		pkg, err := ixgo.NewContext(0).LoadFile("integers.go", source.String())
		if err != nil {
			t.Fatal(err)
		}
		for _, op := range ops {
			r, ok := MakeRunner(pkg.Func(op.name)).(ixgo.TypedArgumentsRunner)
			if !ok {
				t.Fatalf("%s rejected", op.name)
			}
			for _, shift := range []uint64{0, 1, 7, 31, 63, 64, 255} {
				if got, want := r.RunArgs([]ixgo.Value{x, y, shift}), op.want(x, y, shift); got != want {
					t.Fatalf("%s, shift %d: got %#v, want %#v", op.name, shift, got, want)
				}
			}
		}
	})
}

func mixedTestPanic(t *testing.T, f func()) (v interface{}) {
	t.Helper()
	defer func() {
		v = recover()
		if v == nil {
			t.Error("expected panic")
		}
	}()
	f()
	return nil
}

func TestMixedSignedShiftCounts(t *testing.T) {
	for _, count := range []ixgo.Value{int(-1), int8(-1), int16(-1), int32(-1), int64(-1)} {
		t.Run(reflect.TypeOf(count).String(), func(t *testing.T) {
			source := fmt.Sprintf("package main; func left(x int16,n %T) int16 {return x<<n}; func right(x int16,n %T) int16 {return x>>n}", count, count)
			pkg, err := ixgo.NewContext(0).LoadFile("shifts.go", source)
			if err != nil {
				t.Fatal(err)
			}
			n := int64(-1)
			want := mixedTestPanic(t, func() { _ = int16(1) << n })
			for _, name := range []string{"left", "right"} {
				r := MakeRunner(pkg.Func(name)).(ixgo.TypedArgumentsRunner)
				got := mixedTestPanic(t, func() { r.RunArgs([]ixgo.Value{int16(-16), count}) })
				if !reflect.DeepEqual(got, want) {
					t.Fatalf("%s panic: %#v, want %#v", name, got, want)
				}
				positive := reflect.ValueOf(int(2)).Convert(reflect.TypeOf(count)).Interface()
				expected := int16(-4)
				if name == "left" {
					expected = -64
				}
				if got := r.RunArgs([]ixgo.Value{int16(-16), positive}); got != expected {
					t.Fatalf("%s: got %v, want %v", name, got, expected)
				}
			}
		})
	}
}

func TestMixedUnsupportedGraphs(t *testing.T) {
	var source strings.Builder
	source.WriteString(`package main
func external(x int) int
func externalCall(x int) int {return external(x)}
func variadic(x ...int) int {return 0}
func builtin(x int) int {return len("hello")+int(real(complex(float64(x),1)))}
func panicCall(x int) int {panic(x)}
func pointer(p *int) int {return *p}
func boolField(p *struct{b bool}) int {return 0}
func nestedField(p *struct{p *int}) int {return 0}
func nonnumeric(p *struct{s string}) int {return 0}
func capture(x int) func() int {return func() int {return x}}
func manyResults() (int,int,int,int,int,int,int,int,int) {return 0,0,0,0,0,0,0,0,0}
func manyCall(x int) int {a,b,c,d,e,f,g,h,i:=manyResults(); return x+a+b+c+d+e+f+g+h+i}
func discardCall(x int) int {external(x);return x}
`)
	source.WriteString("func frame(x int) int {\n")
	for i := 0; i < 70; i++ {
		fmt.Fprintf(&source, "x=x*%d+%d\n", i+2, i+1)
	}
	source.WriteString("return x}\n")
	pkg, err := ixgo.NewContext(0).LoadFile("unsupported.go", source.String())
	if err != nil {
		t.Fatal(err)
	}
	if MakeRunner(nil) != nil || MakeRunner(&ssa.Function{}) != nil {
		t.Fatal("empty function accepted")
	}
	for _, name := range []string{"external", "externalCall", "variadic", "builtin", "panicCall", "pointer", "boolField", "nestedField", "nonnumeric", "capture", "manyResults", "manyCall", "discardCall", "frame"} {
		t.Run(name, func(t *testing.T) {
			fn := pkg.Func(name)
			if fn == nil {
				t.Fatal("missing test function")
			}
			if MakeRunner(fn) != nil {
				t.Fatal("unsupported graph accepted")
			}
		})
	}
	if MakeRunner(pkg.Func("capture").AnonFuncs[0]) != nil {
		t.Fatal("free variables accepted")
	}
}

func TestMixedConstantValidation(t *testing.T) {
	for _, c := range []*ssa.Const{
		ssa.NewConst(nil, types.NewPointer(types.Typ[types.Int])),
		ssa.NewConst(constant.MakeString("x"), types.Typ[types.String]),
		ssa.NewConst(constant.MakeImag(constant.MakeInt64(1)), types.Typ[types.Complex128]),
	} {
		if _, ok := mixedConstant(c); ok {
			t.Fatalf("unsupported constant accepted: %v", c)
		}
	}
	c := ssa.NewConst(constant.MakeUint64(123), types.Typ[types.Uintptr])
	if got, ok := mixedConstant(c); !ok || got != 123 {
		t.Fatalf("uintptr constant: %v, %v", got, ok)
	}
	if _, ok := mixedConstantValue[float64](constant.MakeBool(true)); ok {
		t.Fatal("boolean float constant accepted")
	}
	if _, ok := mixedConstantValue[int](constant.MakeFloat64(1.5)); ok {
		t.Fatal("fractional integer constant accepted")
	}
	if _, ok := mixedConstantValue[int64](constant.MakeUint64(math.MaxUint64)); ok {
		t.Fatal("overflowing signed constant accepted")
	}
	if _, ok := mixedConstantValue[uint64](constant.MakeInt64(-1)); ok {
		t.Fatal("negative unsigned constant accepted")
	}
	if _, ok := mixedConstantValue[float64](constant.Shift(constant.MakeInt64(1), token.SHL, 2048)); ok {
		t.Fatal("infinite float constant accepted")
	}
}

func TestMixedInvalidOpcodes(t *testing.T) {
	for _, tc := range []struct {
		name string
		call func()
	}{
		{"argument kind", func() { mixedUnbox(1, types.String) }},
		{"result kind", func() { mixedBox(1, types.String) }},
		{"conversion kind", func() { mixedConversion(1, types.Bool) }},
		{"scalar kind", func() { type named int; mixedBits(named(1)) }},
		{"evaluate kind", func() { mixedEvaluate(&mixedOp{scalar: types.String}, 0, 0) }},
		{"numeric opcode", func() { mixedNumeric(&mixedOp{}, 1, 2, 0) }},
	} {
		t.Run(tc.name, func(t *testing.T) { mixedTestPanic(t, tc.call) })
	}
	for _, tc := range []struct {
		kind       mixedOpKind
		x, y, want uint64
	}{
		{mixedBaseNot, 0, 0, 1}, {mixedBaseNot, 1, 0, 0},
		{mixedBaseEqual, 1, 1, 1}, {mixedBaseEqual, 1, 0, 0},
		{mixedBaseNotEqual, 1, 0, 1}, {mixedBaseNotEqual, 1, 1, 0},
	} {
		op := mixedOp{mixedBaseOp: mixedBaseOp{kind: tc.kind}, scalar: types.Bool}
		if got := mixedEvaluate(&op, tc.x, tc.y); got != tc.want {
			t.Fatalf("boolean opcode %d: got %d, want %d", tc.kind, got, tc.want)
		}
	}
}

func TestMixedReceiverValidation(t *testing.T) {
	pkg, err := ixgo.NewContext(0).LoadFile("nilreceiver.go", `package main
	type S struct {x int}
	func read(p *S) int {return p.x}
	`)
	if err != nil {
		t.Fatal(err)
	}
	r := MakeRunner(pkg.Func("read")).(*mixedTypedRunner)
	mixedTestPanic(t, func() { r.RunArgs([]ixgo.Value{&struct{ X int }{1}}) })
	r.BindTypes(func(types.Type) reflect.Type { return reflect.TypeOf((*struct{ X int })(nil)) })
	mixedTestPanic(t, func() { r.RunArgs([]ixgo.Value{&struct{ Y int }{1}}) })
	got := mixedTestPanic(t, func() { r.RunArgs([]ixgo.Value{(*struct{ X int })(nil)}) })
	want := mixedTestPanic(t, func() { var p *struct{ X int }; _ = p.X })
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("nil receiver panic: %#v, want %#v", got, want)
	}
}

func TestMixedMalformedSSA(t *testing.T) {
	const source = `package main
func binary(x,y int) int {return x+y}
func floating(x,y float64) float64 {return x+y}
func boolean(x,y bool) bool {return x==y}
func unary(x int) int {return -x}
func convert(x int) float64 {return float64(x)}
func branch(x int) int {if x>0{return 1}; return 2}
func loop(x int) int {for x>0{x--};return x}
func pair(x int) (int,int) {return x,x+1}
func extract(x int) int {a,b:=pair(x);return a+b}
func builtin(x int) int {return min(x,2)}
`
	one := ssa.NewConst(constant.MakeInt64(1), types.Typ[types.Int])
	float := ssa.NewConst(constant.MakeFloat64(1), types.Typ[types.Float64])
	boolean := ssa.NewConst(constant.MakeBool(true), types.Typ[types.Bool])
	str := ssa.NewConst(constant.MakeString("x"), types.Typ[types.String])
	for _, tc := range []struct {
		name, fn string
		mutate   func(*ssa.Function)
	}{
		{"unknown binary", "binary", func(f *ssa.Function) { f.Blocks[0].Instrs[0].(*ssa.BinOp).Op = token.LAND }},
		{"mismatched binary", "binary", func(f *ssa.Function) { f.Blocks[0].Instrs[0].(*ssa.BinOp).Y = float }},
		{"float remainder", "floating", func(f *ssa.Function) { f.Blocks[0].Instrs[0].(*ssa.BinOp).Op = token.REM }},
		{"float shift", "floating", func(f *ssa.Function) { f.Blocks[0].Instrs[0].(*ssa.BinOp).Op = token.SHL }},
		{"boolean arithmetic", "boolean", func(f *ssa.Function) { f.Blocks[0].Instrs[0].(*ssa.BinOp).Op = token.ADD }},
		{"unknown unary", "unary", func(f *ssa.Function) { f.Blocks[0].Instrs[0].(*ssa.UnOp).Op = token.ADD }},
		{"integer not", "unary", func(f *ssa.Function) { f.Blocks[0].Instrs[0].(*ssa.UnOp).Op = token.NOT }},
		{"boolean negate", "unary", func(f *ssa.Function) { f.Blocks[0].Instrs[0].(*ssa.UnOp).X = boolean }},
		{"float complement", "unary", func(f *ssa.Function) { u := f.Blocks[0].Instrs[0].(*ssa.UnOp); u.Op = token.XOR; u.X = float }},
		{"boolean conversion", "convert", func(f *ssa.Function) { f.Blocks[0].Instrs[0].(*ssa.Convert).X = boolean }},
		{"missing branch successor", "branch", func(f *ssa.Function) { f.Blocks[0].Succs = f.Blocks[0].Succs[:1] }},
		{"missing jump successor", "loop", func(f *ssa.Function) { f.Blocks[0].Succs = nil }},
		{"wrong result count", "binary", func(f *ssa.Function) { f.Blocks[0].Instrs[1].(*ssa.Return).Results = nil }},
		{"unsupported result operand", "binary", func(f *ssa.Function) { f.Blocks[0].Instrs[1].(*ssa.Return).Results = []ssa.Value{str} }},
		{"invalid tuple index", "extract", func(f *ssa.Function) {
			for _, b := range f.Blocks {
				for _, ins := range b.Instrs {
					if e, ok := ins.(*ssa.Extract); ok {
						e.Index = 9
						return
					}
				}
			}
		}},
		{"missing phi edge", "loop", func(f *ssa.Function) {
			for _, b := range f.Blocks {
				for _, ins := range b.Instrs {
					if p, ok := ins.(*ssa.Phi); ok {
						p.Edges = p.Edges[:1]
						return
					}
				}
			}
		}},
		{"invalid integer constant", "binary", func(f *ssa.Function) {
			f.Blocks[0].Instrs[0].(*ssa.BinOp).Y = ssa.NewConst(constant.MakeUint64(math.MaxUint64), types.Typ[types.Int64])
		}},
		{"builtin call", "builtin", func(f *ssa.Function) {}},
		{"unbound binary operand", "binary", func(f *ssa.Function) {
			f.Blocks[0].Instrs[0].(*ssa.BinOp).X = &ssa.Function{Signature: types.NewSignatureType(nil, nil, nil, nil, nil, false)}
		}},
		{"constant operand", "binary", func(f *ssa.Function) {
			f.Blocks[0].Instrs[0].(*ssa.BinOp).Y = one
			f.Blocks[0].Instrs = append(f.Blocks[0].Instrs[:1], &ssa.Panic{X: one})
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pkg, err := ixgo.NewContext(0).LoadFile("malformed.go", source)
			if err != nil {
				t.Fatal(err)
			}
			fn := pkg.Func(tc.fn)
			tc.mutate(fn)
			if MakeRunner(fn) != nil {
				t.Fatal("malformed SSA accepted")
			}
		})
	}
}

func TestMixedParallelPhiCycles(t *testing.T) {
	pkg, err := ixgo.NewContext(ixgo.DisableMemo).LoadFile("cycles.go", `package main
func rotate(n int) int {
 a,b,c:=1,2,3
 for j:=0;j<n;j++ { a,b,c=b,c,a }
 return 100*a+10*b+c
}
func disjoint(n int) int {
 a,b,c,d:=1,2,3,4
 for j:=0;j<n;j++ { a,b,c,d=b,a,d,c }
 return 1000*a+100*b+10*c+d
}
`)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"rotate", "disjoint"} {
		r := MakeRunner(pkg.Func(name)).(*mixedTypedRunner)
		for n := 0; n < 12; n++ {
			a, b, c, d := 1, 2, 3, 4
			for j := 0; j < n; j++ {
				if name == "rotate" {
					a, b, c = b, c, a
				} else {
					a, b, c, d = b, a, d, c
				}
			}
			want := 100*a + 10*b + c
			if name == "disjoint" {
				want = 1000*a + 100*b + 10*c + d
			}
			for _, numeric := range []bool{false, true} {
				r.numeric = numeric
				if got := r.RunArgs([]ixgo.Value{n}); got != want {
					t.Fatalf("%s(%d), numeric=%v: %v, want %d", name, n, numeric, got, want)
				}
			}
		}
	}
}
