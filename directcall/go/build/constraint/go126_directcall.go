// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package constraint

import (
	q "go/build/constraint"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("go/build/constraint", map[string]ixgo.DirectCallAdapter{
		"(*AndExpr).Eval":      method_ptr_AndExpr_Eval,
		"(*AndExpr).String":    method_ptr_AndExpr_String,
		"(*NotExpr).Eval":      method_ptr_NotExpr_Eval,
		"(*NotExpr).String":    method_ptr_NotExpr_String,
		"(*OrExpr).Eval":       method_ptr_OrExpr_Eval,
		"(*OrExpr).String":     method_ptr_OrExpr_String,
		"(*SyntaxError).Error": method_ptr_SyntaxError_Error,
		"(*TagExpr).Eval":      method_ptr_TagExpr_Eval,
		"(*TagExpr).String":    method_ptr_TagExpr_String,
		"GoVersion":            func_GoVersion,
		"IsGoBuild":            func_IsGoBuild,
		"IsPlusBuild":          func_IsPlusBuild,
	})
}

func method_ptr_AndExpr_Eval(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AndExpr).Eval(ixgo.DirectCallArg[*q.AndExpr](ctx, 0), ixgo.DirectCallArg[func(tag string) bool](ctx, 1)))
}

func method_ptr_AndExpr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.AndExpr).String(ixgo.DirectCallArg[*q.AndExpr](ctx, 0)))
}

func func_GoVersion(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.GoVersion(ixgo.DirectCallArg[q.Expr](ctx, 0)))
}

func func_IsGoBuild(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsGoBuild(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_IsPlusBuild(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsPlusBuild(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_NotExpr_Eval(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NotExpr).Eval(ixgo.DirectCallArg[*q.NotExpr](ctx, 0), ixgo.DirectCallArg[func(tag string) bool](ctx, 1)))
}

func method_ptr_NotExpr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NotExpr).String(ixgo.DirectCallArg[*q.NotExpr](ctx, 0)))
}

func method_ptr_OrExpr_Eval(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OrExpr).Eval(ixgo.DirectCallArg[*q.OrExpr](ctx, 0), ixgo.DirectCallArg[func(tag string) bool](ctx, 1)))
}

func method_ptr_OrExpr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OrExpr).String(ixgo.DirectCallArg[*q.OrExpr](ctx, 0)))
}

func method_ptr_SyntaxError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SyntaxError).Error(ixgo.DirectCallArg[*q.SyntaxError](ctx, 0)))
}

func method_ptr_TagExpr_Eval(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TagExpr).Eval(ixgo.DirectCallArg[*q.TagExpr](ctx, 0), ixgo.DirectCallArg[func(tag string) bool](ctx, 1)))
}

func method_ptr_TagExpr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TagExpr).String(ixgo.DirectCallArg[*q.TagExpr](ctx, 0)))
}
