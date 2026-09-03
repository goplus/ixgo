// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package rand

import (
	q "math/rand"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("math/rand", map[string]ixgo.DirectCallAdapter{
		"(*Rand).ExpFloat64":  method_ptr_Rand_ExpFloat64,
		"(*Rand).Float32":     method_ptr_Rand_Float32,
		"(*Rand).Float64":     method_ptr_Rand_Float64,
		"(*Rand).Int":         method_ptr_Rand_Int,
		"(*Rand).Int31":       method_ptr_Rand_Int31,
		"(*Rand).Int31n":      method_ptr_Rand_Int31n,
		"(*Rand).Int63":       method_ptr_Rand_Int63,
		"(*Rand).Int63n":      method_ptr_Rand_Int63n,
		"(*Rand).Intn":        method_ptr_Rand_Intn,
		"(*Rand).NormFloat64": method_ptr_Rand_NormFloat64,
		"(*Rand).Perm":        method_ptr_Rand_Perm,
		"(*Rand).Seed":        method_ptr_Rand_Seed,
		"(*Rand).Shuffle":     method_ptr_Rand_Shuffle,
		"(*Rand).Uint32":      method_ptr_Rand_Uint32,
		"(*Rand).Uint64":      method_ptr_Rand_Uint64,
		"(*Zipf).Uint64":      method_ptr_Zipf_Uint64,
		"ExpFloat64":          func_ExpFloat64,
		"Float32":             func_Float32,
		"Float64":             func_Float64,
		"Int":                 func_Int,
		"Int31":               func_Int31,
		"Int31n":              func_Int31n,
		"Int63":               func_Int63,
		"Int63n":              func_Int63n,
		"Intn":                func_Intn,
		"New":                 func_New,
		"NewSource":           func_NewSource,
		"NewZipf":             func_NewZipf,
		"NormFloat64":         func_NormFloat64,
		"Perm":                func_Perm,
		"Seed":                func_Seed,
		"Shuffle":             func_Shuffle,
		"Uint32":              func_Uint32,
		"Uint64":              func_Uint64,
	})
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

func func_Int31(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int31())
}

func func_Int31n(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int31n(ixgo.DirectCallArg[int32](ctx, 0)))
}

func func_Int63(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int63())
}

func func_Int63n(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int63n(ixgo.DirectCallArg[int64](ctx, 0)))
}

func func_Intn(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Intn(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_New(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.New(ixgo.DirectCallArg[q.Source](ctx, 0)))
}

func func_NewSource(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewSource(ixgo.DirectCallArg[int64](ctx, 0)))
}

func func_NewZipf(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewZipf(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1), ixgo.DirectCallArg[float64](ctx, 2), ixgo.DirectCallArg[uint64](ctx, 3)))
}

func func_NormFloat64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NormFloat64())
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

func method_ptr_Rand_Int31(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Int31(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Int31n(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Int31n(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int32](ctx, 1)))
}

func method_ptr_Rand_Int63(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Int63(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Int63n(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Int63n(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func method_ptr_Rand_Intn(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Intn(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Rand_NormFloat64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).NormFloat64(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Perm(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Perm(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Rand_Seed(ctx ixgo.DirectCallContext) {
	(*q.Rand).Seed(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1))
}

func method_ptr_Rand_Shuffle(ctx ixgo.DirectCallContext) {
	(*q.Rand).Shuffle(ixgo.DirectCallArg[*q.Rand](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[func(i int, j int)](ctx, 2))
}

func method_ptr_Rand_Uint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Uint32(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func method_ptr_Rand_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Rand).Uint64(ixgo.DirectCallArg[*q.Rand](ctx, 0)))
}

func func_Seed(ctx ixgo.DirectCallContext) {
	q.Seed(ixgo.DirectCallArg[int64](ctx, 0))
}

func func_Shuffle(ctx ixgo.DirectCallContext) {
	q.Shuffle(ixgo.DirectCallArg[int](ctx, 0), ixgo.DirectCallArg[func(i int, j int)](ctx, 1))
}

func func_Uint32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint32())
}

func func_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint64())
}

func method_ptr_Zipf_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Zipf).Uint64(ixgo.DirectCallArg[*q.Zipf](ctx, 0)))
}
