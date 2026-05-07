// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25 && js
// +build go1.24,!go1.25,js

package js

import (
	q "syscall/js"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("syscall/js", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "js",
			Path: "syscall/js",
			Deps: map[string]string{
				"internal/synctest": "synctest",
				"runtime":           "runtime",
				"sync":              "sync",
				"unsafe":            "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Error":      reflect.TypeOf((*q.Error)(nil)).Elem(),
				"Func":       reflect.TypeOf((*q.Func)(nil)).Elem(),
				"Type":       reflect.TypeOf((*q.Type)(nil)).Elem(),
				"Value":      reflect.TypeOf((*q.Value)(nil)).Elem(),
				"ValueError": reflect.TypeOf((*q.ValueError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"CopyBytesToGo": q.CopyBytesToGo,
				"CopyBytesToJS": q.CopyBytesToJS,
				"FuncOf":        q.FuncOf,
				"Global":        q.Global,
				"Null":          q.Null,
				"Undefined":     q.Undefined,
				"ValueOf":       q.ValueOf,
			},
			TypedConsts: map[string]interface{}{
				"TypeBoolean":   q.TypeBoolean,
				"TypeFunction":  q.TypeFunction,
				"TypeNull":      q.TypeNull,
				"TypeNumber":    q.TypeNumber,
				"TypeObject":    q.TypeObject,
				"TypeString":    q.TypeString,
				"TypeSymbol":    q.TypeSymbol,
				"TypeUndefined": q.TypeUndefined,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
