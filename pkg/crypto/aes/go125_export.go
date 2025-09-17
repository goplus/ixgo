// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package aes

import (
	q "crypto/aes"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "aes",
		Path: "crypto/aes",
		Deps: map[string]string{
			"crypto/cipher":               "cipher",
			"crypto/internal/boring":      "boring",
			"crypto/internal/fips140/aes": "aes",
			"strconv":                     "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"KeySizeError": reflect.TypeOf((*q.KeySizeError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher": reflect.ValueOf(q.NewCipher),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
		},
	})
}
