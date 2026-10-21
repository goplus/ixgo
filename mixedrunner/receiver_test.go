package mixedrunner

import (
	"go/types"
	"reflect"
	"testing"

	"github.com/goplus/ixgo"
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
