// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package metrics

import (
	q "runtime/metrics"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("runtime/metrics", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"All":  q.All,
				"Read": q.Read,
			},
			TypedConsts: map[string]interface{}{
				"KindBad":              q.KindBad,
				"KindFloat64":          q.KindFloat64,
				"KindFloat64Histogram": q.KindFloat64Histogram,
				"KindUint64":           q.KindUint64,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
