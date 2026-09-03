// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package crc32

import (
	q "hash/crc32"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("hash/crc32", map[string]ixgo.DirectCallAdapter{
		"Checksum":     func_Checksum,
		"ChecksumIEEE": func_ChecksumIEEE,
		"MakeTable":    func_MakeTable,
		"New":          func_New,
		"NewIEEE":      func_NewIEEE,
		"Update":       func_Update,
	})
}

func func_Checksum(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Checksum(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[*q.Table](ctx, 1)))
}

func func_ChecksumIEEE(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ChecksumIEEE(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_MakeTable(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeTable(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[*q.Table](ctx, 0)))
}

func func_NewIEEE(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewIEEE())
}

func func_Update(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Update(ixgo.DirectCallArg[uint32](ctx, 0), ixgo.DirectCallArg[*q.Table](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}
