// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package asn1

import (
	q "encoding/asn1"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/asn1", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "asn1",
			Path: "encoding/asn1",
			Deps: map[string]string{
				"bytes":            "bytes",
				"errors":           "errors",
				"fmt":              "fmt",
				"internal/saferio": "saferio",
				"math":             "math",
				"math/big":         "big",
				"reflect":          "reflect",
				"slices":           "slices",
				"strconv":          "strconv",
				"strings":          "strings",
				"time":             "time",
				"unicode/utf16":    "utf16",
				"unicode/utf8":     "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"BitString":        reflect.TypeOf((*q.BitString)(nil)).Elem(),
				"Enumerated":       reflect.TypeOf((*q.Enumerated)(nil)).Elem(),
				"Flag":             reflect.TypeOf((*q.Flag)(nil)).Elem(),
				"ObjectIdentifier": reflect.TypeOf((*q.ObjectIdentifier)(nil)).Elem(),
				"RawContent":       reflect.TypeOf((*q.RawContent)(nil)).Elem(),
				"RawValue":         reflect.TypeOf((*q.RawValue)(nil)).Elem(),
				"StructuralError":  reflect.TypeOf((*q.StructuralError)(nil)).Elem(),
				"SyntaxError":      reflect.TypeOf((*q.SyntaxError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"NullBytes":    &q.NullBytes,
				"NullRawValue": &q.NullRawValue,
			},
			Funcs: map[string]interface{}{
				"Marshal":             q.Marshal,
				"MarshalWithParams":   q.MarshalWithParams,
				"Unmarshal":           q.Unmarshal,
				"UnmarshalWithParams": q.UnmarshalWithParams,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"ClassApplication":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ClassApplication))},
				"ClassContextSpecific": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ClassContextSpecific))},
				"ClassPrivate":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ClassPrivate))},
				"ClassUniversal":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.ClassUniversal))},
				"TagBMPString":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagBMPString))},
				"TagBitString":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagBitString))},
				"TagBoolean":           {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagBoolean))},
				"TagEnum":              {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagEnum))},
				"TagGeneralString":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagGeneralString))},
				"TagGeneralizedTime":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagGeneralizedTime))},
				"TagIA5String":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagIA5String))},
				"TagInteger":           {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagInteger))},
				"TagNull":              {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagNull))},
				"TagNumericString":     {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagNumericString))},
				"TagOID":               {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagOID))},
				"TagOctetString":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagOctetString))},
				"TagPrintableString":   {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagPrintableString))},
				"TagSequence":          {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagSequence))},
				"TagSet":               {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagSet))},
				"TagT61String":         {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagT61String))},
				"TagUTCTime":           {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagUTCTime))},
				"TagUTF8String":        {Typ: "untyped int", Value: constant.MakeInt64(int64(q.TagUTF8String))},
			},
		}
	})
}
