// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package quotedprintable

import (
	q "mime/quotedprintable"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("mime/quotedprintable", map[string]ixgo.DirectCallAdapter{
		"(*Writer).Close": method_ptr_Writer_Close,
		"NewReader":       func_NewReader,
		"NewWriter":       func_NewWriter,
	})
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func method_ptr_Writer_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Close(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}
