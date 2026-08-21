// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package json

import (
	q "encoding/json/v2"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("encoding/json/v2", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "json",
			Path: "encoding/json/v2",
			Deps: map[string]string{
				"bytes":                            "bytes",
				"cmp":                              "cmp",
				"encoding":                         "encoding",
				"encoding/base32":                  "base32",
				"encoding/base64":                  "base64",
				"encoding/binary":                  "binary",
				"encoding/hex":                     "hex",
				"encoding/json/internal":           "internal",
				"encoding/json/internal/jsonflags": "jsonflags",
				"encoding/json/internal/jsonopts":  "jsonopts",
				"encoding/json/internal/jsonwire":  "jsonwire",
				"encoding/json/jsontext":           "jsontext",
				"errors":                           "errors",
				"fmt":                              "fmt",
				"io":                               "io",
				"math":                             "math",
				"math/bits":                        "bits",
				"reflect":                          "reflect",
				"slices":                           "slices",
				"strconv":                          "strconv",
				"strings":                          "strings",
				"sync":                             "sync",
				"time":                             "time",
				"unicode":                          "unicode",
				"unicode/utf8":                     "utf8",
			},
			Interfaces: map[string]reflect.Type{
				"Marshaler":       reflect.TypeOf((*q.Marshaler)(nil)).Elem(),
				"MarshalerTo":     reflect.TypeOf((*q.MarshalerTo)(nil)).Elem(),
				"Unmarshaler":     reflect.TypeOf((*q.Unmarshaler)(nil)).Elem(),
				"UnmarshalerFrom": reflect.TypeOf((*q.UnmarshalerFrom)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"SemanticError": reflect.TypeOf((*q.SemanticError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"Marshalers":   reflect.TypeOf((*q.Marshalers)(nil)).Elem(),
				"Options":      reflect.TypeOf((*q.Options)(nil)).Elem(),
				"Unmarshalers": reflect.TypeOf((*q.Unmarshalers)(nil)).Elem(),
			},
			Vars: map[string]reflect.Value{
				"ErrUnknownName": reflect.ValueOf(&q.ErrUnknownName),
			},
			Funcs: map[string]reflect.Value{
				"DefaultOptionsV2":          reflect.ValueOf(q.DefaultOptionsV2),
				"Deterministic":             reflect.ValueOf(q.Deterministic),
				"FormatNilMapAsNull":        reflect.ValueOf(q.FormatNilMapAsNull),
				"FormatNilSliceAsNull":      reflect.ValueOf(q.FormatNilSliceAsNull),
				"JoinMarshalers":            reflect.ValueOf(q.JoinMarshalers),
				"JoinOptions":               reflect.ValueOf(q.JoinOptions),
				"JoinUnmarshalers":          reflect.ValueOf(q.JoinUnmarshalers),
				"Marshal":                   reflect.ValueOf(q.Marshal),
				"MarshalEncode":             reflect.ValueOf(q.MarshalEncode),
				"MarshalWrite":              reflect.ValueOf(q.MarshalWrite),
				"MatchCaseInsensitiveNames": reflect.ValueOf(q.MatchCaseInsensitiveNames),
				"OmitZeroStructFields":      reflect.ValueOf(q.OmitZeroStructFields),
				"RejectUnknownMembers":      reflect.ValueOf(q.RejectUnknownMembers),
				"StringifyNumbers":          reflect.ValueOf(q.StringifyNumbers),
				"Unmarshal":                 reflect.ValueOf(q.Unmarshal),
				"UnmarshalDecode":           reflect.ValueOf(q.UnmarshalDecode),
				"UnmarshalRead":             reflect.ValueOf(q.UnmarshalRead),
				"WithMarshalers":            reflect.ValueOf(q.WithMarshalers),
				"WithUnmarshalers":          reflect.ValueOf(q.WithUnmarshalers),
			},
			TypedConsts:   map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
