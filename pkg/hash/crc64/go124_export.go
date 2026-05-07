// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package crc64

import (
	q "hash/crc64"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("hash/crc64", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "crc64",
			Path: "hash/crc64",
			Deps: map[string]string{
				"errors":             "errors",
				"hash":               "hash",
				"internal/byteorder": "byteorder",
				"sync":               "sync",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Table": reflect.TypeOf((*q.Table)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Checksum":  q.Checksum,
				"MakeTable": q.MakeTable,
				"New":       q.New,
				"Update":    q.Update,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"ECMA": {Typ: "untyped int", Value: constant.MakeUint64(uint64(q.ECMA))},
				"ISO":  {Typ: "untyped int", Value: constant.MakeUint64(uint64(q.ISO))},
				"Size": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size))},
			},
		}
	})
}
