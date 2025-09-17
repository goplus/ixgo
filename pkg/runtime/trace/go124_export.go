// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package trace

import (
	q "runtime/trace"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "trace",
		Path: "runtime/trace",
		Deps: map[string]string{
			"context":     "context",
			"fmt":         "fmt",
			"io":          "io",
			"runtime":     "runtime",
			"sync":        "sync",
			"sync/atomic": "atomic",
			"unsafe":      "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Region": reflect.TypeOf((*q.Region)(nil)).Elem(),
			"Task":   reflect.TypeOf((*q.Task)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"IsEnabled":   reflect.ValueOf(q.IsEnabled),
			"Log":         reflect.ValueOf(q.Log),
			"Logf":        reflect.ValueOf(q.Logf),
			"NewTask":     reflect.ValueOf(q.NewTask),
			"Start":       reflect.ValueOf(q.Start),
			"StartRegion": reflect.ValueOf(q.StartRegion),
			"Stop":        reflect.ValueOf(q.Stop),
			"WithRegion":  reflect.ValueOf(q.WithRegion),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
