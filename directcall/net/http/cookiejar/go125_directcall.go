// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package cookiejar

import (
	q "net/http/cookiejar"

	"github.com/goplus/ixgo"
	http "net/http"
	url "net/url"
)

func init() {
	ixgo.RegisterDirectCalls("net/http/cookiejar", map[string]ixgo.DirectCallAdapter{
		"(*Jar).Cookies":    method_ptr_Jar_Cookies,
		"(*Jar).SetCookies": method_ptr_Jar_SetCookies,
	})
}

func method_ptr_Jar_Cookies(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Jar).Cookies(ixgo.DirectCallArg[*q.Jar](ctx, 0), ixgo.DirectCallArg[*url.URL](ctx, 1)))
}

func method_ptr_Jar_SetCookies(ctx ixgo.DirectCallContext) {
	(*q.Jar).SetCookies(ixgo.DirectCallArg[*q.Jar](ctx, 0), ixgo.DirectCallArg[*url.URL](ctx, 1), ixgo.DirectCallArg[[]*http.Cookie](ctx, 2))
}
