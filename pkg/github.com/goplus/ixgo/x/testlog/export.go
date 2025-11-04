// export by github.com/goplus/ixgo/cmd/qexp

package testlog

import (
	q "github.com/goplus/ixgo/x/testlog"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Getenv":          reflect.ValueOf(q.Getenv),
			"Logger":          reflect.ValueOf(q.Logger),
			"Open":            reflect.ValueOf(q.Open),
			"PanicOnExit0":    reflect.ValueOf(q.PanicOnExit0),
			"SetLogger":       reflect.ValueOf(q.SetLogger),
			"SetPanicOnExit0": reflect.ValueOf(q.SetPanicOnExit0),
			"Stat":            reflect.ValueOf(q.Stat),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
