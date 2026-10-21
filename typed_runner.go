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

// TypedRunnerTypeBinder binds interpreter runtime types before installation.
// Returning false rejects the runner. Binding must finish before execution.
type TypedRunnerTypeBinder interface {
	BindTypes(func(types.Type) reflect.Type) bool
}
