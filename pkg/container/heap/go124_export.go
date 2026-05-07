// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package heap

import (
	q "container/heap"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("container/heap", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Fix":    q.Fix,
				"Init":   q.Init,
				"Pop":    q.Pop,
				"Push":   q.Push,
				"Remove": q.Remove,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
