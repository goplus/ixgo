// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package elliptic

import (
	q "crypto/elliptic"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "elliptic",
		Path: "crypto/elliptic",
		Deps: map[string]string{
			"crypto/internal/fips140/nistec": "nistec",
			"errors":                         "errors",
			"io":                             "io",
			"math/big":                       "big",
			"sync":                           "sync",
		},
		Interfaces: map[string]reflect.Type{
			"Curve": reflect.TypeOf((*q.Curve)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"CurveParams": reflect.TypeOf((*q.CurveParams)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey":         reflect.ValueOf(q.GenerateKey),
			"Marshal":             reflect.ValueOf(q.Marshal),
			"MarshalCompressed":   reflect.ValueOf(q.MarshalCompressed),
			"P224":                reflect.ValueOf(q.P224),
			"P256":                reflect.ValueOf(q.P256),
			"P384":                reflect.ValueOf(q.P384),
			"P521":                reflect.ValueOf(q.P521),
			"Unmarshal":           reflect.ValueOf(q.Unmarshal),
			"UnmarshalCompressed": reflect.ValueOf(q.UnmarshalCompressed),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
