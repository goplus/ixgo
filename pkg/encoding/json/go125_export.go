// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package json

import (
	q "encoding/json"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/json", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "json",
			Path: "encoding/json",
			Deps: map[string]string{
				"bytes":           "bytes",
				"cmp":             "cmp",
				"encoding":        "encoding",
				"encoding/base64": "base64",
				"errors":          "errors",
				"fmt":             "fmt",
				"io":              "io",
				"math":            "math",
				"reflect":         "reflect",
				"slices":          "slices",
				"strconv":         "strconv",
				"strings":         "strings",
				"sync":            "sync",
				"unicode":         "unicode",
				"unicode/utf16":   "utf16",
				"unicode/utf8":    "utf8",
				"unsafe":          "unsafe",
			},
			Interfaces: map[string]reflect.Type{
				"Marshaler":   reflect.TypeOf((*q.Marshaler)(nil)).Elem(),
				"Token":       reflect.TypeOf((*q.Token)(nil)).Elem(),
				"Unmarshaler": reflect.TypeOf((*q.Unmarshaler)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Decoder":               reflect.TypeOf((*q.Decoder)(nil)).Elem(),
				"Delim":                 reflect.TypeOf((*q.Delim)(nil)).Elem(),
				"Encoder":               reflect.TypeOf((*q.Encoder)(nil)).Elem(),
				"InvalidUTF8Error":      reflect.TypeOf((*q.InvalidUTF8Error)(nil)).Elem(),
				"InvalidUnmarshalError": reflect.TypeOf((*q.InvalidUnmarshalError)(nil)).Elem(),
				"MarshalerError":        reflect.TypeOf((*q.MarshalerError)(nil)).Elem(),
				"Number":                reflect.TypeOf((*q.Number)(nil)).Elem(),
				"RawMessage":            reflect.TypeOf((*q.RawMessage)(nil)).Elem(),
				"SyntaxError":           reflect.TypeOf((*q.SyntaxError)(nil)).Elem(),
				"UnmarshalFieldError":   reflect.TypeOf((*q.UnmarshalFieldError)(nil)).Elem(),
				"UnmarshalTypeError":    reflect.TypeOf((*q.UnmarshalTypeError)(nil)).Elem(),
				"UnsupportedTypeError":  reflect.TypeOf((*q.UnsupportedTypeError)(nil)).Elem(),
				"UnsupportedValueError": reflect.TypeOf((*q.UnsupportedValueError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Compact":       q.Compact,
				"HTMLEscape":    q.HTMLEscape,
				"Indent":        q.Indent,
				"Marshal":       q.Marshal,
				"MarshalIndent": q.MarshalIndent,
				"NewDecoder":    q.NewDecoder,
				"NewEncoder":    q.NewEncoder,
				"Unmarshal":     q.Unmarshal,
				"Valid":         q.Valid,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
