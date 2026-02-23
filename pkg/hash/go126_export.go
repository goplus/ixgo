// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package hash

import (
	q "hash"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "hash",
		Path: "hash",
		Deps: map[string]string{
			"io": "io",
		},
		Interfaces: map[string]reflect.Type{
			"Cloner": reflect.TypeOf((*q.Cloner)(nil)).Elem(),
			"Hash":   reflect.TypeOf((*q.Hash)(nil)).Elem(),
			"Hash32": reflect.TypeOf((*q.Hash32)(nil)).Elem(),
			"Hash64": reflect.TypeOf((*q.Hash64)(nil)).Elem(),
			"XOF":    reflect.TypeOf((*q.XOF)(nil)).Elem(),
		},
		NamedTypes:    map[string]reflect.Type{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
