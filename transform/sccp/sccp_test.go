package sccp

import (
	"go/ast"
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

type testWriter struct{ t *testing.T }

func (w testWriter) Write(p []byte) (int, error) {
	w.t.Log(string(p))
	return len(p), nil
}
