// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package constant

import (
	q "go/constant"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("go/constant", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "constant",
			Path: "go/constant",
			Deps: map[string]string{
				"fmt":          "fmt",
				"go/token":     "token",
				"math":         "math",
				"math/big":     "big",
				"math/bits":    "bits",
				"strconv":      "strconv",
				"strings":      "strings",
				"sync":         "sync",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{
				"Value": reflect.TypeOf((*q.Value)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Kind": reflect.TypeOf((*q.Kind)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"BinaryOp":        q.BinaryOp,
				"BitLen":          q.BitLen,
				"BoolVal":         q.BoolVal,
				"Bytes":           q.Bytes,
				"Compare":         q.Compare,
				"Denom":           q.Denom,
				"Float32Val":      q.Float32Val,
				"Float64Val":      q.Float64Val,
				"Imag":            q.Imag,
				"Int64Val":        q.Int64Val,
				"Make":            q.Make,
				"MakeBool":        q.MakeBool,
				"MakeFloat64":     q.MakeFloat64,
				"MakeFromBytes":   q.MakeFromBytes,
				"MakeFromLiteral": q.MakeFromLiteral,
				"MakeImag":        q.MakeImag,
				"MakeInt64":       q.MakeInt64,
				"MakeString":      q.MakeString,
				"MakeUint64":      q.MakeUint64,
				"MakeUnknown":     q.MakeUnknown,
				"Num":             q.Num,
				"Real":            q.Real,
				"Shift":           q.Shift,
				"Sign":            q.Sign,
				"StringVal":       q.StringVal,
				"ToComplex":       q.ToComplex,
				"ToFloat":         q.ToFloat,
				"ToInt":           q.ToInt,
				"Uint64Val":       q.Uint64Val,
				"UnaryOp":         q.UnaryOp,
				"Val":             q.Val,
			},
			TypedConsts: map[string]interface{}{
				"Bool":    q.Bool,
				"Complex": q.Complex,
				"Float":   q.Float,
				"Int":     q.Int,
				"String":  q.String,
				"Unknown": q.Unknown,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
