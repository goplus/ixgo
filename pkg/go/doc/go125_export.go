// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package doc

import (
	q "go/doc"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("go/doc", func() *ixgo.Package {
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
			Vars: map[string]reflect.Value{
				"IllegalPrefixes": reflect.ValueOf(&q.IllegalPrefixes),
			},
			Funcs: map[string]reflect.Value{
				"Examples":      reflect.ValueOf(q.Examples),
				"IsPredeclared": reflect.ValueOf(q.IsPredeclared),
				"New":           reflect.ValueOf(q.New),
				"NewFromFiles":  reflect.ValueOf(q.NewFromFiles),
				"Synopsis":      reflect.ValueOf(q.Synopsis),
				"ToHTML":        reflect.ValueOf(q.ToHTML),
				"ToText":        reflect.ValueOf(q.ToText),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"AllDecls":    {Typ: reflect.TypeOf(q.AllDecls), Value: constant.MakeInt64(int64(q.AllDecls))},
				"AllMethods":  {Typ: reflect.TypeOf(q.AllMethods), Value: constant.MakeInt64(int64(q.AllMethods))},
				"PreserveAST": {Typ: reflect.TypeOf(q.PreserveAST), Value: constant.MakeInt64(int64(q.PreserveAST))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
