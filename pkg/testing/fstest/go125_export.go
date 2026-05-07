// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package fstest

import (
	q "testing/fstest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("testing/fstest", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "fstest",
			Path: "testing/fstest",
			Deps: map[string]string{
				"errors":         "errors",
				"fmt":            "fmt",
				"io":             "io",
				"io/fs":          "fs",
				"maps":           "maps",
				"path":           "path",
				"slices":         "slices",
				"strings":        "strings",
				"testing/iotest": "iotest",
				"time":           "time",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"MapFS":   reflect.TypeOf((*q.MapFS)(nil)).Elem(),
				"MapFile": reflect.TypeOf((*q.MapFile)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"TestFS": q.TestFS,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
