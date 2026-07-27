package sccp

import (
	"go/ast"
	"go/constant"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"testing"

	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func buildPackage(t *testing.T, src string) *ssa.Package {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "p.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	pkg := types.NewPackage("p", "")
	ssapkg, _, err := ssautil.BuildPackage(
		&types.Config{Importer: importer.Default()}, fset, pkg,
		[]*ast.File{file}, ssa.SanityCheckFunctions)
	if err != nil {
		t.Fatal(err)
	}
	return ssapkg
}

func runSCCP(t *testing.T, src, name string) *ssa.Function {
	t.Helper()
	pkg := buildPackage(t, src)
	if err := (Pass{}).Run(pkg); err != nil {
		t.Fatal(err)
	}
	fn := pkg.Func(name)
	if fn == nil {
		t.Fatalf("function %s not found", name)
	}
	assertSSASane(t, fn)
	return fn
}

// returnConst returns the constant returned by fn's (single reachable)
// return instruction, or nil.
func returnConst(fn *ssa.Function) *ssa.Const {
	reachable := make(map[*ssa.BasicBlock]bool)
	var mark func(b *ssa.BasicBlock)
	mark = func(b *ssa.BasicBlock) {
		if reachable[b] {
			return
		}
		reachable[b] = true
		for _, s := range b.Succs {
			mark(s)
		}
	}
	mark(fn.Blocks[0])
	for _, b := range fn.Blocks {
		if !reachable[b] {
			continue
		}
		for _, instr := range b.Instrs {
			if ret, ok := instr.(*ssa.Return); ok && len(ret.Results) == 1 {
				if c, ok := ret.Results[0].(*ssa.Const); ok {
					return c
				}
			}
		}
	}
	return nil
}

func TestBranchFolding(t *testing.T) {
	fn := runSCCP(t, `package p
func f() int {
	x := 1
	if x == 1 {
		x = 2
	} else {
		x = 3
	}
	return x
}`, "f")
	c := returnConst(fn)
	if c == nil || c.Int64() != 2 {
		fn.WriteTo(testWriter{t})
		t.Fatalf("want return constant 2, got %v", c)
	}
	if len(fn.Blocks) != 1 {
		fn.WriteTo(testWriter{t})
		t.Fatalf("blocks after SCCP = %d, want 1", len(fn.Blocks))
	}
}

// TestLoopInvariant exercises the "conditional" part of sccp: x==1 can only
// be proven by propagating constants along reachable paths through the loop.
func TestLoopInvariant(t *testing.T) {
	fn := runSCCP(t, `package p
func g(n int) int {
	x := 1
	for i := 0; i < n; i++ {
		if x != 1 {
			x = 2
		}
	}
	return x
}`, "g")
	c := returnConst(fn)
	if c == nil || c.Int64() != 1 {
		fn.WriteTo(testWriter{t})
		t.Fatalf("want return constant 1, got %v", c)
	}
}

func TestDivByZeroNotFolded(t *testing.T) {
	fn := runSCCP(t, `package p
func h() int {
	z := 0
	return 10 / z
}`, "h")
	if c := returnConst(fn); c != nil {
		t.Fatalf("division by zero must not be folded, got constant %v", c)
	}
	// the BinOp must survive so the run-time panic is preserved
	found := false
	for _, b := range fn.Blocks {
		for _, instr := range b.Instrs {
			if bin, ok := instr.(*ssa.BinOp); ok && bin.Op == token.QUO {
				found = true
			}
		}
	}
	if !found {
		t.Fatal("QUO instruction was removed")
	}
}

func TestLiteralCondRewire(t *testing.T) {
	fn := runSCCP(t, `package p
var sink int
func k(a int) int {
	if a > 0 || true {
		sink = 1
	}
	return sink
}`, "k")
	assertSSASane(t, fn)
}

func TestPhiChain(t *testing.T) {
	fn := runSCCP(t, `package p
func m(cond bool) int {
	a := 10
	b := 0
	if cond {
		b = a * 2
	} else {
		b = 20
	}
	return b + 1
}`, "m")
	c := returnConst(fn)
	if c == nil || c.Int64() != 21 {
		fn.WriteTo(testWriter{t})
		t.Fatalf("want return constant 21, got %v", c)
	}
}

func TestUnsignedWrap(t *testing.T) {
	fn := runSCCP(t, `package p
func u() uint8 {
	a := uint8(200)
	b := uint8(100)
	return a + b
}`, "u")
	c := returnConst(fn)
	if c == nil || c.Uint64() != 44 {
		fn.WriteTo(testWriter{t})
		t.Fatalf("want return constant 44, got %v", c)
	}
}

func TestMethodsAndClosures(t *testing.T) {
	pkg := buildPackage(t, `package p
type T int
func (t T) Double() T {
	x := T(2)
	return t * x
}
func closure() func() int {
	x := 3
	return func() int {
		y := 4
		if y > x {
			return 1
		}
		return 0
	}
}`)
	if err := (Pass{}).Run(pkg); err != nil {
		t.Fatal(err)
	}
	fns := functions(pkg)
	// expect init, Double, closure, closure$1
	var foundMethod, foundAnon bool
	for _, fn := range fns {
		assertSSASane(t, fn)
		if fn.Name() == "Double" {
			foundMethod = true
		}
		if fn.Parent() != nil {
			foundAnon = true
			if c := returnConstAny(fn); c == nil || c.Int64() != 1 {
				fn.WriteTo(testWriter{t})
				t.Fatalf("closure: want return constant 1, got %v", c)
			}
		}
	}
	if !foundMethod || !foundAnon {
		t.Fatalf("functions() missed method (%v) or anon func (%v)", foundMethod, foundAnon)
	}
}

// returnConstAny is like returnConst but looks at the first reachable return.
func returnConstAny(fn *ssa.Function) *ssa.Const {
	return returnConst(fn)
}

// fold unit tests ------------------------------------------------------------

func icst(v int64, k types.BasicKind) *ssa.Const {
	return ssa.NewConst(constant.MakeInt64(v), types.Typ[k])
}

func ucst(v uint64, k types.BasicKind) *ssa.Const {
	return ssa.NewConst(constant.MakeUint64(v), types.Typ[k])
}

func fcst(v float64, k types.BasicKind) *ssa.Const {
	return ssa.NewConst(constant.MakeFloat64(v), types.Typ[k])
}

func TestFoldBinOpInt(t *testing.T) {
	tests := []struct {
		op   token.Token
		x, y *ssa.Const
		want interface{} // int64, uint64, bool or nil (must not fold)
	}{
		{token.ADD, icst(127, types.Int8), icst(1, types.Int8), int64(-128)}, // wrap
		{token.MUL, icst(1000, types.Int16), icst(1000, types.Int16), int64(16960)},
		{token.QUO, icst(-9223372036854775808, types.Int64), icst(-1, types.Int64), int64(-9223372036854775808)},
		{token.QUO, icst(7, types.Int64), icst(0, types.Int64), nil}, // div by zero
		{token.REM, icst(7, types.Int64), icst(0, types.Int64), nil},
		{token.SHL, icst(1, types.Int8), icst(10, types.Int64), int64(0)},    // shift out
		{token.SHR, icst(-1, types.Int8), icst(100, types.Int64), int64(-1)}, // arith shift
		{token.SHL, icst(1, types.Int64), icst(-1, types.Int64), nil},        // negative shift
		{token.ADD, ucst(200, types.Uint8), ucst(100, types.Uint8), uint64(44)},
		{token.SUB, ucst(0, types.Uint16), ucst(1, types.Uint16), uint64(65535)},
		{token.LSS, icst(-1, types.Int32), icst(1, types.Int32), true},
		{token.LSS, ucst(0xffffffff, types.Uint32), ucst(1, types.Uint32), false},
		{token.AND_NOT, icst(0xff, types.Int32), icst(0x0f, types.Int32), int64(0xf0)},
	}
	for i, tt := range tests {
		var typ types.Type = tt.x.Type()
		if _, isBool := tt.want.(bool); isBool {
			typ = types.Typ[types.Bool]
		}
		got := foldBinOp(tt.op, tt.x, tt.y, typ)
		if tt.want == nil {
			if got != nil {
				t.Errorf("#%d: %v %v %v: folded to %v, want no fold", i, tt.x, tt.op, tt.y, got)
			}
			continue
		}
		if got == nil {
			t.Errorf("#%d: %v %v %v: not folded", i, tt.x, tt.op, tt.y)
			continue
		}
		switch want := tt.want.(type) {
		case int64:
			if got.Int64() != want {
				t.Errorf("#%d: got %v, want %d", i, got, want)
			}
		case uint64:
			if got.Uint64() != want {
				t.Errorf("#%d: got %v, want %d", i, got, want)
			}
		case bool:
			if constant.BoolVal(got.Value) != want {
				t.Errorf("#%d: got %v, want %v", i, got, want)
			}
		}
	}
}

func TestFoldFloat(t *testing.T) {
	if got := foldBinOp(token.ADD, fcst(1.5, types.Float64), fcst(2.25, types.Float64), types.Typ[types.Float64]); got == nil || got.Float64() != 3.75 {
		t.Errorf("1.5+2.25: got %v", got)
	}
	// float division by zero yields Inf: must not fold
	if got := foldBinOp(token.QUO, fcst(1, types.Float64), fcst(0, types.Float64), types.Typ[types.Float64]); got != nil {
		t.Errorf("1.0/0.0 folded to %v, want no fold", got)
	}
	// negative zero must not be folded (go/constant cannot represent it)
	if got := foldBinOp(token.MUL, fcst(-1, types.Float64), fcst(0, types.Float64), types.Typ[types.Float64]); got != nil {
		t.Errorf("-1.0*0.0 folded to %v, want no fold", got)
	}
	if got := foldUnOp(token.SUB, fcst(0, types.Float64), types.Typ[types.Float64]); got != nil {
		t.Errorf("-(0.0) folded to %v, want no fold", got)
	}
	// float32 rounding
	if got := foldBinOp(token.ADD, fcst(16777216, types.Float32), fcst(1, types.Float32), types.Typ[types.Float32]); got == nil || got.Float64() != 16777216 {
		t.Errorf("float32 rounding: got %v", got)
	}
}

func TestFoldConvert(t *testing.T) {
	// int8(-1) -> uint16 = 65535
	if got := foldConvert(icst(-1, types.Int8), types.Typ[types.Uint16]); got == nil || got.Uint64() != 65535 {
		t.Errorf("uint16(int8(-1)): got %v", got)
	}
	// int32(300) -> int8 = 44
	if got := foldConvert(icst(300, types.Int32), types.Typ[types.Int8]); got == nil || got.Int64() != 44 {
		t.Errorf("int8(300): got %v", got)
	}
	// float64(2.9) -> int = 2
	if got := foldConvert(fcst(2.9, types.Float64), types.Typ[types.Int]); got == nil || got.Int64() != 2 {
		t.Errorf("int(2.9): got %v", got)
	}
	// out of range float -> int: no fold
	if got := foldConvert(fcst(1e30, types.Float64), types.Typ[types.Int32]); got != nil {
		t.Errorf("int32(1e30) folded to %v, want no fold", got)
	}
	// negative float -> uint: no fold
	if got := foldConvert(fcst(-1, types.Float64), types.Typ[types.Uint]); got != nil {
		t.Errorf("uint(-1.0) folded to %v, want no fold", got)
	}
	// int -> float32 rounding
	if got := foldConvert(icst(16777217, types.Int64), types.Typ[types.Float32]); got == nil || got.Float64() != 16777216 {
		t.Errorf("float32(16777217): got %v", got)
	}
}

type testWriter struct{ t *testing.T }

func (w testWriter) Write(p []byte) (int, error) {
	w.t.Log(string(p))
	return len(p), nil
}
