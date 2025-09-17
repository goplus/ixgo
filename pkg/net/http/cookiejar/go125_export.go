// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package cookiejar

import (
	q "net/http/cookiejar"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "cookiejar",
		Path: "net/http/cookiejar",
		Deps: map[string]string{
			"cmp":                     "cmp",
			"errors":                  "errors",
			"fmt":                     "fmt",
			"net":                     "net",
			"net/http":                "http",
			"net/http/internal/ascii": "ascii",
			"net/url":                 "url",
			"slices":                  "slices",
			"strings":                 "strings",
			"sync":                    "sync",
			"time":                    "time",
			"unicode/utf8":            "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"PublicSuffixList": reflect.TypeOf((*q.PublicSuffixList)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Jar":     reflect.TypeOf((*q.Jar)(nil)).Elem(),
			"Options": reflect.TypeOf((*q.Options)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
