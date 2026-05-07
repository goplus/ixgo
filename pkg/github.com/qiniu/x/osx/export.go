// export by github.com/goplus/ixgo/cmd/qexp

package osx

import (
	q "github.com/qiniu/x/osx"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/qiniu/x/osx", func() *ixgo.Package {
		return &ixgo.Package{
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
				"BLineIter":   reflect.TypeOf((*q.BLineIter)(nil)).Elem(),
				"BLineReader": reflect.TypeOf((*q.BLineReader)(nil)).Elem(),
				"LineIter":    reflect.TypeOf((*q.LineIter)(nil)).Elem(),
				"LineReader":  reflect.TypeOf((*q.LineReader)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"BLines":     q.BLines,
				"Check":      q.Check,
				"EnumBLines": q.EnumBLines,
				"EnumLines":  q.EnumLines,
				"Errorln":    q.Errorln,
				"Fatal":      q.Fatal,
				"Lines":      q.Lines,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
