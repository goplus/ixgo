// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package png

import (
	q "image/png"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":       reflect.ValueOf(q.Decode),
			"DecodeConfig": reflect.ValueOf(q.DecodeConfig),
			"Encode":       reflect.ValueOf(q.Encode),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"BestCompression":    {reflect.TypeOf(q.BestCompression), constant.MakeInt64(int64(q.BestCompression))},
			"BestSpeed":          {reflect.TypeOf(q.BestSpeed), constant.MakeInt64(int64(q.BestSpeed))},
			"DefaultCompression": {reflect.TypeOf(q.DefaultCompression), constant.MakeInt64(int64(q.DefaultCompression))},
			"NoCompression":      {reflect.TypeOf(q.NoCompression), constant.MakeInt64(int64(q.NoCompression))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
