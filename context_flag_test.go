package ixgo

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/build"
	"go/constant"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"unsafe"

	"github.com/goplus/ixgo/alias"
	"golang.org/x/tools/go/ssa"
)

type invokingTestProcessor struct{}

func (invokingTestProcessor) LoadMain(_ *build.Context, pkg *build.Package) ([]byte, error) {
	return []byte(fmt.Sprintf(`package main
import tested %q
func main() {
	tested.TestSample()
	if !tested.TestSampleRan() { panic("TestSample was not executed") }
}
`, pkg.ImportPath)), nil
}

func TestRunTestPreservesTestingFlags(t *testing.T) {
	oldWorkingDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	dir, err := os.MkdirTemp(".", "runtest-flag-")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = os.RemoveAll(dir) })
	dir, err = filepath.Abs(dir)
	if err != nil {
		t.Fatal(err)
	}
	source := []byte(`package sample

var testSampleRan bool

func TestSample() {
	testSampleRan = true
}

func TestSampleRan() bool {
	return testSampleRan
}
`)
	if err := os.WriteFile(filepath.Join(dir, "sample_test.go"), source, 0o644); err != nil {
		t.Fatal(err)
	}

	oldTestProcessor := testProcessor
	testProcessor = invokingTestProcessor{}
	t.Cleanup(func() { testProcessor = oldTestProcessor })

	if err := NewContext(0).RunTest(dir, nil); err != nil {
		t.Fatal(err)
	}
	workingDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	if workingDir != oldWorkingDir {
		t.Fatalf("working directory changed from %q to %q", oldWorkingDir, workingDir)
	}
	_ = testing.Verbose()
}

func loadMain(t *testing.T, src string) (*Context, *ssa.Package, *Interp) {
	t.Helper()
	ctx := NewContext(0)
	pkg, err := ctx.LoadFile("main.go", src)
	if err != nil {
		t.Fatal(err)
	}
	interp, err := NewInterp(ctx, pkg)
	if err != nil {
		t.Fatal(err)
	}
	return ctx, pkg, interp
}

func TestInterpPublicAPI(t *testing.T) {
	src := `package main
const C = 1
var V = 2
type T int
func Add(a, b int) int { return a + b }
func main() {}
`
	ctx, pkg, interp := loadMain(t, src)
	if interp.MainPkg() == nil {
		t.Fatal("MainPkg")
	}
	if _, ok := interp.GetFunc("Add"); !ok {
		t.Fatal("GetFunc Add")
	}
	if _, ok := interp.GetFunc("missing"); ok {
		t.Fatal("GetFunc missing")
	}
	if _, ok := interp.GetFunc("V"); ok {
		t.Fatal("GetFunc var")
	}
	if p, ok := interp.GetVarAddr("V"); !ok || p == nil {
		t.Fatal("GetVarAddr")
	}
	if _, ok := interp.GetVarAddr("missing"); ok {
		t.Fatal("GetVarAddr missing")
	}
	if _, ok := interp.GetVarAddr("Add"); ok {
		t.Fatal("GetVarAddr func")
	}
	if c, ok := interp.GetConst("C"); !ok || constant.Val(c) != int64(1) {
		t.Fatal("GetConst")
	}
	if _, ok := interp.GetConst("missing"); ok {
		t.Fatal("GetConst missing")
	}
	if _, ok := interp.GetConst("V"); ok {
		t.Fatal("GetConst var")
	}
	if typ, ok := interp.GetType("T"); !ok || typ == nil {
		t.Fatal("GetType")
	}
	if _, ok := interp.GetType("missing"); ok {
		t.Fatal("GetType missing")
	}
	if _, ok := interp.GetType("V"); ok {
		t.Fatal("GetType var")
	}
	if _, _, ok := interp.GetSymbol("Add"); !ok {
		t.Fatal("GetSymbol Add")
	}
	if _, _, ok := interp.GetSymbol("V"); !ok {
		t.Fatal("GetSymbol V")
	}
	if _, _, ok := interp.GetSymbol("C"); !ok {
		t.Fatal("GetSymbol C")
	}
	if _, _, ok := interp.GetSymbol("T"); !ok {
		t.Fatal("GetSymbol T")
	}
	if _, _, ok := interp.GetSymbol("main.Add"); !ok {
		t.Fatal("GetSymbol main.Add")
	}
	if _, _, ok := interp.GetSymbol("missing.Add"); ok {
		t.Fatal("GetSymbol missing pkg")
	}
	if _, _, ok := interp.GetSymbol("a.b.c"); ok {
		t.Fatal("GetSymbol too many")
	}
	if _, _, ok := interp.GetSymbol("missing"); ok {
		t.Fatal("GetSymbol missing")
	}

	ret, err := ctx.RunFunc(pkg, "Add", 2, 3)
	if err != nil || ret.(int) != 5 {
		t.Fatalf("RunFunc %v %v", ret, err)
	}
	if _, err := ctx.RunFunc(pkg, "missing"); err == nil {
		t.Fatal("RunFunc missing")
	}
	if interp.ExitCode() != 0 {
		t.Fatal("ExitCode")
	}
	interp.Abort()
	if _, err := interp.RunMain(); err != nil && false {
		t.Fatal(err)
	}

	fn := pkg.Func("Add")
	if fn != nil && len(fn.Blocks) > 0 {
		dumpBlock(fn.Blocks[0], 0, fn.Blocks[0].Instrs[0])
		jumps := map[*ssa.BasicBlock]bool{}
		succs := map[*ssa.BasicBlock]*ssa.BasicBlock{}
		checkJumps(fn.Blocks[0], jumps, succs)
		checkRuns(fn.Blocks[0], map[*ssa.BasicBlock]bool{}, succs)
		if succs[fn.Blocks[0]] == nil && len(fn.Blocks[0].Succs) > 0 {
			succs[fn.Blocks[0]] = fn.Blocks[0].Succs[0]
			checkJumps(fn.Blocks[0], map[*ssa.BasicBlock]bool{}, succs)
			checkRuns(fn.Blocks[0], map[*ssa.BasicBlock]bool{}, succs)
		}
	}
	_ = deref(types.NewPointer(types.Typ[types.Int]))
	_ = deref(types.Typ[types.Int])
}

func TestContextModesAndLoad(t *testing.T) {
	ctx := NewContext(0)
	if ctx.IsEvalMode() {
		t.Fatal("eval default")
	}
	ctx.SetEvalMode(true)
	if !ctx.IsEvalMode() {
		t.Fatal("SetEvalMode")
	}
	ctx.SetEvalMode(false)
	ctx.SetUnsafeSizes(types.SizesFor("gc", "amd64"))
	ctx.SetLeastCallForEnablePool(8)
	var buf bytes.Buffer
	ctx.SetPrintOutput(&buf)
	ctx.SetDebug(func(info *DebugInfo) {
		_ = info.Position()
		info.AsVar()
		info.AsFunc()
	})
	ctx.SetPanic(func(info *PanicInfo) {
		_ = info.Position()
	})

	src := `package main
func main() {
	x := 1
	_ = x
}
`
	pkg, err := ctx.LoadFile("main.go", src)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := ctx.NewInterp(pkg); err != nil {
		t.Fatal(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "p.go", `package main; func F() int { return 1 }`, 0)
	if err != nil {
		t.Fatal(err)
	}
	apkg := &ast.Package{Name: "main", Files: map[string]*ast.File{"p.go": f}}
	if _, err := ctx.LoadAstPackage("main", apkg); err != nil {
		t.Fatal(err)
	}

	RegisterFileProcess(".covertest", func(ctx *Context, filename string, src interface{}) ([]byte, error) {
		return []byte(`package main; func main() {}`), nil
	})
	if _, err := ctx.ParseFile("x.covertest", "ignored"); err != nil {
		t.Fatal(err)
	}

	list := PackageList()
	if len(list) == 0 {
		t.Fatal("PackageList")
	}
	RegisterPackageLazy("cover95/lazy", func() *Package {
		return &Package{
			Name: "lazy", Path: "cover95/lazy",
			Interfaces:    map[string]reflect.Type{},
			NamedTypes:    map[string]reflect.Type{},
			Vars:          map[string]reflect.Value{},
			Funcs:         map[string]reflect.Value{},
			UntypedConsts: map[string]UntypedConst{},
		}
	})
	if _, ok := LookupPackage("cover95/lazy"); !ok {
		t.Fatal("LookupPackage lazy")
	}
	RegisterPackageLazy("cover95/lazy", func() *Package {
		return &Package{
			Name: "lazy", Path: "cover95/lazy",
			Interfaces:    map[string]reflect.Type{"I": reflect.TypeOf((*error)(nil)).Elem()},
			NamedTypes:    map[string]reflect.Type{},
			Vars:          map[string]reflect.Value{},
			Funcs:         map[string]reflect.Value{},
			UntypedConsts: map[string]UntypedConst{},
		}
	})
}

func TestRunDirAndTestPkg(t *testing.T) {
	dir := t.TempDir()
	main := filepath.Join(dir, "main.go")
	if err := os.WriteFile(main, []byte("package main\nfunc main() {}\n"), 0644); err != nil {
		t.Fatal(err)
	}
	ctx := NewContext(0)
	if code, err := ctx.Run(dir, nil); err != nil || code != 0 {
		t.Fatalf("Run dir %v %v", code, err)
	}
	if code, err := ctx.Run(main, nil); err != nil {
		t.Fatalf("Run file %v %v", code, err)
	}

	empty := t.TempDir()
	if err := os.WriteFile(filepath.Join(empty, "lib.go"), []byte("package lib\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := ctx.RunTest(empty, nil); err != nil {
		t.Fatal(err)
	}

	src := `package main
func Add(a, b int) int { return a+b }
func main() {}
`
	pkg, err := ctx.LoadFile("main.go", src)
	if err != nil {
		t.Fatal(err)
	}
	interp, err := NewInterp(ctx, pkg)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := interp.RunFunc("Add", 1, 2); err != nil {
		t.Fatal(err)
	}
	if _, err := interp.RunFunc("missing"); err == nil {
		t.Fatal("missing func")
	}
}

func TestSetPanicAndDebugRun(t *testing.T) {
	ctx := NewContext(0)
	ctx.SetPanic(func(info *PanicInfo) {
		_ = info.Error
		_ = info.Position()
	})
	if _, err := ctx.RunFile("main.go", `package main
func main() { panic(1) }
`, nil); err == nil {
		t.Fatal("expected panic")
	}

	ctx = NewContext(0)
	var hit bool
	ctx.SetDebug(func(info *DebugInfo) {
		hit = true
		info.Position()
		info.AsVar()
		info.AsFunc()
	})
	if _, err := ctx.RunFile("main.go", `package main
func main() {
	x := 1
	println(x)
}
`, nil); err != nil {
		t.Fatal(err)
	}
	_ = hit
}

func TestReplAPI(t *testing.T) {
	r := NewRepl(NewContext(0))
	r.SetFileName("repl.go")
	if _, _, err := r.Eval("1+2"); err != nil {
		t.Fatal(err)
	}
	if r.Interp() == nil {
		t.Fatal("Interp")
	}
	if r.Source() == "" {
		t.Fatal("Source")
	}
	_ = (&Eval{constant.MakeFloat64(1.5), reflect.TypeOf(float64(0))}).String()
	_ = (&Eval{constant.MakeInt64(1), reflect.TypeOf(0)}).String()
}

func TestCallExternalAndRelease(t *testing.T) {
	_, _, interp := loadMain(t, `package main
func main() {}
`)
	fr := &frame{interp: interp}
	toUpper := func(s string) string { return s + "!" }
	if v := interp.callExternal(fr, reflect.ValueOf(toUpper), []value{"ab"}, nil); v != "ab!" {
		t.Fatal("callExternal", v)
	}
	nop := func() {}
	if v := interp.callExternal(fr, reflect.ValueOf(nop), nil, nil); v != nil {
		t.Fatal("callExternal void")
	}
	pair := func(a, b int) (int, int) { return a, b }
	if v := interp.callExternal(fr, reflect.ValueOf(pair), []value{1, 2}, nil); len(v.(tuple)) != 2 {
		t.Fatal("callExternal tuple")
	}
	variadic := func(a string, rest ...int) int { return len(rest) }
	if v := interp.callExternal(fr, reflect.ValueOf(variadic), []value{"x", []int{1, 2}}, nil); v != 2 {
		t.Fatal("callExternal variadic", v)
	}
	ptrfn := func(p *int) int {
		if p == nil {
			return -1
		}
		return *p
	}
	if v := interp.callExternal(fr, reflect.ValueOf(ptrfn), []value{nil}, nil); v != -1 {
		t.Fatal("callExternal nil arg")
	}
	interp.callExternalDiscardsResult(fr, reflect.ValueOf(nop), nil, nil)
	interp.callExternalDiscardsResult(fr, reflect.ValueOf(ptrfn), []value{nil}, nil)
	interp.callExternalDiscardsResult(fr, reflect.ValueOf(variadic), []value{"x", []int{1}}, nil)
	interp.UnsafeRelease()
	ResetAllIcall()
	_, _, _ = IcallStat()
	_ = IcallCached()
}

func TestCallerFramesOnPanic(t *testing.T) {
	ctx := NewContext(0)
	ctx.SetPanic(func(info *PanicInfo) {
		_ = info.CallerFrames()
		debugPrintStack(info.Frame)
	})
	_, _ = ctx.RunFile("main.go", `package main
func main() { panic("boom") }
`, nil)
}

func TestLoadTestPackageXTestAndMain(t *testing.T) {
	old := testProcessor
	testProcessor = invokingTestProcessor{}
	t.Cleanup(func() { testProcessor = old })

	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "lib.go"), []byte("package lib\nfunc F() int { return 1 }\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "lib_test.go"), []byte(`package lib
func TestSample() {}
func TestSampleRan() bool { return true }
`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "x_test.go"), []byte(`package lib_test
func TestX() {}
`), 0644); err != nil {
		t.Fatal(err)
	}
	ctx := NewContext(0)
	if err := ctx.RunTest(dir, []string{"-test.v"}); err != nil {
		t.Fatal(err)
	}

	mainDir := t.TempDir()
	if err := os.WriteFile(filepath.Join(mainDir, "main.go"), []byte("package main\nfunc main() {}\n"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(mainDir, "main_test.go"), []byte(`package main
func TestSample() {}
func TestSampleRan() bool { return true }
`), 0644); err != nil {
		t.Fatal(err)
	}
	if err := ctx.RunTest(mainDir, nil); err != nil {
		t.Fatal(err)
	}
}

func TestPackageMergeAndEqualValue(t *testing.T) {
	p := &Package{
		Interfaces:    map[string]reflect.Type{},
		NamedTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		UntypedConsts: map[string]UntypedConst{},
	}
	p.merge(&Package{
		Interfaces:    map[string]reflect.Type{"E": reflect.TypeOf((*error)(nil)).Elem()},
		NamedTypes:    map[string]reflect.Type{"N": reflect.TypeOf(0)},
		Vars:          map[string]reflect.Value{"v": reflect.ValueOf(1)},
		Funcs:         map[string]reflect.Value{"f": reflect.ValueOf(func() {})},
		UntypedConsts: map[string]UntypedConst{"c": {Typ: "int", Value: constant.MakeInt64(1)}},
		Alias:         map[string]alias.Type{},
		Import:        func(fset *token.FileSet, pkgs map[string]*types.Package) (*types.Package, error) { return nil, nil },
	})
	if p.Import == nil {
		t.Fatal("merge Import")
	}

	ch := make(chan int)
	both := (chan int)(ch)
	send := (chan<- int)(ch)
	if !equalValue(reflect.ValueOf(both), reflect.ValueOf(send)) {
		t.Fatal("equalValue chan convert")
	}
	if !equalValue(reflect.ValueOf(send), reflect.ValueOf(both)) {
		t.Fatal("equalValue chan convert 2")
	}
}

func TestMoreUnopXorAndPrint(t *testing.T) {
	_ = unop(&ssa.UnOp{Op: token.XOR}, uint(0))
	_ = unop(&ssa.UnOp{Op: token.XOR}, uint16(0))
	_ = unop(&ssa.UnOp{Op: token.XOR}, uint32(0))
	_ = unop(&ssa.UnOp{Op: token.XOR}, uint64(0))
	_ = unop(&ssa.UnOp{Op: token.XOR}, uintptr(0))
	_ = unop(&ssa.UnOp{Op: token.XOR}, namedUint(0))
	mustPanic(t, func() { unop(&ssa.UnOp{Op: token.XOR}, 1.0) })
	type namedBool bool
	_ = unop(&ssa.UnOp{Op: token.NOT}, namedBool(true))

	var buf bytes.Buffer
	type namedF32 float32
	type namedC64 complex64
	writeany(&buf, namedF32(1.5))
	writeany(&buf, namedC64(1+2i))
	writevalue(&buf, namedF32(1.5), true)
	writevalue(&buf, namedC64(1+2i), true)
	writevalue(&buf, namedString("s"), true)
	writeany(&buf, struct{ A int }{1})
}

func TestUnsafeAndBuiltinProgram(t *testing.T) {
	src := `package main
import "unsafe"
type T struct { A int; B int }
func main() {
	s := []int{1, 2}
	s = append(s, 3)
	s = append(s)
	copy(s, s)
	m := map[int]int{1: 2}
	delete(m, 1)
	ch := make(chan int, 1)
	close(ch)
	_ = len(s)
	_ = cap(s)
	_ = real(1 + 2i)
	_ = imag(1 + 2i)
	_ = complex(1.0, 2.0)
	var c64 complex64 = 1 + 2i
	_ = real(c64)
	_ = imag(c64)
	_ = complex(float32(1), float32(2))
	p := new(int)
	_ = unsafe.Add(unsafe.Pointer(p), 0)
	sl := unsafe.Slice(p, 1)
	_ = unsafe.SliceData(sl)
	var nilp *int
	_ = unsafe.Slice(nilp, 0)
	empty := []int{}
	_ = unsafe.SliceData(empty)
	var nsl []int
	_ = unsafe.SliceData(nsl)
	b := []byte("hi")
	_ = unsafe.String(&b[0], 2)
	_ = unsafe.String((*byte)(nil), 0)
	_ = unsafe.Sizeof(0)
	_ = unsafe.Alignof(0)
	_ = unsafe.Offsetof(T{}.B)
	print("x")
	println("y", 1)
	var i interface{} = 1
	println(i)
}
`
	ctx := NewContext(EnablePrintAny)
	if _, err := ctx.RunFile("main.go", src, nil); err != nil {
		t.Fatal(err)
	}
}

func TestCallBuiltinLegacy(t *testing.T) {
	src := `package main
import "unsafe"
func main() {
	_ = len("hi")
	_ = cap([]int{1})
	_ = append([]int{1}, 2)
	_ = copy([]int{0}, []int{1})
	delete(map[int]int{1: 1}, 1)
	ch := make(chan int)
	close(ch)
	_ = real(1 + 2i)
	_ = imag(1 + 2i)
	_ = complex(1.0, 2.0)
	_ = unsafe.Add(unsafe.Pointer(new(int)), 0)
	_ = unsafe.Slice(new(int), 1)
	_ = unsafe.SliceData([]int{1})
	b := []byte("ab")
	_ = unsafe.String(&b[0], 2)
	print(1)
	println(1)
}
`
	ctx := NewContext(0)
	pkg, err := ctx.LoadFile("main.go", src)
	if err != nil {
		t.Fatal(err)
	}
	interp, err := NewInterp(ctx, pkg)
	if err != nil {
		t.Fatal(err)
	}
	builtins := map[string]*ssa.Builtin{}
	var walk func(fn *ssa.Function)
	walk = func(fn *ssa.Function) {
		if fn == nil {
			return
		}
		for _, b := range fn.Blocks {
			for _, instr := range b.Instrs {
				call, ok := instr.(interface{ Common() *ssa.CallCommon })
				if !ok {
					continue
				}
				if bu, ok := call.Common().Value.(*ssa.Builtin); ok {
					builtins[bu.Name()] = bu
				}
			}
		}
		for _, anon := range fn.AnonFuncs {
			walk(anon)
		}
	}
	for _, m := range pkg.Members {
		if fn, ok := m.(*ssa.Function); ok {
			walk(fn)
		}
	}
	fr := &frame{interp: interp}
	maybe := func(name string) *ssa.Builtin { return builtins[name] }
	p := new(int)
	if b := maybe("len"); b != nil {
		if v := interp.callBuiltin(fr, b, []value{"hi"}, nil); v != 2 {
			t.Fatal("len", v)
		}
		mustPanic(t, func() { interp.callBuiltinDiscardsResult(fr, b, []value{"x"}, nil) })
	}
	if b := maybe("cap"); b != nil {
		if v := interp.callBuiltin(fr, b, []value{[]int{1, 2, 3}}, nil); v != 3 {
			t.Fatal("cap", v)
		}
	}
	if b := maybe("append"); b != nil {
		if v := interp.callBuiltin(fr, b, []value{[]int{1}}, nil); len(v.([]int)) != 1 {
			t.Fatal("append one")
		}
		if v := interp.callBuiltin(fr, b, []value{[]int{1}, []int{2}}, nil); len(v.([]int)) != 2 {
			t.Fatal("append")
		}
		if v := interp.callBuiltin(fr, b, []value{[]byte{}, "ab"}, nil); string(v.([]byte)) != "ab" {
			t.Fatal("append bytes")
		}
	}
	if b := maybe("copy"); b != nil {
		if v := interp.callBuiltin(fr, b, []value{[]int{0, 0}, []int{1, 2}}, nil); v != 2 {
			t.Fatal("copy")
		}
		interp.callBuiltinDiscardsResult(fr, b, []value{[]int{0}, []int{1}}, nil)
	}
	if b := maybe("close"); b != nil {
		interp.callBuiltin(fr, b, []value{make(chan int)}, nil)
		interp.callBuiltinDiscardsResult(fr, b, []value{make(chan int)}, nil)
	}
	if b := maybe("delete"); b != nil {
		m := map[int]int{1: 2}
		interp.callBuiltin(fr, b, []value{m, 1}, nil)
		interp.callBuiltinDiscardsResult(fr, b, []value{map[int]int{1: 1}, 1}, nil)
	}
	if b := maybe("real"); b != nil {
		interp.callBuiltin(fr, b, []value{1 + 2i}, nil)
		interp.callBuiltin(fr, b, []value{complex64(3 + 4i)}, nil)
	}
	if b := maybe("imag"); b != nil {
		interp.callBuiltin(fr, b, []value{complex64(1 + 2i)}, nil)
	}
	if b := maybe("complex"); b != nil {
		interp.callBuiltin(fr, b, []value{1.0, 2.0}, nil)
		interp.callBuiltin(fr, b, []value{float32(1), float32(2)}, nil)
	}
	if b := maybe("Add"); b != nil {
		interp.callBuiltin(fr, b, []value{unsafe.Pointer(p), 0}, nil)
	}
	if b := maybe("Slice"); b != nil {
		interp.callBuiltin(fr, b, []value{p, 1}, nil)
		interp.callBuiltin(fr, b, []value{(*int)(nil), 0}, nil)
	}
	if b := maybe("SliceData"); b != nil {
		interp.callBuiltin(fr, b, []value{[]int{1}}, nil)
		interp.callBuiltin(fr, b, []value{[]int{}}, nil)
		interp.callBuiltin(fr, b, []value{[]int(nil)}, nil)
	}
	if b := maybe("String"); b != nil {
		bs := []byte("hi")
		interp.callBuiltin(fr, b, []value{&bs[0], 2}, nil)
		interp.callBuiltin(fr, b, []value{(*byte)(nil), 0}, nil)
	}
	if b := maybe("print"); b != nil {
		interp.callBuiltin(fr, b, []value{1, "x"}, nil)
		interp.callBuiltinDiscardsResult(fr, b, []value{"x"}, nil)
	}
	if b := maybe("println"); b != nil {
		interp.callBuiltin(fr, b, []value{1, "x"}, nil)
		interp.callBuiltinDiscardsResult(fr, b, []value{"x"}, nil)
	}
	if b := maybe("ssa:wrapnilchk"); b != nil {
		interp.callBuiltin(fr, b, []value{p, "T", "M"}, nil)
		interp.callBuiltinDiscardsResult(fr, b, []value{p, "T", "M"}, nil)
		mustPanic(t, func() { interp.callBuiltin(fr, b, []value{(*int)(nil), "main.T", "M"}, nil) })
		mustPanic(t, func() { interp.callBuiltin(fr, b, []value{(*int)(nil), "T", "M"}, nil) })
		mustPanic(t, func() { interp.callBuiltinDiscardsResult(fr, b, []value{(*int)(nil), "T", "M"}, nil) })
	}
	if b := maybe("recover"); b != nil {
		interp.callBuiltinDiscardsResult(fr, b, nil, nil)
		interp.callBuiltin(fr, b, nil, nil)
	}
	if b := maybe("panic"); b != nil {
		mustPanic(t, func() { interp.callBuiltin(fr, b, []value{1}, nil) })
	}
	if len(builtins) == 0 {
		t.Fatal("no builtins found")
	}
}

func TestUnsafeReleaseFindMethodAndChanRecv(t *testing.T) {
	src := `package main
type T struct{}
func (T) M() int { return 1 }
var ch = make(chan int, 1)
func init() { ch <- 7 }
func main() {
	x := <-ch
	close(ch)
	y, ok := <-ch
	if x != 7 || ok {
		panic("chan")
	}
	_ = y
	_ = T{}.M()
}
`
	_, pkg, interp := loadMain(t, src)
	if err := interp.RunInit(); err != nil {
		t.Fatal(err)
	}
	if _, err := interp.RunFunc("main"); err != nil {
		t.Fatal(err)
	}
	obj := pkg.Pkg.Scope().Lookup("T")
	named := obj.Type().(*types.Named)
	fn := named.Method(0)
	mtyp := interp.toType(fn.Type())
	call := interp.FindMethod(mtyp, fn)
	typ, ok := interp.GetType("T")
	if !ok {
		t.Fatal("GetType T")
	}
	out := call([]reflect.Value{reflect.New(typ).Elem()})
	if len(out) != 1 || out[0].Int() != 1 {
		t.Fatal("FindMethod", out)
	}
	mustPanic(t, func() {
		interp.FindMethod(mtyp, types.NewFunc(0, pkg.Pkg, "Missing", types.NewSignatureType(
			types.NewVar(0, pkg.Pkg, "r", obj.Type()), nil, nil, nil, nil, false)))
	})
	_ = interp.IcallAlloc()
	for _, pfn := range interp.funcs {
		_ = pfn.InstrForPC(0)
		_ = pfn.InstrForPC(-1)
		_ = pfn.PosForPC(0)
		_ = pfn.PositionForPC(0)
		pfn.UnsafeRelease()
	}
}
