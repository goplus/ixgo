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
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"HTMLAutoClose": reflect.ValueOf(&q.HTMLAutoClose),
			"HTMLEntity":    reflect.ValueOf(&q.HTMLEntity),
		},
		Funcs: map[string]reflect.Value{
			"CopyToken":       reflect.ValueOf(q.CopyToken),
			"Escape":          reflect.ValueOf(q.Escape),
			"EscapeText":      reflect.ValueOf(q.EscapeText),
			"Marshal":         reflect.ValueOf(q.Marshal),
			"MarshalIndent":   reflect.ValueOf(q.MarshalIndent),
			"NewDecoder":      reflect.ValueOf(q.NewDecoder),
			"NewEncoder":      reflect.ValueOf(q.NewEncoder),
			"NewTokenDecoder": reflect.ValueOf(q.NewTokenDecoder),
			"Unmarshal":       reflect.ValueOf(q.Unmarshal),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"Header": {"untyped string", constant.MakeString(string(q.Header))},
		},
	})
}
