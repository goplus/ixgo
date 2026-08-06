package main

import (
	"go/constant"
	"reflect"

	"github.com/goplus/ixgo/testdata/direct_call/pkg"
)

type number interface {
	Value() int
}

func value(n number) int {
	return n.Value()
}

func main() {
	if got := pkg.Add(20, 22); got != 42 {
		panic(got)
	}
	if got := pkg.Inspect(reflect.TypeOf(0), constant.MakeInt64(42)); got != 42 {
		panic(got)
	}

	n := pkg.Number(7)
	if got := n.Value(); got != 7 {
		panic(got)
	}
	if got := value(&n); got != 7 {
		panic(got)
	}
}
