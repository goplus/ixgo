// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package printer

import (
	q "go/printer"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("go/printer", func() *ixgo.Package {
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Fprint": q.Fprint,
			},
			TypedConsts: map[string]interface{}{
				"RawFormat": q.RawFormat,
				"SourcePos": q.SourcePos,
				"TabIndent": q.TabIndent,
				"UseSpaces": q.UseSpaces,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
