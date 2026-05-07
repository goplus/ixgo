// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package flag

import (
	q "flag"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("flag", func() *ixgo.Package {
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
			Vars: map[string]interface{}{
				"CommandLine": &q.CommandLine,
				"ErrHelp":     &q.ErrHelp,
				"Usage":       &q.Usage,
			},
			Funcs: map[string]interface{}{
				"Arg":           q.Arg,
				"Args":          q.Args,
				"Bool":          q.Bool,
				"BoolFunc":      q.BoolFunc,
				"BoolVar":       q.BoolVar,
				"Duration":      q.Duration,
				"DurationVar":   q.DurationVar,
				"Float64":       q.Float64,
				"Float64Var":    q.Float64Var,
				"Func":          q.Func,
				"Int":           q.Int,
				"Int64":         q.Int64,
				"Int64Var":      q.Int64Var,
				"IntVar":        q.IntVar,
				"Lookup":        q.Lookup,
				"NArg":          q.NArg,
				"NFlag":         q.NFlag,
				"NewFlagSet":    q.NewFlagSet,
				"Parse":         q.Parse,
				"Parsed":        q.Parsed,
				"PrintDefaults": q.PrintDefaults,
				"Set":           q.Set,
				"String":        q.String,
				"StringVar":     q.StringVar,
				"TextVar":       q.TextVar,
				"Uint":          q.Uint,
				"Uint64":        q.Uint64,
				"Uint64Var":     q.Uint64Var,
				"UintVar":       q.UintVar,
				"UnquoteUsage":  q.UnquoteUsage,
				"Var":           q.Var,
				"Visit":         q.Visit,
				"VisitAll":      q.VisitAll,
			},
			TypedConsts: map[string]interface{}{
				"ContinueOnError": q.ContinueOnError,
				"ExitOnError":     q.ExitOnError,
				"PanicOnError":    q.PanicOnError,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
