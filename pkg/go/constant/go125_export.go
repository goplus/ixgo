// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package constant

import (
	q "go/constant"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"BinaryOp":        reflect.ValueOf(q.BinaryOp),
			"BitLen":          reflect.ValueOf(q.BitLen),
			"BoolVal":         reflect.ValueOf(q.BoolVal),
			"Bytes":           reflect.ValueOf(q.Bytes),
			"Compare":         reflect.ValueOf(q.Compare),
			"Denom":           reflect.ValueOf(q.Denom),
			"Float32Val":      reflect.ValueOf(q.Float32Val),
			"Float64Val":      reflect.ValueOf(q.Float64Val),
			"Imag":            reflect.ValueOf(q.Imag),
			"Int64Val":        reflect.ValueOf(q.Int64Val),
			"Make":            reflect.ValueOf(q.Make),
			"MakeBool":        reflect.ValueOf(q.MakeBool),
			"MakeFloat64":     reflect.ValueOf(q.MakeFloat64),
			"MakeFromBytes":   reflect.ValueOf(q.MakeFromBytes),
			"MakeFromLiteral": reflect.ValueOf(q.MakeFromLiteral),
			"MakeImag":        reflect.ValueOf(q.MakeImag),
			"MakeInt64":       reflect.ValueOf(q.MakeInt64),
			"MakeString":      reflect.ValueOf(q.MakeString),
			"MakeUint64":      reflect.ValueOf(q.MakeUint64),
			"MakeUnknown":     reflect.ValueOf(q.MakeUnknown),
			"Num":             reflect.ValueOf(q.Num),
			"Real":            reflect.ValueOf(q.Real),
			"Shift":           reflect.ValueOf(q.Shift),
			"Sign":            reflect.ValueOf(q.Sign),
			"StringVal":       reflect.ValueOf(q.StringVal),
			"ToComplex":       reflect.ValueOf(q.ToComplex),
			"ToFloat":         reflect.ValueOf(q.ToFloat),
			"ToInt":           reflect.ValueOf(q.ToInt),
			"Uint64Val":       reflect.ValueOf(q.Uint64Val),
			"UnaryOp":         reflect.ValueOf(q.UnaryOp),
			"Val":             reflect.ValueOf(q.Val),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"Bool":    {reflect.TypeOf(q.Bool), constant.MakeInt64(int64(q.Bool))},
			"Complex": {reflect.TypeOf(q.Complex), constant.MakeInt64(int64(q.Complex))},
			"Float":   {reflect.TypeOf(q.Float), constant.MakeInt64(int64(q.Float))},
			"Int":     {reflect.TypeOf(q.Int), constant.MakeInt64(int64(q.Int))},
			"String":  {reflect.TypeOf(q.String), constant.MakeInt64(int64(q.String))},
			"Unknown": {reflect.TypeOf(q.Unknown), constant.MakeInt64(int64(q.Unknown))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
