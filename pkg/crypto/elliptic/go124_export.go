// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package elliptic

import (
	q "crypto/elliptic"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/elliptic", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"GenerateKey":         q.GenerateKey,
				"Marshal":             q.Marshal,
				"MarshalCompressed":   q.MarshalCompressed,
				"P224":                q.P224,
				"P256":                q.P256,
				"P384":                q.P384,
				"P521":                q.P521,
				"Unmarshal":           q.Unmarshal,
				"UnmarshalCompressed": q.UnmarshalCompressed,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
