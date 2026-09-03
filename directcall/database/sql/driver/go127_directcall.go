// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package driver

import (
	q "database/sql/driver"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("database/sql/driver", map[string]ixgo.DirectCallAdapter{
		"IsScanValue": func_IsScanValue,
		"IsValue":     func_IsValue,
	})
}

func func_IsScanValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsScanValue(ixgo.DirectCallArg[interface{}](ctx, 0)))
}

func func_IsValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsValue(ixgo.DirectCallArg[interface{}](ctx, 0)))
}
