// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package gzip

import (
	q "compress/gzip"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("compress/gzip", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "gzip",
			Path: "compress/gzip",
			Deps: map[string]string{
				"bufio":           "bufio",
				"compress/flate":  "flate",
				"encoding/binary": "binary",
				"errors":          "errors",
				"fmt":             "fmt",
				"hash/crc32":      "crc32",
				"io":              "io",
				"time":            "time",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Header": reflect.TypeOf((*q.Header)(nil)).Elem(),
				"Reader": reflect.TypeOf((*q.Reader)(nil)).Elem(),
				"Writer": reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrChecksum": &q.ErrChecksum,
				"ErrHeader":   &q.ErrHeader,
			},
			Funcs: map[string]interface{}{
				"NewReader":      q.NewReader,
				"NewWriter":      q.NewWriter,
				"NewWriterLevel": q.NewWriterLevel,
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
