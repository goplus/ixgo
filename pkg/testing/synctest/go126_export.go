// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package synctest

import (
	q "testing/synctest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("testing/synctest", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "synctest",
			Path: "testing/synctest",
			Deps: map[string]string{
				"internal/synctest": "synctest",
				"testing":           "testing",
				"unsafe":            "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"Test": reflect.ValueOf(q.Test),
				"Wait": reflect.ValueOf(q.Wait),
			},
			TypedConsts:   map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
