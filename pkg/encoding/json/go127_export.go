// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package json

import (
	q "encoding/json"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("encoding/json", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "json",
			Path: "encoding/json",
			Deps: map[string]string{
				"bytes":                            "bytes",
				"cmp":                              "cmp",
				"encoding":                         "encoding",
				"encoding/json/internal":           "internal",
				"encoding/json/internal/jsonflags": "jsonflags",
				"encoding/json/internal/jsonopts":  "jsonopts",
				"encoding/json/internal/jsonwire":  "jsonwire",
				"encoding/json/jsontext":           "jsontext",
				"encoding/json/v2":                 "json",
				"errors":                           "errors",
				"fmt":                              "fmt",
				"io":                               "io",
				"reflect":                          "reflect",
				"strconv":                          "strconv",
				"strings":                          "strings",
			},
			Interfaces: map[string]reflect.Type{
				"Token": reflect.TypeOf((*q.Token)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Decoder":               reflect.TypeOf((*q.Decoder)(nil)).Elem(),
				"Delim":                 reflect.TypeOf((*q.Delim)(nil)).Elem(),
				"Encoder":               reflect.TypeOf((*q.Encoder)(nil)).Elem(),
				"InvalidUTF8Error":      reflect.TypeOf((*q.InvalidUTF8Error)(nil)).Elem(),
				"InvalidUnmarshalError": reflect.TypeOf((*q.InvalidUnmarshalError)(nil)).Elem(),
				"MarshalerError":        reflect.TypeOf((*q.MarshalerError)(nil)).Elem(),
				"Number":                reflect.TypeOf((*q.Number)(nil)).Elem(),
				"SyntaxError":           reflect.TypeOf((*q.SyntaxError)(nil)).Elem(),
				"UnmarshalFieldError":   reflect.TypeOf((*q.UnmarshalFieldError)(nil)).Elem(),
				"UnmarshalTypeError":    reflect.TypeOf((*q.UnmarshalTypeError)(nil)).Elem(),
				"UnsupportedTypeError":  reflect.TypeOf((*q.UnsupportedTypeError)(nil)).Elem(),
				"UnsupportedValueError": reflect.TypeOf((*q.UnsupportedValueError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"Marshaler":   reflect.TypeOf((*q.Marshaler)(nil)).Elem(),
				"Options":     reflect.TypeOf((*q.Options)(nil)).Elem(),
				"RawMessage":  reflect.TypeOf((*q.RawMessage)(nil)).Elem(),
				"Unmarshaler": reflect.TypeOf((*q.Unmarshaler)(nil)).Elem(),
			},
			Vars: map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"CallMethodsWithLegacySemantics":  reflect.ValueOf(q.CallMethodsWithLegacySemantics),
				"Compact":                         reflect.ValueOf(q.Compact),
				"DefaultOptionsV1":                reflect.ValueOf(q.DefaultOptionsV1),
				"FormatByteArrayAsArray":          reflect.ValueOf(q.FormatByteArrayAsArray),
				"FormatBytesWithLegacySemantics":  reflect.ValueOf(q.FormatBytesWithLegacySemantics),
				"FormatDurationAsNano":            reflect.ValueOf(q.FormatDurationAsNano),
				"HTMLEscape":                      reflect.ValueOf(q.HTMLEscape),
				"Indent":                          reflect.ValueOf(q.Indent),
				"Marshal":                         reflect.ValueOf(q.Marshal),
				"MarshalIndent":                   reflect.ValueOf(q.MarshalIndent),
				"MatchCaseSensitiveDelimiter":     reflect.ValueOf(q.MatchCaseSensitiveDelimiter),
				"MergeWithLegacySemantics":        reflect.ValueOf(q.MergeWithLegacySemantics),
				"NewDecoder":                      reflect.ValueOf(q.NewDecoder),
				"NewEncoder":                      reflect.ValueOf(q.NewEncoder),
				"OmitEmptyWithLegacySemantics":    reflect.ValueOf(q.OmitEmptyWithLegacySemantics),
				"ParseBytesWithLooseRFC4648":      reflect.ValueOf(q.ParseBytesWithLooseRFC4648),
				"ParseTimeWithLooseRFC3339":       reflect.ValueOf(q.ParseTimeWithLooseRFC3339),
				"ReportErrorsWithLegacySemantics": reflect.ValueOf(q.ReportErrorsWithLegacySemantics),
				"StringifyWithLegacySemantics":    reflect.ValueOf(q.StringifyWithLegacySemantics),
				"Unmarshal":                       reflect.ValueOf(q.Unmarshal),
				"UnmarshalArrayFromAnyLength":     reflect.ValueOf(q.UnmarshalArrayFromAnyLength),
				"Valid":                           reflect.ValueOf(q.Valid),
			},
			TypedConsts:   map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
