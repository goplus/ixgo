// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package md5

import (
	q "crypto/md5"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/md5", map[string]ixgo.DirectCallAdapter{
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
