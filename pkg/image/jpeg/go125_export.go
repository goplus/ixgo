// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package jpeg

import (
	q "image/jpeg"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "jpeg",
		Path: "image/jpeg",
		Deps: map[string]string{
			"bufio":                    "bufio",
			"errors":                   "errors",
			"image":                    "image",
			"image/color":              "color",
			"image/internal/imageutil": "imageutil",
			"io":                       "io",
		},
		Interfaces: map[string]reflect.Type{
			"Reader": reflect.TypeOf((*q.Reader)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"FormatError":      reflect.TypeOf((*q.FormatError)(nil)).Elem(),
			"Options":          reflect.TypeOf((*q.Options)(nil)).Elem(),
			"UnsupportedError": reflect.TypeOf((*q.UnsupportedError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":       reflect.ValueOf(q.Decode),
			"DecodeConfig": reflect.ValueOf(q.DecodeConfig),
			"Encode":       reflect.ValueOf(q.Encode),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"DefaultQuality": {"untyped int", constant.MakeInt64(int64(q.DefaultQuality))},
		},
	})
}
