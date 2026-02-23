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
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AES128GCM":           reflect.ValueOf(q.AES128GCM),
			"AES256GCM":           reflect.ValueOf(q.AES256GCM),
			"ChaCha20Poly1305":    reflect.ValueOf(q.ChaCha20Poly1305),
			"DHKEM":               reflect.ValueOf(q.DHKEM),
			"ExportOnly":          reflect.ValueOf(q.ExportOnly),
			"HKDFSHA256":          reflect.ValueOf(q.HKDFSHA256),
			"HKDFSHA384":          reflect.ValueOf(q.HKDFSHA384),
			"HKDFSHA512":          reflect.ValueOf(q.HKDFSHA512),
			"MLKEM1024":           reflect.ValueOf(q.MLKEM1024),
			"MLKEM1024P384":       reflect.ValueOf(q.MLKEM1024P384),
			"MLKEM768":            reflect.ValueOf(q.MLKEM768),
			"MLKEM768P256":        reflect.ValueOf(q.MLKEM768P256),
			"MLKEM768X25519":      reflect.ValueOf(q.MLKEM768X25519),
			"NewAEAD":             reflect.ValueOf(q.NewAEAD),
			"NewDHKEMPrivateKey":  reflect.ValueOf(q.NewDHKEMPrivateKey),
			"NewDHKEMPublicKey":   reflect.ValueOf(q.NewDHKEMPublicKey),
			"NewHybridPrivateKey": reflect.ValueOf(q.NewHybridPrivateKey),
			"NewHybridPublicKey":  reflect.ValueOf(q.NewHybridPublicKey),
			"NewKDF":              reflect.ValueOf(q.NewKDF),
			"NewKEM":              reflect.ValueOf(q.NewKEM),
			"NewMLKEMPrivateKey":  reflect.ValueOf(q.NewMLKEMPrivateKey),
			"NewMLKEMPublicKey":   reflect.ValueOf(q.NewMLKEMPublicKey),
			"NewRecipient":        reflect.ValueOf(q.NewRecipient),
			"NewSender":           reflect.ValueOf(q.NewSender),
			"Open":                reflect.ValueOf(q.Open),
			"SHAKE128":            reflect.ValueOf(q.SHAKE128),
			"SHAKE256":            reflect.ValueOf(q.SHAKE256),
			"Seal":                reflect.ValueOf(q.Seal),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
