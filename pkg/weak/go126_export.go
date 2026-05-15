// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package weak

import (
	_ "weak"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "weak",
		Path: "weak",
		Deps: map[string]string{
			"internal/abi": "abi",
			"runtime":      "runtime",
			"unsafe":       "unsafe",
		},
	})
}
