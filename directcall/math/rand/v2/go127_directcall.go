// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package rand

import (
	q "math/rand/v2"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("math/rand/v2", map[string]ixgo.DirectCallAdapter{
		"(*ChaCha8).Seed":            method_ptr_ChaCha8_Seed,
		"(*ChaCha8).Uint64":          method_ptr_ChaCha8_Uint64,
		"(*ChaCha8).UnmarshalBinary": method_ptr_ChaCha8_UnmarshalBinary,
		"(*PCG).Seed":                method_ptr_PCG_Seed,
		"(*PCG).Uint64":              method_ptr_PCG_Uint64,
		"(*PCG).UnmarshalBinary":     method_ptr_PCG_UnmarshalBinary,
		"(*Rand).ExpFloat64":         method_ptr_Rand_ExpFloat64,
		"(*Rand).Float32":            method_ptr_Rand_Float32,
		"(*Rand).Float64":            method_ptr_Rand_Float64,
		"(*Rand).Int":                method_ptr_Rand_Int,
		"(*Rand).Int32":              method_ptr_Rand_Int32,
		"(*Rand).Int32N":             method_ptr_Rand_Int32N,
		"(*Rand).Int64":              method_ptr_Rand_Int64,
		"(*Rand).Int64N":             method_ptr_Rand_Int64N,
		"(*Rand).IntN":               method_ptr_Rand_IntN,
		"(*Rand).NormFloat64":        method_ptr_Rand_NormFloat64,
		"(*Rand).Perm":               method_ptr_Rand_Perm,
		"(*Rand).Shuffle":            method_ptr_Rand_Shuffle,
		"(*Rand).Uint":               method_ptr_Rand_Uint,
		"(*Rand).Uint32":             method_ptr_Rand_Uint32,
		"(*Rand).Uint32N":            method_ptr_Rand_Uint32N,
		"(*Rand).Uint64":             method_ptr_Rand_Uint64,
		"(*Rand).Uint64N":            method_ptr_Rand_Uint64N,
		"(*Rand).UintN":              method_ptr_Rand_UintN,
		"(*Zipf).Uint64":             method_ptr_Zipf_Uint64,
		"ExpFloat64":                 func_ExpFloat64,
		"Float32":                    func_Float32,
		"Float64":                    func_Float64,
		"Int":                        func_Int,
		"Int32":                      func_Int32,
		"Int32N":                     func_Int32N,
		"Int64":                      func_Int64,
		"Int64N":                     func_Int64N,
		"IntN":                       func_IntN,
		"New":                        func_New,
		"NewChaCha8":                 func_NewChaCha8,
		"NewPCG":                     func_NewPCG,
		"NewZipf":                    func_NewZipf,
		"NormFloat64":                func_NormFloat64,
		"Perm":                       func_Perm,
		"Shuffle":                    func_Shuffle,
		"Uint":                       func_Uint,
		"Uint32":                     func_Uint32,
		"Uint32N":                    func_Uint32N,
		"Uint64":                     func_Uint64,
		"Uint64N":                    func_Uint64N,
		"UintN":                      func_UintN,
	})
}

func method_ptr_ChaCha8_Seed(ctx ixgo.DirectCallContext) {
	(*q.ChaCha8).Seed(ixgo.DirectCallArg[*q.ChaCha8](ctx, 0), ixgo.DirectCallArg[[32]byte](ctx, 1))
}

func method_ptr_ChaCha8_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ChaCha8).Uint64(ixgo.DirectCallArg[*q.ChaCha8](ctx, 0)))
}

func method_ptr_ChaCha8_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ChaCha8).UnmarshalBinary(ixgo.DirectCallArg[*q.ChaCha8](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_ExpFloat64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ExpFloat64())
}

func func_Float32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float32())
}

func func_Float64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float64())
}

func func_Int(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int())
}

func func_Int32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int32())
}

func func_Int32N(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int32N(ixgo.DirectCallArg[int32](ctx, 0)))
}

func func_Int64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int64())
}

func func_Int64N(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int64N(ixgo.DirectCallArg[int64](ctx, 0)))
}

func func_IntN(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IntN(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[q.Source](ctx, 0)))
}

func func_NewChaCha8(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewChaCha8(ixgo.DirectCallArg[[32]byte](ctx, 0)))
}

func func_NewPCG(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewPCG(ixgo.DirectCallArg[uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func func_NewZipf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewZipf(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1), ixgo.DirectCallArg[float64](ctx, 2), ixgo.DirectCallArg[uint64](ctx, 3)))
}

func func_NormFloat64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NormFloat64())
}

func method_ptr_PCG_Seed(ctx ixgo.DirectCallContext) {
	(*q.PCG).Seed(ixgo.DirectCallArg[*q.PCG](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1), ixgo.DirectCallArg[uint64](ctx, 2))
}

func method_ptr_PCG_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PCG).Uint64(ixgo.DirectCallArg[*q.PCG](ctx, 0)))
}

func method_ptr_PCG_UnmarshalBinary(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PCG).UnmarshalBinary(ixgo.DirectCallArg[*q.PCG](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1)))
}

func func_Perm(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Perm(ixgo.DirectCallArg[int](ctx, 0)))
}

func method_ptr_Rand_ExpFloat64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).ExpFloat64(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Float32(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Float32(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Float64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Float64(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Int(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Int(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Int32(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Int32(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Int32N(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Int32N(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func method_ptr_Rand_Int64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Int64(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Int64N(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Int64N(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Rand_IntN(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).IntN(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Rand_NormFloat64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).NormFloat64(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Perm(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Perm(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Rand_Shuffle(ctx ixgo.DirectCallContext) {
	(*q.Rand).Shuffle(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[func(i int, j int)](ctx, 2))
}

func method_ptr_Rand_Uint(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Uint(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Uint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Uint32(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Uint32N(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Uint32N(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func method_ptr_Rand_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Uint64(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Uint64N(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Uint64N(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1)))
}

func method_ptr_Rand_UintN(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).UintN(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[uint](ctx, 1)))
}

func func_Shuffle(ctx ixgo.DirectCallContext) {
	q.Shuffle(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[func(i int, j int)](ctx, 1))
}

func func_Uint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint())
}

func func_Uint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint32())
}

func func_Uint32N(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint32N(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint64())
}

func func_Uint64N(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint64N(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_UintN(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.UintN(ixgo.DirectCallArg[uint](ctx, 0)))
}

func method_ptr_Zipf_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Zipf).Uint64(ixgo.DirectCallArg[*q.Zipf](ctx, 0)))
}
