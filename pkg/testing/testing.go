package testing

import (
	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/load/test"
)

func init() {
	ixgo.RegisterTestProcessor(test.Processor{})
}
