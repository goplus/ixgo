// export by github.com/goplus/ixgo/cmd/qexp

package env

import (
	q "github.com/goplus/xgo/env"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"BuildDate":   reflect.ValueOf(q.BuildDate),
			"HOME":        reflect.ValueOf(q.HOME),
			"Installed":   reflect.ValueOf(q.Installed),
			"MainVersion": reflect.ValueOf(q.MainVersion),
			"Version":     reflect.ValueOf(q.Version),
			"XGOROOT":     reflect.ValueOf(q.XGOROOT),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
