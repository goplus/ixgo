// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package jsontext

import (
	q "encoding/json/jsontext"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/json/jsontext", map[string]ixgo.DirectCallAdapter{
		"(*Decoder).InputOffset":     method_ptr_Decoder_InputOffset,
		"(*Decoder).Options":         method_ptr_Decoder_Options,
		"(*Decoder).PeekKind":        method_ptr_Decoder_PeekKind,
		"(*Decoder).Reset":           method_ptr_Decoder_Reset,
		"(*Decoder).SkipValue":       method_ptr_Decoder_SkipValue,
		"(*Decoder).StackDepth":      method_ptr_Decoder_StackDepth,
		"(*Decoder).StackPointer":    method_ptr_Decoder_StackPointer,
		"(*Decoder).UnreadBuffer":    method_ptr_Decoder_UnreadBuffer,
		"(*Encoder).AvailableBuffer": method_ptr_Encoder_AvailableBuffer,
		"(*Encoder).Options":         method_ptr_Encoder_Options,
		"(*Encoder).OutputOffset":    method_ptr_Encoder_OutputOffset,
		"(*Encoder).Reset":           method_ptr_Encoder_Reset,
		"(*Encoder).StackDepth":      method_ptr_Encoder_StackDepth,
		"(*Encoder).StackPointer":    method_ptr_Encoder_StackPointer,
		"(*Encoder).WriteToken":      method_ptr_Encoder_WriteToken,
		"(*Encoder).WriteValue":      method_ptr_Encoder_WriteValue,
		"(*Kind).String":             method_ptr_Kind_String,
		"(*Pointer).AppendToken":     method_ptr_Pointer_AppendToken,
		"(*Pointer).Contains":        method_ptr_Pointer_Contains,
		"(*Pointer).IsValid":         method_ptr_Pointer_IsValid,
		"(*Pointer).LastToken":       method_ptr_Pointer_LastToken,
		"(*Pointer).Parent":          method_ptr_Pointer_Parent,
		"(*Pointer).Tokens":          method_ptr_Pointer_Tokens,
		"(*SyntacticError).Error":    method_ptr_SyntacticError_Error,
		"(*SyntacticError).Unwrap":   method_ptr_SyntacticError_Unwrap,
		"(*Token).Bool":              method_ptr_Token_Bool,
		"(*Token).Clone":             method_ptr_Token_Clone,
		"(*Token).Kind":              method_ptr_Token_Kind,
		"(*Token).String":            method_ptr_Token_String,
		"(*Value).Canonicalize":      method_ptr_Value_Canonicalize,
		"(*Value).Clone":             method_ptr_Value_Clone,
		"(*Value).Compact":           method_ptr_Value_Compact,
		"(*Value).Format":            method_ptr_Value_Format,
		"(*Value).Indent":            method_ptr_Value_Indent,
		"(*Value).IsValid":           method_ptr_Value_IsValid,
		"(*Value).Kind":              method_ptr_Value_Kind,
		"(*Value).String":            method_ptr_Value_String,
		"(*Value).UnmarshalJSON":     method_ptr_Value_UnmarshalJSON,
		"(Kind).String":              method_Kind_String,
		"(Pointer).AppendToken":      method_Pointer_AppendToken,
		"(Pointer).Contains":         method_Pointer_Contains,
		"(Pointer).IsValid":          method_Pointer_IsValid,
		"(Pointer).LastToken":        method_Pointer_LastToken,
		"(Pointer).Parent":           method_Pointer_Parent,
		"(Pointer).Tokens":           method_Pointer_Tokens,
		"(Token).Bool":               method_Token_Bool,
		"(Token).Clone":              method_Token_Clone,
		"(Token).Kind":               method_Token_Kind,
		"(Token).String":             method_Token_String,
		"(Value).Clone":              method_Value_Clone,
		"(Value).IsValid":            method_Value_IsValid,
		"(Value).Kind":               method_Value_Kind,
		"(Value).String":             method_Value_String,
		"AllowDuplicateNames":        func_AllowDuplicateNames,
		"AllowInvalidUTF8":           func_AllowInvalidUTF8,
		"AppendFloat":                func_AppendFloat,
		"Bool":                       func_Bool,
		"CanonicalizeRawFloats":      func_CanonicalizeRawFloats,
		"CanonicalizeRawInts":        func_CanonicalizeRawInts,
		"EscapeForHTML":              func_EscapeForHTML,
		"EscapeForJS":                func_EscapeForJS,
		"Float":                      func_Float,
		"Float32":                    func_Float32,
		"Int":                        func_Int,
		"Multiline":                  func_Multiline,
		"NewDecoder":                 func_NewDecoder,
		"NewEncoder":                 func_NewEncoder,
		"PreserveRawStrings":         func_PreserveRawStrings,
		"ReorderRawObjects":          func_ReorderRawObjects,
		"SpaceAfterColon":            func_SpaceAfterColon,
		"SpaceAfterComma":            func_SpaceAfterComma,
		"String":                     func_String,
		"Uint":                       func_Uint,
		"WithIndent":                 func_WithIndent,
		"WithIndentPrefix":           func_WithIndentPrefix,
	})
}

func func_AllowDuplicateNames(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AllowDuplicateNames(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_AllowInvalidUTF8(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AllowInvalidUTF8(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_AppendFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendFloat(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_Bool(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Bool(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_CanonicalizeRawFloats(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CanonicalizeRawFloats(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_CanonicalizeRawInts(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CanonicalizeRawInts(ixgo.DirectCallArg[bool](ctx, 0)))
}

func method_ptr_Decoder_InputOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).InputOffset(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_Options(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).Options(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_PeekKind(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).PeekKind(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_Reset(ctx ixgo.DirectCallContext) {
	(*q.Decoder).Reset(ixgo.DirectCallArg[*q.Decoder](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1), ixgo.DirectCallArg[[]q.Options](ctx, 2)...)
}

func method_ptr_Decoder_SkipValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).SkipValue(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_StackDepth(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).StackDepth(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_StackPointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).StackPointer(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_UnreadBuffer(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).UnreadBuffer(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Encoder_AvailableBuffer(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).AvailableBuffer(ixgo.DirectCallArg[*q.Encoder](ctx, 0)))
}

func method_ptr_Encoder_Options(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).Options(ixgo.DirectCallArg[*q.Encoder](ctx, 0)))
}

func method_ptr_Encoder_OutputOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).OutputOffset(ixgo.DirectCallArg[*q.Encoder](ctx, 0)))
}

func method_ptr_Encoder_Reset(ctx ixgo.DirectCallContext) {
	(*q.Encoder).Reset(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[[]q.Options](ctx, 2)...)
}

func method_ptr_Encoder_StackDepth(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).StackDepth(ixgo.DirectCallArg[*q.Encoder](ctx, 0)))
}

func method_ptr_Encoder_StackPointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).StackPointer(ixgo.DirectCallArg[*q.Encoder](ctx, 0)))
}

func method_ptr_Encoder_WriteToken(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).WriteToken(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[q.Token](ctx, 1)))
}

func method_ptr_Encoder_WriteValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).WriteValue(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1)))
}

func func_EscapeForHTML(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EscapeForHTML(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_EscapeForJS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EscapeForJS(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_Float(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_Float32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float32(ixgo.DirectCallArg[float32](ctx, 0)))
}

func func_Int(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int(ixgo.DirectCallArg[int64](ctx, 0)))
}

func method_Kind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Kind.String(ixgo.DirectCallArg[q.Kind](ctx, 0)))
}

func method_ptr_Kind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Kind).String(ixgo.DirectCallArg[*q.Kind](ctx, 0)))
}

func func_Multiline(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Multiline(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_NewDecoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewDecoder(ixgo.DirectCallArg[io.Reader](ctx, 0), ixgo.DirectCallArg[[]q.Options](ctx, 1)...))
}

func func_NewEncoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewEncoder(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[[]q.Options](ctx, 1)...))
}

func method_Pointer_AppendToken(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pointer.AppendToken(ixgo.DirectCallArg[q.Pointer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Pointer_AppendToken(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pointer).AppendToken(ixgo.DirectCallArg[*q.Pointer](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_Pointer_Contains(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pointer.Contains(ixgo.DirectCallArg[q.Pointer](ctx, 0), ixgo.DirectCallArg[q.Pointer](ctx, 1)))
}

func method_ptr_Pointer_Contains(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pointer).Contains(ixgo.DirectCallArg[*q.Pointer](ctx, 0), ixgo.DirectCallArg[q.Pointer](ctx, 1)))
}

func method_Pointer_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pointer.IsValid(ixgo.DirectCallArg[q.Pointer](ctx, 0)))
}

func method_ptr_Pointer_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pointer).IsValid(ixgo.DirectCallArg[*q.Pointer](ctx, 0)))
}

func method_Pointer_LastToken(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pointer.LastToken(ixgo.DirectCallArg[q.Pointer](ctx, 0)))
}

func method_ptr_Pointer_LastToken(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pointer).LastToken(ixgo.DirectCallArg[*q.Pointer](ctx, 0)))
}

func method_Pointer_Parent(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pointer.Parent(ixgo.DirectCallArg[q.Pointer](ctx, 0)))
}

func method_ptr_Pointer_Parent(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pointer).Parent(ixgo.DirectCallArg[*q.Pointer](ctx, 0)))
}

func method_Pointer_Tokens(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pointer.Tokens(ixgo.DirectCallArg[q.Pointer](ctx, 0)))
}

func method_ptr_Pointer_Tokens(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pointer).Tokens(ixgo.DirectCallArg[*q.Pointer](ctx, 0)))
}

func func_PreserveRawStrings(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PreserveRawStrings(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_ReorderRawObjects(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ReorderRawObjects(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_SpaceAfterColon(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SpaceAfterColon(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_SpaceAfterComma(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SpaceAfterComma(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.String(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_SyntacticError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SyntacticError).Error(ixgo.DirectCallArg[*q.SyntacticError](ctx, 0)))
}

func method_ptr_SyntacticError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SyntacticError).Unwrap(ixgo.DirectCallArg[*q.SyntacticError](ctx, 0)))
}

func method_Token_Bool(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Token.Bool(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Token_Bool(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Token).Bool(ixgo.DirectCallArg[*q.Token](ctx, 0)))
}

func method_Token_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Token.Clone(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Token_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Token).Clone(ixgo.DirectCallArg[*q.Token](ctx, 0)))
}

func method_Token_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Token.Kind(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Token_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Token).Kind(ixgo.DirectCallArg[*q.Token](ctx, 0)))
}

func method_Token_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Token.String(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Token_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Token).String(ixgo.DirectCallArg[*q.Token](ctx, 0)))
}

func func_Uint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func method_ptr_Value_Canonicalize(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Canonicalize(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Options](ctx, 1)...))
}

func method_Value_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Clone(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Clone(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Clone(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_ptr_Value_Compact(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Compact(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Options](ctx, 1)...))
}

func method_ptr_Value_Format(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Format(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Options](ctx, 1)...))
}

func method_ptr_Value_Indent(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Indent(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Options](ctx, 1)...))
}

func method_Value_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.IsValid(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Options](ctx, 1)...))
}

func method_ptr_Value_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).IsValid(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Options](ctx, 1)...))
}

func method_Value_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Kind(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Kind(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.String(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).String(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_ptr_Value_UnmarshalJSON(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).UnmarshalJSON(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_WithIndent(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WithIndent(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_WithIndentPrefix(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WithIndentPrefix(ixgo.DirectCallArg[string](ctx, 0)))
}
