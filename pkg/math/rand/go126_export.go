// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package rand

import (
	q "math/rand"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("math/rand", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "rand",
			Path: "math/rand",
			Deps: map[string]string{
				"internal/godebug": "godebug",
				"math":             "math",
				"sync":             "sync",
				"sync/atomic":      "atomic",
				"unsafe":           "unsafe",
			},
			Interfaces: map[string]reflect.Type{
				"Source":   reflect.TypeOf((*q.Source)(nil)).Elem(),
				"Source64": reflect.TypeOf((*q.Source64)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Rand": reflect.TypeOf((*q.Rand)(nil)).Elem(),
				"Zipf": reflect.TypeOf((*q.Zipf)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"ExpFloat64":  q.ExpFloat64,
				"Float32":     q.Float32,
				"Float64":     q.Float64,
				"Int":         q.Int,
				"Int31":       q.Int31,
				"Int31n":      q.Int31n,
				"Int63":       q.Int63,
				"Int63n":      q.Int63n,
				"Intn":        q.Intn,
				"New":         q.New,
				"NewSource":   q.NewSource,
				"NewZipf":     q.NewZipf,
				"NormFloat64": q.NormFloat64,
				"Perm":        q.Perm,
				"Read":        q.Read,
				"Seed":        q.Seed,
				"Shuffle":     q.Shuffle,
				"Uint32":      q.Uint32,
				"Uint64":      q.Uint64,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
