// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package dwarf

import (
	q "debug/dwarf"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("debug/dwarf", map[string]ixgo.DirectCallAdapter{
		"(*ArrayType).Size":         method_ptr_ArrayType_Size,
		"(*ArrayType).String":       method_ptr_ArrayType_String,
		"(*Attr).GoString":          method_ptr_Attr_GoString,
		"(*Attr).String":            method_ptr_Attr_String,
		"(*BasicType).Basic":        method_ptr_BasicType_Basic,
		"(*BasicType).String":       method_ptr_BasicType_String,
		"(*Class).GoString":         method_ptr_Class_GoString,
		"(*Class).String":           method_ptr_Class_String,
		"(*CommonType).Common":      method_ptr_CommonType_Common,
		"(*CommonType).Size":        method_ptr_CommonType_Size,
		"(*Data).AddSection":        method_ptr_Data_AddSection,
		"(*Data).AddTypes":          method_ptr_Data_AddTypes,
		"(*Data).Reader":            method_ptr_Data_Reader,
		"(*DecodeError).Error":      method_ptr_DecodeError_Error,
		"(*DotDotDotType).String":   method_ptr_DotDotDotType_String,
		"(*Entry).AttrField":        method_ptr_Entry_AttrField,
		"(*Entry).Val":              method_ptr_Entry_Val,
		"(*EnumType).String":        method_ptr_EnumType_String,
		"(*FuncType).String":        method_ptr_FuncType_String,
		"(*LineReader).Files":       method_ptr_LineReader_Files,
		"(*LineReader).Next":        method_ptr_LineReader_Next,
		"(*LineReader).Reset":       method_ptr_LineReader_Reset,
		"(*LineReader).Seek":        method_ptr_LineReader_Seek,
		"(*LineReader).SeekPC":      method_ptr_LineReader_SeekPC,
		"(*LineReader).Tell":        method_ptr_LineReader_Tell,
		"(*PtrType).String":         method_ptr_PtrType_String,
		"(*QualType).Size":          method_ptr_QualType_Size,
		"(*QualType).String":        method_ptr_QualType_String,
		"(*Reader).AddressSize":     method_ptr_Reader_AddressSize,
		"(*Reader).ByteOrder":       method_ptr_Reader_ByteOrder,
		"(*Reader).Seek":            method_ptr_Reader_Seek,
		"(*Reader).SkipChildren":    method_ptr_Reader_SkipChildren,
		"(*StructType).Defn":        method_ptr_StructType_Defn,
		"(*StructType).String":      method_ptr_StructType_String,
		"(*Tag).GoString":           method_ptr_Tag_GoString,
		"(*Tag).String":             method_ptr_Tag_String,
		"(*TypedefType).Size":       method_ptr_TypedefType_Size,
		"(*TypedefType).String":     method_ptr_TypedefType_String,
		"(*UnsupportedType).String": method_ptr_UnsupportedType_String,
		"(*VoidType).String":        method_ptr_VoidType_String,
		"(Attr).GoString":           method_Attr_GoString,
		"(Attr).String":             method_Attr_String,
		"(Class).GoString":          method_Class_GoString,
		"(Class).String":            method_Class_String,
		"(DecodeError).Error":       method_DecodeError_Error,
		"(Tag).GoString":            method_Tag_GoString,
		"(Tag).String":              method_Tag_String,
	})
}

func method_ptr_ArrayType_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ArrayType).Size(ixgo.DirectCallArg[*q.ArrayType](ctx, 0)))
}

func method_ptr_ArrayType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ArrayType).String(ixgo.DirectCallArg[*q.ArrayType](ctx, 0)))
}

func method_Attr_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Attr.GoString(ixgo.DirectCallArg[q.Attr](ctx, 0)))
}

func method_ptr_Attr_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Attr).GoString(ixgo.DirectCallArg[*q.Attr](ctx, 0)))
}

func method_Attr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Attr.String(ixgo.DirectCallArg[q.Attr](ctx, 0)))
}

func method_ptr_Attr_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Attr).String(ixgo.DirectCallArg[*q.Attr](ctx, 0)))
}

func method_ptr_BasicType_Basic(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BasicType).Basic(ixgo.DirectCallArg[*q.BasicType](ctx, 0)))
}

func method_ptr_BasicType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.BasicType).String(ixgo.DirectCallArg[*q.BasicType](ctx, 0)))
}

func method_Class_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Class.GoString(ixgo.DirectCallArg[q.Class](ctx, 0)))
}

func method_ptr_Class_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Class).GoString(ixgo.DirectCallArg[*q.Class](ctx, 0)))
}

func method_Class_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Class.String(ixgo.DirectCallArg[q.Class](ctx, 0)))
}

func method_ptr_Class_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Class).String(ixgo.DirectCallArg[*q.Class](ctx, 0)))
}

func method_ptr_CommonType_Common(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommonType).Common(ixgo.DirectCallArg[*q.CommonType](ctx, 0)))
}

func method_ptr_CommonType_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CommonType).Size(ixgo.DirectCallArg[*q.CommonType](ctx, 0)))
}

func method_ptr_Data_AddSection(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Data).AddSection(ixgo.DirectCallArg[*q.Data](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}

func method_ptr_Data_AddTypes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Data).AddTypes(ixgo.DirectCallArg[*q.Data](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2)))
}

func method_ptr_Data_Reader(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Data).Reader(ixgo.DirectCallArg[*q.Data](ctx, 0)))
}

func method_DecodeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DecodeError.Error(ixgo.DirectCallArg[q.DecodeError](ctx, 0)))
}

func method_ptr_DecodeError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DecodeError).Error(ixgo.DirectCallArg[*q.DecodeError](ctx, 0)))
}

func method_ptr_DotDotDotType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DotDotDotType).String(ixgo.DirectCallArg[*q.DotDotDotType](ctx, 0)))
}

func method_ptr_Entry_AttrField(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Entry).AttrField(ixgo.DirectCallArg[*q.Entry](ctx, 0), ixgo.DirectCallArg[q.Attr](ctx, 1)))
}

func method_ptr_Entry_Val(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Entry).Val(ixgo.DirectCallArg[*q.Entry](ctx, 0), ixgo.DirectCallArg[q.Attr](ctx, 1)))
}

func method_ptr_EnumType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.EnumType).String(ixgo.DirectCallArg[*q.EnumType](ctx, 0)))
}

func method_ptr_FuncType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FuncType).String(ixgo.DirectCallArg[*q.FuncType](ctx, 0)))
}

func method_ptr_LineReader_Files(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LineReader).Files(ixgo.DirectCallArg[*q.LineReader](ctx, 0)))
}

func method_ptr_LineReader_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LineReader).Next(ixgo.DirectCallArg[*q.LineReader](ctx, 0), ixgo.DirectCallArg[*q.LineEntry](ctx, 1)))
}

func method_ptr_LineReader_Reset(ctx ixgo.DirectCallContext) {
	(*q.LineReader).Reset(ixgo.DirectCallArg[*q.LineReader](ctx, 0))
}

func method_ptr_LineReader_Seek(ctx ixgo.DirectCallContext) {
	(*q.LineReader).Seek(ixgo.DirectCallArg[*q.LineReader](ctx, 0), ixgo.DirectCallArg[q.LineReaderPos](ctx, 1))
}

func method_ptr_LineReader_SeekPC(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LineReader).SeekPC(ixgo.DirectCallArg[*q.LineReader](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1), ixgo.DirectCallArg[*q.LineEntry](ctx, 2)))
}

func method_ptr_LineReader_Tell(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LineReader).Tell(ixgo.DirectCallArg[*q.LineReader](ctx, 0)))
}

func method_ptr_PtrType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PtrType).String(ixgo.DirectCallArg[*q.PtrType](ctx, 0)))
}

func method_ptr_QualType_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QualType).Size(ixgo.DirectCallArg[*q.QualType](ctx, 0)))
}

func method_ptr_QualType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.QualType).String(ixgo.DirectCallArg[*q.QualType](ctx, 0)))
}

func method_ptr_Reader_AddressSize(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).AddressSize(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_ByteOrder(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).ByteOrder(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_Seek(ctx ixgo.DirectCallContext) {
	(*q.Reader).Seek(ixgo.DirectCallArg[*q.Reader](ctx, 0), ixgo.DirectCallArg[q.Offset](ctx, 1))
}

func method_ptr_Reader_SkipChildren(ctx ixgo.DirectCallContext) {
	(*q.Reader).SkipChildren(ixgo.DirectCallArg[*q.Reader](ctx, 0))
}

func method_ptr_StructType_Defn(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StructType).Defn(ixgo.DirectCallArg[*q.StructType](ctx, 0)))
}

func method_ptr_StructType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StructType).String(ixgo.DirectCallArg[*q.StructType](ctx, 0)))
}

func method_Tag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Tag.GoString(ixgo.DirectCallArg[q.Tag](ctx, 0)))
}

func method_ptr_Tag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tag).GoString(ixgo.DirectCallArg[*q.Tag](ctx, 0)))
}

func method_Tag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Tag.String(ixgo.DirectCallArg[q.Tag](ctx, 0)))
}

func method_ptr_Tag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tag).String(ixgo.DirectCallArg[*q.Tag](ctx, 0)))
}

func method_ptr_TypedefType_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypedefType).Size(ixgo.DirectCallArg[*q.TypedefType](ctx, 0)))
}

func method_ptr_TypedefType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypedefType).String(ixgo.DirectCallArg[*q.TypedefType](ctx, 0)))
}

func method_ptr_UnsupportedType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.UnsupportedType).String(ixgo.DirectCallArg[*q.UnsupportedType](ctx, 0)))
}

func method_ptr_VoidType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.VoidType).String(ixgo.DirectCallArg[*q.VoidType](ctx, 0)))
}
