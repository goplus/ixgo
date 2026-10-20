package mixedrunner

import (
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/goplus/ixgo"
	"golang.org/x/tools/go/ssa"
)

func TestMixedSelection(t *testing.T) {
	const source = `package main
func sum(n int,step float64) float64 {s:=0.0;for i:=0;i<n;i++ {s+=step+float64(i)};return s}
func call(n int,step float64) float64 {return sum(n,step)}
func unsupported(n int,step float64) float64 {a:=[2]float64{step,2};return a[n]}
`
	ctx := ixgo.NewContext(0)
	pkg, err := ctx.LoadFile("mixed.go", source)
	if err != nil {
		t.Fatal(err)
	}
	r, ok := MakeRunner(pkg.Func("sum")).(ixgo.TypedArgumentsRunner)
	if !ok || r.RunArgs([]ixgo.Value{5, 0.5}) != 12.5 {
		t.Fatal("mixed fallback not selected")
	}
	r, ok = MakeRunner(pkg.Func("call")).(ixgo.TypedArgumentsRunner)
	if !ok || r.RunArgs([]ixgo.Value{5, 0.5}) != 12.5 {
		t.Fatal("mixed call graph not selected")
	}
	if MakeRunner(pkg.Func("unsupported")) != nil {
		t.Fatal("unsupported memory must fall back")
	}
	interp, err := ixgo.NewContext(0).LoadInterp("mixed.go", source)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	for _, tc := range []struct {
		name string
		n    int
		want float64
	}{{"sum", 5, 12.5}, {"call", 5, 12.5}, {"unsupported", 0, 0.5}} {
		got, err := interp.RunFunc(tc.name, tc.n, 0.5)
		if err != nil || got != tc.want {
			t.Fatalf("%s: %v %v", tc.name, got, err)
		}
	}
}

func TestDivmodMixedSelection(t *testing.T) {
	path := filepath.Join(runtime.GOROOT(), "test", "divmod.go")
	source, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		t.Skip("Go distribution test sources unavailable")
	}
	if err != nil {
		t.Fatal(err)
	}
	pkg, err := ixgo.NewContext(0).LoadFile(path, source)
	if err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{"udiv", "idiv"} {
		fn := pkg.Func(name)
		if fn == nil {
			t.Fatalf("%s: expected mixed graph", name)
		}
		r := MakeRunner(fn)
		if r == nil || reflect.TypeOf(r) != reflect.TypeOf(MakeRunner(fn)) {
			t.Fatalf("%s: mixed runner not selected", name)
		}
	}
	r := MakeRunner(pkg.Func("idiv")).(ixgo.TypedArgumentsRunner)
	if got := r.RunArgs([]ixgo.Value{int64(-17), int64(3)}); !reflect.DeepEqual(got, ixgo.Tuple{int64(-5), int64(-2)}) {
		t.Fatalf("idiv: %v", got)
	}
}

func TestStandaloneInstallation(t *testing.T) {
	selected := make(map[string]bool)
	ixgo.RegisterMixedRunner(func(fn *ssa.Function) ixgo.TypedRunner {
		selected[fn.Name()] = true
		return MakeRunner(fn)
	})
	t.Cleanup(func() { ixgo.RegisterMixedRunner(MakeRunner) })
	i, err := ixgo.NewContext(0).LoadInterp("standalone.go", `package main
func plain(n int) int {return n+1}
func integer(n int,u uint64) int {return n+int(u)}
func floating(x float32,y float64) float64 {return float64(x)+y}
`)
	if err != nil {
		t.Fatal(err)
	}
	defer i.UnsafeRelease()
	for _, tc := range []struct {
		name string
		args []ixgo.Value
		want ixgo.Value
	}{
		{"plain", []ixgo.Value{3}, 4},
		{"integer", []ixgo.Value{3, uint64(4)}, 7},
		{"floating", []ixgo.Value{float32(1.5), 2.5}, 4.0},
	} {
		if !selected[tc.name] {
			t.Fatalf("%s did not reach standalone mixed factory", tc.name)
		}
		if got, err := i.RunFunc(tc.name, tc.args...); err != nil || got != tc.want {
			t.Fatalf("%s: %v %v", tc.name, got, err)
		}
	}
}
