// Package memo enables optional interpreter memoization.
package memo

import (
	"github.com/goplus/ixgo"
	"golang.org/x/tools/go/ssa"
)

func init() { ixgo.RegisterMemoPlugin(install) }

func install(p ixgo.MemoInstaller, fn *ssa.Function) {
	if scalarEligible(fn) {
		p.InstallScalarMemo()
		return
	}
	if plan := analyzeReceiver(fn); plan != nil {
		p.InstallReceiverMemo(&ixgo.MemoReceiverPlan{
			Receiver: &ixgo.MemoReceiverState{
				Fields:   plan.receiver.Fields,
				Additive: plan.receiver.Additive,
				Affine:   plan.receiver.Affine,
				Bitwise:  plan.receiver.Bitwise,
			},
			Functions: plan.Functions,
			Updates:   convertUpdates(plan.Updates),
		})
		return
	}
	if plan := analyzeAccumulator(fn); plan != nil {
		p.InstallAccumulatorMemo(plan.Output, false)
	} else if plan := analyzeState(fn); plan != nil {
		p.InstallAccumulatorMemo(plan.Output, true)
	}
}

func convertUpdates(src map[*ssa.Store]receiverUpdate) map[*ssa.Store]ixgo.MemoReceiverUpdate {
	dst := make(map[*ssa.Store]ixgo.MemoReceiverUpdate, len(src))
	for store, update := range src {
		dst[store] = ixgo.MemoReceiverUpdate{Field: update.Field, Op: update.Op, Operand: update.Operand}
	}
	return dst
}
