// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package big

import (
	q "math/big"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Jacobi":     reflect.ValueOf(q.Jacobi),
			"NewFloat":   reflect.ValueOf(q.NewFloat),
			"NewInt":     reflect.ValueOf(q.NewInt),
			"NewRat":     reflect.ValueOf(q.NewRat),
			"ParseFloat": reflect.ValueOf(q.ParseFloat),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"Above":         {Typ: reflect.TypeOf(q.Above), Value: constant.MakeInt64(int64(q.Above))},
			"AwayFromZero":  {Typ: reflect.TypeOf(q.AwayFromZero), Value: constant.MakeInt64(int64(q.AwayFromZero))},
			"Below":         {Typ: reflect.TypeOf(q.Below), Value: constant.MakeInt64(int64(q.Below))},
			"Exact":         {Typ: reflect.TypeOf(q.Exact), Value: constant.MakeInt64(int64(q.Exact))},
			"ToNearestAway": {Typ: reflect.TypeOf(q.ToNearestAway), Value: constant.MakeInt64(int64(q.ToNearestAway))},
			"ToNearestEven": {Typ: reflect.TypeOf(q.ToNearestEven), Value: constant.MakeInt64(int64(q.ToNearestEven))},
			"ToNegativeInf": {Typ: reflect.TypeOf(q.ToNegativeInf), Value: constant.MakeInt64(int64(q.ToNegativeInf))},
			"ToPositiveInf": {Typ: reflect.TypeOf(q.ToPositiveInf), Value: constant.MakeInt64(int64(q.ToPositiveInf))},
			"ToZero":        {Typ: reflect.TypeOf(q.ToZero), Value: constant.MakeInt64(int64(q.ToZero))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"MaxBase": {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.MaxBase))},
			"MaxExp":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxExp))},
			"MaxPrec": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxPrec))},
			"MinExp":  {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MinExp))},
		},
	})
}
