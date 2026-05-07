// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package bufio

import (
	q "bufio"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("bufio", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "bufio",
			Path: "bufio",
			Deps: map[string]string{
				"bytes":        "bytes",
				"errors":       "errors",
				"io":           "io",
				"strings":      "strings",
				"unicode/utf8": "utf8",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"ReadWriter": reflect.TypeOf((*q.ReadWriter)(nil)).Elem(),
				"Reader":     reflect.TypeOf((*q.Reader)(nil)).Elem(),
				"Scanner":    reflect.TypeOf((*q.Scanner)(nil)).Elem(),
				"SplitFunc":  reflect.TypeOf((*q.SplitFunc)(nil)).Elem(),
				"Writer":     reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]interface{}{
				"ErrAdvanceTooFar":     &q.ErrAdvanceTooFar,
				"ErrBadReadCount":      &q.ErrBadReadCount,
				"ErrBufferFull":        &q.ErrBufferFull,
				"ErrFinalToken":        &q.ErrFinalToken,
				"ErrInvalidUnreadByte": &q.ErrInvalidUnreadByte,
				"ErrInvalidUnreadRune": &q.ErrInvalidUnreadRune,
				"ErrNegativeAdvance":   &q.ErrNegativeAdvance,
				"ErrNegativeCount":     &q.ErrNegativeCount,
				"ErrTooLong":           &q.ErrTooLong,
			},
			Funcs: map[string]interface{}{
				"NewReadWriter": q.NewReadWriter,
				"NewReader":     q.NewReader,
				"NewReaderSize": q.NewReaderSize,
				"NewScanner":    q.NewScanner,
				"NewWriter":     q.NewWriter,
				"NewWriterSize": q.NewWriterSize,
				"ScanBytes":     q.ScanBytes,
				"ScanLines":     q.ScanLines,
				"ScanRunes":     q.ScanRunes,
				"ScanWords":     q.ScanWords,
			},
			TypedConsts: map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"MaxScanTokenSize": {Typ: "untyped int", Value: constant.MakeInt64(int64(q.MaxScanTokenSize))},
			},
		}
	})
}
