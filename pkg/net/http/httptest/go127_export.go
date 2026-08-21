// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package httptest

import (
	q "net/http/httptest"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("net/http/httptest", func() *ixgo.Package {
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
				"internal/nettest":                      "nettest",
				"io":                                    "io",
				"log":                                   "log",
				"net":                                   "net",
				"net/http":                              "http",
				"net/http/internal/testcert":            "testcert",
				"net/textproto":                         "textproto",
				"os":                                    "os",
				"runtime":                               "runtime",
				"strconv":                               "strconv",
				"strings":                               "strings",
				"sync":                                  "sync",
				"testing":                               "testing",
				"time":                                  "time",
				"unsafe":                                "unsafe",
				"vendor/golang.org/x/net/http/httpguts": "httpguts",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"ResponseRecorder": reflect.TypeOf((*q.ResponseRecorder)(nil)).Elem(),
				"Server":           reflect.TypeOf((*q.Server)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"NewRecorder":           reflect.ValueOf(q.NewRecorder),
				"NewRequest":            reflect.ValueOf(q.NewRequest),
				"NewRequestWithContext": reflect.ValueOf(q.NewRequestWithContext),
				"NewServer":             reflect.ValueOf(q.NewServer),
				"NewTLSServer":          reflect.ValueOf(q.NewTLSServer),
				"NewTestServer":         reflect.ValueOf(q.NewTestServer),
				"NewUnstartedServer":    reflect.ValueOf(q.NewUnstartedServer),
			},
			TypedConsts: map[string]ixgo.TypedConst{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"DefaultRemoteAddr": {Typ: "untyped string", Value: constant.MakeString(string(q.DefaultRemoteAddr))},
			},
		}
	})
}
