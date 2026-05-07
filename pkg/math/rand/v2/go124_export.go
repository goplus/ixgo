// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package rand

import (
	q "math/rand/v2"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("math/rand/v2", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "rand",
			Path: "math/rand/v2",
			Deps: map[string]string{
				"errors":               "errors",
				"internal/byteorder":   "byteorder",
				"internal/chacha8rand": "chacha8rand",
				"math":                 "math",
				"math/bits":            "bits",
				"unsafe":               "unsafe",
			},
			Interfaces: map[string]reflect.Type{
				"Source": reflect.TypeOf((*q.Source)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"ChaCha8": reflect.TypeOf((*q.ChaCha8)(nil)).Elem(),
				"PCG":     reflect.TypeOf((*q.PCG)(nil)).Elem(),
				"Rand":    reflect.TypeOf((*q.Rand)(nil)).Elem(),
				"Zipf":    reflect.TypeOf((*q.Zipf)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"ExpFloat64":  q.ExpFloat64,
				"Float32":     q.Float32,
				"Float64":     q.Float64,
				"Int":         q.Int,
				"Int32":       q.Int32,
				"Int32N":      q.Int32N,
				"Int64":       q.Int64,
				"Int64N":      q.Int64N,
				"IntN":        q.IntN,
				"New":         q.New,
				"NewChaCha8":  q.NewChaCha8,
				"NewPCG":      q.NewPCG,
				"NewZipf":     q.NewZipf,
				"NormFloat64": q.NormFloat64,
				"Perm":        q.Perm,
				"Shuffle":     q.Shuffle,
				"Uint":        q.Uint,
				"Uint32":      q.Uint32,
				"Uint32N":     q.Uint32N,
				"Uint64":      q.Uint64,
				"Uint64N":     q.Uint64N,
				"UintN":       q.UintN,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
