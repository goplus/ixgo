// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package rsa

import (
	q "crypto/rsa"

	crypto "crypto"
	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/rsa", map[string]ixgo.DirectCallAdapter{
		"(*PSSOptions).HashFunc":    method_ptr_PSSOptions_HashFunc,
		"(*PrivateKey).Equal":       method_ptr_PrivateKey_Equal,
		"(*PrivateKey).Precompute":  method_ptr_PrivateKey_Precompute,
		"(*PrivateKey).Public":      method_ptr_PrivateKey_Public,
		"(*PrivateKey).Validate":    method_ptr_PrivateKey_Validate,
		"(*PublicKey).Equal":        method_ptr_PublicKey_Equal,
		"(*PublicKey).Size":         method_ptr_PublicKey_Size,
		"DecryptPKCS1v15SessionKey": func_DecryptPKCS1v15SessionKey,
		"VerifyPKCS1v15":            func_VerifyPKCS1v15,
		"VerifyPSS":                 func_VerifyPSS,
	})
}

func func_DecryptPKCS1v15SessionKey(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DecryptPKCS1v15SessionKey(ixgo.DirectCallArg[io.Reader](ctx, 0), ixgo.DirectCallArg[*q.PrivateKey](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2), ixgo.DirectCallArg[[]byte](ctx, 3)))
}

func method_ptr_PSSOptions_HashFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PSSOptions).HashFunc(ixgo.DirectCallArg[*q.PSSOptions](ctx, 0)))
}

func method_ptr_PrivateKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Equal(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0), ixgo.DirectCallArg[crypto.PrivateKey](ctx, 1)))
}

func method_ptr_PrivateKey_Precompute(ctx ixgo.DirectCallContext) {
	(*q.PrivateKey).Precompute(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0))
}

func method_ptr_PrivateKey_Public(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Public(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
}

func method_ptr_PrivateKey_Validate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Validate(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
}

func method_ptr_PublicKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKey).Equal(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[crypto.PublicKey](ctx, 1)))
}

func method_ptr_PublicKey_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKey).Size(ixgo.DirectCallArg[*q.PublicKey](ctx, 0)))
}

func func_VerifyPKCS1v15(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VerifyPKCS1v15(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[crypto.Hash](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2), ixgo.DirectCallArg[[]byte](ctx, 3)))
}

func func_VerifyPSS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VerifyPSS(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[crypto.Hash](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2), ixgo.DirectCallArg[[]byte](ctx, 3), ixgo.DirectCallArg[*q.PSSOptions](ctx, 4)))
}
