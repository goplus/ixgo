// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package format

import (
	q "go/format"

	"github.com/goplus/ixgo"
	token "go/token"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("go/format", map[string]ixgo.DirectCallAdapter{
		"Node": func_Node,
	})
}

func func_Node(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Node(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[*token.FileSet](ctx, 1), ixgo.DirectCallArg[any](ctx, 2)))
}
