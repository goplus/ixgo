// Copyright (c) 2026 The XGo Authors (xgo.dev). All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sccp

import (
	"testing"

	"golang.org/x/tools/go/ssa"
)

func TestSCCPKeepsUnknownBranchesExecutable(t *testing.T) {
	fn := buildSSAFunction(t, `package test
func target(condition bool) int {
	key := 1
	if condition {
		key = 2
	}
	return key
}`, "target")
	runSCCPPassForTest(t, fn.Pkg)

	branches := 0
	for _, block := range fn.Blocks {
		for _, instruction := range block.Instrs {
			if _, ok := instruction.(*ssa.If); ok {
				branches++
			}
		}
	}
	if branches != 1 {
		t.Fatalf("dynamic branches after SCCP = %d, want 1", branches)
	}
}

func TestSCCPFoldsUnaryConversionAndBinaryOperations(t *testing.T) {
	fn := buildSSAFunction(t, `package test
func target(condition bool) int64 {
	value := int8(7)
	if condition {
		value = 7
	}
	return int64(-value) + 2
}`, "target")
	runSCCPPassForTest(t, fn.Pkg)

	returns := reachableReturns(fn)
	if len(returns) != 1 {
		t.Fatalf("reachable returns after SCCP = %d, want 1", len(returns))
	}
	value, ok := returns[0].Results[0].(*ssa.Const)
	if !ok || value.String() != "-5:int64" {
		t.Fatalf("reachable return = %v, want -5", returns[0].Results[0])
	}
}

func TestSCCPPassOptimizesConstantDeadBranch(t *testing.T) {
	fn := buildSSAFunction(t, `package test
func dead() {}
func target() int {
	key := 1
	if false {
		dead()
		key = 2
	}
	return key
}`, "target")
	runSCCPPassForTest(t, fn.Pkg)

	if got := len(fn.Blocks); got != 1 {
		t.Fatalf("blocks after SCCP = %d, want 1", got)
	}
	instructions := fn.Blocks[0].Instrs
	if len(instructions) != 1 {
		t.Fatalf("entry instructions after SCCP = %d, want only Return", len(instructions))
	}
}

func reachableReturns(fn *ssa.Function) []*ssa.Return {
	reachable := make(map[*ssa.BasicBlock]bool)
	var visit func(*ssa.BasicBlock)
	visit = func(block *ssa.BasicBlock) {
		if block == nil || reachable[block] {
			return
		}
		reachable[block] = true
		for _, successor := range block.Succs {
			visit(successor)
		}
	}
	visit(fn.Blocks[0])

	var returns []*ssa.Return
	for block := range reachable {
		for _, instruction := range block.Instrs {
			if ret, ok := instruction.(*ssa.Return); ok {
				returns = append(returns, ret)
			}
		}
	}
	return returns
}

func findPhi(t *testing.T, fn *ssa.Function) *ssa.Phi {
	t.Helper()
	for _, block := range fn.Blocks {
		for _, instruction := range block.Instrs {
			if phi, ok := instruction.(*ssa.Phi); ok {
				return phi
			}
		}
	}
	t.Fatal("phi not found")
	return nil
}
