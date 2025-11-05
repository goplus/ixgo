package tpl

import (
	"github.com/goplus/ixgo"
)

var tpl_patch = `package tpl

// ListOp converts the matching result of (R % ",") to a flat list.
// R % "," means R *("," R)
func ListOp[T any](in []any, fn func(v any) T) []T {
	next := in[1].([]any)
	ret := make([]T, len(next)+1)
	ret[0] = fn(in[0])
	for i, v := range next {
		ret[i+1] = fn(v.([]any)[1])
	}
	return ret
}
`

func init() {
	ixgo.RegisterPatch("github.com/goplus/xgo/tpl", tpl_patch)
}
