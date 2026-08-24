//go:build go1.23

package main

import (
	"os"
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/iter"
)

func TestPatch(t *testing.T) {
	src, e := os.ReadFile("main.go")
	if e != nil {
		t.Fatal(e)
	}
	if _, e = ixgo.NewContext(0).RunFile("main.go", src, nil); e != nil {
		t.Fatal(e)
	}
}
