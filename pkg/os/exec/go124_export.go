// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package exec

import (
	q "os/exec"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("os/exec", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "exec",
			Path: "os/exec",
			Deps: map[string]string{
				"bytes":                    "bytes",
				"context":                  "context",
				"errors":                   "errors",
				"internal/godebug":         "godebug",
				"internal/syscall/execenv": "execenv",
				"internal/syscall/unix":    "unix",
				"io":                       "io",
				"io/fs":                    "fs",
				"os":                       "os",
				"path/filepath":            "filepath",
				"runtime":                  "runtime",
				"strconv":                  "strconv",
				"strings":                  "strings",
				"syscall":                  "syscall",
				"time":                     "time",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Cmd":       reflect.TypeOf((*q.Cmd)(nil)).Elem(),
				"Error":     reflect.TypeOf((*q.Error)(nil)).Elem(),
				"ExitError": reflect.TypeOf((*q.ExitError)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrDot":       &q.ErrDot,
				"ErrNotFound":  &q.ErrNotFound,
				"ErrWaitDelay": &q.ErrWaitDelay,
			},
			Funcs: map[string]interface{}{
				"Command":        q.Command,
				"CommandContext": q.CommandContext,
				"LookPath":       q.LookPath,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
