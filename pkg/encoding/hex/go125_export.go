// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package hex

import (
	q "encoding/hex"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/hex", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "hex",
			Path: "encoding/hex",
			Deps: map[string]string{
				"errors":  "errors",
				"fmt":     "fmt",
				"io":      "io",
				"slices":  "slices",
				"strings": "strings",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"InvalidByteError": reflect.TypeOf((*q.InvalidByteError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrLength": &q.ErrLength,
			},
			Funcs: map[string]interface{}{
				"AppendDecode":   q.AppendDecode,
				"AppendEncode":   q.AppendEncode,
				"Decode":         q.Decode,
				"DecodeString":   q.DecodeString,
				"DecodedLen":     q.DecodedLen,
				"Dump":           q.Dump,
				"Dumper":         q.Dumper,
				"Encode":         q.Encode,
				"EncodeToString": q.EncodeToString,
				"EncodedLen":     q.EncodedLen,
				"NewDecoder":     q.NewDecoder,
				"NewEncoder":     q.NewEncoder,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
