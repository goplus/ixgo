// export by github.com/goplus/ixgo/cmd/qexp

package fetcher

import (
	q "github.com/goplus/xgo/dql/fetcher"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"ErrUnknownPageType": reflect.ValueOf(&q.ErrUnknownPageType),
		},
		Funcs: map[string]reflect.Value{
			"Do":       reflect.ValueOf(q.Do),
			"From":     reflect.ValueOf(q.From),
			"List":     reflect.ValueOf(q.List),
			"Register": reflect.ValueOf(q.Register),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
