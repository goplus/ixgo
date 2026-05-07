// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package sha256

import (
	q "crypto/sha256"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/sha256", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "sha256",
			Path: "crypto/sha256",
			Deps: map[string]string{
				"crypto":                         "crypto",
				"crypto/internal/boring":         "boring",
				"crypto/internal/fips140/sha256": "sha256",
				"hash":                           "hash",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New":    q.New,
				"New224": q.New224,
				"Sum224": q.Sum224,
				"Sum256": q.Sum256,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"BlockSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.BlockSize))},
				"Size":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size))},
				"Size224":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Size224))},
			},
		}
	})
}
