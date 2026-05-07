// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package cmplx

import (
	q "math/cmplx"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("math/cmplx", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "cmplx",
			Path: "math/cmplx",
			Deps: map[string]string{
				"math":      "math",
				"math/bits": "bits",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Abs":   q.Abs,
				"Acos":  q.Acos,
				"Acosh": q.Acosh,
				"Asin":  q.Asin,
				"Asinh": q.Asinh,
				"Atan":  q.Atan,
				"Atanh": q.Atanh,
				"Conj":  q.Conj,
				"Cos":   q.Cos,
				"Cosh":  q.Cosh,
				"Cot":   q.Cot,
				"Exp":   q.Exp,
				"Inf":   q.Inf,
				"IsInf": q.IsInf,
				"IsNaN": q.IsNaN,
				"Log":   q.Log,
				"Log10": q.Log10,
				"NaN":   q.NaN,
				"Phase": q.Phase,
				"Polar": q.Polar,
				"Pow":   q.Pow,
				"Rect":  q.Rect,
				"Sin":   q.Sin,
				"Sinh":  q.Sinh,
				"Sqrt":  q.Sqrt,
				"Tan":   q.Tan,
				"Tanh":  q.Tanh,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
