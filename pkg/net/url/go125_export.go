// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package url

import (
	q "net/url"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("net/url", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "url",
			Path: "net/url",
			Deps: map[string]string{
				"errors":  "errors",
				"fmt":     "fmt",
				"maps":    "maps",
				"path":    "path",
				"slices":  "slices",
				"strconv": "strconv",
				"strings": "strings",
				"unsafe":  "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Error":            reflect.TypeOf((*q.Error)(nil)).Elem(),
				"EscapeError":      reflect.TypeOf((*q.EscapeError)(nil)).Elem(),
				"InvalidHostError": reflect.TypeOf((*q.InvalidHostError)(nil)).Elem(),
				"URL":              reflect.TypeOf((*q.URL)(nil)).Elem(),
				"Userinfo":         reflect.TypeOf((*q.Userinfo)(nil)).Elem(),
				"Values":           reflect.TypeOf((*q.Values)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"JoinPath":        q.JoinPath,
				"Parse":           q.Parse,
				"ParseQuery":      q.ParseQuery,
				"ParseRequestURI": q.ParseRequestURI,
				"PathEscape":      q.PathEscape,
				"PathUnescape":    q.PathUnescape,
				"QueryEscape":     q.QueryEscape,
				"QueryUnescape":   q.QueryUnescape,
				"User":            q.User,
				"UserPassword":    q.UserPassword,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
