// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package uuid

import (
	q "uuid"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("uuid", map[string]ixgo.DirectCallAdapter{
		"(*UUID).Compare":       method_ptr_UUID_Compare,
		"(*UUID).String":        method_ptr_UUID_String,
		"(*UUID).UnmarshalText": method_ptr_UUID_UnmarshalText,
		"(UUID).Compare":        method_UUID_Compare,
		"(UUID).String":         method_UUID_String,
		"Max":                   func_Max,
		"MustParse":             func_MustParse,
		"New":                   func_New,
		"NewV4":                 func_NewV4,
		"NewV7":                 func_NewV7,
		"Nil":                   func_Nil,
	})
}

func func_Max(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Max())
}

func func_MustParse(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MustParse(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New())
}

func func_NewV4(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewV4())
}

func func_NewV7(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewV7())
}

func func_Nil(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Nil())
}

func method_UUID_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UUID.Compare(ixgo.DirectCallArg[q.UUID](ctx, 0), ixgo.DirectCallArg[q.UUID](ctx, 1)))
}

func method_ptr_UUID_Compare(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UUID).Compare(ixgo.DirectCallArg[*q.UUID](ctx, 0), ixgo.DirectCallArg[q.UUID](ctx, 1)))
}

func method_UUID_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UUID.String(ixgo.DirectCallArg[q.UUID](ctx, 0)))
}

func method_ptr_UUID_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UUID).String(ixgo.DirectCallArg[*q.UUID](ctx, 0)))
}

func method_ptr_UUID_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UUID).UnmarshalText(ixgo.DirectCallArg[*q.UUID](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}
