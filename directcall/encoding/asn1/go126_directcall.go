// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package asn1

import (
	q "encoding/asn1"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/asn1", map[string]ixgo.DirectCallAdapter{
		"(*BitString).At":            method_ptr_BitString_At,
		"(*BitString).RightAlign":    method_ptr_BitString_RightAlign,
		"(*ObjectIdentifier).Equal":  method_ptr_ObjectIdentifier_Equal,
		"(*ObjectIdentifier).String": method_ptr_ObjectIdentifier_String,
		"(*StructuralError).Error":   method_ptr_StructuralError_Error,
		"(*SyntaxError).Error":       method_ptr_SyntaxError_Error,
		"(BitString).At":             method_BitString_At,
		"(BitString).RightAlign":     method_BitString_RightAlign,
		"(ObjectIdentifier).Equal":   method_ObjectIdentifier_Equal,
		"(ObjectIdentifier).String":  method_ObjectIdentifier_String,
		"(StructuralError).Error":    method_StructuralError_Error,
		"(SyntaxError).Error":        method_SyntaxError_Error,
	})
}

func method_BitString_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BitString.At(ixgo.DirectCallArg[q.BitString](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_BitString_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BitString).At(ixgo.DirectCallArg[*q.BitString](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_BitString_RightAlign(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.BitString.RightAlign(ixgo.DirectCallArg[q.BitString](ctx, 0)))
}

func method_ptr_BitString_RightAlign(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BitString).RightAlign(ixgo.DirectCallArg[*q.BitString](ctx, 0)))
}

func method_ObjectIdentifier_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ObjectIdentifier.Equal(ixgo.DirectCallArg[q.ObjectIdentifier](ctx, 0), ixgo.DirectCallArg[q.ObjectIdentifier](ctx, 1)))
}

func method_ptr_ObjectIdentifier_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ObjectIdentifier).Equal(ixgo.DirectCallArg[*q.ObjectIdentifier](ctx, 0), ixgo.DirectCallArg[q.ObjectIdentifier](ctx, 1)))
}

func method_ObjectIdentifier_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ObjectIdentifier.String(ixgo.DirectCallArg[q.ObjectIdentifier](ctx, 0)))
}

func method_ptr_ObjectIdentifier_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ObjectIdentifier).String(ixgo.DirectCallArg[*q.ObjectIdentifier](ctx, 0)))
}

func method_StructuralError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StructuralError.Error(ixgo.DirectCallArg[q.StructuralError](ctx, 0)))
}

func method_ptr_StructuralError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StructuralError).Error(ixgo.DirectCallArg[*q.StructuralError](ctx, 0)))
}

func method_SyntaxError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SyntaxError.Error(ixgo.DirectCallArg[q.SyntaxError](ctx, 0)))
}

func method_ptr_SyntaxError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SyntaxError).Error(ixgo.DirectCallArg[*q.SyntaxError](ctx, 0)))
}
