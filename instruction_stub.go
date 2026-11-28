//go:build !instruction

package ixgo

import "golang.org/x/tools/go/ssa"

const instructionSupported = false

// typedInstruction is empty in the default build; descriptions are available
// in the instruction-tagged build.
type typedInstruction struct{}

func describeCallResults(*function, ssa.Value) []register { return nil }
func bindRuntimeRegisters(*typedInstruction)              {}

func describeInstruction(*function, ssa.Instruction, []*ssa.Value, []register) typedInstruction {
	return typedInstruction{}
}

func makeInstruction(*Interp, typedInstruction) func(*frame) { return nil }

func planInstructionLayout(*function) {}

func compileInstruction(*function, ssa.Instruction) (typedInstruction, func(*frame)) {
	panic("instruction backend requires the instruction build tag")
}
