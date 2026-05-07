// export by github.com/goplus/ixgo/cmd/qexp

package scannertest

import (
	q "github.com/goplus/xgo/tpl/scanner/scannertest"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/tpl/scanner/scannertest", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "scannertest",
			Path: "github.com/goplus/xgo/tpl/scanner/scannertest",
			Deps: map[string]string{
				"fmt":                               "fmt",
				"github.com/goplus/xgo/scanner":     "scanner",
				"github.com/goplus/xgo/token":       "token",
				"github.com/goplus/xgo/tpl/scanner": "scanner",
				"github.com/goplus/xgo/tpl/token":   "token",
				"go/scanner":                        "scanner",
				"go/token":                          "token",
				"io":                                "io",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"GoScan":  q.GoScan,
				"GopScan": q.GopScan,
				"Scan":    q.Scan,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
