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
			"RawFormat": {reflect.TypeOf(q.RawFormat), constant.MakeInt64(int64(q.RawFormat))},
			"SourcePos": {reflect.TypeOf(q.SourcePos), constant.MakeInt64(int64(q.SourcePos))},
			"TabIndent": {reflect.TypeOf(q.TabIndent), constant.MakeInt64(int64(q.TabIndent))},
			"UseSpaces": {reflect.TypeOf(q.UseSpaces), constant.MakeInt64(int64(q.UseSpaces))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"DbgFlagAll": {"untyped int", constant.MakeInt64(int64(q.DbgFlagAll))},
		},
	})
}
