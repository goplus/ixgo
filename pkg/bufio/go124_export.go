// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package bufio

import (
	q "bufio"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars: map[string]reflect.Value{
			"ErrAdvanceTooFar":     reflect.ValueOf(&q.ErrAdvanceTooFar),
			"ErrBadReadCount":      reflect.ValueOf(&q.ErrBadReadCount),
			"ErrBufferFull":        reflect.ValueOf(&q.ErrBufferFull),
			"ErrFinalToken":        reflect.ValueOf(&q.ErrFinalToken),
			"ErrInvalidUnreadByte": reflect.ValueOf(&q.ErrInvalidUnreadByte),
			"ErrInvalidUnreadRune": reflect.ValueOf(&q.ErrInvalidUnreadRune),
			"ErrNegativeAdvance":   reflect.ValueOf(&q.ErrNegativeAdvance),
			"ErrNegativeCount":     reflect.ValueOf(&q.ErrNegativeCount),
			"ErrTooLong":           reflect.ValueOf(&q.ErrTooLong),
		},
		Funcs: map[string]reflect.Value{
			"NewReadWriter": reflect.ValueOf(q.NewReadWriter),
			"NewReader":     reflect.ValueOf(q.NewReader),
			"NewReaderSize": reflect.ValueOf(q.NewReaderSize),
			"NewScanner":    reflect.ValueOf(q.NewScanner),
			"NewWriter":     reflect.ValueOf(q.NewWriter),
			"NewWriterSize": reflect.ValueOf(q.NewWriterSize),
			"ScanBytes":     reflect.ValueOf(q.ScanBytes),
			"ScanLines":     reflect.ValueOf(q.ScanLines),
			"ScanRunes":     reflect.ValueOf(q.ScanRunes),
			"ScanWords":     reflect.ValueOf(q.ScanWords),
		},
		TypedConsts: map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"MaxScanTokenSize": {"untyped int", constant.MakeInt64(int64(q.MaxScanTokenSize))},
		},
	})
}
