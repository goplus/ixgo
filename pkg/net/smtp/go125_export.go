// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package smtp

import (
	q "net/smtp"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "smtp",
		Path: "net/smtp",
		Deps: map[string]string{
			"crypto/hmac":     "hmac",
			"crypto/md5":      "md5",
			"crypto/tls":      "tls",
			"encoding/base64": "base64",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"net":             "net",
			"net/textproto":   "textproto",
			"strings":         "strings",
		},
		Interfaces: map[string]reflect.Type{
			"Auth": reflect.TypeOf((*q.Auth)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Client":     reflect.TypeOf((*q.Client)(nil)).Elem(),
			"ServerInfo": reflect.TypeOf((*q.ServerInfo)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"CRAMMD5Auth": reflect.ValueOf(q.CRAMMD5Auth),
			"Dial":        reflect.ValueOf(q.Dial),
			"NewClient":   reflect.ValueOf(q.NewClient),
			"PlainAuth":   reflect.ValueOf(q.PlainAuth),
			"SendMail":    reflect.ValueOf(q.SendMail),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
