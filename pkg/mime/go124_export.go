// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package mime

import (
	q "mime"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("mime", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "mime",
			Path: "mime",
			Deps: map[string]string{
				"bufio":           "bufio",
				"bytes":           "bytes",
				"encoding/base64": "base64",
				"errors":          "errors",
				"fmt":             "fmt",
				"io":              "io",
				"maps":            "maps",
				"os":              "os",
				"slices":          "slices",
				"strings":         "strings",
				"sync":            "sync",
				"unicode":         "unicode",
				"unicode/utf8":    "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"WordDecoder": reflect.TypeOf((*q.WordDecoder)(nil)).Elem(),
				"WordEncoder": reflect.TypeOf((*q.WordEncoder)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrInvalidMediaParameter": &q.ErrInvalidMediaParameter,
			},
			Funcs: map[string]interface{}{
				"AddExtensionType": q.AddExtensionType,
				"ExtensionsByType": q.ExtensionsByType,
				"FormatMediaType":  q.FormatMediaType,
				"ParseMediaType":   q.ParseMediaType,
				"TypeByExtension":  q.TypeByExtension,
			},
			TypedConsts: map[string]interface{}{
				"BEncoding": q.BEncoding,
				"QEncoding": q.QEncoding,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
