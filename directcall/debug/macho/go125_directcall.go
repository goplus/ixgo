// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package macho

import (
	q "debug/macho"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("debug/macho", map[string]ixgo.DirectCallAdapter{
		"(*Cpu).GoString":              method_ptr_Cpu_GoString,
		"(*Cpu).String":                method_ptr_Cpu_String,
		"(*FatFile).Close":             method_ptr_FatFile_Close,
		"(*File).Close":                method_ptr_File_Close,
		"(*File).Section":              method_ptr_File_Section,
		"(*File).Segment":              method_ptr_File_Segment,
		"(*FormatError).Error":         method_ptr_FormatError_Error,
		"(*LoadBytes).Raw":             method_ptr_LoadBytes_Raw,
		"(*LoadCmd).GoString":          method_ptr_LoadCmd_GoString,
		"(*LoadCmd).String":            method_ptr_LoadCmd_String,
		"(*RelocTypeARM).GoString":     method_ptr_RelocTypeARM_GoString,
		"(*RelocTypeARM).String":       method_ptr_RelocTypeARM_String,
		"(*RelocTypeARM64).GoString":   method_ptr_RelocTypeARM64_GoString,
		"(*RelocTypeARM64).String":     method_ptr_RelocTypeARM64_String,
		"(*RelocTypeGeneric).GoString": method_ptr_RelocTypeGeneric_GoString,
		"(*RelocTypeGeneric).String":   method_ptr_RelocTypeGeneric_String,
		"(*RelocTypeX86_64).GoString":  method_ptr_RelocTypeX86_64_GoString,
		"(*RelocTypeX86_64).String":    method_ptr_RelocTypeX86_64_String,
		"(*Section).Open":              method_ptr_Section_Open,
		"(*Segment).Open":              method_ptr_Segment_Open,
		"(*Type).GoString":             method_ptr_Type_GoString,
		"(*Type).String":               method_ptr_Type_String,
		"(Cpu).GoString":               method_Cpu_GoString,
		"(Cpu).String":                 method_Cpu_String,
		"(LoadBytes).Raw":              method_LoadBytes_Raw,
		"(LoadCmd).GoString":           method_LoadCmd_GoString,
		"(LoadCmd).String":             method_LoadCmd_String,
		"(RelocTypeARM).GoString":      method_RelocTypeARM_GoString,
		"(RelocTypeARM).String":        method_RelocTypeARM_String,
		"(RelocTypeARM64).GoString":    method_RelocTypeARM64_GoString,
		"(RelocTypeARM64).String":      method_RelocTypeARM64_String,
		"(RelocTypeGeneric).GoString":  method_RelocTypeGeneric_GoString,
		"(RelocTypeGeneric).String":    method_RelocTypeGeneric_String,
		"(RelocTypeX86_64).GoString":   method_RelocTypeX86_64_GoString,
		"(RelocTypeX86_64).String":     method_RelocTypeX86_64_String,
		"(Type).GoString":              method_Type_GoString,
		"(Type).String":                method_Type_String,
	})
}

func method_Cpu_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Cpu.GoString(ixgo.DirectCallArg[q.Cpu](ctx, 0)))
}

func method_ptr_Cpu_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Cpu).GoString(ixgo.DirectCallArg[*q.Cpu](ctx, 0)))
}

func method_Cpu_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Cpu.String(ixgo.DirectCallArg[q.Cpu](ctx, 0)))
}

func method_ptr_Cpu_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Cpu).String(ixgo.DirectCallArg[*q.Cpu](ctx, 0)))
}

func method_ptr_FatFile_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FatFile).Close(ixgo.DirectCallArg[*q.FatFile](ctx, 0)))
}

func method_ptr_File_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Close(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Section(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Section(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_File_Segment(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Segment(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_FormatError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FormatError).Error(ixgo.DirectCallArg[*q.FormatError](ctx, 0)))
}

func method_LoadBytes_Raw(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LoadBytes.Raw(ixgo.DirectCallArg[q.LoadBytes](ctx, 0)))
}

func method_ptr_LoadBytes_Raw(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LoadBytes).Raw(ixgo.DirectCallArg[*q.LoadBytes](ctx, 0)))
}

func method_LoadCmd_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LoadCmd.GoString(ixgo.DirectCallArg[q.LoadCmd](ctx, 0)))
}

func method_ptr_LoadCmd_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LoadCmd).GoString(ixgo.DirectCallArg[*q.LoadCmd](ctx, 0)))
}

func method_LoadCmd_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LoadCmd.String(ixgo.DirectCallArg[q.LoadCmd](ctx, 0)))
}

func method_ptr_LoadCmd_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LoadCmd).String(ixgo.DirectCallArg[*q.LoadCmd](ctx, 0)))
}

func method_RelocTypeARM_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RelocTypeARM.GoString(ixgo.DirectCallArg[q.RelocTypeARM](ctx, 0)))
}

func method_ptr_RelocTypeARM_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RelocTypeARM).GoString(ixgo.DirectCallArg[*q.RelocTypeARM](ctx, 0)))
}

func method_RelocTypeARM_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RelocTypeARM.String(ixgo.DirectCallArg[q.RelocTypeARM](ctx, 0)))
}

func method_ptr_RelocTypeARM_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RelocTypeARM).String(ixgo.DirectCallArg[*q.RelocTypeARM](ctx, 0)))
}

func method_RelocTypeARM64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RelocTypeARM64.GoString(ixgo.DirectCallArg[q.RelocTypeARM64](ctx, 0)))
}

func method_ptr_RelocTypeARM64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RelocTypeARM64).GoString(ixgo.DirectCallArg[*q.RelocTypeARM64](ctx, 0)))
}

func method_RelocTypeARM64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RelocTypeARM64.String(ixgo.DirectCallArg[q.RelocTypeARM64](ctx, 0)))
}

func method_ptr_RelocTypeARM64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RelocTypeARM64).String(ixgo.DirectCallArg[*q.RelocTypeARM64](ctx, 0)))
}

func method_RelocTypeGeneric_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RelocTypeGeneric.GoString(ixgo.DirectCallArg[q.RelocTypeGeneric](ctx, 0)))
}

func method_ptr_RelocTypeGeneric_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RelocTypeGeneric).GoString(ixgo.DirectCallArg[*q.RelocTypeGeneric](ctx, 0)))
}

func method_RelocTypeGeneric_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RelocTypeGeneric.String(ixgo.DirectCallArg[q.RelocTypeGeneric](ctx, 0)))
}

func method_ptr_RelocTypeGeneric_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RelocTypeGeneric).String(ixgo.DirectCallArg[*q.RelocTypeGeneric](ctx, 0)))
}

func method_RelocTypeX86_64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RelocTypeX86_64.GoString(ixgo.DirectCallArg[q.RelocTypeX86_64](ctx, 0)))
}

func method_ptr_RelocTypeX86_64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RelocTypeX86_64).GoString(ixgo.DirectCallArg[*q.RelocTypeX86_64](ctx, 0)))
}

func method_RelocTypeX86_64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RelocTypeX86_64.String(ixgo.DirectCallArg[q.RelocTypeX86_64](ctx, 0)))
}

func method_ptr_RelocTypeX86_64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.RelocTypeX86_64).String(ixgo.DirectCallArg[*q.RelocTypeX86_64](ctx, 0)))
}

func method_ptr_Section_Open(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Section).Open(ixgo.DirectCallArg[*q.Section](ctx, 0)))
}

func method_ptr_Segment_Open(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Segment).Open(ixgo.DirectCallArg[*q.Segment](ctx, 0)))
}

func method_Type_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Type.GoString(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func method_ptr_Type_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Type).GoString(ixgo.DirectCallArg[*q.Type](ctx, 0)))
}

func method_Type_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Type.String(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func method_ptr_Type_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Type).String(ixgo.DirectCallArg[*q.Type](ctx, 0)))
}
