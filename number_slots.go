package ixgo

import (
	"go/constant"
	"go/token"
	"go/types"

	"golang.org/x/tools/go/ssa"
)

// Deliberately exclude named types, conversions, calls and addressable locals.
func numberKind(t types.Type) types.BasicKind {
	b, ok := t.(*types.Basic)
	if !ok {
		return types.Invalid
	}
	switch b.Kind() {
	case types.Int, types.Uint64, types.Bool:
		return b.Kind()
	}
	return types.Invalid
}

// Setup and boxed argument/result transfers outweigh the savings in tiny helpers.
const minNumberSlotInstructions = 16

func (p *function) compileNumberSlots() {
	ctx := p.Interp.ctx
	if !ctx.NumberSlots || (memoPlugin != nil && ctx.Mode&DisableMemo == 0) || p.memoFlags != 0 || p.Interp.execution != nil || ctx.execution != nil || ctx.evalMode || ctx.evalCallFn != nil || ctx.debugFunc != nil || ctx.entryFunc != nil || ctx.leaveFunc != nil || ctx.panicFunc != nil || ctx.Mode&EnableTracing != 0 || ctx.BuilderMode&ssa.GlobalDebug != 0 || len(p.Fn.FreeVars) != 0 || p.Fn.Recover != nil {
		return
	}
	instructions := 0
	for _, b := range p.Fn.Blocks {
		instructions += len(b.Instrs)
	}
	if instructions < minNumberSlotInstructions {
		return
	}
	slots := make(map[ssa.Value]int)
	var values []ssa.Value
	add := func(v ssa.Value) bool {
		_, isConst := v.(*ssa.Const)
		if v == nil || (numberKind(v.Type()) == types.Invalid && !(isConst && v.Type() == types.Typ[types.Uint])) {
			return false
		}
		if _, ok := slots[v]; !ok {
			slots[v] = len(values)
			values = append(values, v)
		}
		return true
	}
	for _, v := range p.Fn.Params {
		if !add(v) {
			return
		}
	}
	for _, b := range p.Fn.Blocks {
		for _, ins := range b.Instrs {
			switch x := ins.(type) {
			case *ssa.BinOp:
				switch x.Op {
				case token.ADD, token.SUB, token.MUL, token.AND, token.OR, token.XOR, token.AND_NOT, token.SHL, token.SHR, token.EQL, token.NEQ, token.LSS, token.LEQ, token.GTR, token.GEQ:
				default:
					return
				}
			case *ssa.UnOp:
				if x.Op != token.SUB && x.Op != token.XOR && x.Op != token.NOT {
					return
				}
			case *ssa.Phi, *ssa.If, *ssa.Jump, *ssa.Return:
			default:
				return
			}
			if v, ok := ins.(ssa.Value); ok && !add(v) {
				return
			}
			for _, op := range ins.Operands(nil) {
				if op != nil && !add(*op) {
					return
				}
			}
		}
	}
	if len(values) == 0 {
		return
	}
	var code []func(*frame)
	var source []ssa.Instruction
	blocks := make([]int, len(p.Fn.Blocks))
	for _, b := range p.Fn.Blocks {
		blocks[b.Index] = len(code)
		for _, ins := range b.Instrs {
			var f func(*frame)
			switch x := ins.(type) {
			case *ssa.BinOp:
				d, a, c := slots[x], slots[x.X], slots[x.Y]
				op := numberBinary(x.Op, numberKind(x.X.Type()), numberKind(x.Y.Type()))
				f = func(fr *frame) {
					fr.numberSlots()[d] = op(fr.numberSlots()[a], fr.numberSlots()[c])
				}
			case *ssa.UnOp:
				d, a, op := slots[x], slots[x.X], x.Op
				f = func(fr *frame) {
					v := fr.numberSlots()[a]
					switch op {
					case token.SUB:
						v = -v
					case token.XOR:
						v = ^v
					case token.NOT:
						v ^= 1
					}
					if numberKind(x.Type()) == types.Int {
						v = uint64(int(v))
					}
					fr.numberSlots()[d] = v
				}
			case *ssa.Phi:
				if b.Instrs[0] != ins {
					continue
				}
				var dst []int
				var edges [][]int
				for _, in := range b.Instrs {
					ph, ok := in.(*ssa.Phi)
					if !ok {
						break
					}
					dst = append(dst, slots[ph])
					e := make([]int, len(ph.Edges))
					for j, v := range ph.Edges {
						e[j] = slots[v]
					}
					edges = append(edges, e)
				}
				f = func(fr *frame) {
					pred := -1
					for j, bb := range b.Preds {
						if int(fr.pred) == bb.Index {
							pred = j
							break
						}
					}
					if pred < 0 {
						panic("numeric phi predecessor")
					}
					scratch := fr.numberSlots()[len(values):]
					for j := range dst {
						scratch[j] = fr.numberSlots()[edges[j][pred]]
					}
					for j, d := range dst {
						fr.numberSlots()[d] = scratch[j]
					}
				}
			case *ssa.If:
				c := slots[x.Cond]
				f = func(fr *frame) {
					next := 1
					if fr.numberSlots()[c] != 0 {
						next = 0
					}
					fr.pred = int32(b.Index)
					fr.block = b.Succs[next]
					fr.ipc = int32(blocks[fr.block.Index])
				}
			case *ssa.Jump:
				f = func(fr *frame) {
					fr.pred = int32(b.Index)
					fr.block = b.Succs[0]
					fr.ipc = int32(blocks[fr.block.Index])
				}
			case *ssa.Return:
				f = func(fr *frame) {
					for j, v := range x.Results {
						n := fr.numberSlots()[slots[v]]
						switch numberKind(v.Type()) {
						case types.Int:
							fr.stack[j] = int(n)
						case types.Uint64:
							fr.stack[j] = n
						case types.Bool:
							fr.stack[j] = n != 0
						}
					}
					fr.ipc = -1
				}
			}
			code = append(code, f)
			source = append(source, ins)
		}
	}
	initial := make([]uint64, len(values))
	for j, v := range values {
		if c, ok := v.(*ssa.Const); ok {
			switch c.Type().(*types.Basic).Kind() {
			case types.Bool:
				if constant.BoolVal(c.Value) {
					initial[j] = 1
				}
			case types.Int:
				n, _ := constant.Int64Val(c.Value)
				initial[j] = uint64(n)
			case types.Uint64, types.Uint:
				initial[j], _ = constant.Uint64Val(c.Value)
			}
		}
	}
	first := code[0]
	code[0] = func(fr *frame) {
		if len(fr.numberSlots()) != 2*len(values) {
			fr.setNumberSlots(make([]uint64, 2*len(values)))
		}
		copy(fr.numberSlots(), initial)
		for j, v := range p.Fn.Params {
			n := uint64(0)
			switch numberKind(v.Type()) {
			case types.Int:
				n = uint64(fr.stack[p.nres+j].(int))
			case types.Uint64:
				n = fr.stack[p.nres+j].(uint64)
			case types.Bool:
				if fr.stack[p.nres+j].(bool) {
					n = 1
				}
			}
			fr.numberSlots()[slots[v]] = n
		}
		first(fr)
	}
	p.Instrs, p.ssaInstrs, p.Blocks = code, source, blocks
	p.numberSlots = len(values)
}

func (fr *frame) numberSlots() []uint64 {
	if fr.memoState == nil {
		fr.memoState = &interpreterMemoState{}
	}
	return fr.memoState.numberSlots
}

func (fr *frame) setNumberSlots(values []uint64) {
	if fr.memoState == nil {
		fr.memoState = &interpreterMemoState{}
	}
	fr.memoState.numberSlots = values
}

func numberBinary(op token.Token, kind, rhs types.BasicKind) func(uint64, uint64) uint64 {
	boolean := func(b bool) uint64 {
		if b {
			return 1
		}
		return 0
	}
	signed := kind == types.Int
	normalize := func(n uint64) uint64 {
		if signed {
			return uint64(int(n))
		}
		return n
	}
	switch op {
	case token.ADD:
		return func(a, b uint64) uint64 { return normalize(a + b) }
	case token.SUB:
		return func(a, b uint64) uint64 { return normalize(a - b) }
	case token.MUL:
		return func(a, b uint64) uint64 { return normalize(a * b) }
	case token.AND:
		return func(a, b uint64) uint64 { return a & b }
	case token.OR:
		return func(a, b uint64) uint64 { return a | b }
	case token.XOR:
		return func(a, b uint64) uint64 { return a ^ b }
	case token.AND_NOT:
		return func(a, b uint64) uint64 { return a &^ b }
	case token.SHL, token.SHR:
		return func(a, b uint64) uint64 {
			if rhs == types.Int && int(b) < 0 {
				return a << int(b)
			}
			if op == token.SHL {
				return normalize(a << b)
			}
			if signed {
				return uint64(int(a) >> b)
			}
			return a >> b
		}
	case token.EQL:
		return func(a, b uint64) uint64 { return boolean(a == b) }
	case token.NEQ:
		return func(a, b uint64) uint64 { return boolean(a != b) }
	case token.LSS:
		return func(a, b uint64) uint64 {
			if signed {
				return boolean(int(a) < int(b))
			}
			return boolean(a < b)
		}
	case token.LEQ:
		return func(a, b uint64) uint64 {
			if signed {
				return boolean(int(a) <= int(b))
			}
			return boolean(a <= b)
		}
	case token.GTR:
		return func(a, b uint64) uint64 {
			if signed {
				return boolean(int(a) > int(b))
			}
			return boolean(a > b)
		}
	case token.GEQ:
		return func(a, b uint64) uint64 {
			if signed {
				return boolean(int(a) >= int(b))
			}
			return boolean(a >= b)
		}
	}
	panic("unsupported numeric operation")
}
