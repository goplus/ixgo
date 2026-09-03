// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package rand

import (
	q "crypto/rand"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("crypto/rand", map[string]ixgo.DirectCallAdapter{
		"Text": func_Text,
	})
}

func func_Text(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Text())
}
