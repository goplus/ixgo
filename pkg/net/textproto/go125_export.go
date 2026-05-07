// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package textproto

import (
	q "net/textproto"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("net/textproto", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "textproto",
			Path: "net/textproto",
			Deps: map[string]string{
				"bufio":   "bufio",
				"bytes":   "bytes",
				"errors":  "errors",
				"fmt":     "fmt",
				"io":      "io",
				"math":    "math",
				"net":     "net",
				"strconv": "strconv",
				"strings": "strings",
				"sync":    "sync",
				"unsafe":  "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Conn":          reflect.TypeOf((*q.Conn)(nil)).Elem(),
				"Error":         reflect.TypeOf((*q.Error)(nil)).Elem(),
				"MIMEHeader":    reflect.TypeOf((*q.MIMEHeader)(nil)).Elem(),
				"Pipeline":      reflect.TypeOf((*q.Pipeline)(nil)).Elem(),
				"ProtocolError": reflect.TypeOf((*q.ProtocolError)(nil)).Elem(),
				"Reader":        reflect.TypeOf((*q.Reader)(nil)).Elem(),
				"Writer":        reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"CanonicalMIMEHeaderKey": q.CanonicalMIMEHeaderKey,
				"Dial":                   q.Dial,
				"NewConn":                q.NewConn,
				"NewReader":              q.NewReader,
				"NewWriter":              q.NewWriter,
				"TrimBytes":              q.TrimBytes,
				"TrimString":             q.TrimString,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
