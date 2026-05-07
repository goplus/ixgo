// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package build

import (
	q "go/build"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("go/build", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "build",
			Path: "go/build",
			Deps: map[string]string{
				"bufio":               "bufio",
				"bytes":               "bytes",
				"errors":              "errors",
				"fmt":                 "fmt",
				"go/ast":              "ast",
				"go/build/constraint": "constraint",
				"go/doc":              "doc",
				"go/parser":           "parser",
				"go/scanner":          "scanner",
				"go/token":            "token",
				"internal/buildcfg":   "buildcfg",
				"internal/godebug":    "godebug",
				"internal/goroot":     "goroot",
				"internal/goversion":  "goversion",
				"internal/platform":   "platform",
				"internal/syslist":    "syslist",
				"io":                  "io",
				"io/fs":               "fs",
				"os":                  "os",
				"os/exec":             "exec",
				"path":                "path",
				"path/filepath":       "filepath",
				"runtime":             "runtime",
				"slices":              "slices",
				"strconv":             "strconv",
				"strings":             "strings",
				"unicode":             "unicode",
				"unicode/utf8":        "utf8",
				"unsafe":              "unsafe",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Context":              reflect.TypeOf((*q.Context)(nil)).Elem(),
				"Directive":            reflect.TypeOf((*q.Directive)(nil)).Elem(),
				"ImportMode":           reflect.TypeOf((*q.ImportMode)(nil)).Elem(),
				"MultiplePackageError": reflect.TypeOf((*q.MultiplePackageError)(nil)).Elem(),
				"NoGoError":            reflect.TypeOf((*q.NoGoError)(nil)).Elem(),
				"Package":              reflect.TypeOf((*q.Package)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"Default": &q.Default,
				"ToolDir": &q.ToolDir,
			},
			Funcs: map[string]interface{}{
				"ArchChar":      q.ArchChar,
				"Import":        q.Import,
				"ImportDir":     q.ImportDir,
				"IsLocalImport": q.IsLocalImport,
			},
			TypedConsts: map[string]interface{}{
				"AllowBinary":   q.AllowBinary,
				"FindOnly":      q.FindOnly,
				"IgnoreVendor":  q.IgnoreVendor,
				"ImportComment": q.ImportComment,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
