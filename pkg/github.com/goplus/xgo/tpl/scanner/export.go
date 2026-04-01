// export by github.com/goplus/ixgo/cmd/qexp

package scanner

import (
	q "github.com/goplus/xgo/tpl/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"PrintError": reflect.ValueOf(q.PrintError),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"NoInsertSemis": {Typ: reflect.TypeOf(q.NoInsertSemis), Value: constant.MakeInt64(int64(q.NoInsertSemis))},
			"ScanComments":  {Typ: reflect.TypeOf(q.ScanComments), Value: constant.MakeInt64(int64(q.ScanComments))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
