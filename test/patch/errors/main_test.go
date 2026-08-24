//go:build go1.26

package main

import (
	"os"
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/errors"
	_ "github.com/goplus/ixgo/pkg/fmt"
)

func TestPatch(t *testing.T) {
	src, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatal(err)
	}
	if _, err = ixgo.NewContext(0).RunFile("main.go", src, nil); err != nil {
		t.Fatal(err)
	}
}
