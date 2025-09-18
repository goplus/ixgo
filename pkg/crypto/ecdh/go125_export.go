// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package ecdh

import (
	q "crypto/ecdh"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "ecdh",
		Path: "crypto/ecdh",
		Deps: map[string]string{
			"bytes":                        "bytes",
			"crypto":                       "crypto",
			"crypto/internal/boring":       "boring",
			"crypto/internal/fips140/ecdh": "ecdh",
			"crypto/internal/fips140/edwards25519/field": "field",
			"crypto/internal/fips140only":                "fips140only",
			"crypto/internal/randutil":                   "randutil",
			"crypto/subtle":                              "subtle",
			"errors":                                     "errors",
			"io":                                         "io",
		},
		Interfaces: map[string]reflect.Type{
			"Curve": reflect.TypeOf((*q.Curve)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
			"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"P256":   reflect.ValueOf(q.P256),
			"P384":   reflect.ValueOf(q.P384),
			"P521":   reflect.ValueOf(q.P521),
			"X25519": reflect.ValueOf(q.X25519),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
