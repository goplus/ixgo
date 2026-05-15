// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package iter

import (
	_ "iter"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "iter",
		Path: "iter",
		Deps: map[string]string{
			"internal/race": "race",
			"runtime":       "runtime",
			"unsafe":        "unsafe",
		},
	})
}
