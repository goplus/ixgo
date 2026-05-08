// export by github.com/goplus/ixgo/cmd/qexp

package delay

import (
	q "github.com/goplus/xgo/tpl/variant/delay"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("github.com/goplus/xgo/tpl/variant/delay", func() *ixgo.Package {
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
			Vars: map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"Call":         reflect.ValueOf(q.Call),
				"CallObject":   reflect.ValueOf(q.CallObject),
				"ChgValue":     reflect.ValueOf(q.ChgValue),
				"Compare":      reflect.ValueOf(q.Compare),
				"Eval":         reflect.ValueOf(q.Eval),
				"EvalOp":       reflect.ValueOf(q.EvalOp),
				"IfElse":       reflect.ValueOf(q.IfElse),
				"InitUniverse": reflect.ValueOf(q.InitUniverse),
				"List":         reflect.ValueOf(q.List),
				"LogicOp":      reflect.ValueOf(q.LogicOp),
				"MathOp":       reflect.ValueOf(q.MathOp),
				"RangeOp":      reflect.ValueOf(q.RangeOp),
				"RepeatUntil":  reflect.ValueOf(q.RepeatUntil),
				"SetValue":     reflect.ValueOf(q.SetValue),
				"StmtList":     reflect.ValueOf(q.StmtList),
				"UnaryOp":      reflect.ValueOf(q.UnaryOp),
				"ValueOf":      reflect.ValueOf(q.ValueOf),
				"While":        reflect.ValueOf(q.While),
			},
			TypedConsts:   map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
