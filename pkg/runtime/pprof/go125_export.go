// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package pprof

import (
	q "runtime/pprof"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "pprof",
		Path: "runtime/pprof",
		Deps: map[string]string{
			"bufio":                  "bufio",
			"bytes":                  "bytes",
			"cmp":                    "cmp",
			"compress/gzip":          "gzip",
			"context":                "context",
			"encoding/binary":        "binary",
			"errors":                 "errors",
			"fmt":                    "fmt",
			"internal/abi":           "abi",
			"internal/byteorder":     "byteorder",
			"internal/profilerecord": "profilerecord",
			"io":                     "io",
			"math":                   "math",
			"os":                     "os",
			"runtime":                "runtime",
			"slices":                 "slices",
			"sort":                   "sort",
			"strconv":                "strconv",
			"strings":                "strings",
			"sync":                   "sync",
			"syscall":                "syscall",
			"text/tabwriter":         "tabwriter",
			"time":                   "time",
			"unsafe":                 "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"LabelSet": reflect.TypeOf((*q.LabelSet)(nil)).Elem(),
			"Profile":  reflect.TypeOf((*q.Profile)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Do":                 reflect.ValueOf(q.Do),
			"ForLabels":          reflect.ValueOf(q.ForLabels),
			"Label":              reflect.ValueOf(q.Label),
			"Labels":             reflect.ValueOf(q.Labels),
			"Lookup":             reflect.ValueOf(q.Lookup),
			"NewProfile":         reflect.ValueOf(q.NewProfile),
			"Profiles":           reflect.ValueOf(q.Profiles),
			"SetGoroutineLabels": reflect.ValueOf(q.SetGoroutineLabels),
			"StartCPUProfile":    reflect.ValueOf(q.StartCPUProfile),
			"StopCPUProfile":     reflect.ValueOf(q.StopCPUProfile),
			"WithLabels":         reflect.ValueOf(q.WithLabels),
			"WriteHeapProfile":   reflect.ValueOf(q.WriteHeapProfile),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
