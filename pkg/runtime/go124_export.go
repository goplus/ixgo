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
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"MemProfileRate": reflect.ValueOf(&q.MemProfileRate),
		},
		Funcs: map[string]reflect.Value{
			"BlockProfile":            reflect.ValueOf(q.BlockProfile),
			"Breakpoint":              reflect.ValueOf(q.Breakpoint),
			"CPUProfile":              reflect.ValueOf(q.CPUProfile),
			"Caller":                  reflect.ValueOf(q.Caller),
			"Callers":                 reflect.ValueOf(q.Callers),
			"CallersFrames":           reflect.ValueOf(q.CallersFrames),
			"FuncForPC":               reflect.ValueOf(q.FuncForPC),
			"GC":                      reflect.ValueOf(q.GC),
			"GOMAXPROCS":              reflect.ValueOf(q.GOMAXPROCS),
			"GOROOT":                  reflect.ValueOf(q.GOROOT),
			"Goexit":                  reflect.ValueOf(q.Goexit),
			"GoroutineProfile":        reflect.ValueOf(q.GoroutineProfile),
			"Gosched":                 reflect.ValueOf(q.Gosched),
			"KeepAlive":               reflect.ValueOf(q.KeepAlive),
			"LockOSThread":            reflect.ValueOf(q.LockOSThread),
			"MemProfile":              reflect.ValueOf(q.MemProfile),
			"MutexProfile":            reflect.ValueOf(q.MutexProfile),
			"NumCPU":                  reflect.ValueOf(q.NumCPU),
			"NumCgoCall":              reflect.ValueOf(q.NumCgoCall),
			"NumGoroutine":            reflect.ValueOf(q.NumGoroutine),
			"ReadMemStats":            reflect.ValueOf(q.ReadMemStats),
			"ReadTrace":               reflect.ValueOf(q.ReadTrace),
			"SetBlockProfileRate":     reflect.ValueOf(q.SetBlockProfileRate),
			"SetCPUProfileRate":       reflect.ValueOf(q.SetCPUProfileRate),
			"SetCgoTraceback":         reflect.ValueOf(q.SetCgoTraceback),
			"SetFinalizer":            reflect.ValueOf(q.SetFinalizer),
			"SetMutexProfileFraction": reflect.ValueOf(q.SetMutexProfileFraction),
			"Stack":                   reflect.ValueOf(q.Stack),
			"StartTrace":              reflect.ValueOf(q.StartTrace),
			"StopTrace":               reflect.ValueOf(q.StopTrace),
			"ThreadCreateProfile":     reflect.ValueOf(q.ThreadCreateProfile),
			"UnlockOSThread":          reflect.ValueOf(q.UnlockOSThread),
			"Version":                 reflect.ValueOf(q.Version),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"GOARCH": {reflect.TypeOf(q.GOARCH), constant.MakeString(string(q.GOARCH))},
			"GOOS":   {reflect.TypeOf(q.GOOS), constant.MakeString(string(q.GOOS))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"Compiler": {"untyped string", constant.MakeString(string(q.Compiler))},
		},
	})
}
