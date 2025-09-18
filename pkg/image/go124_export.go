// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package image

import (
	q "image"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"Black":       reflect.ValueOf(&q.Black),
			"ErrFormat":   reflect.ValueOf(&q.ErrFormat),
			"Opaque":      reflect.ValueOf(&q.Opaque),
			"Transparent": reflect.ValueOf(&q.Transparent),
			"White":       reflect.ValueOf(&q.White),
			"ZP":          reflect.ValueOf(&q.ZP),
			"ZR":          reflect.ValueOf(&q.ZR),
		},
		Funcs: map[string]reflect.Value{
			"Decode":         reflect.ValueOf(q.Decode),
			"DecodeConfig":   reflect.ValueOf(q.DecodeConfig),
			"NewAlpha":       reflect.ValueOf(q.NewAlpha),
			"NewAlpha16":     reflect.ValueOf(q.NewAlpha16),
			"NewCMYK":        reflect.ValueOf(q.NewCMYK),
			"NewGray":        reflect.ValueOf(q.NewGray),
			"NewGray16":      reflect.ValueOf(q.NewGray16),
			"NewNRGBA":       reflect.ValueOf(q.NewNRGBA),
			"NewNRGBA64":     reflect.ValueOf(q.NewNRGBA64),
			"NewNYCbCrA":     reflect.ValueOf(q.NewNYCbCrA),
			"NewPaletted":    reflect.ValueOf(q.NewPaletted),
			"NewRGBA":        reflect.ValueOf(q.NewRGBA),
			"NewRGBA64":      reflect.ValueOf(q.NewRGBA64),
			"NewUniform":     reflect.ValueOf(q.NewUniform),
			"NewYCbCr":       reflect.ValueOf(q.NewYCbCr),
			"Pt":             reflect.ValueOf(q.Pt),
			"Rect":           reflect.ValueOf(q.Rect),
			"RegisterFormat": reflect.ValueOf(q.RegisterFormat),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"YCbCrSubsampleRatio410": {reflect.TypeOf(q.YCbCrSubsampleRatio410), constant.MakeInt64(int64(q.YCbCrSubsampleRatio410))},
			"YCbCrSubsampleRatio411": {reflect.TypeOf(q.YCbCrSubsampleRatio411), constant.MakeInt64(int64(q.YCbCrSubsampleRatio411))},
			"YCbCrSubsampleRatio420": {reflect.TypeOf(q.YCbCrSubsampleRatio420), constant.MakeInt64(int64(q.YCbCrSubsampleRatio420))},
			"YCbCrSubsampleRatio422": {reflect.TypeOf(q.YCbCrSubsampleRatio422), constant.MakeInt64(int64(q.YCbCrSubsampleRatio422))},
			"YCbCrSubsampleRatio440": {reflect.TypeOf(q.YCbCrSubsampleRatio440), constant.MakeInt64(int64(q.YCbCrSubsampleRatio440))},
			"YCbCrSubsampleRatio444": {reflect.TypeOf(q.YCbCrSubsampleRatio444), constant.MakeInt64(int64(q.YCbCrSubsampleRatio444))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
