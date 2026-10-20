/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package ixgo

import "golang.org/x/tools/go/ssa"

// TypedRunner executes a statically compiled function without interpreter frames.
type TypedRunner interface {
	Run(Value) Value
}

// TypedArgumentsRunner executes a statically compiled function with all arguments.
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

type mixedMemoProbe struct{ enabled bool }

func (p *mixedMemoProbe) InstallScalarMemo()                    { p.enabled = true }
func (p *mixedMemoProbe) InstallReceiverMemo(*MemoReceiverPlan) { p.enabled = true }
func (p *mixedMemoProbe) InstallAccumulatorMemo(int, bool)      { p.enabled = true }

func mixedGraphUsesMemo(fn *ssa.Function, seen map[*ssa.Function]bool) bool {
	if memoPlugin == nil || fn == nil || seen[fn] {
		return false
	}
	seen[fn] = true
	var probe mixedMemoProbe
	memoPlugin(&probe, fn)
	if probe.enabled {
		return true
	}
	for _, block := range fn.Blocks {
		for _, instr := range block.Instrs {
			if call, ok := instr.(*ssa.Call); ok {
				if callee := call.Common().StaticCallee(); callee != nil && mixedGraphUsesMemo(callee, seen) {
					return true
				}
			}
		}
	}
	return false
}
