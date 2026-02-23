// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package coverage

import (
	q "runtime/coverage"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "coverage",
		Path: "runtime/coverage",
		Deps: map[string]string{
			"internal/coverage/cfile": "cfile",
			"io":                      "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ClearCounters":    reflect.ValueOf(q.ClearCounters),
			"WriteCounters":    reflect.ValueOf(q.WriteCounters),
			"WriteCountersDir": reflect.ValueOf(q.WriteCountersDir),
			"WriteMeta":        reflect.ValueOf(q.WriteMeta),
			"WriteMetaDir":     reflect.ValueOf(q.WriteMetaDir),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
