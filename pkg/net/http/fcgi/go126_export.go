// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package fcgi

import (
	q "net/http/fcgi"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "fcgi",
		Path: "net/http/fcgi",
		Deps: map[string]string{
			"bufio":           "bufio",
			"bytes":           "bytes",
			"context":         "context",
			"encoding/binary": "binary",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"net":             "net",
			"net/http":        "http",
			"net/http/cgi":    "cgi",
			"os":              "os",
			"strings":         "strings",
			"sync":            "sync",
			"time":            "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrConnClosed":     reflect.ValueOf(&q.ErrConnClosed),
			"ErrRequestAborted": reflect.ValueOf(&q.ErrRequestAborted),
		},
		Funcs: map[string]reflect.Value{
			"ProcessEnv": reflect.ValueOf(q.ProcessEnv),
			"Serve":      reflect.ValueOf(q.Serve),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
