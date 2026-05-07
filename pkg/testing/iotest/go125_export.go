// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package iotest

import (
	q "testing/iotest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("testing/iotest", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "iotest",
			Path: "testing/iotest",
			Deps: map[string]string{
				"bytes":  "bytes",
				"errors": "errors",
				"fmt":    "fmt",
				"io":     "io",
				"log":    "log",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrTimeout": &q.ErrTimeout,
			},
			Funcs: map[string]interface{}{
				"DataErrReader":  q.DataErrReader,
				"ErrReader":      q.ErrReader,
				"HalfReader":     q.HalfReader,
				"NewReadLogger":  q.NewReadLogger,
				"NewWriteLogger": q.NewWriteLogger,
				"OneByteReader":  q.OneByteReader,
				"TestReader":     q.TestReader,
				"TimeoutReader":  q.TimeoutReader,
				"TruncateWriter": q.TruncateWriter,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
