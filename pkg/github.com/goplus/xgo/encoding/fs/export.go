// export by github.com/goplus/ixgo/cmd/qexp

package fs

import (
	q "github.com/goplus/xgo/encoding/fs"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "fs",
		Path: "github.com/goplus/xgo/encoding/fs",
		Deps: map[string]string{
			"github.com/goplus/xgo/dql/fs": "fs",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
