package xgobuild

import (
	"testing"

	"github.com/goplus/ixgo"
)

func TestModeExtensionBits(t *testing.T) {
	// Existing interpreter flags retain their values when new modes are added.
	if ixgo.EnableCachedReg != 1<<14 {
		t.Fatalf("EnableCachedReg changed: %#x", ixgo.EnableCachedReg)
	}
	modes := []ixgo.Mode{
		ixgo.DisableRecover, ixgo.DisableCustomBuiltin, ixgo.EnableDumpImports,
		ixgo.EnableDumpInstr, ixgo.EnableTracing, ixgo.EnablePrintAny,
		ixgo.EnableNoStrict, ixgo.ExperimentalSupportGC, ixgo.SupportMultipleInterp,
		ixgo.CheckGopOverloadFunc, ixgo.DisableAutoLoadPatchs,
		ixgo.DisableDynamicFuncCallAnalysis, ixgo.OptionLoadRutimeImethod,
		ixgo.OptionLoadAllImethod, ixgo.EnableCachedReg, ixgo.DisableMemo,
		ixgo.DisableMixedRunner,
	}
	var used ixgo.Mode
	for _, mode := range modes {
		if mode == 0 || mode&(mode-1) != 0 || used&mode != 0 {
			t.Fatalf("invalid or overlapping mode: %#x", mode)
		}
		if mode > ixgo.LastMode {
			t.Fatalf("mode %#x exceeds LastMode %#x", mode, ixgo.LastMode)
		}
		used |= mode
	}
	if used&StaticLoad != 0 || used&NormalizeExport != 0 || StaticLoad&NormalizeExport != 0 {
		t.Fatalf("extension modes overlap: core=%#x, static=%#x, normalize=%#x", used, StaticLoad, NormalizeExport)
	}
}
