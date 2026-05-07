// export by github.com/goplus/ixgo/cmd/qexp

package env

import (
	q "github.com/goplus/xgo/env"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/xgo/env", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "env",
			Path: "github.com/goplus/xgo/env",
			Deps: map[string]string{
				"bytes":         "bytes",
				"log":           "log",
				"os":            "os",
				"os/exec":       "exec",
				"path/filepath": "filepath",
				"runtime/debug": "debug",
				"strings":       "strings",
				"syscall":       "syscall",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"BuildDate":   q.BuildDate,
				"HOME":        q.HOME,
				"Installed":   q.Installed,
				"MainVersion": q.MainVersion,
				"Version":     q.Version,
				"XGOROOT":     q.XGOROOT,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
