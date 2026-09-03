// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package json

import (
	q "encoding/json/v2"

	jsontext "encoding/json/jsontext"
	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/json/v2", map[string]ixgo.DirectCallAdapter{
		"(*SemanticError).Error":    method_ptr_SemanticError_Error,
		"(*SemanticError).Unwrap":   method_ptr_SemanticError_Unwrap,
		"DefaultOptionsV2":          func_DefaultOptionsV2,
		"Deterministic":             func_Deterministic,
		"FormatNilMapAsNull":        func_FormatNilMapAsNull,
		"FormatNilSliceAsNull":      func_FormatNilSliceAsNull,
		"JoinOptions":               func_JoinOptions,
		"MarshalEncode":             func_MarshalEncode,
		"MarshalWrite":              func_MarshalWrite,
		"MatchCaseInsensitiveNames": func_MatchCaseInsensitiveNames,
		"OmitZeroStructFields":      func_OmitZeroStructFields,
		"RejectUnknownMembers":      func_RejectUnknownMembers,
		"StringifyNumbers":          func_StringifyNumbers,
		"Unmarshal":                 func_Unmarshal,
		"UnmarshalDecode":           func_UnmarshalDecode,
		"UnmarshalRead":             func_UnmarshalRead,
	})
}

func func_DefaultOptionsV2(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DefaultOptionsV2())
}

func func_Deterministic(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Deterministic(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_FormatNilMapAsNull(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatNilMapAsNull(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_FormatNilSliceAsNull(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatNilSliceAsNull(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_JoinOptions(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.JoinOptions(ixgo.DirectCallArg[[]q.Options](ctx, 0)...))
}

func func_MarshalEncode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MarshalEncode(ixgo.DirectCallArg[*jsontext.Encoder](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1), ixgo.DirectCallArg[[]q.Options](ctx, 2)...))
}

func func_MarshalWrite(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MarshalWrite(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1), ixgo.DirectCallArg[[]q.Options](ctx, 2)...))
}

func func_MatchCaseInsensitiveNames(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MatchCaseInsensitiveNames(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_OmitZeroStructFields(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OmitZeroStructFields(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_RejectUnknownMembers(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RejectUnknownMembers(ixgo.DirectCallArg[bool](ctx, 0)))
}

func method_ptr_SemanticError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SemanticError).Error(ixgo.DirectCallArg[*q.SemanticError](ctx, 0)))
}

func method_ptr_SemanticError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SemanticError).Unwrap(ixgo.DirectCallArg[*q.SemanticError](ctx, 0)))
}

func func_StringifyNumbers(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StringifyNumbers(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_Unmarshal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Unmarshal(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1), ixgo.DirectCallArg[[]q.Options](ctx, 2)...))
}

func func_UnmarshalDecode(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnmarshalDecode(ixgo.DirectCallArg[*jsontext.Decoder](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1), ixgo.DirectCallArg[[]q.Options](ctx, 2)...))
}

func func_UnmarshalRead(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnmarshalRead(ixgo.DirectCallArg[io.Reader](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1), ixgo.DirectCallArg[[]q.Options](ctx, 2)...))
}
