// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package trace

import (
	q "runtime/trace"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "trace",
		Path: "runtime/trace",
		Deps: map[string]string{
			"context":                "context",
			"errors":                 "errors",
			"fmt":                    "fmt",
			"internal/trace/tracev2": "tracev2",
			"io":                     "io",
			"runtime":                "runtime",
			"slices":                 "slices",
			"sync":                   "sync",
			"sync/atomic":            "atomic",
			"time":                   "time",
			"unsafe":                 "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"FlightRecorder":       reflect.TypeOf((*q.FlightRecorder)(nil)).Elem(),
			"FlightRecorderConfig": reflect.TypeOf((*q.FlightRecorderConfig)(nil)).Elem(),
			"Region":               reflect.TypeOf((*q.Region)(nil)).Elem(),
			"Task":                 reflect.TypeOf((*q.Task)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"IsEnabled":         reflect.ValueOf(q.IsEnabled),
			"Log":               reflect.ValueOf(q.Log),
			"Logf":              reflect.ValueOf(q.Logf),
			"NewFlightRecorder": reflect.ValueOf(q.NewFlightRecorder),
			"NewTask":           reflect.ValueOf(q.NewTask),
			"Start":             reflect.ValueOf(q.Start),
			"StartRegion":       reflect.ValueOf(q.StartRegion),
			"Stop":              reflect.ValueOf(q.Stop),
			"WithRegion":        reflect.ValueOf(q.WithRegion),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
