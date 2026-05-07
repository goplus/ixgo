// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package draw

import (
	q "image/draw"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("image/draw", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "draw",
			Path: "image/draw",
			Deps: map[string]string{
				"image":                    "image",
				"image/color":              "color",
				"image/internal/imageutil": "imageutil",
			},
			Interfaces: map[string]reflect.Type{
				"Drawer":      reflect.TypeOf((*q.Drawer)(nil)).Elem(),
				"Image":       reflect.TypeOf((*q.Image)(nil)).Elem(),
				"Quantizer":   reflect.TypeOf((*q.Quantizer)(nil)).Elem(),
				"RGBA64Image": reflect.TypeOf((*q.RGBA64Image)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Op": reflect.TypeOf((*q.Op)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"FloydSteinberg": &q.FloydSteinberg,
			},
			Funcs: map[string]interface{}{
				"Draw":     q.Draw,
				"DrawMask": q.DrawMask,
			},
			TypedConsts: map[string]interface{}{
				"Over": q.Over,
				"Src":  q.Src,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
