// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package cryptotest

import (
	q "testing/cryptotest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "cryptotest",
		Path: "testing/cryptotest",
		Deps: map[string]string{
			"crypto/rand":        "rand",
			"internal/byteorder": "byteorder",
			"io":                 "io",
			"math/rand/v2":       "rand",
			"sync":               "sync",
			"testing":            "testing",
			"unsafe":             "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"SetGlobalRandom": reflect.ValueOf(q.SetGlobalRandom),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
