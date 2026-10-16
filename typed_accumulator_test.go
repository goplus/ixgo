package ixgo

import (
	"os"
	"strings"
	"testing"
)

func TestTypedAccumulatorMemo(t *testing.T) {
	source, err := os.ReadFile("testdata/fibish/uglyfib.go")
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := NewContext(0).LoadFile("ugly.go", source)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"f", "g", "h", "k"} {
		fast := MakeTypedRunnerOf[int](pkg.Func(name)).(*typedArgsRunner[int])
		slow := makeTypedArgumentsRunnerOf[int](pkg.Func(name)).(*typedArgsRunner[int])
		if !fast.accumulatorMemo {
			t.Fatalf("%s: memo disabled", name)
		}
		for _, alias := range []bool{false, true} {
			for n := 0; n <= 18; n++ {
				if alias && n > 3 {
					continue
				}
				a, b, out := n-1, n-2, int(^uint(0)>>1)-3
				a2, b2, out2 := a, b, out
				p, p2 := &out, &out2
				if alias {
					a, a2 = -1, -1
					p, p2 = &a, &a2
				}
				fast.RunArgs([]Value{n, &a, &b, p})
				slow.RunArgs([]Value{n, &a2, &b2, p2})
				if a != a2 || b != b2 || out != out2 {
					t.Fatalf("%s n=%d alias=%v: got %d,%d,%d want %d,%d,%d", name, n, alias, a, b, out, a2, b2, out2)
				}
			}
		}
		out := 7
		fast.RunArgs([]Value{1, (*int)(nil), (*int)(nil), &out})
		if out != 8 {
			t.Fatal(out)
		}
		func() {
			defer func() {
				if recover() == nil {
					t.Error("missing nil panic")
				}
			}()
			fast.RunArgs([]Value{3, (*int)(nil), new(int), new(int)})
		}()
	}
}

func TestTypedAccumulatorMemoRejectsEffects(t *testing.T) {
	base := `package main
func recur(x int, a,b,p *int) {
 if x<2 { *p += x; return }
 BODY
 recur(x-1,a,b,p)
 recur(x-2,a,b,p)
}`
	for _, body := range []string{"*a += 1", "*p *= 2", "*p += *p", "if *p > 0 { x-- }", "recur(x-1,p,b,p)", "if a == p { x-- }"} {
		pkg, err := NewContext(0).LoadFile("reject.go", strings.Replace(base, "BODY", body, 1))
		if err != nil {
			t.Fatal(err)
		}
		r := MakeTypedRunnerOf[int](pkg.Func("recur"))
		if r != nil && r.(*typedArgsRunner[int]).accumulatorMemo {
			t.Fatalf("accepted %s", body)
		}
	}
}

func BenchmarkTypedAccumulatorMemo(b *testing.B) {
	source, err := os.ReadFile("testdata/fibish/uglyfib.go")
	if err != nil {
		b.Fatal(err)
	}
	pkg, err := NewContext(0).LoadFile("ugly.go", source)
	if err != nil {
		b.Fatal(err)
	}
	for _, fast := range []bool{false, true} {
		name := "baseline"
		if fast {
			name = "memo"
		}
		b.Run(name, func(b *testing.B) {
			var r TypedArgumentsRunner
			if fast {
				r = MakeTypedRunnerOf[int](pkg.Func("f")).(TypedArgumentsRunner)
			} else {
				r = makeTypedArgumentsRunnerOf[int](pkg.Func("f")).(TypedArgumentsRunner)
				base := r.(*typedArgsRunner[int])
				fuseTypedArguments(base)
				fuseTypedArrayCopies(base)
				batchTypedConstants(base)
			}
			a, c, out := 19, 18, 0
			args := []Value{20, &a, &c, &out}
			b.ReportAllocs()
			b.ResetTimer()
			for b.Loop() {
				out = 0
				r.RunArgs(args)
				if out != 6765 {
					b.Fatal(out)
				}
			}
		})
	}
}
