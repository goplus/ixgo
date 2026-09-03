// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package image

import (
	q "image"

	"github.com/goplus/ixgo"
	color "image/color"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("image", map[string]ixgo.DirectCallAdapter{
		"(*Alpha).AlphaAt":              method_ptr_Alpha_AlphaAt,
		"(*Alpha).At":                   method_ptr_Alpha_At,
		"(*Alpha).Bounds":               method_ptr_Alpha_Bounds,
		"(*Alpha).ColorModel":           method_ptr_Alpha_ColorModel,
		"(*Alpha).Opaque":               method_ptr_Alpha_Opaque,
		"(*Alpha).PixOffset":            method_ptr_Alpha_PixOffset,
		"(*Alpha).RGBA64At":             method_ptr_Alpha_RGBA64At,
		"(*Alpha).Set":                  method_ptr_Alpha_Set,
		"(*Alpha).SetAlpha":             method_ptr_Alpha_SetAlpha,
		"(*Alpha).SetRGBA64":            method_ptr_Alpha_SetRGBA64,
		"(*Alpha).SubImage":             method_ptr_Alpha_SubImage,
		"(*Alpha16).Alpha16At":          method_ptr_Alpha16_Alpha16At,
		"(*Alpha16).At":                 method_ptr_Alpha16_At,
		"(*Alpha16).Bounds":             method_ptr_Alpha16_Bounds,
		"(*Alpha16).ColorModel":         method_ptr_Alpha16_ColorModel,
		"(*Alpha16).Opaque":             method_ptr_Alpha16_Opaque,
		"(*Alpha16).PixOffset":          method_ptr_Alpha16_PixOffset,
		"(*Alpha16).RGBA64At":           method_ptr_Alpha16_RGBA64At,
		"(*Alpha16).Set":                method_ptr_Alpha16_Set,
		"(*Alpha16).SetAlpha16":         method_ptr_Alpha16_SetAlpha16,
		"(*Alpha16).SetRGBA64":          method_ptr_Alpha16_SetRGBA64,
		"(*Alpha16).SubImage":           method_ptr_Alpha16_SubImage,
		"(*CMYK).At":                    method_ptr_CMYK_At,
		"(*CMYK).Bounds":                method_ptr_CMYK_Bounds,
		"(*CMYK).CMYKAt":                method_ptr_CMYK_CMYKAt,
		"(*CMYK).ColorModel":            method_ptr_CMYK_ColorModel,
		"(*CMYK).Opaque":                method_ptr_CMYK_Opaque,
		"(*CMYK).PixOffset":             method_ptr_CMYK_PixOffset,
		"(*CMYK).RGBA64At":              method_ptr_CMYK_RGBA64At,
		"(*CMYK).Set":                   method_ptr_CMYK_Set,
		"(*CMYK).SetCMYK":               method_ptr_CMYK_SetCMYK,
		"(*CMYK).SetRGBA64":             method_ptr_CMYK_SetRGBA64,
		"(*CMYK).SubImage":              method_ptr_CMYK_SubImage,
		"(*Gray).At":                    method_ptr_Gray_At,
		"(*Gray).Bounds":                method_ptr_Gray_Bounds,
		"(*Gray).ColorModel":            method_ptr_Gray_ColorModel,
		"(*Gray).GrayAt":                method_ptr_Gray_GrayAt,
		"(*Gray).Opaque":                method_ptr_Gray_Opaque,
		"(*Gray).PixOffset":             method_ptr_Gray_PixOffset,
		"(*Gray).RGBA64At":              method_ptr_Gray_RGBA64At,
		"(*Gray).Set":                   method_ptr_Gray_Set,
		"(*Gray).SetGray":               method_ptr_Gray_SetGray,
		"(*Gray).SetRGBA64":             method_ptr_Gray_SetRGBA64,
		"(*Gray).SubImage":              method_ptr_Gray_SubImage,
		"(*Gray16).At":                  method_ptr_Gray16_At,
		"(*Gray16).Bounds":              method_ptr_Gray16_Bounds,
		"(*Gray16).ColorModel":          method_ptr_Gray16_ColorModel,
		"(*Gray16).Gray16At":            method_ptr_Gray16_Gray16At,
		"(*Gray16).Opaque":              method_ptr_Gray16_Opaque,
		"(*Gray16).PixOffset":           method_ptr_Gray16_PixOffset,
		"(*Gray16).RGBA64At":            method_ptr_Gray16_RGBA64At,
		"(*Gray16).Set":                 method_ptr_Gray16_Set,
		"(*Gray16).SetGray16":           method_ptr_Gray16_SetGray16,
		"(*Gray16).SetRGBA64":           method_ptr_Gray16_SetRGBA64,
		"(*Gray16).SubImage":            method_ptr_Gray16_SubImage,
		"(*NRGBA).At":                   method_ptr_NRGBA_At,
		"(*NRGBA).Bounds":               method_ptr_NRGBA_Bounds,
		"(*NRGBA).ColorModel":           method_ptr_NRGBA_ColorModel,
		"(*NRGBA).NRGBAAt":              method_ptr_NRGBA_NRGBAAt,
		"(*NRGBA).Opaque":               method_ptr_NRGBA_Opaque,
		"(*NRGBA).PixOffset":            method_ptr_NRGBA_PixOffset,
		"(*NRGBA).RGBA64At":             method_ptr_NRGBA_RGBA64At,
		"(*NRGBA).Set":                  method_ptr_NRGBA_Set,
		"(*NRGBA).SetNRGBA":             method_ptr_NRGBA_SetNRGBA,
		"(*NRGBA).SetRGBA64":            method_ptr_NRGBA_SetRGBA64,
		"(*NRGBA).SubImage":             method_ptr_NRGBA_SubImage,
		"(*NRGBA64).At":                 method_ptr_NRGBA64_At,
		"(*NRGBA64).Bounds":             method_ptr_NRGBA64_Bounds,
		"(*NRGBA64).ColorModel":         method_ptr_NRGBA64_ColorModel,
		"(*NRGBA64).NRGBA64At":          method_ptr_NRGBA64_NRGBA64At,
		"(*NRGBA64).Opaque":             method_ptr_NRGBA64_Opaque,
		"(*NRGBA64).PixOffset":          method_ptr_NRGBA64_PixOffset,
		"(*NRGBA64).RGBA64At":           method_ptr_NRGBA64_RGBA64At,
		"(*NRGBA64).Set":                method_ptr_NRGBA64_Set,
		"(*NRGBA64).SetNRGBA64":         method_ptr_NRGBA64_SetNRGBA64,
		"(*NRGBA64).SetRGBA64":          method_ptr_NRGBA64_SetRGBA64,
		"(*NRGBA64).SubImage":           method_ptr_NRGBA64_SubImage,
		"(*NYCbCrA).AOffset":            method_ptr_NYCbCrA_AOffset,
		"(*NYCbCrA).At":                 method_ptr_NYCbCrA_At,
		"(*NYCbCrA).ColorModel":         method_ptr_NYCbCrA_ColorModel,
		"(*NYCbCrA).NYCbCrAAt":          method_ptr_NYCbCrA_NYCbCrAAt,
		"(*NYCbCrA).Opaque":             method_ptr_NYCbCrA_Opaque,
		"(*NYCbCrA).RGBA64At":           method_ptr_NYCbCrA_RGBA64At,
		"(*NYCbCrA).SubImage":           method_ptr_NYCbCrA_SubImage,
		"(*Paletted).At":                method_ptr_Paletted_At,
		"(*Paletted).Bounds":            method_ptr_Paletted_Bounds,
		"(*Paletted).ColorIndexAt":      method_ptr_Paletted_ColorIndexAt,
		"(*Paletted).ColorModel":        method_ptr_Paletted_ColorModel,
		"(*Paletted).Opaque":            method_ptr_Paletted_Opaque,
		"(*Paletted).PixOffset":         method_ptr_Paletted_PixOffset,
		"(*Paletted).RGBA64At":          method_ptr_Paletted_RGBA64At,
		"(*Paletted).Set":               method_ptr_Paletted_Set,
		"(*Paletted).SetColorIndex":     method_ptr_Paletted_SetColorIndex,
		"(*Paletted).SetRGBA64":         method_ptr_Paletted_SetRGBA64,
		"(*Paletted).SubImage":          method_ptr_Paletted_SubImage,
		"(*Point).Add":                  method_ptr_Point_Add,
		"(*Point).Div":                  method_ptr_Point_Div,
		"(*Point).Eq":                   method_ptr_Point_Eq,
		"(*Point).In":                   method_ptr_Point_In,
		"(*Point).Mod":                  method_ptr_Point_Mod,
		"(*Point).Mul":                  method_ptr_Point_Mul,
		"(*Point).String":               method_ptr_Point_String,
		"(*Point).Sub":                  method_ptr_Point_Sub,
		"(*RGBA).At":                    method_ptr_RGBA_At,
		"(*RGBA).Bounds":                method_ptr_RGBA_Bounds,
		"(*RGBA).ColorModel":            method_ptr_RGBA_ColorModel,
		"(*RGBA).Opaque":                method_ptr_RGBA_Opaque,
		"(*RGBA).PixOffset":             method_ptr_RGBA_PixOffset,
		"(*RGBA).RGBA64At":              method_ptr_RGBA_RGBA64At,
		"(*RGBA).RGBAAt":                method_ptr_RGBA_RGBAAt,
		"(*RGBA).Set":                   method_ptr_RGBA_Set,
		"(*RGBA).SetRGBA":               method_ptr_RGBA_SetRGBA,
		"(*RGBA).SetRGBA64":             method_ptr_RGBA_SetRGBA64,
		"(*RGBA).SubImage":              method_ptr_RGBA_SubImage,
		"(*RGBA64).At":                  method_ptr_RGBA64_At,
		"(*RGBA64).Bounds":              method_ptr_RGBA64_Bounds,
		"(*RGBA64).ColorModel":          method_ptr_RGBA64_ColorModel,
		"(*RGBA64).Opaque":              method_ptr_RGBA64_Opaque,
		"(*RGBA64).PixOffset":           method_ptr_RGBA64_PixOffset,
		"(*RGBA64).RGBA64At":            method_ptr_RGBA64_RGBA64At,
		"(*RGBA64).Set":                 method_ptr_RGBA64_Set,
		"(*RGBA64).SetRGBA64":           method_ptr_RGBA64_SetRGBA64,
		"(*RGBA64).SubImage":            method_ptr_RGBA64_SubImage,
		"(*Rectangle).Add":              method_ptr_Rectangle_Add,
		"(*Rectangle).At":               method_ptr_Rectangle_At,
		"(*Rectangle).Bounds":           method_ptr_Rectangle_Bounds,
		"(*Rectangle).Canon":            method_ptr_Rectangle_Canon,
		"(*Rectangle).ColorModel":       method_ptr_Rectangle_ColorModel,
		"(*Rectangle).Dx":               method_ptr_Rectangle_Dx,
		"(*Rectangle).Dy":               method_ptr_Rectangle_Dy,
		"(*Rectangle).Empty":            method_ptr_Rectangle_Empty,
		"(*Rectangle).Eq":               method_ptr_Rectangle_Eq,
		"(*Rectangle).In":               method_ptr_Rectangle_In,
		"(*Rectangle).Inset":            method_ptr_Rectangle_Inset,
		"(*Rectangle).Intersect":        method_ptr_Rectangle_Intersect,
		"(*Rectangle).Overlaps":         method_ptr_Rectangle_Overlaps,
		"(*Rectangle).RGBA64At":         method_ptr_Rectangle_RGBA64At,
		"(*Rectangle).Size":             method_ptr_Rectangle_Size,
		"(*Rectangle).String":           method_ptr_Rectangle_String,
		"(*Rectangle).Sub":              method_ptr_Rectangle_Sub,
		"(*Rectangle).Union":            method_ptr_Rectangle_Union,
		"(*Uniform).At":                 method_ptr_Uniform_At,
		"(*Uniform).Bounds":             method_ptr_Uniform_Bounds,
		"(*Uniform).ColorModel":         method_ptr_Uniform_ColorModel,
		"(*Uniform).Convert":            method_ptr_Uniform_Convert,
		"(*Uniform).Opaque":             method_ptr_Uniform_Opaque,
		"(*Uniform).RGBA64At":           method_ptr_Uniform_RGBA64At,
		"(*YCbCr).At":                   method_ptr_YCbCr_At,
		"(*YCbCr).Bounds":               method_ptr_YCbCr_Bounds,
		"(*YCbCr).COffset":              method_ptr_YCbCr_COffset,
		"(*YCbCr).ColorModel":           method_ptr_YCbCr_ColorModel,
		"(*YCbCr).Opaque":               method_ptr_YCbCr_Opaque,
		"(*YCbCr).RGBA64At":             method_ptr_YCbCr_RGBA64At,
		"(*YCbCr).SubImage":             method_ptr_YCbCr_SubImage,
		"(*YCbCr).YCbCrAt":              method_ptr_YCbCr_YCbCrAt,
		"(*YCbCr).YOffset":              method_ptr_YCbCr_YOffset,
		"(*YCbCrSubsampleRatio).String": method_ptr_YCbCrSubsampleRatio_String,
		"(Point).Add":                   method_Point_Add,
		"(Point).Div":                   method_Point_Div,
		"(Point).Eq":                    method_Point_Eq,
		"(Point).In":                    method_Point_In,
		"(Point).Mod":                   method_Point_Mod,
		"(Point).Mul":                   method_Point_Mul,
		"(Point).String":                method_Point_String,
		"(Point).Sub":                   method_Point_Sub,
		"(Rectangle).Add":               method_Rectangle_Add,
		"(Rectangle).At":                method_Rectangle_At,
		"(Rectangle).Bounds":            method_Rectangle_Bounds,
		"(Rectangle).Canon":             method_Rectangle_Canon,
		"(Rectangle).ColorModel":        method_Rectangle_ColorModel,
		"(Rectangle).Dx":                method_Rectangle_Dx,
		"(Rectangle).Dy":                method_Rectangle_Dy,
		"(Rectangle).Empty":             method_Rectangle_Empty,
		"(Rectangle).Eq":                method_Rectangle_Eq,
		"(Rectangle).In":                method_Rectangle_In,
		"(Rectangle).Inset":             method_Rectangle_Inset,
		"(Rectangle).Intersect":         method_Rectangle_Intersect,
		"(Rectangle).Overlaps":          method_Rectangle_Overlaps,
		"(Rectangle).RGBA64At":          method_Rectangle_RGBA64At,
		"(Rectangle).Size":              method_Rectangle_Size,
		"(Rectangle).String":            method_Rectangle_String,
		"(Rectangle).Sub":               method_Rectangle_Sub,
		"(Rectangle).Union":             method_Rectangle_Union,
		"(YCbCrSubsampleRatio).String":  method_YCbCrSubsampleRatio_String,
		"NewAlpha":                      func_NewAlpha,
		"NewAlpha16":                    func_NewAlpha16,
		"NewCMYK":                       func_NewCMYK,
		"NewGray":                       func_NewGray,
		"NewGray16":                     func_NewGray16,
		"NewNRGBA":                      func_NewNRGBA,
		"NewNRGBA64":                    func_NewNRGBA64,
		"NewNYCbCrA":                    func_NewNYCbCrA,
		"NewPaletted":                   func_NewPaletted,
		"NewRGBA":                       func_NewRGBA,
		"NewRGBA64":                     func_NewRGBA64,
		"NewUniform":                    func_NewUniform,
		"NewYCbCr":                      func_NewYCbCr,
		"Pt":                            func_Pt,
		"Rect":                          func_Rect,
		"RegisterFormat":                func_RegisterFormat,
	})
}

func method_ptr_Alpha_AlphaAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha).AlphaAt(ixgo.DirectCallArg[*q.Alpha](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Alpha_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha).At(ixgo.DirectCallArg[*q.Alpha](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Alpha_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha).Bounds(ixgo.DirectCallArg[*q.Alpha](ctx, 0)))
}

func method_ptr_Alpha_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha).ColorModel(ixgo.DirectCallArg[*q.Alpha](ctx, 0)))
}

func method_ptr_Alpha_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha).Opaque(ixgo.DirectCallArg[*q.Alpha](ctx, 0)))
}

func method_ptr_Alpha_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha).PixOffset(ixgo.DirectCallArg[*q.Alpha](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Alpha_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha).RGBA64At(ixgo.DirectCallArg[*q.Alpha](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Alpha_Set(ctx ixgo.DirectCallContext) {
	(*q.Alpha).Set(ixgo.DirectCallArg[*q.Alpha](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_Alpha_SetAlpha(ctx ixgo.DirectCallContext) {
	(*q.Alpha).SetAlpha(ixgo.DirectCallArg[*q.Alpha](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Alpha](ctx, 3))
}

func method_ptr_Alpha_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.Alpha).SetRGBA64(ixgo.DirectCallArg[*q.Alpha](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_Alpha_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha).SubImage(ixgo.DirectCallArg[*q.Alpha](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Alpha16_Alpha16At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha16).Alpha16At(ixgo.DirectCallArg[*q.Alpha16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Alpha16_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha16).At(ixgo.DirectCallArg[*q.Alpha16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Alpha16_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha16).Bounds(ixgo.DirectCallArg[*q.Alpha16](ctx, 0)))
}

func method_ptr_Alpha16_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha16).ColorModel(ixgo.DirectCallArg[*q.Alpha16](ctx, 0)))
}

func method_ptr_Alpha16_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha16).Opaque(ixgo.DirectCallArg[*q.Alpha16](ctx, 0)))
}

func method_ptr_Alpha16_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha16).PixOffset(ixgo.DirectCallArg[*q.Alpha16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Alpha16_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha16).RGBA64At(ixgo.DirectCallArg[*q.Alpha16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Alpha16_Set(ctx ixgo.DirectCallContext) {
	(*q.Alpha16).Set(ixgo.DirectCallArg[*q.Alpha16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_Alpha16_SetAlpha16(ctx ixgo.DirectCallContext) {
	(*q.Alpha16).SetAlpha16(ixgo.DirectCallArg[*q.Alpha16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Alpha16](ctx, 3))
}

func method_ptr_Alpha16_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.Alpha16).SetRGBA64(ixgo.DirectCallArg[*q.Alpha16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_Alpha16_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alpha16).SubImage(ixgo.DirectCallArg[*q.Alpha16](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_CMYK_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CMYK).At(ixgo.DirectCallArg[*q.CMYK](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_CMYK_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CMYK).Bounds(ixgo.DirectCallArg[*q.CMYK](ctx, 0)))
}

func method_ptr_CMYK_CMYKAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CMYK).CMYKAt(ixgo.DirectCallArg[*q.CMYK](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_CMYK_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CMYK).ColorModel(ixgo.DirectCallArg[*q.CMYK](ctx, 0)))
}

func method_ptr_CMYK_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CMYK).Opaque(ixgo.DirectCallArg[*q.CMYK](ctx, 0)))
}

func method_ptr_CMYK_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CMYK).PixOffset(ixgo.DirectCallArg[*q.CMYK](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_CMYK_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CMYK).RGBA64At(ixgo.DirectCallArg[*q.CMYK](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_CMYK_Set(ctx ixgo.DirectCallContext) {
	(*q.CMYK).Set(ixgo.DirectCallArg[*q.CMYK](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_CMYK_SetCMYK(ctx ixgo.DirectCallContext) {
	(*q.CMYK).SetCMYK(ixgo.DirectCallArg[*q.CMYK](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.CMYK](ctx, 3))
}

func method_ptr_CMYK_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.CMYK).SetRGBA64(ixgo.DirectCallArg[*q.CMYK](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_CMYK_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CMYK).SubImage(ixgo.DirectCallArg[*q.CMYK](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Gray_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray).At(ixgo.DirectCallArg[*q.Gray](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Gray_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray).Bounds(ixgo.DirectCallArg[*q.Gray](ctx, 0)))
}

func method_ptr_Gray_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray).ColorModel(ixgo.DirectCallArg[*q.Gray](ctx, 0)))
}

func method_ptr_Gray_GrayAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray).GrayAt(ixgo.DirectCallArg[*q.Gray](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Gray_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray).Opaque(ixgo.DirectCallArg[*q.Gray](ctx, 0)))
}

func method_ptr_Gray_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray).PixOffset(ixgo.DirectCallArg[*q.Gray](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Gray_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray).RGBA64At(ixgo.DirectCallArg[*q.Gray](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Gray_Set(ctx ixgo.DirectCallContext) {
	(*q.Gray).Set(ixgo.DirectCallArg[*q.Gray](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_Gray_SetGray(ctx ixgo.DirectCallContext) {
	(*q.Gray).SetGray(ixgo.DirectCallArg[*q.Gray](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Gray](ctx, 3))
}

func method_ptr_Gray_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.Gray).SetRGBA64(ixgo.DirectCallArg[*q.Gray](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_Gray_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray).SubImage(ixgo.DirectCallArg[*q.Gray](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Gray16_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray16).At(ixgo.DirectCallArg[*q.Gray16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Gray16_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray16).Bounds(ixgo.DirectCallArg[*q.Gray16](ctx, 0)))
}

func method_ptr_Gray16_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray16).ColorModel(ixgo.DirectCallArg[*q.Gray16](ctx, 0)))
}

func method_ptr_Gray16_Gray16At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray16).Gray16At(ixgo.DirectCallArg[*q.Gray16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Gray16_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray16).Opaque(ixgo.DirectCallArg[*q.Gray16](ctx, 0)))
}

func method_ptr_Gray16_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray16).PixOffset(ixgo.DirectCallArg[*q.Gray16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Gray16_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray16).RGBA64At(ixgo.DirectCallArg[*q.Gray16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Gray16_Set(ctx ixgo.DirectCallContext) {
	(*q.Gray16).Set(ixgo.DirectCallArg[*q.Gray16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_Gray16_SetGray16(ctx ixgo.DirectCallContext) {
	(*q.Gray16).SetGray16(ixgo.DirectCallArg[*q.Gray16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Gray16](ctx, 3))
}

func method_ptr_Gray16_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.Gray16).SetRGBA64(ixgo.DirectCallArg[*q.Gray16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_Gray16_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Gray16).SubImage(ixgo.DirectCallArg[*q.Gray16](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_NRGBA_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA).At(ixgo.DirectCallArg[*q.NRGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NRGBA_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA).Bounds(ixgo.DirectCallArg[*q.NRGBA](ctx, 0)))
}

func method_ptr_NRGBA_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA).ColorModel(ixgo.DirectCallArg[*q.NRGBA](ctx, 0)))
}

func method_ptr_NRGBA_NRGBAAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA).NRGBAAt(ixgo.DirectCallArg[*q.NRGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NRGBA_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA).Opaque(ixgo.DirectCallArg[*q.NRGBA](ctx, 0)))
}

func method_ptr_NRGBA_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA).PixOffset(ixgo.DirectCallArg[*q.NRGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NRGBA_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA).RGBA64At(ixgo.DirectCallArg[*q.NRGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NRGBA_Set(ctx ixgo.DirectCallContext) {
	(*q.NRGBA).Set(ixgo.DirectCallArg[*q.NRGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_NRGBA_SetNRGBA(ctx ixgo.DirectCallContext) {
	(*q.NRGBA).SetNRGBA(ixgo.DirectCallArg[*q.NRGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.NRGBA](ctx, 3))
}

func method_ptr_NRGBA_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.NRGBA).SetRGBA64(ixgo.DirectCallArg[*q.NRGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_NRGBA_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA).SubImage(ixgo.DirectCallArg[*q.NRGBA](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_NRGBA64_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA64).At(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NRGBA64_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA64).Bounds(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0)))
}

func method_ptr_NRGBA64_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA64).ColorModel(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0)))
}

func method_ptr_NRGBA64_NRGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA64).NRGBA64At(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NRGBA64_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA64).Opaque(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0)))
}

func method_ptr_NRGBA64_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA64).PixOffset(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NRGBA64_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA64).RGBA64At(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NRGBA64_Set(ctx ixgo.DirectCallContext) {
	(*q.NRGBA64).Set(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_NRGBA64_SetNRGBA64(ctx ixgo.DirectCallContext) {
	(*q.NRGBA64).SetNRGBA64(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.NRGBA64](ctx, 3))
}

func method_ptr_NRGBA64_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.NRGBA64).SetRGBA64(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_NRGBA64_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NRGBA64).SubImage(ixgo.DirectCallArg[*q.NRGBA64](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_NYCbCrA_AOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NYCbCrA).AOffset(ixgo.DirectCallArg[*q.NYCbCrA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NYCbCrA_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NYCbCrA).At(ixgo.DirectCallArg[*q.NYCbCrA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NYCbCrA_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NYCbCrA).ColorModel(ixgo.DirectCallArg[*q.NYCbCrA](ctx, 0)))
}

func method_ptr_NYCbCrA_NYCbCrAAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NYCbCrA).NYCbCrAAt(ixgo.DirectCallArg[*q.NYCbCrA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NYCbCrA_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NYCbCrA).Opaque(ixgo.DirectCallArg[*q.NYCbCrA](ctx, 0)))
}

func method_ptr_NYCbCrA_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NYCbCrA).RGBA64At(ixgo.DirectCallArg[*q.NYCbCrA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_NYCbCrA_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NYCbCrA).SubImage(ixgo.DirectCallArg[*q.NYCbCrA](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func func_NewAlpha(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewAlpha(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func func_NewAlpha16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewAlpha16(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func func_NewCMYK(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewCMYK(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func func_NewGray(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewGray(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func func_NewGray16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewGray16(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func func_NewNRGBA(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewNRGBA(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func func_NewNRGBA64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewNRGBA64(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func func_NewNYCbCrA(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewNYCbCrA(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.YCbCrSubsampleRatio](ctx, 1)))
}

func func_NewPaletted(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewPaletted(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[color.Palette](ctx, 1)))
}

func func_NewRGBA(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewRGBA(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func func_NewRGBA64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewRGBA64(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func func_NewUniform(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewUniform(ixgo.DirectCallArg[color.Color](ctx, 0)))
}

func func_NewYCbCr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewYCbCr(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.YCbCrSubsampleRatio](ctx, 1)))
}

func method_ptr_Paletted_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Paletted).At(ixgo.DirectCallArg[*q.Paletted](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Paletted_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Paletted).Bounds(ixgo.DirectCallArg[*q.Paletted](ctx, 0)))
}

func method_ptr_Paletted_ColorIndexAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Paletted).ColorIndexAt(ixgo.DirectCallArg[*q.Paletted](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Paletted_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Paletted).ColorModel(ixgo.DirectCallArg[*q.Paletted](ctx, 0)))
}

func method_ptr_Paletted_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Paletted).Opaque(ixgo.DirectCallArg[*q.Paletted](ctx, 0)))
}

func method_ptr_Paletted_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Paletted).PixOffset(ixgo.DirectCallArg[*q.Paletted](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Paletted_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Paletted).RGBA64At(ixgo.DirectCallArg[*q.Paletted](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Paletted_Set(ctx ixgo.DirectCallContext) {
	(*q.Paletted).Set(ixgo.DirectCallArg[*q.Paletted](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_Paletted_SetColorIndex(ctx ixgo.DirectCallContext) {
	(*q.Paletted).SetColorIndex(ixgo.DirectCallArg[*q.Paletted](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[uint8](ctx, 3))
}

func method_ptr_Paletted_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.Paletted).SetRGBA64(ixgo.DirectCallArg[*q.Paletted](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_Paletted_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Paletted).SubImage(ixgo.DirectCallArg[*q.Paletted](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_Point_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Point.Add(ixgo.DirectCallArg[q.Point](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func method_ptr_Point_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Point).Add(ixgo.DirectCallArg[*q.Point](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func method_Point_Div(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Point.Div(ixgo.DirectCallArg[q.Point](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Point_Div(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Point).Div(ixgo.DirectCallArg[*q.Point](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_Point_Eq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Point.Eq(ixgo.DirectCallArg[q.Point](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func method_ptr_Point_Eq(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Point).Eq(ixgo.DirectCallArg[*q.Point](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func method_Point_In(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Point.In(ixgo.DirectCallArg[q.Point](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Point_In(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Point).In(ixgo.DirectCallArg[*q.Point](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_Point_Mod(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Point.Mod(ixgo.DirectCallArg[q.Point](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Point_Mod(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Point).Mod(ixgo.DirectCallArg[*q.Point](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_Point_Mul(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Point.Mul(ixgo.DirectCallArg[q.Point](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Point_Mul(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Point).Mul(ixgo.DirectCallArg[*q.Point](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_Point_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Point.String(ixgo.DirectCallArg[q.Point](ctx, 0)))
}

func method_ptr_Point_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Point).String(ixgo.DirectCallArg[*q.Point](ctx, 0)))
}

func method_Point_Sub(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Point.Sub(ixgo.DirectCallArg[q.Point](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func method_ptr_Point_Sub(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Point).Sub(ixgo.DirectCallArg[*q.Point](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func func_Pt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Pt(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_RGBA_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA).At(ixgo.DirectCallArg[*q.RGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_RGBA_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA).Bounds(ixgo.DirectCallArg[*q.RGBA](ctx, 0)))
}

func method_ptr_RGBA_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA).ColorModel(ixgo.DirectCallArg[*q.RGBA](ctx, 0)))
}

func method_ptr_RGBA_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA).Opaque(ixgo.DirectCallArg[*q.RGBA](ctx, 0)))
}

func method_ptr_RGBA_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA).PixOffset(ixgo.DirectCallArg[*q.RGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_RGBA_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA).RGBA64At(ixgo.DirectCallArg[*q.RGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_RGBA_RGBAAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA).RGBAAt(ixgo.DirectCallArg[*q.RGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_RGBA_Set(ctx ixgo.DirectCallContext) {
	(*q.RGBA).Set(ixgo.DirectCallArg[*q.RGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_RGBA_SetRGBA(ctx ixgo.DirectCallContext) {
	(*q.RGBA).SetRGBA(ixgo.DirectCallArg[*q.RGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA](ctx, 3))
}

func method_ptr_RGBA_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.RGBA).SetRGBA64(ixgo.DirectCallArg[*q.RGBA](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_RGBA_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA).SubImage(ixgo.DirectCallArg[*q.RGBA](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_RGBA64_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA64).At(ixgo.DirectCallArg[*q.RGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_RGBA64_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA64).Bounds(ixgo.DirectCallArg[*q.RGBA64](ctx, 0)))
}

func method_ptr_RGBA64_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA64).ColorModel(ixgo.DirectCallArg[*q.RGBA64](ctx, 0)))
}

func method_ptr_RGBA64_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA64).Opaque(ixgo.DirectCallArg[*q.RGBA64](ctx, 0)))
}

func method_ptr_RGBA64_PixOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA64).PixOffset(ixgo.DirectCallArg[*q.RGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_RGBA64_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA64).RGBA64At(ixgo.DirectCallArg[*q.RGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_RGBA64_Set(ctx ixgo.DirectCallContext) {
	(*q.RGBA64).Set(ixgo.DirectCallArg[*q.RGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.Color](ctx, 3))
}

func method_ptr_RGBA64_SetRGBA64(ctx ixgo.DirectCallContext) {
	(*q.RGBA64).SetRGBA64(ixgo.DirectCallArg[*q.RGBA64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[color.RGBA64](ctx, 3))
}

func method_ptr_RGBA64_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RGBA64).SubImage(ixgo.DirectCallArg[*q.RGBA64](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func func_Rect(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rect(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func method_Rectangle_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Add(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func method_ptr_Rectangle_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Add(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func method_Rectangle_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.At(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Rectangle_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).At(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_Rectangle_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Bounds(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func method_ptr_Rectangle_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Bounds(ixgo.DirectCallArg[*q.Rectangle](ctx, 0)))
}

func method_Rectangle_Canon(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Canon(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func method_ptr_Rectangle_Canon(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Canon(ixgo.DirectCallArg[*q.Rectangle](ctx, 0)))
}

func method_Rectangle_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.ColorModel(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func method_ptr_Rectangle_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).ColorModel(ixgo.DirectCallArg[*q.Rectangle](ctx, 0)))
}

func method_Rectangle_Dx(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Dx(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func method_ptr_Rectangle_Dx(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Dx(ixgo.DirectCallArg[*q.Rectangle](ctx, 0)))
}

func method_Rectangle_Dy(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Dy(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func method_ptr_Rectangle_Dy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Dy(ixgo.DirectCallArg[*q.Rectangle](ctx, 0)))
}

func method_Rectangle_Empty(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Empty(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func method_ptr_Rectangle_Empty(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Empty(ixgo.DirectCallArg[*q.Rectangle](ctx, 0)))
}

func method_Rectangle_Eq(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Eq(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Rectangle_Eq(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Eq(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_Rectangle_In(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.In(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Rectangle_In(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).In(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_Rectangle_Inset(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Inset(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Rectangle_Inset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Inset(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_Rectangle_Intersect(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Intersect(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Rectangle_Intersect(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Intersect(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_Rectangle_Overlaps(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Overlaps(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Rectangle_Overlaps(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Overlaps(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_Rectangle_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.RGBA64At(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Rectangle_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).RGBA64At(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_Rectangle_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Size(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func method_ptr_Rectangle_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Size(ixgo.DirectCallArg[*q.Rectangle](ctx, 0)))
}

func method_Rectangle_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.String(ixgo.DirectCallArg[q.Rectangle](ctx, 0)))
}

func method_ptr_Rectangle_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).String(ixgo.DirectCallArg[*q.Rectangle](ctx, 0)))
}

func method_Rectangle_Sub(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Sub(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func method_ptr_Rectangle_Sub(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Sub(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Point](ctx, 1)))
}

func method_Rectangle_Union(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rectangle.Union(ixgo.DirectCallArg[q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_Rectangle_Union(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rectangle).Union(ixgo.DirectCallArg[*q.Rectangle](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func func_RegisterFormat(ctx ixgo.DirectCallContext) {
	q.RegisterFormat(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[func(io.Reader) (q.Image, error)](ctx, 2), ixgo.DirectCallArg[func(io.Reader) (q.Config, error)](ctx, 3))
}

func method_ptr_Uniform_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uniform).At(ixgo.DirectCallArg[*q.Uniform](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Uniform_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uniform).Bounds(ixgo.DirectCallArg[*q.Uniform](ctx, 0)))
}

func method_ptr_Uniform_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uniform).ColorModel(ixgo.DirectCallArg[*q.Uniform](ctx, 0)))
}

func method_ptr_Uniform_Convert(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uniform).Convert(ixgo.DirectCallArg[*q.Uniform](ctx, 0), ixgo.DirectCallArg[color.Color](ctx, 1)))
}

func method_ptr_Uniform_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uniform).Opaque(ixgo.DirectCallArg[*q.Uniform](ctx, 0)))
}

func method_ptr_Uniform_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Uniform).RGBA64At(ixgo.DirectCallArg[*q.Uniform](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_YCbCr_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCr).At(ixgo.DirectCallArg[*q.YCbCr](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_YCbCr_Bounds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCr).Bounds(ixgo.DirectCallArg[*q.YCbCr](ctx, 0)))
}

func method_ptr_YCbCr_COffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCr).COffset(ixgo.DirectCallArg[*q.YCbCr](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_YCbCr_ColorModel(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCr).ColorModel(ixgo.DirectCallArg[*q.YCbCr](ctx, 0)))
}

func method_ptr_YCbCr_Opaque(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCr).Opaque(ixgo.DirectCallArg[*q.YCbCr](ctx, 0)))
}

func method_ptr_YCbCr_RGBA64At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCr).RGBA64At(ixgo.DirectCallArg[*q.YCbCr](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_YCbCr_SubImage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCr).SubImage(ixgo.DirectCallArg[*q.YCbCr](ctx, 0), ixgo.DirectCallArg[q.Rectangle](ctx, 1)))
}

func method_ptr_YCbCr_YCbCrAt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCr).YCbCrAt(ixgo.DirectCallArg[*q.YCbCr](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_YCbCr_YOffset(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCr).YOffset(ixgo.DirectCallArg[*q.YCbCr](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_YCbCrSubsampleRatio_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.YCbCrSubsampleRatio.String(ixgo.DirectCallArg[q.YCbCrSubsampleRatio](ctx, 0)))
}

func method_ptr_YCbCrSubsampleRatio_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.YCbCrSubsampleRatio).String(ixgo.DirectCallArg[*q.YCbCrSubsampleRatio](ctx, 0)))
}
