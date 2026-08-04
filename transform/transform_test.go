package transform_test

import (
	"strings"
	"testing"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/transform"
)

func runSrc(t *testing.T, src string) error {
	t.Helper()
	ctx := ixgo.NewContext(0)
	pkg, err := ctx.LoadFile("main.go", src)
	if err != nil {
		t.Fatal(err)
	}
	if err := transform.Transform(pkg); err != nil {
		t.Fatal(err)
	}
	_, err = ctx.RunPkg(pkg, "main.go", nil)
	return err
}

// TestInterpAfterTransform runs a transformed program under the ixgo
// interpreter and checks the observable behavior is unchanged.
func TestInterpAfterTransform(t *testing.T) {
	err := runSrc(t, `package main

func compute() int {
	x := 1
	if x == 1 {
		x = 2
	} else {
		x = 3
	}
	sum := 0
	for i := 0; i < 3; i++ {
		if x != 2 {
			sum += 100
		}
		sum++
	}
	return x*10 + sum
}

type T int

func (v T) Double() T { return v * T(2) }

func main() {
	if got := compute(); got != 23 {
		panic("compute: unexpected result")
	}
	if got := T(21).Double(); got != 42 {
		panic("method: unexpected result")
	}
	f := func() uint8 {
		a := uint8(200)
		b := uint8(100)
		return a + b
	}
	if got := f(); got != 44 {
		panic("closure: unexpected result")
	}
}
`)
	if err != nil {
		t.Fatal(err)
	}
}

// TestPanicPreserved checks that operations which must panic at run time
// (division by zero, negative shift) are not folded away.
func TestPanicPreserved(t *testing.T) {
	err := runSrc(t, `package main

func main() {
	z := 0
	_ = 10 / z
}
`)
	if err == nil || !strings.Contains(err.Error(), "divide by zero") {
		t.Fatalf("want divide by zero panic, got %v", err)
	}

	err = runSrc(t, `package main

func main() {
	s := -1
	_ = 1 << s
}
`)
	if err == nil || !strings.Contains(err.Error(), "shift") {
		t.Fatalf("want negative shift panic, got %v", err)
	}
}
