// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package xml

import (
	q "encoding/xml"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("encoding/xml", map[string]ixgo.DirectCallAdapter{
		"(*CharData).Copy":              method_ptr_CharData_Copy,
		"(*Comment).Copy":               method_ptr_Comment_Copy,
		"(*Decoder).Decode":             method_ptr_Decoder_Decode,
		"(*Decoder).DecodeElement":      method_ptr_Decoder_DecodeElement,
		"(*Decoder).InputOffset":        method_ptr_Decoder_InputOffset,
		"(*Decoder).Skip":               method_ptr_Decoder_Skip,
		"(*Directive).Copy":             method_ptr_Directive_Copy,
		"(*Encoder).Close":              method_ptr_Encoder_Close,
		"(*Encoder).Encode":             method_ptr_Encoder_Encode,
		"(*Encoder).EncodeElement":      method_ptr_Encoder_EncodeElement,
		"(*Encoder).EncodeToken":        method_ptr_Encoder_EncodeToken,
		"(*Encoder).Flush":              method_ptr_Encoder_Flush,
		"(*Encoder).Indent":             method_ptr_Encoder_Indent,
		"(*ProcInst).Copy":              method_ptr_ProcInst_Copy,
		"(*StartElement).Copy":          method_ptr_StartElement_Copy,
		"(*StartElement).End":           method_ptr_StartElement_End,
		"(*SyntaxError).Error":          method_ptr_SyntaxError_Error,
		"(*TagPathError).Error":         method_ptr_TagPathError_Error,
		"(*UnmarshalError).Error":       method_ptr_UnmarshalError_Error,
		"(*UnsupportedTypeError).Error": method_ptr_UnsupportedTypeError_Error,
		"(CharData).Copy":               method_CharData_Copy,
		"(Comment).Copy":                method_Comment_Copy,
		"(Directive).Copy":              method_Directive_Copy,
		"(ProcInst).Copy":               method_ProcInst_Copy,
		"(StartElement).Copy":           method_StartElement_Copy,
		"(StartElement).End":            method_StartElement_End,
		"(UnmarshalError).Error":        method_UnmarshalError_Error,
		"CopyToken":                     func_CopyToken,
		"Escape":                        func_Escape,
		"EscapeText":                    func_EscapeText,
		"NewDecoder":                    func_NewDecoder,
		"NewEncoder":                    func_NewEncoder,
		"NewTokenDecoder":               func_NewTokenDecoder,
		"Unmarshal":                     func_Unmarshal,
	})
}

func method_CharData_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CharData.Copy(ixgo.DirectCallArg[q.CharData](ctx, 0)))
}

func method_ptr_CharData_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CharData).Copy(ixgo.DirectCallArg[*q.CharData](ctx, 0)))
}

func method_Comment_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Comment.Copy(ixgo.DirectCallArg[q.Comment](ctx, 0)))
}

func method_ptr_Comment_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Comment).Copy(ixgo.DirectCallArg[*q.Comment](ctx, 0)))
}

func func_CopyToken(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CopyToken(ixgo.DirectCallArg[q.Token](ctx, 0)))
}

func method_ptr_Decoder_Decode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).Decode(ixgo.DirectCallArg[*q.Decoder](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_Decoder_DecodeElement(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).DecodeElement(ixgo.DirectCallArg[*q.Decoder](ctx, 0), ixgo.DirectCallArg[any](ctx, 1), ixgo.DirectCallArg[*q.StartElement](ctx, 2)))
}

func method_ptr_Decoder_InputOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).InputOffset(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_ptr_Decoder_Skip(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Decoder).Skip(ixgo.DirectCallArg[*q.Decoder](ctx, 0)))
}

func method_Directive_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Directive.Copy(ixgo.DirectCallArg[q.Directive](ctx, 0)))
}

func method_ptr_Directive_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Directive).Copy(ixgo.DirectCallArg[*q.Directive](ctx, 0)))
}

func method_ptr_Encoder_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).Close(ixgo.DirectCallArg[*q.Encoder](ctx, 0)))
}

func method_ptr_Encoder_Encode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).Encode(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_ptr_Encoder_EncodeElement(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).EncodeElement(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[any](ctx, 1), ixgo.DirectCallArg[q.StartElement](ctx, 2)))
}

func method_ptr_Encoder_EncodeToken(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).EncodeToken(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[q.Token](ctx, 1)))
}

func method_ptr_Encoder_Flush(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Encoder).Flush(ixgo.DirectCallArg[*q.Encoder](ctx, 0)))
}

func method_ptr_Encoder_Indent(ctx ixgo.DirectCallContext) {
	(*q.Encoder).Indent(ixgo.DirectCallArg[*q.Encoder](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func func_Escape(ctx ixgo.DirectCallContext) {
	q.Escape(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1))
}

func func_EscapeText(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.EscapeText(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_NewDecoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewDecoder(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewEncoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewEncoder(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func func_NewTokenDecoder(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewTokenDecoder(ixgo.DirectCallArg[q.TokenReader](ctx, 0)))
}

func method_ProcInst_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ProcInst.Copy(ixgo.DirectCallArg[q.ProcInst](ctx, 0)))
}

func method_ptr_ProcInst_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcInst).Copy(ixgo.DirectCallArg[*q.ProcInst](ctx, 0)))
}

func method_StartElement_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StartElement.Copy(ixgo.DirectCallArg[q.StartElement](ctx, 0)))
}

func method_ptr_StartElement_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StartElement).Copy(ixgo.DirectCallArg[*q.StartElement](ctx, 0)))
}

func method_StartElement_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StartElement.End(ixgo.DirectCallArg[q.StartElement](ctx, 0)))
}

func method_ptr_StartElement_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StartElement).End(ixgo.DirectCallArg[*q.StartElement](ctx, 0)))
}

func method_ptr_SyntaxError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SyntaxError).Error(ixgo.DirectCallArg[*q.SyntaxError](ctx, 0)))
}

func method_ptr_TagPathError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TagPathError).Error(ixgo.DirectCallArg[*q.TagPathError](ctx, 0)))
}

func func_Unmarshal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Unmarshal(ixgo.DirectCallArg[[]byte](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func method_UnmarshalError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UnmarshalError.Error(ixgo.DirectCallArg[q.UnmarshalError](ctx, 0)))
}

func method_ptr_UnmarshalError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnmarshalError).Error(ixgo.DirectCallArg[*q.UnmarshalError](ctx, 0)))
}

func method_ptr_UnsupportedTypeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnsupportedTypeError).Error(ixgo.DirectCallArg[*q.UnsupportedTypeError](ctx, 0)))
}
