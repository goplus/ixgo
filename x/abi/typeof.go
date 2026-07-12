package abi

import "unsafe"

// TypeOf returns the abi.Type of some value.
func TypeOf(a any) *Type {
	eface := *(*EmptyInterface)(unsafe.Pointer(&a))
	// Types are either static (for compiler-created types) or
	// heap-allocated but always reachable (for reflection-created
	// types, held in the central map). So there is no need to
	// escape types. noescape here help avoid unnecessary escape
	// of v.
	return (*Type)(NoEscape(unsafe.Pointer(eface.Type)))
}

// TypeFor returns the abi.Type for a type parameter.
func TypeFor[T any]() *Type {
	var v T
	if t := TypeOf(v); t != nil {
		return t // optimize for T being a non-interface kind
	}
	return TypeOf((*T)(nil)).Elem() // only for an interface kind
}
