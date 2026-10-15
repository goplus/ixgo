/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package ixgo

import (
	"go/constant"
	"go/token"
	"go/types"
	"math"
	"reflect"

	"golang.org/x/tools/go/ssa"
)

// TypedScalar is the scalar set supported by the generic typed runner.
type TypedScalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

const (
	maxTypedResults    = 8
	maxTypedScalarRegs = 32
)

type typedResults[T TypedScalar] struct {
	values [maxTypedResults]T
	n      uint8
}

type typedOpKind uint8

const (
	typedConst typedOpKind = iota
	typedAdd
	typedSub
	typedMul
	typedQuo
	typedRem
	typedAnd
	typedOr
	typedXor
	typedAndNot
	typedShl
	typedShr
	typedLess
	typedLessEqual
	typedGreater
	typedGreaterEqual
	typedEqual
	typedNotEqual
	typedCall
	typedBranch
	typedJump
	typedReturn
	typedMove
	typedBoolMove
	typedBoolConst
	typedNegate
	typedComplement
	typedNot
	typedBoolEqual
	typedBoolNotEqual
)

type typedOp[T TypedScalar] struct {
	kind      typedOpKind
	dst, x, y uint8
	nres      uint8
	results   []uint8
	value     T
	truePC    int
	falsePC   int
	target    int
}

// typedRunner owns immutable code for a scalar SSA subset. Registers are local
// to each run, so recursive calls need no boxed interpreter frame.
type typedRunner[T TypedScalar] struct {
	code      []typedOp[T]
	recursive bool
}

// TypedRunner executes a function compiled for the typed scalar subset.
// Implementations are normally installed by the typedrunner plugin.
type TypedRunner interface {
	Run(Value) Value
}

// TypedRunnerFactory builds a typed runner for an SSA function. Returning nil
// asks the interpreter to use its normal boxed frame path.
type TypedRunnerFactory func(*ssa.Function) TypedRunner

var typedRunnerFactory TypedRunnerFactory

// RegisterTypedRunnerFactory installs the optional typed runner factory.
// Passing nil disables typed runner compilation.
func RegisterTypedRunnerFactory(factory TypedRunnerFactory) {
	typedRunnerFactory = factory
}

func (r *typedRunner[T]) Run(input Value) Value {
	x, ok := input.(T)
	if !ok {
		return nil
	}
	result := r.runSelectedTyped(x)
	if result.n == 0 {
		return nil
	}
	if result.n == 1 {
		return result.values[0]
	}
	ret := make(tuple, result.n)
	for i := range ret {
		ret[i] = result.values[i]
	}
	return ret
}

func (r *typedRunner[T]) runSelectedTyped(input T) typedResults[T] {
	// Select once per entry; keep the ordinary execution loop unchanged.
	if r.recursive {
		return r.runRecursiveTyped(input)
	}
	return r.runTyped(input)
}

func (r *typedRunner[T]) runTyped(input T) typedResults[T] {
	var regs [maxTypedScalarRegs]T
	var bools [maxTypedScalarRegs]bool
	regs[0] = input
	for pc := 0; ; pc++ {
		op := &r.code[pc]
		switch op.kind {
		case typedConst:
			regs[op.dst] = op.value
		case typedBoolConst:
			bools[op.dst] = op.value != 0
		case typedMove:
			regs[op.dst] = regs[op.x]
		case typedBoolMove:
			bools[op.dst] = bools[op.x]
		case typedAdd:
			regs[op.dst] = regs[op.x] + regs[op.y]
		case typedSub:
			regs[op.dst] = regs[op.x] - regs[op.y]
		case typedMul:
			regs[op.dst] = regs[op.x] * regs[op.y]
		case typedQuo:
			regs[op.dst] = regs[op.x] / regs[op.y]
		case typedRem:
			regs[op.dst] = typedRemainder(regs[op.x], regs[op.y])
		case typedAnd:
			regs[op.dst] = T(uint64(regs[op.x]) & uint64(regs[op.y]))
		case typedOr:
			regs[op.dst] = T(uint64(regs[op.x]) | uint64(regs[op.y]))
		case typedXor:
			regs[op.dst] = T(uint64(regs[op.x]) ^ uint64(regs[op.y]))
		case typedAndNot:
			regs[op.dst] = T(uint64(regs[op.x]) &^ uint64(regs[op.y]))
		case typedShl:
			regs[op.dst] = typedShiftLeft(regs[op.x], regs[op.y])
		case typedShr:
			regs[op.dst] = typedShiftRight(regs[op.x], regs[op.y])
		case typedLess:
			bools[op.dst] = regs[op.x] < regs[op.y]
		case typedLessEqual:
			bools[op.dst] = regs[op.x] <= regs[op.y]
		case typedGreater:
			bools[op.dst] = regs[op.x] > regs[op.y]
		case typedGreaterEqual:
			bools[op.dst] = regs[op.x] >= regs[op.y]
		case typedEqual:
			bools[op.dst] = regs[op.x] == regs[op.y]
		case typedNotEqual:
			bools[op.dst] = regs[op.x] != regs[op.y]
		case typedNegate:
			regs[op.dst] = -regs[op.x]
		case typedComplement:
			regs[op.dst] = T(^uint64(regs[op.x]))
		case typedNot:
			bools[op.dst] = !bools[op.x]
		case typedBoolEqual:
			bools[op.dst] = bools[op.x] == bools[op.y]
		case typedBoolNotEqual:
			bools[op.dst] = bools[op.x] != bools[op.y]
		case typedCall:
			result := r.runTyped(regs[op.x])
			for i := 0; i < int(result.n); i++ {
				regs[op.dst+uint8(i)] = result.values[i]
			}
		case typedBranch:
			if bools[op.x] {
				pc = op.truePC - 1
			} else {
				pc = op.falsePC - 1
			}
		case typedJump:
			pc = op.target - 1
		case typedReturn:
			var result typedResults[T]
			result.n = op.nres
			for i, reg := range op.results {
				result.values[i] = regs[reg]
			}
			return result
		}
	}
}

func (r *typedRunner[T]) runRecursiveTyped(input T) typedResults[T] {
	var result typedResults[T]
	var output [maxTypedScalarRegs]T
	result.n = r.runTypedInto(input, &output, 0)
	copy(result.values[:], output[:result.n])
	return result
}

// runTypedInto writes directly into its caller's result registers. Recursive
// calls avoid returning and copying the maximum-size result structure.
func (r *typedRunner[T]) runTypedInto(input T, output *[maxTypedScalarRegs]T, base uint8) uint8 {
	var regs [maxTypedScalarRegs]T
	var bools [maxTypedScalarRegs]bool
	regs[0] = input
	for pc := 0; ; pc++ {
		op := &r.code[pc]
		switch op.kind {
		case typedConst:
			regs[op.dst] = op.value
		case typedBoolConst:
			bools[op.dst] = op.value != 0
		case typedMove:
			regs[op.dst] = regs[op.x]
		case typedBoolMove:
			bools[op.dst] = bools[op.x]
		case typedAdd:
			regs[op.dst] = regs[op.x] + regs[op.y]
		case typedSub:
			regs[op.dst] = regs[op.x] - regs[op.y]
		case typedMul:
			regs[op.dst] = regs[op.x] * regs[op.y]
		case typedQuo:
			regs[op.dst] = regs[op.x] / regs[op.y]
		case typedRem:
			regs[op.dst] = typedRemainder(regs[op.x], regs[op.y])
		case typedAnd:
			regs[op.dst] = T(uint64(regs[op.x]) & uint64(regs[op.y]))
		case typedOr:
			regs[op.dst] = T(uint64(regs[op.x]) | uint64(regs[op.y]))
		case typedXor:
			regs[op.dst] = T(uint64(regs[op.x]) ^ uint64(regs[op.y]))
		case typedAndNot:
			regs[op.dst] = T(uint64(regs[op.x]) &^ uint64(regs[op.y]))
		case typedShl:
			regs[op.dst] = typedShiftLeft(regs[op.x], regs[op.y])
		case typedShr:
			regs[op.dst] = typedShiftRight(regs[op.x], regs[op.y])
		case typedLess:
			bools[op.dst] = regs[op.x] < regs[op.y]
		case typedLessEqual:
			bools[op.dst] = regs[op.x] <= regs[op.y]
		case typedGreater:
			bools[op.dst] = regs[op.x] > regs[op.y]
		case typedGreaterEqual:
			bools[op.dst] = regs[op.x] >= regs[op.y]
		case typedEqual:
			bools[op.dst] = regs[op.x] == regs[op.y]
		case typedNotEqual:
			bools[op.dst] = regs[op.x] != regs[op.y]
		case typedNegate:
			regs[op.dst] = -regs[op.x]
		case typedComplement:
			regs[op.dst] = T(^uint64(regs[op.x]))
		case typedNot:
			bools[op.dst] = !bools[op.x]
		case typedBoolEqual:
			bools[op.dst] = bools[op.x] == bools[op.y]
		case typedBoolNotEqual:
			bools[op.dst] = bools[op.x] != bools[op.y]
		case typedCall:
			r.runTypedInto(regs[op.x], &regs, op.dst)
		case typedBranch:
			if bools[op.x] {
				pc = op.truePC - 1
			} else {
				pc = op.falsePC - 1
			}
		case typedJump:
			pc = op.target - 1
		case typedReturn:
			for i, reg := range op.results {
				output[int(base)+i] = regs[reg]
			}
			return op.nres
		}
	}
}

// MakeTypedRunnerOf compiles an SSA function for one scalar type. A plugin
// chooses the concrete T after inspecting the function signature. The core
// selects the scalar fast path first, then the argument and memory runner.
// The argument and memory runner applies instruction fusion and batched constant
// initialization throughout supported recursive call graphs before publication.
// Unsupported functions return nil for ordinary interpreter execution.
func MakeTypedRunnerOf[T TypedScalar](fn *ssa.Function) TypedRunner {
	if runner := makeTypedScalarRunnerOf[T](fn); runner != nil {
		return runner
	}
	runner := makeTypedArgumentsRunnerOf[T](fn)
	if runner != nil {
		r := runner.(*typedArgsRunner[T])
		fuseTypedArguments(r)
		batchTypedConstants(r)
	}
	return runner
}

func makeTypedScalarRunnerOf[T TypedScalar](fn *ssa.Function) TypedRunner {
	if fn == nil || fn.Signature == nil || fn.Recover != nil || len(fn.FreeVars) != 0 ||
		fn.Signature.Params().Len() < 1 || fn.Signature.Results().Len() > maxTypedResults ||
		!typedBoundaryTypeIs[T](fn.Signature.Params().At(0).Type()) {
		return nil
	}
	for i := 0; i < fn.Signature.Results().Len(); i++ {
		if _, pointer := fn.Signature.Results().At(i).Type().Underlying().(*types.Pointer); pointer || !typedBoundaryTypeIs[T](fn.Signature.Results().At(i).Type()) {
			return nil
		}
	}
	regs := make(map[ssa.Value]uint8)
	tupleRegs := make(map[ssa.Value]uint8)
	regs[fn.Params[0]] = 0
	code := make([]typedOp[T], 0, len(fn.Blocks)*4)
	blockPC := make([]int, len(fn.Blocks))
	type patch struct {
		pc    int
		block *ssa.BasicBlock
	}
	var branches []patch
	var jumps []patch
	nextReg := uint8(1)

	allocReg := func() (uint8, bool) {
		if nextReg >= maxTypedScalarRegs {
			return 0, false
		}
		r := nextReg
		nextReg++
		return r, true
	}
	var regFor func(ssa.Value) (uint8, bool)
	regFor = func(v ssa.Value) (uint8, bool) {
		if r, ok := regs[v]; ok {
			return r, true
		}
		if c, ok := v.(*ssa.Const); ok {
			r, ok := allocReg()
			if !ok {
				return 0, false
			}
			op, ok := typedConstOp[T](c, r, false)
			if !ok {
				return 0, false
			}
			regs[v] = r
			code = append(code, op)
			return r, true
		}
		if e, ok := v.(*ssa.Extract); ok {
			base, ok := tupleRegs[e.Tuple]
			if !ok || e.Index < 0 || e.Index >= fn.Signature.Results().Len() {
				return 0, false
			}
			r := base + uint8(e.Index)
			regs[v] = r
			return r, true
		}
		return 0, false
	}
	fail := func() TypedRunner { return nil }
	// Phi destinations must exist before compiling backedges.
	phis := make(map[*ssa.BasicBlock][]*ssa.Phi)
	maxPhis := 0
	for _, block := range fn.Blocks {
		for _, instr := range block.Instrs {
			if phi, ok := instr.(*ssa.Phi); ok {
				basic, isBasic := phi.Type().Underlying().(*types.Basic)
				if !typedTypeIs[T](phi.Type()) && (!isBasic || basic.Kind() != types.Bool) {
					return fail()
				}
				dst, ok := allocReg()
				if !ok || len(phi.Edges) != len(block.Preds) {
					return fail()
				}
				regs[phi] = dst
				phis[block] = append(phis[block], phi)
			}
		}
		if len(phis[block]) > maxPhis {
			maxPhis = len(phis[block])
		}
	}
	// Phi inputs can come from any predecessor, so their constants must be
	// initialized on every path. Keep the existing constant placement for
	// functions without Phi to avoid extra work in recursive leaf calls.
	if maxPhis > 0 {
		for _, block := range fn.Blocks {
			for _, instr := range block.Instrs {
				for _, operand := range instr.Operands(nil) {
					if _, ok := (*operand).(*ssa.Const); ok {
						if _, ok := regFor(*operand); !ok {
							return fail()
						}
					}
				}
			}
		}
	}
	for _, block := range fn.Blocks {
		blockPC[block.Index] = len(code)
		for _, instr := range block.Instrs {
			switch instr := instr.(type) {
			case *ssa.Phi:
				// Lowered to parallel copies on each incoming control-flow edge.
			case *ssa.UnOp:
				kind, valid := typedUnaryOpKind[T](instr.Op, instr.X.Type())
				if !valid {
					return fail()
				}
				x, ok := regFor(instr.X)
				if !ok {
					return fail()
				}
				dst, ok := allocReg()
				if !ok {
					return fail()
				}
				regs[instr] = dst
				code = append(code, typedOp[T]{kind: kind, dst: dst, x: x})
			case *ssa.BinOp:
				kind, valid := typedValueBinaryOpKind[T](instr.Op, instr.X.Type(), instr.Y.Type())
				if !valid && (instr.Op == token.SHL || instr.Op == token.SHR) && typedTypeIs[T](instr.X.Type()) {
					if c, ok := instr.Y.(*ssa.Const); ok && c.Value != nil && c.Value.Kind() == constant.Int {
						kind, valid = typedBinaryOpKind(instr.Op)
					}
				}
				if !valid {
					return fail()
				}
				x, ok := regFor(instr.X)
				if !ok {
					return fail()
				}
				y, ok := regFor(instr.Y)
				if !ok {
					return fail()
				}
				dst, ok := allocReg()
				if !ok {
					return fail()
				}
				regs[instr] = dst
				code = append(code, typedOp[T]{kind: kind, dst: dst, x: x, y: y})
			case *ssa.Call:
				callee, ok := instr.Call.Value.(*ssa.Function)
				if !ok || callee != fn || len(instr.Call.Args) != len(fn.Params) {
					return fail()
				}
				// Additional parameters may be omitted from the typed frame only
				// when they are forwarded unchanged and never otherwise used.
				// Any other use fails regFor or instruction validation below.
				for j := 1; j < len(fn.Params); j++ {
					if instr.Call.Args[j] != fn.Params[j] {
						return fail()
					}
				}
				x, ok := regFor(instr.Call.Args[0])
				if !ok {
					return fail()
				}
				base, ok := allocReg()
				nres := fn.Signature.Results().Len()
				if !ok || nres > maxTypedResults || (nres > 0 && base+uint8(nres-1) >= maxTypedScalarRegs) {
					return fail()
				}
				if nres > 0 {
					nextReg += uint8(nres - 1)
				}
				if nres == 1 {
					regs[instr] = base
				} else {
					tupleRegs[instr] = base
				}
				code = append(code, typedOp[T]{kind: typedCall, dst: base, x: x, nres: uint8(nres)})
			case *ssa.Extract:
				if _, ok := regFor(instr); !ok {
					return fail()
				}
			case *ssa.If:
				cond, ok := regFor(instr.Cond)
				if !ok || len(instr.Block().Succs) != 2 {
					return fail()
				}
				pc := len(code)
				code = append(code, typedOp[T]{kind: typedBranch, x: cond})
				branches = append(branches, patch{pc: pc, block: block})
			case *ssa.Jump:
				if len(instr.Block().Succs) != 1 {
					return fail()
				}
				pc := len(code)
				code = append(code, typedOp[T]{kind: typedJump})
				jumps = append(jumps, patch{pc: pc, block: block})
			case *ssa.Return:
				if len(instr.Results) > maxTypedResults {
					return fail()
				}
				results := make([]uint8, len(instr.Results))
				for i, value := range instr.Results {
					reg, ok := regFor(value)
					if !ok {
						return fail()
					}
					results[i] = reg
				}
				code = append(code, typedOp[T]{kind: typedReturn, nres: uint8(len(results)), results: results})
			default:
				return fail()
			}
		}
	}
	// Scratch registers are shared by all edges. Read all incoming values
	// before writing destinations, including cycles such as a,b = b,a.
	scratch := make([]uint8, maxPhis)
	for i := range scratch {
		reg, ok := allocReg()
		if !ok {
			return fail()
		}
		scratch[i] = reg
	}
	edgePC := func(from *ssa.BasicBlock, succ int) (int, bool) {
		to := from.Succs[succ]
		incoming := phis[to]
		if len(incoming) == 0 {
			return blockPC[to.Index], true
		}
		occurrence := 0
		for _, previous := range from.Succs[:succ] {
			if previous == to {
				occurrence++
			}
		}
		pred := -1
		for i, predecessor := range to.Preds {
			if predecessor == from {
				if occurrence == 0 {
					pred = i
					break
				}
				occurrence--
			}
		}
		if pred < 0 {
			return 0, false
		}
		pc := len(code)
		for i, phi := range incoming {
			src, ok := regFor(phi.Edges[pred])
			if !ok {
				return 0, false
			}
			kind := typedMove
			if phi.Type().Underlying().(*types.Basic).Kind() == types.Bool {
				kind = typedBoolMove
			}
			code = append(code, typedOp[T]{kind: kind, dst: scratch[i], x: src})
		}
		for i, phi := range incoming {
			kind := typedMove
			if phi.Type().Underlying().(*types.Basic).Kind() == types.Bool {
				kind = typedBoolMove
			}
			code = append(code, typedOp[T]{kind: kind, dst: regs[phi], x: scratch[i]})
		}
		code = append(code, typedOp[T]{kind: typedJump, target: blockPC[to.Index]})
		return pc, true
	}
	for _, p := range branches {
		truePC, ok := edgePC(p.block, 0)
		if !ok {
			return fail()
		}
		falsePC, ok := edgePC(p.block, 1)
		if !ok {
			return fail()
		}
		code[p.pc].truePC, code[p.pc].falsePC = truePC, falsePC
	}
	for _, p := range jumps {
		target, ok := edgePC(p.block, 0)
		if !ok {
			return fail()
		}
		code[p.pc].target = target
	}
	recursive := false
	for _, op := range code {
		if op.kind == typedCall {
			recursive = true
			break
		}
	}
	return &typedRunner[T]{code: code, recursive: recursive}
}

func typedTypeIs[T TypedScalar](typ types.Type) bool {
	b, ok := typ.Underlying().(*types.Basic)
	if !ok {
		return false
	}
	var zero T
	kind := reflect.TypeOf(zero).Kind()
	switch kind {
	case reflect.Int:
		return b.Kind() == types.Int
	case reflect.Int8:
		return b.Kind() == types.Int8
	case reflect.Int16:
		return b.Kind() == types.Int16
	case reflect.Int32:
		return b.Kind() == types.Int32
	case reflect.Int64:
		return b.Kind() == types.Int64
	case reflect.Uint:
		return b.Kind() == types.Uint
	case reflect.Uint8:
		return b.Kind() == types.Uint8
	case reflect.Uint16:
		return b.Kind() == types.Uint16
	case reflect.Uint32:
		return b.Kind() == types.Uint32
	case reflect.Uint64:
		return b.Kind() == types.Uint64
	case reflect.Uintptr:
		return b.Kind() == types.Uintptr
	case reflect.Float32:
		return b.Kind() == types.Float32
	case reflect.Float64:
		return b.Kind() == types.Float64
	default:
		return false
	}
}

// Interpreter-defined named values have distinct runtime types. A runner
// instantiated with a built-in T cannot bind those boxed values as T or *T.
// Explicit named generic instantiations still support their underlying kind.
func typedBoundaryTypeIs[T TypedScalar](typ types.Type) bool {
	if !typedTypeIs[T](typ) {
		return false
	}
	var zero T
	if reflect.TypeOf(zero).PkgPath() != "" {
		return true
	}
	_, basic := types.Unalias(typ).(*types.Basic)
	return basic
}

func isTypedInteger(typ types.Type) bool {
	b, ok := typ.Underlying().(*types.Basic)
	if !ok {
		return false
	}
	switch b.Kind() {
	case types.Int, types.Int8, types.Int16, types.Int32, types.Int64,
		types.Uint, types.Uint8, types.Uint16, types.Uint32, types.Uint64, types.Uintptr:
		return true
	default:
		return false
	}
}

// TypedArgumentsRunner binds all arguments, unlike the scalar runner, whose
// additional arguments can only be forwarded unchanged.
type TypedArgumentsRunner interface {
	TypedRunner
	RunArgs([]Value) Value
}

const typedArgsLimit = 64

type typedMemoryPointer int

type typedArgsFrame[T TypedScalar] struct {
	regs   [typedArgsLimit]T
	ptrs   [typedArgsLimit]typedMemoryPointer
	bools  [typedArgsLimit]bool
	locals [typedArgsLimit]T
	runner *typedArgsRunner[T]
	pc     int
}

const (
	typedAlloc typedOpKind = typedBoolNotEqual + 1 + iota
	typedIndexAddr
	typedLoad
	typedStore
	typedPointerConst
	typedPointerEqual
	typedPointerNotEqual
	typedArrayLoad
	typedArrayStore
	// Superinstructions emitted by the optimization passes.
	typedLessBranch
	typedLoadAddStore
	typedConstBlock
)

type typedArgsOp[T TypedScalar] struct {
	typedOp[T]
	// Calls store (source, destination) register pairs, grouped by argument kind.
	// offset ends the scalar group; length ends the pointer group; bool follows.
	args           []uint8
	callee         *typedArgsRunner[T]
	offset, length int
}

type typedArgKind uint8

const (
	typedArgScalar typedArgKind = iota
	typedArgPointer
	typedArgBool
)

type typedArgsRunner[T TypedScalar] struct {
	code        []typedArgsOp[T]
	params      []uint8
	kinds       []typedArgKind
	resultKinds []typedArgKind
	hasCalls    bool
	constants   []T // batched constant initialization
}

// batchTypedConstants collapses a dense numeric constant prologue to one
// dispatch and copy. Keep the old instructions for any existing PC references.
// All optimization passes mutate code and must run before execution starts.
func batchTypedConstants[T TypedScalar](root *typedArgsRunner[T]) int {
	seen := make(map[*typedArgsRunner[T]]bool)
	var visit func(*typedArgsRunner[T]) int
	visit = func(r *typedArgsRunner[T]) int {
		if seen[r] {
			return 0
		}
		seen[r] = true
		n := 0
		end := 0
		for end < len(r.code) && r.code[end].kind == typedConst && int(r.code[end].dst) == int(r.code[0].dst)+end {
			end++
		}
		if end > 1 {
			r.constants = make([]T, end)
			for j := 0; j < end; j++ {
				r.constants[j] = r.code[j].value
			}
			op := &r.code[0]
			op.kind, op.offset, op.length, op.target = typedConstBlock, int(op.dst), end, end
			n++
		}
		for _, op := range r.code {
			if op.kind == typedCall {
				n += visit(op.callee)
			}
		}
		return n
	}
	return visit(root)
}

// fuseTypedArguments combines common instruction sequences. Preserve instruction
// positions and intermediate results so branch targets and later uses remain
// valid. A visited set handles mutually recursive call graphs.
func fuseTypedArguments[T TypedScalar](root *typedArgsRunner[T]) int {
	seen := make(map[*typedArgsRunner[T]]bool)
	var visit func(*typedArgsRunner[T]) int
	visit = func(r *typedArgsRunner[T]) int {
		if seen[r] {
			return 0
		}
		seen[r] = true
		n := 0
		for i := range r.code {
			op := &r.code[i]
			if op.kind == typedCall {
				n += visit(op.callee)
			}
			if op.kind == typedLess && i+1 < len(r.code) {
				next := &r.code[i+1]
				if next.kind == typedBranch && next.x == op.dst {
					op.kind = typedLessBranch
					op.truePC, op.falsePC = next.truePC, next.falsePC
					n++
				}
			}
			if op.kind == typedLoad && i+2 < len(r.code) {
				add, store := &r.code[i+1], &r.code[i+2]
				if add.kind == typedAdd && store.kind == typedStore && store.x == op.x && store.y == add.dst && (add.x == op.dst || add.y == op.dst) {
					op.kind, op.offset, op.target = typedLoadAddStore, int(add.dst), i+3
					if add.x == op.dst {
						op.y = add.y
					} else {
						op.y, op.length = add.x, 1
					}
					n++
				}
			}
		}
		return n
	}
	return visit(root)
}

func (r *typedArgsRunner[T]) Run(input Value) Value { return r.RunArgs([]Value{input}) }

// Local addresses are frame/slot handles, not Go pointers into the frame slice.
// Growing the slice cannot invalidate addresses. A closed, checked call graph
// guarantees that these handles never escape through returns or external calls.
func (r *typedArgsRunner[T]) RunArgs(args []Value) Value {
	if len(args) != len(r.params) {
		panic("ixgo: invalid typed runner argument count")
	}
	// Leaf functions need only one frame. Reserve recursion capacity only for
	// functions with calls, keeping both initial allocations stack eligible.
	var frames []typedArgsFrame[T]
	if r.hasCalls {
		frames = make([]typedArgsFrame[T], 1, 8)
	} else {
		frames = make([]typedArgsFrame[T], 1)
	}
	external := make([]*T, 0, len(args))
	frames[0].runner = r
	for i, arg := range args {
		dst := r.params[i]
		switch r.kinds[i] {
		case typedArgScalar:
			frames[0].regs[dst] = arg.(T)
		case typedArgPointer:
			p := arg.(*T)
			if p == nil {
				continue
			}
			index := -1
			for j, previous := range external {
				if previous == p {
					index = j
					break
				}
			}
			if index < 0 {
				index = len(external)
				external = append(external, p)
			}
			frames[0].ptrs[dst] = typedMemoryPointer(-index - 1)
		case typedArgBool:
			frames[0].bools[dst] = arg.(bool)
		}
	}
	depth := 0
	for {
		fr := &frames[depth]
		op := &fr.runner.code[fr.pc]
		fr.pc++
		switch op.kind {
		case typedConst:
			fr.regs[op.dst] = op.value
		case typedConstBlock:
			copy(fr.regs[op.offset:op.offset+op.length], fr.runner.constants)
			fr.pc = op.target
		case typedBoolConst:
			fr.bools[op.dst] = op.value != 0
		case typedPointerConst:
			fr.ptrs[op.dst] = 0
		case typedMove:
			fr.regs[op.dst] = fr.regs[op.x]
			fr.ptrs[op.dst] = fr.ptrs[op.x]
			fr.bools[op.dst] = fr.bools[op.x]
		case typedAdd:
			fr.regs[op.dst] = fr.regs[op.x] + fr.regs[op.y]
		case typedSub:
			fr.regs[op.dst] = fr.regs[op.x] - fr.regs[op.y]
		case typedMul:
			fr.regs[op.dst] = fr.regs[op.x] * fr.regs[op.y]
		case typedQuo:
			fr.regs[op.dst] = fr.regs[op.x] / fr.regs[op.y]
		case typedRem:
			fr.regs[op.dst] = typedRemainder(fr.regs[op.x], fr.regs[op.y])
		case typedAnd:
			fr.regs[op.dst] = T(uint64(fr.regs[op.x]) & uint64(fr.regs[op.y]))
		case typedOr:
			fr.regs[op.dst] = T(uint64(fr.regs[op.x]) | uint64(fr.regs[op.y]))
		case typedXor:
			fr.regs[op.dst] = T(uint64(fr.regs[op.x]) ^ uint64(fr.regs[op.y]))
		case typedAndNot:
			fr.regs[op.dst] = T(uint64(fr.regs[op.x]) &^ uint64(fr.regs[op.y]))
		case typedShl:
			fr.regs[op.dst] = typedShiftLeft(fr.regs[op.x], fr.regs[op.y])
		case typedShr:
			fr.regs[op.dst] = typedShiftRight(fr.regs[op.x], fr.regs[op.y])
		case typedLess:
			fr.bools[op.dst] = fr.regs[op.x] < fr.regs[op.y]
		case typedLessBranch:
			fr.bools[op.dst] = fr.regs[op.x] < fr.regs[op.y]
			if fr.bools[op.dst] {
				fr.pc = op.truePC
			} else {
				fr.pc = op.falsePC
			}
		case typedLessEqual:
			fr.bools[op.dst] = fr.regs[op.x] <= fr.regs[op.y]
		case typedGreater:
			fr.bools[op.dst] = fr.regs[op.x] > fr.regs[op.y]
		case typedGreaterEqual:
			fr.bools[op.dst] = fr.regs[op.x] >= fr.regs[op.y]
		case typedEqual:
			fr.bools[op.dst] = fr.regs[op.x] == fr.regs[op.y]
		case typedNotEqual:
			fr.bools[op.dst] = fr.regs[op.x] != fr.regs[op.y]
		case typedBoolEqual:
			fr.bools[op.dst] = fr.bools[op.x] == fr.bools[op.y]
		case typedBoolNotEqual:
			fr.bools[op.dst] = fr.bools[op.x] != fr.bools[op.y]
		case typedPointerEqual:
			fr.bools[op.dst] = fr.ptrs[op.x] == fr.ptrs[op.y]
		case typedPointerNotEqual:
			fr.bools[op.dst] = fr.ptrs[op.x] != fr.ptrs[op.y]
		case typedNegate:
			fr.regs[op.dst] = -fr.regs[op.x]
		case typedComplement:
			fr.regs[op.dst] = T(^uint64(fr.regs[op.x]))
		case typedNot:
			fr.bools[op.dst] = !fr.bools[op.x]
		case typedAlloc:
			clear(fr.locals[op.offset : op.offset+op.length])
			fr.ptrs[op.dst] = typedMemoryPointer(depth*typedArgsLimit + op.offset + 1)
		case typedIndexAddr:
			index := fr.regs[op.y]
			if index < 0 {
				_ = fr.locals[:op.length][int64(index)] // Preserve a runtime bounds panic.
			}
			if uint64(index) >= uint64(op.length) {
				_ = fr.locals[:op.length][uint64(index)]
			}
			fr.ptrs[op.dst] = fr.ptrs[op.x] + typedMemoryPointer(index)
		case typedArrayLoad:
			slot := int(fr.ptrs[op.x]) - 1
			copy(fr.locals[op.offset:op.offset+op.length], frames[slot/typedArgsLimit].locals[slot%typedArgsLimit:slot%typedArgsLimit+op.length])
			fr.ptrs[op.dst] = typedMemoryPointer(depth*typedArgsLimit + op.offset + 1)
		case typedArrayStore:
			dst, src := int(fr.ptrs[op.x])-1, int(fr.ptrs[op.y])-1
			copy(frames[dst/typedArgsLimit].locals[dst%typedArgsLimit:dst%typedArgsLimit+op.length], frames[src/typedArgsLimit].locals[src%typedArgsLimit:src%typedArgsLimit+op.length])
		case typedLoad:
			p := fr.ptrs[op.x]
			if p > 0 {
				slot := int(p) - 1
				fr.regs[op.dst] = frames[slot/typedArgsLimit].locals[slot%typedArgsLimit]
			} else if p < 0 {
				fr.regs[op.dst] = *external[-int(p)-1]
			} else {
				var nilPointer *T
				fr.regs[op.dst] = *nilPointer
			}
		case typedStore:
			p := fr.ptrs[op.x]
			if p > 0 {
				slot := int(p) - 1
				frames[slot/typedArgsLimit].locals[slot%typedArgsLimit] = fr.regs[op.y]
			} else if p < 0 {
				*external[-int(p)-1] = fr.regs[op.y]
			} else {
				var nilPointer *T
				*nilPointer = fr.regs[op.y]
			}
		case typedLoadAddStore:
			handle := fr.ptrs[op.x]
			var p *T
			if handle > 0 {
				slot := int(handle) - 1
				p = &frames[slot/typedArgsLimit].locals[slot%typedArgsLimit]
			} else if handle < 0 {
				p = external[-int(handle)-1]
			}
			fr.regs[op.dst] = *p
			if op.length == 0 {
				fr.regs[op.offset] = fr.regs[op.dst] + fr.regs[op.y]
			} else {
				fr.regs[op.offset] = fr.regs[op.y] + fr.regs[op.dst]
			}
			*p = fr.regs[op.offset]
			fr.pc = op.target
		case typedCall:
			if depth+1 == len(frames) {
				frames = append(frames, typedArgsFrame[T]{})
				fr = &frames[depth] // append may have moved the backing array.
			}
			next := &frames[depth+1]
			next.runner, next.pc = op.callee, 0
			// Argument register moves are grouped by kind at compile time.
			for i := 0; i < op.offset; i += 2 {
				next.regs[op.args[i+1]] = fr.regs[op.args[i]]
			}
			for i := op.offset; i < op.length; i += 2 {
				next.ptrs[op.args[i+1]] = fr.ptrs[op.args[i]]
			}
			for i := op.length; i < len(op.args); i += 2 {
				next.bools[op.args[i+1]] = fr.bools[op.args[i]]
			}
			depth++
		case typedBranch:
			if fr.bools[op.x] {
				fr.pc = op.truePC
			} else {
				fr.pc = op.falsePC
			}
		case typedJump:
			fr.pc = op.target
		case typedReturn:
			if depth == 0 {
				switch len(op.results) {
				case 0:
					return nil
				case 1:
					if fr.runner.resultKinds[0] == typedArgBool {
						return fr.bools[op.results[0]]
					}
					return fr.regs[op.results[0]]
				default:
					result := make(tuple, len(op.results))
					for i, reg := range op.results {
						if fr.runner.resultKinds[i] == typedArgBool {
							result[i] = fr.bools[reg]
						} else {
							result[i] = fr.regs[reg]
						}
					}
					return result
				}
			}
			if len(op.results) != 0 {
				caller := &frames[depth-1]
				call := &caller.runner.code[caller.pc-1]
				for i, src := range op.results {
					if fr.runner.resultKinds[i] == typedArgBool {
						caller.bools[call.results[i]] = fr.bools[src]
					} else {
						caller.regs[call.results[i]] = fr.regs[src]
					}
				}
			}
			depth--
		}
	}
}

// typedArgsCompiler compiles a closed call graph over T, bool, *T,
// and local [N]T arrays. Pointer returns, external calls and unsupported SSA
// cause the whole graph to fall back to the ordinary interpreter. Each run
// owns its frame arena, so compiled runners may be called concurrently.
// Functions are limited to 64 registers and 64 local scalar slots per frame,
// with up to eight scalar or bool results. Use TypedArgumentsRunner.RunArgs to bind
// every parameter. Run is the convenience entry for a single parameter.
type typedArgsCompiler[T TypedScalar] struct {
	compiled map[*ssa.Function]*typedArgsRunner[T]
	depth    int
}

func (c *typedArgsCompiler[T]) build(fn *ssa.Function) (runner *typedArgsRunner[T]) {
	if r, ok := c.compiled[fn]; ok {
		return r
	}
	c.depth++
	defer func() {
		c.depth--
		if runner == nil {
			// A cached callee may still refer to this recursion placeholder.
			// Roll back the entire graph when its root fails compilation.
			if c.depth == 0 {
				clear(c.compiled)
			}
			c.compiled[fn] = nil
		}
	}()
	if fn == nil || fn.Signature == nil || len(fn.Blocks) == 0 || fn.Recover != nil || len(fn.FreeVars) != 0 || fn.Signature.Variadic() || len(fn.Params) > typedArgsLimit || fn.Signature.Results().Len() > maxTypedResults {
		return nil
	}
	r := &typedArgsRunner[T]{}
	c.compiled[fn] = r // Resolve mutual recursion before compiling callees.
	kindFor := func(typ types.Type) (typedArgKind, bool) {
		if typedTypeIs[T](typ) {
			return typedArgScalar, true
		}
		if p, ok := typ.(*types.Pointer); ok && typedTypeIs[T](p.Elem()) {
			return typedArgPointer, true
		}
		if typ == types.Typ[types.Bool] {
			return typedArgBool, true
		}
		return 0, false
	}
	for i := 0; i < fn.Signature.Results().Len(); i++ {
		kind, ok := kindFor(fn.Signature.Results().At(i).Type())
		if !ok || kind == typedArgPointer || (kind == typedArgScalar && !typedBoundaryTypeIs[T](fn.Signature.Results().At(i).Type())) {
			return nil
		}
		r.resultKinds = append(r.resultKinds, kind)
	}
	regs := make(map[ssa.Value]uint8)
	tuples := make(map[ssa.Value][]uint8)
	nextReg := 0
	alloc := func() (uint8, bool) {
		if nextReg == typedArgsLimit {
			return 0, false
		}
		reg := uint8(nextReg)
		nextReg++
		return reg, true
	}
	for _, p := range fn.Params {
		kind, ok := kindFor(p.Type())
		if kind == typedArgScalar && !typedBoundaryTypeIs[T](p.Type()) {
			return nil
		}
		if kind == typedArgPointer {
			pointer := p.Type().(*types.Pointer)
			if !typedBoundaryTypeIs[T](pointer.Elem()) {
				return nil
			}
		}
		if !ok {
			return nil
		}
		reg, ok := alloc()
		if !ok {
			return nil
		}
		regs[p] = reg
		r.params = append(r.params, reg)
		r.kinds = append(r.kinds, kind)
	}
	// Preassign destinations so Phi backedges and extracts can refer forward.
	for _, b := range fn.Blocks {
		for _, instr := range b.Instrs {
			v, ok := instr.(ssa.Value)
			if !ok {
				continue
			}
			if _, ok := instr.(*ssa.Extract); ok {
				continue
			}
			if call, ok := instr.(*ssa.Call); ok {
				n := call.Call.Signature().Results().Len()
				if n != 1 {
					for i := 0; i < n; i++ {
						reg, ok := alloc()
						if !ok {
							return nil
						}
						tuples[v] = append(tuples[v], reg)
					}
					continue
				}
			}
			reg, ok := alloc()
			if !ok {
				return nil
			}
			regs[v] = reg
		}
	}
	emit := func(op typedArgsOp[T]) { r.code = append(r.code, op) }
	var regFor func(ssa.Value) (uint8, bool)
	regFor = func(v ssa.Value) (uint8, bool) {
		if reg, ok := regs[v]; ok {
			return reg, true
		}
		if e, ok := v.(*ssa.Extract); ok {
			results := tuples[e.Tuple]
			if e.Index < 0 || e.Index >= len(results) {
				return 0, false
			}
			regs[v] = results[e.Index]
			return results[e.Index], true
		}
		c, ok := v.(*ssa.Const)
		if !ok {
			return 0, false
		}
		reg, ok := alloc()
		if !ok {
			return 0, false
		}
		base, ok := typedConstOp[T](c, reg, true)
		if !ok || (c.IsNil() && func() bool { _, valid := kindFor(c.Type()); return !valid }()) {
			return 0, false
		}
		op := typedArgsOp[T]{typedOp: base}
		regs[v] = reg
		emit(op)
		return reg, true
	}
	// Constants dominate all paths, including copies on Phi backedges.
	for _, b := range fn.Blocks {
		for _, i := range b.Instrs {
			for _, operand := range i.Operands(nil) {
				if c, ok := (*operand).(*ssa.Const); ok {
					if _, ok := regFor(c); !ok {
						return nil
					}
				}
			}
		}
	}
	blockPC := make([]int, len(fn.Blocks))
	type edge struct {
		pc     int
		from   *ssa.BasicBlock
		succ   int
		branch bool
	}
	var edges []edge
	locals := 0
	for _, b := range fn.Blocks {
		blockPC[b.Index] = len(r.code)
		for _, instr := range b.Instrs {
			op := typedArgsOp[T]{}
			if v, ok := instr.(ssa.Value); ok {
				op.dst = regs[v]
			}
			switch i := instr.(type) {
			case *ssa.Phi:
				if _, ok := kindFor(i.Type()); !ok {
					return nil
				}
				continue
			case *ssa.Extract:
				if _, ok := regFor(i); !ok {
					return nil
				}
				continue
			case *ssa.Alloc:
				p, ok := i.Type().(*types.Pointer)
				if !ok {
					return nil
				}
				n := 1
				if a, ok := p.Elem().(*types.Array); ok {
					if !typedTypeIs[T](a.Elem()) || a.Len() > typedArgsLimit {
						return nil
					}
					n = int(a.Len())
				} else if !typedTypeIs[T](p.Elem()) {
					return nil
				}
				// Alloc inside a cycle could create addresses that outlive this reuse.
				if blockInCycle(b) {
					return nil
				}
				if locals+n > typedArgsLimit {
					return nil
				}
				op.kind, op.offset, op.length = typedAlloc, locals, n
				locals += n
			case *ssa.IndexAddr:
				p, ok := i.X.Type().(*types.Pointer)
				if !ok {
					return nil
				}
				a, ok := p.Elem().(*types.Array)
				if !ok || !typedTypeIs[T](a.Elem()) || a.Len() > typedArgsLimit {
					return nil
				}
				op.kind, op.length = typedIndexAddr, int(a.Len())
				var valid bool
				op.x, valid = regFor(i.X)
				if !valid {
					return nil
				}
				op.y, valid = regFor(i.Index)
				if !valid {
					return nil
				}
				if _, ok := i.Index.(*ssa.Const); !ok && !typedTypeIs[T](i.Index.Type()) {
					return nil
				}
			case *ssa.UnOp:
				var ok bool
				op.x, ok = regFor(i.X)
				if !ok {
					return nil
				}
				switch i.Op {
				case token.MUL:
					if a, ok := i.Type().(*types.Array); ok && typedTypeIs[T](a.Elem()) && a.Len() <= typedArgsLimit {
						n := int(a.Len())
						if locals+n > typedArgsLimit {
							return nil
						}
						op.kind, op.offset, op.length = typedArrayLoad, locals, n
						locals += n
					} else {
						if kind, ok := kindFor(i.X.Type()); !ok || kind != typedArgPointer {
							return nil
						}
						op.kind = typedLoad
					}
				default:
					kind, valid := typedUnaryOpKind[T](i.Op, i.X.Type())
					if !valid {
						return nil
					}
					op.kind = kind
				}
			case *ssa.Store:
				if a, ok := i.Val.Type().(*types.Array); ok && typedTypeIs[T](a.Elem()) && a.Len() <= typedArgsLimit {
					op.kind, op.length = typedArrayStore, int(a.Len())
				} else {
					if kind, ok := kindFor(i.Addr.Type()); !ok || kind != typedArgPointer || !typedTypeIs[T](i.Val.Type()) {
						return nil
					}
					op.kind = typedStore
				}
				var ok bool
				op.x, ok = regFor(i.Addr)
				if !ok {
					return nil
				}
				op.y, ok = regFor(i.Val)
				if !ok {
					return nil
				}
			case *ssa.BinOp:
				var ok bool
				op.x, ok = regFor(i.X)
				if !ok {
					return nil
				}
				op.y, ok = regFor(i.Y)
				if !ok {
					return nil
				}
				if kind, valid := typedValueBinaryOpKind[T](i.Op, i.X.Type(), i.Y.Type()); valid {
					op.kind = kind
				} else if (i.Op == token.SHL || i.Op == token.SHR) && typedTypeIs[T](i.X.Type()) {
					if c, ok := i.Y.(*ssa.Const); ok && c.Value != nil && c.Value.Kind() == constant.Int {
						op.kind, _ = typedBinaryOpKind(i.Op)
					} else {
						return nil
					}
				} else if kind, valid := kindFor(i.X.Type()); valid && kind == typedArgPointer && types.Identical(i.X.Type(), i.Y.Type()) {
					if i.Op == token.EQL {
						op.kind = typedPointerEqual
					} else if i.Op == token.NEQ {
						op.kind = typedPointerNotEqual
					} else {
						return nil
					}
				} else {
					return nil
				}
			case *ssa.Call:
				r.hasCalls = true
				callee, ok := i.Call.Value.(*ssa.Function)
				if !ok || i.Call.IsInvoke() {
					return nil
				}
				op.callee = c.build(callee)
				if op.callee == nil || len(i.Call.Args) != len(op.callee.params) {
					return nil
				}
				for j, arg := range i.Call.Args {
					if !types.Identical(arg.Type(), callee.Params[j].Type()) {
						return nil
					}
					reg, ok := regFor(arg)
					if !ok {
						return nil
					}
					op.args = append(op.args, reg)
				}
				if callee.Signature.Results().Len() == 1 {
					op.results = []uint8{regs[i]}
				} else {
					op.results = tuples[i]
				}
				// Store (source, destination) pairs in scalar/pointer/bool groups.
				// Different frames make reordering these copies safe.
				sources := op.args
				op.args = make([]uint8, 0, 2*len(sources))
				for _, kind := range []typedArgKind{typedArgScalar, typedArgPointer, typedArgBool} {
					for j, src := range sources {
						if op.callee.kinds[j] == kind {
							op.args = append(op.args, src, op.callee.params[j])
						}
					}
					if kind == typedArgScalar {
						op.offset = len(op.args)
					} else if kind == typedArgPointer {
						op.length = len(op.args)
					}
				}
				op.kind = typedCall
			case *ssa.If:
				var ok bool
				op.x, ok = regFor(i.Cond)
				if !ok || i.Cond.Type() != types.Typ[types.Bool] {
					return nil
				}
				op.kind = typedBranch
				edges = append(edges, edge{len(r.code), b, 0, true}, edge{len(r.code), b, 1, true})
			case *ssa.Jump:
				op.kind = typedJump
				edges = append(edges, edge{len(r.code), b, 0, false})
			case *ssa.Return:
				op.kind = typedReturn
				for _, v := range i.Results {
					reg, ok := regFor(v)
					if !ok {
						return nil
					}
					op.results = append(op.results, reg)
				}
			default:
				return nil
			}
			emit(op)
		}
	}
	// Parallel Phi copies share scratch registers across edges.
	var scratch []uint8
	for _, e := range edges {
		to := e.from.Succs[e.succ]
		var phis []*ssa.Phi
		for _, instr := range to.Instrs {
			if p, ok := instr.(*ssa.Phi); ok {
				phis = append(phis, p)
			}
		}
		target := blockPC[to.Index]
		if len(phis) > 0 {
			occurrence := 0
			for _, s := range e.from.Succs[:e.succ] {
				if s == to {
					occurrence++
				}
			}
			pred := -1
			for j, p := range to.Preds {
				if p == e.from {
					if occurrence == 0 {
						pred = j
						break
					}
					occurrence--
				}
			}
			if pred < 0 {
				return nil
			}
			for len(scratch) < len(phis) {
				reg, ok := alloc()
				if !ok {
					return nil
				}
				scratch = append(scratch, reg)
			}
			target = len(r.code)
			for j, p := range phis {
				src, ok := regFor(p.Edges[pred])
				if !ok {
					return nil
				}
				emit(typedArgsOp[T]{typedOp: typedOp[T]{kind: typedMove, dst: scratch[j], x: src}})
			}
			for j, p := range phis {
				emit(typedArgsOp[T]{typedOp: typedOp[T]{kind: typedMove, dst: regs[p], x: scratch[j]}})
			}
			emit(typedArgsOp[T]{typedOp: typedOp[T]{kind: typedJump, target: blockPC[to.Index]}})
		}
		if e.branch {
			if e.succ == 0 {
				r.code[e.pc].truePC = target
			} else {
				r.code[e.pc].falsePC = target
			}
		} else {
			r.code[e.pc].target = target
		}
	}
	return r
}

func makeTypedArgumentsRunnerOf[T TypedScalar](fn *ssa.Function) TypedRunner {
	c := typedArgsCompiler[T]{compiled: make(map[*ssa.Function]*typedArgsRunner[T])}
	runner := c.build(fn)
	if runner == nil {
		return nil
	}
	return runner
}

// Local allocation slots can be reused across calls but not across iterations
// of the same function. Such allocations fall back until fresh-slot lifetimes
// are implemented.
func blockInCycle(start *ssa.BasicBlock) bool {
	seen := map[*ssa.BasicBlock]bool{}
	var visit func(*ssa.BasicBlock) bool
	visit = func(b *ssa.BasicBlock) bool {
		if b == start {
			return true
		}
		if seen[b] {
			return false
		}
		seen[b] = true
		for _, s := range b.Succs {
			if visit(s) {
				return true
			}
		}
		return false
	}
	for _, s := range start.Succs {
		if visit(s) {
			return true
		}
	}
	return false
}

func typedBinaryOpKind(op token.Token) (typedOpKind, bool) {
	switch op {
	case token.ADD:
		return typedAdd, true
	case token.SUB:
		return typedSub, true
	case token.MUL:
		return typedMul, true
	case token.QUO:
		return typedQuo, true
	case token.REM:
		return typedRem, true
	case token.AND:
		return typedAnd, true
	case token.OR:
		return typedOr, true
	case token.XOR:
		return typedXor, true
	case token.AND_NOT:
		return typedAndNot, true
	case token.SHL:
		return typedShl, true
	case token.SHR:
		return typedShr, true
	case token.LSS:
		return typedLess, true
	case token.LEQ:
		return typedLessEqual, true
	case token.GTR:
		return typedGreater, true
	case token.GEQ:
		return typedGreaterEqual, true
	case token.EQL:
		return typedEqual, true
	case token.NEQ:
		return typedNotEqual, true
	}
	return 0, false
}

// Reject constants that cannot be represented by the register type, including
// wider shift counts and array indices. Truncation would change Go semantics.
func typedConstOp[T TypedScalar](c *ssa.Const, dst uint8, allowNil bool) (typedOp[T], bool) {
	op := typedOp[T]{kind: typedConst, dst: dst}
	if c.IsNil() {
		if !allowNil {
			return op, false
		}
		op.kind = typedPointerConst
		return op, true
	}
	if c.Value == nil {
		return op, false
	}
	if c.Value.Kind() == constant.Bool {
		op.kind = typedBoolConst
		if constant.BoolVal(c.Value) {
			op.value = 1
		}
		return op, true
	}
	if !typedTypeIs[T](c.Type()) && (c.Value.Kind() != constant.Int || !isTypedInteger(c.Type())) {
		return op, false
	}
	var ok bool
	op.value, ok = typedConstantValue[T](c.Value)
	return op, ok
}

func typedConstantValue[T TypedScalar](v constant.Value) (T, bool) {
	var zero T
	kind := reflect.TypeOf(zero).Kind()
	if kind == reflect.Float32 || kind == reflect.Float64 {
		if v.Kind() != constant.Int && v.Kind() != constant.Float {
			return zero, false
		}
		var n float64
		var exact bool
		if kind == reflect.Float32 {
			f, ok := constant.Float32Val(v)
			n, exact = float64(f), ok
		} else {
			n, exact = constant.Float64Val(v)
		}
		// Integer constants may also be array indices. Do not round them.
		return T(n), !math.IsInf(n, 0) && (v.Kind() != constant.Int || exact)
	}
	if v.Kind() != constant.Int {
		return zero, false
	}
	bits := reflect.TypeOf(zero).Bits()
	if bits == 0 {
		return zero, false
	}
	signed := reflect.TypeOf(zero).Kind() >= reflect.Int && reflect.TypeOf(zero).Kind() <= reflect.Int64
	if n, ok := constant.Int64Val(v); ok {
		if n < 0 {
			min := -(int64(1) << (bits - 1))
			if !signed || n < min {
				return zero, false
			}
			return T(n), true
		}
		// Unsigned types accept the full nonnegative range.
		if signed {
			if bits < 64 && n >= int64(1)<<(bits-1) {
				return zero, false
			}
		} else if bits < 64 && uint64(n) >= uint64(1)<<bits {
			return zero, false
		}
		return T(n), true
	}
	if n, ok := constant.Uint64Val(v); ok {
		if signed || bits < 64 && n >= uint64(1)<<bits {
			return zero, false
		}
		return T(n), true
	}
	return zero, false
}

func typedUnaryOpKind[T TypedScalar](op token.Token, typ types.Type) (typedOpKind, bool) {
	if typ == types.Typ[types.Bool] && op == token.NOT {
		return typedNot, true
	}
	if typedTypeIs[T](typ) {
		switch op {
		case token.SUB:
			return typedNegate, true
		case token.XOR:
			if isTypedInteger(typ) {
				return typedComplement, true
			}
		}
	}
	return 0, false
}

func typedValueBinaryOpKind[T TypedScalar](op token.Token, x, y types.Type) (typedOpKind, bool) {
	if typedTypeIs[T](x) && typedTypeIs[T](y) {
		if !isTypedInteger(x) {
			switch op {
			case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.LEQ, token.GTR, token.GEQ, token.EQL, token.NEQ:
			default:
				return 0, false
			}
		}
		return typedBinaryOpKind(op)
	}
	if x == types.Typ[types.Bool] && y == types.Typ[types.Bool] {
		switch op {
		case token.EQL:
			return typedBoolEqual, true
		case token.NEQ:
			return typedBoolNotEqual, true
		}
	}
	return 0, false
}

// Integer-only instructions keep their low bits when narrowing back to T.
// The SSA compiler never emits these instructions for floating-point operands.
// Keeping conversions here allows the shared execution loops to accept floats
// without per-instruction callbacks or a second copy of the control flow.
func typedRemainder[T TypedScalar](x, y T) T {
	var zero T
	if zero-1 > zero {
		return T(uint64(x) % uint64(y))
	}
	return T(int64(x) % int64(y))
}

func typedShiftLeft[T TypedScalar](x, y T) T {
	if y < 0 {
		return T(uint64(x) << int64(y)) // Preserve negative shift panics.
	}
	return T(uint64(x) << uint64(y))
}

func typedShiftRight[T TypedScalar](x, y T) T {
	if y < 0 {
		return T(int64(x) >> int64(y))
	}
	var zero T
	if zero-1 > zero {
		return T(uint64(x) >> uint64(y))
	}
	return T(int64(x) >> uint64(y))
}
