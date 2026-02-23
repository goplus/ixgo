// export by github.com/goplus/ixgo/cmd/qexp

package testdeps

import (
	q "github.com/goplus/ixgo/x/testdeps"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"CoverMarkProfileEmittedFunc": reflect.ValueOf(&q.CoverMarkProfileEmittedFunc),
			"CoverMode":                   reflect.ValueOf(&q.CoverMode),
			"CoverProcessTestDirFunc":     reflect.ValueOf(&q.CoverProcessTestDirFunc),
			"CoverSelectedPackages":       reflect.ValueOf(&q.CoverSelectedPackages),
			"CoverSnapshotFunc":           reflect.ValueOf(&q.CoverSnapshotFunc),
			"Covered":                     reflect.ValueOf(&q.Covered),
			"ImportPath":                  reflect.ValueOf(&q.ImportPath),
			"ModulePath":                  reflect.ValueOf(&q.ModulePath),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
