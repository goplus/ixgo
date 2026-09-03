// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package crc64

import (
	q "hash/crc64"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("hash/crc64", map[string]ixgo.DirectCallAdapter{
		"Checksum":  func_Checksum,
		"MakeTable": func_MakeTable,
		"New":       func_New,
		"Update":    func_Update,
	})
}

func func_Checksum(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Checksum(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[*q.Table](ctx, 1)))
}

func func_MakeTable(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeTable(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[*q.Table](ctx, 0)))
}

func func_Update(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Update(ixgo.DirectCallArg[uint64](ctx, 0), ixgo.DirectCallArg[*q.Table](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}
