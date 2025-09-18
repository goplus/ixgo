// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package context

import (
	q "context"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"Canceled":         reflect.ValueOf(&q.Canceled),
			"DeadlineExceeded": reflect.ValueOf(&q.DeadlineExceeded),
		},
		Funcs: map[string]reflect.Value{
			"AfterFunc":         reflect.ValueOf(q.AfterFunc),
			"Background":        reflect.ValueOf(q.Background),
			"Cause":             reflect.ValueOf(q.Cause),
			"TODO":              reflect.ValueOf(q.TODO),
			"WithCancel":        reflect.ValueOf(q.WithCancel),
			"WithCancelCause":   reflect.ValueOf(q.WithCancelCause),
			"WithDeadline":      reflect.ValueOf(q.WithDeadline),
			"WithDeadlineCause": reflect.ValueOf(q.WithDeadlineCause),
			"WithTimeout":       reflect.ValueOf(q.WithTimeout),
			"WithTimeoutCause":  reflect.ValueOf(q.WithTimeoutCause),
			"WithValue":         reflect.ValueOf(q.WithValue),
			"WithoutCancel":     reflect.ValueOf(q.WithoutCancel),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
