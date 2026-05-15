// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package rand

import (
	q "math/rand"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ExpFloat64":  reflect.ValueOf(q.ExpFloat64),
			"Float32":     reflect.ValueOf(q.Float32),
			"Float64":     reflect.ValueOf(q.Float64),
			"Int":         reflect.ValueOf(q.Int),
			"Int31":       reflect.ValueOf(q.Int31),
			"Int31n":      reflect.ValueOf(q.Int31n),
			"Int63":       reflect.ValueOf(q.Int63),
			"Int63n":      reflect.ValueOf(q.Int63n),
			"Intn":        reflect.ValueOf(q.Intn),
			"New":         reflect.ValueOf(q.New),
			"NewSource":   reflect.ValueOf(q.NewSource),
			"NewZipf":     reflect.ValueOf(q.NewZipf),
			"NormFloat64": reflect.ValueOf(q.NormFloat64),
			"Perm":        reflect.ValueOf(q.Perm),
			"Read":        reflect.ValueOf(q.Read),
			"Seed":        reflect.ValueOf(q.Seed),
			"Shuffle":     reflect.ValueOf(q.Shuffle),
			"Uint32":      reflect.ValueOf(q.Uint32),
			"Uint64":      reflect.ValueOf(q.Uint64),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
