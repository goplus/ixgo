// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package fmt

import (
	q "fmt"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("fmt", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Append":       q.Append,
				"Appendf":      q.Appendf,
				"Appendln":     q.Appendln,
				"Errorf":       q.Errorf,
				"FormatString": q.FormatString,
				"Fprint":       q.Fprint,
				"Fprintf":      q.Fprintf,
				"Fprintln":     q.Fprintln,
				"Fscan":        q.Fscan,
				"Fscanf":       q.Fscanf,
				"Fscanln":      q.Fscanln,
				"Print":        q.Print,
				"Printf":       q.Printf,
				"Println":      q.Println,
				"Scan":         q.Scan,
				"Scanf":        q.Scanf,
				"Scanln":       q.Scanln,
				"Sprint":       q.Sprint,
				"Sprintf":      q.Sprintf,
				"Sprintln":     q.Sprintln,
				"Sscan":        q.Sscan,
				"Sscanf":       q.Sscanf,
				"Sscanln":      q.Sscanln,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
