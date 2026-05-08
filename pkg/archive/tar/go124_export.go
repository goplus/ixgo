// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package tar

import (
	q "archive/tar"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("archive/tar", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "tar",
			Path: "archive/tar",
			Deps: map[string]string{
				"bytes":            "bytes",
				"errors":           "errors",
				"fmt":              "fmt",
				"internal/godebug": "godebug",
				"io":               "io",
				"io/fs":            "fs",
				"maps":             "maps",
				"math":             "math",
				"os/user":          "user",
				"path":             "path",
				"path/filepath":    "filepath",
				"reflect":          "reflect",
				"runtime":          "runtime",
				"slices":           "slices",
				"strconv":          "strconv",
				"strings":          "strings",
				"sync":             "sync",
				"syscall":          "syscall",
				"time":             "time",
			},
			Interfaces: map[string]reflect.Type{
				"FileInfoNames": reflect.TypeOf((*q.FileInfoNames)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Format": reflect.TypeOf((*q.Format)(nil)).Elem(),
				"Header": reflect.TypeOf((*q.Header)(nil)).Elem(),
				"Reader": reflect.TypeOf((*q.Reader)(nil)).Elem(),
				"Writer": reflect.TypeOf((*q.Writer)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]reflect.Value{
				"ErrFieldTooLong":    reflect.ValueOf(&q.ErrFieldTooLong),
				"ErrHeader":          reflect.ValueOf(&q.ErrHeader),
				"ErrInsecurePath":    reflect.ValueOf(&q.ErrInsecurePath),
				"ErrWriteAfterClose": reflect.ValueOf(&q.ErrWriteAfterClose),
				"ErrWriteTooLong":    reflect.ValueOf(&q.ErrWriteTooLong),
			},
			Funcs: map[string]reflect.Value{
				"FileInfoHeader": reflect.ValueOf(q.FileInfoHeader),
				"NewReader":      reflect.ValueOf(q.NewReader),
				"NewWriter":      reflect.ValueOf(q.NewWriter),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"FormatGNU":     {Typ: reflect.TypeOf(q.FormatGNU), Value: constant.MakeInt64(int64(q.FormatGNU))},
				"FormatPAX":     {Typ: reflect.TypeOf(q.FormatPAX), Value: constant.MakeInt64(int64(q.FormatPAX))},
				"FormatUSTAR":   {Typ: reflect.TypeOf(q.FormatUSTAR), Value: constant.MakeInt64(int64(q.FormatUSTAR))},
				"FormatUnknown": {Typ: reflect.TypeOf(q.FormatUnknown), Value: constant.MakeInt64(int64(q.FormatUnknown))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{
				"TypeBlock":         {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeBlock))},
				"TypeChar":          {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeChar))},
				"TypeCont":          {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeCont))},
				"TypeDir":           {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeDir))},
				"TypeFifo":          {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeFifo))},
				"TypeGNULongLink":   {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeGNULongLink))},
				"TypeGNULongName":   {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeGNULongName))},
				"TypeGNUSparse":     {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeGNUSparse))},
				"TypeLink":          {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeLink))},
				"TypeReg":           {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeReg))},
				"TypeRegA":          {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeRegA))},
				"TypeSymlink":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeSymlink))},
				"TypeXGlobalHeader": {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeXGlobalHeader))},
				"TypeXHeader":       {Typ: "untyped rune", Value: constant.MakeInt64(int64(q.TypeXHeader))},
			},
		}
	})
}
