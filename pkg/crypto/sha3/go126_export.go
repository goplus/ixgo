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
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New224":       reflect.ValueOf(q.New224),
			"New256":       reflect.ValueOf(q.New256),
			"New384":       reflect.ValueOf(q.New384),
			"New512":       reflect.ValueOf(q.New512),
			"NewCSHAKE128": reflect.ValueOf(q.NewCSHAKE128),
			"NewCSHAKE256": reflect.ValueOf(q.NewCSHAKE256),
			"NewSHAKE128":  reflect.ValueOf(q.NewSHAKE128),
			"NewSHAKE256":  reflect.ValueOf(q.NewSHAKE256),
			"Sum224":       reflect.ValueOf(q.Sum224),
			"Sum256":       reflect.ValueOf(q.Sum256),
			"Sum384":       reflect.ValueOf(q.Sum384),
			"Sum512":       reflect.ValueOf(q.Sum512),
			"SumSHAKE128":  reflect.ValueOf(q.SumSHAKE128),
			"SumSHAKE256":  reflect.ValueOf(q.SumSHAKE256),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
