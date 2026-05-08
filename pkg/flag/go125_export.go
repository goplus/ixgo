// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package flag

import (
	q "flag"

	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackageLazy("flag", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "flag",
			Path: "flag",
			Deps: map[string]string{
				"encoding": "encoding",
				"errors":   "errors",
				"fmt":      "fmt",
				"io":       "io",
				"os":       "os",
				"reflect":  "reflect",
				"runtime":  "runtime",
				"slices":   "slices",
				"strconv":  "strconv",
				"strings":  "strings",
				"time":     "time",
			},
			Interfaces: map[string]reflect.Type{
				"Getter": reflect.TypeOf((*q.Getter)(nil)).Elem(),
				"Value":  reflect.TypeOf((*q.Value)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"ErrorHandling": reflect.TypeOf((*q.ErrorHandling)(nil)).Elem(),
				"Flag":          reflect.TypeOf((*q.Flag)(nil)).Elem(),
				"FlagSet":       reflect.TypeOf((*q.FlagSet)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars: map[string]reflect.Value{
				"CommandLine": reflect.ValueOf(&q.CommandLine),
				"ErrHelp":     reflect.ValueOf(&q.ErrHelp),
				"Usage":       reflect.ValueOf(&q.Usage),
			},
			Funcs: map[string]reflect.Value{
				"Arg":           reflect.ValueOf(q.Arg),
				"Args":          reflect.ValueOf(q.Args),
				"Bool":          reflect.ValueOf(q.Bool),
				"BoolFunc":      reflect.ValueOf(q.BoolFunc),
				"BoolVar":       reflect.ValueOf(q.BoolVar),
				"Duration":      reflect.ValueOf(q.Duration),
				"DurationVar":   reflect.ValueOf(q.DurationVar),
				"Float64":       reflect.ValueOf(q.Float64),
				"Float64Var":    reflect.ValueOf(q.Float64Var),
				"Func":          reflect.ValueOf(q.Func),
				"Int":           reflect.ValueOf(q.Int),
				"Int64":         reflect.ValueOf(q.Int64),
				"Int64Var":      reflect.ValueOf(q.Int64Var),
				"IntVar":        reflect.ValueOf(q.IntVar),
				"Lookup":        reflect.ValueOf(q.Lookup),
				"NArg":          reflect.ValueOf(q.NArg),
				"NFlag":         reflect.ValueOf(q.NFlag),
				"NewFlagSet":    reflect.ValueOf(q.NewFlagSet),
				"Parse":         reflect.ValueOf(q.Parse),
				"Parsed":        reflect.ValueOf(q.Parsed),
				"PrintDefaults": reflect.ValueOf(q.PrintDefaults),
				"Set":           reflect.ValueOf(q.Set),
				"String":        reflect.ValueOf(q.String),
				"StringVar":     reflect.ValueOf(q.StringVar),
				"TextVar":       reflect.ValueOf(q.TextVar),
				"Uint":          reflect.ValueOf(q.Uint),
				"Uint64":        reflect.ValueOf(q.Uint64),
				"Uint64Var":     reflect.ValueOf(q.Uint64Var),
				"UintVar":       reflect.ValueOf(q.UintVar),
				"UnquoteUsage":  reflect.ValueOf(q.UnquoteUsage),
				"Var":           reflect.ValueOf(q.Var),
				"Visit":         reflect.ValueOf(q.Visit),
				"VisitAll":      reflect.ValueOf(q.VisitAll),
			},
			TypedConsts: map[string]ixgo.TypedConst{
				"ContinueOnError": {Typ: reflect.TypeOf(q.ContinueOnError), Value: constant.MakeInt64(int64(q.ContinueOnError))},
				"ExitOnError":     {Typ: reflect.TypeOf(q.ExitOnError), Value: constant.MakeInt64(int64(q.ExitOnError))},
				"PanicOnError":    {Typ: reflect.TypeOf(q.PanicOnError), Value: constant.MakeInt64(int64(q.PanicOnError))},
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
