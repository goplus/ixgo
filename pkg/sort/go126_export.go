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
	ixgo.RegisterPackage(&ixgo.Package{
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
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Find":              reflect.ValueOf(q.Find),
			"Float64s":          reflect.ValueOf(q.Float64s),
			"Float64sAreSorted": reflect.ValueOf(q.Float64sAreSorted),
			"Ints":              reflect.ValueOf(q.Ints),
			"IntsAreSorted":     reflect.ValueOf(q.IntsAreSorted),
			"IsSorted":          reflect.ValueOf(q.IsSorted),
			"Reverse":           reflect.ValueOf(q.Reverse),
			"Search":            reflect.ValueOf(q.Search),
			"SearchFloat64s":    reflect.ValueOf(q.SearchFloat64s),
			"SearchInts":        reflect.ValueOf(q.SearchInts),
			"SearchStrings":     reflect.ValueOf(q.SearchStrings),
			"Slice":             reflect.ValueOf(q.Slice),
			"SliceIsSorted":     reflect.ValueOf(q.SliceIsSorted),
			"SliceStable":       reflect.ValueOf(q.SliceStable),
			"Sort":              reflect.ValueOf(q.Sort),
			"Stable":            reflect.ValueOf(q.Stable),
			"Strings":           reflect.ValueOf(q.Strings),
			"StringsAreSorted":  reflect.ValueOf(q.StringsAreSorted),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
