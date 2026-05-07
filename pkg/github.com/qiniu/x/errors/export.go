// export by github.com/goplus/ixgo/cmd/qexp

package errors

import (
	q "github.com/qiniu/x/errors"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/qiniu/x/errors", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "errors",
			Path: "github.com/qiniu/x/errors",
			Deps: map[string]string{
				"errors":  "errors",
				"fmt":     "fmt",
				"io":      "io",
				"reflect": "reflect",
				"runtime": "runtime",
				"strconv": "strconv",
				"strings": "strings",
			},
			Interfaces: map[string]reflect.Type{},
			NamedTypes: map[string]reflect.Type{
				"Frame":    reflect.TypeOf((*q.Frame)(nil)).Elem(),
				"List":     reflect.TypeOf((*q.List)(nil)).Elem(),
				"NotFound": reflect.TypeOf((*q.NotFound)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{
				"ErrorInfo": reflect.TypeOf((*q.ErrorInfo)(nil)).Elem(),
			},
			Vars: map[string]interface{}{},
			Funcs: map[string]interface{}{
				"As":         q.As,
				"CallDetail": q.CallDetail,
				"Detail":     q.Detail,
				"Err":        q.Err,
				"Info":       q.Info,
				"InfoEx":     q.InfoEx,
				"Is":         q.Is,
				"IsNotFound": q.IsNotFound,
				"New":        q.New,
				"NewFrame":   q.NewFrame,
				"NewWith":    q.NewWith,
				"Summary":    q.Summary,
				"Unwrap":     q.Unwrap,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
