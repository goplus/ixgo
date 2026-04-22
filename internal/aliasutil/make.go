package aliasutil

import (
	"go/types"

	"github.com/goplus/ixgo/alias"
	"github.com/goplus/ixgo/internal/typesutil"
)

func Make(typ types.Type, atyp alias.Type, lookup func(typ types.Type, name string) types.Type) types.Type {
	if atyp == nil {
		return typ
	}
	switch typ := typ.(type) {
	case *types.Basic:
		if t := lookup(typ, atyp.(*alias.Alias).Typ); t != nil {
			return t
		}
	case *types.Named:
		if t := lookup(typ, atyp.(*alias.Alias).Typ); t != nil {
			return t
		}
	case *types.Array:
		atyp := atyp.(*alias.Array)
		return types.NewArray(Make(typ.Elem(), atyp.Elem, lookup), typ.Len())
	case *types.Chan:
		atyp := atyp.(*alias.Chan)
		return types.NewChan(typ.Dir(), Make(typ.Elem(), atyp.Elem, lookup))
	case *types.Map:
		atyp := atyp.(*alias.Map)
		return types.NewMap(Make(typ.Key(), atyp.Key, lookup), Make(typ.Elem(), atyp.Elem, lookup))
	case *types.Slice:
		atyp := atyp.(*alias.Slice)
		return types.NewSlice(Make(typ.Elem(), atyp.Elem, lookup))
	case *types.Signature:
		atyp := atyp.(*alias.Func)
		return types.NewSignature(typ.Recv(), Tuple(typ.Params(), atyp.Params, lookup), Tuple(typ.Results(), atyp.Results, lookup), typ.Variadic())
	case *types.Struct:
		n := typ.NumFields()
		if n == 0 {
			break
		}
		atyp := atyp.(*alias.Struct)
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if ntyp := Make(f.Type(), atyp.Fields[i], lookup); ntyp != nil {
				typesutil.SetVarType(f, ntyp)
			}
		}
	case *types.Interface:
		n := typ.NumMethods()
		if n == 0 {
			break
		}
		atyp := atyp.(*alias.Interface)
		for i := 0; i < n; i++ {
			m := typ.Method(i)
			if afn, ok := atyp.Methods[m.Name()]; ok {
				if ntyp := Make(m.Type(), afn, lookup); ntyp != nil {
					typesutil.SetFuncType(m, ntyp)
				}
			}
		}
		return typ
	}
	return typ
}

func Tuple(tuple *types.Tuple, typs []alias.Type, lookup func(typ types.Type, name string) types.Type) *types.Tuple {
	if tuple == nil || typs == nil {
		return tuple
	}
	n := tuple.Len()
	vars := make([]*types.Var, n)
	for i := 0; i < n; i++ {
		v := tuple.At(i)
		if typs[i] == nil {
			vars[i] = v
		} else {
			vars[i] = types.NewVar(v.Pos(), v.Pkg(), v.Name(), Make(v.Type(), typs[i], lookup))
		}
	}
	return types.NewTuple(vars...)
}
