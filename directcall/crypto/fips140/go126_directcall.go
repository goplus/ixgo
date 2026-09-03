// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package fips140

import (
	q "crypto/fips140"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/fips140", map[string]ixgo.DirectCallAdapter{
		"Enabled":            func_Enabled,
		"Enforced":           func_Enforced,
		"Version":            func_Version,
		"WithoutEnforcement": func_WithoutEnforcement,
	})
}

func func_Enabled(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Enabled())
}

func func_Enforced(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Enforced())
}

func func_Version(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Version())
}

func func_WithoutEnforcement(ctx ixgo.DirectCallContext) {
	q.WithoutEnforcement(ixgo.DirectCallArg[func()](ctx, 0))
}
