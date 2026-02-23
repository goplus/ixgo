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
	"bytes"
	"fmt"
	"reflect"
	"unsafe"
)

// writevalue for write argument pass to print
func writevalue(buf *bytes.Buffer, v interface{}, enableAny bool) {
	switch v := v.(type) {
	case float64:
		writefloat64(buf, v)
	case float32:
		writefloat32(buf, v)
	case complex128:
		writecomplex128(buf, v)
	case complex64:
		writecomplex64(buf, v)
	case nil, bool, int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64, uintptr,
		string:
		fmt.Fprintf(buf, "%v", v)
	default:
		i := reflect.ValueOf(v)
		switch i.Kind() {
		case reflect.Float32:
			writefloat32(buf, float32(i.Float()))
		case reflect.Float64:
			writefloat64(buf, i.Float())
		case reflect.Complex64:
			writecomplex64(buf, complex64(i.Complex()))
		case reflect.Complex128:
			writecomplex128(buf, i.Complex())
		case reflect.Map, reflect.Ptr, reflect.Func, reflect.Chan, reflect.UnsafePointer:
			fmt.Fprintf(buf, "%p", v)
		case reflect.Slice:
			fmt.Fprintf(buf, "[%v/%v]%p", i.Len(), i.Cap(), v)
		case reflect.String:
			fmt.Fprintf(buf, "%v", v)
		case reflect.Interface:
			eface := *(*emptyInterface)(unsafe.Pointer(&i))
			fmt.Fprintf(buf, "(%p,%p)", eface.typ, eface.word)
		case reflect.Struct, reflect.Array:
			if enableAny {
				fmt.Fprintf(buf, "%v", v)
			} else {
				panic(fmt.Sprintf("illegal types for operand: print\n\t%T", v))
			}
		default:
			fmt.Fprintf(buf, "%v", v)
		}
	}
}

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  unsafe.Pointer
	word unsafe.Pointer
}

func writeinterface(out *bytes.Buffer, i interface{}) {
	eface := *(*emptyInterface)(unsafe.Pointer(&i))
	fmt.Fprintf(out, "(%p,%p)", eface.typ, eface.word)
}

// writeany for write argument pass to panic
func writeany(buf *bytes.Buffer, v interface{}) {
	switch v := v.(type) {
	case nil, bool, int, int8, int16, int32, int64,
		float32, float64, complex64, complex128,
		uint, uint8, uint16, uint32, uint64, uintptr,
		string:
		writevalue(buf, v, true)
	case error:
		fmt.Fprintf(buf, "%v", v)
	default:
		i := reflect.ValueOf(v)
		typ := i.Type()
		switch i.Kind() {
		case reflect.String:
			fmt.Fprintf(buf, "%v(%q)", typ, i)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			fmt.Fprintf(buf, "%v(%v)", typ, i)
		case reflect.Float32:
			buf.WriteString(typ.String())
			buf.WriteByte('(')
			writefloat32(buf, float32(i.Float()))
			buf.WriteByte(')')
		case reflect.Float64:
			buf.WriteString(typ.String())
			buf.WriteByte('(')
			writefloat64(buf, i.Float())
			buf.WriteByte(')')
		case reflect.Complex64:
			buf.WriteString(typ.String())
			buf.WriteByte('(')
			writecomplex64(buf, complex64(i.Complex()))
			buf.WriteByte(')')
		case reflect.Complex128:
			buf.WriteString(typ.String())
			buf.WriteByte('(')
			writecomplex128(buf, i.Complex())
			buf.WriteByte(')')
		default:
			eface := *(*emptyInterface)(unsafe.Pointer(&i))
			fmt.Fprintf(buf, "(%v) %p", typ, eface.word)
		}
	}
}
