// export by github.com/goplus/ixgo/cmd/qexp

package matcher

import (
	q "github.com/goplus/xgo/tpl/matcher"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/matcher", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "matcher",
			Path: "github.com/goplus/xgo/tpl/matcher",
			Deps: map[string]string{
				"errors":                          "errors",
				"fmt":                             "fmt",
				"github.com/goplus/xgo/tpl/token": "token",
				"github.com/goplus/xgo/tpl/types": "types",
				"log":                             "log",
			},
			Interfaces: map[string]reflect.Type{
				"Matcher": reflect.TypeOf((*q.Matcher)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Choices":        reflect.TypeOf((*q.Choices)(nil)).Elem(),
				"Context":        reflect.TypeOf((*q.Context)(nil)).Elem(),
				"Error":          reflect.TypeOf((*q.Error)(nil)).Elem(),
				"MatchToken":     reflect.TypeOf((*q.MatchToken)(nil)).Elem(),
				"RecursiveError": reflect.TypeOf((*q.RecursiveError)(nil)).Elem(),
				"Var":            reflect.TypeOf((*q.Var)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"ListRetProc": reflect.TypeOf((*q.ListRetProc)(nil)).Elem(),
				"RetProc":     reflect.TypeOf((*q.RetProc)(nil)).Elem(),
			},
			Vars: map[string]interface{}{
				"ErrVarAssigned": &q.ErrVarAssigned,
			},
			Funcs: map[string]interface{}{
				"Adjoin":     q.Adjoin,
				"Choice":     q.Choice,
				"List":       q.List,
				"Literal":    q.Literal,
				"NewContext": q.NewContext,
				"NewVar":     q.NewVar,
				"Repeat0":    q.Repeat0,
				"Repeat01":   q.Repeat01,
				"Repeat1":    q.Repeat1,
				"Sequence":   q.Sequence,
				"SetDebug":   q.SetDebug,
				"String":     q.String,
				"Token":      q.Token,
				"True":       q.True,
				"WhiteSpace": q.WhiteSpace,
			},
			TypedConsts: map[string]interface{}{
				"DbgFlagAll":      q.DbgFlagAll,
				"DbgFlagMatchVar": q.DbgFlagMatchVar,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
