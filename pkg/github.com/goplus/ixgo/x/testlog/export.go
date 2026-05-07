// export by github.com/goplus/ixgo/cmd/qexp

package testlog

import (
	q "github.com/goplus/ixgo/x/testlog"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/ixgo/x/testlog", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "testlog",
			Path: "github.com/goplus/ixgo/x/testlog",
			Deps: map[string]string{
				"sync":        "sync",
				"sync/atomic": "atomic",
			},
			Interfaces: map[string]reflect.Type{
				"Interface": reflect.TypeOf((*q.Interface)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Getenv":          q.Getenv,
				"Logger":          q.Logger,
				"Open":            q.Open,
				"PanicOnExit0":    q.PanicOnExit0,
				"SetLogger":       q.SetLogger,
				"SetPanicOnExit0": q.SetPanicOnExit0,
				"Stat":            q.Stat,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
