// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

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
				"internal/abi":          "abi",
				"internal/byteorder":    "byteorder",
				"internal/goarch":       "goarch",
				"internal/goexperiment": "goexperiment",
				"math":                  "math",
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
