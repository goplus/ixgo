package tpl

import (
	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/ast"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/cl"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/encoding/csv"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/encoding/json"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/encoding/regexp"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/encoding/regexposix"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/encoding/xml"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/matcher"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/parser"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/scanner"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/token"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/types"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/variant"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/variant/builtin"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/variant/delay"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/variant/math"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl/variant/time"
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

var variant_patch = `
package variant

import . "github.com/goplus/xgo/tpl/variant"

// ListOp converts the matching result of (R % ",") to a flat list.
// R % "," means R *("," R)
func ListOp[T any](in []any, fn func(v any) T) []T {
	next := in[1].([]any)
	ret := make([]T, len(next)+1)
	ret[0] = fn(Eval(in[0]))
	for i, v := range next {
		ret[i+1] = fn(Eval(v.([]any)[1]))
	}
	return ret
}
`

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
	ixgo.RegisterPatch("github.com/goplus/xgo/tpl", tpl_patch)
	ixgo.RegisterPatch("github.com/goplus/xgo/tpl/variant", variant_patch)
	ixgo.RegisterPatch("github.com/goplus/xgo/tpl/variant/delay", delay_patch)
}
