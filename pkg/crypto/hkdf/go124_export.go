// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package hkdf

import (
	_ "crypto/hkdf"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "hkdf",
		Path: "crypto/hkdf",
		Deps: map[string]string{
			"crypto/internal/fips140/hkdf": "hkdf",
			"crypto/internal/fips140hash":  "fips140hash",
			"crypto/internal/fips140only":  "fips140only",
			"errors":                       "errors",
			"hash":                         "hash",
		},
	})
}
