// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package maphash

import (
	q "hash/maphash"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("hash/maphash", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "maphash",
			Path: "hash/maphash",
			Deps: map[string]string{
				"hash":                  "hash",
				"internal/abi":          "abi",
				"internal/runtime/maps": "maps",
				"unsafe":                "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Hash": reflect.TypeOf((*q.Hash)(nil)).Elem(),
				"Seed": reflect.TypeOf((*q.Seed)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Bytes":    q.Bytes,
				"MakeSeed": q.MakeSeed,
				"String":   q.String,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
