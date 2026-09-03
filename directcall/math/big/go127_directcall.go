// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package big

import (
	q "math/big"

	fmt "fmt"
	"github.com/goplus/ixgo"
	rand "math/rand"
)

func init() {
	ixgo.RegisterDirectCalls("math/big", map[string]ixgo.DirectCallAdapter{
		"(*Accuracy).String":      method_ptr_Accuracy_String,
		"(*ErrNaN).Error":         method_ptr_ErrNaN_Error,
		"(*Float).Abs":            method_ptr_Float_Abs,
		"(*Float).Acc":            method_ptr_Float_Acc,
		"(*Float).Add":            method_ptr_Float_Add,
		"(*Float).Append":         method_ptr_Float_Append,
		"(*Float).Cmp":            method_ptr_Float_Cmp,
		"(*Float).Copy":           method_ptr_Float_Copy,
		"(*Float).Format":         method_ptr_Float_Format,
		"(*Float).GobDecode":      method_ptr_Float_GobDecode,
		"(*Float).IsInf":          method_ptr_Float_IsInf,
		"(*Float).IsInt":          method_ptr_Float_IsInt,
		"(*Float).MantExp":        method_ptr_Float_MantExp,
		"(*Float).MinPrec":        method_ptr_Float_MinPrec,
		"(*Float).Mode":           method_ptr_Float_Mode,
		"(*Float).Mul":            method_ptr_Float_Mul,
		"(*Float).Neg":            method_ptr_Float_Neg,
		"(*Float).Prec":           method_ptr_Float_Prec,
		"(*Float).Quo":            method_ptr_Float_Quo,
		"(*Float).Scan":           method_ptr_Float_Scan,
		"(*Float).Set":            method_ptr_Float_Set,
		"(*Float).SetFloat64":     method_ptr_Float_SetFloat64,
		"(*Float).SetInf":         method_ptr_Float_SetInf,
		"(*Float).SetInt":         method_ptr_Float_SetInt,
		"(*Float).SetInt64":       method_ptr_Float_SetInt64,
		"(*Float).SetMantExp":     method_ptr_Float_SetMantExp,
		"(*Float).SetMode":        method_ptr_Float_SetMode,
		"(*Float).SetPrec":        method_ptr_Float_SetPrec,
		"(*Float).SetRat":         method_ptr_Float_SetRat,
		"(*Float).SetUint64":      method_ptr_Float_SetUint64,
		"(*Float).Sign":           method_ptr_Float_Sign,
		"(*Float).Signbit":        method_ptr_Float_Signbit,
		"(*Float).Sqrt":           method_ptr_Float_Sqrt,
		"(*Float).String":         method_ptr_Float_String,
		"(*Float).Sub":            method_ptr_Float_Sub,
		"(*Float).Text":           method_ptr_Float_Text,
		"(*Float).UnmarshalText":  method_ptr_Float_UnmarshalText,
		"(*Int).Abs":              method_ptr_Int_Abs,
		"(*Int).Add":              method_ptr_Int_Add,
		"(*Int).And":              method_ptr_Int_And,
		"(*Int).AndNot":           method_ptr_Int_AndNot,
		"(*Int).Append":           method_ptr_Int_Append,
		"(*Int).Binomial":         method_ptr_Int_Binomial,
		"(*Int).Bit":              method_ptr_Int_Bit,
		"(*Int).BitLen":           method_ptr_Int_BitLen,
		"(*Int).Bits":             method_ptr_Int_Bits,
		"(*Int).Bytes":            method_ptr_Int_Bytes,
		"(*Int).Cmp":              method_ptr_Int_Cmp,
		"(*Int).CmpAbs":           method_ptr_Int_CmpAbs,
		"(*Int).Div":              method_ptr_Int_Div,
		"(*Int).Exp":              method_ptr_Int_Exp,
		"(*Int).FillBytes":        method_ptr_Int_FillBytes,
		"(*Int).Format":           method_ptr_Int_Format,
		"(*Int).GCD":              method_ptr_Int_GCD,
		"(*Int).GobDecode":        method_ptr_Int_GobDecode,
		"(*Int).Int64":            method_ptr_Int_Int64,
		"(*Int).IsInt64":          method_ptr_Int_IsInt64,
		"(*Int).IsUint64":         method_ptr_Int_IsUint64,
		"(*Int).Lsh":              method_ptr_Int_Lsh,
		"(*Int).Mod":              method_ptr_Int_Mod,
		"(*Int).ModInverse":       method_ptr_Int_ModInverse,
		"(*Int).ModSqrt":          method_ptr_Int_ModSqrt,
		"(*Int).Mul":              method_ptr_Int_Mul,
		"(*Int).MulRange":         method_ptr_Int_MulRange,
		"(*Int).Neg":              method_ptr_Int_Neg,
		"(*Int).Not":              method_ptr_Int_Not,
		"(*Int).Or":               method_ptr_Int_Or,
		"(*Int).ProbablyPrime":    method_ptr_Int_ProbablyPrime,
		"(*Int).Quo":              method_ptr_Int_Quo,
		"(*Int).Rand":             method_ptr_Int_Rand,
		"(*Int).Rem":              method_ptr_Int_Rem,
		"(*Int).Rsh":              method_ptr_Int_Rsh,
		"(*Int).Scan":             method_ptr_Int_Scan,
		"(*Int).Set":              method_ptr_Int_Set,
		"(*Int).SetBit":           method_ptr_Int_SetBit,
		"(*Int).SetBits":          method_ptr_Int_SetBits,
		"(*Int).SetBytes":         method_ptr_Int_SetBytes,
		"(*Int).SetInt64":         method_ptr_Int_SetInt64,
		"(*Int).SetUint64":        method_ptr_Int_SetUint64,
		"(*Int).Sign":             method_ptr_Int_Sign,
		"(*Int).Sqrt":             method_ptr_Int_Sqrt,
		"(*Int).String":           method_ptr_Int_String,
		"(*Int).Sub":              method_ptr_Int_Sub,
		"(*Int).Text":             method_ptr_Int_Text,
		"(*Int).TrailingZeroBits": method_ptr_Int_TrailingZeroBits,
		"(*Int).Uint64":           method_ptr_Int_Uint64,
		"(*Int).UnmarshalJSON":    method_ptr_Int_UnmarshalJSON,
		"(*Int).UnmarshalText":    method_ptr_Int_UnmarshalText,
		"(*Int).Xor":              method_ptr_Int_Xor,
		"(*Rat).Abs":              method_ptr_Rat_Abs,
		"(*Rat).Add":              method_ptr_Rat_Add,
		"(*Rat).Cmp":              method_ptr_Rat_Cmp,
		"(*Rat).Denom":            method_ptr_Rat_Denom,
		"(*Rat).FloatString":      method_ptr_Rat_FloatString,
		"(*Rat).GobDecode":        method_ptr_Rat_GobDecode,
		"(*Rat).Inv":              method_ptr_Rat_Inv,
		"(*Rat).IsInt":            method_ptr_Rat_IsInt,
		"(*Rat).Mul":              method_ptr_Rat_Mul,
		"(*Rat).Neg":              method_ptr_Rat_Neg,
		"(*Rat).Num":              method_ptr_Rat_Num,
		"(*Rat).Quo":              method_ptr_Rat_Quo,
		"(*Rat).RatString":        method_ptr_Rat_RatString,
		"(*Rat).Scan":             method_ptr_Rat_Scan,
		"(*Rat).Set":              method_ptr_Rat_Set,
		"(*Rat).SetFloat64":       method_ptr_Rat_SetFloat64,
		"(*Rat).SetFrac":          method_ptr_Rat_SetFrac,
		"(*Rat).SetFrac64":        method_ptr_Rat_SetFrac64,
		"(*Rat).SetInt":           method_ptr_Rat_SetInt,
		"(*Rat).SetInt64":         method_ptr_Rat_SetInt64,
		"(*Rat).SetUint64":        method_ptr_Rat_SetUint64,
		"(*Rat).Sign":             method_ptr_Rat_Sign,
		"(*Rat).String":           method_ptr_Rat_String,
		"(*Rat).Sub":              method_ptr_Rat_Sub,
		"(*Rat).UnmarshalText":    method_ptr_Rat_UnmarshalText,
		"(*RoundingMode).String":  method_ptr_RoundingMode_String,
		"(Accuracy).String":       method_Accuracy_String,
		"(ErrNaN).Error":          method_ErrNaN_Error,
		"(RoundingMode).String":   method_RoundingMode_String,
		"Jacobi":                  func_Jacobi,
		"NewFloat":                func_NewFloat,
		"NewInt":                  func_NewInt,
		"NewRat":                  func_NewRat,
	})
}

func method_Accuracy_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Accuracy.String(ixgo.DirectCallArg[q.Accuracy](ctx, 0)))
}

func method_ptr_Accuracy_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Accuracy).String(ixgo.DirectCallArg[*q.Accuracy](ctx, 0)))
}

func method_ErrNaN_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ErrNaN.Error(ixgo.DirectCallArg[q.ErrNaN](ctx, 0)))
}

func method_ptr_ErrNaN_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ErrNaN).Error(ixgo.DirectCallArg[*q.ErrNaN](ctx, 0)))
}

func method_ptr_Float_Abs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Abs(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1)))
}

func method_ptr_Float_Acc(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Acc(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Add(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1), ixgo.DirectCallArg[*q.Float](ctx, 2)))
}

func method_ptr_Float_Append(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Append(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[byte](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func method_ptr_Float_Cmp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Cmp(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1)))
}

func method_ptr_Float_Copy(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Copy(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1)))
}

func method_ptr_Float_Format(ctx ixgo.DirectCallContext) {
	(*q.Float).Format(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[fmt.State](ctx, 1), ixgo.DirectCallArg[rune](ctx, 2))
}

func method_ptr_Float_GobDecode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).GobDecode(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Float_IsInf(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).IsInf(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_IsInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).IsInt(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_MantExp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).MantExp(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1)))
}

func method_ptr_Float_MinPrec(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).MinPrec(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_Mode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Mode(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_Mul(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Mul(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1), ixgo.DirectCallArg[*q.Float](ctx, 2)))
}

func method_ptr_Float_Neg(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Neg(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1)))
}

func method_ptr_Float_Prec(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Prec(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_Quo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Quo(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1), ixgo.DirectCallArg[*q.Float](ctx, 2)))
}

func method_ptr_Float_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Scan(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[fmt.ScanState](ctx, 1), ixgo.DirectCallArg[rune](ctx, 2)))
}

func method_ptr_Float_Set(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Set(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1)))
}

func method_ptr_Float_SetFloat64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).SetFloat64(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func method_ptr_Float_SetInf(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).SetInf(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1)))
}

func method_ptr_Float_SetInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).SetInt(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func method_ptr_Float_SetInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).SetInt64(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Float_SetMantExp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).SetMantExp(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Float_SetMode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).SetMode(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[q.RoundingMode](ctx, 1)))
}

func method_ptr_Float_SetPrec(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).SetPrec(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[uint](ctx, 1)))
}

func method_ptr_Float_SetRat(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).SetRat(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1)))
}

func method_ptr_Float_SetUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).SetUint64(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Float_Sign(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Sign(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_Signbit(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Signbit(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_Sqrt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Sqrt(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1)))
}

func method_ptr_Float_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).String(ixgo.DirectCallArg[*q.Float](ctx, 0)))
}

func method_ptr_Float_Sub(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Sub(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[*q.Float](ctx, 1), ixgo.DirectCallArg[*q.Float](ctx, 2)))
}

func method_ptr_Float_Text(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).Text(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Float_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Float).UnmarshalText(ixgo.DirectCallArg[*q.Float](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Int_Abs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Abs(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func method_ptr_Int_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Add(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_And(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).And(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_AndNot(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).AndNot(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_Append(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Append(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_Int_Binomial(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Binomial(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2)))
}

func method_ptr_Int_Bit(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Bit(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Int_BitLen(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).BitLen(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_Bits(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Bits(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Bytes(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_Cmp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Cmp(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func method_ptr_Int_CmpAbs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).CmpAbs(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func method_ptr_Int_Div(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Div(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_Exp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Exp(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2), ixgo.DirectCallArg[*q.Int](ctx, 3)))
}

func method_ptr_Int_FillBytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).FillBytes(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Int_Format(ctx ixgo.DirectCallContext) {
	(*q.Int).Format(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[fmt.State](ctx, 1), ixgo.DirectCallArg[rune](ctx, 2))
}

func method_ptr_Int_GCD(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).GCD(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2), ixgo.DirectCallArg[*q.Int](ctx, 3), ixgo.DirectCallArg[*q.Int](ctx, 4)))
}

func method_ptr_Int_GobDecode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).GobDecode(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Int_Int64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Int64(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_IsInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).IsInt64(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_IsUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).IsUint64(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_Lsh(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Lsh(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[uint](ctx, 2)))
}

func method_ptr_Int_Mod(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Mod(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_ModInverse(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).ModInverse(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_ModSqrt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).ModSqrt(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_Mul(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Mul(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_MulRange(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).MulRange(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2)))
}

func method_ptr_Int_Neg(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Neg(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func method_ptr_Int_Not(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Not(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func method_ptr_Int_Or(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Or(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_ProbablyPrime(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).ProbablyPrime(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Int_Quo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Quo(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_Rand(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Rand(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*rand.Rand](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_Rem(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Rem(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_Rsh(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Rsh(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[uint](ctx, 2)))
}

func method_ptr_Int_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Scan(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[fmt.ScanState](ctx, 1), ixgo.DirectCallArg[rune](ctx, 2)))
}

func method_ptr_Int_Set(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Set(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func method_ptr_Int_SetBit(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).SetBit(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[uint](ctx, 3)))
}

func method_ptr_Int_SetBits(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).SetBits(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[[]q.Word](ctx, 1)))
}

func method_ptr_Int_SetBytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).SetBytes(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Int_SetInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).SetInt64(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Int_SetUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).SetUint64(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Int_Sign(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Sign(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_Sqrt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Sqrt(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func method_ptr_Int_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).String(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_Sub(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Sub(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Int_Text(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Text(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Int_TrailingZeroBits(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).TrailingZeroBits(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Uint64(ixgo.DirectCallArg[*q.Int](ctx, 0)))
}

func method_ptr_Int_UnmarshalJSON(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).UnmarshalJSON(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Int_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).UnmarshalText(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Int_Xor(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Int).Xor(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func func_Jacobi(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Jacobi(ixgo.DirectCallArg[*q.Int](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func func_NewFloat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewFloat(ixgo.DirectCallArg[float64](ctx, 0)))
}

func func_NewInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewInt(ixgo.DirectCallArg[int64](ctx, 0)))
}

func func_NewRat(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewRat(ixgo.DirectCallArg[int64](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Rat_Abs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Abs(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1)))
}

func method_ptr_Rat_Add(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Add(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1), ixgo.DirectCallArg[*q.Rat](ctx, 2)))
}

func method_ptr_Rat_Cmp(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Cmp(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1)))
}

func method_ptr_Rat_Denom(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Denom(ixgo.DirectCallArg[*q.Rat](ctx, 0)))
}

func method_ptr_Rat_FloatString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).FloatString(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Rat_GobDecode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).GobDecode(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_ptr_Rat_Inv(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Inv(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1)))
}

func method_ptr_Rat_IsInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).IsInt(ixgo.DirectCallArg[*q.Rat](ctx, 0)))
}

func method_ptr_Rat_Mul(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Mul(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1), ixgo.DirectCallArg[*q.Rat](ctx, 2)))
}

func method_ptr_Rat_Neg(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Neg(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1)))
}

func method_ptr_Rat_Num(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Num(ixgo.DirectCallArg[*q.Rat](ctx, 0)))
}

func method_ptr_Rat_Quo(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Quo(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1), ixgo.DirectCallArg[*q.Rat](ctx, 2)))
}

func method_ptr_Rat_RatString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).RatString(ixgo.DirectCallArg[*q.Rat](ctx, 0)))
}

func method_ptr_Rat_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Scan(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[fmt.ScanState](ctx, 1), ixgo.DirectCallArg[rune](ctx, 2)))
}

func method_ptr_Rat_Set(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Set(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1)))
}

func method_ptr_Rat_SetFloat64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).SetFloat64(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1)))
}

func method_ptr_Rat_SetFrac(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).SetFrac(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1), ixgo.DirectCallArg[*q.Int](ctx, 2)))
}

func method_ptr_Rat_SetFrac64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).SetFrac64(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2)))
}

func method_ptr_Rat_SetInt(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).SetInt(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Int](ctx, 1)))
}

func method_ptr_Rat_SetInt64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).SetInt64(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Rat_SetUint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).SetUint64(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Rat_Sign(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Sign(ixgo.DirectCallArg[*q.Rat](ctx, 0)))
}

func method_ptr_Rat_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).String(ixgo.DirectCallArg[*q.Rat](ctx, 0)))
}

func method_ptr_Rat_Sub(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).Sub(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[*q.Rat](ctx, 1), ixgo.DirectCallArg[*q.Rat](ctx, 2)))
}

func method_ptr_Rat_UnmarshalText(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rat).UnmarshalText(ixgo.DirectCallArg[*q.Rat](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func method_RoundingMode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RoundingMode.String(ixgo.DirectCallArg[q.RoundingMode](ctx, 0)))
}

func method_ptr_RoundingMode_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RoundingMode).String(ixgo.DirectCallArg[*q.RoundingMode](ctx, 0)))
}
