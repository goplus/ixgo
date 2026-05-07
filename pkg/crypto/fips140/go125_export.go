// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package fips140

import (
	q "crypto/fips140"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/fips140", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "fips140",
			Path: "crypto/fips140",
			Deps: map[string]string{
				"crypto/internal/fips140":       "fips140",
				"crypto/internal/fips140/check": "check",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Enabled": q.Enabled,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
