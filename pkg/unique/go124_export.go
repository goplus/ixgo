// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

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
			"internal/stringslite": "stringslite",
			"internal/sync":        "sync",
			"runtime":              "runtime",
			"sync":                 "sync",
			"unsafe":               "unsafe",
			"weak":                 "weak",
		},
	})
}
