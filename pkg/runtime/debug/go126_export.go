// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package debug

import (
	q "runtime/debug"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("runtime/debug", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "debug",
			Path: "runtime/debug",
			Deps: map[string]string{
				"fmt":           "fmt",
				"internal/poll": "poll",
				"os":            "os",
				"runtime":       "runtime",
				"slices":        "slices",
				"strconv":       "strconv",
				"strings":       "strings",
				"time":          "time",
				"unsafe":        "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"BuildInfo":    reflect.TypeOf((*q.BuildInfo)(nil)).Elem(),
				"BuildSetting": reflect.TypeOf((*q.BuildSetting)(nil)).Elem(),
				"CrashOptions": reflect.TypeOf((*q.CrashOptions)(nil)).Elem(),
				"GCStats":      reflect.TypeOf((*q.GCStats)(nil)).Elem(),
				"Module":       reflect.TypeOf((*q.Module)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"FreeOSMemory":    q.FreeOSMemory,
				"ParseBuildInfo":  q.ParseBuildInfo,
				"PrintStack":      q.PrintStack,
				"ReadBuildInfo":   q.ReadBuildInfo,
				"ReadGCStats":     q.ReadGCStats,
				"SetCrashOutput":  q.SetCrashOutput,
				"SetGCPercent":    q.SetGCPercent,
				"SetMaxStack":     q.SetMaxStack,
				"SetMaxThreads":   q.SetMaxThreads,
				"SetMemoryLimit":  q.SetMemoryLimit,
				"SetPanicOnFault": q.SetPanicOnFault,
				"SetTraceback":    q.SetTraceback,
				"Stack":           q.Stack,
				"WriteHeapDump":   q.WriteHeapDump,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
