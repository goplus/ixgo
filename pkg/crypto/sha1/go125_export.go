// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package sha1

import (
	q "crypto/sha1"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/sha1", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "sha1",
			Path: "crypto/sha1",
			Deps: map[string]string{
				"crypto":                      "crypto",
				"crypto/internal/boring":      "boring",
				"crypto/internal/fips140only": "fips140only",
				"crypto/internal/impl":        "impl",
				"errors":                      "errors",
				"hash":                        "hash",
				"internal/byteorder":          "byteorder",
				"internal/cpu":                "cpu",
				"math/bits":                   "bits",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New": q.New,
				"Sum": q.Sum,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"BlockSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.BlockSize))},
				"Size":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size))},
			},
		}
	})
}
