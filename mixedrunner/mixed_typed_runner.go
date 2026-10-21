/*
 * Copyright (c) 2026 The GoPlus Authors (goplus.org). All rights reserved.
 * Licensed under the Apache License, Version 2.0.
 */

package mixedrunner

import (
	"go/constant"
	"go/token"
	"go/types"
	"math"
	"reflect"
	"strconv"

	"github.com/goplus/ixgo"
	"golang.org/x/tools/go/ssa"
)

// mixedScalar is the numeric set supported by the mixed runner.
type mixedScalar interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr | ~float32 | ~float64
}

const mixedArgsLimit = 64
const mixedMaxResults = 8

// Register indices and result counts must fit in uint8.
const _ uint8 = mixedArgsLimit - 1
const _ uint8 = mixedMaxResults

type mixedOpKind uint8

const (
	mixedNop mixedOpKind = iota
	mixedBaseAdd
	mixedBaseSub
	mixedBaseMul
	mixedBaseQuo
	mixedBaseRem
	mixedBaseAnd
	mixedBaseOr
	mixedBaseXor
	mixedBaseAndNot
	mixedBaseShl
	mixedBaseShr
	mixedBaseLess
	mixedBaseLessEqual
	mixedBaseGreater
	mixedBaseGreaterEqual
	mixedBaseEqual
	mixedBaseNotEqual
	mixedBaseCall
	mixedBaseBranch
	mixedBaseJump
	mixedBaseReturn
	mixedBaseMove
	mixedBaseBoolMove
	mixedBaseBoolConst
	mixedBaseNegate
	mixedBaseComplement
	mixedBaseNot
	mixedBaseBoolEqual
	mixedBaseBoolNotEqual
)

type mixedBaseOp[T mixedScalar] struct {
	kind                    mixedOpKind
	dst, x, y               uint8
	nres                    uint8
	results                 []uint8
	value                   T
	truePC, falsePC, target int
}

const (
	mixedConvert mixedOpKind = mixedBaseBoolNotEqual + 1 + iota
	mixedAdd
	mixedSub
	mixedOr
	mixedShl
	mixedUnsignedShr
	mixedSignedShr
	mixedUnsignedLess
	mixedSignedLess
	mixedUnsignedLessEqual
	mixedSignedLessEqual
	mixedUnsignedGreater
	mixedSignedGreater
	mixedUnsignedGreaterEqual
	mixedSignedGreaterEqual
	mixedEqual
	mixedNotEqual
	mixedNegate
	mixedNot
)

// Every register carries raw bits; the instruction supplies its numeric kind.
type mixedOp struct {
	mixedBaseOp[uint64]
	scalar, other types.BasicKind
	callee        *mixedTypedRunner
	args          []uint8
}

type mixedTypedRunner struct {
	code            []mixedOp
	constants       []uint64
	params, results []types.BasicKind
}

func mixedKind(t types.Type) (types.BasicKind, bool) {
	b, ok := t.Underlying().(*types.Basic)
	if !ok {
		return 0, false
	}
	k := b.Kind()
	return k, k == types.Bool || k >= types.Int && k <= types.Float64
}

func mixedBoundaryKind(t types.Type) (types.BasicKind, bool) {
	if _, ok := types.Unalias(t).(*types.Basic); !ok {
		return 0, false
	}
	return mixedKind(t)
}

// MakeRunner compiles closed numeric and boolean static call graphs.
// Unsupported functions, named boundary types and oversized frames return nil.
// Direct construction does not install memo or debug hooks.
func MakeRunner(fn *ssa.Function) ixgo.TypedRunner {
	c := mixedCompiler{compiled: make(map[*ssa.Function]*mixedTypedRunner)}
	if r := c.build(fn); r != nil {
		return r
	}
	return nil
}

type mixedCompiler struct {
	compiled map[*ssa.Function]*mixedTypedRunner
}

func (c *mixedCompiler) build(fn *ssa.Function) *mixedTypedRunner {
	if r, ok := c.compiled[fn]; ok {
		return r
	}
	if fn == nil || fn.Signature == nil || fn.Recover != nil || len(fn.FreeVars) != 0 || len(fn.Blocks) == 0 || fn.Signature.Variadic() || fn.Signature.Results().Len() > mixedMaxResults {
		return nil
	}
	r := &mixedTypedRunner{}
	regs := make(map[ssa.Value]uint8)
	alloc := func(v ssa.Value) bool {
		if _, ok := regs[v]; ok {
			return true
		}
		if _, ok := mixedKind(v.Type()); !ok || len(r.constants) >= mixedArgsLimit {
			return false
		}
		regs[v] = uint8(len(r.constants))
		r.constants = append(r.constants, 0)
		return true
	}
	for _, p := range fn.Params {
		k, ok := mixedBoundaryKind(p.Type())
		if !ok || !alloc(p) {
			return nil
		}
		r.params = append(r.params, k)
	}
	for i := 0; i < fn.Signature.Results().Len(); i++ {
		k, ok := mixedBoundaryKind(fn.Signature.Results().At(i).Type())
		if !ok {
			return nil
		}
		r.results = append(r.results, k)
	}
	// Publish signatures before resolving recursive edges. A failed graph is discarded.
	c.compiled[fn] = r
	tuples := make(map[ssa.Value][]uint8)
	phis := make(map[*ssa.BasicBlock][]*ssa.Phi)
	maxPhis := 0
	// Allocate definitions and constants before lowering edges, including backedges.
	for _, b := range fn.Blocks {
		for _, ins := range b.Instrs {
			if call, ok := ins.(*ssa.Call); ok && call.Call.Signature().Results().Len() != 1 {
				n := call.Call.Signature().Results().Len()
				if n > mixedMaxResults || len(r.constants)+n > mixedArgsLimit {
					return nil
				}
				for j := 0; j < n; j++ {
					tuples[call] = append(tuples[call], uint8(len(r.constants)))
					r.constants = append(r.constants, 0)
				}
			} else if v, ok := ins.(ssa.Value); ok && !alloc(v) {
				return nil
			}
			if phi, ok := ins.(*ssa.Phi); ok {
				if len(phi.Edges) != len(b.Preds) {
					return nil
				}
				phis[b] = append(phis[b], phi)
			}
			for _, operand := range ins.Operands(nil) {
				if operand == nil || *operand == nil {
					continue
				}
				if c, ok := (*operand).(*ssa.Const); ok {
					if !alloc(c) {
						return nil
					}
					bits, ok := mixedConstant(c)
					if !ok {
						return nil
					}
					r.constants[regs[c]] = bits
				}
			}
		}
		if len(phis[b]) > maxPhis {
			maxPhis = len(phis[b])
		}
	}
	scratch := len(r.constants)
	if scratch+maxPhis > mixedArgsLimit {
		return nil
	}
	type patch struct {
		pc    int
		block *ssa.BasicBlock
	}
	var patches []patch
	blockPC := make(map[*ssa.BasicBlock]int)
	reg := func(v ssa.Value) (uint8, bool) { n, ok := regs[v]; return n, ok }
	for _, b := range fn.Blocks {
		blockPC[b] = len(r.code)
		for _, ins := range b.Instrs {
			var op mixedOp
			switch in := ins.(type) {
			case *ssa.Phi:
				continue
			case *ssa.Call:
				callee, ok := in.Call.Value.(*ssa.Function)
				if !ok || in.Call.IsInvoke() {
					return nil
				}
				op.callee = c.build(callee)
				if op.callee == nil || len(in.Call.Args) != len(op.callee.params) {
					return nil
				}
				for j, arg := range in.Call.Args {
					if !types.Identical(arg.Type(), callee.Params[j].Type()) {
						return nil
					}
					x, ok := reg(arg)
					if !ok {
						return nil
					}
					op.args = append(op.args, x)
				}
				op.kind = mixedBaseCall
				if len(op.callee.results) == 1 {
					op.results = []uint8{regs[in]}
				} else {
					op.results = tuples[in]
				}
			case *ssa.Extract:
				results := tuples[in.Tuple]
				if in.Index < 0 || in.Index >= len(results) {
					return nil
				}
				op = mixedOp{mixedBaseOp: mixedBaseOp[uint64]{kind: mixedBaseMove, dst: regs[in], x: results[in.Index]}}
			case *ssa.BinOp:
				k, ok := mixedKind(in.X.Type())
				if !ok {
					return nil
				}
				yk, ok := mixedKind(in.Y.Type())
				if !ok {
					return nil
				}
				kind, ok := mixedBinaryOpKind(in.Op)
				if !ok {
					return nil
				}
				if in.Op == token.SHL || in.Op == token.SHR {
					if !mixedIsInteger(in.X.Type()) || !mixedIsInteger(in.Y.Type()) {
						return nil
					}
				} else if k != yk {
					return nil
				}
				if k == types.Bool && in.Op != token.EQL && in.Op != token.NEQ {
					return nil
				}
				if k == types.Float32 || k == types.Float64 {
					switch in.Op {
					case token.ADD, token.SUB, token.MUL, token.QUO, token.LSS, token.LEQ, token.GTR, token.GEQ, token.EQL, token.NEQ:
					default:
						return nil
					}
				}
				x, ok := reg(in.X)
				if !ok {
					return nil
				}
				y, ok := reg(in.Y)
				if !ok {
					return nil
				}
				op = mixedOp{mixedBaseOp: mixedBaseOp[uint64]{kind: kind, dst: regs[in], x: x, y: y}, scalar: k, other: yk}
				if in.Op == token.SHL || in.Op == token.SHR {
					if count, ok := in.Y.(*ssa.Const); ok && constant.Sign(count.Value) >= 0 {
						op.other = types.Uint64
					}
				}
			case *ssa.UnOp:
				k, ok := mixedKind(in.X.Type())
				if !ok {
					return nil
				}
				var kind mixedOpKind
				switch in.Op {
				case token.NOT:
					if k != types.Bool {
						return nil
					}
					kind = mixedBaseNot
				case token.SUB:
					if k == types.Bool {
						return nil
					}
					kind = mixedBaseNegate
				case token.XOR:
					if !mixedIsInteger(in.X.Type()) {
						return nil
					}
					kind = mixedBaseComplement
				default:
					return nil
				}
				x, ok := reg(in.X)
				if !ok {
					return nil
				}
				op = mixedOp{mixedBaseOp: mixedBaseOp[uint64]{kind: kind, dst: regs[in], x: x}, scalar: k}
			case *ssa.Convert:
				k, ok := mixedKind(in.X.Type())
				if !ok || k == types.Bool {
					return nil
				}
				to, ok := mixedKind(in.Type())
				if !ok || to == types.Bool {
					return nil
				}
				x, ok := reg(in.X)
				if !ok {
					return nil
				}
				op = mixedOp{mixedBaseOp: mixedBaseOp[uint64]{kind: mixedConvert, dst: regs[in], x: x}, scalar: k, other: to}
			case *ssa.If:
				x, ok := reg(in.Cond)
				if !ok || len(b.Succs) != 2 {
					return nil
				}
				op = mixedOp{mixedBaseOp: mixedBaseOp[uint64]{kind: mixedBaseBranch, x: x}}
				patches = append(patches, patch{len(r.code), b})
			case *ssa.Jump:
				if len(b.Succs) != 1 {
					return nil
				}
				op.kind = mixedBaseJump
				patches = append(patches, patch{len(r.code), b})
			case *ssa.Return:
				if len(in.Results) != len(r.results) {
					return nil
				}
				op.kind = mixedBaseReturn
				for _, v := range in.Results {
					x, ok := reg(v)
					if !ok {
						return nil
					}
					op.results = append(op.results, x)
				}
			default:
				return nil
			}
			r.code = append(r.code, op)
		}
	}
	edge := func(from *ssa.BasicBlock, succ int) (int, bool) {
		to := from.Succs[succ]
		if len(phis[to]) == 0 {
			return blockPC[to], true
		}
		occurrence := 0
		for _, b := range from.Succs[:succ] {
			if b == to {
				occurrence++
			}
		}
		pred := -1
		for i, b := range to.Preds {
			if b == from {
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
		pc := len(r.code)
		for i, phi := range phis[to] {
			x, ok := reg(phi.Edges[pred])
			if !ok {
				return 0, false
			}
			r.code = append(r.code, mixedOp{mixedBaseOp: mixedBaseOp[uint64]{kind: mixedBaseMove, dst: uint8(scratch + i), x: x}})
		}
		for i, phi := range phis[to] {
			r.code = append(r.code, mixedOp{mixedBaseOp: mixedBaseOp[uint64]{kind: mixedBaseMove, dst: regs[phi], x: uint8(scratch + i)}})
		}
		r.code = append(r.code, mixedOp{mixedBaseOp: mixedBaseOp[uint64]{kind: mixedBaseJump, target: blockPC[to]}})
		return pc, true
	}
	for _, p := range patches {
		target, ok := edge(p.block, 0)
		if !ok {
			return nil
		}
		if r.code[p.pc].kind == mixedBaseJump {
			r.code[p.pc].target = target
		} else {
			other, ok := edge(p.block, 1)
			if !ok {
				return nil
			}
			r.code[p.pc].truePC = target
			r.code[p.pc].falsePC = other
		}
	}
	for i := range r.code {
		specializeMixedOp(&r.code[i])
	}
	return r
}

// Full-width integer operations need no normalization or dynamic scalar dispatch.
func specializeMixedOp(op *mixedOp) {
	if op.scalar == types.Bool {
		switch op.kind {
		case mixedBaseNot:
			op.kind = mixedNot
		case mixedBaseEqual:
			op.kind = mixedEqual
		case mixedBaseNotEqual:
			op.kind = mixedNotEqual
		}
		return
	}
	signed := op.scalar == types.Int64 || op.scalar == types.Int
	if op.scalar != types.Uint64 && op.scalar != types.Int64 && !(strconv.IntSize == 64 && (op.scalar == types.Int || op.scalar == types.Uint || op.scalar == types.Uintptr)) {
		return
	}
	switch op.kind {
	case mixedBaseAdd:
		op.kind = mixedAdd
	case mixedBaseSub:
		op.kind = mixedSub
	case mixedBaseOr:
		op.kind = mixedOr
	case mixedBaseEqual:
		op.kind = mixedEqual
	case mixedBaseNotEqual:
		op.kind = mixedNotEqual
	case mixedBaseNegate:
		op.kind = mixedNegate
	case mixedBaseLess:
		if signed {
			op.kind = mixedSignedLess
		} else {
			op.kind = mixedUnsignedLess
		}
	case mixedBaseLessEqual:
		if signed {
			op.kind = mixedSignedLessEqual
		} else {
			op.kind = mixedUnsignedLessEqual
		}
	case mixedBaseGreater:
		if signed {
			op.kind = mixedSignedGreater
		} else {
			op.kind = mixedUnsignedGreater
		}
	case mixedBaseGreaterEqual:
		if signed {
			op.kind = mixedSignedGreaterEqual
		} else {
			op.kind = mixedUnsignedGreaterEqual
		}
	case mixedBaseShl, mixedBaseShr:
		// Only unsigned or proven nonnegative constant counts bypass the shift check.
		if op.other >= types.Uint && op.other <= types.Uintptr {
			if op.kind == mixedBaseShl {
				op.kind = mixedShl
			} else if signed {
				op.kind = mixedSignedShr
			} else {
				op.kind = mixedUnsignedShr
			}
		}
	}
}

func (r *mixedTypedRunner) Run(input ixgo.Value) ixgo.Value { return r.RunArgs([]ixgo.Value{input}) }

// RunArgs requires the exact parameter count and exact built-in dynamic types.
// Named types and convertible values are not accepted.
func (r *mixedTypedRunner) RunArgs(args []ixgo.Value) ixgo.Value {
	if len(args) != len(r.params) {
		panic("ixgo: invalid mixed runner argument count")
	}
	var initial [3]mixedFrame
	frames := initial[:]
	depth := 0
	regs := &frames[0].regs
	frames[0].runner = r
	copy(regs[:], r.constants)
	for i, k := range r.params {
		regs[i] = mixedUnbox(args[i], k)
	}
	// Frame registers stay in place across calls and returns.
	code, pc := r.code, 0
	for {
		op := &code[pc]
		switch op.kind {
		case mixedAdd:
			regs[op.dst] = regs[op.x] + regs[op.y]
		case mixedSub:
			regs[op.dst] = regs[op.x] - regs[op.y]
		case mixedOr:
			regs[op.dst] = regs[op.x] | regs[op.y]
		case mixedShl:
			regs[op.dst] = regs[op.x] << regs[op.y]
		case mixedUnsignedShr:
			regs[op.dst] = regs[op.x] >> regs[op.y]
		case mixedSignedShr:
			regs[op.dst] = uint64(int64(regs[op.x]) >> regs[op.y])
		case mixedUnsignedLess:
			regs[op.dst] = mixedBool(regs[op.x] < regs[op.y])
		case mixedSignedLess:
			regs[op.dst] = mixedBool(int64(regs[op.x]) < int64(regs[op.y]))
		case mixedUnsignedLessEqual:
			regs[op.dst] = mixedBool(regs[op.x] <= regs[op.y])
		case mixedSignedLessEqual:
			regs[op.dst] = mixedBool(int64(regs[op.x]) <= int64(regs[op.y]))
		case mixedUnsignedGreater:
			regs[op.dst] = mixedBool(regs[op.x] > regs[op.y])
		case mixedSignedGreater:
			regs[op.dst] = mixedBool(int64(regs[op.x]) > int64(regs[op.y]))
		case mixedUnsignedGreaterEqual:
			regs[op.dst] = mixedBool(regs[op.x] >= regs[op.y])
		case mixedSignedGreaterEqual:
			regs[op.dst] = mixedBool(int64(regs[op.x]) >= int64(regs[op.y]))
		case mixedEqual:
			regs[op.dst] = mixedBool(regs[op.x] == regs[op.y])
		case mixedNotEqual:
			regs[op.dst] = mixedBool(regs[op.x] != regs[op.y])
		case mixedNegate:
			regs[op.dst] = -regs[op.x]
		case mixedNot:
			regs[op.dst] = mixedBool(regs[op.x] == 0)
		case mixedBaseMove:
			regs[op.dst] = regs[op.x]
		case mixedBaseBranch:
			if regs[op.x] != 0 {
				pc = op.truePC
			} else {
				pc = op.falsePC
			}
			continue
		case mixedBaseJump:
			pc = op.target
			continue
		case mixedBaseCall:
			frames[depth].code, frames[depth].pc = code, pc
			depth++
			if depth == len(frames) {
				frames = append(frames, mixedFrame{})
			}
			caller := &frames[depth-1]
			next := &frames[depth]
			regs = &next.regs
			// Constants are never destinations; SSA definitions and Phi scratch
			// are written before use, so the same callee needs no reinitialization.
			if next.runner != op.callee {
				copy(regs[:], op.callee.constants)
				next.runner = op.callee
			}
			for i, src := range op.args {
				regs[i] = caller.regs[src]
			}
			code, pc = op.callee.code, 0
			continue
		case mixedBaseReturn:
			if depth == 0 {
				if len(op.results) == 0 {
					return nil
				}
				if len(op.results) == 1 {
					return mixedBox(regs[op.results[0]], r.results[0])
				}
				ret := make(ixgo.Tuple, len(op.results))
				for i, src := range op.results {
					ret[i] = mixedBox(regs[src], r.results[i])
				}
				return ret
			}
			caller := &frames[depth-1]
			destinations := caller.code[caller.pc].results
			for i, src := range op.results {
				caller.regs[destinations[i]] = regs[src]
			}
			code, pc, regs = caller.code, caller.pc+1, &caller.regs
			depth--
			continue
		default:
			regs[op.dst] = mixedEvaluate(op, regs[op.x], regs[op.y])
		}
		pc++
	}
}

// Frames belong to one invocation; compiled graphs are immutable and shareable.
type mixedFrame struct {
	runner *mixedTypedRunner
	code   []mixedOp
	pc     int
	regs   [mixedArgsLimit]uint64
}

func mixedBool(b bool) uint64 {
	if b {
		return 1
	}
	return 0
}

func mixedUnbox(v ixgo.Value, k types.BasicKind) uint64 {
	switch k {
	case types.Bool:
		return mixedBool(v.(bool))
	case types.Int:
		return uint64(v.(int))
	case types.Int8:
		return uint64(v.(int8))
	case types.Int16:
		return uint64(v.(int16))
	case types.Int32:
		return uint64(v.(int32))
	case types.Int64:
		return uint64(v.(int64))
	case types.Uint:
		return uint64(v.(uint))
	case types.Uint8:
		return uint64(v.(uint8))
	case types.Uint16:
		return uint64(v.(uint16))
	case types.Uint32:
		return uint64(v.(uint32))
	case types.Uint64:
		return uint64(v.(uint64))
	case types.Uintptr:
		return uint64(v.(uintptr))
	case types.Float32:
		return uint64(math.Float32bits(v.(float32)))
	case types.Float64:
		return math.Float64bits(v.(float64))
	}
	panic("ixgo: invalid mixed argument kind")
}

func mixedBox(v uint64, k types.BasicKind) ixgo.Value {
	switch k {
	case types.Bool:
		return v != 0
	case types.Int:
		return int(v)
	case types.Int8:
		return int8(v)
	case types.Int16:
		return int16(v)
	case types.Int32:
		return int32(v)
	case types.Int64:
		return int64(v)
	case types.Uint:
		return uint(v)
	case types.Uint8:
		return uint8(v)
	case types.Uint16:
		return uint16(v)
	case types.Uint32:
		return uint32(v)
	case types.Uint64:
		return uint64(v)
	case types.Uintptr:
		return uintptr(v)
	case types.Float32:
		return math.Float32frombits(uint32(v))
	case types.Float64:
		return math.Float64frombits(v)
	}
	panic("ixgo: invalid mixed result kind")
}

func mixedBits[T mixedScalar](v T) uint64 {
	switch n := any(v).(type) {
	case int:
		return uint64(n)
	case int8:
		return uint64(n)
	case int16:
		return uint64(n)
	case int32:
		return uint64(n)
	case int64:
		return uint64(n)
	case uint:
		return uint64(n)
	case uint8:
		return uint64(n)
	case uint16:
		return uint64(n)
	case uint32:
		return uint64(n)
	case uint64:
		return uint64(n)
	case uintptr:
		return uint64(n)
	case float32:
		return uint64(math.Float32bits(n))
	case float64:
		return math.Float64bits(n)
	}
	panic("ixgo: invalid mixed scalar")
}

func mixedConstant(c *ssa.Const) (uint64, bool) {
	if c.Value == nil {
		return 0, false
	}
	k, ok := mixedKind(c.Type())
	if !ok {
		return 0, false
	}
	switch k {
	case types.Bool:
		return mixedBool(constant.BoolVal(c.Value)), true
	case types.Int:
		v, ok := mixedConstantValue[int](c.Value)
		return uint64(v), ok
	case types.Int8:
		v, ok := mixedConstantValue[int8](c.Value)
		return uint64(v), ok
	case types.Int16:
		v, ok := mixedConstantValue[int16](c.Value)
		return uint64(v), ok
	case types.Int32:
		v, ok := mixedConstantValue[int32](c.Value)
		return uint64(v), ok
	case types.Int64:
		v, ok := mixedConstantValue[int64](c.Value)
		return uint64(v), ok
	case types.Uint:
		v, ok := mixedConstantValue[uint](c.Value)
		return uint64(v), ok
	case types.Uint8:
		v, ok := mixedConstantValue[uint8](c.Value)
		return uint64(v), ok
	case types.Uint16:
		v, ok := mixedConstantValue[uint16](c.Value)
		return uint64(v), ok
	case types.Uint32:
		v, ok := mixedConstantValue[uint32](c.Value)
		return uint64(v), ok
	case types.Uint64:
		v, ok := mixedConstantValue[uint64](c.Value)
		return uint64(v), ok
	case types.Uintptr:
		v, ok := mixedConstantValue[uintptr](c.Value)
		return uint64(v), ok
	case types.Float32:
		v, ok := mixedConstantValue[float32](c.Value)
		return uint64(math.Float32bits(v)), ok
	case types.Float64:
		v, ok := mixedConstantValue[float64](c.Value)
		return math.Float64bits(v), ok
	}
	return 0, false
}

func mixedConversion[T mixedScalar](v T, k types.BasicKind) uint64 {
	switch k {
	case types.Int:
		return uint64(int(v))
	case types.Int8:
		return uint64(int8(v))
	case types.Int16:
		return uint64(int16(v))
	case types.Int32:
		return uint64(int32(v))
	case types.Int64:
		return uint64(int64(v))
	case types.Uint:
		return uint64(uint(v))
	case types.Uint8:
		return uint64(uint8(v))
	case types.Uint16:
		return uint64(uint16(v))
	case types.Uint32:
		return uint64(uint32(v))
	case types.Uint64:
		return uint64(uint64(v))
	case types.Uintptr:
		return uint64(uintptr(v))
	case types.Float32:
		return uint64(math.Float32bits(float32(v)))
	case types.Float64:
		return math.Float64bits(float64(v))
	}
	panic("ixgo: invalid mixed conversion")
}

func mixedEvaluate(op *mixedOp, x, y uint64) uint64 {
	if op.scalar == types.Bool {
		switch op.kind {
		case mixedBaseNot:
			return mixedBool(x == 0)
		case mixedBaseEqual:
			return mixedBool(x == y)
		case mixedBaseNotEqual:
			return mixedBool(x != y)
		}
	}
	// Shift counts retain their own width and signedness, not the left operand's.
	if op.kind == mixedBaseShl || op.kind == mixedBaseShr {
		switch op.other {
		case types.Int:
			if int(y) < 0 {
				return uint64(1) << int64(y)
			}
		case types.Int8:
			if int8(y) < 0 {
				return uint64(1) << int64(y)
			}
		case types.Int16:
			if int16(y) < 0 {
				return uint64(1) << int64(y)
			}
		case types.Int32:
			if int32(y) < 0 {
				return uint64(1) << int64(y)
			}
		case types.Int64:
			if int64(y) < 0 {
				return uint64(1) << int64(y)
			}
		}
	}
	switch op.scalar {
	case types.Int:
		return mixedNumeric(op, int(x), int(y), y)
	case types.Int8:
		return mixedNumeric(op, int8(x), int8(y), y)
	case types.Int16:
		return mixedNumeric(op, int16(x), int16(y), y)
	case types.Int32:
		return mixedNumeric(op, int32(x), int32(y), y)
	case types.Int64:
		return mixedNumeric(op, int64(x), int64(y), y)
	case types.Uint:
		return mixedNumeric(op, uint(x), uint(y), y)
	case types.Uint8:
		return mixedNumeric(op, uint8(x), uint8(y), y)
	case types.Uint16:
		return mixedNumeric(op, uint16(x), uint16(y), y)
	case types.Uint32:
		return mixedNumeric(op, uint32(x), uint32(y), y)
	case types.Uint64:
		return mixedNumeric(op, uint64(x), uint64(y), y)
	case types.Uintptr:
		return mixedNumeric(op, uintptr(x), uintptr(y), y)
	case types.Float32:
		return mixedNumeric(op, math.Float32frombits(uint32(x)), math.Float32frombits(uint32(y)), y)
	case types.Float64:
		return mixedNumeric(op, math.Float64frombits(x), math.Float64frombits(y), y)
	}
	panic("ixgo: invalid mixed opcode")
}

func mixedNumeric[T mixedScalar](op *mixedOp, x, y T, shift uint64) uint64 {
	switch op.kind {
	case mixedConvert:
		return mixedConversion(x, op.other)
	case mixedBaseAdd:
		return mixedBits(x + y)
	case mixedBaseSub:
		return mixedBits(x - y)
	case mixedBaseMul:
		return mixedBits(x * y)
	case mixedBaseQuo:
		return mixedBits(x / y)
	case mixedBaseRem:
		return mixedBits(mixedRemainder(x, y))
	case mixedBaseAnd:
		return mixedBits(T(uint64(x) & uint64(y)))
	case mixedBaseOr:
		return mixedBits(T(uint64(x) | uint64(y)))
	case mixedBaseXor:
		return mixedBits(T(uint64(x) ^ uint64(y)))
	case mixedBaseAndNot:
		return mixedBits(T(uint64(x) &^ uint64(y)))
	case mixedBaseShl:
		return mixedBits(T(uint64(x) << shift))
	case mixedBaseShr:
		var zero T
		if zero-1 > zero {
			return mixedBits(T(uint64(x) >> shift))
		}
		return mixedBits(T(int64(x) >> shift))
	case mixedBaseLess:
		return mixedBool(x < y)
	case mixedBaseLessEqual:
		return mixedBool(x <= y)
	case mixedBaseGreater:
		return mixedBool(x > y)
	case mixedBaseGreaterEqual:
		return mixedBool(x >= y)
	case mixedBaseEqual:
		return mixedBool(x == y)
	case mixedBaseNotEqual:
		return mixedBool(x != y)
	case mixedBaseNegate:
		return mixedBits(-x)
	case mixedBaseComplement:
		return mixedBits(T(^uint64(x)))
	}
	panic("ixgo: invalid mixed numeric opcode")
}

var mixedBinaryOps = map[token.Token]mixedOpKind{token.ADD: mixedBaseAdd, token.SUB: mixedBaseSub, token.MUL: mixedBaseMul, token.QUO: mixedBaseQuo, token.REM: mixedBaseRem, token.AND: mixedBaseAnd, token.OR: mixedBaseOr, token.XOR: mixedBaseXor, token.AND_NOT: mixedBaseAndNot, token.SHL: mixedBaseShl, token.SHR: mixedBaseShr, token.LSS: mixedBaseLess, token.LEQ: mixedBaseLessEqual, token.GTR: mixedBaseGreater, token.GEQ: mixedBaseGreaterEqual, token.EQL: mixedBaseEqual, token.NEQ: mixedBaseNotEqual}

func mixedBinaryOpKind(op token.Token) (mixedOpKind, bool) {
	k, ok := mixedBinaryOps[op]
	return k, ok
}
func mixedIsInteger(t types.Type) bool {
	b, ok := t.Underlying().(*types.Basic)
	return ok && b.Info()&types.IsInteger != 0
}
func mixedRemainder[T mixedScalar](x, y T) T {
	var z T
	if z-1 > z {
		return T(uint64(x) % uint64(y))
	}
	return T(int64(x) % int64(y))
}
func mixedConstantValue[T mixedScalar](v constant.Value) (T, bool) {
	var z T
	k := reflect.TypeOf(z).Kind()
	if k == reflect.Float32 || k == reflect.Float64 {
		if v.Kind() != constant.Int && v.Kind() != constant.Float {
			return z, false
		}
		n, ok := constant.Float64Val(v)
		return T(n), ok && !math.IsInf(n, 0)
	}
	if v.Kind() != constant.Int {
		return z, false
	}
	if k >= reflect.Uint && k <= reflect.Uintptr {
		n, ok := constant.Uint64Val(v)
		return T(n), ok
	}
	n, ok := constant.Int64Val(v)
	if !ok {
		return z, false
	}
	return T(n), true
}
