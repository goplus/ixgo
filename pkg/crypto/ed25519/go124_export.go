// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package ed25519

import (
	q "crypto/ed25519"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/ed25519", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "ed25519",
			Path: "crypto/ed25519",
			Deps: map[string]string{
				"crypto":                          "crypto",
				"crypto/internal/fips140/ed25519": "ed25519",
				"crypto/internal/fips140only":     "fips140only",
				"crypto/rand":                     "rand",
				"crypto/subtle":                   "subtle",
				"errors":                          "errors",
				"io":                              "io",
				"strconv":                         "strconv",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Options":    reflect.TypeOf((*q.Options)(nil)).Elem(),
				"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
				"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"GenerateKey":       q.GenerateKey,
				"NewKeyFromSeed":    q.NewKeyFromSeed,
				"Sign":              q.Sign,
				"Verify":            q.Verify,
				"VerifyWithOptions": q.VerifyWithOptions,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"PrivateKeySize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.PrivateKeySize))},
				"PublicKeySize":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.PublicKeySize))},
				"SeedSize":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SeedSize))},
				"SignatureSize":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SignatureSize))},
			},
		}
	})
}
