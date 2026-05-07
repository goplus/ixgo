// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package des

import (
	q "crypto/des"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/des", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "des",
			Path: "crypto/des",
			Deps: map[string]string{
				"crypto/cipher":                 "cipher",
				"crypto/internal/fips140/alias": "alias",
				"crypto/internal/fips140only":   "fips140only",
				"errors":                        "errors",
				"internal/byteorder":            "byteorder",
				"strconv":                       "strconv",
				"sync":                          "sync",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"KeySizeError": reflect.TypeOf((*q.KeySizeError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewCipher":          q.NewCipher,
				"NewTripleDESCipher": q.NewTripleDESCipher,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"BlockSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.BlockSize))},
			},
		}
	})
}
