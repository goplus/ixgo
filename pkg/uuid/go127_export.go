// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package uuid

import (
	q "uuid"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "uuid",
		Path: "uuid",
		Deps: map[string]string{
			"bytes":           "bytes",
			"cmp":             "cmp",
			"crypto/rand":     "rand",
			"encoding/binary": "binary",
			"encoding/hex":    "hex",
			"errors":          "errors",
			"sync":            "sync",
			"time":            "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"UUID": reflect.TypeOf((*q.UUID)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Max":       reflect.ValueOf(q.Max),
			"MustParse": reflect.ValueOf(q.MustParse),
			"New":       reflect.ValueOf(q.New),
			"NewV4":     reflect.ValueOf(q.NewV4),
			"NewV7":     reflect.ValueOf(q.NewV7),
			"Nil":       reflect.ValueOf(q.Nil),
			"Parse":     reflect.ValueOf(q.Parse),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
