// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package jsontext

import (
	q "encoding/json/jsontext"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("encoding/json/jsontext", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "jsontext",
			Path: "encoding/json/jsontext",
			Deps: map[string]string{
				"bytes":                            "bytes",
				"encoding/json/internal":           "internal",
				"encoding/json/internal/jsonflags": "jsonflags",
				"encoding/json/internal/jsonopts":  "jsonopts",
				"encoding/json/internal/jsonwire":  "jsonwire",
				"errors":                           "errors",
				"io":                               "io",
				"iter":                             "iter",
				"math":                             "math",
				"math/bits":                        "bits",
				"slices":                           "slices",
				"strconv":                          "strconv",
				"strings":                          "strings",
				"sync":                             "sync",
				"unicode/utf8":                     "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Decoder":        reflect.TypeOf((*q.Decoder)(nil)).Elem(),
				"Encoder":        reflect.TypeOf((*q.Encoder)(nil)).Elem(),
				"Kind":           reflect.TypeOf((*q.Kind)(nil)).Elem(),
				"Pointer":        reflect.TypeOf((*q.Pointer)(nil)).Elem(),
				"SyntacticError": reflect.TypeOf((*q.SyntacticError)(nil)).Elem(),
				"Token":          reflect.TypeOf((*q.Token)(nil)).Elem(),
				"Value":          reflect.TypeOf((*q.Value)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"Options": reflect.TypeOf((*q.Options)(nil)).Elem(),
			},
			Vars: map[string]reflect.Value{
				"BeginArray":       reflect.ValueOf(&q.BeginArray),
				"BeginObject":      reflect.ValueOf(&q.BeginObject),
				"EndArray":         reflect.ValueOf(&q.EndArray),
				"EndObject":        reflect.ValueOf(&q.EndObject),
				"ErrDuplicateName": reflect.ValueOf(&q.ErrDuplicateName),
				"ErrNonStringName": reflect.ValueOf(&q.ErrNonStringName),
				"False":            reflect.ValueOf(&q.False),
				"Internal":         reflect.ValueOf(&q.Internal),
				"Null":             reflect.ValueOf(&q.Null),
				"True":             reflect.ValueOf(&q.True),
			},
			Funcs: map[string]reflect.Value{
				"AllowDuplicateNames":   reflect.ValueOf(q.AllowDuplicateNames),
				"AllowInvalidUTF8":      reflect.ValueOf(q.AllowInvalidUTF8),
				"AppendFloat":           reflect.ValueOf(q.AppendFloat),
				"Bool":                  reflect.ValueOf(q.Bool),
				"CanonicalizeRawFloats": reflect.ValueOf(q.CanonicalizeRawFloats),
				"CanonicalizeRawInts":   reflect.ValueOf(q.CanonicalizeRawInts),
				"EscapeForHTML":         reflect.ValueOf(q.EscapeForHTML),
				"EscapeForJS":           reflect.ValueOf(q.EscapeForJS),
				"Float":                 reflect.ValueOf(q.Float),
				"Float32":               reflect.ValueOf(q.Float32),
				"Int":                   reflect.ValueOf(q.Int),
				"Multiline":             reflect.ValueOf(q.Multiline),
				"NewDecoder":            reflect.ValueOf(q.NewDecoder),
				"NewEncoder":            reflect.ValueOf(q.NewEncoder),
				"PreserveRawStrings":    reflect.ValueOf(q.PreserveRawStrings),
				"ReorderRawObjects":     reflect.ValueOf(q.ReorderRawObjects),
				"SpaceAfterColon":       reflect.ValueOf(q.SpaceAfterColon),
				"SpaceAfterComma":       reflect.ValueOf(q.SpaceAfterComma),
				"String":                reflect.ValueOf(q.String),
				"Uint":                  reflect.ValueOf(q.Uint),
				"WithIndent":            reflect.ValueOf(q.WithIndent),
				"WithIndentPrefix":      reflect.ValueOf(q.WithIndentPrefix),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"KindBeginArray":  {Typ: reflect.TypeOf(q.KindBeginArray), Value: constant.MakeInt64(int64(q.KindBeginArray))},
				"KindBeginObject": {Typ: reflect.TypeOf(q.KindBeginObject), Value: constant.MakeInt64(int64(q.KindBeginObject))},
				"KindEndArray":    {Typ: reflect.TypeOf(q.KindEndArray), Value: constant.MakeInt64(int64(q.KindEndArray))},
				"KindEndObject":   {Typ: reflect.TypeOf(q.KindEndObject), Value: constant.MakeInt64(int64(q.KindEndObject))},
				"KindFalse":       {Typ: reflect.TypeOf(q.KindFalse), Value: constant.MakeInt64(int64(q.KindFalse))},
				"KindInvalid":     {Typ: reflect.TypeOf(q.KindInvalid), Value: constant.MakeInt64(int64(q.KindInvalid))},
				"KindNull":        {Typ: reflect.TypeOf(q.KindNull), Value: constant.MakeInt64(int64(q.KindNull))},
				"KindNumber":      {Typ: reflect.TypeOf(q.KindNumber), Value: constant.MakeInt64(int64(q.KindNumber))},
				"KindString":      {Typ: reflect.TypeOf(q.KindString), Value: constant.MakeInt64(int64(q.KindString))},
				"KindTrue":        {Typ: reflect.TypeOf(q.KindTrue), Value: constant.MakeInt64(int64(q.KindTrue))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
