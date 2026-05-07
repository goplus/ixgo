// export by github.com/goplus/ixgo/cmd/qexp

package race

import (
	q "github.com/goplus/ixgo/x/race"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/ixgo/x/race", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "race",
			Path: "github.com/goplus/ixgo/x/race",
			Deps: map[string]string{
				"unsafe": "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Acquire":      q.Acquire,
				"Disable":      q.Disable,
				"Enable":       q.Enable,
				"Errors":       q.Errors,
				"Read":         q.Read,
				"ReadRange":    q.ReadRange,
				"Release":      q.Release,
				"ReleaseMerge": q.ReleaseMerge,
				"Write":        q.Write,
				"WriteRange":   q.WriteRange,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Enabled": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.Enabled))},
			},
		}
	})
}
