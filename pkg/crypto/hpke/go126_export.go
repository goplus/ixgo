// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package hpke

import (
	q "crypto/hpke"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto/hpke", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "hpke",
			Path: "crypto/hpke",
			Deps: map[string]string{
				"bytes":                           "bytes",
				"crypto":                          "crypto",
				"crypto/cipher":                   "cipher",
				"crypto/ecdh":                     "ecdh",
				"crypto/fips140":                  "fips140",
				"crypto/hkdf":                     "hkdf",
				"crypto/internal/fips140/aes":     "aes",
				"crypto/internal/fips140/aes/gcm": "gcm",
				"crypto/internal/fips140/drbg":    "drbg",
				"crypto/internal/rand":            "rand",
				"crypto/mlkem":                    "mlkem",
				"crypto/sha256":                   "sha256",
				"crypto/sha3":                     "sha3",
				"crypto/sha512":                   "sha512",
				"errors":                          "errors",
				"fmt":                             "fmt",
				"hash":                            "hash",
				"internal/byteorder":              "byteorder",
				"slices":                          "slices",
				"vendor/golang.org/x/crypto/chacha20poly1305": "chacha20poly1305",
			},
			Interfaces: map[string]reflect.Type{
				"AEAD":       reflect.TypeOf((*q.AEAD)(nil)).Elem(),
				"KDF":        reflect.TypeOf((*q.KDF)(nil)).Elem(),
				"KEM":        reflect.TypeOf((*q.KEM)(nil)).Elem(),
				"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
				"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Recipient": reflect.TypeOf((*q.Recipient)(nil)).Elem(),
				"Sender":    reflect.TypeOf((*q.Sender)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"AES128GCM":           q.AES128GCM,
				"AES256GCM":           q.AES256GCM,
				"ChaCha20Poly1305":    q.ChaCha20Poly1305,
				"DHKEM":               q.DHKEM,
				"ExportOnly":          q.ExportOnly,
				"HKDFSHA256":          q.HKDFSHA256,
				"HKDFSHA384":          q.HKDFSHA384,
				"HKDFSHA512":          q.HKDFSHA512,
				"MLKEM1024":           q.MLKEM1024,
				"MLKEM1024P384":       q.MLKEM1024P384,
				"MLKEM768":            q.MLKEM768,
				"MLKEM768P256":        q.MLKEM768P256,
				"MLKEM768X25519":      q.MLKEM768X25519,
				"NewAEAD":             q.NewAEAD,
				"NewDHKEMPrivateKey":  q.NewDHKEMPrivateKey,
				"NewDHKEMPublicKey":   q.NewDHKEMPublicKey,
				"NewHybridPrivateKey": q.NewHybridPrivateKey,
				"NewHybridPublicKey":  q.NewHybridPublicKey,
				"NewKDF":              q.NewKDF,
				"NewKEM":              q.NewKEM,
				"NewMLKEMPrivateKey":  q.NewMLKEMPrivateKey,
				"NewMLKEMPublicKey":   q.NewMLKEMPublicKey,
				"NewRecipient":        q.NewRecipient,
				"NewSender":           q.NewSender,
				"Open":                q.Open,
				"SHAKE128":            q.SHAKE128,
				"SHAKE256":            q.SHAKE256,
				"Seal":                q.Seal,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
