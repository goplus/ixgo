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

package repl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sort"
	"strings"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/internal/constant"
)

// parseSymbol breaks str apart into a pkg, symbol and method
// fmt.Println => fmt Println
// fmt.Stringer.String => fmt Stringer.String
func parseSymbol(str string) (pkg, symbol string, method string, err error) {
	if str == "" {
		return
	}
	elem := strings.Split(str, ".")
	switch len(elem) {
	case 1:
	case 2:
		symbol = elem[1]
	case 3:
		symbol = elem[1]
		method = elem[2]
	default:
		err = errors.New("too many periods in symbol specification")
		return
	}
	pkg = elem[0]
	return
}

func findPkg(name string) (pkg string, found bool) {
	list := ixgo.PackageList()
	for _, v := range list {
		if name == v {
			return v, true
		}
	}
	for _, v := range list {
		if strings.HasSuffix(v, "/"+name) {
			return v, true
		}
	}
	return
}

func lookupSymbol(p *ixgo.Package, sym string) (info string, found bool) {
	if v, ok := p.UntypedConsts[sym]; ok {
		return fmt.Sprintf("const %v.%v %v = %v", p.Name, sym, v.Typ, constant.ExactConstant(v.Value)), true
	}
	if v, ok := p.TypedConsts[sym]; ok {
		return fmt.Sprintf("const %v.%v %v = %v", p.Name, sym, v.Typ, constant.ExactConstant(v.Value)), true
	}
	if v, ok := p.Vars[sym]; ok {
		return fmt.Sprintf("var %v.%v %v", p.Name, sym, v.Type().Elem()), true
	}
	if v, ok := p.Funcs[sym]; ok {
		var buf bytes.Buffer
		writeFunc(&buf, p.Name+"."+sym, v.Type())
		return buf.String(), true
	}
	if t, ok := p.NamedTypes[sym]; ok {
		var buf bytes.Buffer
		if t.Kind() == reflect.Struct {
			writeStruct(&buf, p.Name+"."+sym, t)
			buf.WriteByte('\n')
		} else {
			fmt.Fprintf(&buf, "type %v.%v %v\n", p.Name, sym, t.Kind())
		}
		n := t.NumMethod()
		skip := make(map[string]bool)
		for i := 0; i < n; i++ {
			m := t.Method(i)
			skip[m.Name] = true
			writeMethod(&buf, m)
			buf.WriteByte('\n')
		}
		ptyp := reflect.PtrTo(t)
		n = ptyp.NumMethod()
		for i := 0; i < n; i++ {
			m := ptyp.Method(i)
			if skip[m.Name] {
				continue
			}
			writeMethod(&buf, m)
			buf.WriteByte('\n')
		}
		return buf.String(), true
	}
	if t, ok := p.Interfaces[sym]; ok {
		var buf bytes.Buffer
		writeInterface(&buf, p.Name+"."+sym, t)
		buf.WriteByte('\n')
		return buf.String(), true
	}
	return
}

func dumpPkg(p *ixgo.Package) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "package %v // import %q\n\n", p.Name, p.Path)
	// untyped const
	var uconst []string
	for v := range p.UntypedConsts {
		uconst = append(uconst, v)
	}
	sort.Strings(uconst)
	for _, v := range uconst {
		t := p.UntypedConsts[v]
		fmt.Fprintf(&buf, "const %v = %v\n", v, constant.ExactConstant(t.Value))
	}
	// typed const
	var tconst []string
	for v := range p.TypedConsts {
		tconst = append(tconst, v)
	}
	sort.Strings(tconst)
	for _, v := range tconst {
		t := p.TypedConsts[v]
		fmt.Fprintf(&buf, "const %v %v = %v\n", v, t.Typ, constant.ExactConstant(t.Value))
	}
	// var
	var vars []string
	for v := range p.Vars {
		vars = append(vars, v)
	}
	sort.Strings(vars)
	for _, v := range vars {
		t := p.Vars[v]
		fmt.Fprintf(&buf, "var %v %v\n", v, t.Elem().Type())
	}
	// funcs
	var funcs []string
	for v := range p.Funcs {
		funcs = append(funcs, v)
	}
	sort.Strings(funcs)
	for _, v := range funcs {
		f := p.Funcs[v]
		writeFunc(&buf, v, f.Type())
		buf.WriteByte('\n')
	}
	// named types
	var types []string
	for v := range p.NamedTypes {
		types = append(types, v)
	}
	sort.Strings(types)
	for _, v := range types {
		t := p.NamedTypes[v]
		fmt.Fprintf(&buf, "type %v %v\n", v, t.Kind())
		skip := make(map[string]bool)
		n := t.NumMethod()
		for i := 0; i < n; i++ {
			m := t.Method(i)
			skip[m.Name] = true
			buf.WriteString("    ")
			writeMethod(&buf, m)
			buf.WriteByte('\n')
		}
		ptyp := reflect.PtrTo(t)
		n = ptyp.NumMethod()
		for i := 0; i < n; i++ {
			m := ptyp.Method(i)
			if skip[m.Name] {
				continue
			}
			buf.WriteString("    ")
			writeMethod(&buf, m)
			buf.WriteByte('\n')
		}
	}
	// interface
	var ifaces []string
	for v := range p.Interfaces {
		ifaces = append(ifaces, v)
	}
	sort.Strings(ifaces)
	for _, v := range ifaces {
		t := p.Interfaces[v]
		fmt.Fprintf(&buf, "type %v interface\n", v)
		n := t.NumMethod()
		for i := 0; i < n; i++ {
			m := t.Method(i)
			buf.WriteString("    ")
			writeFunc(&buf, m.Name, m.Type)
			buf.WriteByte('\n')
		}
	}
	return buf.String()
}

func writeInterface(w io.Writer, name string, typ reflect.Type) {
	n := typ.NumMethod()
	if n == 0 {
		fmt.Fprintf(w, "type %v interface{}", name)
		return
	}
	fmt.Fprintf(w, "type %v interface{\n", name)
	for i := 0; i < n; i++ {
		m := typ.Method(i)
		w.Write([]byte("    "))
		writeFunc(w, m.Name, m.Type)
		w.Write([]byte{'\n'})
	}
	w.Write([]byte{'}'})
}

func writeStruct(w io.Writer, name string, typ reflect.Type) {
	n := typ.NumField()
	if n == 0 {
		fmt.Fprintf(w, "type %v struct{}", name)
		return
	}
	fmt.Fprintf(w, "type %v struct{\n", name)
	for i := 0; i < n; i++ {
		f := typ.Field(i)
		if f.PkgPath != "" {
			fmt.Fprintf(w, "    // ... other fields elided ...\n")
			break
		}
		if f.Anonymous {
			fmt.Fprintf(w, "    %v\n", f.Type)
		} else {
			fmt.Fprintf(w, "    %v %v\n", f.Name, f.Type)
		}
	}
	fmt.Fprintf(w, "}")
}

func writeMethod(w io.Writer, m reflect.Method) {
	typ := m.Type
	numIn := typ.NumIn()
	numOut := typ.NumOut()
	var ins []string
	if typ.IsVariadic() {
		for i := 1; i < numIn-1; i++ {
			ins = append(ins, typ.In(i).String())
		}
		ins = append(ins, "..."+typ.In(numIn-1).Elem().String())
	} else {
		for i := 1; i < numIn; i++ {
			ins = append(ins, typ.In(i).String())
		}
	}
	switch numOut {
	case 0:
		fmt.Fprintf(w, "func (%v) %v(%v)", typ.In(0), m.Name, strings.Join(ins, ", "))
	case 1:
		fmt.Fprintf(w, "func (%v) %v(%v) %v", typ.In(0), m.Name, strings.Join(ins, ", "), typ.Out(0).String())
	default:
		var outs []string
		for i := 0; i < numOut; i++ {
			outs = append(outs, typ.Out(i).String())
		}
		fmt.Fprintf(w, "func (%v) %v(%v) (%v)", typ.In(0), m.Name, strings.Join(ins, ", "), strings.Join(outs, ", "))
	}
}

func writeFunc(w io.Writer, name string, typ reflect.Type) {
	numIn := typ.NumIn()
	numOut := typ.NumOut()
	var ins []string
	if typ.IsVariadic() {
		for i := 0; i < numIn-1; i++ {
			ins = append(ins, typ.In(i).String())
		}
		ins = append(ins, "..."+typ.In(numIn-1).Elem().String())
	} else {
		for i := 0; i < numIn; i++ {
			ins = append(ins, typ.In(i).String())
		}
	}
	switch numOut {
	case 0:
		fmt.Fprintf(w, "func %v(%v)", name, strings.Join(ins, ", "))
	case 1:
		fmt.Fprintf(w, "func %v(%v) %v", name, strings.Join(ins, ", "), typ.Out(0).String())
	default:
		var outs []string
		for i := 0; i < numOut; i++ {
			outs = append(outs, typ.Out(i).String())
		}
		fmt.Fprintf(w, "func %v(%v) (%v)", name, strings.Join(ins, ", "), strings.Join(outs, ", "))
	}
}
