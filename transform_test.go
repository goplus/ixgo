package ixgo_test

import (
	"path/filepath"
	"testing"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/transform"
)

func TestTestdataFilesAfterTransform(t *testing.T) {
	for _, name := range testdataTests {
		t.Run(name, func(t *testing.T) {
			input := filepath.Join("testdata", name)
			ctx := ixgo.NewContext(0)
			pkg, err := ctx.LoadFile(input, nil)
			if err != nil {
				t.Fatal(err)
			}
			if err := transform.Transform(pkg); err != nil {
				t.Fatal(err)
			}
			if _, err := ctx.RunPkg(pkg, input, nil); err != nil {
				t.Fatal(err)
			}
		})
	}
}
