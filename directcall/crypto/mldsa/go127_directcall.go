// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package mldsa

import (
	q "crypto/mldsa"

	crypto "crypto"
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/mldsa", map[string]ixgo.DirectCallAdapter{
		"(*Options).HashFunc":         method_ptr_Options_HashFunc,
		"(*Parameters).PublicKeySize": method_ptr_Parameters_PublicKeySize,
		"(*Parameters).SignatureSize": method_ptr_Parameters_SignatureSize,
		"(*Parameters).String":        method_ptr_Parameters_String,
		"(*PrivateKey).Bytes":         method_ptr_PrivateKey_Bytes,
		"(*PrivateKey).Equal":         method_ptr_PrivateKey_Equal,
		"(*PrivateKey).Public":        method_ptr_PrivateKey_Public,
		"(*PrivateKey).PublicKey":     method_ptr_PrivateKey_PublicKey,
		"(*PublicKey).Bytes":          method_ptr_PublicKey_Bytes,
		"(*PublicKey).Equal":          method_ptr_PublicKey_Equal,
		"(*PublicKey).Parameters":     method_ptr_PublicKey_Parameters,
		"(Parameters).PublicKeySize":  method_Parameters_PublicKeySize,
		"(Parameters).SignatureSize":  method_Parameters_SignatureSize,
		"(Parameters).String":         method_Parameters_String,
		"MLDSA44":                     func_MLDSA44,
		"MLDSA65":                     func_MLDSA65,
		"MLDSA87":                     func_MLDSA87,
		"Verify":                      func_Verify,
	})
}

func func_MLDSA44(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MLDSA44())
}

func func_MLDSA65(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MLDSA65())
}

func func_MLDSA87(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MLDSA87())
}

func method_ptr_Options_HashFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Options).HashFunc(ixgo.DirectCallArg[*q.Options](ctx, 0)))
}

func method_Parameters_PublicKeySize(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Parameters.PublicKeySize(ixgo.DirectCallArg[q.Parameters](ctx, 0)))
}

func method_ptr_Parameters_PublicKeySize(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Parameters).PublicKeySize(ixgo.DirectCallArg[*q.Parameters](ctx, 0)))
}

func method_Parameters_SignatureSize(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Parameters.SignatureSize(ixgo.DirectCallArg[q.Parameters](ctx, 0)))
}

func method_ptr_Parameters_SignatureSize(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Parameters).SignatureSize(ixgo.DirectCallArg[*q.Parameters](ctx, 0)))
}

func method_Parameters_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Parameters.String(ixgo.DirectCallArg[q.Parameters](ctx, 0)))
}

func method_ptr_Parameters_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Parameters).String(ixgo.DirectCallArg[*q.Parameters](ctx, 0)))
}

func method_ptr_PrivateKey_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Bytes(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
}

func method_ptr_PrivateKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Equal(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0), ixgo.DirectCallArg[crypto.PrivateKey](ctx, 1)))
}

func method_ptr_PrivateKey_Public(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Public(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
}

func method_ptr_PrivateKey_PublicKey(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).PublicKey(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
}

func method_ptr_PublicKey_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKey).Bytes(ixgo.DirectCallArg[*q.PublicKey](ctx, 0)))
}

func method_ptr_PublicKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKey).Equal(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[crypto.PublicKey](ctx, 1)))
}

func method_ptr_PublicKey_Parameters(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKey).Parameters(ixgo.DirectCallArg[*q.PublicKey](ctx, 0)))
}

func func_Verify(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Verify(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2), ixgo.DirectCallArg[*q.Options](ctx, 3)))
}
