// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package buildinfo

import (
	q "debug/buildinfo"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "buildinfo",
		Path: "debug/buildinfo",
		Deps: map[string]string{
			"bytes":            "bytes",
			"debug/elf":        "elf",
			"debug/macho":      "macho",
			"debug/pe":         "pe",
			"debug/plan9obj":   "plan9obj",
			"encoding/binary":  "binary",
			"errors":           "errors",
			"fmt":              "fmt",
			"internal/saferio": "saferio",
			"internal/xcoff":   "xcoff",
			"io":               "io",
			"io/fs":            "fs",
			"os":               "os",
			"runtime/debug":    "debug",
			"unsafe":           "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"BuildInfo": reflect.TypeOf((*q.BuildInfo)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Read":     reflect.ValueOf(q.Read),
			"ReadFile": reflect.ValueOf(q.ReadFile),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
