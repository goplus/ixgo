// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package draw

import (
	q "image/draw"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("image/draw", func() *ixgo.Package {
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
			Vars: map[string]reflect.Value{
				"FloydSteinberg": reflect.ValueOf(&q.FloydSteinberg),
			},
			Funcs: map[string]reflect.Value{
				"Draw":     reflect.ValueOf(q.Draw),
				"DrawMask": reflect.ValueOf(q.DrawMask),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"Over": {Typ: reflect.TypeOf(q.Over), Value: constant.MakeInt64(int64(q.Over))},
				"Src":  {Typ: reflect.TypeOf(q.Src), Value: constant.MakeInt64(int64(q.Src))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
