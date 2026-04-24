// export by github.com/goplus/ixgo/cmd/qexp

package msg

import (
	q "github.com/goplus/ixgo/testdata/alias/msg"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name:       "msg",
		Path:       "github.com/goplus/ixgo/testdata/alias/msg",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"Message": reflect.TypeOf((*q.Message)(nil)).Elem(),
		},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
