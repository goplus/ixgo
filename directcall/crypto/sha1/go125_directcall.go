// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package sha1

import (
	q "crypto/sha1"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/sha1", map[string]ixgo.DirectCallAdapter{
		"New": func_New,
		"Sum": func_Sum,
	})
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New())
}

func func_Sum(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Sum(ixgo.DirectCallArg[[]byte](ctx, 0)))
}
