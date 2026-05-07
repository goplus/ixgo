// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package gob

import (
	q "encoding/gob"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("encoding/gob", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "gob",
			Path: "encoding/gob",
			Deps: map[string]string{
				"bufio":            "bufio",
				"encoding":         "encoding",
				"encoding/binary":  "binary",
				"errors":           "errors",
				"fmt":              "fmt",
				"internal/saferio": "saferio",
				"io":               "io",
				"maps":             "maps",
				"math":             "math",
				"math/bits":        "bits",
				"os":               "os",
				"reflect":          "reflect",
				"sync":             "sync",
				"sync/atomic":      "atomic",
				"unicode":          "unicode",
				"unicode/utf8":     "utf8",
			},
			Interfaces: map[string]reflect.Type{
				"GobDecoder": reflect.TypeOf((*q.GobDecoder)(nil)).Elem(),
				"GobEncoder": reflect.TypeOf((*q.GobEncoder)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"CommonType": reflect.TypeOf((*q.CommonType)(nil)).Elem(),
				"Decoder":    reflect.TypeOf((*q.Decoder)(nil)).Elem(),
				"Encoder":    reflect.TypeOf((*q.Encoder)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewDecoder":   q.NewDecoder,
				"NewEncoder":   q.NewEncoder,
				"Register":     q.Register,
				"RegisterName": q.RegisterName,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
