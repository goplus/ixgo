// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package types

import (
	q "go/types"

	bytes "bytes"
	"github.com/goplus/ixgo"
	ast "go/ast"
	constant "go/constant"
	token "go/token"
	maphash "hash/maphash"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("go/types", map[string]ixgo.DirectCallAdapter{
		"(*Alias).Obj":                    method_ptr_Alias_Obj,
		"(*Alias).Origin":                 method_ptr_Alias_Origin,
		"(*Alias).Rhs":                    method_ptr_Alias_Rhs,
		"(*Alias).SetTypeParams":          method_ptr_Alias_SetTypeParams,
		"(*Alias).String":                 method_ptr_Alias_String,
		"(*Alias).TypeArgs":               method_ptr_Alias_TypeArgs,
		"(*Alias).TypeParams":             method_ptr_Alias_TypeParams,
		"(*Alias).Underlying":             method_ptr_Alias_Underlying,
		"(*ArgumentError).Error":          method_ptr_ArgumentError_Error,
		"(*ArgumentError).Unwrap":         method_ptr_ArgumentError_Unwrap,
		"(*Array).Elem":                   method_ptr_Array_Elem,
		"(*Array).Len":                    method_ptr_Array_Len,
		"(*Array).String":                 method_ptr_Array_String,
		"(*Array).Underlying":             method_ptr_Array_Underlying,
		"(*Basic).Info":                   method_ptr_Basic_Info,
		"(*Basic).Kind":                   method_ptr_Basic_Kind,
		"(*Basic).Name":                   method_ptr_Basic_Name,
		"(*Basic).String":                 method_ptr_Basic_String,
		"(*Basic).Underlying":             method_ptr_Basic_Underlying,
		"(*Builtin).String":               method_ptr_Builtin_String,
		"(*Chan).Dir":                     method_ptr_Chan_Dir,
		"(*Chan).Elem":                    method_ptr_Chan_Elem,
		"(*Chan).String":                  method_ptr_Chan_String,
		"(*Chan).Underlying":              method_ptr_Chan_Underlying,
		"(*Checker).Files":                method_ptr_Checker_Files,
		"(*Const).String":                 method_ptr_Const_String,
		"(*Const).Val":                    method_ptr_Const_Val,
		"(*Error).Error":                  method_ptr_Error_Error,
		"(*Func).FullName":                method_ptr_Func_FullName,
		"(*Func).Origin":                  method_ptr_Func_Origin,
		"(*Func).Pkg":                     method_ptr_Func_Pkg,
		"(*Func).Scope":                   method_ptr_Func_Scope,
		"(*Func).Signature":               method_ptr_Func_Signature,
		"(*Func).String":                  method_ptr_Func_String,
		"(*Hasher).Equal":                 method_ptr_Hasher_Equal,
		"(*Hasher).Hash":                  method_ptr_Hasher_Hash,
		"(*HasherIgnoreTags).Equal":       method_ptr_HasherIgnoreTags_Equal,
		"(*HasherIgnoreTags).Hash":        method_ptr_HasherIgnoreTags_Hash,
		"(*Info).ObjectOf":                method_ptr_Info_ObjectOf,
		"(*Info).PkgNameOf":               method_ptr_Info_PkgNameOf,
		"(*Info).TypeOf":                  method_ptr_Info_TypeOf,
		"(*Initializer).String":           method_ptr_Initializer_String,
		"(*Instance).String":              method_ptr_Instance_String,
		"(*Interface).Complete":           method_ptr_Interface_Complete,
		"(*Interface).Embedded":           method_ptr_Interface_Embedded,
		"(*Interface).EmbeddedType":       method_ptr_Interface_EmbeddedType,
		"(*Interface).EmbeddedTypes":      method_ptr_Interface_EmbeddedTypes,
		"(*Interface).Empty":              method_ptr_Interface_Empty,
		"(*Interface).ExplicitMethod":     method_ptr_Interface_ExplicitMethod,
		"(*Interface).ExplicitMethods":    method_ptr_Interface_ExplicitMethods,
		"(*Interface).IsComparable":       method_ptr_Interface_IsComparable,
		"(*Interface).IsImplicit":         method_ptr_Interface_IsImplicit,
		"(*Interface).IsMethodSet":        method_ptr_Interface_IsMethodSet,
		"(*Interface).MarkImplicit":       method_ptr_Interface_MarkImplicit,
		"(*Interface).Method":             method_ptr_Interface_Method,
		"(*Interface).Methods":            method_ptr_Interface_Methods,
		"(*Interface).NumEmbeddeds":       method_ptr_Interface_NumEmbeddeds,
		"(*Interface).NumExplicitMethods": method_ptr_Interface_NumExplicitMethods,
		"(*Interface).NumMethods":         method_ptr_Interface_NumMethods,
		"(*Interface).String":             method_ptr_Interface_String,
		"(*Interface).Underlying":         method_ptr_Interface_Underlying,
		"(*Label).String":                 method_ptr_Label_String,
		"(*Map).Elem":                     method_ptr_Map_Elem,
		"(*Map).Key":                      method_ptr_Map_Key,
		"(*Map).String":                   method_ptr_Map_String,
		"(*Map).Underlying":               method_ptr_Map_Underlying,
		"(*MethodSet).At":                 method_ptr_MethodSet_At,
		"(*MethodSet).Len":                method_ptr_MethodSet_Len,
		"(*MethodSet).Lookup":             method_ptr_MethodSet_Lookup,
		"(*MethodSet).Methods":            method_ptr_MethodSet_Methods,
		"(*MethodSet).String":             method_ptr_MethodSet_String,
		"(*Named).AddMethod":              method_ptr_Named_AddMethod,
		"(*Named).Method":                 method_ptr_Named_Method,
		"(*Named).Methods":                method_ptr_Named_Methods,
		"(*Named).NumMethods":             method_ptr_Named_NumMethods,
		"(*Named).Obj":                    method_ptr_Named_Obj,
		"(*Named).Origin":                 method_ptr_Named_Origin,
		"(*Named).SetTypeParams":          method_ptr_Named_SetTypeParams,
		"(*Named).SetUnderlying":          method_ptr_Named_SetUnderlying,
		"(*Named).String":                 method_ptr_Named_String,
		"(*Named).TypeArgs":               method_ptr_Named_TypeArgs,
		"(*Named).TypeParams":             method_ptr_Named_TypeParams,
		"(*Named).Underlying":             method_ptr_Named_Underlying,
		"(*Nil).String":                   method_ptr_Nil_String,
		"(*Package).Complete":             method_ptr_Package_Complete,
		"(*Package).GoVersion":            method_ptr_Package_GoVersion,
		"(*Package).Imports":              method_ptr_Package_Imports,
		"(*Package).MarkComplete":         method_ptr_Package_MarkComplete,
		"(*Package).Name":                 method_ptr_Package_Name,
		"(*Package).Path":                 method_ptr_Package_Path,
		"(*Package).Scope":                method_ptr_Package_Scope,
		"(*Package).SetImports":           method_ptr_Package_SetImports,
		"(*Package).SetName":              method_ptr_Package_SetName,
		"(*Package).String":               method_ptr_Package_String,
		"(*PkgName).Imported":             method_ptr_PkgName_Imported,
		"(*PkgName).String":               method_ptr_PkgName_String,
		"(*Pointer).Elem":                 method_ptr_Pointer_Elem,
		"(*Pointer).String":               method_ptr_Pointer_String,
		"(*Pointer).Underlying":           method_ptr_Pointer_Underlying,
		"(*Scope).Child":                  method_ptr_Scope_Child,
		"(*Scope).Children":               method_ptr_Scope_Children,
		"(*Scope).Contains":               method_ptr_Scope_Contains,
		"(*Scope).End":                    method_ptr_Scope_End,
		"(*Scope).Innermost":              method_ptr_Scope_Innermost,
		"(*Scope).Insert":                 method_ptr_Scope_Insert,
		"(*Scope).Len":                    method_ptr_Scope_Len,
		"(*Scope).Lookup":                 method_ptr_Scope_Lookup,
		"(*Scope).Names":                  method_ptr_Scope_Names,
		"(*Scope).NumChildren":            method_ptr_Scope_NumChildren,
		"(*Scope).Parent":                 method_ptr_Scope_Parent,
		"(*Scope).Pos":                    method_ptr_Scope_Pos,
		"(*Scope).String":                 method_ptr_Scope_String,
		"(*Scope).WriteTo":                method_ptr_Scope_WriteTo,
		"(*Selection).Index":              method_ptr_Selection_Index,
		"(*Selection).Indirect":           method_ptr_Selection_Indirect,
		"(*Selection).Kind":               method_ptr_Selection_Kind,
		"(*Selection).Obj":                method_ptr_Selection_Obj,
		"(*Selection).Recv":               method_ptr_Selection_Recv,
		"(*Selection).String":             method_ptr_Selection_String,
		"(*Selection).Type":               method_ptr_Selection_Type,
		"(*Signature).Params":             method_ptr_Signature_Params,
		"(*Signature).Recv":               method_ptr_Signature_Recv,
		"(*Signature).RecvTypeParams":     method_ptr_Signature_RecvTypeParams,
		"(*Signature).Results":            method_ptr_Signature_Results,
		"(*Signature).String":             method_ptr_Signature_String,
		"(*Signature).TypeParams":         method_ptr_Signature_TypeParams,
		"(*Signature).Underlying":         method_ptr_Signature_Underlying,
		"(*Signature).Variadic":           method_ptr_Signature_Variadic,
		"(*Slice).Elem":                   method_ptr_Slice_Elem,
		"(*Slice).String":                 method_ptr_Slice_String,
		"(*Slice).Underlying":             method_ptr_Slice_Underlying,
		"(*StdSizes).Alignof":             method_ptr_StdSizes_Alignof,
		"(*StdSizes).Offsetsof":           method_ptr_StdSizes_Offsetsof,
		"(*StdSizes).Sizeof":              method_ptr_StdSizes_Sizeof,
		"(*Struct).Field":                 method_ptr_Struct_Field,
		"(*Struct).Fields":                method_ptr_Struct_Fields,
		"(*Struct).NumFields":             method_ptr_Struct_NumFields,
		"(*Struct).String":                method_ptr_Struct_String,
		"(*Struct).Tag":                   method_ptr_Struct_Tag,
		"(*Struct).Underlying":            method_ptr_Struct_Underlying,
		"(*Term).String":                  method_ptr_Term_String,
		"(*Term).Tilde":                   method_ptr_Term_Tilde,
		"(*Term).Type":                    method_ptr_Term_Type,
		"(*Tuple).At":                     method_ptr_Tuple_At,
		"(*Tuple).Len":                    method_ptr_Tuple_Len,
		"(*Tuple).String":                 method_ptr_Tuple_String,
		"(*Tuple).Underlying":             method_ptr_Tuple_Underlying,
		"(*Tuple).Variables":              method_ptr_Tuple_Variables,
		"(*TypeAndValue).Addressable":     method_ptr_TypeAndValue_Addressable,
		"(*TypeAndValue).Assignable":      method_ptr_TypeAndValue_Assignable,
		"(*TypeAndValue).HasOk":           method_ptr_TypeAndValue_HasOk,
		"(*TypeAndValue).IsBuiltin":       method_ptr_TypeAndValue_IsBuiltin,
		"(*TypeAndValue).IsNil":           method_ptr_TypeAndValue_IsNil,
		"(*TypeAndValue).IsType":          method_ptr_TypeAndValue_IsType,
		"(*TypeAndValue).IsValue":         method_ptr_TypeAndValue_IsValue,
		"(*TypeAndValue).IsVoid":          method_ptr_TypeAndValue_IsVoid,
		"(*TypeList).At":                  method_ptr_TypeList_At,
		"(*TypeList).Len":                 method_ptr_TypeList_Len,
		"(*TypeList).String":              method_ptr_TypeList_String,
		"(*TypeList).Types":               method_ptr_TypeList_Types,
		"(*TypeName).IsAlias":             method_ptr_TypeName_IsAlias,
		"(*TypeName).String":              method_ptr_TypeName_String,
		"(*TypeParam).Constraint":         method_ptr_TypeParam_Constraint,
		"(*TypeParam).Index":              method_ptr_TypeParam_Index,
		"(*TypeParam).Obj":                method_ptr_TypeParam_Obj,
		"(*TypeParam).SetConstraint":      method_ptr_TypeParam_SetConstraint,
		"(*TypeParam).String":             method_ptr_TypeParam_String,
		"(*TypeParam).Underlying":         method_ptr_TypeParam_Underlying,
		"(*TypeParamList).At":             method_ptr_TypeParamList_At,
		"(*TypeParamList).Len":            method_ptr_TypeParamList_Len,
		"(*TypeParamList).String":         method_ptr_TypeParamList_String,
		"(*TypeParamList).TypeParams":     method_ptr_TypeParamList_TypeParams,
		"(*Union).Len":                    method_ptr_Union_Len,
		"(*Union).String":                 method_ptr_Union_String,
		"(*Union).Term":                   method_ptr_Union_Term,
		"(*Union).Terms":                  method_ptr_Union_Terms,
		"(*Union).Underlying":             method_ptr_Union_Underlying,
		"(*Var).Anonymous":                method_ptr_Var_Anonymous,
		"(*Var).Embedded":                 method_ptr_Var_Embedded,
		"(*Var).IsField":                  method_ptr_Var_IsField,
		"(*Var).Kind":                     method_ptr_Var_Kind,
		"(*Var).Origin":                   method_ptr_Var_Origin,
		"(*Var).SetKind":                  method_ptr_Var_SetKind,
		"(*Var).String":                   method_ptr_Var_String,
		"(*VarKind).String":               method_ptr_VarKind_String,
		"(Error).Error":                   method_Error_Error,
		"(Hasher).Equal":                  method_Hasher_Equal,
		"(Hasher).Hash":                   method_Hasher_Hash,
		"(HasherIgnoreTags).Equal":        method_HasherIgnoreTags_Equal,
		"(HasherIgnoreTags).Hash":         method_HasherIgnoreTags_Hash,
		"(Instance).String":               method_Instance_String,
		"(TypeAndValue).Addressable":      method_TypeAndValue_Addressable,
		"(TypeAndValue).Assignable":       method_TypeAndValue_Assignable,
		"(TypeAndValue).HasOk":            method_TypeAndValue_HasOk,
		"(TypeAndValue).IsBuiltin":        method_TypeAndValue_IsBuiltin,
		"(TypeAndValue).IsNil":            method_TypeAndValue_IsNil,
		"(TypeAndValue).IsType":           method_TypeAndValue_IsType,
		"(TypeAndValue).IsValue":          method_TypeAndValue_IsValue,
		"(TypeAndValue).IsVoid":           method_TypeAndValue_IsVoid,
		"(VarKind).String":                method_VarKind_String,
		"AssertableTo":                    func_AssertableTo,
		"AssignableTo":                    func_AssignableTo,
		"CheckExpr":                       func_CheckExpr,
		"Comparable":                      func_Comparable,
		"ConvertibleTo":                   func_ConvertibleTo,
		"DefPredeclaredTestFuncs":         func_DefPredeclaredTestFuncs,
		"Default":                         func_Default,
		"ExprString":                      func_ExprString,
		"Id":                              func_Id,
		"Identical":                       func_Identical,
		"IdenticalIgnoreTags":             func_IdenticalIgnoreTags,
		"Implements":                      func_Implements,
		"IsInterface":                     func_IsInterface,
		"NewAlias":                        func_NewAlias,
		"NewArray":                        func_NewArray,
		"NewChan":                         func_NewChan,
		"NewChecker":                      func_NewChecker,
		"NewConst":                        func_NewConst,
		"NewContext":                      func_NewContext,
		"NewField":                        func_NewField,
		"NewFunc":                         func_NewFunc,
		"NewInterface":                    func_NewInterface,
		"NewInterfaceType":                func_NewInterfaceType,
		"NewLabel":                        func_NewLabel,
		"NewMap":                          func_NewMap,
		"NewMethodSet":                    func_NewMethodSet,
		"NewNamed":                        func_NewNamed,
		"NewPackage":                      func_NewPackage,
		"NewParam":                        func_NewParam,
		"NewPkgName":                      func_NewPkgName,
		"NewPointer":                      func_NewPointer,
		"NewScope":                        func_NewScope,
		"NewSignature":                    func_NewSignature,
		"NewSignatureType":                func_NewSignatureType,
		"NewSlice":                        func_NewSlice,
		"NewStruct":                       func_NewStruct,
		"NewTerm":                         func_NewTerm,
		"NewTuple":                        func_NewTuple,
		"NewTypeName":                     func_NewTypeName,
		"NewTypeParam":                    func_NewTypeParam,
		"NewUnion":                        func_NewUnion,
		"NewVar":                          func_NewVar,
		"ObjectString":                    func_ObjectString,
		"RelativeTo":                      func_RelativeTo,
		"Satisfies":                       func_Satisfies,
		"SelectionString":                 func_SelectionString,
		"SizesFor":                        func_SizesFor,
		"TypeString":                      func_TypeString,
		"Unalias":                         func_Unalias,
		"WriteExpr":                       func_WriteExpr,
		"WriteSignature":                  func_WriteSignature,
		"WriteType":                       func_WriteType,
	})
}

func method_ptr_Alias_Obj(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alias).Obj(ixgo.DirectCallArg[*q.Alias](ctx, 0)))
}

func method_ptr_Alias_Origin(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alias).Origin(ixgo.DirectCallArg[*q.Alias](ctx, 0)))
}

func method_ptr_Alias_Rhs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alias).Rhs(ixgo.DirectCallArg[*q.Alias](ctx, 0)))
}

func method_ptr_Alias_SetTypeParams(ctx ixgo.DirectCallContext) {
	(*q.Alias).SetTypeParams(ixgo.DirectCallArg[*q.Alias](ctx, 0), ixgo.DirectCallArg[[]*q.TypeParam](ctx, 1))
}

func method_ptr_Alias_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alias).String(ixgo.DirectCallArg[*q.Alias](ctx, 0)))
}

func method_ptr_Alias_TypeArgs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alias).TypeArgs(ixgo.DirectCallArg[*q.Alias](ctx, 0)))
}

func method_ptr_Alias_TypeParams(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alias).TypeParams(ixgo.DirectCallArg[*q.Alias](ctx, 0)))
}

func method_ptr_Alias_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Alias).Underlying(ixgo.DirectCallArg[*q.Alias](ctx, 0)))
}

func method_ptr_ArgumentError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ArgumentError).Error(ixgo.DirectCallArg[*q.ArgumentError](ctx, 0)))
}

func method_ptr_ArgumentError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ArgumentError).Unwrap(ixgo.DirectCallArg[*q.ArgumentError](ctx, 0)))
}

func method_ptr_Array_Elem(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Array).Elem(ixgo.DirectCallArg[*q.Array](ctx, 0)))
}

func method_ptr_Array_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Array).Len(ixgo.DirectCallArg[*q.Array](ctx, 0)))
}

func method_ptr_Array_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Array).String(ixgo.DirectCallArg[*q.Array](ctx, 0)))
}

func method_ptr_Array_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Array).Underlying(ixgo.DirectCallArg[*q.Array](ctx, 0)))
}

func func_AssertableTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AssertableTo(ixgo.DirectCallArg[*q.Interface](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_AssignableTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.AssignableTo(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func method_ptr_Basic_Info(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Basic).Info(ixgo.DirectCallArg[*q.Basic](ctx, 0)))
}

func method_ptr_Basic_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Basic).Kind(ixgo.DirectCallArg[*q.Basic](ctx, 0)))
}

func method_ptr_Basic_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Basic).Name(ixgo.DirectCallArg[*q.Basic](ctx, 0)))
}

func method_ptr_Basic_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Basic).String(ixgo.DirectCallArg[*q.Basic](ctx, 0)))
}

func method_ptr_Basic_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Basic).Underlying(ixgo.DirectCallArg[*q.Basic](ctx, 0)))
}

func method_ptr_Builtin_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Builtin).String(ixgo.DirectCallArg[*q.Builtin](ctx, 0)))
}

func method_ptr_Chan_Dir(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Chan).Dir(ixgo.DirectCallArg[*q.Chan](ctx, 0)))
}

func method_ptr_Chan_Elem(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Chan).Elem(ixgo.DirectCallArg[*q.Chan](ctx, 0)))
}

func method_ptr_Chan_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Chan).String(ixgo.DirectCallArg[*q.Chan](ctx, 0)))
}

func method_ptr_Chan_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Chan).Underlying(ixgo.DirectCallArg[*q.Chan](ctx, 0)))
}

func func_CheckExpr(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CheckExpr(ixgo.DirectCallArg[*token.FileSet](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[token.Pos](ctx, 2), ixgo.DirectCallArg[ast.Expr](ctx, 3), ixgo.DirectCallArg[*q.Info](ctx, 4)))
}

func method_ptr_Checker_Files(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Checker).Files(ixgo.DirectCallArg[*q.Checker](ctx, 0), ixgo.DirectCallArg[[]*ast.File](ctx, 1)))
}

func func_Comparable(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Comparable(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func method_ptr_Const_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Const).String(ixgo.DirectCallArg[*q.Const](ctx, 0)))
}

func method_ptr_Const_Val(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Const).Val(ixgo.DirectCallArg[*q.Const](ctx, 0)))
}

func func_ConvertibleTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ConvertibleTo(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_DefPredeclaredTestFuncs(ctx ixgo.DirectCallContext) {
	q.DefPredeclaredTestFuncs()
}

func func_Default(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Default(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func method_Error_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Error.Error(ixgo.DirectCallArg[q.Error](ctx, 0)))
}

func method_ptr_Error_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Error).Error(ixgo.DirectCallArg[*q.Error](ctx, 0)))
}

func func_ExprString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ExprString(ixgo.DirectCallArg[ast.Expr](ctx, 0)))
}

func method_ptr_Func_FullName(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).FullName(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func method_ptr_Func_Origin(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).Origin(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func method_ptr_Func_Pkg(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).Pkg(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func method_ptr_Func_Scope(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).Scope(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func method_ptr_Func_Signature(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).Signature(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func method_ptr_Func_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Func).String(ixgo.DirectCallArg[*q.Func](ctx, 0)))
}

func method_Hasher_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Hasher.Equal(ixgo.DirectCallArg[q.Hasher](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1), ixgo.DirectCallArg[q.Type](ctx, 2)))
}

func method_ptr_Hasher_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Hasher).Equal(ixgo.DirectCallArg[*q.Hasher](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1), ixgo.DirectCallArg[q.Type](ctx, 2)))
}

func method_Hasher_Hash(ctx ixgo.DirectCallContext) {
	q.Hasher.Hash(ixgo.DirectCallArg[q.Hasher](ctx, 0), ixgo.DirectCallArg[*maphash.Hash](ctx, 1), ixgo.DirectCallArg[q.Type](ctx, 2))
}

func method_ptr_Hasher_Hash(ctx ixgo.DirectCallContext) {
	(*q.Hasher).Hash(ixgo.DirectCallArg[*q.Hasher](ctx, 0), ixgo.DirectCallArg[*maphash.Hash](ctx, 1), ixgo.DirectCallArg[q.Type](ctx, 2))
}

func method_HasherIgnoreTags_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.HasherIgnoreTags.Equal(ixgo.DirectCallArg[q.HasherIgnoreTags](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1), ixgo.DirectCallArg[q.Type](ctx, 2)))
}

func method_ptr_HasherIgnoreTags_Equal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.HasherIgnoreTags).Equal(ixgo.DirectCallArg[*q.HasherIgnoreTags](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1), ixgo.DirectCallArg[q.Type](ctx, 2)))
}

func method_HasherIgnoreTags_Hash(ctx ixgo.DirectCallContext) {
	q.HasherIgnoreTags.Hash(ixgo.DirectCallArg[q.HasherIgnoreTags](ctx, 0), ixgo.DirectCallArg[*maphash.Hash](ctx, 1), ixgo.DirectCallArg[q.Type](ctx, 2))
}

func method_ptr_HasherIgnoreTags_Hash(ctx ixgo.DirectCallContext) {
	(*q.HasherIgnoreTags).Hash(ixgo.DirectCallArg[*q.HasherIgnoreTags](ctx, 0), ixgo.DirectCallArg[*maphash.Hash](ctx, 1), ixgo.DirectCallArg[q.Type](ctx, 2))
}

func func_Id(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Id(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Identical(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Identical(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_IdenticalIgnoreTags(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IdenticalIgnoreTags(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_Implements(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Implements(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[*q.Interface](ctx, 1)))
}

func method_ptr_Info_ObjectOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Info).ObjectOf(ixgo.DirectCallArg[*q.Info](ctx, 0), ixgo.DirectCallArg[*ast.Ident](ctx, 1)))
}

func method_ptr_Info_PkgNameOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Info).PkgNameOf(ixgo.DirectCallArg[*q.Info](ctx, 0), ixgo.DirectCallArg[*ast.ImportSpec](ctx, 1)))
}

func method_ptr_Info_TypeOf(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Info).TypeOf(ixgo.DirectCallArg[*q.Info](ctx, 0), ixgo.DirectCallArg[ast.Expr](ctx, 1)))
}

func method_ptr_Initializer_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Initializer).String(ixgo.DirectCallArg[*q.Initializer](ctx, 0)))
}

func method_Instance_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Instance.String(ixgo.DirectCallArg[q.Instance](ctx, 0)))
}

func method_ptr_Instance_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Instance).String(ixgo.DirectCallArg[*q.Instance](ctx, 0)))
}

func method_ptr_Interface_Complete(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).Complete(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_Embedded(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).Embedded(ixgo.DirectCallArg[*q.Interface](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Interface_EmbeddedType(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).EmbeddedType(ixgo.DirectCallArg[*q.Interface](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Interface_EmbeddedTypes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).EmbeddedTypes(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_Empty(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).Empty(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_ExplicitMethod(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).ExplicitMethod(ixgo.DirectCallArg[*q.Interface](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Interface_ExplicitMethods(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).ExplicitMethods(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_IsComparable(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).IsComparable(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_IsImplicit(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).IsImplicit(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_IsMethodSet(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).IsMethodSet(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_MarkImplicit(ctx ixgo.DirectCallContext) {
	(*q.Interface).MarkImplicit(ixgo.DirectCallArg[*q.Interface](ctx, 0))
}

func method_ptr_Interface_Method(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).Method(ixgo.DirectCallArg[*q.Interface](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Interface_Methods(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).Methods(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_NumEmbeddeds(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).NumEmbeddeds(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_NumExplicitMethods(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).NumExplicitMethods(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_NumMethods(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).NumMethods(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).String(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func method_ptr_Interface_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Interface).Underlying(ixgo.DirectCallArg[*q.Interface](ctx, 0)))
}

func func_IsInterface(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsInterface(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func method_ptr_Label_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Label).String(ixgo.DirectCallArg[*q.Label](ctx, 0)))
}

func method_ptr_Map_Elem(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Map).Elem(ixgo.DirectCallArg[*q.Map](ctx, 0)))
}

func method_ptr_Map_Key(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Map).Key(ixgo.DirectCallArg[*q.Map](ctx, 0)))
}

func method_ptr_Map_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Map).String(ixgo.DirectCallArg[*q.Map](ctx, 0)))
}

func method_ptr_Map_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Map).Underlying(ixgo.DirectCallArg[*q.Map](ctx, 0)))
}

func method_ptr_MethodSet_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MethodSet).At(ixgo.DirectCallArg[*q.MethodSet](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_MethodSet_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MethodSet).Len(ixgo.DirectCallArg[*q.MethodSet](ctx, 0)))
}

func method_ptr_MethodSet_Lookup(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MethodSet).Lookup(ixgo.DirectCallArg[*q.MethodSet](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_MethodSet_Methods(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MethodSet).Methods(ixgo.DirectCallArg[*q.MethodSet](ctx, 0)))
}

func method_ptr_MethodSet_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.MethodSet).String(ixgo.DirectCallArg[*q.MethodSet](ctx, 0)))
}

func method_ptr_Named_AddMethod(ctx ixgo.DirectCallContext) {
	(*q.Named).AddMethod(ixgo.DirectCallArg[*q.Named](ctx, 0), ixgo.DirectCallArg[*q.Func](ctx, 1))
}

func method_ptr_Named_Method(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Named).Method(ixgo.DirectCallArg[*q.Named](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Named_Methods(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Named).Methods(ixgo.DirectCallArg[*q.Named](ctx, 0)))
}

func method_ptr_Named_NumMethods(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Named).NumMethods(ixgo.DirectCallArg[*q.Named](ctx, 0)))
}

func method_ptr_Named_Obj(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Named).Obj(ixgo.DirectCallArg[*q.Named](ctx, 0)))
}

func method_ptr_Named_Origin(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Named).Origin(ixgo.DirectCallArg[*q.Named](ctx, 0)))
}

func method_ptr_Named_SetTypeParams(ctx ixgo.DirectCallContext) {
	(*q.Named).SetTypeParams(ixgo.DirectCallArg[*q.Named](ctx, 0), ixgo.DirectCallArg[[]*q.TypeParam](ctx, 1))
}

func method_ptr_Named_SetUnderlying(ctx ixgo.DirectCallContext) {
	(*q.Named).SetUnderlying(ixgo.DirectCallArg[*q.Named](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1))
}

func method_ptr_Named_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Named).String(ixgo.DirectCallArg[*q.Named](ctx, 0)))
}

func method_ptr_Named_TypeArgs(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Named).TypeArgs(ixgo.DirectCallArg[*q.Named](ctx, 0)))
}

func method_ptr_Named_TypeParams(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Named).TypeParams(ixgo.DirectCallArg[*q.Named](ctx, 0)))
}

func method_ptr_Named_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Named).Underlying(ixgo.DirectCallArg[*q.Named](ctx, 0)))
}

func func_NewAlias(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewAlias(ixgo.DirectCallArg[*q.TypeName](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_NewArray(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewArray(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_NewChan(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewChan(ixgo.DirectCallArg[q.ChanDir](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_NewChecker(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewChecker(ixgo.DirectCallArg[*q.Config](ctx, 0), ixgo.DirectCallArg[*token.FileSet](ctx, 1), ixgo.DirectCallArg[*q.Package](ctx, 2), ixgo.DirectCallArg[*q.Info](ctx, 3)))
}

func func_NewConst(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewConst(ixgo.DirectCallArg[token.Pos](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[q.Type](ctx, 3), ixgo.DirectCallArg[constant.Value](ctx, 4)))
}

func func_NewContext(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewContext())
}

func func_NewField(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewField(ixgo.DirectCallArg[token.Pos](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[q.Type](ctx, 3), ixgo.DirectCallArg[bool](ctx, 4)))
}

func func_NewFunc(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewFunc(ixgo.DirectCallArg[token.Pos](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[*q.Signature](ctx, 3)))
}

func func_NewInterface(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewInterface(ixgo.DirectCallArg[[]*q.Func](ctx, 0), ixgo.DirectCallArg[[]*q.Named](ctx, 1)))
}

func func_NewInterfaceType(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewInterfaceType(ixgo.DirectCallArg[[]*q.Func](ctx, 0), ixgo.DirectCallArg[[]q.Type](ctx, 1)))
}

func func_NewLabel(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewLabel(ixgo.DirectCallArg[token.Pos](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func func_NewMap(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewMap(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_NewMethodSet(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewMethodSet(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func func_NewNamed(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewNamed(ixgo.DirectCallArg[*q.TypeName](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1), ixgo.DirectCallArg[[]*q.Func](ctx, 2)))
}

func func_NewPackage(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewPackage(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_NewParam(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewParam(ixgo.DirectCallArg[token.Pos](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[q.Type](ctx, 3)))
}

func func_NewPkgName(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewPkgName(ixgo.DirectCallArg[token.Pos](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[*q.Package](ctx, 3)))
}

func func_NewPointer(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewPointer(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func func_NewScope(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewScope(ixgo.DirectCallArg[*q.Scope](ctx, 0), ixgo.DirectCallArg[token.Pos](ctx, 1), ixgo.DirectCallArg[token.Pos](ctx, 2), ixgo.DirectCallArg[string](ctx, 3)))
}

func func_NewSignature(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewSignature(ixgo.DirectCallArg[*q.Var](ctx, 0), ixgo.DirectCallArg[*q.Tuple](ctx, 1), ixgo.DirectCallArg[*q.Tuple](ctx, 2), ixgo.DirectCallArg[bool](ctx, 3)))
}

func func_NewSignatureType(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewSignatureType(ixgo.DirectCallArg[*q.Var](ctx, 0), ixgo.DirectCallArg[[]*q.TypeParam](ctx, 1), ixgo.DirectCallArg[[]*q.TypeParam](ctx, 2), ixgo.DirectCallArg[*q.Tuple](ctx, 3), ixgo.DirectCallArg[*q.Tuple](ctx, 4), ixgo.DirectCallArg[bool](ctx, 5)))
}

func func_NewSlice(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewSlice(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func func_NewStruct(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewStruct(ixgo.DirectCallArg[[]*q.Var](ctx, 0), ixgo.DirectCallArg[[]string](ctx, 1)))
}

func func_NewTerm(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewTerm(ixgo.DirectCallArg[bool](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_NewTuple(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewTuple(ixgo.DirectCallArg[[]*q.Var](ctx, 0)...))
}

func func_NewTypeName(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewTypeName(ixgo.DirectCallArg[token.Pos](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[q.Type](ctx, 3)))
}

func func_NewTypeParam(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewTypeParam(ixgo.DirectCallArg[*q.TypeName](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func func_NewUnion(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewUnion(ixgo.DirectCallArg[[]*q.Term](ctx, 0)))
}

func func_NewVar(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewVar(ixgo.DirectCallArg[token.Pos](ctx, 0), ixgo.DirectCallArg[*q.Package](ctx, 1), ixgo.DirectCallArg[string](ctx, 2), ixgo.DirectCallArg[q.Type](ctx, 3)))
}

func method_ptr_Nil_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Nil).String(ixgo.DirectCallArg[*q.Nil](ctx, 0)))
}

func func_ObjectString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ObjectString(ixgo.DirectCallArg[q.Object](ctx, 0), ixgo.DirectCallArg[q.Qualifier](ctx, 1)))
}

func method_ptr_Package_Complete(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Complete(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_Package_GoVersion(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).GoVersion(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_Package_Imports(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Imports(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_Package_MarkComplete(ctx ixgo.DirectCallContext) {
	(*q.Package).MarkComplete(ixgo.DirectCallArg[*q.Package](ctx, 0))
}

func method_ptr_Package_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Name(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_Package_Path(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Path(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_Package_Scope(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).Scope(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_Package_SetImports(ctx ixgo.DirectCallContext) {
	(*q.Package).SetImports(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[[]*q.Package](ctx, 1))
}

func method_ptr_Package_SetName(ctx ixgo.DirectCallContext) {
	(*q.Package).SetName(ixgo.DirectCallArg[*q.Package](ctx, 0), ixgo.DirectCallArg[string](ctx, 1))
}

func method_ptr_Package_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Package).String(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func method_ptr_PkgName_Imported(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PkgName).Imported(ixgo.DirectCallArg[*q.PkgName](ctx, 0)))
}

func method_ptr_PkgName_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.PkgName).String(ixgo.DirectCallArg[*q.PkgName](ctx, 0)))
}

func method_ptr_Pointer_Elem(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pointer).Elem(ixgo.DirectCallArg[*q.Pointer](ctx, 0)))
}

func method_ptr_Pointer_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pointer).String(ixgo.DirectCallArg[*q.Pointer](ctx, 0)))
}

func method_ptr_Pointer_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Pointer).Underlying(ixgo.DirectCallArg[*q.Pointer](ctx, 0)))
}

func func_RelativeTo(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RelativeTo(ixgo.DirectCallArg[*q.Package](ctx, 0)))
}

func func_Satisfies(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Satisfies(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[*q.Interface](ctx, 1)))
}

func method_ptr_Scope_Child(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Child(ixgo.DirectCallArg[*q.Scope](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Scope_Children(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Children(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func method_ptr_Scope_Contains(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Contains(ixgo.DirectCallArg[*q.Scope](ctx, 0), ixgo.DirectCallArg[token.Pos](ctx, 1)))
}

func method_ptr_Scope_End(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).End(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func method_ptr_Scope_Innermost(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Innermost(ixgo.DirectCallArg[*q.Scope](ctx, 0), ixgo.DirectCallArg[token.Pos](ctx, 1)))
}

func method_ptr_Scope_Insert(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Insert(ixgo.DirectCallArg[*q.Scope](ctx, 0), ixgo.DirectCallArg[q.Object](ctx, 1)))
}

func method_ptr_Scope_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Len(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func method_ptr_Scope_Lookup(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Lookup(ixgo.DirectCallArg[*q.Scope](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Scope_Names(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Names(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func method_ptr_Scope_NumChildren(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).NumChildren(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func method_ptr_Scope_Parent(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Parent(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func method_ptr_Scope_Pos(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).Pos(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func method_ptr_Scope_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scope).String(ixgo.DirectCallArg[*q.Scope](ctx, 0)))
}

func method_ptr_Scope_WriteTo(ctx ixgo.DirectCallContext) {
	(*q.Scope).WriteTo(ixgo.DirectCallArg[*q.Scope](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[bool](ctx, 3))
}

func method_ptr_Selection_Index(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Selection).Index(ixgo.DirectCallArg[*q.Selection](ctx, 0)))
}

func method_ptr_Selection_Indirect(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Selection).Indirect(ixgo.DirectCallArg[*q.Selection](ctx, 0)))
}

func method_ptr_Selection_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Selection).Kind(ixgo.DirectCallArg[*q.Selection](ctx, 0)))
}

func method_ptr_Selection_Obj(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Selection).Obj(ixgo.DirectCallArg[*q.Selection](ctx, 0)))
}

func method_ptr_Selection_Recv(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Selection).Recv(ixgo.DirectCallArg[*q.Selection](ctx, 0)))
}

func method_ptr_Selection_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Selection).String(ixgo.DirectCallArg[*q.Selection](ctx, 0)))
}

func method_ptr_Selection_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Selection).Type(ixgo.DirectCallArg[*q.Selection](ctx, 0)))
}

func func_SelectionString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SelectionString(ixgo.DirectCallArg[*q.Selection](ctx, 0), ixgo.DirectCallArg[q.Qualifier](ctx, 1)))
}

func method_ptr_Signature_Params(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Signature).Params(ixgo.DirectCallArg[*q.Signature](ctx, 0)))
}

func method_ptr_Signature_Recv(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Signature).Recv(ixgo.DirectCallArg[*q.Signature](ctx, 0)))
}

func method_ptr_Signature_RecvTypeParams(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Signature).RecvTypeParams(ixgo.DirectCallArg[*q.Signature](ctx, 0)))
}

func method_ptr_Signature_Results(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Signature).Results(ixgo.DirectCallArg[*q.Signature](ctx, 0)))
}

func method_ptr_Signature_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Signature).String(ixgo.DirectCallArg[*q.Signature](ctx, 0)))
}

func method_ptr_Signature_TypeParams(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Signature).TypeParams(ixgo.DirectCallArg[*q.Signature](ctx, 0)))
}

func method_ptr_Signature_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Signature).Underlying(ixgo.DirectCallArg[*q.Signature](ctx, 0)))
}

func method_ptr_Signature_Variadic(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Signature).Variadic(ixgo.DirectCallArg[*q.Signature](ctx, 0)))
}

func func_SizesFor(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SizesFor(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Slice_Elem(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Slice).Elem(ixgo.DirectCallArg[*q.Slice](ctx, 0)))
}

func method_ptr_Slice_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Slice).String(ixgo.DirectCallArg[*q.Slice](ctx, 0)))
}

func method_ptr_Slice_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Slice).Underlying(ixgo.DirectCallArg[*q.Slice](ctx, 0)))
}

func method_ptr_StdSizes_Alignof(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StdSizes).Alignof(ixgo.DirectCallArg[*q.StdSizes](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func method_ptr_StdSizes_Offsetsof(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StdSizes).Offsetsof(ixgo.DirectCallArg[*q.StdSizes](ctx, 0), ixgo.DirectCallArg[[]*q.Var](ctx, 1)))
}

func method_ptr_StdSizes_Sizeof(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.StdSizes).Sizeof(ixgo.DirectCallArg[*q.StdSizes](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1)))
}

func method_ptr_Struct_Field(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Struct).Field(ixgo.DirectCallArg[*q.Struct](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Struct_Fields(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Struct).Fields(ixgo.DirectCallArg[*q.Struct](ctx, 0)))
}

func method_ptr_Struct_NumFields(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Struct).NumFields(ixgo.DirectCallArg[*q.Struct](ctx, 0)))
}

func method_ptr_Struct_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Struct).String(ixgo.DirectCallArg[*q.Struct](ctx, 0)))
}

func method_ptr_Struct_Tag(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Struct).Tag(ixgo.DirectCallArg[*q.Struct](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Struct_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Struct).Underlying(ixgo.DirectCallArg[*q.Struct](ctx, 0)))
}

func method_ptr_Term_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Term).String(ixgo.DirectCallArg[*q.Term](ctx, 0)))
}

func method_ptr_Term_Tilde(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Term).Tilde(ixgo.DirectCallArg[*q.Term](ctx, 0)))
}

func method_ptr_Term_Type(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Term).Type(ixgo.DirectCallArg[*q.Term](ctx, 0)))
}

func method_ptr_Tuple_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tuple).At(ixgo.DirectCallArg[*q.Tuple](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Tuple_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tuple).Len(ixgo.DirectCallArg[*q.Tuple](ctx, 0)))
}

func method_ptr_Tuple_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tuple).String(ixgo.DirectCallArg[*q.Tuple](ctx, 0)))
}

func method_ptr_Tuple_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tuple).Underlying(ixgo.DirectCallArg[*q.Tuple](ctx, 0)))
}

func method_ptr_Tuple_Variables(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Tuple).Variables(ixgo.DirectCallArg[*q.Tuple](ctx, 0)))
}

func method_TypeAndValue_Addressable(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeAndValue.Addressable(ixgo.DirectCallArg[q.TypeAndValue](ctx, 0)))
}

func method_ptr_TypeAndValue_Addressable(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAndValue).Addressable(ixgo.DirectCallArg[*q.TypeAndValue](ctx, 0)))
}

func method_TypeAndValue_Assignable(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeAndValue.Assignable(ixgo.DirectCallArg[q.TypeAndValue](ctx, 0)))
}

func method_ptr_TypeAndValue_Assignable(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAndValue).Assignable(ixgo.DirectCallArg[*q.TypeAndValue](ctx, 0)))
}

func method_TypeAndValue_HasOk(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeAndValue.HasOk(ixgo.DirectCallArg[q.TypeAndValue](ctx, 0)))
}

func method_ptr_TypeAndValue_HasOk(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAndValue).HasOk(ixgo.DirectCallArg[*q.TypeAndValue](ctx, 0)))
}

func method_TypeAndValue_IsBuiltin(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeAndValue.IsBuiltin(ixgo.DirectCallArg[q.TypeAndValue](ctx, 0)))
}

func method_ptr_TypeAndValue_IsBuiltin(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAndValue).IsBuiltin(ixgo.DirectCallArg[*q.TypeAndValue](ctx, 0)))
}

func method_TypeAndValue_IsNil(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeAndValue.IsNil(ixgo.DirectCallArg[q.TypeAndValue](ctx, 0)))
}

func method_ptr_TypeAndValue_IsNil(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAndValue).IsNil(ixgo.DirectCallArg[*q.TypeAndValue](ctx, 0)))
}

func method_TypeAndValue_IsType(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeAndValue.IsType(ixgo.DirectCallArg[q.TypeAndValue](ctx, 0)))
}

func method_ptr_TypeAndValue_IsType(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAndValue).IsType(ixgo.DirectCallArg[*q.TypeAndValue](ctx, 0)))
}

func method_TypeAndValue_IsValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeAndValue.IsValue(ixgo.DirectCallArg[q.TypeAndValue](ctx, 0)))
}

func method_ptr_TypeAndValue_IsValue(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAndValue).IsValue(ixgo.DirectCallArg[*q.TypeAndValue](ctx, 0)))
}

func method_TypeAndValue_IsVoid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeAndValue.IsVoid(ixgo.DirectCallArg[q.TypeAndValue](ctx, 0)))
}

func method_ptr_TypeAndValue_IsVoid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeAndValue).IsVoid(ixgo.DirectCallArg[*q.TypeAndValue](ctx, 0)))
}

func method_ptr_TypeList_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeList).At(ixgo.DirectCallArg[*q.TypeList](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_TypeList_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeList).Len(ixgo.DirectCallArg[*q.TypeList](ctx, 0)))
}

func method_ptr_TypeList_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeList).String(ixgo.DirectCallArg[*q.TypeList](ctx, 0)))
}

func method_ptr_TypeList_Types(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeList).Types(ixgo.DirectCallArg[*q.TypeList](ctx, 0)))
}

func method_ptr_TypeName_IsAlias(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeName).IsAlias(ixgo.DirectCallArg[*q.TypeName](ctx, 0)))
}

func method_ptr_TypeName_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeName).String(ixgo.DirectCallArg[*q.TypeName](ctx, 0)))
}

func method_ptr_TypeParam_Constraint(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeParam).Constraint(ixgo.DirectCallArg[*q.TypeParam](ctx, 0)))
}

func method_ptr_TypeParam_Index(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeParam).Index(ixgo.DirectCallArg[*q.TypeParam](ctx, 0)))
}

func method_ptr_TypeParam_Obj(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeParam).Obj(ixgo.DirectCallArg[*q.TypeParam](ctx, 0)))
}

func method_ptr_TypeParam_SetConstraint(ctx ixgo.DirectCallContext) {
	(*q.TypeParam).SetConstraint(ixgo.DirectCallArg[*q.TypeParam](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1))
}

func method_ptr_TypeParam_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeParam).String(ixgo.DirectCallArg[*q.TypeParam](ctx, 0)))
}

func method_ptr_TypeParam_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeParam).Underlying(ixgo.DirectCallArg[*q.TypeParam](ctx, 0)))
}

func method_ptr_TypeParamList_At(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeParamList).At(ixgo.DirectCallArg[*q.TypeParamList](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_TypeParamList_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeParamList).Len(ixgo.DirectCallArg[*q.TypeParamList](ctx, 0)))
}

func method_ptr_TypeParamList_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeParamList).String(ixgo.DirectCallArg[*q.TypeParamList](ctx, 0)))
}

func method_ptr_TypeParamList_TypeParams(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.TypeParamList).TypeParams(ixgo.DirectCallArg[*q.TypeParamList](ctx, 0)))
}

func func_TypeString(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TypeString(ixgo.DirectCallArg[q.Type](ctx, 0), ixgo.DirectCallArg[q.Qualifier](ctx, 1)))
}

func func_Unalias(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Unalias(ixgo.DirectCallArg[q.Type](ctx, 0)))
}

func method_ptr_Union_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Union).Len(ixgo.DirectCallArg[*q.Union](ctx, 0)))
}

func method_ptr_Union_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Union).String(ixgo.DirectCallArg[*q.Union](ctx, 0)))
}

func method_ptr_Union_Term(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Union).Term(ixgo.DirectCallArg[*q.Union](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Union_Terms(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Union).Terms(ixgo.DirectCallArg[*q.Union](ctx, 0)))
}

func method_ptr_Union_Underlying(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Union).Underlying(ixgo.DirectCallArg[*q.Union](ctx, 0)))
}

func method_ptr_Var_Anonymous(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Var).Anonymous(ixgo.DirectCallArg[*q.Var](ctx, 0)))
}

func method_ptr_Var_Embedded(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Var).Embedded(ixgo.DirectCallArg[*q.Var](ctx, 0)))
}

func method_ptr_Var_IsField(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Var).IsField(ixgo.DirectCallArg[*q.Var](ctx, 0)))
}

func method_ptr_Var_Kind(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Var).Kind(ixgo.DirectCallArg[*q.Var](ctx, 0)))
}

func method_ptr_Var_Origin(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Var).Origin(ixgo.DirectCallArg[*q.Var](ctx, 0)))
}

func method_ptr_Var_SetKind(ctx ixgo.DirectCallContext) {
	(*q.Var).SetKind(ixgo.DirectCallArg[*q.Var](ctx, 0), ixgo.DirectCallArg[q.VarKind](ctx, 1))
}

func method_ptr_Var_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Var).String(ixgo.DirectCallArg[*q.Var](ctx, 0)))
}

func method_VarKind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.VarKind.String(ixgo.DirectCallArg[q.VarKind](ctx, 0)))
}

func method_ptr_VarKind_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.VarKind).String(ixgo.DirectCallArg[*q.VarKind](ctx, 0)))
}

func func_WriteExpr(ctx ixgo.DirectCallContext) {
	q.WriteExpr(ixgo.DirectCallArg[*bytes.Buffer](ctx, 0), ixgo.DirectCallArg[ast.Expr](ctx, 1))
}

func func_WriteSignature(ctx ixgo.DirectCallContext) {
	q.WriteSignature(ixgo.DirectCallArg[*bytes.Buffer](ctx, 0), ixgo.DirectCallArg[*q.Signature](ctx, 1), ixgo.DirectCallArg[q.Qualifier](ctx, 2))
}

func func_WriteType(ctx ixgo.DirectCallContext) {
	q.WriteType(ixgo.DirectCallArg[*bytes.Buffer](ctx, 0), ixgo.DirectCallArg[q.Type](ctx, 1), ixgo.DirectCallArg[q.Qualifier](ctx, 2))
}
