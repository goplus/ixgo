// export by github.com/goplus/ixgo/cmd/qexp

package ng

import (
	q "github.com/qiniu/x/xgo/ng"

	"go/constant"
	"go/token"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/qiniu/x/xgo/ng", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "ng",
			Path: "github.com/qiniu/x/xgo/ng",
			Deps: map[string]string{
				"fmt":       "fmt",
				"log":       "log",
				"math/big":  "big",
				"math/bits": "bits",
				"strconv":   "strconv",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Bigfloat":        reflect.TypeOf((*q.Bigfloat)(nil)).Elem(),
				"Bigint":          reflect.TypeOf((*q.Bigint)(nil)).Elem(),
				"Bigrat":          reflect.TypeOf((*q.Bigrat)(nil)).Elem(),
				"Int128":          reflect.TypeOf((*q.Int128)(nil)).Elem(),
				"Uint128":         reflect.TypeOf((*q.Uint128)(nil)).Elem(),
				"UntypedBigfloat": reflect.TypeOf((*q.UntypedBigfloat)(nil)).Elem(),
				"UntypedBigint":   reflect.TypeOf((*q.UntypedBigint)(nil)).Elem(),
				"UntypedBigrat":   reflect.TypeOf((*q.UntypedBigrat)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"UntypedBigfloat_Default": reflect.TypeOf((*q.UntypedBigfloat_Default)(nil)).Elem(),
				"UntypedBigint_Default":   reflect.TypeOf((*q.UntypedBigint_Default)(nil)).Elem(),
				"UntypedBigrat_Default":   reflect.TypeOf((*q.UntypedBigrat_Default)(nil)).Elem(),
				"XGo_ninteger":            reflect.TypeOf((*q.XGo_ninteger)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Bigint_Cast__0":        q.Bigint_Cast__0,
				"Bigint_Cast__1":        q.Bigint_Cast__1,
				"Bigint_Cast__2":        q.Bigint_Cast__2,
				"Bigint_Cast__3":        q.Bigint_Cast__3,
				"Bigint_Cast__4":        q.Bigint_Cast__4,
				"Bigint_Cast__5":        q.Bigint_Cast__5,
				"Bigint_Cast__6":        q.Bigint_Cast__6,
				"Bigint_Cast__7":        q.Bigint_Cast__7,
				"Bigint_Init__0":        q.Bigint_Init__0,
				"Bigint_Init__1":        q.Bigint_Init__1,
				"Bigint_Init__2":        q.Bigint_Init__2,
				"Bigrat_Cast__0":        q.Bigrat_Cast__0,
				"Bigrat_Cast__1":        q.Bigrat_Cast__1,
				"Bigrat_Cast__2":        q.Bigrat_Cast__2,
				"Bigrat_Cast__3":        q.Bigrat_Cast__3,
				"Bigrat_Cast__4":        q.Bigrat_Cast__4,
				"Bigrat_Cast__5":        q.Bigrat_Cast__5,
				"Bigrat_Cast__6":        q.Bigrat_Cast__6,
				"Bigrat_Init__0":        q.Bigrat_Init__0,
				"Bigrat_Init__1":        q.Bigrat_Init__1,
				"Bigrat_Init__2":        q.Bigrat_Init__2,
				"FormatInt128":          q.FormatInt128,
				"FormatUint128":         q.FormatUint128,
				"Int128_Cast__0":        q.Int128_Cast__0,
				"Int128_Cast__1":        q.Int128_Cast__1,
				"Int128_Cast__2":        q.Int128_Cast__2,
				"Int128_Cast__3":        q.Int128_Cast__3,
				"Int128_Cast__4":        q.Int128_Cast__4,
				"Int128_Cast__5":        q.Int128_Cast__5,
				"Int128_Cast__6":        q.Int128_Cast__6,
				"Int128_Cast__7":        q.Int128_Cast__7,
				"Int128_Cast__8":        q.Int128_Cast__8,
				"Int128_Cast__9":        q.Int128_Cast__9,
				"Int128_Cast__a":        q.Int128_Cast__a,
				"Int128_Init__0":        q.Int128_Init__0,
				"Int128_Init__1":        q.Int128_Init__1,
				"ParseInt128":           q.ParseInt128,
				"ParseUint128":          q.ParseUint128,
				"Uint128_Cast__0":       q.Uint128_Cast__0,
				"Uint128_Cast__1":       q.Uint128_Cast__1,
				"Uint128_Cast__2":       q.Uint128_Cast__2,
				"Uint128_Cast__3":       q.Uint128_Cast__3,
				"Uint128_Cast__4":       q.Uint128_Cast__4,
				"Uint128_Cast__5":       q.Uint128_Cast__5,
				"Uint128_Cast__6":       q.Uint128_Cast__6,
				"Uint128_Cast__7":       q.Uint128_Cast__7,
				"Uint128_Cast__8":       q.Uint128_Cast__8,
				"Uint128_Cast__9":       q.Uint128_Cast__9,
				"Uint128_Cast__a":       q.Uint128_Cast__a,
				"Uint128_Cast__b":       q.Uint128_Cast__b,
				"Uint128_Cast__c":       q.Uint128_Cast__c,
				"Uint128_Init__0":       q.Uint128_Init__0,
				"Uint128_Init__1":       q.Uint128_Init__1,
				"UntypedBigint_Init__0": q.UntypedBigint_Init__0,
				"UntypedBigrat_Init__0": q.UntypedBigrat_Init__0,
				"UntypedBigrat_Init__1": q.UntypedBigrat_Init__1,
				"XGo_istmp":             q.XGo_istmp,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"Int128_IsUntyped":  {Typ: "untyped bool", Value: constant.MakeBool(bool(q.Int128_IsUntyped))},
				"Int128_Max":        {Typ: "untyped int", Value: constant.MakeFromLiteral("170141183460469231731687303715884105727", token.INT, 0)},
				"Int128_Min":        {Typ: "untyped int", Value: constant.MakeFromLiteral("-170141183460469231731687303715884105728", token.INT, 0)},
				"Uint128_IsUntyped": {Typ: "untyped bool", Value: constant.MakeBool(bool(q.Uint128_IsUntyped))},
				"Uint128_Max":       {Typ: "untyped int", Value: constant.MakeFromLiteral("340282366920938463463374607431768211455", token.INT, 0)},
				"Uint128_Min":       {Typ: "untyped int", Value: constant.MakeInt64(int64(q.Uint128_Min))},
				"XGoPackage":        {Typ: "untyped bool", Value: constant.MakeBool(bool(q.XGoPackage))},
			},
		}
	})
}
