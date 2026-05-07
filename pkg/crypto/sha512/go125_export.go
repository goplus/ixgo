// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package sha512

import (
	q "crypto/sha512"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/sha512", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "sha512",
			Path: "crypto/sha512",
			Deps: map[string]string{
				"crypto":                         "crypto",
				"crypto/internal/boring":         "boring",
				"crypto/internal/fips140/sha512": "sha512",
				"hash":                           "hash",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New":        q.New,
				"New384":     q.New384,
				"New512_224": q.New512_224,
				"New512_256": q.New512_256,
				"Sum384":     q.Sum384,
				"Sum512":     q.Sum512,
				"Sum512_224": q.Sum512_224,
				"Sum512_256": q.Sum512_256,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"BlockSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.BlockSize))},
				"Size":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size))},
				"Size224":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size224))},
				"Size256":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size256))},
				"Size384":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size384))},
			},
		}
	})
}
