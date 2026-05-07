// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package expvar

import (
	q "expvar"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("expvar", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "expvar",
			Path: "expvar",
			Deps: map[string]string{
				"encoding/json":    "json",
				"internal/godebug": "godebug",
				"log":              "log",
				"math":             "math",
				"net/http":         "http",
				"os":               "os",
				"runtime":          "runtime",
				"slices":           "slices",
				"strconv":          "strconv",
				"sync":             "sync",
				"sync/atomic":      "atomic",
				"unicode/utf8":     "utf8",
			},
			Interfaces: map[string]reflect.Type{
				"Var": reflect.TypeOf((*q.Var)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Float":    reflect.TypeOf((*q.Float)(nil)).Elem(),
				"Func":     reflect.TypeOf((*q.Func)(nil)).Elem(),
				"Int":      reflect.TypeOf((*q.Int)(nil)).Elem(),
				"KeyValue": reflect.TypeOf((*q.KeyValue)(nil)).Elem(),
				"Map":      reflect.TypeOf((*q.Map)(nil)).Elem(),
				"String":   reflect.TypeOf((*q.String)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Do":        q.Do,
				"Get":       q.Get,
				"Handler":   q.Handler,
				"NewFloat":  q.NewFloat,
				"NewInt":    q.NewInt,
				"NewMap":    q.NewMap,
				"NewString": q.NewString,
				"Publish":   q.Publish,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
