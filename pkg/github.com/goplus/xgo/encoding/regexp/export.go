// export by github.com/goplus/ixgo/cmd/qexp

package regexp

import (
	q "github.com/goplus/xgo/encoding/regexp"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/encoding/regexp", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"New": q.New,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
