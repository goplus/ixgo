// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package mime

import (
	q "mime"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"ErrInvalidMediaParameter": reflect.ValueOf(&q.ErrInvalidMediaParameter),
		},
		Funcs: map[string]reflect.Value{
			"AddExtensionType": reflect.ValueOf(q.AddExtensionType),
			"ExtensionsByType": reflect.ValueOf(q.ExtensionsByType),
			"FormatMediaType":  reflect.ValueOf(q.FormatMediaType),
			"ParseMediaType":   reflect.ValueOf(q.ParseMediaType),
			"TypeByExtension":  reflect.ValueOf(q.TypeByExtension),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"BEncoding": {reflect.TypeOf(q.BEncoding), constant.MakeInt64(int64(q.BEncoding))},
			"QEncoding": {reflect.TypeOf(q.QEncoding), constant.MakeInt64(int64(q.QEncoding))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
