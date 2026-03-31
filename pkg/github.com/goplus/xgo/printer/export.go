// export by github.com/goplus/ixgo/cmd/qexp

package printer

import (
	q "github.com/goplus/xgo/printer"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Fprint":   reflect.ValueOf(q.Fprint),
			"SetDebug": reflect.ValueOf(q.SetDebug),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"RawFormat": {Typ: reflect.TypeOf(q.RawFormat), Value: constant.MakeInt64(int64(q.RawFormat))},
			"SourcePos": {Typ: reflect.TypeOf(q.SourcePos), Value: constant.MakeInt64(int64(q.SourcePos))},
			"TabIndent": {Typ: reflect.TypeOf(q.TabIndent), Value: constant.MakeInt64(int64(q.TabIndent))},
			"UseSpaces": {Typ: reflect.TypeOf(q.UseSpaces), Value: constant.MakeInt64(int64(q.UseSpaces))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"DbgFlagAll": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DbgFlagAll))},
		},
	})
}
