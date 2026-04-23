//go:build go1.26
// +build go1.26

package typesutil

import (
	"go/token"
	"go/types"
)

type object struct {
	parent    *types.Scope
	pos       token.Pos
	pkg       *types.Package
	name      string
	typ       types.Type
	order_    uint32
	scopePos_ token.Pos
}
