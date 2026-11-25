package ixgo

// makeInstruction builds supported closures solely from execution metadata.
// It returns nil when the description has no executor implementation.
// Unsupported instructions retain their existing lowering.
func makeInstruction(d typedInstruction) func(*frame) {
	switch d.Op {
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

// phiDirectAssignment checks that each source survives all earlier writes for
// every predecessor. Self assignments preserve the register's old value.
func phiDirectAssignment(results []register, edges [][]register) bool {
	if len(results) == 0 || len(edges) == 0 {
		return false
	}
	for pred := range edges[0] {
		written := make(map[register]bool, len(results))
		for i, result := range results {
			source := edges[i][pred]
			if written[source] {
				return false
			}
			if source != result {
				written[result] = true
			}
		}
	}
	return true
}
