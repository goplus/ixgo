// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package crc32

import (
	q "hash/crc32"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("hash/crc32", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "crc32",
			Path: "hash/crc32",
			Deps: map[string]string{
				"errors":             "errors",
				"hash":               "hash",
				"internal/byteorder": "byteorder",
				"internal/cpu":       "cpu",
				"sync":               "sync",
				"sync/atomic":        "atomic",
				"unsafe":             "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Table": reflect.TypeOf((*q.Table)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"IEEETable": &q.IEEETable,
			},
			Funcs: map[string]interface{}{
				"Checksum":     q.Checksum,
				"ChecksumIEEE": q.ChecksumIEEE,
				"MakeTable":    q.MakeTable,
				"New":          q.New,
				"NewIEEE":      q.NewIEEE,
				"Update":       q.Update,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Castagnoli": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Castagnoli))},
				"IEEE":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.IEEE))},
				"Koopman":    {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Koopman))},
				"Size":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size))},
			},
		}
	})
}
