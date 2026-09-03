// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package fips140

import (
	q "crypto/fips140"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/fips140", map[string]ixgo.DirectCallAdapter{
		"Enabled": func_Enabled,
	})
}

func func_Enabled(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Enabled())
}
