// export by github.com/goplus/ixgo/cmd/qexp

package pkg

import (
	q "github.com/goplus/ixgo/testdata/alias/pkg"

	"github.com/goplus/ixgo/alias"
	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage("github.com/goplus/ixgo/testdata/alias/pkg", func() *ixgo.Package {
		var (
			alias_Int         = &alias.Alias{Typ: "Int"}
			alias_M           = &alias.Alias{Typ: "M"}
			alias_Msg         = &alias.Alias{Typ: "Msg"}
			alias_S           = &alias.Alias{Typ: "S"}
			alias_any         = &alias.Builtin{Typ: "any"}
			alias_msg_Message = &alias.Alias{Typ: "github.com/goplus/ixgo/testdata/alias/msg.Message"}
			alias_myInt       = &alias.Alias{Typ: "myInt"}
		)
		return &ixgo.Package{
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
			Vars: map[string]interface{}{
				"V1": &q.V1,
				"V2": &q.V2,
				"V3": &q.V3,
				"V4": &q.V4,
			},
			Funcs: map[string]interface{}{
				"Demo1": q.Demo1,
				"Demo2": q.Demo2,
				"Demo3": q.Demo3,
				"Demo4": q.Demo4,
				"Demo5": q.Demo5,
				"Demo6": q.Demo6,
				"Demo7": q.Demo7,
			},
			TypedConsts: map[string]interface{}{
				"A1": q.A1,
				"A2": q.A2,
			},
			UntypedConsts: map[string]ixgo.UntypedConst{},
			Alias: map[string]alias.Type{
				"A1": alias_Int,
				"Addable": &alias.Named{
					Underlying: &alias.Interface{
						Methods: map[string]*alias.Func{
							"Add": &alias.Func{
								Params:  []alias.Type{alias_Int, nil},
								Results: nil,
							},
						},
					},
				},
				"Demo1": &alias.Func{
					Params:  []alias.Type{alias_Int, nil},
					Results: []alias.Type{alias_Int},
				},
				"Demo2": &alias.Func{
					Params:  []alias.Type{alias_myInt, nil},
					Results: []alias.Type{alias_Int},
				},
				"Demo3": &alias.Func{
					Params:  []alias.Type{alias_any},
					Results: nil,
				},
				"Demo4": &alias.Func{
					Params:  []alias.Type{nil, alias_M},
					Results: nil,
				},
				"Demo5": &alias.Func{
					Params:  []alias.Type{nil, &alias.Map{Key: alias_Int, Elem: nil}},
					Results: nil,
				},
				"Demo6": &alias.Func{
					Params:  []alias.Type{nil, alias_S, alias_Msg},
					Results: nil,
				},
				"Demo7": &alias.Func{
					Params:  []alias.Type{nil, &alias.Slice{Elem: alias_Int}},
					Results: nil,
				},
				"Point": &alias.Named{
					Underlying: &alias.Struct{Fields: []alias.Type{alias_Int, nil}},
					Methods: map[string]*alias.Func{
						"Demo1": &alias.Func{
							Params:  []alias.Type{alias_Int, nil},
							Results: []alias.Type{alias_Int},
						},
						"Demo2": &alias.Func{
							Params:  []alias.Type{alias_myInt, nil},
							Results: []alias.Type{alias_Int},
						},
						"Demo3": &alias.Func{
							Params:  []alias.Type{alias_any},
							Results: nil,
						},
						"Demo4": &alias.Func{
							Params:  []alias.Type{nil, alias_M},
							Results: nil,
						},
						"Demo5": &alias.Func{
							Params:  []alias.Type{nil, &alias.Map{Key: alias_Int, Elem: alias_msg_Message}},
							Results: []alias.Type{nil, alias_msg_Message},
						},
					},
				},
				"V1": alias_Int,
				"V2": alias_M,
				"V3": &alias.Slice{Elem: alias_Int},
				"V4": &alias.Struct{Fields: []alias.Type{alias_Int, &alias.Array{Elem: alias_Int}, &alias.Chan{Elem: alias_Int}, nil, &alias.Func{
					Params:  []alias.Type{alias_Int},
					Results: nil,
				}, &alias.Pointer{Elem: &alias.Struct{Fields: []alias.Type{alias_Int}}}}},
			},
		}
	})
}
