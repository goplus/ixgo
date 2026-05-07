// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package ascii85

import (
	q "encoding/ascii85"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/ascii85", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "ascii85",
			Path: "encoding/ascii85",
			Deps: map[string]string{
				"io":      "io",
				"strconv": "strconv",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"CorruptInputError": reflect.TypeOf((*q.CorruptInputError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Decode":        q.Decode,
				"Encode":        q.Encode,
				"MaxEncodedLen": q.MaxEncodedLen,
				"NewDecoder":    q.NewDecoder,
				"NewEncoder":    q.NewEncoder,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
