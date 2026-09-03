// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package adler32

import (
	q "hash/adler32"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("hash/adler32", map[string]ixgo.DirectCallAdapter{
		"Checksum": func_Checksum,
		"New":      func_New,
	})
}

func func_Checksum(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Checksum(ixgo.DirectCallArg[[]byte](ctx, 0)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New())
}
