// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package cipher

import (
	q "crypto/cipher"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/cipher", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "cipher",
			Path: "crypto/cipher",
			Deps: map[string]string{
				"bytes":                           "bytes",
				"crypto/internal/fips140/aes":     "aes",
				"crypto/internal/fips140/aes/gcm": "gcm",
				"crypto/internal/fips140/alias":   "alias",
				"crypto/internal/fips140only":     "fips140only",
				"crypto/subtle":                   "subtle",
				"errors":                          "errors",
				"internal/byteorder":              "byteorder",
				"io":                              "io",
			},
			Interfaces: map[string]reflect.Type{
				"AEAD":      reflect.TypeOf((*q.AEAD)(nil)).Elem(),
				"Block":     reflect.TypeOf((*q.Block)(nil)).Elem(),
				"BlockMode": reflect.TypeOf((*q.BlockMode)(nil)).Elem(),
				"Stream":    reflect.TypeOf((*q.Stream)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"StreamReader": reflect.TypeOf((*q.StreamReader)(nil)).Elem(),
				"StreamWriter": reflect.TypeOf((*q.StreamWriter)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"NewCBCDecrypter":       q.NewCBCDecrypter,
				"NewCBCEncrypter":       q.NewCBCEncrypter,
				"NewCFBDecrypter":       q.NewCFBDecrypter,
				"NewCFBEncrypter":       q.NewCFBEncrypter,
				"NewCTR":                q.NewCTR,
				"NewGCM":                q.NewGCM,
				"NewGCMWithNonceSize":   q.NewGCMWithNonceSize,
				"NewGCMWithRandomNonce": q.NewGCMWithRandomNonce,
				"NewGCMWithTagSize":     q.NewGCMWithTagSize,
				"NewOFB":                q.NewOFB,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
