// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package rand

import (
	q "crypto/rand"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/rand", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars: map[string]interface{}{
				"Reader": &q.Reader,
			},
			Funcs: map[string]interface{}{
				"Int":   q.Int,
				"Prime": q.Prime,
				"Read":  q.Read,
				"Text":  q.Text,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
