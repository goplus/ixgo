// export by github.com/goplus/ixgo/cmd/qexp

package pkg

import (
	q "github.com/goplus/ixgo/testdata/direct_call/pkg"

	"github.com/goplus/ixgo"
	"go/constant"
	"reflect"
)

func init() {
	ixgo.RegisterPackageLazy("github.com/goplus/ixgo/testdata/direct_call/pkg", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "pkg",
			Path: "github.com/goplus/ixgo/testdata/direct_call/pkg",
			Deps: map[string]string{
				"go/constant": "constant",
				"reflect":     "reflect",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Number": reflect.TypeOf((*q.Number)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"Add":     reflect.ValueOf(q.Add),
				"Inspect": reflect.ValueOf(q.Inspect),
			},
			TypedConsts: map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Marker": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Marker))},
			},
		}
	})
	ixgo.RegisterDirectCalls("github.com/goplus/ixgo/testdata/direct_call/pkg", map[string]ixgo.DirectCallBinding{
		"(*Number).Value": {Target: reflect.ValueOf((*q.Number).Value), Adapter: method_ptr_Number_Value},
		"Add":             {Target: reflect.ValueOf(q.Add), Adapter: func_Add},
		"Inspect":         {Target: reflect.ValueOf(q.Inspect), Adapter: func_Inspect},
		"Number.Value":    {Target: reflect.ValueOf(q.Number.Value), Adapter: method_Number_Value},
	})
}

func func_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Add(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_Inspect(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Inspect(ixgo.DirectCallArg[reflect.Type](ctx, 0), ixgo.DirectCallArg[constant.Value](ctx, 1)))
}

func method_Number_Value(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Number.Value(ixgo.DirectCallArg[q.Number](ctx, 0)))
}

func method_ptr_Number_Value(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Number).Value(ixgo.DirectCallArg[*q.Number](ctx, 0)))
}
