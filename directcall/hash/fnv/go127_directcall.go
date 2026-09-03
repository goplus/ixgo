// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package fnv

import (
	q "hash/fnv"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("hash/fnv", map[string]ixgo.DirectCallAdapter{
		"New128":  func_New128,
		"New128a": func_New128a,
		"New32":   func_New32,
		"New32a":  func_New32a,
		"New64":   func_New64,
		"New64a":  func_New64a,
	})
}

func func_New128(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New128())
}

func func_New128a(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New128a())
}

func func_New32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New32())
}

func func_New32a(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New32a())
}

func func_New64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New64())
}

func func_New64a(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New64a())
}
