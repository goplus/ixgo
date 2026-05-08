// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package printer

import (
	q "go/printer"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("go/printer", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "printer",
			Path: "go/printer",
			Deps: map[string]string{
				"fmt":                 "fmt",
				"go/ast":              "ast",
				"go/build/constraint": "constraint",
				"go/doc/comment":      "comment",
				"go/token":            "token",
				"io":                  "io",
				"math":                "math",
				"os":                  "os",
				"slices":              "slices",
				"strconv":             "strconv",
				"strings":             "strings",
				"sync":                "sync",
				"text/tabwriter":      "tabwriter",
				"unicode":             "unicode",
				"unicode/utf8":        "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"CommentedNode": reflect.TypeOf((*q.CommentedNode)(nil)).Elem(),
				"Config":        reflect.TypeOf((*q.Config)(nil)).Elem(),
				"Mode":          reflect.TypeOf((*q.Mode)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"Fprint": reflect.ValueOf(q.Fprint),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"RawFormat": {Typ: reflect.TypeOf(q.RawFormat), Value: constant.MakeInt64(int64(q.RawFormat))},
				"SourcePos": {Typ: reflect.TypeOf(q.SourcePos), Value: constant.MakeInt64(int64(q.SourcePos))},
				"TabIndent": {Typ: reflect.TypeOf(q.TabIndent), Value: constant.MakeInt64(int64(q.TabIndent))},
				"UseSpaces": {Typ: reflect.TypeOf(q.UseSpaces), Value: constant.MakeInt64(int64(q.UseSpaces))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
