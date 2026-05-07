// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package context

import (
	q "context"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("context", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "context",
			Path: "context",
			Deps: map[string]string{
				"errors":               "errors",
				"internal/reflectlite": "reflectlite",
				"sync":                 "sync",
				"sync/atomic":          "atomic",
				"time":                 "time",
			},
			Interfaces: map[string]reflect.Type{
				"Context": reflect.TypeOf((*q.Context)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"CancelCauseFunc": reflect.TypeOf((*q.CancelCauseFunc)(nil)).Elem(),
				"CancelFunc":      reflect.TypeOf((*q.CancelFunc)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"Canceled":         &q.Canceled,
				"DeadlineExceeded": &q.DeadlineExceeded,
			},
			Funcs: map[string]interface{}{
				"AfterFunc":         q.AfterFunc,
				"Background":        q.Background,
				"Cause":             q.Cause,
				"TODO":              q.TODO,
				"WithCancel":        q.WithCancel,
				"WithCancelCause":   q.WithCancelCause,
				"WithDeadline":      q.WithDeadline,
				"WithDeadlineCause": q.WithDeadlineCause,
				"WithTimeout":       q.WithTimeout,
				"WithTimeoutCause":  q.WithTimeoutCause,
				"WithValue":         q.WithValue,
				"WithoutCancel":     q.WithoutCancel,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
