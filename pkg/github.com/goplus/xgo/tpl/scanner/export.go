// export by github.com/goplus/ixgo/cmd/qexp

package scanner

import (
	q "github.com/goplus/xgo/tpl/scanner"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/scanner", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "scanner",
			Path: "github.com/goplus/xgo/tpl/scanner",
			Deps: map[string]string{
				"bytes":                           "bytes",
				"fmt":                             "fmt",
				"github.com/goplus/xgo/tpl/token": "token",
				"github.com/goplus/xgo/tpl/types": "types",
				"go/scanner":                      "scanner",
				"io":                              "io",
				"path/filepath":                   "filepath",
				"strconv":                         "strconv",
				"unicode":                         "unicode",
				"unicode/utf8":                    "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"ErrorHandler": reflect.TypeOf((*q.ErrorHandler)(nil)).Elem(),
				"Mode":         reflect.TypeOf((*q.Mode)(nil)).Elem(),
				"Scanner":      reflect.TypeOf((*q.Scanner)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"Error":     reflect.TypeOf((*q.Error)(nil)).Elem(),
				"ErrorList": reflect.TypeOf((*q.ErrorList)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"PrintError": q.PrintError,
			},
			TypedConsts: map[string]interface{}{
				"NoInsertSemis": q.NoInsertSemis,
				"ScanComments":  q.ScanComments,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
