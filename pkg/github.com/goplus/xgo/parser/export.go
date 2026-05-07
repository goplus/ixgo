// export by github.com/goplus/ixgo/cmd/qexp

package parser

import (
	q "github.com/goplus/xgo/parser"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/parser", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars: map[string]interface{}{
				"ErrUnknownFileKind": &q.ErrUnknownFileKind,
			},
			Funcs: map[string]interface{}{
				"Parse":          q.Parse,
				"ParseDir":       q.ParseDir,
				"ParseDirEx":     q.ParseDirEx,
				"ParseEntries":   q.ParseEntries,
				"ParseEntry":     q.ParseEntry,
				"ParseExpr":      q.ParseExpr,
				"ParseExprEx":    q.ParseExprEx,
				"ParseExprFrom":  q.ParseExprFrom,
				"ParseFSDir":     q.ParseFSDir,
				"ParseFSEntries": q.ParseFSEntries,
				"ParseFSEntry":   q.ParseFSEntry,
				"ParseFSFile":    q.ParseFSFile,
				"ParseFSFiles":   q.ParseFSFiles,
				"ParseFile":      q.ParseFile,
				"ParseFiles":     q.ParseFiles,
				"SetDebug":       q.SetDebug,
			},
			TypedConsts: map[string]interface{}{
				"AllErrors":         q.AllErrors,
				"DeclarationErrors": q.DeclarationErrors,
				"ImportsOnly":       q.ImportsOnly,
				"PackageClauseOnly": q.PackageClauseOnly,
				"ParseComments":     q.ParseComments,
				"ParseGoAsGoPlus":   q.ParseGoAsGoPlus,
				"ParseGoAsXGo":      q.ParseGoAsXGo,
				"ParseGoPlusClass":  q.ParseGoPlusClass,
				"ParseXGoClass":     q.ParseXGoClass,
				"SaveAbsFile":       q.SaveAbsFile,
				"Trace":             q.Trace,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"DbgFlagAll":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DbgFlagAll))},
				"DbgFlagParseError":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DbgFlagParseError))},
				"DbgFlagParseOutput": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DbgFlagParseOutput))},
			},
		}
	})
}
