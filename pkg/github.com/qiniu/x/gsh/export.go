// export by github.com/goplus/ixgo/cmd/qexp

package gsh

import (
	q "github.com/qiniu/x/gsh"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/qiniu/x/gsh", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars: map[string]interface{}{
				"Sys": &q.Sys,
			},
			Funcs: map[string]interface{}{
				"Getenv":        q.Getenv,
				"InitApp":       q.InitApp,
				"Setenv__0":     q.Setenv__0,
				"Setenv__1":     q.Setenv__1,
				"Setenv__2":     q.Setenv__2,
				"XGot_App_Main": q.XGot_App_Main,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"XGoPackage": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.XGoPackage))},
			},
		}
	})
}
