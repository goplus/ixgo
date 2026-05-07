// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package sync

import (
	q "sync"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("sync", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "sync",
			Path: "sync",
			Deps: map[string]string{
				"internal/race":     "race",
				"internal/sync":     "sync",
				"internal/synctest": "synctest",
				"runtime":           "runtime",
				"sync/atomic":       "atomic",
				"unsafe":            "unsafe",
			},
			Interfaces: map[string]reflect.Type{
				"Locker": reflect.TypeOf((*q.Locker)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Cond":      reflect.TypeOf((*q.Cond)(nil)).Elem(),
				"Map":       reflect.TypeOf((*q.Map)(nil)).Elem(),
				"Mutex":     reflect.TypeOf((*q.Mutex)(nil)).Elem(),
				"Once":      reflect.TypeOf((*q.Once)(nil)).Elem(),
				"Pool":      reflect.TypeOf((*q.Pool)(nil)).Elem(),
				"RWMutex":   reflect.TypeOf((*q.RWMutex)(nil)).Elem(),
				"WaitGroup": reflect.TypeOf((*q.WaitGroup)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewCond":  q.NewCond,
				"OnceFunc": q.OnceFunc,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
