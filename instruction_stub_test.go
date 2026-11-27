//go:build !instruction

package ixgo

import "testing"

func TestInstructionStub(t *testing.T) {
	d := describeInstruction(nil, nil, nil, nil)
	if makeInstruction(nil, d) != nil {
		t.Fatal("stub returned an executor")
	}
	for _, mode := range []Mode{EnableInstructionDescriptions, EnableInstructionExecution, EnableInstructionDescriptions | EnableInstructionExecution} {
		i, err := NewContext(DisableMemo|mode).LoadInterp("instruction_stub.go", `package main
 func calc(x int) int {if x>0{return x+1};return x-1}
 `)
		if err != nil {
			t.Fatal(err)
		}
		for _, p := range i.funcs {
			if p.Code != nil {
				t.Fatal("default build allocated instruction metadata")
			}
		}
		if got, err := i.RunFunc("calc", 2); err != nil || got != 3 {
			t.Fatalf("calc = %v,%v", got, err)
		}
		i.UnsafeRelease()
	}
}
