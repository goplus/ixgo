// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package quick

import (
	q "testing/quick"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "quick",
		Path: "testing/quick",
		Deps: map[string]string{
			"flag":      "flag",
			"fmt":       "fmt",
			"math":      "math",
			"math/rand": "rand",
			"reflect":   "reflect",
			"strings":   "strings",
			"time":      "time",
		},
		Interfaces: map[string]reflect.Type{
			"Generator": reflect.TypeOf((*q.Generator)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"CheckEqualError": reflect.TypeOf((*q.CheckEqualError)(nil)).Elem(),
			"CheckError":      reflect.TypeOf((*q.CheckError)(nil)).Elem(),
			"Config":          reflect.TypeOf((*q.Config)(nil)).Elem(),
			"SetupError":      reflect.TypeOf((*q.SetupError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Check":      reflect.ValueOf(q.Check),
			"CheckEqual": reflect.ValueOf(q.CheckEqual),
			"Value":      reflect.ValueOf(q.Value),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
