// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package pprof

import (
	q "net/http/pprof"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "pprof",
		Path: "net/http/pprof",
		Deps: map[string]string{
			"bufio":                 "bufio",
			"bytes":                 "bytes",
			"context":               "context",
			"fmt":                   "fmt",
			"html":                  "html",
			"internal/godebug":      "godebug",
			"internal/goexperiment": "goexperiment",
			"internal/profile":      "profile",
			"io":                    "io",
			"log":                   "log",
			"net/http":              "http",
			"net/url":               "url",
			"os":                    "os",
			"runtime":               "runtime",
			"runtime/pprof":         "pprof",
			"runtime/trace":         "trace",
			"slices":                "slices",
			"strconv":               "strconv",
			"strings":               "strings",
			"time":                  "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Cmdline": reflect.ValueOf(q.Cmdline),
			"Handler": reflect.ValueOf(q.Handler),
			"Index":   reflect.ValueOf(q.Index),
			"Profile": reflect.ValueOf(q.Profile),
			"Symbol":  reflect.ValueOf(q.Symbol),
			"Trace":   reflect.ValueOf(q.Trace),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
