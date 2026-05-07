// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package utf8

import (
	q "unicode/utf8"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("unicode/utf8", func() *ixgo.Package {
		return &ixgo.Package{
			Name:       "utf8",
			Path:       "unicode/utf8",
			Deps:       map[string]string{},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"AppendRune":             q.AppendRune,
				"DecodeLastRune":         q.DecodeLastRune,
				"DecodeLastRuneInString": q.DecodeLastRuneInString,
				"DecodeRune":             q.DecodeRune,
				"DecodeRuneInString":     q.DecodeRuneInString,
				"EncodeRune":             q.EncodeRune,
				"FullRune":               q.FullRune,
				"FullRuneInString":       q.FullRuneInString,
				"RuneCount":              q.RuneCount,
				"RuneCountInString":      q.RuneCountInString,
				"RuneLen":                q.RuneLen,
				"RuneStart":              q.RuneStart,
				"Valid":                  q.Valid,
				"ValidRune":              q.ValidRune,
				"ValidString":            q.ValidString,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"MaxRune":   {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.MaxRune))},
				"RuneError": {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.RuneError))},
				"RuneSelf":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.RuneSelf))},
				"UTFMax":    {Typ: "untyped int", Value: constant.MakeInt64(int64(q.UTFMax))},
			},
		}
	})
}
