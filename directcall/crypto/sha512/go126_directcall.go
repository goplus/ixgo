// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package sha512

import (
	q "crypto/sha512"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/sha512", map[string]ixgo.DirectCallAdapter{
		"New":        func_New,
		"New384":     func_New384,
		"New512_224": func_New512_224,
		"New512_256": func_New512_256,
		"Sum384":     func_Sum384,
		"Sum512":     func_Sum512,
		"Sum512_224": func_Sum512_224,
		"Sum512_256": func_Sum512_256,
	})
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New())
}

func func_New384(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New384())
}

func func_New512_224(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New512_224())
}

func func_New512_256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New512_256())
}

func func_Sum384(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum384(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_Sum512(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum512(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_Sum512_224(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum512_224(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_Sum512_256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum512_256(ixgo.DirectCallArg[[]byte](ctx, 0)))
}
