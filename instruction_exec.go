//go:build instruction

package ixgo

// makeInstruction builds supported closures solely from execution metadata.
// It returns nil when the description has no executor implementation.
// Unsupported instructions retain their existing lowering.
func makeInstruction(d typedInstruction) func(*frame) {
	switch d.Op {
	case opMakeClosure:
		aux := d.Aux.(closureAux)
		if aux.Function.Body == nil {
			return nil
		}
		body, signature, bindings, result := aux.Function.Body, aux.Signature, aux.Bindings, d.Result
		return func(fr *frame) {
			var env []value
			for _, binding := range bindings {
				env = append(env, fr.reg(binding))
			}
			fr.setReg(result, body.makeFunction(signature, env).Interface())
		}
	case opCall:
		aux := d.Aux.(callAux)
		body := aux.Function.Body
		if aux.Kind != callStatic || body == nil {
			return nil
		}
		interp, args, result, dests := body.Interp, aux.Arguments, d.Result, aux.ExtractResults
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
