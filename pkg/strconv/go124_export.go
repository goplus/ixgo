// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package strconv

import (
	q "strconv"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("strconv", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "strconv",
			Path: "strconv",
			Deps: map[string]string{
				"errors":               "errors",
				"internal/bytealg":     "bytealg",
				"internal/stringslite": "stringslite",
				"math":                 "math",
				"math/bits":            "bits",
				"unicode/utf8":         "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"NumError": reflect.TypeOf((*q.NumError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrRange":  &q.ErrRange,
				"ErrSyntax": &q.ErrSyntax,
			},
			Funcs: map[string]interface{}{
				"AppendBool":               q.AppendBool,
				"AppendFloat":              q.AppendFloat,
				"AppendInt":                q.AppendInt,
				"AppendQuote":              q.AppendQuote,
				"AppendQuoteRune":          q.AppendQuoteRune,
				"AppendQuoteRuneToASCII":   q.AppendQuoteRuneToASCII,
				"AppendQuoteRuneToGraphic": q.AppendQuoteRuneToGraphic,
				"AppendQuoteToASCII":       q.AppendQuoteToASCII,
				"AppendQuoteToGraphic":     q.AppendQuoteToGraphic,
				"AppendUint":               q.AppendUint,
				"Atoi":                     q.Atoi,
				"CanBackquote":             q.CanBackquote,
				"FormatBool":               q.FormatBool,
				"FormatComplex":            q.FormatComplex,
				"FormatFloat":              q.FormatFloat,
				"FormatInt":                q.FormatInt,
				"FormatUint":               q.FormatUint,
				"IsGraphic":                q.IsGraphic,
				"IsPrint":                  q.IsPrint,
				"Itoa":                     q.Itoa,
				"ParseBool":                q.ParseBool,
				"ParseComplex":             q.ParseComplex,
				"ParseFloat":               q.ParseFloat,
				"ParseInt":                 q.ParseInt,
				"ParseUint":                q.ParseUint,
				"Quote":                    q.Quote,
				"QuoteRune":                q.QuoteRune,
				"QuoteRuneToASCII":         q.QuoteRuneToASCII,
				"QuoteRuneToGraphic":       q.QuoteRuneToGraphic,
				"QuoteToASCII":             q.QuoteToASCII,
				"QuoteToGraphic":           q.QuoteToGraphic,
				"QuotedPrefix":             q.QuotedPrefix,
				"Unquote":                  q.Unquote,
				"UnquoteChar":              q.UnquoteChar,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"IntSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.IntSize))},
			},
		}
	})
}
