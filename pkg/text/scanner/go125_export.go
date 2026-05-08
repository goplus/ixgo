// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package scanner

import (
	q "text/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("text/scanner", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "scanner",
			Path: "text/scanner",
			Deps: map[string]string{
				"bytes":        "bytes",
				"fmt":          "fmt",
				"io":           "io",
				"os":           "os",
				"unicode":      "unicode",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Position": reflect.TypeOf((*q.Position)(nil)).Elem(),
				"Scanner":  reflect.TypeOf((*q.Scanner)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"TokenString": reflect.ValueOf(q.TokenString),
			},
			TypedConsts: map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Char":           {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Char))},
				"Comment":        {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Comment))},
				"EOF":            {Typ: "untyped int", Value: constant.MakeInt64(int64(q.EOF))},
				"Float":          {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Float))},
				"GoTokens":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.GoTokens))},
				"GoWhitespace":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.GoWhitespace))},
				"Ident":          {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Ident))},
				"Int":            {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Int))},
				"RawString":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.RawString))},
				"ScanChars":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ScanChars))},
				"ScanComments":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ScanComments))},
				"ScanFloats":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ScanFloats))},
				"ScanIdents":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ScanIdents))},
				"ScanInts":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ScanInts))},
				"ScanRawStrings": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ScanRawStrings))},
				"ScanStrings":    {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ScanStrings))},
				"SkipComments":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.SkipComments))},
				"String":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.String))},
			},
		}
	})
}
