// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package image

import (
	q "image"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("image", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "image",
			Path: "image",
			Deps: map[string]string{
				"bufio":       "bufio",
				"errors":      "errors",
				"image/color": "color",
				"io":          "io",
				"math/bits":   "bits",
				"strconv":     "strconv",
				"sync":        "sync",
				"sync/atomic": "atomic",
			},
			Interfaces: map[string]reflect.Type{
				"Image":         reflect.TypeOf((*q.Image)(nil)).Elem(),
				"PalettedImage": reflect.TypeOf((*q.PalettedImage)(nil)).Elem(),
				"RGBA64Image":   reflect.TypeOf((*q.RGBA64Image)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Alpha":               reflect.TypeOf((*q.Alpha)(nil)).Elem(),
				"Alpha16":             reflect.TypeOf((*q.Alpha16)(nil)).Elem(),
				"CMYK":                reflect.TypeOf((*q.CMYK)(nil)).Elem(),
				"Config":              reflect.TypeOf((*q.Config)(nil)).Elem(),
				"Gray":                reflect.TypeOf((*q.Gray)(nil)).Elem(),
				"Gray16":              reflect.TypeOf((*q.Gray16)(nil)).Elem(),
				"NRGBA":               reflect.TypeOf((*q.NRGBA)(nil)).Elem(),
				"NRGBA64":             reflect.TypeOf((*q.NRGBA64)(nil)).Elem(),
				"NYCbCrA":             reflect.TypeOf((*q.NYCbCrA)(nil)).Elem(),
				"Paletted":            reflect.TypeOf((*q.Paletted)(nil)).Elem(),
				"Point":               reflect.TypeOf((*q.Point)(nil)).Elem(),
				"RGBA":                reflect.TypeOf((*q.RGBA)(nil)).Elem(),
				"RGBA64":              reflect.TypeOf((*q.RGBA64)(nil)).Elem(),
				"Rectangle":           reflect.TypeOf((*q.Rectangle)(nil)).Elem(),
				"Uniform":             reflect.TypeOf((*q.Uniform)(nil)).Elem(),
				"YCbCr":               reflect.TypeOf((*q.YCbCr)(nil)).Elem(),
				"YCbCrSubsampleRatio": reflect.TypeOf((*q.YCbCrSubsampleRatio)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"Black":       &q.Black,
				"ErrFormat":   &q.ErrFormat,
				"Opaque":      &q.Opaque,
				"Transparent": &q.Transparent,
				"White":       &q.White,
				"ZP":          &q.ZP,
				"ZR":          &q.ZR,
			},
			Funcs: map[string]interface{}{
				"Decode":         q.Decode,
				"DecodeConfig":   q.DecodeConfig,
				"NewAlpha":       q.NewAlpha,
				"NewAlpha16":     q.NewAlpha16,
				"NewCMYK":        q.NewCMYK,
				"NewGray":        q.NewGray,
				"NewGray16":      q.NewGray16,
				"NewNRGBA":       q.NewNRGBA,
				"NewNRGBA64":     q.NewNRGBA64,
				"NewNYCbCrA":     q.NewNYCbCrA,
				"NewPaletted":    q.NewPaletted,
				"NewRGBA":        q.NewRGBA,
				"NewRGBA64":      q.NewRGBA64,
				"NewUniform":     q.NewUniform,
				"NewYCbCr":       q.NewYCbCr,
				"Pt":             q.Pt,
				"Rect":           q.Rect,
				"RegisterFormat": q.RegisterFormat,
			},
			TypedConsts: map[string]interface{}{
				"YCbCrSubsampleRatio410": q.YCbCrSubsampleRatio410,
				"YCbCrSubsampleRatio411": q.YCbCrSubsampleRatio411,
				"YCbCrSubsampleRatio420": q.YCbCrSubsampleRatio420,
				"YCbCrSubsampleRatio422": q.YCbCrSubsampleRatio422,
				"YCbCrSubsampleRatio440": q.YCbCrSubsampleRatio440,
				"YCbCrSubsampleRatio444": q.YCbCrSubsampleRatio444,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
