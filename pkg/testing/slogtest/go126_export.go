// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package slogtest

import (
	q "testing/slogtest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Run":         reflect.ValueOf(q.Run),
			"TestHandler": reflect.ValueOf(q.TestHandler),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
