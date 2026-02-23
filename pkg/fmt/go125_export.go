// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package fmt

import (
	q "fmt"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "fmt",
		Path: "fmt",
		Deps: map[string]string{
			"errors":           "errors",
			"internal/fmtsort": "fmtsort",
			"io":               "io",
			"math":             "math",
			"os":               "os",
			"reflect":          "reflect",
			"slices":           "slices",
			"strconv":          "strconv",
			"sync":             "sync",
			"unicode/utf8":     "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Formatter":  reflect.TypeOf((*q.Formatter)(nil)).Elem(),
			"GoStringer": reflect.TypeOf((*q.GoStringer)(nil)).Elem(),
			"ScanState":  reflect.TypeOf((*q.ScanState)(nil)).Elem(),
			"Scanner":    reflect.TypeOf((*q.Scanner)(nil)).Elem(),
			"State":      reflect.TypeOf((*q.State)(nil)).Elem(),
			"Stringer":   reflect.TypeOf((*q.Stringer)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Append":       reflect.ValueOf(q.Append),
			"Appendf":      reflect.ValueOf(q.Appendf),
			"Appendln":     reflect.ValueOf(q.Appendln),
			"Errorf":       reflect.ValueOf(q.Errorf),
			"FormatString": reflect.ValueOf(q.FormatString),
			"Fprint":       reflect.ValueOf(q.Fprint),
			"Fprintf":      reflect.ValueOf(q.Fprintf),
			"Fprintln":     reflect.ValueOf(q.Fprintln),
			"Fscan":        reflect.ValueOf(q.Fscan),
			"Fscanf":       reflect.ValueOf(q.Fscanf),
			"Fscanln":      reflect.ValueOf(q.Fscanln),
			"Print":        reflect.ValueOf(q.Print),
			"Printf":       reflect.ValueOf(q.Printf),
			"Println":      reflect.ValueOf(q.Println),
			"Scan":         reflect.ValueOf(q.Scan),
			"Scanf":        reflect.ValueOf(q.Scanf),
			"Scanln":       reflect.ValueOf(q.Scanln),
			"Sprint":       reflect.ValueOf(q.Sprint),
			"Sprintf":      reflect.ValueOf(q.Sprintf),
			"Sprintln":     reflect.ValueOf(q.Sprintln),
			"Sscan":        reflect.ValueOf(q.Sscan),
			"Sscanf":       reflect.ValueOf(q.Sscanf),
			"Sscanln":      reflect.ValueOf(q.Sscanln),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
