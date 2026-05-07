// export by github.com/goplus/ixgo/cmd/qexp

package fs

import (
	q "github.com/goplus/xgo/encoding/fs"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/encoding/fs", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "fs",
			Path: "github.com/goplus/xgo/encoding/fs",
			Deps: map[string]string{
				"github.com/goplus/xgo/dql/fs": "fs",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New": q.New,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
