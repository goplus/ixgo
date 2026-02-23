// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package cmplx

import (
	q "math/cmplx"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "cmplx",
		Path: "math/cmplx",
		Deps: map[string]string{
			"math":      "math",
			"math/bits": "bits",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Abs":   reflect.ValueOf(q.Abs),
			"Acos":  reflect.ValueOf(q.Acos),
			"Acosh": reflect.ValueOf(q.Acosh),
			"Asin":  reflect.ValueOf(q.Asin),
			"Asinh": reflect.ValueOf(q.Asinh),
			"Atan":  reflect.ValueOf(q.Atan),
			"Atanh": reflect.ValueOf(q.Atanh),
			"Conj":  reflect.ValueOf(q.Conj),
			"Cos":   reflect.ValueOf(q.Cos),
			"Cosh":  reflect.ValueOf(q.Cosh),
			"Cot":   reflect.ValueOf(q.Cot),
			"Exp":   reflect.ValueOf(q.Exp),
			"Inf":   reflect.ValueOf(q.Inf),
			"IsInf": reflect.ValueOf(q.IsInf),
			"IsNaN": reflect.ValueOf(q.IsNaN),
			"Log":   reflect.ValueOf(q.Log),
			"Log10": reflect.ValueOf(q.Log10),
			"NaN":   reflect.ValueOf(q.NaN),
			"Phase": reflect.ValueOf(q.Phase),
			"Polar": reflect.ValueOf(q.Polar),
			"Pow":   reflect.ValueOf(q.Pow),
			"Rect":  reflect.ValueOf(q.Rect),
			"Sin":   reflect.ValueOf(q.Sin),
			"Sinh":  reflect.ValueOf(q.Sinh),
			"Sqrt":  reflect.ValueOf(q.Sqrt),
			"Tan":   reflect.ValueOf(q.Tan),
			"Tanh":  reflect.ValueOf(q.Tanh),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
