// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package json

import (
	q "encoding/json"

	bytes "bytes"
	jsontext "encoding/json/jsontext"
	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/json", map[string]ixgo.DirectCallAdapter{
		"(*Decoder).Buffered":              method_ptr_Decoder_Buffered,
		"(*Decoder).Decode":                method_ptr_Decoder_Decode,
		"(*Decoder).DisallowUnknownFields": method_ptr_Decoder_DisallowUnknownFields,
		"(*Decoder).InputOffset":           method_ptr_Decoder_InputOffset,
		"(*Decoder).More":                  method_ptr_Decoder_More,
		"(*Decoder).UseNumber":             method_ptr_Decoder_UseNumber,
		"(*Delim).String":                  method_ptr_Delim_String,
		"(*Encoder).Encode":                method_ptr_Encoder_Encode,
		"(*Encoder).SetEscapeHTML":         method_ptr_Encoder_SetEscapeHTML,
		"(*Encoder).SetIndent":             method_ptr_Encoder_SetIndent,
		"(*InvalidUTF8Error).Error":        method_ptr_InvalidUTF8Error_Error,
		"(*InvalidUnmarshalError).Error":   method_ptr_InvalidUnmarshalError_Error,
		"(*MarshalerError).Error":          method_ptr_MarshalerError_Error,
		"(*MarshalerError).Unwrap":         method_ptr_MarshalerError_Unwrap,
		"(*Number).MarshalJSONTo":          method_ptr_Number_MarshalJSONTo,
		"(*Number).String":                 method_ptr_Number_String,
		"(*Number).UnmarshalJSONFrom":      method_ptr_Number_UnmarshalJSONFrom,
		"(*SyntaxError).Error":             method_ptr_SyntaxError_Error,
		"(*UnmarshalFieldError).Error":     method_ptr_UnmarshalFieldError_Error,
		"(*UnmarshalTypeError).Error":      method_ptr_UnmarshalTypeError_Error,
		"(*UnmarshalTypeError).Unwrap":     method_ptr_UnmarshalTypeError_Unwrap,
		"(*UnsupportedTypeError).Error":    method_ptr_UnsupportedTypeError_Error,
		"(*UnsupportedValueError).Error":   method_ptr_UnsupportedValueError_Error,
		"(Delim).String":                   method_Delim_String,
		"(Number).MarshalJSONTo":           method_Number_MarshalJSONTo,
		"(Number).String":                  method_Number_String,
		"CallMethodsWithLegacySemantics":   func_CallMethodsWithLegacySemantics,
		"Compact":                          func_Compact,
		"DefaultOptionsV1":                 func_DefaultOptionsV1,
		"FormatByteArrayAsArray":           func_FormatByteArrayAsArray,
		"FormatBytesWithLegacySemantics":   func_FormatBytesWithLegacySemantics,
		"FormatDurationAsNano":             func_FormatDurationAsNano,
		"HTMLEscape":                       func_HTMLEscape,
		"Indent":                           func_Indent,
		"MatchCaseSensitiveDelimiter":      func_MatchCaseSensitiveDelimiter,
		"MergeWithLegacySemantics":         func_MergeWithLegacySemantics,
		"NewDecoder":                       func_NewDecoder,
		"NewEncoder":                       func_NewEncoder,
		"OmitEmptyWithLegacySemantics":     func_OmitEmptyWithLegacySemantics,
		"ParseBytesWithLooseRFC4648":       func_ParseBytesWithLooseRFC4648,
		"ParseTimeWithLooseRFC3339":        func_ParseTimeWithLooseRFC3339,
		"ReportErrorsWithLegacySemantics":  func_ReportErrorsWithLegacySemantics,
		"StringifyWithLegacySemantics":     func_StringifyWithLegacySemantics,
		"Unmarshal":                        func_Unmarshal,
		"UnmarshalArrayFromAnyLength":      func_UnmarshalArrayFromAnyLength,
		"Valid":                            func_Valid,
	})
}

func func_CallMethodsWithLegacySemantics(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CallMethodsWithLegacySemantics(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_Compact(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Compact(ixgo.DirectCallArg[*bytes.Buffer](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Decoder_Buffered(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).Buffered(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_Decode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).Decode(ixgo.DirectCallArg[*q.Decoder](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1)))
}

func method_ptr_Decoder_DisallowUnknownFields(ctx ixgo.DirectCallContext) {
	(*q.Decoder).DisallowUnknownFields(ixgo.DirectCallArg[*q.Decoder](ctx, 0))
}

func method_ptr_Decoder_InputOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).InputOffset(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_More(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).More(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_UseNumber(ctx ixgo.DirectCallContext) {
	(*q.Decoder).UseNumber(ixgo.DirectCallArg[*q.Decoder](ctx, 0))
}

func func_DefaultOptionsV1(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DefaultOptionsV1())
}

func method_Delim_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Delim.String(ixgo.DirectCallArg[q.Delim](ctx, 0)))
}

func method_ptr_Delim_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Delim).String(ixgo.DirectCallArg[*q.Delim](ctx, 0)))
}

func method_ptr_Encoder_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).Encode(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1)))
}

func method_ptr_Encoder_SetEscapeHTML(ctx ixgo.DirectCallContext) {
	(*q.Encoder).SetEscapeHTML(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_ptr_Encoder_SetIndent(ctx ixgo.DirectCallContext) {
	(*q.Encoder).SetIndent(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func func_FormatByteArrayAsArray(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatByteArrayAsArray(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_FormatBytesWithLegacySemantics(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatBytesWithLegacySemantics(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_FormatDurationAsNano(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FormatDurationAsNano(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_HTMLEscape(ctx ixgo.DirectCallContext) {
	q.HTMLEscape(ixgo.DirectCallArg[*bytes.Buffer](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1))
}

func func_Indent(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Indent(ixgo.DirectCallArg[*bytes.Buffer](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_InvalidUTF8Error_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InvalidUTF8Error).Error(ixgo.DirectCallArg[*q.InvalidUTF8Error](ctx, 0)))
}

func method_ptr_InvalidUnmarshalError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.InvalidUnmarshalError).Error(ixgo.DirectCallArg[*q.InvalidUnmarshalError](ctx, 0)))
}

func method_ptr_MarshalerError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MarshalerError).Error(ixgo.DirectCallArg[*q.MarshalerError](ctx, 0)))
}

func method_ptr_MarshalerError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MarshalerError).Unwrap(ixgo.DirectCallArg[*q.MarshalerError](ctx, 0)))
}

func func_MatchCaseSensitiveDelimiter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MatchCaseSensitiveDelimiter(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_MergeWithLegacySemantics(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MergeWithLegacySemantics(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_NewDecoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewDecoder(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewEncoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewEncoder(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func method_Number_MarshalJSONTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Number.MarshalJSONTo(ixgo.DirectCallArg[q.Number](ctx, 0), ixgo.DirectCallArg[*jsontext.Encoder](ctx, 1)))
}

func method_ptr_Number_MarshalJSONTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Number).MarshalJSONTo(ixgo.DirectCallArg[*q.Number](ctx, 0), ixgo.DirectCallArg[*jsontext.Encoder](ctx, 1)))
}

func method_Number_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Number.String(ixgo.DirectCallArg[q.Number](ctx, 0)))
}

func method_ptr_Number_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Number).String(ixgo.DirectCallArg[*q.Number](ctx, 0)))
}

func method_ptr_Number_UnmarshalJSONFrom(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Number).UnmarshalJSONFrom(ixgo.DirectCallArg[*q.Number](ctx, 0), ixgo.DirectCallArg[*jsontext.Decoder](ctx, 1)))
}

func func_OmitEmptyWithLegacySemantics(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OmitEmptyWithLegacySemantics(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_ParseBytesWithLooseRFC4648(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ParseBytesWithLooseRFC4648(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_ParseTimeWithLooseRFC3339(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ParseTimeWithLooseRFC3339(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_ReportErrorsWithLegacySemantics(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ReportErrorsWithLegacySemantics(ixgo.DirectCallArg[bool](ctx, 0)))
}

func func_StringifyWithLegacySemantics(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StringifyWithLegacySemantics(ixgo.DirectCallArg[bool](ctx, 0)))
}

func method_ptr_SyntaxError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SyntaxError).Error(ixgo.DirectCallArg[*q.SyntaxError](ctx, 0)))
}

func func_Unmarshal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Unmarshal(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[interface{}](ctx, 1)))
}

func func_UnmarshalArrayFromAnyLength(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnmarshalArrayFromAnyLength(ixgo.DirectCallArg[bool](ctx, 0)))
}

func method_ptr_UnmarshalFieldError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnmarshalFieldError).Error(ixgo.DirectCallArg[*q.UnmarshalFieldError](ctx, 0)))
}

func method_ptr_UnmarshalTypeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnmarshalTypeError).Error(ixgo.DirectCallArg[*q.UnmarshalTypeError](ctx, 0)))
}

func method_ptr_UnmarshalTypeError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnmarshalTypeError).Unwrap(ixgo.DirectCallArg[*q.UnmarshalTypeError](ctx, 0)))
}

func method_ptr_UnsupportedTypeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnsupportedTypeError).Error(ixgo.DirectCallArg[*q.UnsupportedTypeError](ctx, 0)))
}

func method_ptr_UnsupportedValueError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnsupportedValueError).Error(ixgo.DirectCallArg[*q.UnsupportedValueError](ctx, 0)))
}

func func_Valid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Valid(ixgo.DirectCallArg[[]byte](ctx, 0)))
}
