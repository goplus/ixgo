// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package sha256

import (
	q "crypto/sha256"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/sha256", map[string]ixgo.DirectCallAdapter{
		"New":    func_New,
		"New224": func_New224,
		"Sum224": func_Sum224,
		"Sum256": func_Sum256,
	})
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New())
}

func func_New224(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New224())
}

func func_Sum224(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum224(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_Sum256(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum256(ixgo.DirectCallArg[[]byte](ctx, 0)))
}
