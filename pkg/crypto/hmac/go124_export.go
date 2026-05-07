// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package hmac

import (
	q "crypto/hmac"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/hmac", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "hmac",
			Path: "crypto/hmac",
			Deps: map[string]string{
				"crypto/internal/boring":       "boring",
				"crypto/internal/fips140/hmac": "hmac",
				"crypto/internal/fips140hash":  "fips140hash",
				"crypto/internal/fips140only":  "fips140only",
				"crypto/subtle":                "subtle",
				"hash":                         "hash",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Equal": q.Equal,
				"New":   q.New,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
