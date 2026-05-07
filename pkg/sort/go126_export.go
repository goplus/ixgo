// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.26
// +build go1.26

package sort

import (
	q "sort"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("sort", func() *ixgo.Package {
		return &ixgo.Package{
			Name: "sort",
			Path: "sort",
			Deps: map[string]string{
				"internal/reflectlite": "reflectlite",
				"math/bits":            "bits",
				"slices":               "slices",
			},
			Interfaces: map[string]reflect.Type{
				"Interface": reflect.TypeOf((*q.Interface)(nil)).Elem(),
			},
			NamedTypes: map[string]reflect.Type{
				"Float64Slice": reflect.TypeOf((*q.Float64Slice)(nil)).Elem(),
				"IntSlice":     reflect.TypeOf((*q.IntSlice)(nil)).Elem(),
				"StringSlice":  reflect.TypeOf((*q.StringSlice)(nil)).Elem(),
			},
			AliasTypes: map[string]reflect.Type{},
			Vars:       map[string]interface{}{},
			Funcs: map[string]interface{}{
				"Find":              q.Find,
				"Float64s":          q.Float64s,
				"Float64sAreSorted": q.Float64sAreSorted,
				"Ints":              q.Ints,
				"IntsAreSorted":     q.IntsAreSorted,
				"IsSorted":          q.IsSorted,
				"Reverse":           q.Reverse,
				"Search":            q.Search,
				"SearchFloat64s":    q.SearchFloat64s,
				"SearchInts":        q.SearchInts,
				"SearchStrings":     q.SearchStrings,
				"Slice":             q.Slice,
				"SliceIsSorted":     q.SliceIsSorted,
				"SliceStable":       q.SliceStable,
				"Sort":              q.Sort,
				"Stable":            q.Stable,
				"Strings":           q.Strings,
				"StringsAreSorted":  q.StringsAreSorted,
			},
			TypedConsts:   map[string]interface{}{},
			UntypedConsts: map[string]ixgo.UntypedConst{},
		}
	})
}
