package embed

import (
	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/load/embed"
)

func init() {
	ixgo.RegisterEmbedProcessor(embed.Processor{})
}
