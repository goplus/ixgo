// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package plugin

import (
	q "plugin"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "plugin",
		Path: "plugin",
		Deps: map[string]string{
			"errors":      "errors",
			"runtime/cgo": "cgo",
			"sync":        "sync",
			"syscall":     "syscall",
			"unsafe":      "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Symbol": reflect.TypeOf((*q.Symbol)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Plugin": reflect.TypeOf((*q.Plugin)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Open": reflect.ValueOf(q.Open),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
