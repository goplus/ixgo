package sccp

import (
	"bytes"
	"io"
	"testing"
	_ "unsafe"

	"golang.org/x/tools/go/ssa"
)

//go:linkname ssaSanityCheck golang.org/x/tools/go/ssa.sanityCheck
func ssaSanityCheck(*ssa.Function, io.Writer) bool

func assertSSASane(t *testing.T, fn *ssa.Function) {
	t.Helper()
	var diagnostics bytes.Buffer
	if !ssaSanityCheck(fn, &diagnostics) {
		t.Fatalf("invalid SSA after SCCP for %s:\n%s", fn, diagnostics.Bytes())
	}
}
