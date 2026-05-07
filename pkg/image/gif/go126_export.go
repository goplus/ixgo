// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package gif

import (
	q "image/gif"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("image/gif", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "gif",
			Path: "image/gif",
			Deps: map[string]string{
				"bufio":               "bufio",
				"bytes":               "bytes",
				"compress/lzw":        "lzw",
				"errors":              "errors",
				"fmt":                 "fmt",
				"image":               "image",
				"image/color":         "color",
				"image/color/palette": "palette",
				"image/draw":          "draw",
				"internal/byteorder":  "byteorder",
				"io":                  "io",
				"math/bits":           "bits",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"GIF":     reflect.TypeOf((*q.GIF)(nil)).Elem(),
				"Options": reflect.TypeOf((*q.Options)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Decode":       q.Decode,
				"DecodeAll":    q.DecodeAll,
				"DecodeConfig": q.DecodeConfig,
				"Encode":       q.Encode,
				"EncodeAll":    q.EncodeAll,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"DisposalBackground": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DisposalBackground))},
				"DisposalNone":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DisposalNone))},
				"DisposalPrevious":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DisposalPrevious))},
			},
		}
	})
}
