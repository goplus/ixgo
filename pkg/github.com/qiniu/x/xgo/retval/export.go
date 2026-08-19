// export by github.com/goplus/ixgo/cmd/qexp

package retval

import (
	q "github.com/qiniu/x/xgo/retval"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name:       "retval",
		Path:       "github.com/qiniu/x/xgo/retval",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Get": reflect.ValueOf(q.Get),
			"Set": reflect.ValueOf(q.Set),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
