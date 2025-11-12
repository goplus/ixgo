// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.22 && !go1.23
// +build go1.22,!go1.23

package internal

import (
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name:   "internal",
		Path:   "log/slog/internal",
		Deps:   map[string]string{},
		Source: source,
	})
}

var source = "package internal\n\nvar IgnorePC = false\n"
