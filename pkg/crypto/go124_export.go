// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package crypto

import (
	q "crypto"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("crypto", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "crypto",
			Path: "crypto",
			Deps: map[string]string{
				"hash":    "hash",
				"io":      "io",
				"strconv": "strconv",
			},
			Interfaces: map[string]reflect.Type{
				"Decrypter":     reflect.TypeOf((*q.Decrypter)(nil)).Elem(),
				"DecrypterOpts": reflect.TypeOf((*q.DecrypterOpts)(nil)).Elem(),
				"PrivateKey":    reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
				"PublicKey":     reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
				"Signer":        reflect.TypeOf((*q.Signer)(nil)).Elem(),
				"SignerOpts":    reflect.TypeOf((*q.SignerOpts)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Hash": reflect.TypeOf((*q.Hash)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"RegisterHash": q.RegisterHash,
			},
			TypedConsts: map[string]interface{}{
				"BLAKE2b_256": q.BLAKE2b_256,
				"BLAKE2b_384": q.BLAKE2b_384,
				"BLAKE2b_512": q.BLAKE2b_512,
				"BLAKE2s_256": q.BLAKE2s_256,
				"MD4":         q.MD4,
				"MD5":         q.MD5,
				"MD5SHA1":     q.MD5SHA1,
				"RIPEMD160":   q.RIPEMD160,
				"SHA1":        q.SHA1,
				"SHA224":      q.SHA224,
				"SHA256":      q.SHA256,
				"SHA384":      q.SHA384,
				"SHA3_224":    q.SHA3_224,
				"SHA3_256":    q.SHA3_256,
				"SHA3_384":    q.SHA3_384,
				"SHA3_512":    q.SHA3_512,
				"SHA512":      q.SHA512,
				"SHA512_224":  q.SHA512_224,
				"SHA512_256":  q.SHA512_256,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
