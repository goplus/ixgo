// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package cmp

import (
	_ "cmp"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name:   "cmp",
		Path:   "cmp",
		Deps:   map[string]string{},
		Source: source,
	})
}

var source = "package cmp\n\ntype Ordered interface {\n\t~int | ~int8 | ~int16 | ~int32 | ~int64 |\n\t\t~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |\n\t\t~float32 | ~float64 |\n\t\t~string\n}\n\nfunc Less[T Ordered](x, y T) bool {\n\treturn (isNaN(x) && !isNaN(y)) || x < y\n}\n\nfunc Compare[T Ordered](x, y T) int {\n\txNaN := isNaN(x)\n\tyNaN := isNaN(y)\n\tif xNaN {\n\t\tif yNaN {\n\t\t\treturn 0\n\t\t}\n\t\treturn -1\n\t}\n\tif yNaN {\n\t\treturn +1\n\t}\n\tif x < y {\n\t\treturn -1\n\t}\n\tif x > y {\n\t\treturn +1\n\t}\n\treturn 0\n}\n\nfunc isNaN[T Ordered](x T) bool {\n\treturn x != x\n}\n\nfunc Or[T comparable](vals ...T) T {\n\tvar zero T\n\tfor _, val := range vals {\n\t\tif val != zero {\n\t\t\treturn val\n\t\t}\n\t}\n\treturn zero\n}\n"
