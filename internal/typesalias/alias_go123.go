//go:build go1.23
// +build go1.23

package typesalias

import "go/types"

type Alias = types.Alias

func NewAlias(obj *types.TypeName, rhs types.Type) *Alias {
	return types.NewAlias(obj, rhs)
}

func Unalias(t types.Type) types.Type {
	return types.Unalias(t)
}

func Rhs(t *Alias) types.Type {
	return t.Rhs()
}
