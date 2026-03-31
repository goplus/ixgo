// export by github.com/goplus/ixgo/cmd/qexp

package matcher

import (
	q "github.com/goplus/xgo/tpl/matcher"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"ErrVarAssigned": reflect.ValueOf(&q.ErrVarAssigned),
		},
		Funcs: map[string]reflect.Value{
			"Adjoin":     reflect.ValueOf(q.Adjoin),
			"Choice":     reflect.ValueOf(q.Choice),
			"List":       reflect.ValueOf(q.List),
			"Literal":    reflect.ValueOf(q.Literal),
			"NewContext": reflect.ValueOf(q.NewContext),
			"NewVar":     reflect.ValueOf(q.NewVar),
			"Repeat0":    reflect.ValueOf(q.Repeat0),
			"Repeat01":   reflect.ValueOf(q.Repeat01),
			"Repeat1":    reflect.ValueOf(q.Repeat1),
			"Sequence":   reflect.ValueOf(q.Sequence),
			"SetDebug":   reflect.ValueOf(q.SetDebug),
			"String":     reflect.ValueOf(q.String),
			"Token":      reflect.ValueOf(q.Token),
			"True":       reflect.ValueOf(q.True),
			"WhiteSpace": reflect.ValueOf(q.WhiteSpace),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"DbgFlagAll":      {Typ: reflect.TypeOf(q.DbgFlagAll), Value: constant.MakeInt64(int64(q.DbgFlagAll))},
			"DbgFlagMatchVar": {Typ: reflect.TypeOf(q.DbgFlagMatchVar), Value: constant.MakeInt64(int64(q.DbgFlagMatchVar))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
