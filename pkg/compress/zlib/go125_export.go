// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package zlib

import (
	q "compress/zlib"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("compress/zlib", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "zlib",
			Path: "compress/zlib",
			Deps: map[string]string{
				"bufio":           "bufio",
				"compress/flate":  "flate",
				"encoding/binary": "binary",
				"errors":          "errors",
				"fmt":             "fmt",
				"hash":            "hash",
				"hash/adler32":    "adler32",
				"io":              "io",
			},
			Interfaces: map[string]reflect.Type{
				"Resetter": reflect.TypeOf((*q.Resetter)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Writer": reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrChecksum":   &q.ErrChecksum,
				"ErrDictionary": &q.ErrDictionary,
				"ErrHeader":     &q.ErrHeader,
			},
			Funcs: map[string]interface{}{
				"NewReader":          q.NewReader,
				"NewReaderDict":      q.NewReaderDict,
				"NewWriter":          q.NewWriter,
				"NewWriterLevel":     q.NewWriterLevel,
				"NewWriterLevelDict": q.NewWriterLevelDict,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"BestCompression":    {Typ: "untyped int", Value: constant.MakeInt64(int64(q.BestCompression))},
				"BestSpeed":          {Typ: "untyped int", Value: constant.MakeInt64(int64(q.BestSpeed))},
				"DefaultCompression": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.DefaultCompression))},
				"HuffmanOnly":        {Typ: "untyped int", Value: constant.MakeInt64(int64(q.HuffmanOnly))},
				"NoCompression":      {Typ: "untyped int", Value: constant.MakeInt64(int64(q.NoCompression))},
			},
		}
	})
}
