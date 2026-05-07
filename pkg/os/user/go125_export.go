// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package user

import (
	q "os/user"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("os/user", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "user",
			Path: "os/user",
			Deps: map[string]string{
				"fmt":                   "fmt",
				"internal/syscall/unix": "unix",
				"runtime":               "runtime",
				"strconv":               "strconv",
				"strings":               "strings",
				"sync":                  "sync",
				"syscall":               "syscall",
				"unsafe":                "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Group":               reflect.TypeOf((*q.Group)(nil)).Elem(),
				"UnknownGroupError":   reflect.TypeOf((*q.UnknownGroupError)(nil)).Elem(),
				"UnknownGroupIdError": reflect.TypeOf((*q.UnknownGroupIdError)(nil)).Elem(),
				"UnknownUserError":    reflect.TypeOf((*q.UnknownUserError)(nil)).Elem(),
				"UnknownUserIdError":  reflect.TypeOf((*q.UnknownUserIdError)(nil)).Elem(),
				"User":                reflect.TypeOf((*q.User)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Current":       q.Current,
				"Lookup":        q.Lookup,
				"LookupGroup":   q.LookupGroup,
				"LookupGroupId": q.LookupGroupId,
				"LookupId":      q.LookupId,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
