// export by github.com/goplus/ixgo/cmd/qexp

package scanner

import (
	q "github.com/goplus/xgo/scanner"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/scanner", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "scanner",
			Path: "github.com/goplus/xgo/scanner",
			Deps: map[string]string{
				"bytes":                       "bytes",
				"fmt":                         "fmt",
				"github.com/goplus/xgo/token": "token",
				"github.com/qiniu/x/byteutil": "byteutil",
				"go/scanner":                  "scanner",
				"io":                          "io",
				"path/filepath":               "filepath",
				"strconv":                     "strconv",
				"unicode":                     "unicode",
				"unicode/utf8":                "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Mode":    reflect.TypeOf((*q.Mode)(nil)).Elem(),
				"Scanner": reflect.TypeOf((*q.Scanner)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"Error":        reflect.TypeOf((*q.Error)(nil)).Elem(),
				"ErrorHandler": reflect.TypeOf((*q.ErrorHandler)(nil)).Elem(),
				"ErrorList":    reflect.TypeOf((*q.ErrorList)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New":        q.New,
				"PrintError": q.PrintError,
			},
			TypedConsts: map[string]interface{}{
				"ScanComments": q.ScanComments,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
