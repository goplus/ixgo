// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package mlkemtest

import (
	q "crypto/mlkem/mlkemtest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "mlkemtest",
		Path: "crypto/mlkem/mlkemtest",
		Deps: map[string]string{
			"crypto/internal/fips140/mlkem": "mlkem",
			"crypto/internal/fips140only":   "fips140only",
			"crypto/mlkem":                  "mlkem",
			"errors":                        "errors",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Encapsulate1024": reflect.ValueOf(q.Encapsulate1024),
			"Encapsulate768":  reflect.ValueOf(q.Encapsulate768),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
