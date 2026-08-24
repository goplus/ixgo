//go:build go1.18

package main

import "github.com/goplus/xgo/tpl/variant/delay"

func main() {
	n := 0
	f := delay.ListOp([]any{1, []any{[]any{",", 2}}}, func(v any) int { return v.(int) }, func(v []int) { n = len(v) })
	f.(func() any)()
	if n != 2 {
		panic("ListOp")
	}
}
