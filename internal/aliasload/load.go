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
	if atyp, ok := atyp.(*alias.Alias); ok {
		if obj := types.Universe.Lookup(atyp.Typ); obj != nil {
			return obj.Type()
		}
		if t := lookup(typ, atyp.Typ); t != nil {
			return t
		}
	}
	switch typ := typ.(type) {
	case *types.Basic:
		// skip
	case *types.Named:
		// skip
	case *types.Alias:
		// skip
	case *types.Array:
		if atyp, ok := atyp.(*alias.Array); ok {
			return types.NewArray(Make(typ.Elem(), atyp.Elem, lookup), typ.Len())
		}
	case *types.Chan:
		if atyp, ok := atyp.(*alias.Chan); ok {
			return types.NewChan(typ.Dir(), Make(typ.Elem(), atyp.Elem, lookup))
		}
	case *types.Map:
		if atyp, ok := atyp.(*alias.Map); ok {
			return types.NewMap(Make(typ.Key(), atyp.Key, lookup), Make(typ.Elem(), atyp.Elem, lookup))
		}
	case *types.Slice:
		if atyp, ok := atyp.(*alias.Slice); ok {
			return types.NewSlice(Make(typ.Elem(), atyp.Elem, lookup))
		}
	case *types.Pointer:
		if atyp, ok := atyp.(*alias.Pointer); ok {
			return types.NewPointer(Make(typ.Elem(), atyp.Elem, lookup))
		}
	case *types.Signature:
		if atyp, ok := atyp.(*alias.Func); ok {
			return types.NewSignature(typ.Recv(), Tuple(typ.Params(), atyp.Params, lookup), Tuple(typ.Results(), atyp.Results, lookup), typ.Variadic())
		}
	case *types.Struct:
		n := typ.NumFields()
		if n == 0 {
			break
		}
		if atyp, ok := atyp.(*alias.Struct); ok && len(atyp.Fields) == n {
			for i := 0; i < n; i++ {
				f := typ.Field(i)
				if ntyp := Make(f.Type(), atyp.Fields[i], lookup); ntyp != nil {
					typesutil.SetVarType(f, ntyp)
				}
			}
		}
	case *types.Interface:
		n := typ.NumMethods()
		if n == 0 {
			break
		}
		if atyp, ok := atyp.(*alias.Interface); ok {
			for i := 0; i < n; i++ {
				m := typ.Method(i)
				if afn, ok := atyp.Methods[m.Name()]; ok {
					if ntyp := Make(m.Type(), afn, lookup); ntyp != nil {
						typesutil.SetFuncType(m, ntyp)
					}
				}
			}
		}
	}
	return typ
}

func Tuple(tuple *types.Tuple, atyps []alias.Type, lookup func(typ types.Type, name string) types.Type) *types.Tuple {
	if tuple == nil || atyps == nil {
		return tuple
	}
	n := tuple.Len()
	if n != len(atyps) {
		return tuple
	}
	vars := make([]*types.Var, n)
	for i := 0; i < n; i++ {
		v := tuple.At(i)
		if atyps[i] == nil {
			vars[i] = v
		} else {
			vars[i] = types.NewVar(v.Pos(), v.Pkg(), v.Name(), Make(v.Type(), atyps[i], lookup))
		}
	}
	return types.NewTuple(vars...)
}

func Load(p *types.Package, atyps map[string]alias.Type, lookup func(typ types.Type, name string) types.Type) {
	scope := p.Scope()
	for name, atyp := range atyps {
		obj := scope.Lookup(name)
		if obj == nil {
			continue
		}
		switch obj := obj.(type) {
		case *types.Const:
			if ntyp := Make(obj.Type(), atyp, lookup); ntyp != nil {
				typesutil.SetConstType(obj, ntyp)
			}
		case *types.Var:
			if ntyp := Make(obj.Type(), atyp, lookup); ntyp != nil {
				typesutil.SetVarType(obj, ntyp)
			}
		case *types.Func:
			if ntyp := Make(obj.Type(), atyp, lookup); ntyp != nil {
				typesutil.SetFuncType(obj, ntyp)
			}
		case *types.TypeName:
			if named, ok := obj.Type().(*types.Named); ok {
				if atyp, ok := atyp.(*alias.Named); ok {
					for i := 0; i < named.NumMethods(); i++ {
						m := named.Method(i)
						if afn, ok := atyp.Methods[m.Name()]; ok {
							if ntyp := Make(m.Type(), afn, lookup); ntyp != nil {
								typesutil.SetFuncType(m, ntyp)
							}
						}
					}
					if atyp.Underlying != nil {
						if t := Make(named.Underlying(), atyp.Underlying, lookup); t != nil {
							named.SetUnderlying(t)
						}
					}
				}
			}
		}
	}
}
