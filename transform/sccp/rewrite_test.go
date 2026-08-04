// Copyright (c) 2026 The XGo Authors (xgo.dev). All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sccp

import (
	goConstant "go/constant"
	"testing"

	"github.com/goplus/ixgo"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func TestSCCPRewriteCFGForms(t *testing.T) {
	tests := []struct {
		name   string
		source string
		want   int
	}{
		{
			name: "true edge with phi",
			source: `package test
func target() int {
	value := 1
	if true {
		value = 2
	} else {
		value = 3
	}
	return value
}`,
			want: 2,
		},
		{
			name: "false edge with phi",
			source: `package test
func target() int {
	value := 1
	if false {
		value = 2
	} else {
		value = 3
	}
	return value
}`,
			want: 3,
		},
		{
			name: "direct return",
			source: `package test
func target() int {
	if true {
		return 4
	}
	return 5
}`,
			want: 4,
		},
		{
			name: "nested joins",
			source: `package test
func target() int {
	value := 0
	if true {
		if false {
			value = 1
		} else {
			value = 2
		}
	} else {
		if true {
			value = 3
		} else {
			value = 4
		}
	}
	return value
}`,
			want: 2,
		},
		{
			name: "short circuit",
			source: `package test
func target() int {
	value := 1
	if true && false {
		value = 2
	} else {
		value = 3
	}
	return value
}`,
			want: 3,
		},
		{
			name: "switch chain",
			source: `package test
func target() int {
	value := 2
	switch value {
	case 1:
		return 1
	case 2:
		return 2
	default:
		return 3
	}
}`,
			want: 2,
		},
		{
			name: "switch phi join",
			source: `package test
func target() int {
	value := 0
	switch 2 {
	case 1:
		value = 1
	case 2:
		value = 2
	default:
		value = 3
	}
	return value
}`,
			want: 2,
		},
		{
			name: "zero iteration loop",
			source: `package test
func target() int {
	value := 1
	for false {
		value++
	}
	return value
}`,
			want: 1,
		},
		{
			name: "branch inside loop",
			source: `package test
func target() int {
	sum := 0
	for i := 0; i < 3; i++ {
		if true {
			sum += i
		} else {
			sum += 100
		}
	}
	return sum
}`,
			want: 3,
		},
		{
			name: "unreachable cycle",
			source: `package test
func target() int {
	if false {
		for {
		}
	}
	return 7
}`,
			want: 7,
		},
		{
			name: "labeled break",
			source: `package test
func target() int {
	value := 0
outer:
	for {
		if true {
			value = 6
			break outer
		}
		value = 7
	}
	return value
}`,
			want: 6,
		},
		{
			name: "recover root",
			source: `package test
func target() (result int) {
	defer func() {
		if recover() != nil {
			result = 9
		}
	}()
	if false {
		panic("dead")
	}
	return 8
}`,
			want: 8,
		},
		{
			name: "goto cross edges",
			source: `package test
func target() int {
	value := 0
	if true {
		goto left
	}
	goto right
right:
	value = 9
	goto left
left:
	if false {
		goto right
	}
	value = 5
	return value
}`,
			want: 5,
		},
		{
			name: "closure function",
			source: `package test
func target() int {
	closure := func() int {
		if true {
			return 10
		}
		return 11
	}
	return closure()
}`,
			want: 10,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			original := buildSSAFunction(t, test.source, "target").Pkg
			if got := runSSAInt(t, original); got != test.want {
				t.Fatalf("original target() = %d, want %d", got, test.want)
			}

			target := buildSSAFunction(t, test.source, "target")
			blocksBefore := countPackageBlocks(target.Pkg)
			runSCCPPassForTest(t, target.Pkg)
			if blocksAfter := countPackageBlocks(target.Pkg); blocksAfter >= blocksBefore {
				t.Errorf("blocks after SCCP = %d, want fewer than %d", blocksAfter, blocksBefore)
			}
			for fn := range ssautil.AllFunctions(target.Prog) {
				if fn.Pkg != target.Pkg || fn.Blocks == nil {
					continue
				}
				assertRewrittenSSA(t, fn)
				assertSSAUsesDominated(t, fn)
			}
			if got := runSSAInt(t, target.Pkg); got != test.want {
				t.Fatalf("rewritten target() = %d, want %d", got, test.want)
			}
		})
	}
}

func countPackageBlocks(pkg *ssa.Package) int {
	blocks := 0
	for fn := range ssautil.AllFunctions(pkg.Prog) {
		if fn.Pkg == pkg {
			blocks += len(fn.Blocks)
		}
	}
	return blocks
}

func TestSCCPRewriteInfiniteLoop(t *testing.T) {
	fn := buildSSAFunction(t, `package test
func target() {
	for true {
	}
}`, "target")
	blocksBefore := len(fn.Blocks)
	runSCCPPassForTest(t, fn.Pkg)
	if len(fn.Blocks) >= blocksBefore {
		t.Fatalf("blocks after SCCP = %d, want fewer than %d", len(fn.Blocks), blocksBefore)
	}
	assertRewrittenSSA(t, fn)
	assertSSAUsesDominated(t, fn)
}

func TestSCCPRewriteKeepsDynamicBranches(t *testing.T) {
	fn := buildSSAFunction(t, `package test
func target(condition bool) int {
	if condition {
		return 1
	}
	return 2
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
	assertRewrittenSSA(t, fn)
	assertSSAUsesDominated(t, fn)
	if got := runSSAInt(t, fn.Pkg, true); got != 1 {
		t.Fatalf("target(true) = %d, want 1", got)
	}
	if got := runSSAInt(t, fn.Pkg, false); got != 2 {
		t.Fatalf("target(false) = %d, want 2", got)
	}
}

func runSSAInt(t *testing.T, pkg *ssa.Package, args ...ixgo.Value) int {
	t.Helper()
	interp, err := ixgo.NewContext(0).NewInterp(pkg)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	if err := interp.RunInit(); err != nil {
		t.Fatal(err)
	}
	value, err := interp.RunFunc("target", args...)
	if err != nil {
		t.Fatal(err)
	}
	result, ok := value.(int)
	if !ok {
		t.Fatalf("target result has type %T, want int", value)
	}
	return result
}

func assertRewrittenSSA(t *testing.T, fn *ssa.Function) {
	t.Helper()
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
	visit(fn.Recover)

	instructions := make(map[ssa.Instruction]bool)
	values := make(map[ssa.Value]bool)
	for _, parameter := range fn.Params {
		values[parameter] = true
	}
	for _, freeVar := range fn.FreeVars {
		values[freeVar] = true
	}
	for i, block := range fn.Blocks {
		if !reachable[block] {
			t.Errorf("block %d is unreachable", block.Index)
		}
		if block.Index != i {
			t.Errorf("block index = %d, want %d", block.Index, i)
		}
		for _, instruction := range block.Instrs {
			instructions[instruction] = true
			if value, ok := instruction.(ssa.Value); ok {
				values[value] = true
			}
			for _, operand := range instruction.Operands(nil) {
				if operand != nil && *operand != nil {
					values[*operand] = true
				}
			}
			if instruction.Block() != block {
				t.Errorf("%s reports block %v, want %d", instruction, instruction.Block(), block.Index)
			}
			if phi, ok := instruction.(*ssa.Phi); ok && len(phi.Edges) != len(block.Preds) {
				t.Errorf("block %d phi edges = %d, want %d predecessors", block.Index, len(phi.Edges), len(block.Preds))
			}
		}
		if len(block.Instrs) == 0 {
			t.Errorf("block %d has no terminator", block.Index)
			continue
		}
		switch block.Instrs[len(block.Instrs)-1].(type) {
		case *ssa.If:
			if len(block.Succs) != 2 {
				t.Errorf("If block %d has %d successors", block.Index, len(block.Succs))
			}
		case *ssa.Jump:
			if len(block.Succs) != 1 {
				t.Errorf("Jump block %d has %d successors", block.Index, len(block.Succs))
			}
		case *ssa.Return, *ssa.Panic:
			if len(block.Succs) != 0 {
				t.Errorf("terminal block %d has %d successors", block.Index, len(block.Succs))
			}
		default:
			t.Errorf("block %d has invalid terminator %T", block.Index, block.Instrs[len(block.Instrs)-1])
		}

		for _, successor := range block.Succs {
			if countBlocks(block.Succs, successor) != countBlocks(successor.Preds, block) {
				t.Errorf("CFG edge multiplicity between %d and %d differs", block.Index, successor.Index)
			}
		}
		for _, predecessor := range block.Preds {
			if countBlocks(block.Preds, predecessor) != countBlocks(predecessor.Succs, block) {
				t.Errorf("CFG edge multiplicity between %d and %d differs", predecessor.Index, block.Index)
			}
		}
	}

	expectedReferrers := make(map[ssa.Value]map[ssa.Instruction]int)
	for instruction := range instructions {
		if branch, ok := instruction.(*ssa.If); ok {
			if condition, ok := branch.Cond.(*ssa.Const); ok && condition.Value != nil && condition.Value.Kind() == goConstant.Bool {
				t.Errorf("constant branch remains in block %d", branch.Block().Index)
			}
		}
		for _, operand := range instruction.Operands(nil) {
			if operand == nil || *operand == nil {
				continue
			}
			if (*operand).Referrers() != nil {
				if expectedReferrers[*operand] == nil {
					expectedReferrers[*operand] = make(map[ssa.Instruction]int)
				}
				expectedReferrers[*operand][instruction]++
			}
		}
	}
	for value := range values {
		referrers := value.Referrers()
		if referrers == nil {
			continue
		}
		actual := make(map[ssa.Instruction]int)
		for _, referrer := range *referrers {
			actual[referrer]++
		}
		for referrer, count := range expectedReferrers[value] {
			if actual[referrer] != count {
				t.Errorf("%s referrer %s count = %d, want %d", value, referrer, actual[referrer], count)
			}
			delete(actual, referrer)
		}
		for referrer, count := range actual {
			if !instructions[referrer] {
				t.Errorf("%s retains deleted referrer %s", value, referrer)
			} else {
				t.Errorf("%s has unexpected referrer %s count %d", value, referrer, count)
			}
		}
	}
}

func assertSSAUsesDominated(t *testing.T, fn *ssa.Function) {
	t.Helper()
	for _, block := range fn.Blocks {
		positions := make(map[ssa.Instruction]int, len(block.Instrs))
		for i, instruction := range block.Instrs {
			positions[instruction] = i
		}
		for i, instruction := range block.Instrs {
			if phi, ok := instruction.(*ssa.Phi); ok {
				if len(phi.Edges) != len(block.Preds) {
					t.Errorf("block %d phi edges = %d, want %d predecessors", block.Index, len(phi.Edges), len(block.Preds))
					continue
				}
				for edge, value := range phi.Edges {
					definition, ok := value.(ssa.Instruction)
					if ok && !definition.Block().Dominates(block.Preds[edge]) {
						t.Errorf("%s does not dominate phi predecessor block %d", definition, block.Preds[edge].Index)
					}
				}
				continue
			}
			for _, operand := range instruction.Operands(nil) {
				if operand == nil {
					continue
				}
				definition, ok := (*operand).(ssa.Instruction)
				if !ok {
					continue
				}
				if definition.Parent() != fn {
					t.Errorf("%s uses definition %s from another function", instruction, definition)
					continue
				}
				if definition.Block() == block {
					position, exists := positions[definition]
					if !exists {
						t.Errorf("%s is not present in its reported block %d", definition, block.Index)
					} else if position >= i {
						t.Errorf("%s does not precede its use in block %d", definition, block.Index)
					}
					continue
				}
				// go/ssa models Recover as a second CFG root even though entry
				// allocations, such as named results, are available there.
				if definition.Block() == fn.Blocks[0] && blockReachableFrom(fn.Recover, block) {
					continue
				}
				if !definition.Block().Dominates(block) {
					t.Errorf("%s does not dominate its use %s", definition, instruction)
				}
			}
		}
	}
}

func blockReachableFrom(root, target *ssa.BasicBlock) bool {
	seen := make(map[*ssa.BasicBlock]bool)
	var visit func(*ssa.BasicBlock) bool
	visit = func(block *ssa.BasicBlock) bool {
		if block == nil || seen[block] {
			return false
		}
		if block == target {
			return true
		}
		seen[block] = true
		for _, successor := range block.Succs {
			if visit(successor) {
				return true
			}
		}
		return false
	}
	return visit(root)
}

func countBlocks(blocks []*ssa.BasicBlock, target *ssa.BasicBlock) int {
	count := 0
	for _, block := range blocks {
		if block == target {
			count++
		}
	}
	return count
}

func runSCCPPassForTest(t *testing.T, pkg *ssa.Package) {
	t.Helper()
	if err := (Pass{}).Run(pkg); err != nil {
		t.Fatal(err)
	}
	for fn := range ssautil.AllFunctions(pkg.Prog) {
		if fn.Blocks != nil {
			assertSSASane(t, fn)
		}
	}
}

func buildSSAFunction(t *testing.T, source, name string) *ssa.Function {
	t.Helper()
	fn := buildPackage(t, source).Func(name)
	if fn == nil {
		t.Fatalf("function %q not found", name)
	}
	return fn
}
