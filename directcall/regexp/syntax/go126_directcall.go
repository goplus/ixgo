// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package syntax

import (
	q "regexp/syntax"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("regexp/syntax", map[string]ixgo.DirectCallAdapter{
		"(*Error).Error":          method_ptr_Error_Error,
		"(*ErrorCode).String":     method_ptr_ErrorCode_String,
		"(*Inst).MatchEmptyWidth": method_ptr_Inst_MatchEmptyWidth,
		"(*Inst).MatchRune":       method_ptr_Inst_MatchRune,
		"(*Inst).MatchRunePos":    method_ptr_Inst_MatchRunePos,
		"(*Inst).String":          method_ptr_Inst_String,
		"(*InstOp).String":        method_ptr_InstOp_String,
		"(*Op).String":            method_ptr_Op_String,
		"(*Prog).StartCond":       method_ptr_Prog_StartCond,
		"(*Prog).String":          method_ptr_Prog_String,
		"(*Regexp).CapNames":      method_ptr_Regexp_CapNames,
		"(*Regexp).Equal":         method_ptr_Regexp_Equal,
		"(*Regexp).MaxCap":        method_ptr_Regexp_MaxCap,
		"(*Regexp).Simplify":      method_ptr_Regexp_Simplify,
		"(*Regexp).String":        method_ptr_Regexp_String,
		"(ErrorCode).String":      method_ErrorCode_String,
		"(InstOp).String":         method_InstOp_String,
		"(Op).String":             method_Op_String,
		"EmptyOpContext":          func_EmptyOpContext,
		"IsWordChar":              func_IsWordChar,
	})
}

func func_EmptyOpContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EmptyOpContext(ixgo.DirectCallArg[rune](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func method_ptr_Error_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Error(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func method_ErrorCode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ErrorCode.String(ixgo.DirectCallArg[q.ErrorCode](ctx, 0)))
}

func method_ptr_ErrorCode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ErrorCode).String(ixgo.DirectCallArg[*q.ErrorCode](ctx, 0)))
}

func method_ptr_Inst_MatchEmptyWidth(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Inst).MatchEmptyWidth(ixgo.DirectCallArg[*q.Inst](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1), ixgo.DirectCallArg[rune](ctx, 2)))
}

func method_ptr_Inst_MatchRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Inst).MatchRune(ixgo.DirectCallArg[*q.Inst](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func method_ptr_Inst_MatchRunePos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Inst).MatchRunePos(ixgo.DirectCallArg[*q.Inst](ctx, 0), ixgo.DirectCallArg[rune](ctx, 1)))
}

func method_ptr_Inst_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Inst).String(ixgo.DirectCallArg[*q.Inst](ctx, 0)))
}

func method_InstOp_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.InstOp.String(ixgo.DirectCallArg[q.InstOp](ctx, 0)))
}

func method_ptr_InstOp_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InstOp).String(ixgo.DirectCallArg[*q.InstOp](ctx, 0)))
}

func func_IsWordChar(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsWordChar(ixgo.DirectCallArg[rune](ctx, 0)))
}

func method_Op_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Op.String(ixgo.DirectCallArg[q.Op](ctx, 0)))
}

func method_ptr_Op_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Op).String(ixgo.DirectCallArg[*q.Op](ctx, 0)))
}

func method_ptr_Prog_StartCond(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prog).StartCond(ixgo.DirectCallArg[*q.Prog](ctx, 0)))
}

func method_ptr_Prog_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prog).String(ixgo.DirectCallArg[*q.Prog](ctx, 0)))
}

func method_ptr_Regexp_CapNames(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).CapNames(ixgo.DirectCallArg[*q.Regexp](ctx, 0)))
}

func method_ptr_Regexp_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).Equal(ixgo.DirectCallArg[*q.Regexp](ctx, 0), ixgo.DirectCallArg[*q.Regexp](ctx, 1)))
}

func method_ptr_Regexp_MaxCap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).MaxCap(ixgo.DirectCallArg[*q.Regexp](ctx, 0)))
}

func method_ptr_Regexp_Simplify(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).Simplify(ixgo.DirectCallArg[*q.Regexp](ctx, 0)))
}

func method_ptr_Regexp_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Regexp).String(ixgo.DirectCallArg[*q.Regexp](ctx, 0)))
}
