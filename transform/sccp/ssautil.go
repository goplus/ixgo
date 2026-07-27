// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sccp

import (
	"fmt"
	"reflect"
	"unsafe"

	"golang.org/x/tools/go/ssa"
)

// This file supplies the small pieces of compiler SSA infrastructure that the
// original sccp.go relies on (Edge, Block.removeEdge, Block.removePhiArg,
// Value.reset/replacement) on top of go/ssa.

// Edge represents a CFG edge, mirroring cmd/compile's ssa.Edge: b is the
// destination block and i is the index of this edge in b.Preds.
type Edge struct {
	b *ssa.BasicBlock
	i int
}

func rebuildMetadata(fn *ssa.Function) {
	clear := func(value ssa.Value) {
		if value == nil {
			return
		}
		if referrers := value.Referrers(); referrers != nil {
			*referrers = nil
		}
	}
	for _, parameter := range fn.Params {
		clear(parameter)
	}
	for _, freeVar := range fn.FreeVars {
		clear(freeVar)
	}
	for _, block := range fn.Blocks {
		for _, instruction := range block.Instrs {
			if value, ok := instruction.(ssa.Value); ok {
				clear(value)
			}
			for _, operand := range instruction.Operands(nil) {
				if operand != nil {
					clear(*operand)
				}
			}
		}
	}
	ssaBuildReferrers(fn)
	ssaBuildDomTree(fn)
}

// controlValue returns the control value of a block, i.e. the condition of
// its trailing If instruction, mirroring Block.ControlValues() of a BlockIf.
func controlValue(block *ssa.BasicBlock) (ssa.Value, bool) {
	if n := len(block.Instrs); n > 0 {
		if instr, ok := block.Instrs[n-1].(*ssa.If); ok {
			return instr.Cond, true
		}
	}
	return nil, false
}

// succEdge returns the Edge for b.Succs[i]. go/ssa does not record the
// pred-index of an edge, so it is recovered by occurrence counting: edges are
// appended pairwise to Succs and Preds in creation order, hence the k-th
// occurrence of c in b.Succs corresponds to the k-th occurrence of b in
// c.Preds (this matters when both successors of an If are the same block).
func succEdge(b *ssa.BasicBlock, i int) Edge {
	c := b.Succs[i]
	rank := 0
	for k := 0; k < i; k++ {
		if b.Succs[k] == c {
			rank++
		}
	}
	for j, p := range c.Preds {
		if p == b {
			if rank == 0 {
				return Edge{c, j}
			}
			rank--
		}
	}
	panic(fmt.Sprintf("sccp: no pred edge for %v.Succs[%d]", b, i))
}

// removeEdge removes the i'th outgoing edge from b (and the corresponding
// incoming edge from b.Succs[i]), mirroring cmd/compile's Block.removeEdge.
func (t *worklist) removeEdge(b *ssa.BasicBlock, i int) {
	e := succEdge(b, i)
	c, j := e.b, e.i
	// Adjust b.Succs
	b.Succs = append(b.Succs[:i], b.Succs[i+1:]...)
	// Adjust c.Preds
	c.Preds = append(c.Preds[:j], c.Preds[j+1:]...)
	// Remove phi args from c's phis.
	for _, instr := range c.Instrs {
		phi, ok := instr.(*ssa.Phi)
		if !ok {
			break // phis are grouped at the start of the block
		}
		t.removePhiArg(phi, j)
	}
}

// removePhiArg removes the i'th arg from phi, mirroring cmd/compile's
// Block.removePhiArg. If exactly one arg remains, the phi degenerates into
// the compiler's OpCopy; since go/ssa has no copy instruction, the phi is
// eliminated by substituting the remaining value into all its uses.
func (t *worklist) removePhiArg(phi *ssa.Phi, i int) {
	arg := phi.Edges[i]
	phi.Edges = append(phi.Edges[:i], phi.Edges[i+1:]...)
	removeReferrer(arg, phi)
	if len(phi.Edges) == 1 {
		w := phi.Edges[0]
		// If w itself was already proven constant (and thus may already have
		// been removed from its block by replaceConst), substitute the
		// constant instead to avoid resurrecting a deleted instruction.
		if lt, ok := t.latticeCells[w]; ok && lt.tag == constLat && !isConst(w) {
			w = lt.val
		}
		if w != phi {
			replaceUses(phi, w)
		}
		deleteInstr(phi)
	}
}

// replaceUses rewrites all uses of old to new, maintaining referrer lists.
// It is the go/ssa analog of the compiler's in-place Value.reset to a const op.
func replaceUses(old, new ssa.Value) {
	refs := old.Referrers()
	if refs == nil {
		return
	}
	for _, instr := range *refs {
		for _, rand := range instr.Operands(nil) {
			if *rand == old {
				*rand = new
			}
		}
		if nrefs := new.Referrers(); nrefs != nil {
			*nrefs = append(*nrefs, instr)
		}
	}
	*refs = (*refs)[:0]
}

// deleteInstr unlinks instr from its block and from the referrer lists of its
// operands. It is a no-op if instr has already been deleted.
func deleteInstr(instr ssa.Instruction) {
	b := instr.Block()
	idx := -1
	for i, ins := range b.Instrs {
		if ins == instr {
			idx = i
			break
		}
	}
	if idx < 0 {
		return // already deleted
	}
	b.Instrs = append(b.Instrs[:idx], b.Instrs[idx+1:]...)
	for _, rand := range instr.Operands(nil) {
		if *rand != nil {
			removeReferrer(*rand, instr)
		}
	}
}

// removeReferrer removes one occurrence of instr from v's referrer list.
func removeReferrer(v ssa.Value, instr ssa.Instruction) {
	refs := v.Referrers()
	if refs == nil {
		return
	}
	for i, ref := range *refs {
		if ref == instr {
			*refs = append((*refs)[:i], (*refs)[i+1:]...)
			return
		}
	}
}

// newJump creates a Jump terminator for block. go/ssa provides no public
// constructor for instructions, so the unexported anInstruction.block field
// is set through reflection.
func newJump(block *ssa.BasicBlock) *ssa.Jump {
	jump := new(ssa.Jump)
	setBlock(jump, block)
	return jump
}

func setBlock(instr ssa.Instruction, block *ssa.BasicBlock) {
	v := reflect.ValueOf(instr).Elem().FieldByName("block")
	if !v.IsValid() {
		panic("sccp: ssa.Instruction has no block field")
	}
	reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem().
		Set(reflect.ValueOf(block))
}
