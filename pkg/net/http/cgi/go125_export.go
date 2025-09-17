// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package cgi

import (
	q "net/http/cgi"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "cgi",
		Path: "net/http/cgi",
		Deps: map[string]string{
			"bufio":                                 "bufio",
			"crypto/tls":                            "tls",
			"errors":                                "errors",
			"fmt":                                   "fmt",
			"io":                                    "io",
			"log":                                   "log",
			"maps":                                  "maps",
			"net":                                   "net",
			"net/http":                              "http",
			"net/textproto":                         "textproto",
			"net/url":                               "url",
			"os":                                    "os",
			"os/exec":                               "exec",
			"path":                                  "path",
			"path/filepath":                         "filepath",
			"regexp":                                "regexp",
			"runtime":                               "runtime",
			"slices":                                "slices",
			"strconv":                               "strconv",
			"strings":                               "strings",
			"time":                                  "time",
			"vendor/golang.org/x/net/http/httpguts": "httpguts",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Handler": reflect.TypeOf((*q.Handler)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Request":        reflect.ValueOf(q.Request),
			"RequestFromMap": reflect.ValueOf(q.RequestFromMap),
			"Serve":          reflect.ValueOf(q.Serve),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
