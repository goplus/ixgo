// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.21 && !go1.22
// +build go1.21,!go1.22

package internal

import (
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name:   "internal",
		Path:   "log/internal",
		Deps:   map[string]string{},
		Source: source,
	})
}

var source = "package internal\n\nvar DefaultOutput func(pc uintptr, data []byte) error\n"
