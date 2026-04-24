// export by github.com/goplus/ixgo/cmd/qexp

package pkg

import (
	q "github.com/goplus/ixgo/testdata/alias/pkg"

	"github.com/goplus/ixgo/alias"
	"go/constant"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "pkg",
		Path: "github.com/goplus/ixgo/testdata/alias/pkg",
		Deps: map[string]string{
			"github.com/goplus/ixgo/testdata/alias/msg": "msg",
		},
		Interfaces: map[string]reflect.Type{
			"Addable": reflect.TypeOf((*q.Addable)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Point": reflect.TypeOf((*q.Point)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Int": reflect.TypeOf((*q.Int)(nil)).Elem(),
			"M":   reflect.TypeOf((*q.M)(nil)).Elem(),
			"Msg": reflect.TypeOf((*q.Msg)(nil)).Elem(),
			"S":   reflect.TypeOf((*q.S)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"V1": reflect.ValueOf(&q.V1),
			"V2": reflect.ValueOf(&q.V2),
			"V3": reflect.ValueOf(&q.V3),
			"V4": reflect.ValueOf(&q.V4),
		},
		Funcs: map[string]reflect.Value{
			"Demo1": reflect.ValueOf(q.Demo1),
			"Demo2": reflect.ValueOf(q.Demo2),
			"Demo3": reflect.ValueOf(q.Demo3),
			"Demo4": reflect.ValueOf(q.Demo4),
			"Demo5": reflect.ValueOf(q.Demo5),
			"Demo6": reflect.ValueOf(q.Demo6),
			"Demo7": reflect.ValueOf(q.Demo7),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"A1": {Typ: reflect.TypeOf(q.A1), Value: constant.MakeInt64(int64(q.A1))},
			"A2": {Typ: reflect.TypeOf(q.A2), Value: constant.MakeInt64(int64(q.A2))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
		Alias: map[string]alias.Type{
			"A1": &alias.Alias{Typ: "Int"},
			"Addable": &alias.Named{
				Underlying: &alias.Interface{
					Methods: map[string]*alias.Func{
						"Add": &alias.Func{
							Params:  []alias.Type{&alias.Alias{Typ: "Int"}, nil},
							Results: nil,
						},
					},
				},
			},
			"Demo1": &alias.Func{
				Params:  []alias.Type{&alias.Alias{Typ: "Int"}, nil},
				Results: []alias.Type{&alias.Alias{Typ: "Int"}},
			},
			"Demo2": &alias.Func{
				Params:  []alias.Type{&alias.Alias{Typ: "myInt"}, nil},
				Results: []alias.Type{&alias.Alias{Typ: "Int"}},
			},
			"Demo3": &alias.Func{
				Params:  []alias.Type{&alias.Builtin{Typ: "any"}},
				Results: nil,
			},
			"Demo4": &alias.Func{
				Params:  []alias.Type{nil, &alias.Alias{Typ: "M"}},
				Results: nil,
			},
			"Demo5": &alias.Func{
				Params:  []alias.Type{nil, &alias.Map{Key: &alias.Alias{Typ: "Int"}, Elem: nil}},
				Results: nil,
			},
			"Demo6": &alias.Func{
				Params:  []alias.Type{nil, &alias.Alias{Typ: "S"}, &alias.Alias{Typ: "Msg"}},
				Results: nil,
			},
			"Demo7": &alias.Func{
				Params:  []alias.Type{nil, &alias.Slice{Elem: &alias.Alias{Typ: "Int"}}},
				Results: nil,
			},
			"Point": &alias.Named{
				Underlying: &alias.Struct{Fields: []alias.Type{&alias.Alias{Typ: "Int"}, nil}},
				Methods: map[string]*alias.Func{
					"Demo1": &alias.Func{
						Params:  []alias.Type{&alias.Alias{Typ: "Int"}, nil},
						Results: []alias.Type{&alias.Alias{Typ: "Int"}},
					},
					"Demo2": &alias.Func{
						Params:  []alias.Type{&alias.Alias{Typ: "myInt"}, nil},
						Results: []alias.Type{&alias.Alias{Typ: "Int"}},
					},
					"Demo3": &alias.Func{
						Params:  []alias.Type{&alias.Builtin{Typ: "any"}},
						Results: nil,
					},
					"Demo4": &alias.Func{
						Params:  []alias.Type{nil, &alias.Alias{Typ: "M"}},
						Results: nil,
					},
					"Demo5": &alias.Func{
						Params:  []alias.Type{nil, &alias.Map{Key: &alias.Alias{Typ: "Int"}, Elem: &alias.Alias{Typ: "github.com/goplus/ixgo/testdata/alias/msg.Message"}}},
						Results: []alias.Type{nil, &alias.Alias{Typ: "github.com/goplus/ixgo/testdata/alias/msg.Message"}},
					},
				},
			},
			"V1": &alias.Alias{Typ: "Int"},
			"V2": &alias.Alias{Typ: "M"},
			"V3": &alias.Slice{Elem: &alias.Alias{Typ: "Int"}},
			"V4": &alias.Struct{Fields: []alias.Type{&alias.Alias{Typ: "Int"}, &alias.Array{Elem: &alias.Alias{Typ: "Int"}}, &alias.Chan{Elem: &alias.Alias{Typ: "Int"}}, nil, &alias.Func{
				Params:  []alias.Type{&alias.Alias{Typ: "Int"}},
				Results: nil,
			}, &alias.Pointer{Elem: &alias.Struct{Fields: []alias.Type{&alias.Alias{Typ: "Int"}}}}}},
		},
	})
}
