// export by github.com/goplus/ixgo/cmd/qexp

package tpl

import (
	q "github.com/goplus/xgo/tpl"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "tpl",
			Path: "github.com/goplus/xgo/tpl",
			Deps: map[string]string{
				"fmt":                               "fmt",
				"github.com/goplus/xgo/tpl/ast":     "ast",
				"github.com/goplus/xgo/tpl/cl":      "cl",
				"github.com/goplus/xgo/tpl/matcher": "matcher",
				"github.com/goplus/xgo/tpl/parser":  "parser",
				"github.com/goplus/xgo/tpl/scanner": "scanner",
				"github.com/goplus/xgo/tpl/token":   "token",
				"github.com/goplus/xgo/tpl/types":   "types",
				"github.com/qiniu/x/errors":         "errors",
				"github.com/qiniu/x/stream":         "stream",
				"io":                                "io",
				"os":                                "os",
				"reflect":                           "reflect",
			},
			Interfaces: map[string]reflect.Type{
				"Scanner": reflect.TypeOf((*q.Scanner)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Compiler":   reflect.TypeOf((*q.Compiler)(nil)).Elem(),
				"Config":     reflect.TypeOf((*q.Config)(nil)).Elem(),
				"MatchState": reflect.TypeOf((*q.MatchState)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"Error": reflect.TypeOf((*q.Error)(nil)).Elem(),
				"Token": reflect.TypeOf((*q.Token)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"BasicLit":     q.BasicLit,
				"BinaryExpr":   q.BinaryExpr,
				"BinaryExprNR": q.BinaryExprNR,
				"BinaryExprR":  q.BinaryExprR,
				"BinaryOp":     q.BinaryOp,
				"BinaryOpNR":   q.BinaryOpNR,
				"BinaryOpR":    q.BinaryOpR,
				"Dump":         q.Dump,
				"Fdump":        q.Fdump,
				"FromFile":     q.FromFile,
				"Ident":        q.Ident,
				"List":         q.List,
				"New":          q.New,
				"NewEx":        q.NewEx,
				"Panic":        q.Panic,
				"RangeOp":      q.RangeOp,
				"Relocate":     q.Relocate,
				"ShowConflict": q.ShowConflict,
				"UnaryExpr":    q.UnaryExpr,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
