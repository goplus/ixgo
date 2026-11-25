package ixgo

// makeControlInstruction builds a closure solely from execution metadata.
// Other opcodes retain their existing lowering until their descriptions can
// supply all execution inputs.
func makeControlInstruction(d typedInstruction) (func(*frame), bool) {
	switch d.Op {
	case opJump:
		target := d.Aux.(branchAux).Targets[0]
		return func(fr *frame) { fr.jump(target) }, true
	case opIf:
		targets := d.Aux.(branchAux).Targets
		condition, yes, no := d.Args[0], targets[0], targets[1]
		return func(fr *frame) {
			target := no
			if fr.bool(condition) {
				target = yes
			}
			fr.jump(target)
		}, true
	case opPhi:
		aux := d.Aux.(phiAux)
		results, edges, preds := aux.Results, aux.Edges, aux.Preds
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
			if cap(fr.phiVals) < len(results) {
				fr.phiVals = make([]value, len(results))
			}
			// Read all edges before writing results to preserve cyclic assignments.
			values := fr.phiVals[:len(results)]
			for i := range results {
				values[i] = fr.reg(edges[i][predIndex])
			}
			for i := range results {
				fr.setReg(results[i], values[i])
			}
		}, true
	default:
		return nil, false
	}
}
