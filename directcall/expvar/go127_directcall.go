// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package expvar

import (
	q "expvar"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("expvar", map[string]ixgo.DirectCallAdapter{
		"(*Float).Add":     method_ptr_Float_Add,
		"(*Float).Set":     method_ptr_Float_Set,
		"(*Float).String":  method_ptr_Float_String,
		"(*Float).Value":   method_ptr_Float_Value,
		"(*Func).String":   method_ptr_Func_String,
		"(*Func).Value":    method_ptr_Func_Value,
		"(*Int).Add":       method_ptr_Int_Add,
		"(*Int).Set":       method_ptr_Int_Set,
		"(*Int).String":    method_ptr_Int_String,
		"(*Int).Value":     method_ptr_Int_Value,
		"(*Map).Add":       method_ptr_Map_Add,
		"(*Map).AddFloat":  method_ptr_Map_AddFloat,
		"(*Map).Delete":    method_ptr_Map_Delete,
		"(*Map).Do":        method_ptr_Map_Do,
		"(*Map).Get":       method_ptr_Map_Get,
		"(*Map).Init":      method_ptr_Map_Init,
		"(*Map).Set":       method_ptr_Map_Set,
		"(*Map).String":    method_ptr_Map_String,
		"(*String).Set":    method_ptr_String_Set,
		"(*String).String": method_ptr_String_String,
		"(*String).Value":  method_ptr_String_Value,
		"(Func).String":    method_Func_String,
		"(Func).Value":     method_Func_Value,
		"Do":               func_Do,
		"Get":              func_Get,
		"Handler":          func_Handler,
		"NewFloat":         func_NewFloat,
		"NewInt":           func_NewInt,
		"NewMap":           func_NewMap,
		"NewString":        func_NewString,
		"Publish":          func_Publish,
	})
}

func func_Do(ctx ixgo.DirectCallContext) {
	q.Do(ixgo.DirectCallArg[func(q.KeyValue)](ctx, 0))
}

func method_ptr_Float_Add(ctx ixgo.DirectCallContext) {
	(*q.Float).Add(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1))
}

func method_ptr_Float_Set(ctx ixgo.DirectCallContext) {
	(*q.Float).Set(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1))
}

func method_ptr_Float_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).String(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_Value(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Value(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_Func_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Func.String(ixgo.DirectCallArg[q.Func](ctx, 0)))
}

func method_ptr_Func_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).String(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func method_Func_Value(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Func.Value(ixgo.DirectCallArg[q.Func](ctx, 0)))
}

func method_ptr_Func_Value(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).Value(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func func_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Get(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Handler(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Handler())
}

func method_ptr_Int_Add(ctx ixgo.DirectCallContext) {
	(*q.Int).Add(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1))
}

func method_ptr_Int_Set(ctx ixgo.DirectCallContext) {
	(*q.Int).Set(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1))
}

func method_ptr_Int_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).String(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_Value(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Value(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Map_Add(ctx ixgo.DirectCallContext) {
	(*q.Map).Add(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2))
}

func method_ptr_Map_AddFloat(ctx ixgo.DirectCallContext) {
	(*q.Map).AddFloat(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[float64](ctx, 2))
}

func method_ptr_Map_Delete(ctx ixgo.DirectCallContext) {
	(*q.Map).Delete(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_Map_Do(ctx ixgo.DirectCallContext) {
	(*q.Map).Do(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[func(q.KeyValue)](ctx, 1))
}

func method_ptr_Map_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Map).Get(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Map_Init(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Map).Init(ixgo.DirectCallArg[*q.Map](ctx, 0)))
}

func method_ptr_Map_Set(ctx ixgo.DirectCallContext) {
	(*q.Map).Set(ixgo.DirectCallArg[*q.Map](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[q.Var](ctx, 2))
}

func method_ptr_Map_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Map).String(ixgo.DirectCallArg[*q.Map](ctx, 0)))
}

func func_NewFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewFloat(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_NewInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewInt(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_NewMap(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewMap(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_NewString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewString(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Publish(ctx ixgo.DirectCallContext) {
	q.Publish(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[q.Var](ctx, 1))
}

func method_ptr_String_Set(ctx ixgo.DirectCallContext) {
	(*q.String).Set(ixgo.DirectCallArg[*q.String](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_String_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.String).String(ixgo.DirectCallArg[*q.String](ctx, 0)))
}

func method_ptr_String_Value(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.String).Value(ixgo.DirectCallArg[*q.String](ctx, 0)))
}
