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
	ixgo.RegisterPackage("runtime/pprof", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Do":                 q.Do,
				"ForLabels":          q.ForLabels,
				"Label":              q.Label,
				"Labels":             q.Labels,
				"Lookup":             q.Lookup,
				"NewProfile":         q.NewProfile,
				"Profiles":           q.Profiles,
				"SetGoroutineLabels": q.SetGoroutineLabels,
				"StartCPUProfile":    q.StartCPUProfile,
				"StopCPUProfile":     q.StopCPUProfile,
				"WithLabels":         q.WithLabels,
				"WriteHeapProfile":   q.WriteHeapProfile,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
