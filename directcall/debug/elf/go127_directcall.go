// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package elf

import (
	q "debug/elf"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("debug/elf", map[string]ixgo.DirectCallAdapter{
		"(*Class).GoString":           method_ptr_Class_GoString,
		"(*Class).String":             method_ptr_Class_String,
		"(*CompressionType).GoString": method_ptr_CompressionType_GoString,
		"(*CompressionType).String":   method_ptr_CompressionType_String,
		"(*Data).GoString":            method_ptr_Data_GoString,
		"(*Data).String":              method_ptr_Data_String,
		"(*DynFlag).GoString":         method_ptr_DynFlag_GoString,
		"(*DynFlag).String":           method_ptr_DynFlag_String,
		"(*DynFlag1).GoString":        method_ptr_DynFlag1_GoString,
		"(*DynFlag1).String":          method_ptr_DynFlag1_String,
		"(*DynTag).GoString":          method_ptr_DynTag_GoString,
		"(*DynTag).String":            method_ptr_DynTag_String,
		"(*File).Close":               method_ptr_File_Close,
		"(*File).Section":             method_ptr_File_Section,
		"(*File).SectionByType":       method_ptr_File_SectionByType,
		"(*FormatError).Error":        method_ptr_FormatError_Error,
		"(*Machine).GoString":         method_ptr_Machine_GoString,
		"(*Machine).String":           method_ptr_Machine_String,
		"(*NType).GoString":           method_ptr_NType_GoString,
		"(*NType).String":             method_ptr_NType_String,
		"(*OSABI).GoString":           method_ptr_OSABI_GoString,
		"(*OSABI).String":             method_ptr_OSABI_String,
		"(*Prog).Open":                method_ptr_Prog_Open,
		"(*ProgFlag).GoString":        method_ptr_ProgFlag_GoString,
		"(*ProgFlag).String":          method_ptr_ProgFlag_String,
		"(*ProgType).GoString":        method_ptr_ProgType_GoString,
		"(*ProgType).String":          method_ptr_ProgType_String,
		"(*R_386).GoString":           method_ptr_R_386_GoString,
		"(*R_386).String":             method_ptr_R_386_String,
		"(*R_390).GoString":           method_ptr_R_390_GoString,
		"(*R_390).String":             method_ptr_R_390_String,
		"(*R_AARCH64).GoString":       method_ptr_R_AARCH64_GoString,
		"(*R_AARCH64).String":         method_ptr_R_AARCH64_String,
		"(*R_ALPHA).GoString":         method_ptr_R_ALPHA_GoString,
		"(*R_ALPHA).String":           method_ptr_R_ALPHA_String,
		"(*R_ARM).GoString":           method_ptr_R_ARM_GoString,
		"(*R_ARM).String":             method_ptr_R_ARM_String,
		"(*R_LARCH).GoString":         method_ptr_R_LARCH_GoString,
		"(*R_LARCH).String":           method_ptr_R_LARCH_String,
		"(*R_MIPS).GoString":          method_ptr_R_MIPS_GoString,
		"(*R_MIPS).String":            method_ptr_R_MIPS_String,
		"(*R_PPC).GoString":           method_ptr_R_PPC_GoString,
		"(*R_PPC).String":             method_ptr_R_PPC_String,
		"(*R_PPC64).GoString":         method_ptr_R_PPC64_GoString,
		"(*R_PPC64).String":           method_ptr_R_PPC64_String,
		"(*R_RISCV).GoString":         method_ptr_R_RISCV_GoString,
		"(*R_RISCV).String":           method_ptr_R_RISCV_String,
		"(*R_SPARC).GoString":         method_ptr_R_SPARC_GoString,
		"(*R_SPARC).String":           method_ptr_R_SPARC_String,
		"(*R_X86_64).GoString":        method_ptr_R_X86_64_GoString,
		"(*R_X86_64).String":          method_ptr_R_X86_64_String,
		"(*Section).Open":             method_ptr_Section_Open,
		"(*SectionFlag).GoString":     method_ptr_SectionFlag_GoString,
		"(*SectionFlag).String":       method_ptr_SectionFlag_String,
		"(*SectionIndex).GoString":    method_ptr_SectionIndex_GoString,
		"(*SectionIndex).String":      method_ptr_SectionIndex_String,
		"(*SectionType).GoString":     method_ptr_SectionType_GoString,
		"(*SectionType).String":       method_ptr_SectionType_String,
		"(*SymBind).GoString":         method_ptr_SymBind_GoString,
		"(*SymBind).String":           method_ptr_SymBind_String,
		"(*SymType).GoString":         method_ptr_SymType_GoString,
		"(*SymType).String":           method_ptr_SymType_String,
		"(*SymVis).GoString":          method_ptr_SymVis_GoString,
		"(*SymVis).String":            method_ptr_SymVis_String,
		"(*Type).GoString":            method_ptr_Type_GoString,
		"(*Type).String":              method_ptr_Type_String,
		"(*Version).GoString":         method_ptr_Version_GoString,
		"(*Version).String":           method_ptr_Version_String,
		"(*VersionIndex).Index":       method_ptr_VersionIndex_Index,
		"(*VersionIndex).IsHidden":    method_ptr_VersionIndex_IsHidden,
		"(Class).GoString":            method_Class_GoString,
		"(Class).String":              method_Class_String,
		"(CompressionType).GoString":  method_CompressionType_GoString,
		"(CompressionType).String":    method_CompressionType_String,
		"(Data).GoString":             method_Data_GoString,
		"(Data).String":               method_Data_String,
		"(DynFlag).GoString":          method_DynFlag_GoString,
		"(DynFlag).String":            method_DynFlag_String,
		"(DynFlag1).GoString":         method_DynFlag1_GoString,
		"(DynFlag1).String":           method_DynFlag1_String,
		"(DynTag).GoString":           method_DynTag_GoString,
		"(DynTag).String":             method_DynTag_String,
		"(Machine).GoString":          method_Machine_GoString,
		"(Machine).String":            method_Machine_String,
		"(NType).GoString":            method_NType_GoString,
		"(NType).String":              method_NType_String,
		"(OSABI).GoString":            method_OSABI_GoString,
		"(OSABI).String":              method_OSABI_String,
		"(ProgFlag).GoString":         method_ProgFlag_GoString,
		"(ProgFlag).String":           method_ProgFlag_String,
		"(ProgType).GoString":         method_ProgType_GoString,
		"(ProgType).String":           method_ProgType_String,
		"(R_386).GoString":            method_R_386_GoString,
		"(R_386).String":              method_R_386_String,
		"(R_390).GoString":            method_R_390_GoString,
		"(R_390).String":              method_R_390_String,
		"(R_AARCH64).GoString":        method_R_AARCH64_GoString,
		"(R_AARCH64).String":          method_R_AARCH64_String,
		"(R_ALPHA).GoString":          method_R_ALPHA_GoString,
		"(R_ALPHA).String":            method_R_ALPHA_String,
		"(R_ARM).GoString":            method_R_ARM_GoString,
		"(R_ARM).String":              method_R_ARM_String,
		"(R_LARCH).GoString":          method_R_LARCH_GoString,
		"(R_LARCH).String":            method_R_LARCH_String,
		"(R_MIPS).GoString":           method_R_MIPS_GoString,
		"(R_MIPS).String":             method_R_MIPS_String,
		"(R_PPC).GoString":            method_R_PPC_GoString,
		"(R_PPC).String":              method_R_PPC_String,
		"(R_PPC64).GoString":          method_R_PPC64_GoString,
		"(R_PPC64).String":            method_R_PPC64_String,
		"(R_RISCV).GoString":          method_R_RISCV_GoString,
		"(R_RISCV).String":            method_R_RISCV_String,
		"(R_SPARC).GoString":          method_R_SPARC_GoString,
		"(R_SPARC).String":            method_R_SPARC_String,
		"(R_X86_64).GoString":         method_R_X86_64_GoString,
		"(R_X86_64).String":           method_R_X86_64_String,
		"(SectionFlag).GoString":      method_SectionFlag_GoString,
		"(SectionFlag).String":        method_SectionFlag_String,
		"(SectionIndex).GoString":     method_SectionIndex_GoString,
		"(SectionIndex).String":       method_SectionIndex_String,
		"(SectionType).GoString":      method_SectionType_GoString,
		"(SectionType).String":        method_SectionType_String,
		"(SymBind).GoString":          method_SymBind_GoString,
		"(SymBind).String":            method_SymBind_String,
		"(SymType).GoString":          method_SymType_GoString,
		"(SymType).String":            method_SymType_String,
		"(SymVis).GoString":           method_SymVis_GoString,
		"(SymVis).String":             method_SymVis_String,
		"(Type).GoString":             method_Type_GoString,
		"(Type).String":               method_Type_String,
		"(Version).GoString":          method_Version_GoString,
		"(Version).String":            method_Version_String,
		"(VersionIndex).Index":        method_VersionIndex_Index,
		"(VersionIndex).IsHidden":     method_VersionIndex_IsHidden,
		"R_INFO":                      func_R_INFO,
		"R_INFO32":                    func_R_INFO32,
		"R_SYM32":                     func_R_SYM32,
		"R_SYM64":                     func_R_SYM64,
		"R_TYPE32":                    func_R_TYPE32,
		"R_TYPE64":                    func_R_TYPE64,
		"ST_BIND":                     func_ST_BIND,
		"ST_INFO":                     func_ST_INFO,
		"ST_TYPE":                     func_ST_TYPE,
		"ST_VISIBILITY":               func_ST_VISIBILITY,
	})
}

func method_Class_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Class.GoString(ixgo.DirectCallArg[q.Class](ctx, 0)))
}

func method_ptr_Class_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Class).GoString(ixgo.DirectCallArg[*q.Class](ctx, 0)))
}

func method_Class_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Class.String(ixgo.DirectCallArg[q.Class](ctx, 0)))
}

func method_ptr_Class_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Class).String(ixgo.DirectCallArg[*q.Class](ctx, 0)))
}

func method_CompressionType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CompressionType.GoString(ixgo.DirectCallArg[q.CompressionType](ctx, 0)))
}

func method_ptr_CompressionType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CompressionType).GoString(ixgo.DirectCallArg[*q.CompressionType](ctx, 0)))
}

func method_CompressionType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CompressionType.String(ixgo.DirectCallArg[q.CompressionType](ctx, 0)))
}

func method_ptr_CompressionType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.CompressionType).String(ixgo.DirectCallArg[*q.CompressionType](ctx, 0)))
}

func method_Data_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Data.GoString(ixgo.DirectCallArg[q.Data](ctx, 0)))
}

func method_ptr_Data_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Data).GoString(ixgo.DirectCallArg[*q.Data](ctx, 0)))
}

func method_Data_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Data.String(ixgo.DirectCallArg[q.Data](ctx, 0)))
}

func method_ptr_Data_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Data).String(ixgo.DirectCallArg[*q.Data](ctx, 0)))
}

func method_DynFlag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DynFlag.GoString(ixgo.DirectCallArg[q.DynFlag](ctx, 0)))
}

func method_ptr_DynFlag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DynFlag).GoString(ixgo.DirectCallArg[*q.DynFlag](ctx, 0)))
}

func method_DynFlag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DynFlag.String(ixgo.DirectCallArg[q.DynFlag](ctx, 0)))
}

func method_ptr_DynFlag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DynFlag).String(ixgo.DirectCallArg[*q.DynFlag](ctx, 0)))
}

func method_DynFlag1_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DynFlag1.GoString(ixgo.DirectCallArg[q.DynFlag1](ctx, 0)))
}

func method_ptr_DynFlag1_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DynFlag1).GoString(ixgo.DirectCallArg[*q.DynFlag1](ctx, 0)))
}

func method_DynFlag1_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DynFlag1.String(ixgo.DirectCallArg[q.DynFlag1](ctx, 0)))
}

func method_ptr_DynFlag1_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DynFlag1).String(ixgo.DirectCallArg[*q.DynFlag1](ctx, 0)))
}

func method_DynTag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DynTag.GoString(ixgo.DirectCallArg[q.DynTag](ctx, 0)))
}

func method_ptr_DynTag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DynTag).GoString(ixgo.DirectCallArg[*q.DynTag](ctx, 0)))
}

func method_DynTag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DynTag.String(ixgo.DirectCallArg[q.DynTag](ctx, 0)))
}

func method_ptr_DynTag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.DynTag).String(ixgo.DirectCallArg[*q.DynTag](ctx, 0)))
}

func method_ptr_File_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Close(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Section(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Section(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_File_SectionByType(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).SectionByType(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[q.SectionType](ctx, 1)))
}

func method_ptr_FormatError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.FormatError).Error(ixgo.DirectCallArg[*q.FormatError](ctx, 0)))
}

func method_Machine_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Machine.GoString(ixgo.DirectCallArg[q.Machine](ctx, 0)))
}

func method_ptr_Machine_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Machine).GoString(ixgo.DirectCallArg[*q.Machine](ctx, 0)))
}

func method_Machine_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Machine.String(ixgo.DirectCallArg[q.Machine](ctx, 0)))
}

func method_ptr_Machine_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Machine).String(ixgo.DirectCallArg[*q.Machine](ctx, 0)))
}

func method_NType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NType.GoString(ixgo.DirectCallArg[q.NType](ctx, 0)))
}

func method_ptr_NType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NType).GoString(ixgo.DirectCallArg[*q.NType](ctx, 0)))
}

func method_NType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NType.String(ixgo.DirectCallArg[q.NType](ctx, 0)))
}

func method_ptr_NType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.NType).String(ixgo.DirectCallArg[*q.NType](ctx, 0)))
}

func method_OSABI_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OSABI.GoString(ixgo.DirectCallArg[q.OSABI](ctx, 0)))
}

func method_ptr_OSABI_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OSABI).GoString(ixgo.DirectCallArg[*q.OSABI](ctx, 0)))
}

func method_OSABI_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OSABI.String(ixgo.DirectCallArg[q.OSABI](ctx, 0)))
}

func method_ptr_OSABI_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.OSABI).String(ixgo.DirectCallArg[*q.OSABI](ctx, 0)))
}

func method_ptr_Prog_Open(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Prog).Open(ixgo.DirectCallArg[*q.Prog](ctx, 0)))
}

func method_ProgFlag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ProgFlag.GoString(ixgo.DirectCallArg[q.ProgFlag](ctx, 0)))
}

func method_ptr_ProgFlag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProgFlag).GoString(ixgo.DirectCallArg[*q.ProgFlag](ctx, 0)))
}

func method_ProgFlag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ProgFlag.String(ixgo.DirectCallArg[q.ProgFlag](ctx, 0)))
}

func method_ptr_ProgFlag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProgFlag).String(ixgo.DirectCallArg[*q.ProgFlag](ctx, 0)))
}

func method_ProgType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ProgType.GoString(ixgo.DirectCallArg[q.ProgType](ctx, 0)))
}

func method_ptr_ProgType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProgType).GoString(ixgo.DirectCallArg[*q.ProgType](ctx, 0)))
}

func method_ProgType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ProgType.String(ixgo.DirectCallArg[q.ProgType](ctx, 0)))
}

func method_ptr_ProgType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProgType).String(ixgo.DirectCallArg[*q.ProgType](ctx, 0)))
}

func method_R_386_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_386.GoString(ixgo.DirectCallArg[q.R_386](ctx, 0)))
}

func method_ptr_R_386_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_386).GoString(ixgo.DirectCallArg[*q.R_386](ctx, 0)))
}

func method_R_386_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_386.String(ixgo.DirectCallArg[q.R_386](ctx, 0)))
}

func method_ptr_R_386_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_386).String(ixgo.DirectCallArg[*q.R_386](ctx, 0)))
}

func method_R_390_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_390.GoString(ixgo.DirectCallArg[q.R_390](ctx, 0)))
}

func method_ptr_R_390_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_390).GoString(ixgo.DirectCallArg[*q.R_390](ctx, 0)))
}

func method_R_390_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_390.String(ixgo.DirectCallArg[q.R_390](ctx, 0)))
}

func method_ptr_R_390_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_390).String(ixgo.DirectCallArg[*q.R_390](ctx, 0)))
}

func method_R_AARCH64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_AARCH64.GoString(ixgo.DirectCallArg[q.R_AARCH64](ctx, 0)))
}

func method_ptr_R_AARCH64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_AARCH64).GoString(ixgo.DirectCallArg[*q.R_AARCH64](ctx, 0)))
}

func method_R_AARCH64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_AARCH64.String(ixgo.DirectCallArg[q.R_AARCH64](ctx, 0)))
}

func method_ptr_R_AARCH64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_AARCH64).String(ixgo.DirectCallArg[*q.R_AARCH64](ctx, 0)))
}

func method_R_ALPHA_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_ALPHA.GoString(ixgo.DirectCallArg[q.R_ALPHA](ctx, 0)))
}

func method_ptr_R_ALPHA_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_ALPHA).GoString(ixgo.DirectCallArg[*q.R_ALPHA](ctx, 0)))
}

func method_R_ALPHA_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_ALPHA.String(ixgo.DirectCallArg[q.R_ALPHA](ctx, 0)))
}

func method_ptr_R_ALPHA_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_ALPHA).String(ixgo.DirectCallArg[*q.R_ALPHA](ctx, 0)))
}

func method_R_ARM_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_ARM.GoString(ixgo.DirectCallArg[q.R_ARM](ctx, 0)))
}

func method_ptr_R_ARM_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_ARM).GoString(ixgo.DirectCallArg[*q.R_ARM](ctx, 0)))
}

func method_R_ARM_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_ARM.String(ixgo.DirectCallArg[q.R_ARM](ctx, 0)))
}

func method_ptr_R_ARM_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_ARM).String(ixgo.DirectCallArg[*q.R_ARM](ctx, 0)))
}

func func_R_INFO(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_INFO(ixgo.DirectCallArg[uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func func_R_INFO32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_INFO32(ixgo.DirectCallArg[uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1)))
}

func method_R_LARCH_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_LARCH.GoString(ixgo.DirectCallArg[q.R_LARCH](ctx, 0)))
}

func method_ptr_R_LARCH_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_LARCH).GoString(ixgo.DirectCallArg[*q.R_LARCH](ctx, 0)))
}

func method_R_LARCH_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_LARCH.String(ixgo.DirectCallArg[q.R_LARCH](ctx, 0)))
}

func method_ptr_R_LARCH_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_LARCH).String(ixgo.DirectCallArg[*q.R_LARCH](ctx, 0)))
}

func method_R_MIPS_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_MIPS.GoString(ixgo.DirectCallArg[q.R_MIPS](ctx, 0)))
}

func method_ptr_R_MIPS_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_MIPS).GoString(ixgo.DirectCallArg[*q.R_MIPS](ctx, 0)))
}

func method_R_MIPS_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_MIPS.String(ixgo.DirectCallArg[q.R_MIPS](ctx, 0)))
}

func method_ptr_R_MIPS_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_MIPS).String(ixgo.DirectCallArg[*q.R_MIPS](ctx, 0)))
}

func method_R_PPC_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_PPC.GoString(ixgo.DirectCallArg[q.R_PPC](ctx, 0)))
}

func method_ptr_R_PPC_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_PPC).GoString(ixgo.DirectCallArg[*q.R_PPC](ctx, 0)))
}

func method_R_PPC_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_PPC.String(ixgo.DirectCallArg[q.R_PPC](ctx, 0)))
}

func method_ptr_R_PPC_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_PPC).String(ixgo.DirectCallArg[*q.R_PPC](ctx, 0)))
}

func method_R_PPC64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_PPC64.GoString(ixgo.DirectCallArg[q.R_PPC64](ctx, 0)))
}

func method_ptr_R_PPC64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_PPC64).GoString(ixgo.DirectCallArg[*q.R_PPC64](ctx, 0)))
}

func method_R_PPC64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_PPC64.String(ixgo.DirectCallArg[q.R_PPC64](ctx, 0)))
}

func method_ptr_R_PPC64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_PPC64).String(ixgo.DirectCallArg[*q.R_PPC64](ctx, 0)))
}

func method_R_RISCV_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_RISCV.GoString(ixgo.DirectCallArg[q.R_RISCV](ctx, 0)))
}

func method_ptr_R_RISCV_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_RISCV).GoString(ixgo.DirectCallArg[*q.R_RISCV](ctx, 0)))
}

func method_R_RISCV_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_RISCV.String(ixgo.DirectCallArg[q.R_RISCV](ctx, 0)))
}

func method_ptr_R_RISCV_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_RISCV).String(ixgo.DirectCallArg[*q.R_RISCV](ctx, 0)))
}

func method_R_SPARC_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_SPARC.GoString(ixgo.DirectCallArg[q.R_SPARC](ctx, 0)))
}

func method_ptr_R_SPARC_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_SPARC).GoString(ixgo.DirectCallArg[*q.R_SPARC](ctx, 0)))
}

func method_R_SPARC_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_SPARC.String(ixgo.DirectCallArg[q.R_SPARC](ctx, 0)))
}

func method_ptr_R_SPARC_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_SPARC).String(ixgo.DirectCallArg[*q.R_SPARC](ctx, 0)))
}

func func_R_SYM32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_SYM32(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_R_SYM64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_SYM64(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_R_TYPE32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_TYPE32(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_R_TYPE64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_TYPE64(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func method_R_X86_64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_X86_64.GoString(ixgo.DirectCallArg[q.R_X86_64](ctx, 0)))
}

func method_ptr_R_X86_64_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_X86_64).GoString(ixgo.DirectCallArg[*q.R_X86_64](ctx, 0)))
}

func method_R_X86_64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.R_X86_64.String(ixgo.DirectCallArg[q.R_X86_64](ctx, 0)))
}

func method_ptr_R_X86_64_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.R_X86_64).String(ixgo.DirectCallArg[*q.R_X86_64](ctx, 0)))
}

func func_ST_BIND(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ST_BIND(ixgo.DirectCallArg[uint8](ctx, 0)))
}

func func_ST_INFO(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ST_INFO(ixgo.DirectCallArg[q.SymBind](ctx, 0), ixgo.DirectCallArg[q.SymType](ctx, 1)))
}

func func_ST_TYPE(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ST_TYPE(ixgo.DirectCallArg[uint8](ctx, 0)))
}

func func_ST_VISIBILITY(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ST_VISIBILITY(ixgo.DirectCallArg[uint8](ctx, 0)))
}

func method_ptr_Section_Open(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Section).Open(ixgo.DirectCallArg[*q.Section](ctx, 0)))
}

func method_SectionFlag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SectionFlag.GoString(ixgo.DirectCallArg[q.SectionFlag](ctx, 0)))
}

func method_ptr_SectionFlag_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SectionFlag).GoString(ixgo.DirectCallArg[*q.SectionFlag](ctx, 0)))
}

func method_SectionFlag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SectionFlag.String(ixgo.DirectCallArg[q.SectionFlag](ctx, 0)))
}

func method_ptr_SectionFlag_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SectionFlag).String(ixgo.DirectCallArg[*q.SectionFlag](ctx, 0)))
}

func method_SectionIndex_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SectionIndex.GoString(ixgo.DirectCallArg[q.SectionIndex](ctx, 0)))
}

func method_ptr_SectionIndex_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SectionIndex).GoString(ixgo.DirectCallArg[*q.SectionIndex](ctx, 0)))
}

func method_SectionIndex_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SectionIndex.String(ixgo.DirectCallArg[q.SectionIndex](ctx, 0)))
}

func method_ptr_SectionIndex_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SectionIndex).String(ixgo.DirectCallArg[*q.SectionIndex](ctx, 0)))
}

func method_SectionType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SectionType.GoString(ixgo.DirectCallArg[q.SectionType](ctx, 0)))
}

func method_ptr_SectionType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SectionType).GoString(ixgo.DirectCallArg[*q.SectionType](ctx, 0)))
}

func method_SectionType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SectionType.String(ixgo.DirectCallArg[q.SectionType](ctx, 0)))
}

func method_ptr_SectionType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SectionType).String(ixgo.DirectCallArg[*q.SectionType](ctx, 0)))
}

func method_SymBind_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SymBind.GoString(ixgo.DirectCallArg[q.SymBind](ctx, 0)))
}

func method_ptr_SymBind_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SymBind).GoString(ixgo.DirectCallArg[*q.SymBind](ctx, 0)))
}

func method_SymBind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SymBind.String(ixgo.DirectCallArg[q.SymBind](ctx, 0)))
}

func method_ptr_SymBind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SymBind).String(ixgo.DirectCallArg[*q.SymBind](ctx, 0)))
}

func method_SymType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SymType.GoString(ixgo.DirectCallArg[q.SymType](ctx, 0)))
}

func method_ptr_SymType_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SymType).GoString(ixgo.DirectCallArg[*q.SymType](ctx, 0)))
}

func method_SymType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SymType.String(ixgo.DirectCallArg[q.SymType](ctx, 0)))
}

func method_ptr_SymType_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SymType).String(ixgo.DirectCallArg[*q.SymType](ctx, 0)))
}

func method_SymVis_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SymVis.GoString(ixgo.DirectCallArg[q.SymVis](ctx, 0)))
}

func method_ptr_SymVis_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SymVis).GoString(ixgo.DirectCallArg[*q.SymVis](ctx, 0)))
}

func method_SymVis_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SymVis.String(ixgo.DirectCallArg[q.SymVis](ctx, 0)))
}

func method_ptr_SymVis_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SymVis).String(ixgo.DirectCallArg[*q.SymVis](ctx, 0)))
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

func method_Version_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Version.GoString(ixgo.DirectCallArg[q.Version](ctx, 0)))
}

func method_ptr_Version_GoString(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Version).GoString(ixgo.DirectCallArg[*q.Version](ctx, 0)))
}

func method_Version_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Version.String(ixgo.DirectCallArg[q.Version](ctx, 0)))
}

func method_ptr_Version_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Version).String(ixgo.DirectCallArg[*q.Version](ctx, 0)))
}

func method_VersionIndex_Index(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VersionIndex.Index(ixgo.DirectCallArg[q.VersionIndex](ctx, 0)))
}

func method_ptr_VersionIndex_Index(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.VersionIndex).Index(ixgo.DirectCallArg[*q.VersionIndex](ctx, 0)))
}

func method_VersionIndex_IsHidden(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VersionIndex.IsHidden(ixgo.DirectCallArg[q.VersionIndex](ctx, 0)))
}

func method_ptr_VersionIndex_IsHidden(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.VersionIndex).IsHidden(ixgo.DirectCallArg[*q.VersionIndex](ctx, 0)))
}
