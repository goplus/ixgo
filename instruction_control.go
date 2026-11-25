package ixgo

// makeControlInstruction builds a closure solely from execution metadata.
// Other opcodes retain their existing lowering until their descriptions can
// supply all execution inputs.
func makeControlInstruction(d typedInstruction) (func(*frame), bool) {
	switch d.Op {
	case opJump:
		target := d.Aux.(jumpAux).Target
		return func(fr *frame) { fr.jump(target) }, true
	case opIf:
		aux := d.Aux.(ifAux)
		if aux.Constant {
			if aux.Condition {
				return func(fr *frame) { fr.jump(aux.Targets[0]) }, true
			}
			return func(fr *frame) { fr.jump(aux.Targets[1]) }, true
		}
		condition := d.Args[0]
		if aux.BasicBool {
			return func(fr *frame) {
				if fr.reg(condition).(bool) {
					fr.jump(aux.Targets[0])
				} else {
					fr.jump(aux.Targets[1])
				}
			}, true
		}
		return func(fr *frame) {
			if fr.bool(condition) {
				fr.jump(aux.Targets[0])
			} else {
				fr.jump(aux.Targets[1])
			}
		}, true
	case opPhi:
		aux := d.Aux.(phiAux)
		ir, ie, preds := aux.Results, aux.Edges, aux.Preds
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
		}, true
	default:
		return nil, false
	}
}
