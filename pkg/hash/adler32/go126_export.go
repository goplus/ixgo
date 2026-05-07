// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package adler32

import (
	q "hash/adler32"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("hash/adler32", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "adler32",
			Path: "hash/adler32",
			Deps: map[string]string{
				"errors":             "errors",
				"hash":               "hash",
				"internal/byteorder": "byteorder",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Checksum": q.Checksum,
				"New":      q.New,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Size": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size))},
			},
		}
	})
}
