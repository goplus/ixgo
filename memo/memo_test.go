package memo_test

import (
	_ "github.com/goplus/ixgo/memo"
	"testing"

	"github.com/goplus/ixgo"
)

func TestImportEnablesMemoPackage(t *testing.T) {
	ctx := ixgo.NewContext(0)
	interp, err := ctx.LoadInterp("memo.go", `package main
func fib(n int) int { if n < 2 { return n }; return fib(n-1)+fib(n-2) }
`)
	if err != nil {
		t.Fatal(err)
	}
	defer interp.UnsafeRelease()
	got, err := interp.RunFunc("fib", 20)
	if err != nil || got != 6765 {
		t.Fatalf("fib(20)=%v, err=%v", got, err)
	}
}
