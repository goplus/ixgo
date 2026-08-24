//go:build go1.18

package main

import (
	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/xgo/tpl"
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
