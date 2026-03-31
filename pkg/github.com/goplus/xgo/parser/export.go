// export by github.com/goplus/ixgo/cmd/qexp

package parser

import (
	q "github.com/goplus/xgo/parser"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "parser",
		Path: "github.com/goplus/xgo/parser",
		Deps: map[string]string{
			"errors":                           "errors",
			"fmt":                              "fmt",
			"github.com/goplus/xgo/ast":        "ast",
			"github.com/goplus/xgo/parser/fsx": "fsx",
			"github.com/goplus/xgo/scanner":    "scanner",
			"github.com/goplus/xgo/token":      "token",
			"github.com/goplus/xgo/tpl/ast":    "ast",
			"github.com/goplus/xgo/tpl/parser": "parser",
			"github.com/qiniu/x/log":           "log",
			"github.com/qiniu/x/stream":        "stream",
			"go/ast":                           "ast",
			"go/parser":                        "parser",
			"go/scanner":                       "scanner",
			"io/fs":                            "fs",
			"path":                             "path",
			"strconv":                          "strconv",
			"strings":                          "strings",
			"unicode":                          "unicode",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Config": reflect.TypeOf((*q.Config)(nil)).Elem(),
			"Mode":   reflect.TypeOf((*q.Mode)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"FileSystem": reflect.TypeOf((*q.FileSystem)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ErrUnknownFileKind": reflect.ValueOf(&q.ErrUnknownFileKind),
		},
		Funcs: map[string]reflect.Value{
			"Parse":          reflect.ValueOf(q.Parse),
			"ParseDir":       reflect.ValueOf(q.ParseDir),
			"ParseDirEx":     reflect.ValueOf(q.ParseDirEx),
			"ParseEntries":   reflect.ValueOf(q.ParseEntries),
			"ParseEntry":     reflect.ValueOf(q.ParseEntry),
			"ParseExpr":      reflect.ValueOf(q.ParseExpr),
			"ParseExprEx":    reflect.ValueOf(q.ParseExprEx),
			"ParseExprFrom":  reflect.ValueOf(q.ParseExprFrom),
			"ParseFSDir":     reflect.ValueOf(q.ParseFSDir),
			"ParseFSEntries": reflect.ValueOf(q.ParseFSEntries),
			"ParseFSEntry":   reflect.ValueOf(q.ParseFSEntry),
			"ParseFSFile":    reflect.ValueOf(q.ParseFSFile),
			"ParseFSFiles":   reflect.ValueOf(q.ParseFSFiles),
			"ParseFile":      reflect.ValueOf(q.ParseFile),
			"ParseFiles":     reflect.ValueOf(q.ParseFiles),
			"SetDebug":       reflect.ValueOf(q.SetDebug),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"AllErrors":         {Typ: reflect.TypeOf(q.AllErrors), Value: constant.MakeInt64(int64(q.AllErrors))},
			"DeclarationErrors": {Typ: reflect.TypeOf(q.DeclarationErrors), Value: constant.MakeInt64(int64(q.DeclarationErrors))},
			"ImportsOnly":       {Typ: reflect.TypeOf(q.ImportsOnly), Value: constant.MakeInt64(int64(q.ImportsOnly))},
			"PackageClauseOnly": {Typ: reflect.TypeOf(q.PackageClauseOnly), Value: constant.MakeInt64(int64(q.PackageClauseOnly))},
			"ParseComments":     {Typ: reflect.TypeOf(q.ParseComments), Value: constant.MakeInt64(int64(q.ParseComments))},
			"ParseGoAsGoPlus":   {Typ: reflect.TypeOf(q.ParseGoAsGoPlus), Value: constant.MakeInt64(int64(q.ParseGoAsGoPlus))},
			"ParseGoAsXGo":      {Typ: reflect.TypeOf(q.ParseGoAsXGo), Value: constant.MakeInt64(int64(q.ParseGoAsXGo))},
			"ParseGoPlusClass":  {Typ: reflect.TypeOf(q.ParseGoPlusClass), Value: constant.MakeInt64(int64(q.ParseGoPlusClass))},
			"ParseXGoClass":     {Typ: reflect.TypeOf(q.ParseXGoClass), Value: constant.MakeInt64(int64(q.ParseXGoClass))},
			"SaveAbsFile":       {Typ: reflect.TypeOf(q.SaveAbsFile), Value: constant.MakeInt64(int64(q.SaveAbsFile))},
			"Trace":             {Typ: reflect.TypeOf(q.Trace), Value: constant.MakeInt64(int64(q.Trace))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"DbgFlagAll":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DbgFlagAll))},
			"DbgFlagParseError":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DbgFlagParseError))},
			"DbgFlagParseOutput": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DbgFlagParseOutput))},
		},
	})
}
