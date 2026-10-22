/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package ixgo

import (
	"go/types"
	"golang.org/x/tools/go/ssa"
	"reflect"
)

// TypedRunner executes a statically compiled function without interpreter frames.
type TypedRunner interface {
	Run(Value) Value
}

// TypedArgumentsRunner executes a statically compiled function with all arguments.
// RunArgs requires exactly the compiled parameter count. Each argument must have
// the exact built-in Go dynamic type of its parameter (not a named type or a
// merely convertible numeric type). Violating these preconditions may panic.
// Implementations may retain the argument slice after returning.
type TypedArgumentsRunner interface {
	TypedRunner
	RunArgs([]Value) Value
}

// TypedRunnerFactory builds an optional runner for an SSA function.
type TypedRunnerFactory func(*ssa.Function) TypedRunner

var mixedRunnerFactory TypedRunnerFactory

// RegisterMixedRunner installs the optional mixed runner factory.
func RegisterMixedRunner(factory TypedRunnerFactory) {
	mixedRunnerFactory = factory
}

type mixedCallPatch struct {
	caller *function
	callee *function
	index  int
	args   []register
	dests  []register
}

func (visit *visitor) deferMixedCall(pfn *function, call *ssa.Call, index int) {
	ctx := visit.intp.ctx
	if ctx.Mode&EnableTracing != 0 || ctx.evalMode || ctx.evalCallFn != nil || ctx.debugFunc != nil || ctx.entryFunc != nil || ctx.leaveFunc != nil {
		return
	}
	callee := call.Common().StaticCallee()
	if callee == nil || callee.Blocks == nil {
		return
	}
	ifn := visit.intp.funcs[callee]
	if ifn == nil || ifn.Recover != nil || ifn.canRunMixed(len(call.Common().Args), nil) {
		return
	}
	_, args, bindings := getCallIndex(pfn, call.Common())
	args = append(args, bindings...)
	var dests []register
	switch ifn.nres {
	case 0:
	case 1:
		dests = []register{pfn.regIndex(call)}
	default:
		dests = make([]register, ifn.nres)
		for j := range dests {
			dests[j] = -1
		}
		if refs := call.Referrers(); refs != nil {
			for _, ref := range *refs {
				switch ex := ref.(type) {
				case *ssa.Extract:
					if _, elided := pfn.elidedExtract[ex]; !elided {
						// The original instruction still produces a tuple for Extract.
						return
					}
					if ex.Referrers() != nil && len(*ex.Referrers()) > 0 {
						dests[ex.Index] = pfn.regIndex(ex)
					}
				case *ssa.DebugRef:
				default:
					return
				}
			}
		}
	}
	visit.mixedCalls = append(visit.mixedCalls, mixedCallPatch{pfn, ifn, index, args, dests})
}

// Recursive callees finish before their ancestors install memo wrappers.
// Wait for the traversal to unwind before checking installation-only flags.
func (visit *visitor) installMixedRunners() {
	pending := visit.mixedPending
	visit.mixedPending = nil
	for _, pfn := range pending {
		if mixedGraphUsesMemo(visit.intp, pfn.Fn, make(map[*ssa.Function]bool)) {
			continue
		}
		pfn.mixedRunner = mixedRunnerFactory(pfn.Fn)
		if binder, ok := pfn.mixedRunner.(TypedRunnerTypeBinder); ok && !binder.BindTypes(visit.intp.toType) {
			pfn.mixedRunner = nil
		}
	}
	for _, patch := range visit.mixedCalls {
		// Memo installation may have wrapped the caller instruction.
		if patch.caller.memoFlags == 0 && patch.callee.Recover == nil && patch.callee.canRunMixed(len(patch.args), nil) {
			if run := makeCallSpecializerInstr(visit.intp, patch.callee, patch.args, patch.dests); run != nil {
				patch.caller.Instrs[patch.index] = run
			}
		}
	}
	visit.mixedCalls = nil
}

func mixedGraphUsesMemo(interp *Interp, fn *ssa.Function, seen map[*ssa.Function]bool) bool {
	if interp == nil || fn == nil || seen[fn] {
		return false
	}
	seen[fn] = true
	if pfn := interp.funcs[fn]; pfn != nil && pfn.memoFlags != 0 {
		return true
	}
	for _, block := range fn.Blocks {
		for _, instr := range block.Instrs {
			if call, ok := instr.(*ssa.Call); ok {
				if callee := call.Common().StaticCallee(); callee != nil && mixedGraphUsesMemo(interp, callee, seen) {
					return true
				}
			}
		}
	}
	return false
}

// canRunMixed also excludes closure environments from custom factories.
func (pfn *function) canRunMixed(narg int, env []value) bool {
	if pfn.mixedRunner == nil || pfn.nenv != 0 || env != nil || narg != pfn.narg {
		return false
	}
	_, multiple := pfn.mixedRunner.(TypedArgumentsRunner)
	return multiple || narg == 1 && pfn.nres == 1
}

func (pfn *function) runMixed(args []value, env []value) (value, bool) {
	if !pfn.canRunMixed(len(args), env) {
		return nil, false
	}
	if runner, ok := pfn.mixedRunner.(TypedArgumentsRunner); ok {
		return runner.RunArgs(args), true
	}
	return pfn.mixedRunner.Run(args[0]), true
}

// runMixedBorrowed keeps the caller's argument slice private. Public runners
// may retain RunArgs slices after returning.
func (pfn *function) runMixedBorrowed(args []value, env []value) (value, bool) {
	if !pfn.canRunMixed(len(args), env) {
		return nil, false
	}
	if runner, ok := pfn.mixedRunner.(TypedArgumentsRunner); ok {
		values := make([]value, len(args))
		copy(values, args)
		return runner.RunArgs(values), true
	}
	return pfn.mixedRunner.Run(args[0]), true
}

// TypedRunnerTypeBinder binds interpreter runtime types before installation.
// Returning false rejects the runner. Binding must finish before execution.
type TypedRunnerTypeBinder interface {
	BindTypes(func(types.Type) reflect.Type) bool
}
