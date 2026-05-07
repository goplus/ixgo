// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package httptest

import (
	q "net/http/httptest"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("net/http/httptest", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "httptest",
			Path: "net/http/httptest",
			Deps: map[string]string{
				"bufio":                                 "bufio",
				"bytes":                                 "bytes",
				"context":                               "context",
				"crypto/tls":                            "tls",
				"crypto/x509":                           "x509",
				"flag":                                  "flag",
				"fmt":                                   "fmt",
				"io":                                    "io",
				"log":                                   "log",
				"net":                                   "net",
				"net/http":                              "http",
				"net/http/internal/testcert":            "testcert",
				"net/textproto":                         "textproto",
				"os":                                    "os",
				"strconv":                               "strconv",
				"strings":                               "strings",
				"sync":                                  "sync",
				"time":                                  "time",
				"vendor/golang.org/x/net/http/httpguts": "httpguts",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"ResponseRecorder": reflect.TypeOf((*q.ResponseRecorder)(nil)).Elem(),
				"Server":           reflect.TypeOf((*q.Server)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewRecorder":           q.NewRecorder,
				"NewRequest":            q.NewRequest,
				"NewRequestWithContext": q.NewRequestWithContext,
				"NewServer":             q.NewServer,
				"NewTLSServer":          q.NewTLSServer,
				"NewUnstartedServer":    q.NewUnstartedServer,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"DefaultRemoteAddr": {Typ: "untyped string", Value: constant.MakeString(string(q.DefaultRemoteAddr))},
			},
		}
	})
}
