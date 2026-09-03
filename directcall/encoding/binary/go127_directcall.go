// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package binary

import (
	q "encoding/binary"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/binary", map[string]ixgo.DirectCallAdapter{
		"AppendUvarint": func_AppendUvarint,
		"AppendVarint":  func_AppendVarint,
		"PutUvarint":    func_PutUvarint,
		"PutVarint":     func_PutVarint,
		"Read":          func_Read,
		"Size":          func_Size,
		"Write":         func_Write,
	})
}

func func_AppendUvarint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendUvarint(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func func_AppendVarint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendVarint(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_PutUvarint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PutUvarint(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func func_PutVarint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PutVarint(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_Read(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Read(ixgo.DirectCallArg[io.Reader](ctx, 0), ixgo.DirectCallArg[q.ByteOrder](ctx, 1), ixgo.DirectCallArg[interface{}](ctx, 2)))
}

func func_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Size(ixgo.DirectCallArg[interface{}](ctx, 0)))
}

func func_Write(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Write(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[q.ByteOrder](ctx, 1), ixgo.DirectCallArg[interface{}](ctx, 2)))
}
