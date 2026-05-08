// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26 && js
// +build go1.26,js

package js

import (
	q "syscall/js"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "js",
		Path: "syscall/js",
		Deps: map[string]string{
			"internal/synctest": "synctest",
			"runtime":           "runtime",
			"sync":              "sync",
			"unsafe":            "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Error":      reflect.TypeOf((*q.Error)(nil)).Elem(),
			"Func":       reflect.TypeOf((*q.Func)(nil)).Elem(),
			"Type":       reflect.TypeOf((*q.Type)(nil)).Elem(),
			"Value":      reflect.TypeOf((*q.Value)(nil)).Elem(),
			"ValueError": reflect.TypeOf((*q.ValueError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"CopyBytesToGo": reflect.ValueOf(q.CopyBytesToGo),
			"CopyBytesToJS": reflect.ValueOf(q.CopyBytesToJS),
			"FuncOf":        reflect.ValueOf(q.FuncOf),
			"Global":        reflect.ValueOf(q.Global),
			"Null":          reflect.ValueOf(q.Null),
			"Undefined":     reflect.ValueOf(q.Undefined),
			"ValueOf":       reflect.ValueOf(q.ValueOf),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"TypeBoolean":   {Typ: reflect.TypeOf(q.TypeBoolean), Value: constant.MakeInt64(int64(q.TypeBoolean))},
			"TypeFunction":  {Typ: reflect.TypeOf(q.TypeFunction), Value: constant.MakeInt64(int64(q.TypeFunction))},
			"TypeNull":      {Typ: reflect.TypeOf(q.TypeNull), Value: constant.MakeInt64(int64(q.TypeNull))},
			"TypeNumber":    {Typ: reflect.TypeOf(q.TypeNumber), Value: constant.MakeInt64(int64(q.TypeNumber))},
			"TypeObject":    {Typ: reflect.TypeOf(q.TypeObject), Value: constant.MakeInt64(int64(q.TypeObject))},
			"TypeString":    {Typ: reflect.TypeOf(q.TypeString), Value: constant.MakeInt64(int64(q.TypeString))},
			"TypeSymbol":    {Typ: reflect.TypeOf(q.TypeSymbol), Value: constant.MakeInt64(int64(q.TypeSymbol))},
			"TypeUndefined": {Typ: reflect.TypeOf(q.TypeUndefined), Value: constant.MakeInt64(int64(q.TypeUndefined))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
