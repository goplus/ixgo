package ixgo

import (
	"testing"

	"golang.org/x/tools/go/ssa"
)

func TestMapNilInterfaces(t *testing.T) {
	cases := map[string]string{
		"constant_value": `func main() {
			m := map[string]any{"x": 1}
			m["x"] = nil
			if v, ok := m["x"]; !ok || v != nil || len(m) != 1 { panic("nil assignment lost entry") }
		}`,
		"dynamic_value": `func assign(m map[string]any, v any) { m["x"] = v }
		func main() {
			m := map[string]any{"x": 1}
			assign(m, nil)
			if v, ok := m["x"]; !ok || v != nil || len(m) != 1 { panic("nil assignment lost entry") }
		}`,
		"nil_map_constant_value": `func main() {
			defer func() {
				if err, ok := recover().(error); !ok || err.Error() != "assignment to entry in nil map" {
					panic("expected assignment to entry in nil map")
				}
			}()
			var m map[string]any
			m["x"] = nil
		}`,
		"nil_map_dynamic_value": `func assign(m map[string]any, v any) { m["x"] = v }
		func main() {
			defer func() {
				if err, ok := recover().(error); !ok || err.Error() != "assignment to entry in nil map" {
					panic("expected assignment to entry in nil map")
				}
			}()
			assign(nil, nil)
		}`,
		"nil_key": `func main() {
			m := map[any]any{nil: nil}
			if v, ok := m[nil]; !ok || v != nil { panic("nil key not present") }
			m[nil] = 7
			if m[nil] != 7 || len(m) != 1 { panic("nil key not updated") }
			delete(m, nil)
			if _, ok := m[nil]; ok || len(m) != 0 { panic("nil key not deleted") }
		}`,
		"dynamic_nil_key": `func check(k any) {
			m := make(map[any]int)
			m[k] = 7
			if v, ok := m[k]; !ok || v != 7 { panic("dynamic nil key not present") }
			delete(m, k)
			if m[k] != 0 || len(m) != 0 { panic("dynamic nil key not deleted") }
		}
		func main() { check(nil) }`,
		"nil_map_lookup_and_delete": `func main() {
			var m map[any]int
			if v, ok := m[nil]; ok || v != 0 { panic("unexpected nil map lookup") }
			delete(m, nil)
		}`,
		"deferred_delete": `func remove(m map[any]int) { defer delete(m, nil) }
		func main() {
			m := map[any]int{nil: 7}
			remove(m)
			if len(m) != 0 { panic("deferred delete failed") }
		}`,
		"typed_nil": `func main() {
			var p *int
			m := map[any]any{nil: nil, p: p}
			if len(m) != 2 || m[p] != p || m[p] == nil { panic("typed nil lost its type") }
			delete(m, nil)
			if len(m) != 1 || m[p] != p { panic("typed nil key confused with nil") }
		}`,
	}
	for name, body := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := NewContext(0).RunFile("main.go", "package main\n"+body, nil); err != nil {
				t.Fatal(err)
			}
		})
	}
}

// Exercise both generic builtin dispatchers in addition to the compiled
// instruction and deferred-call paths covered above.
func TestDeleteNilKeyBuiltinDispatch(t *testing.T) {
	ctx := NewContext(0)
	pkg, err := ctx.LoadFile("main.go", `package main
func remove(m map[any]int) { delete(m, nil) }
func main() {}`)
	if err != nil {
		t.Fatal(err)
	}
	var builtin *ssa.Builtin
findDelete:
	for _, block := range pkg.Func("remove").Blocks {
		for _, instr := range block.Instrs {
			if call, ok := instr.(*ssa.Call); ok {
				if candidate, ok := call.Common().Value.(*ssa.Builtin); ok && candidate.Name() == "delete" {
					builtin = candidate
					break findDelete
				}
			}
		}
	}
	if builtin == nil {
		t.Fatal("delete builtin not found")
	}
	// delete uses only the evaluated args, so an empty interpreter and nil
	// frame and SSA args suffice.
	inter := &Interp{}
	for _, discard := range []bool{false, true} {
		m := map[any]int{nil: 1, "keep": 2}
		args := []value{m, nil}
		if discard {
			inter.callBuiltinDiscardsResult(nil, builtin, args, nil)
		} else {
			inter.callBuiltin(nil, builtin, args, nil)
		}
		if len(m) != 1 || m["keep"] != 2 {
			t.Fatalf("discard=%v: delete result = %v", discard, m)
		}
	}
}
