// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.24 && !go1.25
// +build go1.24,!go1.25

package slog

import (
	q "log/slog"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "slog",
		Path: "log/slog",
		Deps: map[string]string{
			"bytes":                    "bytes",
			"context":                  "context",
			"encoding":                 "encoding",
			"encoding/json":            "json",
			"errors":                   "errors",
			"fmt":                      "fmt",
			"io":                       "io",
			"log":                      "log",
			"log/internal":             "internal",
			"log/slog/internal":        "internal",
			"log/slog/internal/buffer": "buffer",
			"math":                     "math",
			"reflect":                  "reflect",
			"runtime":                  "runtime",
			"slices":                   "slices",
			"strconv":                  "strconv",
			"strings":                  "strings",
			"sync":                     "sync",
			"sync/atomic":              "atomic",
			"time":                     "time",
			"unicode":                  "unicode",
			"unicode/utf8":             "utf8",
			"unsafe":                   "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"Handler":   reflect.TypeOf((*q.Handler)(nil)).Elem(),
			"Leveler":   reflect.TypeOf((*q.Leveler)(nil)).Elem(),
			"LogValuer": reflect.TypeOf((*q.LogValuer)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Attr":           reflect.TypeOf((*q.Attr)(nil)).Elem(),
			"HandlerOptions": reflect.TypeOf((*q.HandlerOptions)(nil)).Elem(),
			"JSONHandler":    reflect.TypeOf((*q.JSONHandler)(nil)).Elem(),
			"Kind":           reflect.TypeOf((*q.Kind)(nil)).Elem(),
			"Level":          reflect.TypeOf((*q.Level)(nil)).Elem(),
			"LevelVar":       reflect.TypeOf((*q.LevelVar)(nil)).Elem(),
			"Logger":         reflect.TypeOf((*q.Logger)(nil)).Elem(),
			"Record":         reflect.TypeOf((*q.Record)(nil)).Elem(),
			"Source":         reflect.TypeOf((*q.Source)(nil)).Elem(),
			"TextHandler":    reflect.TypeOf((*q.TextHandler)(nil)).Elem(),
			"Value":          reflect.TypeOf((*q.Value)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"DiscardHandler": reflect.ValueOf(&q.DiscardHandler),
		},
		Funcs: map[string]reflect.Value{
			"Any":               reflect.ValueOf(q.Any),
			"AnyValue":          reflect.ValueOf(q.AnyValue),
			"Bool":              reflect.ValueOf(q.Bool),
			"BoolValue":         reflect.ValueOf(q.BoolValue),
			"Debug":             reflect.ValueOf(q.Debug),
			"DebugContext":      reflect.ValueOf(q.DebugContext),
			"Default":           reflect.ValueOf(q.Default),
			"Duration":          reflect.ValueOf(q.Duration),
			"DurationValue":     reflect.ValueOf(q.DurationValue),
			"Error":             reflect.ValueOf(q.Error),
			"ErrorContext":      reflect.ValueOf(q.ErrorContext),
			"Float64":           reflect.ValueOf(q.Float64),
			"Float64Value":      reflect.ValueOf(q.Float64Value),
			"Group":             reflect.ValueOf(q.Group),
			"GroupValue":        reflect.ValueOf(q.GroupValue),
			"Info":              reflect.ValueOf(q.Info),
			"InfoContext":       reflect.ValueOf(q.InfoContext),
			"Int":               reflect.ValueOf(q.Int),
			"Int64":             reflect.ValueOf(q.Int64),
			"Int64Value":        reflect.ValueOf(q.Int64Value),
			"IntValue":          reflect.ValueOf(q.IntValue),
			"Log":               reflect.ValueOf(q.Log),
			"LogAttrs":          reflect.ValueOf(q.LogAttrs),
			"New":               reflect.ValueOf(q.New),
			"NewJSONHandler":    reflect.ValueOf(q.NewJSONHandler),
			"NewLogLogger":      reflect.ValueOf(q.NewLogLogger),
			"NewRecord":         reflect.ValueOf(q.NewRecord),
			"NewTextHandler":    reflect.ValueOf(q.NewTextHandler),
			"SetDefault":        reflect.ValueOf(q.SetDefault),
			"SetLogLoggerLevel": reflect.ValueOf(q.SetLogLoggerLevel),
			"String":            reflect.ValueOf(q.String),
			"StringValue":       reflect.ValueOf(q.StringValue),
			"Time":              reflect.ValueOf(q.Time),
			"TimeValue":         reflect.ValueOf(q.TimeValue),
			"Uint64":            reflect.ValueOf(q.Uint64),
			"Uint64Value":       reflect.ValueOf(q.Uint64Value),
			"Warn":              reflect.ValueOf(q.Warn),
			"WarnContext":       reflect.ValueOf(q.WarnContext),
			"With":              reflect.ValueOf(q.With),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"KindAny":       {reflect.TypeOf(q.KindAny), constant.MakeInt64(int64(q.KindAny))},
			"KindBool":      {reflect.TypeOf(q.KindBool), constant.MakeInt64(int64(q.KindBool))},
			"KindDuration":  {reflect.TypeOf(q.KindDuration), constant.MakeInt64(int64(q.KindDuration))},
			"KindFloat64":   {reflect.TypeOf(q.KindFloat64), constant.MakeInt64(int64(q.KindFloat64))},
			"KindGroup":     {reflect.TypeOf(q.KindGroup), constant.MakeInt64(int64(q.KindGroup))},
			"KindInt64":     {reflect.TypeOf(q.KindInt64), constant.MakeInt64(int64(q.KindInt64))},
			"KindLogValuer": {reflect.TypeOf(q.KindLogValuer), constant.MakeInt64(int64(q.KindLogValuer))},
			"KindString":    {reflect.TypeOf(q.KindString), constant.MakeInt64(int64(q.KindString))},
			"KindTime":      {reflect.TypeOf(q.KindTime), constant.MakeInt64(int64(q.KindTime))},
			"KindUint64":    {reflect.TypeOf(q.KindUint64), constant.MakeInt64(int64(q.KindUint64))},
			"LevelDebug":    {reflect.TypeOf(q.LevelDebug), constant.MakeInt64(int64(q.LevelDebug))},
			"LevelError":    {reflect.TypeOf(q.LevelError), constant.MakeInt64(int64(q.LevelError))},
			"LevelInfo":     {reflect.TypeOf(q.LevelInfo), constant.MakeInt64(int64(q.LevelInfo))},
			"LevelWarn":     {reflect.TypeOf(q.LevelWarn), constant.MakeInt64(int64(q.LevelWarn))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{
			"LevelKey":   {"untyped string", constant.MakeString(string(q.LevelKey))},
			"MessageKey": {"untyped string", constant.MakeString(string(q.MessageKey))},
			"SourceKey":  {"untyped string", constant.MakeString(string(q.SourceKey))},
			"TimeKey":    {"untyped string", constant.MakeString(string(q.TimeKey))},
		},
	})
}
