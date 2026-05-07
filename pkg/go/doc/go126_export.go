// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package doc

import (
	q "go/doc"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("go/doc", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "doc",
			Path: "go/doc",
			Deps: map[string]string{
				"cmp":                 "cmp",
				"fmt":                 "fmt",
				"go/ast":              "ast",
				"go/doc/comment":      "comment",
				"go/token":            "token",
				"internal/lazyregexp": "lazyregexp",
				"io":                  "io",
				"path":                "path",
				"slices":              "slices",
				"strconv":             "strconv",
				"strings":             "strings",
				"unicode":             "unicode",
				"unicode/utf8":        "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Example": reflect.TypeOf((*q.Example)(nil)).Elem(),
				"Filter":  reflect.TypeOf((*q.Filter)(nil)).Elem(),
				"Func":    reflect.TypeOf((*q.Func)(nil)).Elem(),
				"Mode":    reflect.TypeOf((*q.Mode)(nil)).Elem(),
				"Note":    reflect.TypeOf((*q.Note)(nil)).Elem(),
				"Package": reflect.TypeOf((*q.Package)(nil)).Elem(),
				"Type":    reflect.TypeOf((*q.Type)(nil)).Elem(),
				"Value":   reflect.TypeOf((*q.Value)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"IllegalPrefixes": &q.IllegalPrefixes,
			},
			Funcs: map[string]interface{}{
				"Examples":      q.Examples,
				"IsPredeclared": q.IsPredeclared,
				"New":           q.New,
				"NewFromFiles":  q.NewFromFiles,
				"Synopsis":      q.Synopsis,
				"ToHTML":        q.ToHTML,
				"ToText":        q.ToText,
			},
			TypedConsts: map[string]interface{}{
				"AllDecls":    q.AllDecls,
				"AllMethods":  q.AllMethods,
				"PreserveAST": q.PreserveAST,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
