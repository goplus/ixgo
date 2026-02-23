// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package heap

import (
	q "container/heap"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "heap",
		Path: "container/heap",
		Deps: map[string]string{
			"sort": "sort",
		},
		Interfaces: map[string]reflect.Type{
			"Interface": reflect.TypeOf((*q.Interface)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Fix":    reflect.ValueOf(q.Fix),
			"Init":   reflect.ValueOf(q.Init),
			"Pop":    reflect.ValueOf(q.Pop),
			"Push":   reflect.ValueOf(q.Push),
			"Remove": reflect.ValueOf(q.Remove),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
