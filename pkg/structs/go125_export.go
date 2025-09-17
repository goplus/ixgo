// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package structs

import (
	q "structs"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name:       "structs",
		Path:       "structs",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"HostLayout": reflect.TypeOf((*q.HostLayout)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
