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
	ixgo.RegisterPackage("runtime/trace", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"IsEnabled":   q.IsEnabled,
				"Log":         q.Log,
				"Logf":        q.Logf,
				"NewTask":     q.NewTask,
				"Start":       q.Start,
				"StartRegion": q.StartRegion,
				"Stop":        q.Stop,
				"WithRegion":  q.WithRegion,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
