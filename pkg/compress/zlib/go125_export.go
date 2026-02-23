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
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"ErrChecksum":   reflect.ValueOf(&q.ErrChecksum),
			"ErrDictionary": reflect.ValueOf(&q.ErrDictionary),
			"ErrHeader":     reflect.ValueOf(&q.ErrHeader),
		},
		Funcs: map[string]reflect.Value{
			"NewReader":          reflect.ValueOf(q.NewReader),
			"NewReaderDict":      reflect.ValueOf(q.NewReaderDict),
			"NewWriter":          reflect.ValueOf(q.NewWriter),
			"NewWriterLevel":     reflect.ValueOf(q.NewWriterLevel),
			"NewWriterLevelDict": reflect.ValueOf(q.NewWriterLevelDict),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"BestCompression":    {"untyped int", constant.MakeInt64(int64(q.BestCompression))},
			"BestSpeed":          {"untyped int", constant.MakeInt64(int64(q.BestSpeed))},
			"DefaultCompression": {"untyped int", constant.MakeInt64(int64(q.DefaultCompression))},
			"HuffmanOnly":        {"untyped int", constant.MakeInt64(int64(q.HuffmanOnly))},
			"NoCompression":      {"untyped int", constant.MakeInt64(int64(q.NoCompression))},
		},
	})
}
