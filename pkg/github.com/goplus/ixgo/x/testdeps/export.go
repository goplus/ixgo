// export by github.com/goplus/ixgo/cmd/qexp

package testdeps

import (
	q "github.com/goplus/ixgo/x/testdeps"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/ixgo/x/testdeps", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "testdeps",
			Path: "github.com/goplus/ixgo/x/testdeps",
			Deps: map[string]string{
				"bufio":                            "bufio",
				"github.com/goplus/ixgo/x/testlog": "testlog",
				"io":                               "io",
				"os":                               "os",
				"reflect":                          "reflect",
				"regexp":                           "regexp",
				"runtime/pprof":                    "pprof",
				"strings":                          "strings",
				"sync":                             "sync",
				"time":                             "time",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"TestDeps": reflect.TypeOf((*q.TestDeps)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"CorpusEntry": reflect.TypeOf((*q.CorpusEntry)(nil)).Elem(),
			},
			Vars: map[string]interface{}{
				"CoverMarkProfileEmittedFunc": &q.CoverMarkProfileEmittedFunc,
				"CoverMode":                   &q.CoverMode,
				"CoverProcessTestDirFunc":     &q.CoverProcessTestDirFunc,
				"CoverSelectedPackages":       &q.CoverSelectedPackages,
				"CoverSnapshotFunc":           &q.CoverSnapshotFunc,
				"Covered":                     &q.Covered,
				"ImportPath":                  &q.ImportPath,
				"ModulePath":                  &q.ModulePath,
			},
			Funcs:         map[string]interface{}{},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
