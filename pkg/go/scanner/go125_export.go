// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package scanner

import (
	q "go/scanner"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "scanner",
		Path: "go/scanner",
		Deps: map[string]string{
			"bytes":         "bytes",
			"fmt":           "fmt",
			"go/token":      "token",
			"io":            "io",
			"path/filepath": "filepath",
			"sort":          "sort",
			"strconv":       "strconv",
			"unicode":       "unicode",
			"unicode/utf8":  "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Error":        reflect.TypeOf((*q.Error)(nil)).Elem(),
			"ErrorHandler": reflect.TypeOf((*q.ErrorHandler)(nil)).Elem(),
			"ErrorList":    reflect.TypeOf((*q.ErrorList)(nil)).Elem(),
			"Mode":         reflect.TypeOf((*q.Mode)(nil)).Elem(),
			"Scanner":      reflect.TypeOf((*q.Scanner)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"PrintError": reflect.ValueOf(q.PrintError),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"ScanComments": {reflect.TypeOf(q.ScanComments), constant.MakeInt64(int64(q.ScanComments))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
