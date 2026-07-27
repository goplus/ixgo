// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sccp implements sparse conditional constant propagation for
// golang.org/x/tools/go/ssa. It is a faithful adaptation of the Go
// compiler's implementation in cmd/compile/internal/ssa/sccp.go; the
// algorithm, structure and comments follow the original closely, with
// only the necessary changes to map the compiler's value-based IR
// (Op/AuxInt) onto go/ssa's instruction-based IR (token.Token/constant.Value).
package sccp

import (
	"fmt"

	"go/constant"
	"go/token"
	"go/types"
	_ "unsafe"

	"golang.org/x/tools/go/ssa"
)

//go:linkname ssaOptimizeBlocks golang.org/x/tools/go/ssa.optimizeBlocks
func ssaOptimizeBlocks(*ssa.Function)

//go:linkname ssaBuildReferrers golang.org/x/tools/go/ssa.buildReferrers
func ssaBuildReferrers(*ssa.Function)

//go:linkname ssaBuildDomTree golang.org/x/tools/go/ssa.buildDomTree
func ssaBuildDomTree(*ssa.Function)

// ----------------------------------------------------------------------------
// Sparse Conditional Constant Propagation
//
// Described in
// Mark N. Wegman, F. Kenneth Zadeck: Constant Propagation with Conditional Branches.
// TOPLAS 1991.
//
// This algorithm uses three level lattice for SSA value
//
//      Top        undefined
//     / | \
// .. 1  2  3 ..   constant
//     \ | /
//     Bottom      not constant
//
// It starts with optimistically assuming that all SSA values are initially Top
// and then propagates constant facts only along reachable control flow paths.
// Since some basic blocks are not visited yet, corresponding inputs of phi become
// Top, we use the meet(phi) to compute its lattice.
//
// 	  Top ∩ any = any
// 	  Bottom ∩ any = Bottom
// 	  ConstantA ∩ ConstantA = ConstantA
// 	  ConstantA ∩ ConstantB = Bottom
//
// Each lattice value is lowered most twice(Top to Constant, Constant to Bottom)
// due to lattice depth, resulting in a fast convergence speed of the algorithm.
// In this way, sccp can discover optimization opportunities that cannot be found
// by just combining constant folding and constant propagation and dead code
// elimination separately.

// debug mirrors the compiler's f.pass.debug flag.
var debug = false

// Three level lattice holds compile time knowledge about SSA value
const (
	top      int8 = iota // undefined
	constLat             // constant
	bottom               // not a constant
)

type lattice struct {
	tag int8       // lattice type
	val *ssa.Const // constant value
}

type worklist struct {
	f            *ssa.Function                   // the target function to be optimized out
	edges        []Edge                          // propagate constant facts through edges
	uses         []ssa.Value                     // re-visiting set
	visited      map[Edge]bool                   // visited edges
	latticeCells map[ssa.Value]lattice           // constant lattices
	defUse       map[ssa.Value][]ssa.Value       // def-use chains for some values
	defBlock     map[ssa.Value][]*ssa.BasicBlock // use blocks of def
	visitedBlock []bool                          // visited block
}

// Pass runs sparse conditional constant propagation on every function of
// a built SSA package.
type Pass struct{}

// Run applies sccp to all functions (including anonymous functions and
// declared methods) of pkg.
func (Pass) Run(pkg *ssa.Package) error {
	for _, fn := range functions(pkg) {
		if err := runFunc(fn); err != nil {
			return err
		}
	}
	return nil
}

// functions collects package-level functions, methods of named types and,
// recursively, their anonymous functions.
func functions(pkg *ssa.Package) []*ssa.Function {
	var fns []*ssa.Function
	var add func(f *ssa.Function)
	add = func(f *ssa.Function) {
		fns = append(fns, f)
		for _, anon := range f.AnonFuncs {
			add(anon)
		}
	}
	for _, mem := range pkg.Members {
		switch mem := mem.(type) {
		case *ssa.Function:
			add(mem)
		case *ssa.Type:
			if named, ok := mem.Type().(*types.Named); ok {
				for i := 0; i < named.NumMethods(); i++ {
					if fn := pkg.Prog.FuncValue(named.Method(i)); fn != nil {
						add(fn)
					}
				}
			}
		}
	}
	return fns
}

func runFunc(fn *ssa.Function) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("sccp: %s: %v", fn, r)
		}
	}()
	sccp(fn)
	return nil
}

// sccp stands for sparse conditional constant propagation, it propagates constants
// through CFG conditionally and applies constant folding, constant replacement and
// dead code elimination all together.
func sccp(f *ssa.Function) {
	if len(f.Blocks) == 0 {
		return
	}
	var t worklist
	t.f = f
	t.edges = make([]Edge, 0)
	t.visited = make(map[Edge]bool)
	t.edges = append(t.edges, Edge{f.Blocks[0], 0})
	t.defUse = make(map[ssa.Value][]ssa.Value)
	t.defBlock = make(map[ssa.Value][]*ssa.BasicBlock)
	t.latticeCells = make(map[ssa.Value]lattice)
	t.visitedBlock = make([]bool, len(f.Blocks))

	// build it early since we rely heavily on the def-use chain later
	t.buildDefUses()

	// pick up either an edge or SSA value from worklist, process it
	for {
		if len(t.edges) > 0 {
			edge := t.edges[0]
			t.edges = t.edges[1:]
			if _, exist := t.visited[edge]; !exist {
				dest := edge.b
				destVisited := t.visitedBlock[dest.Index]

				// mark edge as visited
				t.visited[edge] = true
				t.visitedBlock[dest.Index] = true
				for _, instr := range dest.Instrs {
					val, ok := instr.(ssa.Value)
					if !ok {
						continue
					}
					if _, isPhi := instr.(*ssa.Phi); isPhi || !destVisited {
						t.visitValue(val)
					}
				}
				// propagates constants facts through CFG, taking condition test
				// into account
				if !destVisited {
					t.propagate(dest)
				}
			}
			continue
		}
		if len(t.uses) > 0 {
			use := t.uses[0]
			t.uses = t.uses[1:]
			t.visitValue(use)
			continue
		}
		break
	}

	// apply optimizations based on discovered constants
	constCnt, rewireCnt := t.replaceConst()
	if constCnt > 0 || rewireCnt > 0 {
		ssaOptimizeBlocks(f)
		rebuildMetadata(f)
	}
	if debug {
		if constCnt > 0 || rewireCnt > 0 {
			fmt.Printf("Phase SCCP for %v : %v constants, %v dce\n", f, constCnt, rewireCnt)
		}
	}
}

func equals(a, b lattice) bool {
	if a == b {
		// fast path
		return true
	}
	if a.tag != b.tag {
		return false
	}
	if a.tag == constLat {
		// The same content of const value may be different Const objects, we
		// should compare their types and constant values instead (the compiler
		// compares Op and AuxInt here)
		return constEquals(a.val, b.val)
	}
	return true
}

// constEquals reports whether two foldable constants are identical. It plays
// the role of the compiler's `v1.Op == v2.Op && v1.AuxInt == v2.AuxInt`
// comparison: integers compare by their truncated machine value, floats by
// bit pattern.
func constEquals(v1, v2 *ssa.Const) bool {
	if v1 == v2 {
		return true
	}
	if !types.Identical(v1.Type(), v2.Type()) {
		return false
	}
	if v1.Value == nil || v2.Value == nil {
		return v1.Value == v2.Value
	}
	b, ok := basicOf(v1.Type())
	if !ok {
		return false
	}
	switch {
	case b.Info()&types.IsBoolean != 0:
		return constant.BoolVal(v1.Value) == constant.BoolVal(v2.Value)
	case b.Info()&types.IsInteger != 0:
		if b.Info()&types.IsUnsigned != 0 {
			return v1.Uint64() == v2.Uint64()
		}
		return v1.Int64() == v2.Int64()
	case b.Info()&types.IsFloat != 0:
		return float64Bits(v1.Float64()) == float64Bits(v2.Float64())
	}
	return false
}

// possibleConst checks if Value can be folded to const. For those Values that can
// never become constants(e.g. Call), we don't make futile efforts.
func possibleConst(val ssa.Value) bool {
	if isConst(val) {
		return true
	}
	switch val := val.(type) {
	case *ssa.ChangeType:
		// analog of OpCopy: same value with another static type
		return true
	case *ssa.Phi:
		return true
	case *ssa.UnOp:
		switch val.Op {
		case
			// negate
			token.SUB,
			// bitwise complement
			token.XOR,
			// not
			token.NOT:
			return true
		}
		return false
	case *ssa.Convert:
		// conversion, the analog of the compiler's Trunc/ZeroExt/SignExt/Cvt ops
		return foldableConvert(val.X.Type(), val.Type())
	case *ssa.BinOp:
		switch val.Op {
		case
			// add
			token.ADD,
			// sub
			token.SUB,
			// mul
			token.MUL,
			// div
			token.QUO,
			// mod
			token.REM,
			// compare
			token.EQL, token.NEQ, token.LSS, token.LEQ, token.GTR, token.GEQ,
			// shift
			token.SHL, token.SHR,
			// bit
			token.AND, token.OR, token.XOR, token.AND_NOT:
			return true
		}
		return false
	}
	return false
}

func (t *worklist) getLatticeCell(val ssa.Value) lattice {
	if !possibleConst(val) {
		// they are always worst
		return lattice{bottom, nil}
	}
	if c, ok := val.(*ssa.Const); ok {
		// In go/ssa constants are free-floating values rather than block
		// instructions, so they are never visited by visitValue; their
		// lattice is constant from the start.
		return lattice{constLat, c}
	}
	lt, exist := t.latticeCells[val]
	if !exist {
		return lattice{top, nil} // optimistically for un-visited value
	}
	return lt
}

// isConst reports whether val is a constant the pass knows how to fold:
// a *ssa.Const of boolean, (non-untyped) integer or float type. This is the
// analog of the compiler's OpConst64/OpConst32/.../OpConstBool/OpConst64F
// check. TODO: support nil/string/complex constants.
func isConst(val ssa.Value) bool {
	c, ok := val.(*ssa.Const)
	if !ok || c.Value == nil {
		return false
	}
	switch c.Value.Kind() {
	case constant.Bool, constant.Int, constant.Float:
		return foldableType(c.Type())
	}
	return false
}

// buildDefUses builds def-use chain for some values early, because once the
// lattice of a value is changed, we need to update lattices of use. But we don't
// need all uses of it, only uses that can become constants would be added into
// re-visit worklist since no matter how many times they are revisited, uses which
// can't become constants lattice remains unchanged, i.e. Bottom.
func (t *worklist) buildDefUses() {
	for _, block := range t.f.Blocks {
		for _, instr := range block.Instrs {
			val, ok := instr.(ssa.Value)
			if !ok {
				continue
			}
			for _, argp := range instr.Operands(nil) {
				arg := *argp
				if arg == nil {
					continue
				}
				// find its uses, only uses that can become constants take into account
				if possibleConst(arg) && possibleConst(val) {
					t.defUse[arg] = append(t.defUse[arg], val)
				}
			}
		}
		// for control values that can become constants, find their use blocks
		if ctl, ok := controlValue(block); ok {
			if possibleConst(ctl) {
				t.defBlock[ctl] = append(t.defBlock[ctl], block)
			}
		}
	}
}

// addUses finds all uses of value and appends them into work list for further process
func (t *worklist) addUses(val ssa.Value) {
	for _, use := range t.defUse[val] {
		if val == use {
			// Phi may refer to itself as uses, ignore them to avoid re-visiting phi
			// for performance reason
			continue
		}
		t.uses = append(t.uses, use)
	}
	for _, block := range t.defBlock[val] {
		if t.visitedBlock[block.Index] {
			t.propagate(block)
		}
	}
}

// meet meets all of phi arguments and computes result lattice
func (t *worklist) meet(val *ssa.Phi) lattice {
	optimisticLt := lattice{top, nil}
	for i := 0; i < len(val.Edges); i++ {
		edge := Edge{val.Block(), i}
		// If incoming edge for phi is not visited, assume top optimistically.
		// According to rules of meet:
		// 		Top ∩ any = any
		// Top participates in meet() but does not affect the result, so here
		// we will ignore Top and only take other lattices into consideration.
		if _, exist := t.visited[edge]; exist {
			lt := t.getLatticeCell(val.Edges[i])
			if lt.tag == constLat {
				if optimisticLt.tag == top {
					optimisticLt = lt
				} else {
					if !equals(optimisticLt, lt) {
						// ConstantA ∩ ConstantB = Bottom
						return lattice{bottom, nil}
					}
				}
			} else if lt.tag == bottom {
				// Bottom ∩ any = Bottom
				return lattice{bottom, nil}
			} else {
				// Top ∩ any = any
			}
		} else {
			// Top ∩ any = any
		}
	}

	// ConstantA ∩ ConstantA = ConstantA or Top ∩ any = any
	return optimisticLt
}

// computeLattice folds a value whose arguments are all constants and returns
// the resulting lattice.
//
// The compiler reuses its generic rewrite rules (rewriteValuegeneric) on a
// temporary value here; go/ssa has no rewrite rules, so the equivalent
// compile-time evaluation with machine semantics lives in fold.go. Just like
// the generic rules, folding fails (returns Bottom) when the operation does
// not satisfy additional constraints, e.g. division by zero or negative
// shift count, whose panics must be preserved for run time.
func computeLattice(f *ssa.Function, val ssa.Value, args ...*ssa.Const) lattice {
	var constValue *ssa.Const
	switch val := val.(type) {
	case *ssa.UnOp:
		constValue = foldUnOp(val.Op, args[0], val.Type())
	case *ssa.Convert:
		constValue = foldConvert(args[0], val.Type())
	case *ssa.BinOp:
		constValue = foldBinOp(val.Op, args[0], args[1], val.Type())
	}
	if constValue != nil {
		return lattice{constLat, constValue}
	}
	// Either we can not fold the given value or it does not satisfy additional
	// constraints(e.g. divide by zero), treat it as non-constant.
	return lattice{bottom, nil}
}

func (t *worklist) visitValue(val ssa.Value) {
	if !possibleConst(val) {
		// fast fail for always worst Values, i.e. there is no lowering happen
		// on them, their lattices must be initially worse Bottom.
		return
	}

	oldLt := t.getLatticeCell(val)
	defer func() {
		// re-visit all uses of value if its lattice is changed
		newLt := t.getLatticeCell(val)
		if !equals(newLt, oldLt) {
			if oldLt.tag > newLt.tag {
				panic("Must lower lattice")
			}
			t.addUses(val)
		}
	}()

	switch val := val.(type) {
	// Note: unlike the compiler, constants (*ssa.Const) are not block
	// instructions in go/ssa, so there is no OpConst case here; they are
	// handled directly in getLatticeCell.

	// lattice value of copy(x) actually means lattice value of (x)
	case *ssa.ChangeType:
		lt := t.getLatticeCell(val.X)
		if lt.tag == constLat {
			// unlike OpCopy, ChangeType changes the static type, so the
			// constant must be re-typed to the result type
			t.latticeCells[val] = lattice{constLat, ssa.NewConst(lt.val.Value, val.Type())}
		} else {
			t.latticeCells[val] = lattice{lt.tag, nil}
		}
	// phi should be processed specially
	case *ssa.Phi:
		t.latticeCells[val] = t.meet(val)
	// fold 1-input operations:
	case *ssa.UnOp:
		lt1 := t.getLatticeCell(val.X)

		if lt1.tag == constLat {
			// here we take a shortcut by reusing fold.go to fold constants
			t.latticeCells[val] = computeLattice(t.f, val, lt1.val)
		} else {
			t.latticeCells[val] = lattice{lt1.tag, nil}
		}
	// conversion is also a 1-input operation
	case *ssa.Convert:
		lt1 := t.getLatticeCell(val.X)

		if lt1.tag == constLat {
			t.latticeCells[val] = computeLattice(t.f, val, lt1.val)
		} else {
			t.latticeCells[val] = lattice{lt1.tag, nil}
		}
	// fold 2-input operations
	case *ssa.BinOp:
		lt1 := t.getLatticeCell(val.X)
		lt2 := t.getLatticeCell(val.Y)

		if lt1.tag == constLat && lt2.tag == constLat {
			// here we take a shortcut by reusing fold.go to fold constants
			t.latticeCells[val] = computeLattice(t.f, val, lt1.val, lt2.val)
		} else {
			if lt1.tag == bottom || lt2.tag == bottom {
				t.latticeCells[val] = lattice{bottom, nil}
			} else {
				t.latticeCells[val] = lattice{top, nil}
			}
		}
	default:
		// Any other type of value cannot be a constant, they are always worst(Bottom)
	}
}

// propagate propagates constants facts through CFG. If the block has single successor,
// add the successor anyway. If the block has multiple successors, only add the
// branch destination corresponding to lattice value of condition value.
func (t *worklist) propagate(block *ssa.BasicBlock) {
	if len(block.Instrs) == 0 {
		return
	}
	switch instr := block.Instrs[len(block.Instrs)-1].(type) {
	case *ssa.Return, *ssa.Panic:
		// control flow ends, do nothing then
		break
	case *ssa.Jump:
		// the analog of BlockPlain: always takes the single branch
		t.edges = append(t.edges, succEdge(block, 0))
	case *ssa.If:
		cond := instr.Cond
		condLattice := t.getLatticeCell(cond)
		if condLattice.tag == bottom {
			// we know nothing about control flow, add all branch destinations
			t.edges = append(t.edges, succEdge(block, 0), succEdge(block, 1))
		} else if condLattice.tag == constLat {
			// add branchIdx destinations depends on its condition
			var branchIdx int
			if constant.BoolVal(condLattice.val.Value) {
				branchIdx = 0
			} else {
				branchIdx = 1
			}
			t.edges = append(t.edges, succEdge(block, branchIdx))
		} else {
			// condition value is not visited yet, don't propagate it now
		}
	default:
		// unknown terminator, conservatively add all branch destinations
		for i := range block.Succs {
			t.edges = append(t.edges, succEdge(block, i))
		}
	}
}

// rewireSuccessor rewires corresponding successors according to constant value
// discovered by previous analysis. As the result, some successors become unreachable
// and thus can be removed in further deadcode phase
func (t *worklist) rewireSuccessor(block *ssa.BasicBlock, constVal *ssa.Const) bool {
	n := len(block.Instrs)
	if n == 0 {
		return false
	}
	switch instr := block.Instrs[n-1].(type) {
	case *ssa.If:
		// Succs[0] is taken when cond is true, so remove the other edge; this
		// matches the compiler's block.removeEdge(int(constVal.AuxInt))
		if constant.BoolVal(constVal.Value) {
			t.removeEdge(block, 1)
		} else {
			t.removeEdge(block, 0)
		}
		// the analog of Kind = BlockPlain + ResetControls: replace the If
		// terminator with a Jump
		deleteInstr(instr)
		block.Instrs = append(block.Instrs, newJump(block))
		return true
	default:
		return false
	}
}

// replaceConst will replace non-constant values that have been proven by sccp
// to be constants.
func (t *worklist) replaceConst() (int, int) {
	constCnt, rewireCnt := 0, 0
	for val, lt := range t.latticeCells {
		if lt.tag == constLat {
			if !isConst(val) {
				if debug {
					fmt.Printf("Replace %v with %v\n", val, lt.val)
				}
				replaceUses(val, lt.val)
				deleteInstr(val.(ssa.Instruction))
				constCnt++
			}
			// If const value controls this block, rewires successors according to its value
			ctrlBlock := t.defBlock[val]
			for _, block := range ctrlBlock {
				if t.rewireSuccessor(block, lt.val) {
					rewireCnt++
					if debug {
						fmt.Printf("Rewire %v successors\n", block)
					}
				}
			}
		}
	}
	// In go/ssa, constants are not block instructions and thus never enter
	// latticeCells, so branches controlled by a literal constant (e.g. if
	// true {...}) are rewired here.
	for val, blocks := range t.defBlock {
		if c, ok := val.(*ssa.Const); ok && isConst(c) {
			for _, block := range blocks {
				if t.rewireSuccessor(block, c) {
					rewireCnt++
					if debug {
						fmt.Printf("Rewire %v successors\n", block)
					}
				}
			}
		}
	}
	return constCnt, rewireCnt
}
