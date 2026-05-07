// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package pprof

import (
	q "net/http/pprof"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("net/http/pprof", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "pprof",
			Path: "net/http/pprof",
			Deps: map[string]string{
				"bufio":            "bufio",
				"bytes":            "bytes",
				"context":          "context",
				"fmt":              "fmt",
				"html":             "html",
				"internal/godebug": "godebug",
				"internal/profile": "profile",
				"io":               "io",
				"log":              "log",
				"net/http":         "http",
				"net/url":          "url",
				"os":               "os",
				"runtime":          "runtime",
				"runtime/pprof":    "pprof",
				"runtime/trace":    "trace",
				"slices":           "slices",
				"strconv":          "strconv",
				"strings":          "strings",
				"time":             "time",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Cmdline": q.Cmdline,
				"Handler": q.Handler,
				"Index":   q.Index,
				"Profile": q.Profile,
				"Symbol":  q.Symbol,
				"Trace":   q.Trace,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
