// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package trace

import (
	q "runtime/trace"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("runtime/trace", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"IsEnabled":         q.IsEnabled,
				"Log":               q.Log,
				"Logf":              q.Logf,
				"NewFlightRecorder": q.NewFlightRecorder,
				"NewTask":           q.NewTask,
				"Start":             q.Start,
				"StartRegion":       q.StartRegion,
				"Stop":              q.Stop,
				"WithRegion":        q.WithRegion,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
