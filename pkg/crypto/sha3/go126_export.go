// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package sha3

import (
	q "crypto/sha3"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/sha3", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "sha3",
			Path: "crypto/sha3",
			Deps: map[string]string{
				"crypto":                       "crypto",
				"crypto/internal/fips140/sha3": "sha3",
				"hash":                         "hash",
				"unsafe":                       "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"SHA3":  reflect.TypeOf((*q.SHA3)(nil)).Elem(),
				"SHAKE": reflect.TypeOf((*q.SHAKE)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New224":       q.New224,
				"New256":       q.New256,
				"New384":       q.New384,
				"New512":       q.New512,
				"NewCSHAKE128": q.NewCSHAKE128,
				"NewCSHAKE256": q.NewCSHAKE256,
				"NewSHAKE128":  q.NewSHAKE128,
				"NewSHAKE256":  q.NewSHAKE256,
				"Sum224":       q.Sum224,
				"Sum256":       q.Sum256,
				"Sum384":       q.Sum384,
				"Sum512":       q.Sum512,
				"SumSHAKE128":  q.SumSHAKE128,
				"SumSHAKE256":  q.SumSHAKE256,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
