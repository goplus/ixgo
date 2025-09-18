// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package sha256

import (
	q "crypto/sha256"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "sha256",
		Path: "crypto/sha256",
		Deps: map[string]string{
			"crypto":                         "crypto",
			"crypto/internal/boring":         "boring",
			"crypto/internal/fips140/sha256": "sha256",
			"hash":                           "hash",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":    reflect.ValueOf(q.New),
			"New224": reflect.ValueOf(q.New224),
			"Sum224": reflect.ValueOf(q.Sum224),
			"Sum256": reflect.ValueOf(q.Sum256),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"Size":      {"untyped int", constant.MakeInt64(int64(q.Size))},
			"Size224":   {"untyped int", constant.MakeInt64(int64(q.Size224))},
		},
	})
}
