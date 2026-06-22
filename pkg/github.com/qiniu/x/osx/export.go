// export by github.com/goplus/ixgo/cmd/qexp

package osx

import (
	q "github.com/qiniu/x/osx"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "osx",
		Path: "github.com/qiniu/x/osx",
		Deps: map[string]string{
			"bufio": "bufio",
			"fmt":   "fmt",
			"io":    "io",
			"os":    "os",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"BLineIter": reflect.TypeOf((*q.BLineIter)(nil)).Elem(),
			"LineIter":  reflect.TypeOf((*q.LineIter)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"BLines":     reflect.ValueOf(q.BLines),
			"Check":      reflect.ValueOf(q.Check),
			"EnumBLines": reflect.ValueOf(q.EnumBLines),
			"EnumLines":  reflect.ValueOf(q.EnumLines),
			"Errorln":    reflect.ValueOf(q.Errorln),
			"Fatal":      reflect.ValueOf(q.Fatal),
			"Lines":      reflect.ValueOf(q.Lines),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
