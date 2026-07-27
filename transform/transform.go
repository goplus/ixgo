package transform

import (
	"github.com/goplus/ixgo/transform/sccp"
	"golang.org/x/tools/go/ssa"
)

// Pass transforms a built SSA package.
type Pass interface {
	Run(*ssa.Package) error
}

var passes = []Pass{
	sccp.Pass{},
}

// Transform runs the default SSA transformation pipeline.
func Transform(pkg *ssa.Package) error {
	for _, pass := range passes {
		if err := pass.Run(pkg); err != nil {
			return err
		}
	}
	return nil
}
