// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package maphash

import (
	q "hash/maphash"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("hash/maphash", map[string]ixgo.DirectCallAdapter{
		"(*Hash).BlockSize": method_ptr_Hash_BlockSize,
		"(*Hash).Reset":     method_ptr_Hash_Reset,
		"(*Hash).Seed":      method_ptr_Hash_Seed,
		"(*Hash).SetSeed":   method_ptr_Hash_SetSeed,
		"(*Hash).Size":      method_ptr_Hash_Size,
		"(*Hash).Sum":       method_ptr_Hash_Sum,
		"(*Hash).Sum64":     method_ptr_Hash_Sum64,
		"(*Hash).WriteByte": method_ptr_Hash_WriteByte,
		"Bytes":             func_Bytes,
		"MakeSeed":          func_MakeSeed,
		"String":            func_String,
	})
}

func func_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Bytes(ixgo.DirectCallArg[q.Seed](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Hash_BlockSize(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).BlockSize(ixgo.DirectCallArg[*q.Hash](ctx, 0)))
}

func method_ptr_Hash_Reset(ctx ixgo.DirectCallContext) {
	(*q.Hash).Reset(ixgo.DirectCallArg[*q.Hash](ctx, 0))
}

func method_ptr_Hash_Seed(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).Seed(ixgo.DirectCallArg[*q.Hash](ctx, 0)))
}

func method_ptr_Hash_SetSeed(ctx ixgo.DirectCallContext) {
	(*q.Hash).SetSeed(ixgo.DirectCallArg[*q.Hash](ctx, 0), ixgo.DirectCallArg[q.Seed](ctx, 1))
}

func method_ptr_Hash_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).Size(ixgo.DirectCallArg[*q.Hash](ctx, 0)))
}

func method_ptr_Hash_Sum(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).Sum(ixgo.DirectCallArg[*q.Hash](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Hash_Sum64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).Sum64(ixgo.DirectCallArg[*q.Hash](ctx, 0)))
}

func method_ptr_Hash_WriteByte(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).WriteByte(ixgo.DirectCallArg[*q.Hash](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1)))
}

func func_MakeSeed(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeSeed())
}

func func_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.String(ixgo.DirectCallArg[q.Seed](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}
