// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package maphash

import (
	q "hash/maphash"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "maphash",
		Path: "hash/maphash",
		Deps: map[string]string{
			"hash":                  "hash",
			"internal/abi":          "abi",
			"internal/goarch":       "goarch",
			"internal/goexperiment": "goexperiment",
			"unsafe":                "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Hash": reflect.TypeOf((*q.Hash)(nil)).Elem(),
			"Seed": reflect.TypeOf((*q.Seed)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Bytes":    reflect.ValueOf(q.Bytes),
			"MakeSeed": reflect.ValueOf(q.MakeSeed),
			"String":   reflect.ValueOf(q.String),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
