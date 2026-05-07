// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package runtime

import (
	q "runtime"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("runtime", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "runtime",
			Path: "runtime",
			Deps: map[string]string{
				"internal/abi":              "abi",
				"internal/bytealg":          "bytealg",
				"internal/byteorder":        "byteorder",
				"internal/chacha8rand":      "chacha8rand",
				"internal/coverage/rtcov":   "rtcov",
				"internal/cpu":              "cpu",
				"internal/goarch":           "goarch",
				"internal/godebugs":         "godebugs",
				"internal/goexperiment":     "goexperiment",
				"internal/goos":             "goos",
				"internal/profilerecord":    "profilerecord",
				"internal/runtime/atomic":   "atomic",
				"internal/runtime/exithook": "exithook",
				"internal/runtime/maps":     "maps",
				"internal/runtime/math":     "math",
				"internal/runtime/sys":      "sys",
				"internal/stringslite":      "stringslite",
				"unsafe":                    "unsafe",
			},
			Interfaces: map[string]reflect.Type{
				"Error": reflect.TypeOf((*q.Error)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"BlockProfileRecord": reflect.TypeOf((*q.BlockProfileRecord)(nil)).Elem(),
				"Cleanup":            reflect.TypeOf((*q.Cleanup)(nil)).Elem(),
				"Frame":              reflect.TypeOf((*q.Frame)(nil)).Elem(),
				"Frames":             reflect.TypeOf((*q.Frames)(nil)).Elem(),
				"Func":               reflect.TypeOf((*q.Func)(nil)).Elem(),
				"MemProfileRecord":   reflect.TypeOf((*q.MemProfileRecord)(nil)).Elem(),
				"MemStats":           reflect.TypeOf((*q.MemStats)(nil)).Elem(),
				"PanicNilError":      reflect.TypeOf((*q.PanicNilError)(nil)).Elem(),
				"Pinner":             reflect.TypeOf((*q.Pinner)(nil)).Elem(),
				"StackRecord":        reflect.TypeOf((*q.StackRecord)(nil)).Elem(),
				"TypeAssertionError": reflect.TypeOf((*q.TypeAssertionError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"MemProfileRate": &q.MemProfileRate,
			},
			Funcs: map[string]interface{}{
				"BlockProfile":            q.BlockProfile,
				"Breakpoint":              q.Breakpoint,
				"CPUProfile":              q.CPUProfile,
				"Caller":                  q.Caller,
				"Callers":                 q.Callers,
				"CallersFrames":           q.CallersFrames,
				"FuncForPC":               q.FuncForPC,
				"GC":                      q.GC,
				"GOMAXPROCS":              q.GOMAXPROCS,
				"GOROOT":                  q.GOROOT,
				"Goexit":                  q.Goexit,
				"GoroutineProfile":        q.GoroutineProfile,
				"Gosched":                 q.Gosched,
				"KeepAlive":               q.KeepAlive,
				"LockOSThread":            q.LockOSThread,
				"MemProfile":              q.MemProfile,
				"MutexProfile":            q.MutexProfile,
				"NumCPU":                  q.NumCPU,
				"NumCgoCall":              q.NumCgoCall,
				"NumGoroutine":            q.NumGoroutine,
				"ReadMemStats":            q.ReadMemStats,
				"ReadTrace":               q.ReadTrace,
				"SetBlockProfileRate":     q.SetBlockProfileRate,
				"SetCPUProfileRate":       q.SetCPUProfileRate,
				"SetCgoTraceback":         q.SetCgoTraceback,
				"SetFinalizer":            q.SetFinalizer,
				"SetMutexProfileFraction": q.SetMutexProfileFraction,
				"Stack":                   q.Stack,
				"StartTrace":              q.StartTrace,
				"StopTrace":               q.StopTrace,
				"ThreadCreateProfile":     q.ThreadCreateProfile,
				"UnlockOSThread":          q.UnlockOSThread,
				"Version":                 q.Version,
			},
			TypedConsts: map[string]interface{}{
				"GOARCH": q.GOARCH,
				"GOOS":   q.GOOS,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Compiler": {Typ: "untyped string", Value: constant.MakeString(string(q.Compiler))},
			},
		}
	})
}
