// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package metrics

import (
	q "runtime/metrics"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "metrics",
		Path: "runtime/metrics",
		Deps: map[string]string{
			"internal/godebugs": "godebugs",
			"math":              "math",
			"runtime":           "runtime",
			"unsafe":            "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Description":      reflect.TypeOf((*q.Description)(nil)).Elem(),
			"Float64Histogram": reflect.TypeOf((*q.Float64Histogram)(nil)).Elem(),
			"Sample":           reflect.TypeOf((*q.Sample)(nil)).Elem(),
			"Value":            reflect.TypeOf((*q.Value)(nil)).Elem(),
			"ValueKind":        reflect.TypeOf((*q.ValueKind)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"All":  reflect.ValueOf(q.All),
			"Read": reflect.ValueOf(q.Read),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"KindBad":              {reflect.TypeOf(q.KindBad), constant.MakeInt64(int64(q.KindBad))},
			"KindFloat64":          {reflect.TypeOf(q.KindFloat64), constant.MakeInt64(int64(q.KindFloat64))},
			"KindFloat64Histogram": {reflect.TypeOf(q.KindFloat64Histogram), constant.MakeInt64(int64(q.KindFloat64Histogram))},
			"KindUint64":           {reflect.TypeOf(q.KindUint64), constant.MakeInt64(int64(q.KindUint64))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
