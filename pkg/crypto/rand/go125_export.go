// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package rand

import (
	q "crypto/rand"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "rand",
		Path: "crypto/rand",
		Deps: map[string]string{
			"crypto/internal/boring":       "boring",
			"crypto/internal/fips140":      "fips140",
			"crypto/internal/fips140/drbg": "drbg",
			"crypto/internal/fips140only":  "fips140only",
			"crypto/internal/randutil":     "randutil",
			"crypto/internal/sysrand":      "sysrand",
			"errors":                       "errors",
			"io":                           "io",
			"math/big":                     "big",
			"unsafe":                       "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Reader": reflect.ValueOf(&q.Reader),
		},
		Funcs: map[string]reflect.Value{
			"Int":   reflect.ValueOf(q.Int),
			"Prime": reflect.ValueOf(q.Prime),
			"Read":  reflect.ValueOf(q.Read),
			"Text":  reflect.ValueOf(q.Text),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
