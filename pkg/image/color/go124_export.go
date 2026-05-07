// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package color

import (
	q "image/color"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("image/color", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "color",
			Path: "image/color",
			Deps: map[string]string{},
			Interfaces: map[string]reflect.Type{
				"Color": reflect.TypeOf((*q.Color)(nil)).Elem(),
				"Model": reflect.TypeOf((*q.Model)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Alpha":   reflect.TypeOf((*q.Alpha)(nil)).Elem(),
				"Alpha16": reflect.TypeOf((*q.Alpha16)(nil)).Elem(),
				"CMYK":    reflect.TypeOf((*q.CMYK)(nil)).Elem(),
				"Gray":    reflect.TypeOf((*q.Gray)(nil)).Elem(),
				"Gray16":  reflect.TypeOf((*q.Gray16)(nil)).Elem(),
				"NRGBA":   reflect.TypeOf((*q.NRGBA)(nil)).Elem(),
				"NRGBA64": reflect.TypeOf((*q.NRGBA64)(nil)).Elem(),
				"NYCbCrA": reflect.TypeOf((*q.NYCbCrA)(nil)).Elem(),
				"Palette": reflect.TypeOf((*q.Palette)(nil)).Elem(),
				"RGBA":    reflect.TypeOf((*q.RGBA)(nil)).Elem(),
				"RGBA64":  reflect.TypeOf((*q.RGBA64)(nil)).Elem(),
				"YCbCr":   reflect.TypeOf((*q.YCbCr)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"Alpha16Model": &q.Alpha16Model,
				"AlphaModel":   &q.AlphaModel,
				"Black":        &q.Black,
				"CMYKModel":    &q.CMYKModel,
				"Gray16Model":  &q.Gray16Model,
				"GrayModel":    &q.GrayModel,
				"NRGBA64Model": &q.NRGBA64Model,
				"NRGBAModel":   &q.NRGBAModel,
				"NYCbCrAModel": &q.NYCbCrAModel,
				"Opaque":       &q.Opaque,
				"RGBA64Model":  &q.RGBA64Model,
				"RGBAModel":    &q.RGBAModel,
				"Transparent":  &q.Transparent,
				"White":        &q.White,
				"YCbCrModel":   &q.YCbCrModel,
			},
			Funcs: map[string]interface{}{
				"CMYKToRGB":  q.CMYKToRGB,
				"ModelFunc":  q.ModelFunc,
				"RGBToCMYK":  q.RGBToCMYK,
				"RGBToYCbCr": q.RGBToYCbCr,
				"YCbCrToRGB": q.YCbCrToRGB,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
