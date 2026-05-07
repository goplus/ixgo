// export by github.com/goplus/ixgo/cmd/qexp

package fetcher

import (
	q "github.com/goplus/xgo/dql/fetcher"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/dql/fetcher", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "fetcher",
			Path: "github.com/goplus/xgo/dql/fetcher",
			Deps: map[string]string{
				"errors":                         "errors",
				"github.com/goplus/xgo/dql/html": "html",
				"reflect":                        "reflect",
				"sort":                           "sort",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{
				"Conv": reflect.TypeOf((*q.Conv)(nil)).Elem(),
			},
			Vars: map[string]interface{}{
				"ErrUnknownFetchType": &q.ErrUnknownFetchType,
			},
			Funcs: map[string]interface{}{
				"Do":       q.Do,
				"From":     q.From,
				"List":     q.List,
				"Register": q.Register,
				"URL":      q.URL,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
