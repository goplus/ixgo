// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package coverage

import (
	q "runtime/coverage"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("runtime/coverage", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "coverage",
			Path: "runtime/coverage",
			Deps: map[string]string{
				"internal/coverage/cfile": "cfile",
				"io":                      "io",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"ClearCounters":    q.ClearCounters,
				"WriteCounters":    q.WriteCounters,
				"WriteCountersDir": q.WriteCountersDir,
				"WriteMeta":        q.WriteMeta,
				"WriteMetaDir":     q.WriteMetaDir,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
