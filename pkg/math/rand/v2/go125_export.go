// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package rand

import (
	q "math/rand/v2"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ExpFloat64":  reflect.ValueOf(q.ExpFloat64),
			"Float32":     reflect.ValueOf(q.Float32),
			"Float64":     reflect.ValueOf(q.Float64),
			"Int":         reflect.ValueOf(q.Int),
			"Int32":       reflect.ValueOf(q.Int32),
			"Int32N":      reflect.ValueOf(q.Int32N),
			"Int64":       reflect.ValueOf(q.Int64),
			"Int64N":      reflect.ValueOf(q.Int64N),
			"IntN":        reflect.ValueOf(q.IntN),
			"New":         reflect.ValueOf(q.New),
			"NewChaCha8":  reflect.ValueOf(q.NewChaCha8),
			"NewPCG":      reflect.ValueOf(q.NewPCG),
			"NewZipf":     reflect.ValueOf(q.NewZipf),
			"NormFloat64": reflect.ValueOf(q.NormFloat64),
			"Perm":        reflect.ValueOf(q.Perm),
			"Shuffle":     reflect.ValueOf(q.Shuffle),
			"Uint":        reflect.ValueOf(q.Uint),
			"Uint32":      reflect.ValueOf(q.Uint32),
			"Uint32N":     reflect.ValueOf(q.Uint32N),
			"Uint64":      reflect.ValueOf(q.Uint64),
			"Uint64N":     reflect.ValueOf(q.Uint64N),
			"UintN":       reflect.ValueOf(q.UintN),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
