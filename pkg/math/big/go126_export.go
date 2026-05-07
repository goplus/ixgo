// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package big

import (
	q "math/big"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("math/big", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "big",
			Path: "math/big",
			Deps: map[string]string{
				"bytes":              "bytes",
				"errors":             "errors",
				"fmt":                "fmt",
				"internal/byteorder": "byteorder",
				"internal/cpu":       "cpu",
				"io":                 "io",
				"math":               "math",
				"math/bits":          "bits",
				"math/rand":          "rand",
				"slices":             "slices",
				"strconv":            "strconv",
				"strings":            "strings",
				"sync":               "sync",
				"unsafe":             "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Accuracy":     reflect.TypeOf((*q.Accuracy)(nil)).Elem(),
				"ErrNaN":       reflect.TypeOf((*q.ErrNaN)(nil)).Elem(),
				"Float":        reflect.TypeOf((*q.Float)(nil)).Elem(),
				"Int":          reflect.TypeOf((*q.Int)(nil)).Elem(),
				"Rat":          reflect.TypeOf((*q.Rat)(nil)).Elem(),
				"RoundingMode": reflect.TypeOf((*q.RoundingMode)(nil)).Elem(),
				"Word":         reflect.TypeOf((*q.Word)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Jacobi":     q.Jacobi,
				"NewFloat":   q.NewFloat,
				"NewInt":     q.NewInt,
				"NewRat":     q.NewRat,
				"ParseFloat": q.ParseFloat,
			},
			TypedConsts: map[string]interface{}{
				"Above":         q.Above,
				"AwayFromZero":  q.AwayFromZero,
				"Below":         q.Below,
				"Exact":         q.Exact,
				"ToNearestAway": q.ToNearestAway,
				"ToNearestEven": q.ToNearestEven,
				"ToNegativeInf": q.ToNegativeInf,
				"ToPositiveInf": q.ToPositiveInf,
				"ToZero":        q.ToZero,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"MaxBase": {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.MaxBase))},
				"MaxExp":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxExp))},
				"MaxPrec": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxPrec))},
				"MinExp":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MinExp))},
			},
		}
	})
}
