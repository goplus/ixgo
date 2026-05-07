// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package zip

import (
	q "archive/zip"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("archive/zip", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars: map[string]interface{}{
				"ErrAlgorithm":    &q.ErrAlgorithm,
				"ErrChecksum":     &q.ErrChecksum,
				"ErrFormat":       &q.ErrFormat,
				"ErrInsecurePath": &q.ErrInsecurePath,
			},
			Funcs: map[string]interface{}{
				"FileInfoHeader":       q.FileInfoHeader,
				"NewReader":            q.NewReader,
				"NewWriter":            q.NewWriter,
				"OpenReader":           q.OpenReader,
				"RegisterCompressor":   q.RegisterCompressor,
				"RegisterDecompressor": q.RegisterDecompressor,
			},
			TypedConsts: map[string]interface{}{
				"Deflate": q.Deflate,
				"Store":   q.Store,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
