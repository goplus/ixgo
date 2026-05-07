// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package ecdh

import (
	q "crypto/ecdh"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/ecdh", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "ecdh",
			Path: "crypto/ecdh",
			Deps: map[string]string{
				"bytes":                        "bytes",
				"crypto":                       "crypto",
				"crypto/internal/boring":       "boring",
				"crypto/internal/fips140/ecdh": "ecdh",
				"crypto/internal/fips140/edwards25519/field": "field",
				"crypto/internal/fips140only":                "fips140only",
				"crypto/internal/rand":                       "rand",
				"crypto/subtle":                              "subtle",
				"errors":                                     "errors",
				"io":                                         "io",
			},
			Interfaces: map[string]reflect.Type{
				"Curve":        reflect.TypeOf((*q.Curve)(nil)).Elem(),
				"KeyExchanger": reflect.TypeOf((*q.KeyExchanger)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
				"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"P256":   q.P256,
				"P384":   q.P384,
				"P521":   q.P521,
				"X25519": q.X25519,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
