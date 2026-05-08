// export by github.com/goplus/ixgo/cmd/qexp

package variant

import (
	q "github.com/goplus/xgo/tpl/variant"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("github.com/goplus/xgo/tpl/variant", func() *ixgo.Package {
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
			Vars: map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"Call":         reflect.ValueOf(q.Call),
				"CallObject":   reflect.ValueOf(q.CallObject),
				"Compare":      reflect.ValueOf(q.Compare),
				"Eval":         reflect.ValueOf(q.Eval),
				"Float":        reflect.ValueOf(q.Float),
				"InitUniverse": reflect.ValueOf(q.InitUniverse),
				"Int":          reflect.ValueOf(q.Int),
				"List":         reflect.ValueOf(q.List),
				"LogicOp":      reflect.ValueOf(q.LogicOp),
				"MathOp":       reflect.ValueOf(q.MathOp),
				"NewModule":    reflect.ValueOf(q.NewModule),
				"RangeOp":      reflect.ValueOf(q.RangeOp),
				"UnaryOp":      reflect.ValueOf(q.UnaryOp),
			},
			TypedConsts:   map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
