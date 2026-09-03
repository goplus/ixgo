// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package ed25519

import (
	q "crypto/ed25519"

	crypto "crypto"
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/ed25519", map[string]ixgo.DirectCallAdapter{
		"(*Options).HashFunc":  method_ptr_Options_HashFunc,
		"(*PrivateKey).Equal":  method_ptr_PrivateKey_Equal,
		"(*PrivateKey).Public": method_ptr_PrivateKey_Public,
		"(*PrivateKey).Seed":   method_ptr_PrivateKey_Seed,
		"(*PublicKey).Equal":   method_ptr_PublicKey_Equal,
		"(PrivateKey).Equal":   method_PrivateKey_Equal,
		"(PrivateKey).Public":  method_PrivateKey_Public,
		"(PrivateKey).Seed":    method_PrivateKey_Seed,
		"(PublicKey).Equal":    method_PublicKey_Equal,
		"NewKeyFromSeed":       func_NewKeyFromSeed,
		"Sign":                 func_Sign,
		"Verify":               func_Verify,
		"VerifyWithOptions":    func_VerifyWithOptions,
	})
}

func func_NewKeyFromSeed(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewKeyFromSeed(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func method_ptr_Options_HashFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Options).HashFunc(ixgo.DirectCallArg[*q.Options](ctx, 0)))
}

func method_PrivateKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PrivateKey.Equal(ixgo.DirectCallArg[q.PrivateKey](ctx, 0), ixgo.DirectCallArg[crypto.PrivateKey](ctx, 1)))
}

func method_ptr_PrivateKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Equal(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0), ixgo.DirectCallArg[crypto.PrivateKey](ctx, 1)))
}

func method_PrivateKey_Public(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PrivateKey.Public(ixgo.DirectCallArg[q.PrivateKey](ctx, 0)))
}

func method_ptr_PrivateKey_Public(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Public(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
}

func method_PrivateKey_Seed(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PrivateKey.Seed(ixgo.DirectCallArg[q.PrivateKey](ctx, 0)))
}

func method_ptr_PrivateKey_Seed(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PrivateKey).Seed(ixgo.DirectCallArg[*q.PrivateKey](ctx, 0)))
}

func method_PublicKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PublicKey.Equal(ixgo.DirectCallArg[q.PublicKey](ctx, 0), ixgo.DirectCallArg[crypto.PublicKey](ctx, 1)))
}

func method_ptr_PublicKey_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PublicKey).Equal(ixgo.DirectCallArg[*q.PublicKey](ctx, 0), ixgo.DirectCallArg[crypto.PublicKey](ctx, 1)))
}

func func_Sign(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sign(ixgo.DirectCallArg[q.PrivateKey](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_Verify(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Verify(ixgo.DirectCallArg[q.PublicKey](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}

func func_VerifyWithOptions(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VerifyWithOptions(ixgo.DirectCallArg[q.PublicKey](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2), ixgo.DirectCallArg[*q.Options](ctx, 3)))
}
