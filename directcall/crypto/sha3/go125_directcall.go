// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package sha3

import (
	q "crypto/sha3"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/sha3", map[string]ixgo.DirectCallAdapter{
		"(*SHA3).BlockSize":        method_ptr_SHA3_BlockSize,
		"(*SHA3).Reset":            method_ptr_SHA3_Reset,
		"(*SHA3).Size":             method_ptr_SHA3_Size,
		"(*SHA3).Sum":              method_ptr_SHA3_Sum,
		"(*SHA3).UnmarshalBinary":  method_ptr_SHA3_UnmarshalBinary,
		"(*SHAKE).BlockSize":       method_ptr_SHAKE_BlockSize,
		"(*SHAKE).Reset":           method_ptr_SHAKE_Reset,
		"(*SHAKE).UnmarshalBinary": method_ptr_SHAKE_UnmarshalBinary,
		"New224":                   func_New224,
		"New256":                   func_New256,
		"New384":                   func_New384,
		"New512":                   func_New512,
		"NewCSHAKE128":             func_NewCSHAKE128,
		"NewCSHAKE256":             func_NewCSHAKE256,
		"NewSHAKE128":              func_NewSHAKE128,
		"NewSHAKE256":              func_NewSHAKE256,
		"Sum224":                   func_Sum224,
		"Sum256":                   func_Sum256,
		"Sum384":                   func_Sum384,
		"Sum512":                   func_Sum512,
		"SumSHAKE128":              func_SumSHAKE128,
		"SumSHAKE256":              func_SumSHAKE256,
	})
}

func func_New224(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New224())
}

func func_New256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New256())
}

func func_New384(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New384())
}

func func_New512(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New512())
}

func func_NewCSHAKE128(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCSHAKE128(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_NewCSHAKE256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCSHAKE256(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_NewSHAKE128(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewSHAKE128())
}

func func_NewSHAKE256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewSHAKE256())
}

func method_ptr_SHA3_BlockSize(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SHA3).BlockSize(ixgo.DirectCallArg[*q.SHA3](ctx, 0)))
}

func method_ptr_SHA3_Reset(ctx ixgo.DirectCallContext) {
	(*q.SHA3).Reset(ixgo.DirectCallArg[*q.SHA3](ctx, 0))
}

func method_ptr_SHA3_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SHA3).Size(ixgo.DirectCallArg[*q.SHA3](ctx, 0)))
}

func method_ptr_SHA3_Sum(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SHA3).Sum(ixgo.DirectCallArg[*q.SHA3](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_SHA3_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SHA3).UnmarshalBinary(ixgo.DirectCallArg[*q.SHA3](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_SHAKE_BlockSize(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SHAKE).BlockSize(ixgo.DirectCallArg[*q.SHAKE](ctx, 0)))
}

func method_ptr_SHAKE_Reset(ctx ixgo.DirectCallContext) {
	(*q.SHAKE).Reset(ixgo.DirectCallArg[*q.SHAKE](ctx, 0))
}

func method_ptr_SHAKE_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SHAKE).UnmarshalBinary(ixgo.DirectCallArg[*q.SHAKE](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_Sum224(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum224(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_Sum256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum256(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_Sum384(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum384(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_Sum512(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum512(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_SumSHAKE128(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SumSHAKE128(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_SumSHAKE256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SumSHAKE256(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}
