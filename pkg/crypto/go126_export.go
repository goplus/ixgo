// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package crypto

import (
	q "crypto"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "crypto",
		Path: "crypto",
		Deps: map[string]string{
			"hash":    "hash",
			"io":      "io",
			"strconv": "strconv",
		},
		Interfaces: map[string]reflect.Type{
			"Decapsulator":  reflect.TypeOf((*q.Decapsulator)(nil)).Elem(),
			"Decrypter":     reflect.TypeOf((*q.Decrypter)(nil)).Elem(),
			"DecrypterOpts": reflect.TypeOf((*q.DecrypterOpts)(nil)).Elem(),
			"Encapsulator":  reflect.TypeOf((*q.Encapsulator)(nil)).Elem(),
			"MessageSigner": reflect.TypeOf((*q.MessageSigner)(nil)).Elem(),
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
			"SignMessage":  reflect.ValueOf(q.SignMessage),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"BLAKE2b_256": {reflect.TypeOf(q.BLAKE2b_256), constant.MakeInt64(int64(q.BLAKE2b_256))},
			"BLAKE2b_384": {reflect.TypeOf(q.BLAKE2b_384), constant.MakeInt64(int64(q.BLAKE2b_384))},
			"BLAKE2b_512": {reflect.TypeOf(q.BLAKE2b_512), constant.MakeInt64(int64(q.BLAKE2b_512))},
			"BLAKE2s_256": {reflect.TypeOf(q.BLAKE2s_256), constant.MakeInt64(int64(q.BLAKE2s_256))},
			"MD4":         {reflect.TypeOf(q.MD4), constant.MakeInt64(int64(q.MD4))},
			"MD5":         {reflect.TypeOf(q.MD5), constant.MakeInt64(int64(q.MD5))},
			"MD5SHA1":     {reflect.TypeOf(q.MD5SHA1), constant.MakeInt64(int64(q.MD5SHA1))},
			"RIPEMD160":   {reflect.TypeOf(q.RIPEMD160), constant.MakeInt64(int64(q.RIPEMD160))},
			"SHA1":        {reflect.TypeOf(q.SHA1), constant.MakeInt64(int64(q.SHA1))},
			"SHA224":      {reflect.TypeOf(q.SHA224), constant.MakeInt64(int64(q.SHA224))},
			"SHA256":      {reflect.TypeOf(q.SHA256), constant.MakeInt64(int64(q.SHA256))},
			"SHA384":      {reflect.TypeOf(q.SHA384), constant.MakeInt64(int64(q.SHA384))},
			"SHA3_224":    {reflect.TypeOf(q.SHA3_224), constant.MakeInt64(int64(q.SHA3_224))},
			"SHA3_256":    {reflect.TypeOf(q.SHA3_256), constant.MakeInt64(int64(q.SHA3_256))},
			"SHA3_384":    {reflect.TypeOf(q.SHA3_384), constant.MakeInt64(int64(q.SHA3_384))},
			"SHA3_512":    {reflect.TypeOf(q.SHA3_512), constant.MakeInt64(int64(q.SHA3_512))},
			"SHA512":      {reflect.TypeOf(q.SHA512), constant.MakeInt64(int64(q.SHA512))},
			"SHA512_224":  {reflect.TypeOf(q.SHA512_224), constant.MakeInt64(int64(q.SHA512_224))},
			"SHA512_256":  {reflect.TypeOf(q.SHA512_256), constant.MakeInt64(int64(q.SHA512_256))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
