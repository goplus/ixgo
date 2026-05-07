// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package mlkem

import (
	q "crypto/mlkem"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/mlkem", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "mlkem",
			Path: "crypto/mlkem",
			Deps: map[string]string{
				"crypto":                        "crypto",
				"crypto/internal/fips140/mlkem": "mlkem",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"DecapsulationKey1024": reflect.TypeOf((*q.DecapsulationKey1024)(nil)).Elem(),
				"DecapsulationKey768":  reflect.TypeOf((*q.DecapsulationKey768)(nil)).Elem(),
				"EncapsulationKey1024": reflect.TypeOf((*q.EncapsulationKey1024)(nil)).Elem(),
				"EncapsulationKey768":  reflect.TypeOf((*q.EncapsulationKey768)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"GenerateKey1024":         q.GenerateKey1024,
				"GenerateKey768":          q.GenerateKey768,
				"NewDecapsulationKey1024": q.NewDecapsulationKey1024,
				"NewDecapsulationKey768":  q.NewDecapsulationKey768,
				"NewEncapsulationKey1024": q.NewEncapsulationKey1024,
				"NewEncapsulationKey768":  q.NewEncapsulationKey768,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"CiphertextSize1024":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.CiphertextSize1024))},
				"CiphertextSize768":        {Typ: "untyped int", Value: constant.MakeInt64(int64(q.CiphertextSize768))},
				"EncapsulationKeySize1024": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.EncapsulationKeySize1024))},
				"EncapsulationKeySize768":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.EncapsulationKeySize768))},
				"SeedSize":                 {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SeedSize))},
				"SharedKeySize":            {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SharedKeySize))},
			},
		}
	})
}
