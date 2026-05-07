// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package base32

import (
	q "encoding/base32"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/base32", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "base32",
			Path: "encoding/base32",
			Deps: map[string]string{
				"io":      "io",
				"slices":  "slices",
				"strconv": "strconv",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"CorruptInputError": reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(),
				"Encoding":          reflect.TypeOf((*q.Encoding)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"HexEncoding": &q.HexEncoding,
				"StdEncoding": &q.StdEncoding,
			},
			Funcs: map[string]interface{}{
				"NewDecoder":  q.NewDecoder,
				"NewEncoder":  q.NewEncoder,
				"NewEncoding": q.NewEncoding,
			},
			TypedConsts: map[string]interface{}{
				"NoPadding":  q.NoPadding,
				"StdPadding": q.StdPadding,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
