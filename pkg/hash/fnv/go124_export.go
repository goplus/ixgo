// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package fnv

import (
	q "hash/fnv"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("hash/fnv", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "fnv",
			Path: "hash/fnv",
			Deps: map[string]string{
				"errors":             "errors",
				"hash":               "hash",
				"internal/byteorder": "byteorder",
				"math/bits":          "bits",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New128":  q.New128,
				"New128a": q.New128a,
				"New32":   q.New32,
				"New32a":  q.New32a,
				"New64":   q.New64,
				"New64a":  q.New64a,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
