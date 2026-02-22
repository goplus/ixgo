// export by github.com/goplus/ixgo/cmd/qexp

package dql

import (
	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "dql",
		Path: "github.com/goplus/xgo/dql",
		Deps: map[string]string{
			"errors":  "errors",
			"strconv": "strconv",
			"strings": "strings",
		},
		Source: source,
	})
}

var source = "package dql\n\nimport (\n\t\"errors\"\n\t\"strconv\"\n\t\"strings\"\n)\n\nconst (\n\tXGoPackage = true\n)\n\nvar (\n\tErrNotFound\t\t= errors.New(\"entity not found\")\n\tErrMultiEntities\t= errors.New(\"too many entities found\")\n)\n\nfunc NopIter[T any](yield func(T) bool)\t{}\n\nfunc First[T any, Seq ~func(func(T) bool)](seq Seq) (ret T, err error) {\n\terr = ErrNotFound\n\tseq(func(item T) bool {\n\t\tret, err = item, nil\n\t\treturn false\n\t})\n\treturn\n}\n\nfunc Single[T any, Seq ~func(func(T) bool)](seq Seq) (ret T, err error) {\n\terr = ErrNotFound\n\tfirst := true\n\tseq(func(item T) bool {\n\t\tif first {\n\t\t\tret, err = item, nil\n\t\t\tfirst = false\n\t\t\treturn true\n\t\t}\n\t\terr = ErrMultiEntities\n\t\treturn false\n\t})\n\treturn\n}\n\nfunc Collect[T any, Seq ~func(func(T) bool)](seq Seq) []T {\n\tret := make([]T, 0, 8)\n\tseq(func(item T) bool {\n\t\tret = append(ret, item)\n\t\treturn true\n\t})\n\treturn ret\n}\n\nfunc Int(text string) (int, error) {\n\treturn strconv.Atoi(strings.ReplaceAll(strings.TrimSpace(text), \",\", \"\"))\n}\n"
