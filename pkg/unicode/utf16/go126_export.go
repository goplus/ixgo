// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package utf16

import (
	q "unicode/utf16"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("unicode/utf16", func() *ixgo.Package {
		return &ixgo.Package{
			Name:       "utf16",
			Path:       "unicode/utf16",
			Deps:       map[string]string{},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"AppendRune":  q.AppendRune,
				"Decode":      q.Decode,
				"DecodeRune":  q.DecodeRune,
				"Encode":      q.Encode,
				"EncodeRune":  q.EncodeRune,
				"IsSurrogate": q.IsSurrogate,
				"RuneLen":     q.RuneLen,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
