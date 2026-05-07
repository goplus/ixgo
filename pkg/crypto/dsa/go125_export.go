// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package dsa

import (
	q "crypto/dsa"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/dsa", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "dsa",
			Path: "crypto/dsa",
			Deps: map[string]string{
				"crypto/internal/fips140only": "fips140only",
				"crypto/internal/randutil":    "randutil",
				"errors":                      "errors",
				"io":                          "io",
				"math/big":                    "big",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"ParameterSizes": reflect.TypeOf((*q.ParameterSizes)(nil)).Elem(),
				"Parameters":     reflect.TypeOf((*q.Parameters)(nil)).Elem(),
				"PrivateKey":     reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
				"PublicKey":      reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrInvalidPublicKey": &q.ErrInvalidPublicKey,
			},
			Funcs: map[string]interface{}{
				"GenerateKey":        q.GenerateKey,
				"GenerateParameters": q.GenerateParameters,
				"Sign":               q.Sign,
				"Verify":             q.Verify,
			},
			TypedConsts: map[string]interface{}{
				"L1024N160": q.L1024N160,
				"L2048N224": q.L2048N224,
				"L2048N256": q.L2048N256,
				"L3072N256": q.L3072N256,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
