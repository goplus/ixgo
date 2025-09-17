// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package mail

import (
	q "net/mail"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "mail",
		Path: "net/mail",
		Deps: map[string]string{
			"bufio":         "bufio",
			"errors":        "errors",
			"fmt":           "fmt",
			"io":            "io",
			"log":           "log",
			"mime":          "mime",
			"net":           "net",
			"net/textproto": "textproto",
			"strings":       "strings",
			"sync":          "sync",
			"time":          "time",
			"unicode/utf8":  "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Address":       reflect.TypeOf((*q.Address)(nil)).Elem(),
			"AddressParser": reflect.TypeOf((*q.AddressParser)(nil)).Elem(),
			"Header":        reflect.TypeOf((*q.Header)(nil)).Elem(),
			"Message":       reflect.TypeOf((*q.Message)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrHeaderNotPresent": reflect.ValueOf(&q.ErrHeaderNotPresent),
		},
		Funcs: map[string]reflect.Value{
			"ParseAddress":     reflect.ValueOf(q.ParseAddress),
			"ParseAddressList": reflect.ValueOf(q.ParseAddressList),
			"ParseDate":        reflect.ValueOf(q.ParseDate),
			"ReadMessage":      reflect.ValueOf(q.ReadMessage),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
