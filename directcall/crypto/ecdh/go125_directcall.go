// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package ecdh

import (
	q "crypto/ecdh"

	crypto "crypto"
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/ecdh", map[string]ixgo.DirectCallAdapter{
		"(*PrivateKey).Bytes":     method_ptr_PrivateKey_Bytes,
		"(*PrivateKey).Curve":     method_ptr_PrivateKey_Curve,
		"(*PrivateKey).Equal":     method_ptr_PrivateKey_Equal,
		"(*PrivateKey).Public":    method_ptr_PrivateKey_Public,
		"(*PrivateKey).PublicKey": method_ptr_PrivateKey_PublicKey,
		"(*PublicKey).Bytes":      method_ptr_PublicKey_Bytes,
		"(*PublicKey).Curve":      method_ptr_PublicKey_Curve,
		"(*PublicKey).Equal":      method_ptr_PublicKey_Equal,
		"P256":                    func_P256,
		"P384":                    func_P384,
		"P521":                    func_P521,
		"X25519":                  func_X25519,
	})
}

func func_P256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.P256())
}

func func_P384(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.P384())
}

func func_P521(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.P521())
}

func method_ptr_PrivateKey_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Bytes(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
}

func method_ptr_PrivateKey_Curve(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Curve(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
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

func method_ptr_PublicKey_Curve(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKey).Curve(ixgo.DirectCallArg[*q.PublicKey](ctx, 0)))
}

func method_ptr_PublicKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKey).Equal(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[crypto.PublicKey](ctx, 1)))
}

func func_X25519(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.X25519())
}
