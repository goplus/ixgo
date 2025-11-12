/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ixgo

import (
	"fmt"
	"go/token"
	"go/types"
	"log"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/goplus/ixgo/internal/typesalias"
	"github.com/goplus/reflectx"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/types/typeutil"
)

var (
	tyEmptyStruct = reflect.TypeOf((*struct{})(nil)).Elem()
	tyEmptyPtr    = reflect.TypeOf((*struct{})(nil))
	tyEmptyMap    = reflect.TypeOf((*map[struct{}]struct{})(nil)).Elem()
	tyEmptySlice  = reflect.TypeOf((*[]struct{})(nil)).Elem()
	tyEmptyArray  = reflect.TypeOf((*[0]struct{})(nil)).Elem()
	tyEmptyChan   = reflect.TypeOf((*chan struct{})(nil)).Elem()
	tyEmptyFunc   = reflect.TypeOf((*func())(nil)).Elem()
	tyUnusedFunc  = reflect.TypeOf((*func(_unused))(nil)).Elem()
)

type _unused struct{}

/*
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
*/

func emptyType(kind reflect.Kind) reflect.Type {
	switch kind {
	case reflect.Array:
		return tyEmptyArray
	case reflect.Chan:
		return tyEmptyChan
	case reflect.Func:
		return tyEmptyFunc
	case reflect.Interface:
		return tyEmptyInterface
	case reflect.Map:
		return tyEmptyMap
	case reflect.Ptr:
		return tyEmptyPtr
	case reflect.Slice:
		return tyEmptySlice
	case reflect.Struct:
		return tyEmptyStruct
	default:
		return xtypeTypes[kind]
	}
}

func toMockType(typ types.Type) reflect.Type {
	switch t := typ.(type) {
	case *types.Basic:
		kind := t.Kind()
		if kind > types.Invalid && kind < types.UntypedNil {
			return xtypeTypes[kind]
		}
		panic(fmt.Errorf("toMockType: invalid type %v", typ))
	case *types.Pointer:
		return tyEmptyPtr
	case *types.Slice:
		return tyEmptySlice
	case *types.Array:
		e := toMockType(t.Elem())
		return reflect.ArrayOf(int(t.Len()), e)
	case *types.Map:
		return tyEmptyMap
	case *types.Chan:
		return tyEmptyChan
	case *types.Struct:
		n := t.NumFields()
		fs := make([]reflect.StructField, n)
		for i := 0; i < n; i++ {
			ft := t.Field(i)
			fs[i].Name = "F" + strconv.Itoa(i)
			fs[i].Type = toMockType(ft.Type())
			fs[i].Anonymous = ft.Embedded()
		}
		return reflect.StructOf(fs)
	case *types.Named:
		return toMockType(typ.Underlying())
	case *types.Interface:
		return tyEmptyInterface
	case *types.Signature:
		in := t.Params().Len()
		out := t.Results().Len()
		if in+out == 0 {
			return tyEmptyFunc
		}
		ins := make([]reflect.Type, in)
		outs := make([]reflect.Type, out)
		variadic := t.Variadic()
		if variadic {
			for i := 0; i < in-1; i++ {
				ins[i] = tyEmptyStruct
			}
			ins[in-1] = tyEmptySlice
		} else {
			for i := 0; i < in; i++ {
				ins[i] = tyEmptyStruct
			}
		}
		for i := 0; i < out; i++ {
			outs[i] = tyEmptyStruct
		}
		return reflect.FuncOf(ins, outs, variadic)
	case *typesalias.Alias:
		return toMockType(typesalias.Unalias(t))
	default:
		panic(fmt.Errorf("toEmptyType: unreachable %T", typ))
	}
}

var xtypeTypes = [...]reflect.Type{
	types.Bool:          reflect.TypeOf(false),
	types.Int:           reflect.TypeOf(0),
	types.Int8:          reflect.TypeOf(int8(0)),
	types.Int16:         reflect.TypeOf(int16(0)),
	types.Int32:         reflect.TypeOf(int32(0)),
	types.Int64:         reflect.TypeOf(int64(0)),
	types.Uint:          reflect.TypeOf(uint(0)),
	types.Uint8:         reflect.TypeOf(uint8(0)),
	types.Uint16:        reflect.TypeOf(uint16(0)),
	types.Uint32:        reflect.TypeOf(uint32(0)),
	types.Uint64:        reflect.TypeOf(uint64(0)),
	types.Uintptr:       reflect.TypeOf(uintptr(0)),
	types.Float32:       reflect.TypeOf(float32(0)),
	types.Float64:       reflect.TypeOf(float64(0)),
	types.Complex64:     reflect.TypeOf(complex64(0)),
	types.Complex128:    reflect.TypeOf(complex128(0)),
	types.String:        reflect.TypeOf(""),
	types.UnsafePointer: reflect.TypeOf(unsafe.Pointer(nil)),

	types.UntypedBool:    reflect.TypeOf(false),
	types.UntypedInt:     reflect.TypeOf(0),
	types.UntypedRune:    reflect.TypeOf('a'),
	types.UntypedFloat:   reflect.TypeOf(0.1),
	types.UntypedComplex: reflect.TypeOf(0 + 1i),
	types.UntypedString:  reflect.TypeOf(""),
}

type FindMethod interface {
	FindMethod(mtyp reflect.Type, fn *types.Func) func([]reflect.Value) []reflect.Value
}

type TypesRecord struct {
	ctx     *Context
	rctx    *reflectx.Context // reflectx context
	loader  Loader
	finder  FindMethod
	rcache  map[reflect.Type]types.Type
	tcache  *typeutil.Map
	ncache  *typeutil.Map //nested cache
	fntargs string        //reflect type arguments used to instantiate the current func
	nested  map[*types.Named]int
	nstack  nestedStack
}

type mfnKey struct {
	typ      reflect.Type
	name     string
	indexs   string
	pointer  bool
	indirect bool
}

type mfnValue struct {
	fn func([]reflect.Value) []reflect.Value
	id int
}

func indexsToString(index []int) string {
	var sb strings.Builder
	for _, v := range index {
		sb.WriteByte('.')
		sb.WriteString(strconv.Itoa(v))
	}
	return sb.String()
}

func (r *TypesRecord) Release() {
	r.loader = nil
	r.rcache = nil
	r.tcache = nil
	r.ncache = nil
	r.finder = nil
	r.nested = nil
}

func NewTypesRecord(ctx *Context, rctx *reflectx.Context, loader Loader, finder FindMethod, nested map[*types.Named]int) *TypesRecord {
	return &TypesRecord{
		ctx:    ctx,
		rctx:   rctx,
		loader: loader,
		finder: finder,
		rcache: make(map[reflect.Type]types.Type),
		tcache: &typeutil.Map{},
		nested: nested,
	}
}

func (r *TypesRecord) LookupLocalTypes(rt reflect.Type) (typ types.Type, ok bool) {
	typ, ok = r.rcache[rt]
	return
}

func (r *TypesRecord) LookupTypes(rt reflect.Type) (typ types.Type, ok bool) {
	typ, ok = r.loader.LookupTypes(rt)
	if !ok {
		typ, ok = r.rcache[rt]
	}
	return
}

func (r *TypesRecord) saveType(typ types.Type, rt reflect.Type, nested bool) {
	r.rcache[rt] = typ
	if nested {
		r.ncache.Set(typ, rt)
		return
	}
	r.tcache.Set(typ, rt)
	return
}

func (r *TypesRecord) ToType(typ types.Type) (reflect.Type, bool) {
	if rt, ok, nested := r.LookupReflect(typ); ok {
		return rt, nested
	}
	var nested bool
	var rt reflect.Type
	switch t := typ.(type) {
	case *types.Basic:
		kind := t.Kind()
		if kind > types.Invalid && kind < types.UntypedNil {
			rt = xtypeTypes[kind]
		}
	case *types.Pointer:
		var elem reflect.Type
		elem, nested = r.ToType(t.Elem())
		rt = reflect.PtrTo(elem)
	case *types.Slice:
		var elem reflect.Type
		elem, nested = r.ToType(t.Elem())
		rt = reflect.SliceOf(elem)
	case *types.Array:
		var elem reflect.Type
		elem, nested = r.ToType(t.Elem())
		rt = reflect.ArrayOf(int(t.Len()), elem)
	case *types.Map:
		key, n1 := r.ToType(t.Key())
		elem, n2 := r.ToType(t.Elem())
		nested = n1 || n2
		rt = reflect.MapOf(key, elem)
	case *types.Chan:
		var elem reflect.Type
		elem, nested = r.ToType(t.Elem())
		rt = reflect.ChanOf(toReflectChanDir(t.Dir()), elem)
	case *types.Struct:
		rt, nested = r.toStructType(t, true)
	case *types.Named:
		rt, nested = r.toNamedType(t)
	case *types.Interface:
		rt, nested = r.toInterfaceType(t)
	case *types.Signature:
		in, n1 := r.ToTypeList(t.Params())
		out, n2 := r.ToTypeList(t.Results())
		nested = n1 || n2
		b := t.Variadic()
		if b && len(in) > 0 {
			last := in[len(in)-1]
			if last.Kind() == reflect.String {
				in[len(in)-1] = reflect.TypeOf([]byte{})
			}
		}
		rt = reflect.FuncOf(in, out, b)
	case *types.Tuple:
		_, nested = r.ToTypeList(t)
		rt = reflect.TypeOf((*_tuple)(nil)).Elem()
	case *typesalias.Alias:
		return r.ToType(typesalias.Unalias(t))
	default:
		panic(fmt.Errorf("ToType: not handled %T", typ))
	}
	r.saveType(typ, rt, nested)
	return rt, nested
}

type _tuple struct{}

func (r *TypesRecord) toInterfaceType(t *types.Interface) (reflect.Type, bool) {
	n := t.NumMethods()
	if n == 0 {
		return tyEmptyInterface, false
	}
	var nested bool
	ms := make([]reflect.Method, n)
	for i := 0; i < n; i++ {
		fn := t.Method(i)
		mtyp, n := r.ToType(fn.Type())
		if n {
			nested = true
		}
		ms[i] = reflect.Method{
			Name: fn.Name(),
			Type: mtyp,
		}
		if pkg := fn.Pkg(); pkg != nil {
			ms[i].PkgPath = pkg.Path()
		}
	}
	return r.rctx.InterfaceOf(nil, ms), nested
}

func (r *TypesRecord) toNamedType(t *types.Named) (reflect.Type, bool) {
	ut := t.Underlying()
	pkgpath, name, typeargs, nested := r.extractNamed(t, false)
	if pkgpath == "" {
		if name == "error" {
			return tyErrorInterface, false
		}
		return r.ToType(ut)
	}
	methods := typeutil.IntuitiveMethodSet(t, nil)
	hasMethod := len(methods) > 0
	etyp := toMockType(ut)
	typ := reflectx.NamedTypeOf(pkgpath, name, etyp)
	if hasMethod {
		var mcount, pcount int
		for i := 0; i < len(methods); i++ {
			sig := methods[i].Type().(*types.Signature)
			if !isPointer(sig.Recv().Type()) {
				mcount++
			}
			pcount++
		}
		typ = r.rctx.NewMethodSet(typ, mcount, pcount)
	}
	r.saveType(t, typ, nested)
	var utype reflect.Type
	if st, ok := ut.(*types.Struct); ok {
		utype, _ = r.toStructType(st, false)
	} else {
		utype, _ = r.ToType(ut)
	}
	reflectx.SetUnderlying(typ, utype)
	if typeargs {
		pkgpath, name, _, _ = r.extractNamed(t, true)
		reflectx.SetTypeName(typ, pkgpath, name)
	}
	if hasMethod && typ.Kind() != reflect.Interface {
		r.setMethods(typ, methods)
	}
	return typ, nested
}

func (r *TypesRecord) toStructType(t *types.Struct, mset bool) (reflect.Type, bool) {
	n := t.NumFields()
	if n == 0 {
		return tyEmptyStruct, false
	}
	var nested bool
	flds := make([]reflect.StructField, n)
	for i := 0; i < n; i++ {
		f := t.Field(i)
		typ, n := r.ToType(f.Type())
		if n {
			nested = true
		}
		flds[i] = r.toStructField(f, typ, t.Tag(i))
	}
	typ := r.rctx.StructOf(flds)
	if mset {
		methods := typeutil.IntuitiveMethodSet(t, nil)
		if numMethods := len(methods); numMethods != 0 {
			// anonymous structs with methods. struct { T }
			var mcount, pcount int
			for i := 0; i < numMethods; i++ {
				sig := methods[i].Type().(*types.Signature)
				if !isPointer(sig.Recv().Type()) {
					mcount++
				}
				pcount++
			}
			typ = r.rctx.NewMethodSet(typ, mcount, pcount)
			r.setMethods(typ, methods)
		}
	}
	return typ, nested
}

func (r *TypesRecord) toStructField(v *types.Var, typ reflect.Type, tag string) reflect.StructField {
	name := v.Name()
	fld := reflect.StructField{
		Name:      name,
		Type:      typ,
		Tag:       reflect.StructTag(tag),
		Anonymous: v.Anonymous(),
	}
	if !token.IsExported(name) {
		if pkg := v.Pkg(); pkg != nil {
			fld.PkgPath = pkg.Path()
		} else {
			// fix golang.org/x/tools@v0.29 or higher invalid prog.RuntimeTypes struct{in string; out string}
			fld.PkgPath = "_"
		}
	}
	return fld
}

func (r *TypesRecord) ToTypeList(tuple *types.Tuple) ([]reflect.Type, bool) {
	n := tuple.Len()
	if n == 0 {
		return nil, false
	}
	var nested bool
	var ne bool
	list := make([]reflect.Type, n)
	for i := 0; i < n; i++ {
		list[i], ne = r.ToType(tuple.At(i).Type())
		if ne {
			nested = true
		}
	}
	return list, nested
}

func isPointer(typ types.Type) bool {
	_, ok := typ.Underlying().(*types.Pointer)
	return ok
}

func embedFunc(typ reflect.Type, name string, idx []int, pointer bool, indirect bool, variadic bool) func([]reflect.Value) []reflect.Value {
	switch kind := typ.Kind(); kind {
	case reflect.Interface:
		if variadic {
			return func(args []reflect.Value) []reflect.Value {
				v := reflectx.FieldByIndexX(args[0], idx[:len(idx)-1])
				return v.MethodByName(name).CallSlice(args[1:])
			}
		}
		return func(args []reflect.Value) []reflect.Value {
			v := reflectx.FieldByIndexX(args[0], idx[:len(idx)-1])
			return v.MethodByName(name).Call(args[1:])
		}
	case reflect.Ptr:
		if pointer {
			m, _ := reflectx.MethodByName(typ, name)
			if variadic {
				return func(args []reflect.Value) []reflect.Value {
					args[0] = reflectx.FieldByIndexX(args[0].Elem(), idx[:len(idx)-1]).Addr()
					return m.Func.CallSlice(args)
				}
			}
			return func(args []reflect.Value) []reflect.Value {
				args[0] = reflectx.FieldByIndexX(args[0].Elem(), idx[:len(idx)-1]).Addr()
				return m.Func.Call(args)
			}
		}
		fallthrough
	default:
		if indirect && kind != reflect.Ptr {
			typ = reflect.PointerTo(typ)
		}
		m, _ := reflectx.MethodByName(typ, name)
		if variadic {
			return func(args []reflect.Value) []reflect.Value {
				args[0] = reflectx.FieldByIndexX(args[0], idx[:len(idx)-1])
				return m.Func.CallSlice(args)
			}
		}
		return func(args []reflect.Value) []reflect.Value {
			args[0] = reflectx.FieldByIndexX(args[0], idx[:len(idx)-1])
			return m.Func.Call(args)
		}
	}
}

func (r *TypesRecord) setMethods(typ reflect.Type, methods []*types.Selection) {
	numMethods := len(methods)
	var ms []reflectx.Method
	for i := 0; i < numMethods; i++ {
		fn := methods[i].Obj().(*types.Func)
		sig := methods[i].Type().(*types.Signature)
		pointer := isPointer(sig.Recv().Type())
		mtyp, _ := r.ToType(sig)
		var mfn func(args []reflect.Value) []reflect.Value
		idx := methods[i].Index()
		var mid int
		if len(idx) > 1 {
			indirect := methods[i].Indirect()
			rtyp := fn.Type().Underlying().(*types.Signature).Recv().Type()
			rt, _ := r.ToType(rtyp)
			key := mfnKey{typ: rt, name: fn.Name(), indexs: indexsToString(idx), pointer: pointer, indirect: indirect}
			if v, ok := r.ctx.mfnmap[key]; ok {
				mfn = v.fn
				mid = v.id
			} else {
				variadic := mtyp.IsVariadic()
				r.ctx.mfnid++
				mid = r.ctx.mfnid
				mfn = embedFunc(rt, fn.Name(), idx, pointer, indirect, variadic)
				r.ctx.mfnmap[key] = mfnValue{fn: mfn, id: mid}
			}
		} else {
			mfn = r.finder.FindMethod(mtyp, fn)
		}
		var pkgpath string
		if pkg := fn.Pkg(); pkg != nil {
			pkgpath = pkg.Path()
		}
		method := reflectx.MakeMethod(fn.Name(), pkgpath, pointer, mtyp, mfn)
		method.FuncId = mid
		ms = append(ms, method)
	}
	err := r.rctx.SetMethodSet(typ, ms, false)
	if err != nil {
		log.Fatalf("SetMethodSet %v err, %v\n", typ, err)
	}
}

func toReflectChanDir(d types.ChanDir) reflect.ChanDir {
	switch d {
	case types.SendRecv:
		return reflect.BothDir
	case types.SendOnly:
		return reflect.SendDir
	case types.RecvOnly:
		return reflect.RecvDir
	}
	return 0
}

func (r *TypesRecord) Load(pkg *ssa.Package) {
	checked := make(map[types.Type]bool)
	for _, v := range pkg.Members {
		typ := v.Type()
		if checked[typ] {
			continue
		}
		checked[typ] = true
		if hasTypeParam(typ) {
			continue
		}
		r.ToType(typ)
	}
}
