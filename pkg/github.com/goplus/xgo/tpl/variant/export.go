// export by github.com/goplus/ixgo/cmd/qexp

package variant

import (
	q "github.com/goplus/xgo/tpl/variant"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/variant", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "variant",
			Path: "github.com/goplus/xgo/tpl/variant",
			Deps: map[string]string{
				"github.com/goplus/xgo/tpl":       "tpl",
				"github.com/goplus/xgo/tpl/token": "token",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Module": reflect.TypeOf((*q.Module)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"DelayValue": reflect.TypeOf((*q.DelayValue)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Call":         q.Call,
				"CallObject":   q.CallObject,
				"Compare":      q.Compare,
				"Eval":         q.Eval,
				"Float":        q.Float,
				"InitUniverse": q.InitUniverse,
				"Int":          q.Int,
				"List":         q.List,
				"LogicOp":      q.LogicOp,
				"MathOp":       q.MathOp,
				"NewModule":    q.NewModule,
				"RangeOp":      q.RangeOp,
				"UnaryOp":      q.UnaryOp,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
