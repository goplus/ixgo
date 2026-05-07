// export by github.com/goplus/ixgo/cmd/qexp

package delay

import (
	q "github.com/goplus/xgo/tpl/variant/delay"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/variant/delay", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "delay",
			Path: "github.com/goplus/xgo/tpl/variant/delay",
			Deps: map[string]string{
				"github.com/goplus/xgo/tpl/token":   "token",
				"github.com/goplus/xgo/tpl/variant": "variant",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{
				"Value": reflect.TypeOf((*q.Value)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Call":         q.Call,
				"CallObject":   q.CallObject,
				"ChgValue":     q.ChgValue,
				"Compare":      q.Compare,
				"Eval":         q.Eval,
				"EvalOp":       q.EvalOp,
				"IfElse":       q.IfElse,
				"InitUniverse": q.InitUniverse,
				"List":         q.List,
				"LogicOp":      q.LogicOp,
				"MathOp":       q.MathOp,
				"RangeOp":      q.RangeOp,
				"RepeatUntil":  q.RepeatUntil,
				"SetValue":     q.SetValue,
				"StmtList":     q.StmtList,
				"UnaryOp":      q.UnaryOp,
				"ValueOf":      q.ValueOf,
				"While":        q.While,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
