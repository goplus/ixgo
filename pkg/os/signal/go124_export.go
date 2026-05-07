// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package signal

import (
	q "os/signal"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("os/signal", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "signal",
			Path: "os/signal",
			Deps: map[string]string{
				"context": "context",
				"os":      "os",
				"slices":  "slices",
				"sync":    "sync",
				"syscall": "syscall",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Ignore":        q.Ignore,
				"Ignored":       q.Ignored,
				"Notify":        q.Notify,
				"NotifyContext": q.NotifyContext,
				"Reset":         q.Reset,
				"Stop":          q.Stop,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
