// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package xml

import (
	q "encoding/xml"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/xml", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "xml",
			Path: "encoding/xml",
			Deps: map[string]string{
				"bufio":        "bufio",
				"bytes":        "bytes",
				"encoding":     "encoding",
				"errors":       "errors",
				"fmt":          "fmt",
				"io":           "io",
				"reflect":      "reflect",
				"runtime":      "runtime",
				"strconv":      "strconv",
				"strings":      "strings",
				"sync":         "sync",
				"unicode":      "unicode",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{
				"Marshaler":       reflect.TypeOf((*q.Marshaler)(nil)).Elem(),
				"MarshalerAttr":   reflect.TypeOf((*q.MarshalerAttr)(nil)).Elem(),
				"Token":           reflect.TypeOf((*q.Token)(nil)).Elem(),
				"TokenReader":     reflect.TypeOf((*q.TokenReader)(nil)).Elem(),
				"Unmarshaler":     reflect.TypeOf((*q.Unmarshaler)(nil)).Elem(),
				"UnmarshalerAttr": reflect.TypeOf((*q.UnmarshalerAttr)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Attr":                 reflect.TypeOf((*q.Attr)(nil)).Elem(),
				"CharData":             reflect.TypeOf((*q.CharData)(nil)).Elem(),
				"Comment":              reflect.TypeOf((*q.Comment)(nil)).Elem(),
				"Decoder":              reflect.TypeOf((*q.Decoder)(nil)).Elem(),
				"Directive":            reflect.TypeOf((*q.Directive)(nil)).Elem(),
				"Encoder":              reflect.TypeOf((*q.Encoder)(nil)).Elem(),
				"EndElement":           reflect.TypeOf((*q.EndElement)(nil)).Elem(),
				"Name":                 reflect.TypeOf((*q.Name)(nil)).Elem(),
				"ProcInst":             reflect.TypeOf((*q.ProcInst)(nil)).Elem(),
				"StartElement":         reflect.TypeOf((*q.StartElement)(nil)).Elem(),
				"SyntaxError":          reflect.TypeOf((*q.SyntaxError)(nil)).Elem(),
				"TagPathError":         reflect.TypeOf((*q.TagPathError)(nil)).Elem(),
				"UnmarshalError":       reflect.TypeOf((*q.UnmarshalError)(nil)).Elem(),
				"UnsupportedTypeError": reflect.TypeOf((*q.UnsupportedTypeError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"HTMLAutoClose": &q.HTMLAutoClose,
				"HTMLEntity":    &q.HTMLEntity,
			},
			Funcs: map[string]interface{}{
				"CopyToken":       q.CopyToken,
				"Escape":          q.Escape,
				"EscapeText":      q.EscapeText,
				"Marshal":         q.Marshal,
				"MarshalIndent":   q.MarshalIndent,
				"NewDecoder":      q.NewDecoder,
				"NewEncoder":      q.NewEncoder,
				"NewTokenDecoder": q.NewTokenDecoder,
				"Unmarshal":       q.Unmarshal,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Header": {Typ: "untyped string", Value: constant.MakeString(string(q.Header))},
			},
		}
	})
}
