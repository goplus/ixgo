// export by github.com/goplus/ixgo/cmd/qexp

package ast

import (
	q "github.com/goplus/xgo/tpl/ast"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/ast", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "ast",
			Path: "github.com/goplus/xgo/tpl/ast",
			Deps: map[string]string{
				"github.com/goplus/xgo/tpl/token": "token",
			},
			Interfaces: map[string]reflect.Type{
				"Decl": reflect.TypeOf((*q.Decl)(nil)).Elem(),
				"Expr": reflect.TypeOf((*q.Expr)(nil)).Elem(),
				"Node": reflect.TypeOf((*q.Node)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"BasicLit":   reflect.TypeOf((*q.BasicLit)(nil)).Elem(),
				"BinaryExpr": reflect.TypeOf((*q.BinaryExpr)(nil)).Elem(),
				"Choice":     reflect.TypeOf((*q.Choice)(nil)).Elem(),
				"File":       reflect.TypeOf((*q.File)(nil)).Elem(),
				"Ident":      reflect.TypeOf((*q.Ident)(nil)).Elem(),
				"Rule":       reflect.TypeOf((*q.Rule)(nil)).Elem(),
				"Sequence":   reflect.TypeOf((*q.Sequence)(nil)).Elem(),
				"UnaryExpr":  reflect.TypeOf((*q.UnaryExpr)(nil)).Elem(),
			},
			AliasTypes:    map[string]reflect.Type{},
			Vars:          map[string]interface{}{},
			Funcs:         map[string]interface{}{},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
