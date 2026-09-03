// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && !go1.27
// +build go1.26,!go1.27

package reflect

import (
	q "reflect"

	"github.com/goplus/ixgo"
	unsafe "unsafe"
)

func init() {
	ixgo.RegisterDirectCalls("reflect", map[string]ixgo.DirectCallAdapter{
		"(*ChanDir).String":         method_ptr_ChanDir_String,
		"(*Kind).String":            method_ptr_Kind_String,
		"(*MapIter).Key":            method_ptr_MapIter_Key,
		"(*MapIter).Next":           method_ptr_MapIter_Next,
		"(*MapIter).Reset":          method_ptr_MapIter_Reset,
		"(*MapIter).Value":          method_ptr_MapIter_Value,
		"(*Method).IsExported":      method_ptr_Method_IsExported,
		"(*StructField).IsExported": method_ptr_StructField_IsExported,
		"(*StructTag).Get":          method_ptr_StructTag_Get,
		"(*Value).Addr":             method_ptr_Value_Addr,
		"(*Value).Bool":             method_ptr_Value_Bool,
		"(*Value).Bytes":            method_ptr_Value_Bytes,
		"(*Value).Call":             method_ptr_Value_Call,
		"(*Value).CallSlice":        method_ptr_Value_CallSlice,
		"(*Value).CanAddr":          method_ptr_Value_CanAddr,
		"(*Value).CanComplex":       method_ptr_Value_CanComplex,
		"(*Value).CanConvert":       method_ptr_Value_CanConvert,
		"(*Value).CanFloat":         method_ptr_Value_CanFloat,
		"(*Value).CanInt":           method_ptr_Value_CanInt,
		"(*Value).CanInterface":     method_ptr_Value_CanInterface,
		"(*Value).CanSet":           method_ptr_Value_CanSet,
		"(*Value).CanUint":          method_ptr_Value_CanUint,
		"(*Value).Cap":              method_ptr_Value_Cap,
		"(*Value).Clear":            method_ptr_Value_Clear,
		"(*Value).Close":            method_ptr_Value_Close,
		"(*Value).Comparable":       method_ptr_Value_Comparable,
		"(*Value).Complex":          method_ptr_Value_Complex,
		"(*Value).Convert":          method_ptr_Value_Convert,
		"(*Value).Elem":             method_ptr_Value_Elem,
		"(*Value).Equal":            method_ptr_Value_Equal,
		"(*Value).Field":            method_ptr_Value_Field,
		"(*Value).FieldByIndex":     method_ptr_Value_FieldByIndex,
		"(*Value).FieldByName":      method_ptr_Value_FieldByName,
		"(*Value).FieldByNameFunc":  method_ptr_Value_FieldByNameFunc,
		"(*Value).Fields":           method_ptr_Value_Fields,
		"(*Value).Float":            method_ptr_Value_Float,
		"(*Value).Grow":             method_ptr_Value_Grow,
		"(*Value).Index":            method_ptr_Value_Index,
		"(*Value).Int":              method_ptr_Value_Int,
		"(*Value).Interface":        method_ptr_Value_Interface,
		"(*Value).InterfaceData":    method_ptr_Value_InterfaceData,
		"(*Value).IsNil":            method_ptr_Value_IsNil,
		"(*Value).IsValid":          method_ptr_Value_IsValid,
		"(*Value).IsZero":           method_ptr_Value_IsZero,
		"(*Value).Kind":             method_ptr_Value_Kind,
		"(*Value).Len":              method_ptr_Value_Len,
		"(*Value).MapIndex":         method_ptr_Value_MapIndex,
		"(*Value).MapKeys":          method_ptr_Value_MapKeys,
		"(*Value).MapRange":         method_ptr_Value_MapRange,
		"(*Value).Method":           method_ptr_Value_Method,
		"(*Value).MethodByName":     method_ptr_Value_MethodByName,
		"(*Value).Methods":          method_ptr_Value_Methods,
		"(*Value).NumField":         method_ptr_Value_NumField,
		"(*Value).NumMethod":        method_ptr_Value_NumMethod,
		"(*Value).OverflowComplex":  method_ptr_Value_OverflowComplex,
		"(*Value).OverflowFloat":    method_ptr_Value_OverflowFloat,
		"(*Value).OverflowInt":      method_ptr_Value_OverflowInt,
		"(*Value).OverflowUint":     method_ptr_Value_OverflowUint,
		"(*Value).Pointer":          method_ptr_Value_Pointer,
		"(*Value).Send":             method_ptr_Value_Send,
		"(*Value).Seq":              method_ptr_Value_Seq,
		"(*Value).Seq2":             method_ptr_Value_Seq2,
		"(*Value).Set":              method_ptr_Value_Set,
		"(*Value).SetBool":          method_ptr_Value_SetBool,
		"(*Value).SetBytes":         method_ptr_Value_SetBytes,
		"(*Value).SetCap":           method_ptr_Value_SetCap,
		"(*Value).SetComplex":       method_ptr_Value_SetComplex,
		"(*Value).SetFloat":         method_ptr_Value_SetFloat,
		"(*Value).SetInt":           method_ptr_Value_SetInt,
		"(*Value).SetIterKey":       method_ptr_Value_SetIterKey,
		"(*Value).SetIterValue":     method_ptr_Value_SetIterValue,
		"(*Value).SetLen":           method_ptr_Value_SetLen,
		"(*Value).SetMapIndex":      method_ptr_Value_SetMapIndex,
		"(*Value).SetPointer":       method_ptr_Value_SetPointer,
		"(*Value).SetString":        method_ptr_Value_SetString,
		"(*Value).SetUint":          method_ptr_Value_SetUint,
		"(*Value).SetZero":          method_ptr_Value_SetZero,
		"(*Value).Slice":            method_ptr_Value_Slice,
		"(*Value).Slice3":           method_ptr_Value_Slice3,
		"(*Value).String":           method_ptr_Value_String,
		"(*Value).TrySend":          method_ptr_Value_TrySend,
		"(*Value).Type":             method_ptr_Value_Type,
		"(*Value).Uint":             method_ptr_Value_Uint,
		"(*Value).UnsafeAddr":       method_ptr_Value_UnsafeAddr,
		"(*Value).UnsafePointer":    method_ptr_Value_UnsafePointer,
		"(*ValueError).Error":       method_ptr_ValueError_Error,
		"(ChanDir).String":          method_ChanDir_String,
		"(Kind).String":             method_Kind_String,
		"(Method).IsExported":       method_Method_IsExported,
		"(StructField).IsExported":  method_StructField_IsExported,
		"(StructTag).Get":           method_StructTag_Get,
		"(Value).Addr":              method_Value_Addr,
		"(Value).Bool":              method_Value_Bool,
		"(Value).Bytes":             method_Value_Bytes,
		"(Value).Call":              method_Value_Call,
		"(Value).CallSlice":         method_Value_CallSlice,
		"(Value).CanAddr":           method_Value_CanAddr,
		"(Value).CanComplex":        method_Value_CanComplex,
		"(Value).CanConvert":        method_Value_CanConvert,
		"(Value).CanFloat":          method_Value_CanFloat,
		"(Value).CanInt":            method_Value_CanInt,
		"(Value).CanInterface":      method_Value_CanInterface,
		"(Value).CanSet":            method_Value_CanSet,
		"(Value).CanUint":           method_Value_CanUint,
		"(Value).Cap":               method_Value_Cap,
		"(Value).Clear":             method_Value_Clear,
		"(Value).Close":             method_Value_Close,
		"(Value).Comparable":        method_Value_Comparable,
		"(Value).Complex":           method_Value_Complex,
		"(Value).Convert":           method_Value_Convert,
		"(Value).Elem":              method_Value_Elem,
		"(Value).Equal":             method_Value_Equal,
		"(Value).Field":             method_Value_Field,
		"(Value).FieldByIndex":      method_Value_FieldByIndex,
		"(Value).FieldByName":       method_Value_FieldByName,
		"(Value).FieldByNameFunc":   method_Value_FieldByNameFunc,
		"(Value).Fields":            method_Value_Fields,
		"(Value).Float":             method_Value_Float,
		"(Value).Grow":              method_Value_Grow,
		"(Value).Index":             method_Value_Index,
		"(Value).Int":               method_Value_Int,
		"(Value).Interface":         method_Value_Interface,
		"(Value).InterfaceData":     method_Value_InterfaceData,
		"(Value).IsNil":             method_Value_IsNil,
		"(Value).IsValid":           method_Value_IsValid,
		"(Value).IsZero":            method_Value_IsZero,
		"(Value).Kind":              method_Value_Kind,
		"(Value).Len":               method_Value_Len,
		"(Value).MapIndex":          method_Value_MapIndex,
		"(Value).MapKeys":           method_Value_MapKeys,
		"(Value).MapRange":          method_Value_MapRange,
		"(Value).Method":            method_Value_Method,
		"(Value).MethodByName":      method_Value_MethodByName,
		"(Value).Methods":           method_Value_Methods,
		"(Value).NumField":          method_Value_NumField,
		"(Value).NumMethod":         method_Value_NumMethod,
		"(Value).OverflowComplex":   method_Value_OverflowComplex,
		"(Value).OverflowFloat":     method_Value_OverflowFloat,
		"(Value).OverflowInt":       method_Value_OverflowInt,
		"(Value).OverflowUint":      method_Value_OverflowUint,
		"(Value).Pointer":           method_Value_Pointer,
		"(Value).Send":              method_Value_Send,
		"(Value).Seq":               method_Value_Seq,
		"(Value).Seq2":              method_Value_Seq2,
		"(Value).Set":               method_Value_Set,
		"(Value).SetBool":           method_Value_SetBool,
		"(Value).SetBytes":          method_Value_SetBytes,
		"(Value).SetCap":            method_Value_SetCap,
		"(Value).SetComplex":        method_Value_SetComplex,
		"(Value).SetFloat":          method_Value_SetFloat,
		"(Value).SetInt":            method_Value_SetInt,
		"(Value).SetIterKey":        method_Value_SetIterKey,
		"(Value).SetIterValue":      method_Value_SetIterValue,
		"(Value).SetLen":            method_Value_SetLen,
		"(Value).SetMapIndex":       method_Value_SetMapIndex,
		"(Value).SetPointer":        method_Value_SetPointer,
		"(Value).SetString":         method_Value_SetString,
		"(Value).SetUint":           method_Value_SetUint,
		"(Value).SetZero":           method_Value_SetZero,
		"(Value).Slice":             method_Value_Slice,
		"(Value).Slice3":            method_Value_Slice3,
		"(Value).String":            method_Value_String,
		"(Value).TrySend":           method_Value_TrySend,
		"(Value).Type":              method_Value_Type,
		"(Value).Uint":              method_Value_Uint,
		"(Value).UnsafeAddr":        method_Value_UnsafeAddr,
		"(Value).UnsafePointer":     method_Value_UnsafePointer,
		"Append":                    func_Append,
		"AppendSlice":               func_AppendSlice,
		"ArrayOf":                   func_ArrayOf,
		"ChanOf":                    func_ChanOf,
		"Copy":                      func_Copy,
		"DeepEqual":                 func_DeepEqual,
		"FuncOf":                    func_FuncOf,
		"Indirect":                  func_Indirect,
		"MakeChan":                  func_MakeChan,
		"MakeFunc":                  func_MakeFunc,
		"MakeMap":                   func_MakeMap,
		"MakeMapWithSize":           func_MakeMapWithSize,
		"MakeSlice":                 func_MakeSlice,
		"MapOf":                     func_MapOf,
		"New":                       func_New,
		"NewAt":                     func_NewAt,
		"PointerTo":                 func_PointerTo,
		"PtrTo":                     func_PtrTo,
		"SliceAt":                   func_SliceAt,
		"SliceOf":                   func_SliceOf,
		"StructOf":                  func_StructOf,
		"Swapper":                   func_Swapper,
		"TypeOf":                    func_TypeOf,
		"ValueOf":                   func_ValueOf,
		"VisibleFields":             func_VisibleFields,
		"Zero":                      func_Zero,
	})
}

func func_Append(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Append(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Value](ctx, 1)...))
}

func func_AppendSlice(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AppendSlice(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1)))
}

func func_ArrayOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ArrayOf(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func method_ChanDir_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ChanDir.String(ixgo.DirectCallArg[q.ChanDir](ctx, 0)))
}

func method_ptr_ChanDir_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ChanDir).String(ixgo.DirectCallArg[*q.ChanDir](ctx, 0)))
}

func func_ChanOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ChanOf(ixgo.DirectCallArg[q.ChanDir](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Copy(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1)))
}

func func_DeepEqual(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DeepEqual(ixgo.DirectCallArg[any](ctx, 0), ixgo.DirectCallArg[any](ctx, 1)))
}

func func_FuncOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.FuncOf(ixgo.DirectCallArg[[]q.Type](ctx, 0), ixgo.DirectCallArg[[]q.Type](ctx, 1), ixgo.DirectCallArg[bool](ctx, 2)))
}

func func_Indirect(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Indirect(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_Kind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Kind.String(ixgo.DirectCallArg[q.Kind](ctx, 0)))
}

func method_ptr_Kind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Kind).String(ixgo.DirectCallArg[*q.Kind](ctx, 0)))
}

func func_MakeChan(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeChan(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_MakeFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeFunc(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[func(args []q.Value) (results []q.Value)](ctx, 1)))
}

func func_MakeMap(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeMap(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func func_MakeMapWithSize(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeMapWithSize(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_MakeSlice(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MakeSlice(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_MapIter_Key(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MapIter).Key(ixgo.DirectCallArg[*q.MapIter](ctx, 0)))
}

func method_ptr_MapIter_Next(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MapIter).Next(ixgo.DirectCallArg[*q.MapIter](ctx, 0)))
}

func method_ptr_MapIter_Reset(ctx ixgo.DirectCallContext) {
	(*q.MapIter).Reset(ixgo.DirectCallArg[*q.MapIter](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1))
}

func method_ptr_MapIter_Value(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MapIter).Value(ixgo.DirectCallArg[*q.MapIter](ctx, 0)))
}

func func_MapOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MapOf(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func method_Method_IsExported(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Method.IsExported(ixgo.DirectCallArg[q.Method](ctx, 0)))
}

func method_ptr_Method_IsExported(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Method).IsExported(ixgo.DirectCallArg[*q.Method](ctx, 0)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func func_NewAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewAt(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[unsafe.Pointer](ctx, 1)))
}

func func_PointerTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PointerTo(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func func_PtrTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.PtrTo(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func func_SliceAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SliceAt(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[unsafe.Pointer](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_SliceOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SliceOf(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func method_StructField_IsExported(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StructField.IsExported(ixgo.DirectCallArg[q.StructField](ctx, 0)))
}

func method_ptr_StructField_IsExported(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StructField).IsExported(ixgo.DirectCallArg[*q.StructField](ctx, 0)))
}

func func_StructOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StructOf(ixgo.DirectCallArg[[]q.StructField](ctx, 0)))
}

func method_StructTag_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.StructTag.Get(ixgo.DirectCallArg[q.StructTag](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_StructTag_Get(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StructTag).Get(ixgo.DirectCallArg[*q.StructTag](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Swapper(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Swapper(ixgo.DirectCallArg[any](ctx, 0)))
}

func func_TypeOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeOf(ixgo.DirectCallArg[any](ctx, 0)))
}

func method_Value_Addr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Addr(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Addr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Addr(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Bool(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Bool(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Bool(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Bool(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Bytes(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Bytes(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Call(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Call(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Value](ctx, 1)))
}

func method_ptr_Value_Call(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Call(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Value](ctx, 1)))
}

func method_Value_CallSlice(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.CallSlice(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Value](ctx, 1)))
}

func method_ptr_Value_CallSlice(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CallSlice(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]q.Value](ctx, 1)))
}

func method_Value_CanAddr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.CanAddr(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_CanAddr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CanAddr(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_CanComplex(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.CanComplex(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_CanComplex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CanComplex(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_CanConvert(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.CanConvert(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func method_ptr_Value_CanConvert(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CanConvert(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func method_Value_CanFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.CanFloat(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_CanFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CanFloat(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_CanInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.CanInt(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_CanInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CanInt(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_CanInterface(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.CanInterface(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_CanInterface(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CanInterface(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_CanSet(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.CanSet(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_CanSet(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CanSet(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_CanUint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.CanUint(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_CanUint(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).CanUint(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Cap(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Cap(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Cap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Cap(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Clear(ctx ixgo.DirectCallContext) {
	q.Value.Clear(ixgo.DirectCallArg[q.Value](ctx, 0))
}

func method_ptr_Value_Clear(ctx ixgo.DirectCallContext) {
	(*q.Value).Clear(ixgo.DirectCallArg[*q.Value](ctx, 0))
}

func method_Value_Close(ctx ixgo.DirectCallContext) {
	q.Value.Close(ixgo.DirectCallArg[q.Value](ctx, 0))
}

func method_ptr_Value_Close(ctx ixgo.DirectCallContext) {
	(*q.Value).Close(ixgo.DirectCallArg[*q.Value](ctx, 0))
}

func method_Value_Comparable(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Comparable(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Comparable(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Comparable(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Complex(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Complex(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Complex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Complex(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Convert(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Convert(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func method_ptr_Value_Convert(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Convert(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func method_Value_Elem(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Elem(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Elem(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Elem(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Equal(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1)))
}

func method_ptr_Value_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Equal(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1)))
}

func method_Value_Field(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Field(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Value_Field(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Field(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_Value_FieldByIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.FieldByIndex(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[[]int](ctx, 1)))
}

func method_ptr_Value_FieldByIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).FieldByIndex(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]int](ctx, 1)))
}

func method_Value_FieldByName(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.FieldByName(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Value_FieldByName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).FieldByName(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_Value_FieldByNameFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.FieldByNameFunc(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[func(string) bool](ctx, 1)))
}

func method_ptr_Value_FieldByNameFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).FieldByNameFunc(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[func(string) bool](ctx, 1)))
}

func method_Value_Fields(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Fields(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Fields(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Fields(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Float(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Float(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Float(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Float(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Grow(ctx ixgo.DirectCallContext) {
	q.Value.Grow(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_Value_Grow(ctx ixgo.DirectCallContext) {
	(*q.Value).Grow(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_Value_Index(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Index(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Value_Index(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Index(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_Value_Int(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Int(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Int(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Int(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Interface(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Interface(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Interface(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Interface(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_InterfaceData(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.InterfaceData(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_InterfaceData(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).InterfaceData(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_IsNil(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.IsNil(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_IsNil(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).IsNil(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.IsValid(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_IsValid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).IsValid(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_IsZero(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.IsZero(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_IsZero(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).IsZero(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Kind(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Kind(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Len(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Len(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_MapIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.MapIndex(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1)))
}

func method_ptr_Value_MapIndex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).MapIndex(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1)))
}

func method_Value_MapKeys(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.MapKeys(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_MapKeys(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).MapKeys(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_MapRange(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.MapRange(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_MapRange(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).MapRange(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Method(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Method(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Value_Method(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Method(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_Value_MethodByName(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.MethodByName(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Value_MethodByName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).MethodByName(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_Value_Methods(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Methods(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Methods(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Methods(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_NumField(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.NumField(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_NumField(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).NumField(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_NumMethod(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.NumMethod(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_NumMethod(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).NumMethod(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_OverflowComplex(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.OverflowComplex(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[complex128](ctx, 1)))
}

func method_ptr_Value_OverflowComplex(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).OverflowComplex(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[complex128](ctx, 1)))
}

func method_Value_OverflowFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.OverflowFloat(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func method_ptr_Value_OverflowFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).OverflowFloat(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func method_Value_OverflowInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.OverflowInt(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Value_OverflowInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).OverflowInt(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_Value_OverflowUint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.OverflowUint(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Value_OverflowUint(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).OverflowUint(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_Value_Pointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Pointer(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Pointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Pointer(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Send(ctx ixgo.DirectCallContext) {
	q.Value.Send(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1))
}

func method_ptr_Value_Send(ctx ixgo.DirectCallContext) {
	(*q.Value).Send(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1))
}

func method_Value_Seq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Seq(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Seq(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Seq(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Seq2(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Seq2(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Seq2(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Seq2(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Set(ctx ixgo.DirectCallContext) {
	q.Value.Set(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1))
}

func method_ptr_Value_Set(ctx ixgo.DirectCallContext) {
	(*q.Value).Set(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1))
}

func method_Value_SetBool(ctx ixgo.DirectCallContext) {
	q.Value.SetBool(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_ptr_Value_SetBool(ctx ixgo.DirectCallContext) {
	(*q.Value).SetBool(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1))
}

func method_Value_SetBytes(ctx ixgo.DirectCallContext) {
	q.Value.SetBytes(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1))
}

func method_ptr_Value_SetBytes(ctx ixgo.DirectCallContext) {
	(*q.Value).SetBytes(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1))
}

func method_Value_SetCap(ctx ixgo.DirectCallContext) {
	q.Value.SetCap(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_Value_SetCap(ctx ixgo.DirectCallContext) {
	(*q.Value).SetCap(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_Value_SetComplex(ctx ixgo.DirectCallContext) {
	q.Value.SetComplex(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[complex128](ctx, 1))
}

func method_ptr_Value_SetComplex(ctx ixgo.DirectCallContext) {
	(*q.Value).SetComplex(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[complex128](ctx, 1))
}

func method_Value_SetFloat(ctx ixgo.DirectCallContext) {
	q.Value.SetFloat(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1))
}

func method_ptr_Value_SetFloat(ctx ixgo.DirectCallContext) {
	(*q.Value).SetFloat(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1))
}

func method_Value_SetInt(ctx ixgo.DirectCallContext) {
	q.Value.SetInt(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1))
}

func method_ptr_Value_SetInt(ctx ixgo.DirectCallContext) {
	(*q.Value).SetInt(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1))
}

func method_Value_SetIterKey(ctx ixgo.DirectCallContext) {
	q.Value.SetIterKey(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[*q.MapIter](ctx, 1))
}

func method_ptr_Value_SetIterKey(ctx ixgo.DirectCallContext) {
	(*q.Value).SetIterKey(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[*q.MapIter](ctx, 1))
}

func method_Value_SetIterValue(ctx ixgo.DirectCallContext) {
	q.Value.SetIterValue(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[*q.MapIter](ctx, 1))
}

func method_ptr_Value_SetIterValue(ctx ixgo.DirectCallContext) {
	(*q.Value).SetIterValue(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[*q.MapIter](ctx, 1))
}

func method_Value_SetLen(ctx ixgo.DirectCallContext) {
	q.Value.SetLen(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_ptr_Value_SetLen(ctx ixgo.DirectCallContext) {
	(*q.Value).SetLen(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1))
}

func method_Value_SetMapIndex(ctx ixgo.DirectCallContext) {
	q.Value.SetMapIndex(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1), ixgo.DirectCallArg[q.Value](ctx, 2))
}

func method_ptr_Value_SetMapIndex(ctx ixgo.DirectCallContext) {
	(*q.Value).SetMapIndex(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1), ixgo.DirectCallArg[q.Value](ctx, 2))
}

func method_Value_SetPointer(ctx ixgo.DirectCallContext) {
	q.Value.SetPointer(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[unsafe.Pointer](ctx, 1))
}

func method_ptr_Value_SetPointer(ctx ixgo.DirectCallContext) {
	(*q.Value).SetPointer(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[unsafe.Pointer](ctx, 1))
}

func method_Value_SetString(ctx ixgo.DirectCallContext) {
	q.Value.SetString(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_Value_SetString(ctx ixgo.DirectCallContext) {
	(*q.Value).SetString(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_Value_SetUint(ctx ixgo.DirectCallContext) {
	q.Value.SetUint(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1))
}

func method_ptr_Value_SetUint(ctx ixgo.DirectCallContext) {
	(*q.Value).SetUint(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1))
}

func method_Value_SetZero(ctx ixgo.DirectCallContext) {
	q.Value.SetZero(ixgo.DirectCallArg[q.Value](ctx, 0))
}

func method_ptr_Value_SetZero(ctx ixgo.DirectCallContext) {
	(*q.Value).SetZero(ixgo.DirectCallArg[*q.Value](ctx, 0))
}

func method_Value_Slice(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Slice(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Value_Slice(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Slice(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_Value_Slice3(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Slice3(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func method_ptr_Value_Slice3(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Slice3(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func method_Value_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.String(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).String(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_TrySend(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.TrySend(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1)))
}

func method_ptr_Value_TrySend(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).TrySend(ixgo.DirectCallArg[*q.Value](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1)))
}

func method_Value_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Type(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Type(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_Uint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.Uint(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_Uint(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).Uint(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_UnsafeAddr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.UnsafeAddr(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_UnsafeAddr(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).UnsafeAddr(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_Value_UnsafePointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Value.UnsafePointer(ixgo.DirectCallArg[q.Value](ctx, 0)))
}

func method_ptr_Value_UnsafePointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Value).UnsafePointer(ixgo.DirectCallArg[*q.Value](ctx, 0)))
}

func method_ptr_ValueError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ValueError).Error(ixgo.DirectCallArg[*q.ValueError](ctx, 0)))
}

func func_ValueOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ValueOf(ixgo.DirectCallArg[any](ctx, 0)))
}

func func_VisibleFields(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VisibleFields(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func func_Zero(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Zero(ixgo.DirectCallArg[q.Type](ctx, 0)))
}
