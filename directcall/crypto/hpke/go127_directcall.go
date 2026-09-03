// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package hpke

import (
	q "crypto/hpke"

	ecdh "crypto/ecdh"
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/hpke", map[string]ixgo.DirectCallAdapter{
		"AES128GCM":        func_AES128GCM,
		"AES256GCM":        func_AES256GCM,
		"ChaCha20Poly1305": func_ChaCha20Poly1305,
		"DHKEM":            func_DHKEM,
		"ExportOnly":       func_ExportOnly,
		"HKDFSHA256":       func_HKDFSHA256,
		"HKDFSHA384":       func_HKDFSHA384,
		"HKDFSHA512":       func_HKDFSHA512,
		"MLKEM1024":        func_MLKEM1024,
		"MLKEM1024P384":    func_MLKEM1024P384,
		"MLKEM768":         func_MLKEM768,
		"MLKEM768P256":     func_MLKEM768P256,
		"MLKEM768X25519":   func_MLKEM768X25519,
		"SHAKE128":         func_SHAKE128,
		"SHAKE256":         func_SHAKE256,
	})
}

func func_AES128GCM(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AES128GCM())
}

func func_AES256GCM(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AES256GCM())
}

func func_ChaCha20Poly1305(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ChaCha20Poly1305())
}

func func_DHKEM(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DHKEM(ixgo.DirectCallArg[ecdh.Curve](ctx, 0)))
}

func func_ExportOnly(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ExportOnly())
}

func func_HKDFSHA256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HKDFSHA256())
}

func func_HKDFSHA384(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HKDFSHA384())
}

func func_HKDFSHA512(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HKDFSHA512())
}

func func_MLKEM1024(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MLKEM1024())
}

func func_MLKEM1024P384(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MLKEM1024P384())
}

func func_MLKEM768(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MLKEM768())
}

func func_MLKEM768P256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MLKEM768P256())
}

func func_MLKEM768X25519(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MLKEM768X25519())
}

func func_SHAKE128(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SHAKE128())
}

func func_SHAKE256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SHAKE256())
}
