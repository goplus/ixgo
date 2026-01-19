// export by github.com/goplus/ixgo/cmd/qexp

package gsh

import (
	q "github.com/qiniu/x/gsh"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "gsh",
		Path: "github.com/qiniu/x/gsh",
		Deps: map[string]string{
			"bytes":   "bytes",
			"errors":  "errors",
			"io":      "io",
			"os":      "os",
			"os/exec": "exec",
			"strings": "strings",
		},
		Interfaces: map[string]reflect.Type{
			"OS": reflect.TypeOf((*q.OS)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"App": reflect.TypeOf((*q.App)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Sys": reflect.ValueOf(&q.Sys),
		},
		Funcs: map[string]reflect.Value{
			"Getenv":        reflect.ValueOf(q.Getenv),
			"InitApp":       reflect.ValueOf(q.InitApp),
			"Setenv__0":     reflect.ValueOf(q.Setenv__0),
			"Setenv__1":     reflect.ValueOf(q.Setenv__1),
			"Setenv__2":     reflect.ValueOf(q.Setenv__2),
			"XGot_App_Main": reflect.ValueOf(q.XGot_App_Main),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"XGoPackage": {"untyped bool", constant.MakeBool(bool(q.XGoPackage))},
		},
	})
}
