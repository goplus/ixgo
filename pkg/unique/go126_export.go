// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package unique

import (
	_ "unique"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "unique",
		Path: "unique",
		Deps: map[string]string{
			"internal/abi":         "abi",
			"internal/goarch":      "goarch",
			"internal/stringslite": "stringslite",
			"internal/sync":        "sync",
			"runtime":              "runtime",
			"sync":                 "sync",
			"sync/atomic":          "atomic",
			"unsafe":               "unsafe",
			"weak":                 "weak",
		},
	})
}
