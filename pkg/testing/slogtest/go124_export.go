// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package slogtest

import (
	q "testing/slogtest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("testing/slogtest", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "slogtest",
			Path: "testing/slogtest",
			Deps: map[string]string{
				"context":  "context",
				"errors":   "errors",
				"fmt":      "fmt",
				"log/slog": "slog",
				"reflect":  "reflect",
				"runtime":  "runtime",
				"testing":  "testing",
				"time":     "time",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Run":         q.Run,
				"TestHandler": q.TestHandler,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
