// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package mlkem

import (
	q "crypto/mlkem"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "mlkem",
		Path: "crypto/mlkem",
		Deps: map[string]string{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey1024":         reflect.ValueOf(q.GenerateKey1024),
			"GenerateKey768":          reflect.ValueOf(q.GenerateKey768),
			"NewDecapsulationKey1024": reflect.ValueOf(q.NewDecapsulationKey1024),
			"NewDecapsulationKey768":  reflect.ValueOf(q.NewDecapsulationKey768),
			"NewEncapsulationKey1024": reflect.ValueOf(q.NewEncapsulationKey1024),
			"NewEncapsulationKey768":  reflect.ValueOf(q.NewEncapsulationKey768),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"CiphertextSize1024":       {"untyped int", constant.MakeInt64(int64(q.CiphertextSize1024))},
			"CiphertextSize768":        {"untyped int", constant.MakeInt64(int64(q.CiphertextSize768))},
			"EncapsulationKeySize1024": {"untyped int", constant.MakeInt64(int64(q.EncapsulationKeySize1024))},
			"EncapsulationKeySize768":  {"untyped int", constant.MakeInt64(int64(q.EncapsulationKeySize768))},
			"SeedSize":                 {"untyped int", constant.MakeInt64(int64(q.SeedSize))},
			"SharedKeySize":            {"untyped int", constant.MakeInt64(int64(q.SharedKeySize))},
		},
	})
}
