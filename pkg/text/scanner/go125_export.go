// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package scanner

import (
	q "text/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
			"Char":           {"untyped int", constant.MakeInt64(int64(q.Char))},
			"Comment":        {"untyped int", constant.MakeInt64(int64(q.Comment))},
			"EOF":            {"untyped int", constant.MakeInt64(int64(q.EOF))},
			"Float":          {"untyped int", constant.MakeInt64(int64(q.Float))},
			"GoTokens":       {"untyped int", constant.MakeInt64(int64(q.GoTokens))},
			"GoWhitespace":   {"untyped int", constant.MakeInt64(int64(q.GoWhitespace))},
			"Ident":          {"untyped int", constant.MakeInt64(int64(q.Ident))},
			"Int":            {"untyped int", constant.MakeInt64(int64(q.Int))},
			"RawString":      {"untyped int", constant.MakeInt64(int64(q.RawString))},
			"ScanChars":      {"untyped int", constant.MakeInt64(int64(q.ScanChars))},
			"ScanComments":   {"untyped int", constant.MakeInt64(int64(q.ScanComments))},
			"ScanFloats":     {"untyped int", constant.MakeInt64(int64(q.ScanFloats))},
			"ScanIdents":     {"untyped int", constant.MakeInt64(int64(q.ScanIdents))},
			"ScanInts":       {"untyped int", constant.MakeInt64(int64(q.ScanInts))},
			"ScanRawStrings": {"untyped int", constant.MakeInt64(int64(q.ScanRawStrings))},
			"ScanStrings":    {"untyped int", constant.MakeInt64(int64(q.ScanStrings))},
			"SkipComments":   {"untyped int", constant.MakeInt64(int64(q.SkipComments))},
			"String":         {"untyped int", constant.MakeInt64(int64(q.String))},
		},
	})
}
