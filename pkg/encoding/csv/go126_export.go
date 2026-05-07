// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package csv

import (
	q "encoding/csv"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/csv", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "csv",
			Path: "encoding/csv",
			Deps: map[string]string{
				"bufio":        "bufio",
				"bytes":        "bytes",
				"errors":       "errors",
				"fmt":          "fmt",
				"io":           "io",
				"strings":      "strings",
				"unicode":      "unicode",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"ParseError": reflect.TypeOf((*q.ParseError)(nil)).Elem(),
				"Reader":     reflect.TypeOf((*q.Reader)(nil)).Elem(),
				"Writer":     reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrBareQuote":     &q.ErrBareQuote,
				"ErrFieldCount":    &q.ErrFieldCount,
				"ErrQuote":         &q.ErrQuote,
				"ErrTrailingComma": &q.ErrTrailingComma,
			},
			Funcs: map[string]interface{}{
				"NewReader": q.NewReader,
				"NewWriter": q.NewWriter,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
