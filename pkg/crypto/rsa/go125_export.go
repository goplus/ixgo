// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package rsa

import (
	q "crypto/rsa"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/rsa", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "rsa",
			Path: "crypto/rsa",
			Deps: map[string]string{
				"crypto":                         "crypto",
				"crypto/internal/boring":         "boring",
				"crypto/internal/boring/bbig":    "bbig",
				"crypto/internal/fips140/bigmod": "bigmod",
				"crypto/internal/fips140/rsa":    "rsa",
				"crypto/internal/fips140hash":    "fips140hash",
				"crypto/internal/fips140only":    "fips140only",
				"crypto/internal/randutil":       "randutil",
				"crypto/rand":                    "rand",
				"crypto/subtle":                  "subtle",
				"errors":                         "errors",
				"fmt":                            "fmt",
				"hash":                           "hash",
				"internal/godebug":               "godebug",
				"io":                             "io",
				"math":                           "math",
				"math/big":                       "big",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"CRTValue":               reflect.TypeOf((*q.CRTValue)(nil)).Elem(),
				"OAEPOptions":            reflect.TypeOf((*q.OAEPOptions)(nil)).Elem(),
				"PKCS1v15DecryptOptions": reflect.TypeOf((*q.PKCS1v15DecryptOptions)(nil)).Elem(),
				"PSSOptions":             reflect.TypeOf((*q.PSSOptions)(nil)).Elem(),
				"PrecomputedValues":      reflect.TypeOf((*q.PrecomputedValues)(nil)).Elem(),
				"PrivateKey":             reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
				"PublicKey":              reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrDecryption":     &q.ErrDecryption,
				"ErrMessageTooLong": &q.ErrMessageTooLong,
				"ErrVerification":   &q.ErrVerification,
			},
			Funcs: map[string]interface{}{
				"DecryptOAEP":               q.DecryptOAEP,
				"DecryptPKCS1v15":           q.DecryptPKCS1v15,
				"DecryptPKCS1v15SessionKey": q.DecryptPKCS1v15SessionKey,
				"EncryptOAEP":               q.EncryptOAEP,
				"EncryptPKCS1v15":           q.EncryptPKCS1v15,
				"GenerateKey":               q.GenerateKey,
				"GenerateMultiPrimeKey":     q.GenerateMultiPrimeKey,
				"SignPKCS1v15":              q.SignPKCS1v15,
				"SignPSS":                   q.SignPSS,
				"VerifyPKCS1v15":            q.VerifyPKCS1v15,
				"VerifyPSS":                 q.VerifyPSS,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"PSSSaltLengthAuto":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.PSSSaltLengthAuto))},
				"PSSSaltLengthEqualsHash": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.PSSSaltLengthEqualsHash))},
			},
		}
	})
}
