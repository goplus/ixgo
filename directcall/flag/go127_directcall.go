// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package flag

import (
	q "flag"

	encoding "encoding"
	"github.com/goplus/ixgo"
	io "io"
	time "time"
)

func init() {
	ixgo.RegisterDirectCalls("flag", map[string]ixgo.DirectCallAdapter{
		"(*FlagSet).Arg":           method_ptr_FlagSet_Arg,
		"(*FlagSet).Args":          method_ptr_FlagSet_Args,
		"(*FlagSet).Bool":          method_ptr_FlagSet_Bool,
		"(*FlagSet).BoolFunc":      method_ptr_FlagSet_BoolFunc,
		"(*FlagSet).BoolVar":       method_ptr_FlagSet_BoolVar,
		"(*FlagSet).Duration":      method_ptr_FlagSet_Duration,
		"(*FlagSet).DurationVar":   method_ptr_FlagSet_DurationVar,
		"(*FlagSet).ErrorHandling": method_ptr_FlagSet_ErrorHandling,
		"(*FlagSet).Float64":       method_ptr_FlagSet_Float64,
		"(*FlagSet).Float64Var":    method_ptr_FlagSet_Float64Var,
		"(*FlagSet).Func":          method_ptr_FlagSet_Func,
		"(*FlagSet).Init":          method_ptr_FlagSet_Init,
		"(*FlagSet).Int":           method_ptr_FlagSet_Int,
		"(*FlagSet).Int64":         method_ptr_FlagSet_Int64,
		"(*FlagSet).Int64Var":      method_ptr_FlagSet_Int64Var,
		"(*FlagSet).IntVar":        method_ptr_FlagSet_IntVar,
		"(*FlagSet).Lookup":        method_ptr_FlagSet_Lookup,
		"(*FlagSet).NArg":          method_ptr_FlagSet_NArg,
		"(*FlagSet).NFlag":         method_ptr_FlagSet_NFlag,
		"(*FlagSet).Name":          method_ptr_FlagSet_Name,
		"(*FlagSet).Output":        method_ptr_FlagSet_Output,
		"(*FlagSet).Parse":         method_ptr_FlagSet_Parse,
		"(*FlagSet).Parsed":        method_ptr_FlagSet_Parsed,
		"(*FlagSet).PrintDefaults": method_ptr_FlagSet_PrintDefaults,
		"(*FlagSet).Set":           method_ptr_FlagSet_Set,
		"(*FlagSet).SetOutput":     method_ptr_FlagSet_SetOutput,
		"(*FlagSet).String":        method_ptr_FlagSet_String,
		"(*FlagSet).StringVar":     method_ptr_FlagSet_StringVar,
		"(*FlagSet).TextVar":       method_ptr_FlagSet_TextVar,
		"(*FlagSet).Uint":          method_ptr_FlagSet_Uint,
		"(*FlagSet).Uint64":        method_ptr_FlagSet_Uint64,
		"(*FlagSet).Uint64Var":     method_ptr_FlagSet_Uint64Var,
		"(*FlagSet).UintVar":       method_ptr_FlagSet_UintVar,
		"(*FlagSet).Var":           method_ptr_FlagSet_Var,
		"(*FlagSet).Visit":         method_ptr_FlagSet_Visit,
		"(*FlagSet).VisitAll":      method_ptr_FlagSet_VisitAll,
		"Arg":                      func_Arg,
		"Args":                     func_Args,
		"Bool":                     func_Bool,
		"BoolFunc":                 func_BoolFunc,
		"BoolVar":                  func_BoolVar,
		"Duration":                 func_Duration,
		"DurationVar":              func_DurationVar,
		"Float64":                  func_Float64,
		"Float64Var":               func_Float64Var,
		"Func":                     func_Func,
		"Int":                      func_Int,
		"Int64":                    func_Int64,
		"Int64Var":                 func_Int64Var,
		"IntVar":                   func_IntVar,
		"Lookup":                   func_Lookup,
		"NArg":                     func_NArg,
		"NFlag":                    func_NFlag,
		"NewFlagSet":               func_NewFlagSet,
		"Parse":                    func_Parse,
		"Parsed":                   func_Parsed,
		"PrintDefaults":            func_PrintDefaults,
		"Set":                      func_Set,
		"String":                   func_String,
		"StringVar":                func_StringVar,
		"TextVar":                  func_TextVar,
		"Uint":                     func_Uint,
		"Uint64":                   func_Uint64,
		"Uint64Var":                func_Uint64Var,
		"UintVar":                  func_UintVar,
		"Var":                      func_Var,
		"Visit":                    func_Visit,
		"VisitAll":                 func_VisitAll,
	})
}

func func_Arg(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Arg(ixgo.DirectCallArg[int](ctx, 0)))
}

func func_Args(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Args())
}

func func_Bool(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Bool(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[bool](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func func_BoolFunc(ctx ixgo.DirectCallContext) {
	q.BoolFunc(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[func(string) error](ctx, 2))
}

func func_BoolVar(ctx ixgo.DirectCallContext) {
	q.BoolVar(ixgo.DirectCallArg[*bool](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[bool](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func func_Duration(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Duration(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[time.Duration](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func func_DurationVar(ctx ixgo.DirectCallContext) {
	q.DurationVar(ixgo.DirectCallArg[*time.Duration](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[time.Duration](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func method_ptr_FlagSet_Arg(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Arg(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_FlagSet_Args(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Args(ixgo.DirectCallArg[*q.FlagSet](ctx, 0)))
}

func method_ptr_FlagSet_Bool(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Bool(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[bool](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_FlagSet_BoolFunc(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).BoolFunc(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[func(string) error](ctx, 3))
}

func method_ptr_FlagSet_BoolVar(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).BoolVar(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[*bool](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[bool](ctx, 3), ixgo.DirectCallArg[string](ctx, 4))
}

func method_ptr_FlagSet_Duration(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Duration(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[time.Duration](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_FlagSet_DurationVar(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).DurationVar(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[*time.Duration](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[time.Duration](ctx, 3), ixgo.DirectCallArg[string](ctx, 4))
}

func method_ptr_FlagSet_ErrorHandling(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).ErrorHandling(ixgo.DirectCallArg[*q.FlagSet](ctx, 0)))
}

func method_ptr_FlagSet_Float64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Float64(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[float64](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_FlagSet_Float64Var(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).Float64Var(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[*float64](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[float64](ctx, 3), ixgo.DirectCallArg[string](ctx, 4))
}

func method_ptr_FlagSet_Func(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).Func(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[func(string) error](ctx, 3))
}

func method_ptr_FlagSet_Init(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).Init(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[q.ErrorHandling](ctx, 2))
}

func method_ptr_FlagSet_Int(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Int(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_FlagSet_Int64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Int64(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_FlagSet_Int64Var(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).Int64Var(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[*int64](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[int64](ctx, 3), ixgo.DirectCallArg[string](ctx, 4))
}

func method_ptr_FlagSet_IntVar(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).IntVar(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[*int](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[int](ctx, 3), ixgo.DirectCallArg[string](ctx, 4))
}

func method_ptr_FlagSet_Lookup(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Lookup(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_FlagSet_NArg(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).NArg(ixgo.DirectCallArg[*q.FlagSet](ctx, 0)))
}

func method_ptr_FlagSet_NFlag(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).NFlag(ixgo.DirectCallArg[*q.FlagSet](ctx, 0)))
}

func method_ptr_FlagSet_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Name(ixgo.DirectCallArg[*q.FlagSet](ctx, 0)))
}

func method_ptr_FlagSet_Output(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Output(ixgo.DirectCallArg[*q.FlagSet](ctx, 0)))
}

func method_ptr_FlagSet_Parse(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Parse(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[[]string](ctx, 1)))
}

func method_ptr_FlagSet_Parsed(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Parsed(ixgo.DirectCallArg[*q.FlagSet](ctx, 0)))
}

func method_ptr_FlagSet_PrintDefaults(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).PrintDefaults(ixgo.DirectCallArg[*q.FlagSet](ctx, 0))
}

func method_ptr_FlagSet_Set(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Set(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_FlagSet_SetOutput(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).SetOutput(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1))
}

func method_ptr_FlagSet_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).String(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_FlagSet_StringVar(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).StringVar(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[*string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3), ixgo.DirectCallArg[string](ctx, 4))
}

func method_ptr_FlagSet_TextVar(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).TextVar(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[encoding.TextUnmarshaler](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[encoding.TextMarshaler](ctx, 3), ixgo.DirectCallArg[string](ctx, 4))
}

func method_ptr_FlagSet_Uint(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Uint(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[uint](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_FlagSet_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FlagSet).Uint64(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[uint64](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func method_ptr_FlagSet_Uint64Var(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).Uint64Var(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[*uint64](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[uint64](ctx, 3), ixgo.DirectCallArg[string](ctx, 4))
}

func method_ptr_FlagSet_UintVar(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).UintVar(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[*uint](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[uint](ctx, 3), ixgo.DirectCallArg[string](ctx, 4))
}

func method_ptr_FlagSet_Var(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).Var(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[q.Value](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func method_ptr_FlagSet_Visit(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).Visit(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[func(*q.Flag)](ctx, 1))
}

func method_ptr_FlagSet_VisitAll(ctx ixgo.DirectCallContext) {
	(*q.FlagSet).VisitAll(ixgo.DirectCallArg[*q.FlagSet](ctx, 0), ixgo.DirectCallArg[func(*q.Flag)](ctx, 1))
}

func func_Float64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Float64(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[float64](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func func_Float64Var(ctx ixgo.DirectCallContext) {
	q.Float64Var(ixgo.DirectCallArg[*float64](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[float64](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func func_Func(ctx ixgo.DirectCallContext) {
	q.Func(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[func(string) error](ctx, 2))
}

func func_Int(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func func_Int64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Int64(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func func_Int64Var(ctx ixgo.DirectCallContext) {
	q.Int64Var(ixgo.DirectCallArg[*int64](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int64](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func func_IntVar(ctx ixgo.DirectCallContext) {
	q.IntVar(ixgo.DirectCallArg[*int](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func func_Lookup(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Lookup(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_NArg(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NArg())
}

func func_NFlag(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NFlag())
}

func func_NewFlagSet(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewFlagSet(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[q.ErrorHandling](ctx, 1)))
}

func func_Parse(ctx ixgo.DirectCallContext) {
	q.Parse()
}

func func_Parsed(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Parsed())
}

func func_PrintDefaults(ctx ixgo.DirectCallContext) {
	q.PrintDefaults()
}

func func_Set(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Set(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.String(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func func_StringVar(ctx ixgo.DirectCallContext) {
	q.StringVar(ixgo.DirectCallArg[*string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func func_TextVar(ctx ixgo.DirectCallContext) {
	q.TextVar(ixgo.DirectCallArg[encoding.TextUnmarshaler](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[encoding.TextMarshaler](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func func_Uint(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[uint](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func func_Uint64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Uint64(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func func_Uint64Var(ctx ixgo.DirectCallContext) {
	q.Uint64Var(ixgo.DirectCallArg[*uint64](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[uint64](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func func_UintVar(ctx ixgo.DirectCallContext) {
	q.UintVar(ixgo.DirectCallArg[*uint](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[uint](ctx, 2), ixgo.DirectCallArg[string](ctx, 3))
}

func func_Var(ctx ixgo.DirectCallContext) {
	q.Var(ixgo.DirectCallArg[q.Value](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2))
}

func func_Visit(ctx ixgo.DirectCallContext) {
	q.Visit(ixgo.DirectCallArg[func(*q.Flag)](ctx, 0))
}

func func_VisitAll(ctx ixgo.DirectCallContext) {
	q.VisitAll(ixgo.DirectCallArg[func(*q.Flag)](ctx, 0))
}
