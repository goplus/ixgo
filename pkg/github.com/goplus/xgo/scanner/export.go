// export by github.com/goplus/ixgo/cmd/qexp

package scanner

import (
	q "github.com/goplus/xgo/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":        reflect.ValueOf(q.New),
			"PrintError": reflect.ValueOf(q.PrintError),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"ScanComments": {Typ: reflect.TypeOf(q.ScanComments), Value: constant.MakeInt64(int64(q.ScanComments))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
