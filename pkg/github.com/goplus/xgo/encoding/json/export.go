// export by github.com/goplus/ixgo/cmd/qexp

package json

import (
	q "github.com/goplus/xgo/encoding/json"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/encoding/json", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "json",
			Path: "github.com/goplus/xgo/encoding/json",
			Deps: map[string]string{
				"encoding/json": "json",
				"strings":       "strings",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{
				"Object": reflect.TypeOf((*q.Object)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New": q.New,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
