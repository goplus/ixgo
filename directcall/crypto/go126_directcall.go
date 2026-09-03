// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package crypto

import (
	q "crypto"

	"github.com/goplus/ixgo"
	hash "hash"
)

func init() {
	ixgo.RegisterDirectCalls("crypto", map[string]ixgo.DirectCallAdapter{
		"(*Hash).Available": method_ptr_Hash_Available,
		"(*Hash).HashFunc":  method_ptr_Hash_HashFunc,
		"(*Hash).New":       method_ptr_Hash_New,
		"(*Hash).Size":      method_ptr_Hash_Size,
		"(*Hash).String":    method_ptr_Hash_String,
		"(Hash).Available":  method_Hash_Available,
		"(Hash).HashFunc":   method_Hash_HashFunc,
		"(Hash).New":        method_Hash_New,
		"(Hash).Size":       method_Hash_Size,
		"(Hash).String":     method_Hash_String,
		"RegisterHash":      func_RegisterHash,
	})
}

func method_Hash_Available(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Hash.Available(ixgo.DirectCallArg[q.Hash](ctx, 0)))
}

func method_ptr_Hash_Available(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).Available(ixgo.DirectCallArg[*q.Hash](ctx, 0)))
}

func method_Hash_HashFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Hash.HashFunc(ixgo.DirectCallArg[q.Hash](ctx, 0)))
}

func method_ptr_Hash_HashFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).HashFunc(ixgo.DirectCallArg[*q.Hash](ctx, 0)))
}

func method_Hash_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Hash.New(ixgo.DirectCallArg[q.Hash](ctx, 0)))
}

func method_ptr_Hash_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).New(ixgo.DirectCallArg[*q.Hash](ctx, 0)))
}

func method_Hash_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Hash.Size(ixgo.DirectCallArg[q.Hash](ctx, 0)))
}

func method_ptr_Hash_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).Size(ixgo.DirectCallArg[*q.Hash](ctx, 0)))
}

func method_Hash_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Hash.String(ixgo.DirectCallArg[q.Hash](ctx, 0)))
}

func method_ptr_Hash_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hash).String(ixgo.DirectCallArg[*q.Hash](ctx, 0)))
}

func func_RegisterHash(ctx ixgo.DirectCallContext) {
	q.RegisterHash(ixgo.DirectCallArg[q.Hash](ctx, 0), ixgo.DirectCallArg[func() hash.Hash](ctx, 1))
}
