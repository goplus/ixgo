// export by github.com/goplus/ixgo/cmd/qexp

package yaml

import (
	q "github.com/goplus/xgo/encoding/yaml"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/encoding/yaml", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "yaml",
			Path: "github.com/goplus/xgo/encoding/yaml",
			Deps: map[string]string{
				"github.com/goccy/go-yaml": "yaml",
				"strings":                  "strings",
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
