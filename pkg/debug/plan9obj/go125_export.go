// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package plan9obj

import (
	q "debug/plan9obj"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "plan9obj",
		Path: "debug/plan9obj",
		Deps: map[string]string{
			"encoding/binary":  "binary",
			"errors":           "errors",
			"fmt":              "fmt",
			"internal/saferio": "saferio",
			"io":               "io",
			"os":               "os",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"File":          reflect.TypeOf((*q.File)(nil)).Elem(),
			"FileHeader":    reflect.TypeOf((*q.FileHeader)(nil)).Elem(),
			"Section":       reflect.TypeOf((*q.Section)(nil)).Elem(),
			"SectionHeader": reflect.TypeOf((*q.SectionHeader)(nil)).Elem(),
			"Sym":           reflect.TypeOf((*q.Sym)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrNoSymbols": reflect.ValueOf(&q.ErrNoSymbols),
		},
		Funcs: map[string]reflect.Value{
			"NewFile": reflect.ValueOf(q.NewFile),
			"Open":    reflect.ValueOf(q.Open),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"Magic386":   {"untyped int", constant.MakeInt64(int64(q.Magic386))},
			"Magic64":    {"untyped int", constant.MakeInt64(int64(q.Magic64))},
			"MagicAMD64": {"untyped int", constant.MakeInt64(int64(q.MagicAMD64))},
			"MagicARM":   {"untyped int", constant.MakeInt64(int64(q.MagicARM))},
		},
	})
}
