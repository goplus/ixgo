// export by github.com/goplus/ixgo/cmd/qexp

package printer

import (
	q "github.com/goplus/xgo/printer"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/printer", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "printer",
			Path: "github.com/goplus/xgo/printer",
			Deps: map[string]string{
				"bytes":                       "bytes",
				"fmt":                         "fmt",
				"github.com/goplus/xgo/ast":   "ast",
				"github.com/goplus/xgo/token": "token",
				"io":                          "io",
				"log":                         "log",
				"math":                        "math",
				"os":                          "os",
				"strconv":                     "strconv",
				"strings":                     "strings",
				"text/tabwriter":              "tabwriter",
				"unicode":                     "unicode",
				"unicode/utf8":                "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"CommentedNode":  reflect.TypeOf((*q.CommentedNode)(nil)).Elem(),
				"CommentedNodes": reflect.TypeOf((*q.CommentedNodes)(nil)).Elem(),
				"Config":         reflect.TypeOf((*q.Config)(nil)).Elem(),
				"Mode":           reflect.TypeOf((*q.Mode)(nil)).Elem(),
				"NewlineStmt":    reflect.TypeOf((*q.NewlineStmt)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Fprint":   q.Fprint,
				"SetDebug": q.SetDebug,
			},
			TypedConsts: map[string]interface{}{
				"RawFormat": q.RawFormat,
				"SourcePos": q.SourcePos,
				"TabIndent": q.TabIndent,
				"UseSpaces": q.UseSpaces,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"DbgFlagAll": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DbgFlagAll))},
			},
		}
	})
}
