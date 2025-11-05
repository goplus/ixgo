package delay

import (
	"github.com/goplus/ixgo"
)

var delay_patch = `package delay

import "github.com/goplus/xgo/tpl/variant"

// ListOp delays to convert the matching result of (R % ",") to a flat list.
// R % "," means R *("," R)
func ListOp[T any](in []any, op func(v any) T, fn func(flat []T)) any {
	return func() any {
		fn(variant.ListOp(in, op))
		return nil
	}
}`

func init() {
	ixgo.RegisterPatch("github.com/goplus/xgo/tpl/variant/delay", delay_patch)
}
