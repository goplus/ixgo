// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package multipart

import (
	q "mime/multipart"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("mime/multipart", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "multipart",
			Path: "mime/multipart",
			Deps: map[string]string{
				"bufio":                "bufio",
				"bytes":                "bytes",
				"crypto/rand":          "rand",
				"errors":               "errors",
				"fmt":                  "fmt",
				"internal/godebug":     "godebug",
				"io":                   "io",
				"maps":                 "maps",
				"math":                 "math",
				"mime":                 "mime",
				"mime/quotedprintable": "quotedprintable",
				"net/textproto":        "textproto",
				"os":                   "os",
				"path/filepath":        "filepath",
				"slices":               "slices",
				"strconv":              "strconv",
				"strings":              "strings",
				"unsafe":               "unsafe",
			},
			Interfaces: map[string]reflect.Type{
				"File": reflect.TypeOf((*q.File)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"FileHeader": reflect.TypeOf((*q.FileHeader)(nil)).Elem(),
				"Form":       reflect.TypeOf((*q.Form)(nil)).Elem(),
				"Part":       reflect.TypeOf((*q.Part)(nil)).Elem(),
				"Reader":     reflect.TypeOf((*q.Reader)(nil)).Elem(),
				"Writer":     reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrMessageTooLarge": &q.ErrMessageTooLarge,
			},
			Funcs: map[string]interface{}{
				"FileContentDisposition": q.FileContentDisposition,
				"NewReader":              q.NewReader,
				"NewWriter":              q.NewWriter,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
