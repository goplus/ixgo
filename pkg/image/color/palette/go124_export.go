// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package palette

import (
	q "image/color/palette"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("image/color/palette", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "palette",
			Path: "image/color/palette",
			Deps: map[string]string{
				"image/color": "color",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"Plan9":   &q.Plan9,
				"WebSafe": &q.WebSafe,
			},
			Funcs:         map[string]interface{}{},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
