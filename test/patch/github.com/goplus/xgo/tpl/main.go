//go:build go1.18

package main

import "github.com/goplus/xgo/tpl"

func main() {
	out := tpl.ListOp([]any{1, []any{[]any{",", 2}}}, func(v any) int { return v.(int) })
	if len(out) != 2 || out[0] != 1 || out[1] != 2 {
		panic("ListOp")
	}
}
