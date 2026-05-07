// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package subtle

import (
	q "crypto/subtle"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/subtle", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "subtle",
			Path: "crypto/subtle",
			Deps: map[string]string{
				"crypto/internal/constanttime":   "constanttime",
				"crypto/internal/fips140/subtle": "subtle",
				"internal/runtime/sys":           "sys",
				"unsafe":                         "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"ConstantTimeByteEq":        q.ConstantTimeByteEq,
				"ConstantTimeCompare":       q.ConstantTimeCompare,
				"ConstantTimeCopy":          q.ConstantTimeCopy,
				"ConstantTimeEq":            q.ConstantTimeEq,
				"ConstantTimeLessOrEq":      q.ConstantTimeLessOrEq,
				"ConstantTimeSelect":        q.ConstantTimeSelect,
				"WithDataIndependentTiming": q.WithDataIndependentTiming,
				"XORBytes":                  q.XORBytes,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
