// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package crypto

import (
	q "crypto"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("crypto", func() *ixgo.Package {
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
			Vars:       map[string]reflect.Value{},
			Funcs: map[string]reflect.Value{
				"RegisterHash": reflect.ValueOf(q.RegisterHash),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"BLAKE2b_256": {Typ: reflect.TypeOf(q.BLAKE2b_256), Value: constant.MakeInt64(int64(q.BLAKE2b_256))},
				"BLAKE2b_384": {Typ: reflect.TypeOf(q.BLAKE2b_384), Value: constant.MakeInt64(int64(q.BLAKE2b_384))},
				"BLAKE2b_512": {Typ: reflect.TypeOf(q.BLAKE2b_512), Value: constant.MakeInt64(int64(q.BLAKE2b_512))},
				"BLAKE2s_256": {Typ: reflect.TypeOf(q.BLAKE2s_256), Value: constant.MakeInt64(int64(q.BLAKE2s_256))},
				"MD4":         {Typ: reflect.TypeOf(q.MD4), Value: constant.MakeInt64(int64(q.MD4))},
				"MD5":         {Typ: reflect.TypeOf(q.MD5), Value: constant.MakeInt64(int64(q.MD5))},
				"MD5SHA1":     {Typ: reflect.TypeOf(q.MD5SHA1), Value: constant.MakeInt64(int64(q.MD5SHA1))},
				"RIPEMD160":   {Typ: reflect.TypeOf(q.RIPEMD160), Value: constant.MakeInt64(int64(q.RIPEMD160))},
				"SHA1":        {Typ: reflect.TypeOf(q.SHA1), Value: constant.MakeInt64(int64(q.SHA1))},
				"SHA224":      {Typ: reflect.TypeOf(q.SHA224), Value: constant.MakeInt64(int64(q.SHA224))},
				"SHA256":      {Typ: reflect.TypeOf(q.SHA256), Value: constant.MakeInt64(int64(q.SHA256))},
				"SHA384":      {Typ: reflect.TypeOf(q.SHA384), Value: constant.MakeInt64(int64(q.SHA384))},
				"SHA3_224":    {Typ: reflect.TypeOf(q.SHA3_224), Value: constant.MakeInt64(int64(q.SHA3_224))},
				"SHA3_256":    {Typ: reflect.TypeOf(q.SHA3_256), Value: constant.MakeInt64(int64(q.SHA3_256))},
				"SHA3_384":    {Typ: reflect.TypeOf(q.SHA3_384), Value: constant.MakeInt64(int64(q.SHA3_384))},
				"SHA3_512":    {Typ: reflect.TypeOf(q.SHA3_512), Value: constant.MakeInt64(int64(q.SHA3_512))},
				"SHA512":      {Typ: reflect.TypeOf(q.SHA512), Value: constant.MakeInt64(int64(q.SHA512))},
				"SHA512_224":  {Typ: reflect.TypeOf(q.SHA512_224), Value: constant.MakeInt64(int64(q.SHA512_224))},
				"SHA512_256":  {Typ: reflect.TypeOf(q.SHA512_256), Value: constant.MakeInt64(int64(q.SHA512_256))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
