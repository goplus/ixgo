// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package color

import (
	q "image/color"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"Alpha16Model": reflect.ValueOf(&q.Alpha16Model),
			"AlphaModel":   reflect.ValueOf(&q.AlphaModel),
			"Black":        reflect.ValueOf(&q.Black),
			"CMYKModel":    reflect.ValueOf(&q.CMYKModel),
			"Gray16Model":  reflect.ValueOf(&q.Gray16Model),
			"GrayModel":    reflect.ValueOf(&q.GrayModel),
			"NRGBA64Model": reflect.ValueOf(&q.NRGBA64Model),
			"NRGBAModel":   reflect.ValueOf(&q.NRGBAModel),
			"NYCbCrAModel": reflect.ValueOf(&q.NYCbCrAModel),
			"Opaque":       reflect.ValueOf(&q.Opaque),
			"RGBA64Model":  reflect.ValueOf(&q.RGBA64Model),
			"RGBAModel":    reflect.ValueOf(&q.RGBAModel),
			"Transparent":  reflect.ValueOf(&q.Transparent),
			"White":        reflect.ValueOf(&q.White),
			"YCbCrModel":   reflect.ValueOf(&q.YCbCrModel),
		},
		Funcs: map[string]reflect.Value{
			"CMYKToRGB":  reflect.ValueOf(q.CMYKToRGB),
			"ModelFunc":  reflect.ValueOf(q.ModelFunc),
			"RGBToCMYK":  reflect.ValueOf(q.RGBToCMYK),
			"RGBToYCbCr": reflect.ValueOf(q.RGBToYCbCr),
			"YCbCrToRGB": reflect.ValueOf(q.YCbCrToRGB),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
