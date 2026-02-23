// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package zip

import (
	q "archive/zip"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "zip",
		Path: "archive/zip",
		Deps: map[string]string{
			"bufio":            "bufio",
			"compress/flate":   "flate",
			"encoding/binary":  "binary",
			"errors":           "errors",
			"fmt":              "fmt",
			"hash":             "hash",
			"hash/crc32":       "crc32",
			"internal/godebug": "godebug",
			"io":               "io",
			"io/fs":            "fs",
			"os":               "os",
			"path":             "path",
			"path/filepath":    "filepath",
			"slices":           "slices",
			"strings":          "strings",
			"sync":             "sync",
			"time":             "time",
			"unicode/utf8":     "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Compressor":   reflect.TypeOf((*q.Compressor)(nil)).Elem(),
			"Decompressor": reflect.TypeOf((*q.Decompressor)(nil)).Elem(),
			"File":         reflect.TypeOf((*q.File)(nil)).Elem(),
			"FileHeader":   reflect.TypeOf((*q.FileHeader)(nil)).Elem(),
			"ReadCloser":   reflect.TypeOf((*q.ReadCloser)(nil)).Elem(),
			"Reader":       reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"Writer":       reflect.TypeOf((*q.Writer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrAlgorithm":    reflect.ValueOf(&q.ErrAlgorithm),
			"ErrChecksum":     reflect.ValueOf(&q.ErrChecksum),
			"ErrFormat":       reflect.ValueOf(&q.ErrFormat),
			"ErrInsecurePath": reflect.ValueOf(&q.ErrInsecurePath),
		},
		Funcs: map[string]reflect.Value{
			"FileInfoHeader":       reflect.ValueOf(q.FileInfoHeader),
			"NewReader":            reflect.ValueOf(q.NewReader),
			"NewWriter":            reflect.ValueOf(q.NewWriter),
			"OpenReader":           reflect.ValueOf(q.OpenReader),
			"RegisterCompressor":   reflect.ValueOf(q.RegisterCompressor),
			"RegisterDecompressor": reflect.ValueOf(q.RegisterDecompressor),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"Deflate": {reflect.TypeOf(q.Deflate), constant.MakeInt64(int64(q.Deflate))},
			"Store":   {reflect.TypeOf(q.Store), constant.MakeInt64(int64(q.Store))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
