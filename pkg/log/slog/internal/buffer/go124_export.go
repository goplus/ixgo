// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package buffer

import (
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "buffer",
		Path: "log/slog/internal/buffer",
		Deps: map[string]string{
			"sync": "sync",
		},
		Source: source,
	})
}

var source = "package buffer\n\nimport \"sync\"\n\ntype Buffer []byte\n\nvar bufPool = sync.Pool{\n\tNew: func() any {\n\t\tb := make([]byte, 0, 1024)\n\t\treturn (*Buffer)(&b)\n\t},\n}\n\nfunc New() *Buffer {\n\treturn bufPool.Get().(*Buffer)\n}\n\nfunc (b *Buffer) Free() {\n\n\tconst maxBufferSize = 16 << 10\n\tif cap(*b) <= maxBufferSize {\n\t\t*b = (*b)[:0]\n\t\tbufPool.Put(b)\n\t}\n}\n\nfunc (b *Buffer) Reset() {\n\tb.SetLen(0)\n}\n\nfunc (b *Buffer) Write(p []byte) (int, error) {\n\t*b = append(*b, p...)\n\treturn len(p), nil\n}\n\nfunc (b *Buffer) WriteString(s string) (int, error) {\n\t*b = append(*b, s...)\n\treturn len(s), nil\n}\n\nfunc (b *Buffer) WriteByte(c byte) error {\n\t*b = append(*b, c)\n\treturn nil\n}\n\nfunc (b *Buffer) String() string {\n\treturn string(*b)\n}\n\nfunc (b *Buffer) Len() int {\n\treturn len(*b)\n}\n\nfunc (b *Buffer) SetLen(n int) {\n\t*b = (*b)[:n]\n}\n"
