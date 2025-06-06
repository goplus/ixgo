// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.19 && !go1.20
// +build go1.19,!go1.20

package rc4

import (
	q "crypto/rc4"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "rc4",
		Path: "crypto/rc4",
		Deps: map[string]string{
			"crypto/internal/subtle": "subtle",
			"strconv":                "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Cipher":       reflect.TypeOf((*q.Cipher)(nil)).Elem(),
			"KeySizeError": reflect.TypeOf((*q.KeySizeError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher": reflect.ValueOf(q.NewCipher),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
