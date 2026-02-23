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
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"ErrDecryption":     reflect.ValueOf(&q.ErrDecryption),
			"ErrMessageTooLong": reflect.ValueOf(&q.ErrMessageTooLong),
			"ErrVerification":   reflect.ValueOf(&q.ErrVerification),
		},
		Funcs: map[string]reflect.Value{
			"DecryptOAEP":               reflect.ValueOf(q.DecryptOAEP),
			"DecryptPKCS1v15":           reflect.ValueOf(q.DecryptPKCS1v15),
			"DecryptPKCS1v15SessionKey": reflect.ValueOf(q.DecryptPKCS1v15SessionKey),
			"EncryptOAEP":               reflect.ValueOf(q.EncryptOAEP),
			"EncryptPKCS1v15":           reflect.ValueOf(q.EncryptPKCS1v15),
			"GenerateKey":               reflect.ValueOf(q.GenerateKey),
			"GenerateMultiPrimeKey":     reflect.ValueOf(q.GenerateMultiPrimeKey),
			"SignPKCS1v15":              reflect.ValueOf(q.SignPKCS1v15),
			"SignPSS":                   reflect.ValueOf(q.SignPSS),
			"VerifyPKCS1v15":            reflect.ValueOf(q.VerifyPKCS1v15),
			"VerifyPSS":                 reflect.ValueOf(q.VerifyPSS),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"PSSSaltLengthAuto":       {"untyped int", constant.MakeInt64(int64(q.PSSSaltLengthAuto))},
			"PSSSaltLengthEqualsHash": {"untyped int", constant.MakeInt64(int64(q.PSSSaltLengthEqualsHash))},
		},
	})
}
