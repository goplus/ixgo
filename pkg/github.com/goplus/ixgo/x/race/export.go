// export by github.com/goplus/ixgo/cmd/qexp

package race

import (
	q "github.com/goplus/ixgo/x/race"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "race",
		Path: "github.com/goplus/ixgo/x/race",
		Deps: map[string]string{
			"unsafe": "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Acquire":      reflect.ValueOf(q.Acquire),
			"Disable":      reflect.ValueOf(q.Disable),
			"Enable":       reflect.ValueOf(q.Enable),
			"Errors":       reflect.ValueOf(q.Errors),
			"Read":         reflect.ValueOf(q.Read),
			"ReadRange":    reflect.ValueOf(q.ReadRange),
			"Release":      reflect.ValueOf(q.Release),
			"ReleaseMerge": reflect.ValueOf(q.ReleaseMerge),
			"Write":        reflect.ValueOf(q.Write),
			"WriteRange":   reflect.ValueOf(q.WriteRange),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"Enabled": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.Enabled))},
		},
	})
}
