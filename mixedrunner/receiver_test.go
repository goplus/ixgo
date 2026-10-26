package mixedrunner

import (
	"go/constant"
	"go/types"
	"reflect"
	"runtime"
	"sync"
	"testing"

	"github.com/goplus/ixgo"
	"golang.org/x/tools/go/ssa"
)

func TestNumericReceiverFields(t *testing.T) {
	for _, typ := range []string{"int", "uint", "int8", "uint8", "int64", "uint64", "float32", "float64"} {
		t.Run(typ, func(t *testing.T) {
			source := `package main
 type S struct { x ` + typ + ` }
 func read(p *S) ` + typ + ` { return p.x }
 func update(p *S, n int) ` + typ + ` {
 if n == 0 { return read(p) }
 p.x += 2
 return update(p, n-1)
 }
 func check() bool { p := &S{x: 3}; return update(p, 4) == 11 && p.x == 11 }
 func nilRead() { var p *S; read(p) }
 `
			pkg, err := ixgo.NewContext(ixgo.DisableMemo).LoadFile("receiver.go", source)
			if err != nil {
				t.Fatal(err)
			}
			r := MakeRunner(pkg.Func("update"))
			if r == nil {
				t.Fatal("receiver graph rejected")
			}
			binder := r.(ixgo.TypedRunnerTypeBinder)
			if binder.BindTypes(func(_ types.Type) reflect.Type { return nil }) {
				t.Fatal("unresolved receiver type accepted")
			}
			fieldType := reflect.TypeOf(int(0))
			switch typ {
			case "uint":
				fieldType = reflect.TypeOf(uint(0))
			case "int8":
				fieldType = reflect.TypeOf(int8(0))
			case "uint8":
				fieldType = reflect.TypeOf(uint8(0))
			case "int64":
				fieldType = reflect.TypeOf(int64(0))
			case "uint64":
				fieldType = reflect.TypeOf(uint64(0))
			case "float32":
				fieldType = reflect.TypeOf(float32(0))
			case "float64":
				fieldType = reflect.TypeOf(float64(0))
			}
			structType := reflect.StructOf([]reflect.StructField{{Name: "X", Type: fieldType}})
			if !binder.BindTypes(func(types.Type) reflect.Type { return reflect.PointerTo(structType) }) {
				t.Fatal("receiver binding failed")
			}
			ptr := reflect.New(structType)
			ptr.Elem().Field(0).Set(reflect.ValueOf(3).Convert(fieldType))
			got := r.(ixgo.TypedArgumentsRunner).RunArgs([]ixgo.Value{ptr.Interface(), 4})
			want := reflect.ValueOf(11).Convert(fieldType).Interface()
			if got != want || ptr.Elem().Field(0).Interface() != want {
				t.Fatalf("direct receiver update: %v, field %v, want %v", got, ptr.Elem().Field(0), want)
			}
			i, err := ixgo.NewContext(ixgo.DisableMemo).LoadInterp("receiver.go", source)
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			if got, err := i.RunFunc("check"); err != nil || got != true {
				t.Fatalf("field updates: %v, %v", got, err)
			}
			if _, err := i.RunFunc("nilRead"); err == nil {
				t.Fatal("nil receiver did not panic")
			}
		})
	}
}

func TestMixedFrameReferences(t *testing.T) {
	const source = `package main
 type S struct { x int }
 func numeric(n int) int { if n == 0 { return 1 }; return numeric(n-1)+1 }
 func first(p *S, n int) int { p.x += n; return p.x }
 func second(n int, p *S) int { p.x += n; return p.x }
 func recurse(p *S, n int) int {
 if n == 0 { return p.x }
 x := numeric(4)
 first(p, x)
 second(x, p)
 return recurse(p, n-1)
 }
 `
	pkg, err := ixgo.NewContext(ixgo.DisableMemo).LoadFile("refs.go", source)
	if err != nil {
		t.Fatal(err)
	}
	numeric := MakeRunner(pkg.Func("numeric")).(*mixedTypedRunner)
	if numeric.refSlots != 0 || len(numeric.refParams) != 0 {
		t.Fatal("numeric runner requires reference storage")
	}
	if !numeric.numeric {
		t.Fatal("numeric graph did not select the numeric loop")
	}
	for _, fast := range []bool{true, false} {
		numeric.numeric = fast
		if got := numeric.RunArgs([]ixgo.Value{32}); got != 33 {
			t.Fatalf("numeric loop %v: got %v, want 33", fast, got)
		}
	}
	numeric.numeric = true
	if size := reflect.TypeOf(mixedNumericFrame{}).Size(); size > mixedArgsLimit*8+64 {
		t.Fatalf("numeric frame is too large: %d bytes", size)
	}
	// A numeric frame needs only its registers and small call metadata.
	if size := reflect.TypeOf(mixedFrame{}).Size(); size > mixedArgsLimit*8+128 {
		t.Fatalf("numeric frame is too large: %d bytes", size)
	}
	r := MakeRunner(pkg.Func("recurse")).(*mixedTypedRunner)
	if r.numeric {
		t.Fatal("receiver graph selected the numeric loop")
	}
	for _, node := range r.graph {
		if node.refSlots > len(node.constants) {
			t.Fatal("reference storage exceeds register count")
		}
	}
	structType := reflect.StructOf([]reflect.StructField{{Name: "X", Type: reflect.TypeOf(0)}})
	if !r.BindTypes(func(types.Type) reflect.Type { return reflect.PointerTo(structType) }) {
		t.Fatal("receiver binding failed")
	}
	var wg sync.WaitGroup
	for range 4 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 20 {
				ptr := reflect.New(structType)
				if got := r.RunArgs([]ixgo.Value{ptr.Interface(), 32}); got != 320 {
					t.Errorf("mixed pointer/numeric calls: got %v, want 320", got)
				}
				if got := ptr.Elem().Field(0).Int(); got != 320 {
					t.Errorf("receiver field: got %v, want 320", got)
				}
			}
		}()
	}
	for range 10 {
		runtime.GC()
	}
	wg.Wait()
}

func TestMixedRejectBoolStore(t *testing.T) {
	pkg, err := ixgo.NewContext(0).LoadFile("store.go", `package main
 type S struct { x int }
 func store(p *S) { p.x = 1 }
 `)
	if err != nil {
		t.Fatal(err)
	}
	fn := pkg.Func("store")
	for _, block := range fn.Blocks {
		for _, instr := range block.Instrs {
			if store, ok := instr.(*ssa.Store); ok {
				// A malformed SSA store must not reach the numeric field writer.
				store.Val = ssa.NewConst(constant.MakeBool(true), types.Typ[types.Bool])
				if MakeRunner(fn) != nil {
					t.Fatal("bool store accepted")
				}
				return
			}
		}
	}
	t.Fatal("missing store instruction")
}

func BenchmarkMixedNumericFrames(b *testing.B) {
	pkg, err := ixgo.NewContext(0).LoadFile("frames.go", `package main
 func leaf(n int) int { return n+1 }
 func shallow(n int) int { return leaf(n)+leaf(n) }
 func recursive(n int) int { if n == 0 { return 0 }; return recursive(n-1)+1 }
 `)
	if err != nil {
		b.Fatal(err)
	}
	for _, name := range []string{"leaf", "shallow", "recursive"} {
		r := MakeRunner(pkg.Func(name)).(ixgo.TypedArgumentsRunner)
		args := []ixgo.Value{16}
		b.Run(name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				if got := r.RunArgs(args); got == nil {
					b.Fatal("missing result")
				}
			}
		})
	}
}

func TestMixedReceiverFieldIndex(t *testing.T) {
	const source = `package main
 type S struct { unused int; x int }
 func read(p *S) int { return p.x }
 func write(p *S, n int) int { p.x = n; return read(p) }
 `
	pkg, err := ixgo.NewContext(0).LoadFile("fields.go", source)
	if err != nil {
		t.Fatal(err)
	}
	structType := reflect.StructOf([]reflect.StructField{{Name: "Unused", Type: reflect.TypeOf(0)}, {Name: "X", Type: reflect.TypeOf(0)}})
	for _, name := range []string{"read", "write"} {
		r := MakeRunner(pkg.Func(name)).(*mixedTypedRunner)
		if !r.BindTypes(func(types.Type) reflect.Type { return reflect.PointerTo(structType) }) {
			t.Fatal("receiver binding failed")
		}
		ptr := reflect.New(structType)
		ptr.Elem().Field(0).SetInt(99)
		ptr.Elem().Field(1).SetInt(42)
		args := []ixgo.Value{ptr.Interface()}
		if name == "write" {
			args = append(args, 42)
		}
		if got := r.RunArgs(args); got != 42 || ptr.Elem().Field(0).Int() != 99 {
			t.Fatalf("%s: wrong field access: %v, %v", name, got, ptr.Elem())
		}
	}
	if size := reflect.TypeOf(mixedOp{}).Size(); size > 104 {
		t.Fatalf("mixed instruction is too large: %d bytes", size)
	}
}
