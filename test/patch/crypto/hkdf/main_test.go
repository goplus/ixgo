//go:build go1.24

package main

import (
	"os"
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/bytes"
	_ "github.com/goplus/ixgo/pkg/crypto/hkdf"
	_ "github.com/goplus/ixgo/pkg/crypto/sha256"
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
