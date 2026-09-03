// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package url

import (
	q "net/url"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("net/url", map[string]ixgo.DirectCallAdapter{
		"(*Error).Error":            method_ptr_Error_Error,
		"(*Error).Temporary":        method_ptr_Error_Temporary,
		"(*Error).Timeout":          method_ptr_Error_Timeout,
		"(*Error).Unwrap":           method_ptr_Error_Unwrap,
		"(*EscapeError).Error":      method_ptr_EscapeError_Error,
		"(*InvalidHostError).Error": method_ptr_InvalidHostError_Error,
		"(*URL).Clone":              method_ptr_URL_Clone,
		"(*URL).EscapedFragment":    method_ptr_URL_EscapedFragment,
		"(*URL).EscapedPath":        method_ptr_URL_EscapedPath,
		"(*URL).Hostname":           method_ptr_URL_Hostname,
		"(*URL).IsAbs":              method_ptr_URL_IsAbs,
		"(*URL).JoinPath":           method_ptr_URL_JoinPath,
		"(*URL).Port":               method_ptr_URL_Port,
		"(*URL).Query":              method_ptr_URL_Query,
		"(*URL).Redacted":           method_ptr_URL_Redacted,
		"(*URL).RequestURI":         method_ptr_URL_RequestURI,
		"(*URL).ResolveReference":   method_ptr_URL_ResolveReference,
		"(*URL).String":             method_ptr_URL_String,
		"(*URL).UnmarshalBinary":    method_ptr_URL_UnmarshalBinary,
		"(*Userinfo).String":        method_ptr_Userinfo_String,
		"(*Userinfo).Username":      method_ptr_Userinfo_Username,
		"(*Values).Add":             method_ptr_Values_Add,
		"(*Values).Clone":           method_ptr_Values_Clone,
		"(*Values).Del":             method_ptr_Values_Del,
		"(*Values).Encode":          method_ptr_Values_Encode,
		"(*Values).Get":             method_ptr_Values_Get,
		"(*Values).Has":             method_ptr_Values_Has,
		"(*Values).Set":             method_ptr_Values_Set,
		"(EscapeError).Error":       method_EscapeError_Error,
		"(InvalidHostError).Error":  method_InvalidHostError_Error,
		"(Values).Add":              method_Values_Add,
		"(Values).Clone":            method_Values_Clone,
		"(Values).Del":              method_Values_Del,
		"(Values).Encode":           method_Values_Encode,
		"(Values).Get":              method_Values_Get,
		"(Values).Has":              method_Values_Has,
		"(Values).Set":              method_Values_Set,
		"PathEscape":                func_PathEscape,
		"QueryEscape":               func_QueryEscape,
		"User":                      func_User,
		"UserPassword":              func_UserPassword,
	})
}

func method_ptr_Error_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Error(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func method_ptr_Error_Temporary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Temporary(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func method_ptr_Error_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Timeout(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func method_ptr_Error_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Unwrap(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func method_EscapeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EscapeError.Error(ixgo.DirectCallArg[q.EscapeError](ctx, 0)))
}

func method_ptr_EscapeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.EscapeError).Error(ixgo.DirectCallArg[*q.EscapeError](ctx, 0)))
}

func method_InvalidHostError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.InvalidHostError.Error(ixgo.DirectCallArg[q.InvalidHostError](ctx, 0)))
}

func method_ptr_InvalidHostError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InvalidHostError).Error(ixgo.DirectCallArg[*q.InvalidHostError](ctx, 0)))
}

func func_PathEscape(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PathEscape(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_QueryEscape(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.QueryEscape(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_URL_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).Clone(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_EscapedFragment(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).EscapedFragment(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_EscapedPath(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).EscapedPath(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_Hostname(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).Hostname(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_IsAbs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).IsAbs(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_JoinPath(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).JoinPath(ixgo.DirectCallArg[*q.URL](ctx, 0), ixgo.DirectCallArg[[]string](ctx, 1)...))
}

func method_ptr_URL_Port(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).Port(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_Query(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).Query(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_Redacted(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).Redacted(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_RequestURI(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).RequestURI(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_ResolveReference(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).ResolveReference(ixgo.DirectCallArg[*q.URL](ctx, 0), ixgo.DirectCallArg[*q.URL](ctx, 1)))
}

func method_ptr_URL_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).String(ixgo.DirectCallArg[*q.URL](ctx, 0)))
}

func method_ptr_URL_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.URL).UnmarshalBinary(ixgo.DirectCallArg[*q.URL](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_User(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.User(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_UserPassword(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UserPassword(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Userinfo_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Userinfo).String(ixgo.DirectCallArg[*q.Userinfo](ctx, 0)))
}

func method_ptr_Userinfo_Username(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Userinfo).Username(ixgo.DirectCallArg[*q.Userinfo](ctx, 0)))
}

func method_Values_Add(ctx ixgo.DirectCallContext) {
	q.Values.Add(ixgo.DirectCallArg[q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_Values_Add(ctx ixgo.DirectCallContext) {
	(*q.Values).Add(ixgo.DirectCallArg[*q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_Values_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Values.Clone(ixgo.DirectCallArg[q.Values](ctx, 0)))
}

func method_ptr_Values_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Values).Clone(ixgo.DirectCallArg[*q.Values](ctx, 0)))
}

func method_Values_Del(ctx ixgo.DirectCallContext) {
	q.Values.Del(ixgo.DirectCallArg[q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_Values_Del(ctx ixgo.DirectCallContext) {
	(*q.Values).Del(ixgo.DirectCallArg[*q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_Values_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Values.Encode(ixgo.DirectCallArg[q.Values](ctx, 0)))
}

func method_ptr_Values_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Values).Encode(ixgo.DirectCallArg[*q.Values](ctx, 0)))
}

func method_Values_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Values.Get(ixgo.DirectCallArg[q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Values_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Values).Get(ixgo.DirectCallArg[*q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_Values_Has(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Values.Has(ixgo.DirectCallArg[q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Values_Has(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Values).Has(ixgo.DirectCallArg[*q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_Values_Set(ctx ixgo.DirectCallContext) {
	q.Values.Set(ixgo.DirectCallArg[q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func method_ptr_Values_Set(ctx ixgo.DirectCallContext) {
	(*q.Values).Set(ixgo.DirectCallArg[*q.Values](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}
