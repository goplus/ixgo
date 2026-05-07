// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package png

import (
	q "image/png"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("image/png", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "png",
			Path: "image/png",
			Deps: map[string]string{
				"bufio":           "bufio",
				"compress/zlib":   "zlib",
				"encoding/binary": "binary",
				"fmt":             "fmt",
				"hash":            "hash",
				"hash/crc32":      "crc32",
				"image":           "image",
				"image/color":     "color",
				"io":              "io",
				"strconv":         "strconv",
			},
			Interfaces: map[string]reflect.Type{
				"EncoderBufferPool": reflect.TypeOf((*q.EncoderBufferPool)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"CompressionLevel": reflect.TypeOf((*q.CompressionLevel)(nil)).Elem(),
				"Encoder":          reflect.TypeOf((*q.Encoder)(nil)).Elem(),
				"EncoderBuffer":    reflect.TypeOf((*q.EncoderBuffer)(nil)).Elem(),
				"FormatError":      reflect.TypeOf((*q.FormatError)(nil)).Elem(),
				"UnsupportedError": reflect.TypeOf((*q.UnsupportedError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Decode":       q.Decode,
				"DecodeConfig": q.DecodeConfig,
				"Encode":       q.Encode,
			},
			TypedConsts: map[string]interface{}{
				"BestCompression":    q.BestCompression,
				"BestSpeed":          q.BestSpeed,
				"DefaultCompression": q.DefaultCompression,
				"NoCompression":      q.NoCompression,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
