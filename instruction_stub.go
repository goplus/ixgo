//go:build !instruction

package ixgo

import "golang.org/x/tools/go/ssa"

// typedInstruction is empty in the default build; descriptions are available
// in the instruction-tagged build.
type typedInstruction struct{}

func describeInstruction(*function, ssa.Instruction, []*ssa.Value, []register) typedInstruction {
	return typedInstruction{}
}

func makeInstruction(typedInstruction) func(*frame) { return nil }
