// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package pbkdf2

import (
	_ "crypto/pbkdf2"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "pbkdf2",
		Path: "crypto/pbkdf2",
		Deps: map[string]string{
			"crypto/internal/fips140/pbkdf2": "pbkdf2",
			"crypto/internal/fips140hash":    "fips140hash",
			"crypto/internal/fips140only":    "fips140only",
			"errors":                         "errors",
			"hash":                           "hash",
		},
	})
}
