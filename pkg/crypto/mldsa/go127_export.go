// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package mldsa

import (
	q "crypto/mldsa"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("crypto/mldsa", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "mldsa",
			Path: "crypto/mldsa",
			Deps: map[string]string{
				"crypto":                        "crypto",
				"crypto/internal/fips140/mldsa": "mldsa",
				"errors":                        "errors",
				"io":                            "io",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Options":    reflect.TypeOf((*q.Options)(nil)).Elem(),
				"Parameters": reflect.TypeOf((*q.Parameters)(nil)).Elem(),
				"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
				"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"GenerateKey":   reflect.ValueOf(q.GenerateKey),
				"MLDSA44":       reflect.ValueOf(q.MLDSA44),
				"MLDSA65":       reflect.ValueOf(q.MLDSA65),
				"MLDSA87":       reflect.ValueOf(q.MLDSA87),
				"NewPrivateKey": reflect.ValueOf(q.NewPrivateKey),
				"NewPublicKey":  reflect.ValueOf(q.NewPublicKey),
				"Verify":        reflect.ValueOf(q.Verify),
			},
			TypedConsts: map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"MLDSA44PublicKeySize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MLDSA44PublicKeySize))},
				"MLDSA44SignatureSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MLDSA44SignatureSize))},
				"MLDSA65PublicKeySize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MLDSA65PublicKeySize))},
				"MLDSA65SignatureSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MLDSA65SignatureSize))},
				"MLDSA87PublicKeySize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MLDSA87PublicKeySize))},
				"MLDSA87SignatureSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MLDSA87SignatureSize))},
				"PrivateKeySize":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.PrivateKeySize))},
			},
		}
	})
}
