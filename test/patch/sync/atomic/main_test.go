//go:build go1.19

package main

import (
	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/sync/atomic"
	"os"
	"testing"
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
