// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package quick

import (
	q "testing/quick"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("testing/quick", func() *ixgo.Package {
		return &ixgo.Package{
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
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Check":      q.Check,
				"CheckEqual": q.CheckEqual,
				"Value":      q.Value,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
