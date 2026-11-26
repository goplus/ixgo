//go:build instruction

package ixgo

import (
	"reflect"
	"sync/atomic"
)

// makeInstruction binds execution metadata to the supplied runtime context.
// It returns nil when the description has no executor implementation.
// Unsupported instructions retain their existing lowering.
func instructionGoCall(interp *Interp, desc callDescription) func(*frame) {
	return func(fr *frame) {
		fn, args := desc.prepare(interp, fr)
		root := rootFrame()
		run := interp.prepareGoCall(root, fn, args, nil)
		atomic.AddInt32(&interp.goroutines, 1)
		go func() {
			defer atomic.AddInt32(&interp.goroutines, -1)
			if interp.ctx.RunContext != nil {
				defer func() {
					if e := recover(); e != nil {
						interp.cherror <- PanicError{stack: debugStack(root), Value: e}
					}
				}()
			}
			run()
		}()
	}
}

func instructionDeferCall(interp *Interp, desc callDescription, id register) func(*frame) {
	if id < 0 {
		return func(fr *frame) {
			fn, args := desc.prepare(interp, fr)
			fr._defer = &_defer{fn: fn, args: args, tail: fr._defer}
		}
	}
	return func(fr *frame) {
		fn, args := desc.prepare(interp, fr)
		d := &_defer{fn: fn, args: args}
		head, ok := deferStackHead(fr.reg(id))
		if !ok {
			panic("range-over-func defer: missing deferstack token")
		}
		pushDeferAt(head, d)
	}
}

func makeInstruction(interp *Interp, d typedInstruction) func(*frame) {
	switch d.Op {
	case opMakeClosure:
		aux := d.Aux.(closureAux)
		if aux.Body == nil {
			return nil
		}
		body, signature, bindings, result := aux.Body, aux.Signature, aux.Bindings, d.Result
		return func(fr *frame) {
			var env []value
			for _, binding := range bindings {
				env = append(env, fr.reg(binding))
			}
			fr.setReg(result, body.makeFunction(signature, env).Interface())
		}

	case opGo, opDefer:
		var desc callDescription
		var stack register
		switch aux := d.Aux.(type) {
		case dynamicCallAux:
			if aux.Target < 0 {
				return nil
			}
			desc = callDescription{receiver: aux.Target, args: aux.Arguments}
			stack = aux.DeferStack
		case bodyCallAux:
			if aux.Body == nil {
				return nil
			}
			desc = callDescription{target: aux.Body, args: aux.Arguments}
			stack = aux.DeferStack
		case externalCallAux:
			if !aux.External.IsValid() {
				return nil
			}
			desc = callDescription{target: aux.External, args: aux.Arguments}
			stack = aux.DeferStack
		default:
			return nil
		}
		if d.Op == opGo {
			return instructionGoCall(interp, desc)
		}
		return instructionDeferCall(interp, desc, stack)

	case opCall:
		switch aux := d.Aux.(type) {
		case dynamicCallAux:
			if aux.Target < 0 {
				return nil
			}
			target, result, args, dests := aux.Target, d.Result, aux.Arguments, aux.ExtractResults
			if supportFuncVal && interp.ctx.Mode&DisableDynamicFuncCallAnalysis == 0 {
				return dynamicFunCall(interp, target, result, args, dests)
			}
			return func(fr *frame) {
				interp.callExternalByStack(fr, reflect.ValueOf(fr.reg(target)), result, args, dests)
			}
		case directCallAux:
			if aux.Adapter == nil {
				return nil
			}
			adapter, args, result := aux.Adapter, aux.Arguments, d.Result
			return func(fr *frame) {
				interp.invokeDirectCall(fr, adapter, result, args)
			}
		case externalCallAux:
			ext, args, result, dests := aux.External, aux.Arguments, d.Result, aux.ExtractResults
			if !ext.IsValid() || ext.Kind() != reflect.Func {
				return nil
			}
			if typ := ext.Type(); typ.NumIn() > 0 && typ.In(0) == typFramePtr {
				return func(fr *frame) { interp.callExternalWithFrameByStack(fr, ext, result, args, dests) }
			}
			return func(fr *frame) { interp.callExternalByStack(fr, ext, result, args, dests) }
		case bodyCallAux:
			body := aux.Body
			if body == nil {
				return nil
			}
			args, result, dests := aux.Arguments, d.Result, aux.ExtractResults
			if body.Recover == nil {
				if run := makeCallSpecializerInstr(interp, body, args, dests); run != nil {
					return run
				}
				switch body.nres {
				case 0:
					return func(fr *frame) { interp.callFunctionByStackNoRecover0(fr, body, result, args) }
				case 1:
					return func(fr *frame) { interp.callFunctionByStackNoRecover1(fr, body, result, args) }
				default:
					return func(fr *frame) { interp.callFunctionByStackNoRecoverN(fr, body, result, args, dests) }
				}
			}
			switch body.nres {
			case 0:
				return func(fr *frame) { interp.callFunctionByStack0(fr, body, result, args) }
			case 1:
				return func(fr *frame) { interp.callFunctionByStack1(fr, body, result, args) }
			default:
				return func(fr *frame) { interp.callFunctionByStackN(fr, body, result, args, dests) }
			}
		}
		return nil
	case opJump:
		target := d.Aux.(jumpAux).Target
		return func(fr *frame) { fr.jump(target) }
	case opIf:
		aux := d.Aux.(ifAux)
		if aux.Constant {
			if aux.Condition {
				return func(fr *frame) { fr.jump(aux.Targets[0]) }
			}
			return func(fr *frame) { fr.jump(aux.Targets[1]) }
		}
		condition := d.Args[0]
		if aux.BasicBool {
			return func(fr *frame) {
				if fr.reg(condition).(bool) {
					fr.jump(aux.Targets[0])
				} else {
					fr.jump(aux.Targets[1])
				}
			}
		}
		return func(fr *frame) {
			if fr.bool(condition) {
				fr.jump(aux.Targets[0])
			} else {
				fr.jump(aux.Targets[1])
			}
		}
	case opPhi:
		aux := d.Aux.(phiAux)
		ir, ie, preds := aux.Results, aux.Edges, aux.Preds
		if len(ir) == 0 || len(ie) == 0 {
			return func(*frame) {}
		}
		if len(ir) == 1 {
			result, edges := ir[0], ie[0]
			return func(fr *frame) {
				for i, pred := range preds {
					if fr.pred == pred {
						fr.setReg(result, fr.reg(edges[i]))
						return
					}
				}
			}
		}
		if phiDirectAssignment(ir, ie) {
			return func(fr *frame) {
				for j, pred := range preds {
					if fr.pred == pred {
						for i, result := range ir {
							fr.setReg(result, fr.reg(ie[i][j]))
						}
						return
					}
				}
			}
		}
		return func(fr *frame) {
			predIndex := -1
			for i, pred := range preds {
				if fr.pred == pred {
					predIndex = i
					break
				}
			}
			if predIndex < 0 {
				return
			}
			if cap(fr.phiVals) < len(ir) {
				fr.phiVals = make([]value, len(ir))
			}
			// Read all edges before writing results to preserve cyclic assignments.
			values := fr.phiVals[:len(ir)]
			for i := range ir {
				values[i] = fr.reg(ie[i][predIndex])
			}
			for i := range ir {
				fr.setReg(ir[i], values[i])
			}
		}
	case opBinOp:
		return makeBinaryInstruction(d)
	case opUnOp:
		return makeUnaryInstruction(d)
	}
	return nil
}
