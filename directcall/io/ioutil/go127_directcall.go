// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package ioutil

import (
	q "io/ioutil"

	"github.com/goplus/ixgo"
	io "io"
	fs "io/fs"
)

func init() {
	ixgo.RegisterDirectCalls("io/ioutil", map[string]ixgo.DirectCallAdapter{
		"NopCloser": func_NopCloser,
		"WriteFile": func_WriteFile,
	})
}

func func_NopCloser(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NopCloser(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_WriteFile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WriteFile(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[fs.FileMode](ctx, 2)))
}
