// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package ecdsa

import (
	q "crypto/ecdsa"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/ecdsa", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "ecdsa",
			Path: "crypto/ecdsa",
			Deps: map[string]string{
				"crypto":                                "crypto",
				"crypto/ecdh":                           "ecdh",
				"crypto/elliptic":                       "elliptic",
				"crypto/internal/boring":                "boring",
				"crypto/internal/boring/bbig":           "bbig",
				"crypto/internal/fips140/ecdsa":         "ecdsa",
				"crypto/internal/fips140hash":           "fips140hash",
				"crypto/internal/fips140only":           "fips140only",
				"crypto/internal/randutil":              "randutil",
				"crypto/sha512":                         "sha512",
				"crypto/subtle":                         "subtle",
				"errors":                                "errors",
				"io":                                    "io",
				"math/big":                              "big",
				"math/rand/v2":                          "rand",
				"vendor/golang.org/x/crypto/cryptobyte": "cryptobyte",
				"vendor/golang.org/x/crypto/cryptobyte/asn1": "asn1",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
				"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"GenerateKey": q.GenerateKey,
				"Sign":        q.Sign,
				"SignASN1":    q.SignASN1,
				"Verify":      q.Verify,
				"VerifyASN1":  q.VerifyASN1,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
