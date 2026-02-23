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
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Compact":       reflect.ValueOf(q.Compact),
			"HTMLEscape":    reflect.ValueOf(q.HTMLEscape),
			"Indent":        reflect.ValueOf(q.Indent),
			"Marshal":       reflect.ValueOf(q.Marshal),
			"MarshalIndent": reflect.ValueOf(q.MarshalIndent),
			"NewDecoder":    reflect.ValueOf(q.NewDecoder),
			"NewEncoder":    reflect.ValueOf(q.NewEncoder),
			"Unmarshal":     reflect.ValueOf(q.Unmarshal),
			"Valid":         reflect.ValueOf(q.Valid),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
