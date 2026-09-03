// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package fstest

import (
	q "testing/fstest"

	"github.com/goplus/ixgo"
	fs "io/fs"
)

func init() {
	ixgo.RegisterDirectCalls("testing/fstest", map[string]ixgo.DirectCallAdapter{
		"TestFS": func_TestFS,
	})
}

func func_TestFS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TestFS(ixgo.DirectCallArg[fs.FS](ctx, 0), ixgo.DirectCallArg[[]string](ctx, 1)...))
}
