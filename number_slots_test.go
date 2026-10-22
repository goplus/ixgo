package ixgo

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"testing"

	"golang.org/x/tools/go/ssa"
)

func numberDivmodSource(tb testing.TB) string {
	tb.Helper()
	data, err := os.ReadFile(filepath.Join(runtime.GOROOT(), "test", "divmod.go"))
	if err != nil {
		tb.Fatal(err)
	}
	return string(data)
}
func numberInterp(tb testing.TB, enabled bool, source string) *Interp {
	tb.Helper()
	ctx := NewContext(DisableMemo)
	ctx.NumberSlots = enabled
	ctx.SetLeastCallForEnablePool(0)
	i, err := ctx.LoadInterp("numbers.go", source)
	if err != nil {
		tb.Fatal(err)
	}
	tb.Cleanup(i.UnsafeRelease)
	return i
}
func numberSelected(i *Interp, name string) int {
	for fn, p := range i.funcs {
		if fn.Name() == name {
			return p.numberSlots
		}
	}
	return 0
}

func TestNumberSlots(t *testing.T) {
	src := numberDivmodSource(t)
	for _, on := range []bool{false, true} {
		t.Run(fmt.Sprint(on), func(t *testing.T) {
			i := numberInterp(t, on, src)
			t.Logf("NumberSlots=%v udiv=%d idiv=%d memo=%v", on, numberSelected(i, "udiv"), numberSelected(i, "idiv"), i.funcs[i.mainpkg.Func("udiv")] != nil && i.funcs[i.mainpkg.Func("udiv")].memoFlags != 0)
			if (numberSelected(i, "udiv") > 0) != on || numberSelected(i, "idiv") != 0 {
				i.mainpkg.Func("udiv").WriteTo(os.Stderr)
				t.Fatal("selection mismatch")
			}
			for _, x := range []uint64{0, 1, 255, 256, 1 << 32, 1 << 63, ^uint64(0)} {
				for _, y := range []uint64{1, 2, 3, 255, 1 << 32, 1 << 63, ^uint64(0)} {
					got, err := i.RunFunc("udiv", x, y)
					want := Tuple{x / y, x % y}
					if err != nil || !reflect.DeepEqual(got, want) {
						t.Fatalf("%d/%d: %v %v want %v", x, y, got, err, want)
					}
				}
			}
			rng := rand.New(rand.NewSource(1))
			for n := 0; n < 1000; n++ {
				x, y := rng.Uint64(), rng.Uint64()|1
				got, err := i.RunFunc("udiv", x, y)
				if err != nil || !reflect.DeepEqual(got, Tuple{x / y, x % y}) {
					t.Fatalf("random udiv %d/%d: %v %v", x, y, got, err)
				}
			}
			for _, x := range []int64{-1 << 63, -65537, -1, 0, 1, 65537, 1<<63 - 1} {
				for _, y := range []int64{-1 << 63, -3, -1, 1, 3, 1<<63 - 1} {
					got, err := i.RunFunc("idiv", x, y)
					want := Tuple{x / y, x % y}
					if err != nil || !reflect.DeepEqual(got, want) {
						t.Fatalf("idiv %d/%d: %v %v", x, y, got, err)
					}
				}
			}
		})
	}
}

func TestNumberSlotsFallback(t *testing.T) {
	for _, mode := range []string{"memo", "hooks", "leave", "panic", "debug", "trace", "globaldebug", "execution", "eval", "evalcall"} {
		t.Run(mode, func(t *testing.T) {
			oldMemo := memoPlugin
			defer func() { memoPlugin = oldMemo }()
			ctx := NewContext(DisableMemo)
			ctx.NumberSlots = true
			switch mode {
			case "memo":
				ctx.Mode &^= DisableMemo
				memoPlugin = func(MemoInstaller, *ssa.Function) {}
			case "hooks":
				ctx.entryFunc = func(*FuncInfo) {}
			case "leave":
				ctx.leaveFunc = func(*FuncInfo) {}
			case "panic":
				ctx.panicFunc = func(*PanicInfo) {}
			case "debug":
				ctx.debugFunc = func(*DebugInfo) {}
			case "trace":
				ctx.Mode |= EnableTracing
			case "globaldebug":
				ctx.BuilderMode |= ssa.GlobalDebug
			case "execution":
				tracker := NewExecutionTracker(func(ExecutionPosition) {})
				defer tracker.Close()
				ctx.SetExecutionTracker(tracker)
			case "eval":
				ctx.evalMode = true
			case "evalcall":
				ctx.evalCallFn = func(*Interp, *ssa.Call, ...interface{}) {}
			}
			i, err := ctx.LoadInterp("numbers.go", numberDivmodSource(t))
			if err != nil {
				t.Fatal(err)
			}
			defer i.UnsafeRelease()
			if numberSelected(i, "udiv") != 0 {
				t.Fatal("instrumented/plugin function selected")
			}
		})
	}
	i := numberInterp(t, true, frameCallsSource)
	if numberSelected(i, "fibish") != 0 || numberSelected(i, "add") != 0 {
		t.Fatal("recursive or short function selected")
	}
	got, err := i.RunFunc("fibish", 20)
	if err != nil || !reflect.DeepEqual(got, Tuple{2584, 8362}) {
		t.Fatal(got, err)
	}
}

func TestNumberSlotsPhi(t *testing.T) {
	src := `package main
func swap(a,b uint64,n int)(uint64,uint64) { for ; n>0; n-- { a,b=b,a; n=n^7; n=n^7; n=n+3; n=n-3; n=n^9; n=n^9; n=n+5; n=n-5 }; return a,b }
func short(n int) int { return n+1 }
func shift(a uint64,n int) uint64 { for j:=0;j<1;j++ { a=a^7; a=a^7; a=a+3; a=a-3; a=a^9; a=a^9; a=a+5; a=a-5 }; return a<<n }
func unsupported(a uint64) uint64 { return a/3 }
`
	i := numberInterp(t, true, src)
	for _, n := range []int{1, 2, 3, 100} {
		got, err := i.RunFunc("swap", uint64(999), uint64(888), n)
		want := Tuple{uint64(999), uint64(888)}
		if n%2 != 0 {
			want[0], want[1] = want[1], want[0]
		}
		if err != nil || !reflect.DeepEqual(got, want) {
			t.Fatal(got, err, want)
		}
	}
	if numberSelected(i, "swap") == 0 || numberSelected(i, "shift") == 0 || numberSelected(i, "short") != 0 || numberSelected(i, "unsupported") != 0 {
		t.Fatalf("selection mismatch: swap=%d shift=%d", numberSelected(i, "swap"), numberSelected(i, "shift"))
	}
	for _, n := range []int{0, 1, 63, 64, 100} {
		got, err := i.RunFunc("shift", uint64(3), n)
		if err != nil || got != uint64(3)<<n {
			t.Fatal(got, err)
		}
	}
	if _, err := i.RunFunc("shift", uint64(3), -1); err == nil {
		t.Fatal("missing negative shift panic")
	}
}

func TestNumberSlotsConcurrent(t *testing.T) {
	i := numberInterp(t, true, `package main
 func leaf(x uint64) uint64 { for n:=0;n<100;n++ { x=x*3+1; x=x^7; x=x^7; x=x+3; x=x-3; x=x^9; x=x^9 }; return x }
 func main() { ch:=make(chan uint64,8); for n:=0;n<8;n++ { go func(){ ch<-leaf(123) }() }; want:=leaf(123); for n:=0;n<8;n++ { if <-ch!=want { panic("numeric frame reuse") } } }
 `)
	if numberSelected(i, "leaf") == 0 {
		t.Fatal("leaf not selected")
	}
	if _, err := i.RunFunc("main"); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkNumberSlots(b *testing.B) {
	for _, work := range []string{"udiv", "idiv", "short", "fibish20", "fibish40", "divmod"} {
		for _, on := range []bool{false, true} {
			b.Run(fmt.Sprintf("%s/%v", work, on), func(b *testing.B) {
				src := frameCallsSource
				name := work
				var args []Value
				switch work {
				case "udiv":
					src = numberDivmodSource(b)
					args = []Value{uint64(0xfedcba9876543210), uint64(3)}
				case "idiv":
					src = numberDivmodSource(b)
					args = []Value{int64(-0x7edcba9876543210), int64(3)}
				case "divmod":
					src = numberDivmodSource(b)
					name = "main"
				case "short":
					name = "add"
					args = []Value{1000}
				case "fibish20":
					name = "fibish"
					args = []Value{20}
				case "fibish40":
					name = "fibish"
					args = []Value{40}
				}
				i := numberInterp(b, on, src)
				var output bytes.Buffer
				i.ctx.SetPrintOutput(&output)
				var selected []string
				for fn, p := range i.funcs {
					if p.numberSlots > 0 {
						selected = append(selected, fmt.Sprintf("%s:%d", fn.String(), p.numberSlots))
					}
					if p.memoFlags != 0 {
						b.Fatal("optimizer unexpectedly active", fn)
					}
				}
				sort.Strings(selected)
				b.Logf("selected=%v", selected)
				b.Logf("NumberSlots=%v target=%s slots=%d udiv=%d memo=%v", on, name, numberSelected(i, name), numberSelected(i, "udiv"), i.funcs[i.mainpkg.Func("udiv")] != nil && i.funcs[i.mainpkg.Func("udiv")].memoFlags != 0)
				b.ReportAllocs()
				b.ResetTimer()
				for n := 0; n < b.N; n++ {
					if _, err := i.RunFunc(name, args...); err != nil {
						b.Fatal(err)
					}
				}
				b.StopTimer()
				if output.Len() != 0 {
					b.Fatalf("unexpected divmod diagnostics: %s", output.String())
				}
			})
		}
	}
}

func TestNumberSlotsArithmetic(t *testing.T) {
	const source = `package main
func signed(a,b int, flag bool) (int,int,int,bool) {
 for j:=0;j<2;j++ { a=a*3+b; b=(b^a)&^7; a=-a; b=^b }
 return a+b, a-b, a>>2, !(flag != (a<b))
}
func unsigned(a,b uint64) (uint64,uint64,bool) {
 for j:=0;j<2;j++ { a=a*3+b; b=(b^a)&^7; a=-a; b=^b }
 return a+b, a-b, a>=b
}
func shifts(a,n int) (int,int) {
 for j:=0;j<1;j++ { a=a^7; a=a^7; a=a+3; a=a-3; a=a^9; a=a^9; a=a+5; a=a-5 }
 return a<<n,a>>n
}
func ushifts(a,n uint64) (uint64,uint64) {
 for j:=0;j<1;j++ { a=a^7; a=a^7; a=a+3; a=a-3; a=a^9; a=a^9; a=a+5; a=a-5 }
 return a<<n,a>>n
}`
	boxed, numeric := numberInterp(t, false, source), numberInterp(t, true, source)
	for _, name := range []string{"signed", "unsigned", "shifts", "ushifts"} {
		if numberSelected(numeric, name) == 0 {
			t.Fatalf("%s not selected", name)
		}
	}
	check := func(name string, args ...Value) {
		t.Helper()
		want, we := boxed.RunFunc(name, args...)
		got, ge := numeric.RunFunc(name, args...)
		if (we == nil) != (ge == nil) || (we == nil && !reflect.DeepEqual(want, got)) {
			t.Fatalf("%s(%v): got %v, %v; want %v, %v", name, args, got, ge, want, we)
		}
	}
	max := int(^uint(0) >> 1)
	for _, a := range []int{-max - 1, -max, -1, 0, 1, max} {
		for _, b := range []int{-max - 1, -1, 0, 1, max} {
			check("signed", a, b, true)
			check("signed", a, b, false)
		}
		for _, n := range []int{-1, 0, 1, 31, 32, 63, 64, 100, max} {
			check("shifts", a, n)
		}
	}
	for _, a := range []uint64{0, 1, 1 << 63, ^uint64(0)} {
		for _, b := range []uint64{0, 1, 1 << 63, ^uint64(0)} {
			check("unsigned", a, b)
		}
		for _, n := range []uint64{0, 1, 63, 64, 100, ^uint64(0)} {
			check("ushifts", a, n)
		}
	}
}

func TestNumberSlotsExactTypeFallback(t *testing.T) {
	const body = ` { for j:=0;j<2;j++ { a=a*3+1; a=a^7; a=a^7; a=a+3; a=a-3; a=a^9; a=a^9 }; return a }`
	source := "package main\ntype named uint64\ntype alias = uint64\n"
	for _, typ := range []string{"int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uintptr", "named", "alias"} {
		source += "func f" + typ + "(a " + typ + ") " + typ + body + "\n"
	}
	i := numberInterp(t, true, source)
	for _, typ := range []string{"int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uintptr", "named", "alias"} {
		if selected := numberSelected(i, "f"+typ) > 0; selected {
			t.Fatalf("unexpected selection for %s: %v", typ, selected)
		}
	}
}

func TestNumberSlotsThreshold(t *testing.T) {
	for _, count := range []int{minNumberSlotInstructions - 1, minNumberSlotInstructions} {
		source := "package main\nfunc leaf(a uint64) uint64 {"
		for n := 0; n < count-1; n++ {
			source += "a=a+3;"
		}
		source += "return a}"
		i := numberInterp(t, true, source)
		fn := i.mainpkg.Func("leaf")
		instructions := 0
		for _, b := range fn.Blocks {
			instructions += len(b.Instrs)
		}
		if instructions != count || (numberSelected(i, "leaf") > 0) != (count >= minNumberSlotInstructions) {
			t.Fatalf("threshold mismatch: instructions=%d selected=%d", instructions, numberSelected(i, "leaf"))
		}
	}
	if NewContext(0).NumberSlots {
		t.Fatal("number slots enabled by default")
	}
}

func TestNumberSlotsGC(t *testing.T) {
	ctx := NewContext(ExperimentalSupportGC | DisableMemo)
	ctx.NumberSlots = true
	ctx.SetLeastCallForEnablePool(0)
	i, err := ctx.LoadInterp("numbers.go", `package main
import "runtime"
func leaf(x uint64) uint64 { for n:=0;n<10;n++ { x=x*3+1; x=x^7; x=x^7; x=x+3; x=x-3; x=x^9; x=x^9 }; return x }
func main() { x:=new(uint64); *x=123; want:=leaf(*x); for n:=0;n<3;n++ { runtime.GC(); if leaf(*x)!=want { panic("numeric GC") } } }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	if !i.gcEnabled || numberSelected(i, "leaf") == 0 {
		t.Fatal("GC or numeric path not enabled")
	}
	if _, err := i.RunFunc("main"); err != nil {
		t.Fatal(err)
	}
}

func TestNumberSlotsFrameReuse(t *testing.T) {
	i := numberInterp(t, true, `package main
func leaf(x uint64) uint64 { for n:=0;n<2;n++ { x=x*3+1; x=x^7; x=x^7; x=x+3; x=x-3; x=x^9; x=x^9 }; return x }
`)
	p := i.funcs[i.mainpkg.Func("leaf")]
	if p.numberSlots == 0 {
		t.Fatal("leaf not selected")
	}
	fr := &frame{pfn: p, block: p.Main, stack: append([]value(nil), p.stack...)}
	var storage *uint64
	for _, x := range []uint64{0, 123, ^uint64(0), 7} {
		fr.ipc, fr.pred, fr.block = 0, 0, p.Main
		fr.stack[p.nres] = x
		for fr.ipc >= 0 {
			pc := fr.ipc
			fr.ipc++
			p.Instrs[pc](fr)
		}
		if got, want := fr.stack[0], (x*3+1)*3+1; got != want {
			t.Fatalf("got %v, want %v", got, want)
		}
		if storage != nil && storage != &fr.numberSlots()[0] {
			t.Fatal("numeric storage was not reused")
		}
		storage = &fr.numberSlots()[0]
		if fr.stack[p.nres] != x {
			t.Fatal("boxed argument changed")
		}
	}
}
