// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package synctest

import (
	q "testing/synctest"

	"github.com/goplus/ixgo"
	testing "testing"
)

func init() {
	ixgo.RegisterDirectCalls("testing/synctest", map[string]ixgo.DirectCallAdapter{
		"Test": func_Test,
		"Wait": func_Wait,
	})
}

func func_Test(ctx ixgo.DirectCallContext) {
	q.Test(ixgo.DirectCallArg[*testing.T](ctx, 0), ixgo.DirectCallArg[func(*testing.T)](ctx, 1))
}

func func_Wait(ctx ixgo.DirectCallContext) {
	q.Wait()
}
