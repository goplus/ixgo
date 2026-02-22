// export by github.com/goplus/ixgo/cmd/qexp

package regexp

import (
	q "github.com/goplus/xgo/encoding/regexp"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "regexp",
		Path: "github.com/goplus/xgo/encoding/regexp",
		Deps: map[string]string{
			"regexp": "regexp",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"Object": reflect.TypeOf((*q.Object)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
