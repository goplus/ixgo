// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package base64

import (
	q "encoding/base64"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/base64", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "base64",
			Path: "encoding/base64",
			Deps: map[string]string{
				"internal/byteorder": "byteorder",
				"io":                 "io",
				"slices":             "slices",
				"strconv":            "strconv",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"CorruptInputError": reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(),
				"Encoding":          reflect.TypeOf((*q.Encoding)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"RawStdEncoding": &q.RawStdEncoding,
				"RawURLEncoding": &q.RawURLEncoding,
				"StdEncoding":    &q.StdEncoding,
				"URLEncoding":    &q.URLEncoding,
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
