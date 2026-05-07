// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package rc4

import (
	q "crypto/rc4"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/rc4", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "rc4",
			Path: "crypto/rc4",
			Deps: map[string]string{
				"crypto/internal/fips140/alias": "alias",
				"crypto/internal/fips140only":   "fips140only",
				"errors":                        "errors",
				"strconv":                       "strconv",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Cipher":       reflect.TypeOf((*q.Cipher)(nil)).Elem(),
				"KeySizeError": reflect.TypeOf((*q.KeySizeError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewCipher": q.NewCipher,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
