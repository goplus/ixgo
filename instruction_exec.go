//go:build instruction

package ixgo

import (
	"bytes"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/visualfc/xtype"
)

// makeInstruction binds execution metadata to the supplied runtime context.
// It returns nil when the description has no executor implementation.
// Unsupported instructions retain their existing lowering.
func makeInstruction(interp *Interp, d typedInstruction) func(*frame) {
	switch d.Op {
	case opMakeClosure:
		aux := d.Aux.(closureAux)
		if aux.Body == nil {
			return nil
		}
		body, signature, bindings, result := aux.Body, aux.Signature, aux.Bindings, d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		return func(fr *frame) {
			var env []value
			for _, binding := range bindings {
				env = append(env, fr.reg(binding))
			}
			fr.setReg(result, body.makeFunction(signature, env).Interface())
		}
	case opGo, opDefer:
		var desc callDescription
		var stack register
		switch aux := d.Aux.(type) {
		case closureCallAux:
			if aux.Body == nil {
				return nil
			}
			args, bindings := aux.Arguments, aux.Bindings
			stack = aux.DeferStack
			if d.RuntimeBound {
				args = aux.runtime.Arguments
				bindings = args[len(aux.Arguments):]
				stack = aux.runtime.DeferStack
			}
			desc = callDescription{closure: aux.Body, bindings: bindings, args: args[:len(args)-len(bindings)]}
		case invokeCallAux:
			target, args := aux.Target, aux.Arguments
			stack = aux.DeferStack
			if d.RuntimeBound {
				target, args, stack = aux.runtime.Target, aux.runtime.Arguments, aux.runtime.DeferStack
			}
			if target < 0 || aux.Method == "" {
				return nil
			}
			desc = callDescription{receiver: target, method: aux.Method, args: args}
		case dynamicCallAux:
			target, args := aux.Target, aux.Arguments
			stack = aux.DeferStack
			if d.RuntimeBound {
				target, args, stack = aux.runtime.Target, aux.runtime.Arguments, aux.runtime.DeferStack
			}
			if target < 0 {
				return nil
			}
			desc = callDescription{receiver: target, args: args}
		case bodyCallAux:
			args := aux.Arguments
			stack = aux.DeferStack
			if d.RuntimeBound {
				args, stack = aux.runtime.Arguments, aux.runtime.DeferStack
			}
			if aux.Body == nil {
				return nil
			}
			desc = callDescription{target: aux.Body, args: args}
		case externalCallAux:
			args := aux.Arguments
			stack = aux.DeferStack
			if d.RuntimeBound {
				args, stack = aux.runtime.Arguments, aux.runtime.DeferStack
			}
			if !aux.External.IsValid() {
				return nil
			}
			desc = callDescription{target: aux.External, args: args}
		default:
			return nil
		}
		if d.Op == opGo {
			return instructionGoCall(interp, desc)
		}
		return instructionDeferCall(interp, desc, stack)
	case opCall:
		switch aux := d.Aux.(type) {
		case invokeCallAux:
			if aux.Target < 0 || aux.Method == "" {
				return nil
			}
			return makeInvokeInstruction(interp, d, aux)
		case builtinCallAux:
			return makeBuiltinInstruction(interp, d, aux)
		case dynamicCallAux:
			target, result, args, dests := aux.Target, d.Result, aux.Arguments, aux.ExtractResults
			if d.RuntimeBound {
				target, result, args, dests = aux.runtime.Target, d.RuntimeResult, aux.runtime.Arguments, aux.runtime.ExtractResults
			}
			if target < 0 {
				return nil
			}
			if supportFuncVal && interp.ctx.Mode&DisableDynamicFuncCallAnalysis == 0 {
				return dynamicFunCall(interp, target, result, args, dests)
			}
			return func(fr *frame) {
				interp.callExternalByStack(fr, reflect.ValueOf(fr.reg(target)), result, args, dests)
			}
		case directCallAux:
			if aux.Adapter == nil {
				return nil
			}
			args, result := aux.Arguments, d.Result
			if d.RuntimeBound {
				args, result = aux.runtime.Arguments, d.RuntimeResult
			}
			adapter := aux.Adapter
			return func(fr *frame) {
				interp.invokeDirectCall(fr, adapter, result, args)
			}
		case externalCallAux:
			ext, args, result, dests := aux.External, aux.Arguments, d.Result, aux.ExtractResults
			if d.RuntimeBound {
				args, result, dests = aux.runtime.Arguments, d.RuntimeResult, aux.runtime.ExtractResults
			}
			if !ext.IsValid() || ext.Kind() != reflect.Func {
				return nil
			}
			if typ := ext.Type(); typ.NumIn() > 0 && typ.In(0) == typFramePtr {
				return func(fr *frame) { interp.callExternalWithFrameByStack(fr, ext, result, args, dests) }
			}
			return func(fr *frame) { interp.callExternalByStack(fr, ext, result, args, dests) }
		case closureCallAux:
			result, args, dests := d.Result, append(append([]register(nil), aux.Arguments...), aux.Bindings...), aux.ExtractResults
			if d.RuntimeBound {
				result, args, dests = d.RuntimeResult, aux.runtime.Arguments, aux.runtime.ExtractResults
			}
			return makeBodyCallInstruction(interp, aux.Body, result, args, dests)
		case bodyCallAux:
			result, args, dests := d.Result, aux.Arguments, aux.ExtractResults
			if d.RuntimeBound {
				result, args, dests = d.RuntimeResult, aux.runtime.Arguments, aux.runtime.ExtractResults
			}
			return makeBodyCallInstruction(interp, aux.Body, result, args, dests)
		}
		return nil
	case opJump:
		target := d.Aux.(jumpAux).Target
		return func(fr *frame) { fr.jump(target) }
	case opIf:
		aux := d.Aux.(ifAux)
		if aux.Constant {
			if aux.Condition {
				return func(fr *frame) { fr.jump(aux.Targets[0]) }
			}
			return func(fr *frame) { fr.jump(aux.Targets[1]) }
		}
		condition := d.Args[0]
		if d.RuntimeBound && len(d.RuntimeArgs) > 0 {
			condition = d.RuntimeArgs[0]
		}
		if aux.BasicBool {
			return func(fr *frame) {
				if fr.reg(condition).(bool) {
					fr.jump(aux.Targets[0])
				} else {
					fr.jump(aux.Targets[1])
				}
			}
		}
		return func(fr *frame) {
			if fr.bool(condition) {
				fr.jump(aux.Targets[0])
			} else {
				fr.jump(aux.Targets[1])
			}
		}
	case opField, opFieldAddr:
		if d.Op == opFieldAddr && interp.ctx.Mode&EnableCachedReg != 0 {
			return nil
		}
		aux := d.Aux.(fieldAux)
		source := aux.RuntimeSource
		if !d.RuntimeBound && len(d.Args) > 0 {
			source = d.Args[0]
		}
		result := d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		return func(fr *frame) {
			var v value
			var err error
			if d.Op == opFieldAddr {
				v, err = fieldAddrAt(fr.reg(source), aux.Offset, aux.Field)
			} else {
				v, err = fieldX(fr.reg(source), aux.Index)
			}
			if err != nil {
				panic(fr.runtimeError(fr.pfn.InstrForPC(int(fr.ipc)-1), err.Error()))
			}
			fr.setReg(result, v)
		}
	case opIndex, opIndexAddr:
		aux := d.Aux.(indexAux)
		source, index := aux.RuntimeSource, aux.RuntimeIndex
		if !d.RuntimeBound && len(d.Args) >= 2 {
			source, index = d.Args[0], d.Args[1]
		}
		result := d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		return func(fr *frame) {
			v := reflect.ValueOf(fr.reg(source))
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			idx := asBound(fr.reg(index))
			if !idx.ok || idx.n >= v.Len() {
				panicIndexBound(fr, fr.pfn.InstrForPC(int(fr.ipc)-1), idx, v.Len())
			}
			if d.Op == opIndexAddr {
				fr.setReg(result, v.Index(idx.n).Addr().Interface())
			} else {
				fr.setReg(result, v.Index(idx.n).Interface())
			}
		}
	case opLookup:
		aux := d.Aux.(lookupAux)
		source, key := aux.RuntimeSource, aux.RuntimeKey
		if !d.RuntimeBound && len(d.Args) >= 2 {
			source, key = d.Args[0], d.Args[1]
		}
		result := d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		return func(fr *frame) {
			container, k := fr.reg(source), fr.reg(key)
			rv := reflect.ValueOf(container)
			if rv.Kind() == reflect.String {
				fr.setReg(result, rv.String()[asInt(k)])
				return
			}
			mv := rv.MapIndex(mapKeyValue(rv, k))
			ok := mv.IsValid()
			var value value
			if ok {
				value = mv.Interface()
			} else {
				value = reflect.Zero(rv.Type().Elem()).Interface()
			}
			if aux.CommaOk {
				fr.setReg(result, tuple{value, ok})
			} else {
				fr.setReg(result, value)
			}
		}
	case opSlice:
		aux := d.Aux.(sliceAux)
		x, low, high, max := aux.RuntimeSource, aux.RuntimeLow, aux.RuntimeHigh, aux.RuntimeMax
		if !d.RuntimeBound {
			x, low, high, max = -1, -1, -1, -1
			if len(d.Args) > 0 {
				x = d.Args[0]
			}
			pos := 1
			if aux.Low && pos < len(d.Args) {
				low = d.Args[pos]
				pos++
			}
			if aux.High && pos < len(d.Args) {
				high = d.Args[pos]
				pos++
			}
			if aux.Max && pos < len(d.Args) {
				max = d.Args[pos]
			}
		}
		// Static slice bounds still require the stack lowering path until
		// their constant values have an explicit bytecode representation.
		if x < 0 || (aux.Low && low < 0) || (aux.High && high < 0) || (aux.Max && max < 0) {
			return nil
		}
		result := d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		isNamed := aux.Target.Kind() == reflect.Slice && aux.Target != reflect.SliceOf(aux.Target.Elem())
		return func(fr *frame) {
			v := slice(fr, fr.pfn.InstrForPC(int(fr.ipc)-1), sliceOperands{low: aux.Low, high: aux.High, max: aux.Max}, aux.MakeSliceCheck, x, high, low, max)
			if isNamed {
				v = v.Convert(aux.Target)
			}
			fr.setReg(result, v.Interface())
		}
	case opChangeType:
		aux := d.Aux.(conversionAux)
		source := aux.RuntimeSource
		if !d.RuntimeBound && len(d.Args) > 0 {
			source = d.Args[0]
		}
		if source < 0 || aux.Source == nil || aux.Target == nil {
			return nil
		}
		// ChangeType preserves representation; unsafe and interface changes use fallback.
		if aux.Source.Kind() == reflect.Interface || aux.Target.Kind() == reflect.Interface || aux.Source.Kind() == reflect.UnsafePointer || aux.Target.Kind() == reflect.UnsafePointer {
			return nil
		}
		result := d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		return func(fr *frame) {
			v := fr.reg(source)
			if v == nil {
				fr.setReg(result, nil)
				return
			}
			rv := reflect.ValueOf(v)
			if rv.Type().ConvertibleTo(aux.Target) {
				fr.setReg(result, rv.Convert(aux.Target).Interface())
			} else {
				fr.setReg(result, v)
			}
		}
	case opConvert:
		aux := d.Aux.(conversionAux)
		if aux.Source != nil && aux.Target != nil && (aux.Source.Kind() == reflect.UnsafePointer || aux.Target.Kind() == reflect.UnsafePointer) {
			source := aux.RuntimeSource
			if !d.RuntimeBound && len(d.Args) > 0 {
				source = d.Args[0]
			}
			if source < 0 {
				return nil
			}
			result := d.Result
			if d.RuntimeBound {
				result = d.RuntimeResult
			}
			sourceKind, targetKind := aux.Source.Kind(), aux.Target.Kind()
			targetType := xtype.TypeOfType(aux.Target)
			return func(fr *frame) {
				v := fr.reg(source)
				switch {
				case targetKind == reflect.UnsafePointer && sourceKind == reflect.Uintptr:
					fr.setReg(result, xtype.ConvertPtr(targetType, unsafe.Pointer(uintptr(fr.uintptr(source)))))
				case targetKind == reflect.UnsafePointer && sourceKind == reflect.Ptr:
					fr.setReg(result, xtype.ConvertPtr(targetType, v))
				case targetKind == reflect.Uintptr && sourceKind == reflect.UnsafePointer:
					fr.setReg(result, xtype.MakeUintptr(targetType, uintptr(fr.pointer(source))))
				case targetKind == reflect.Ptr && sourceKind == reflect.UnsafePointer:
					fr.setReg(result, xtype.Make(targetType, v))
				default:
					return
				}
			}
		}
		source := aux.RuntimeSource
		if !d.RuntimeBound && len(d.Args) > 0 {
			source = d.Args[0]
		}
		if source < 0 || aux.Source == nil || aux.Target == nil {
			return nil
		}
		if aux.Source.Kind() == reflect.UnsafePointer || aux.Target.Kind() == reflect.UnsafePointer {
			return nil
		}
		result := d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		return func(fr *frame) {
			v := fr.reg(source)
			if v == nil {
				fr.setReg(result, nil)
				return
			}
			rv := reflect.ValueOf(v)
			if !rv.Type().ConvertibleTo(aux.Target) {
				return
			}
			fr.setReg(result, rv.Convert(aux.Target).Interface())
		}
	case opSliceToArrayPointer:
		aux := d.Aux.(conversionAux)
		source := aux.RuntimeSource
		if !d.RuntimeBound && len(d.Args) > 0 {
			source = d.Args[0]
		}
		if source < 0 || aux.Target == nil || aux.Target.Kind() != reflect.Pointer || aux.Target.Elem().Kind() != reflect.Array {
			return nil
		}
		result := d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		return func(fr *frame) {
			rv := reflect.ValueOf(fr.reg(source))
			if rv.Len() < aux.Target.Elem().Len() {
				panic(fr.runtimeError(fr.pfn.InstrForPC(int(fr.ipc)-1), fmt.Sprintf(errSliceToArrayPointer, rv.Len(), aux.Target.Elem().Len())))
			}
			fr.setReg(result, rv.Convert(aux.Target).Interface())
		}
	case opChangeInterface, opMakeInterface:
		aux := d.Aux.(conversionAux)
		source := aux.RuntimeSource
		if !d.RuntimeBound && len(d.Args) > 0 {
			source = d.Args[0]
		}
		if source < 0 {
			return nil
		}
		result := d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		return func(fr *frame) { fr.setReg(result, fr.reg(source)) }
	case opTypeAssert:
		aux := d.Aux.(typeAssertAux)
		source := aux.RuntimeSource
		if !d.RuntimeBound && len(d.Args) > 0 {
			source = d.Args[0]
		}
		if source < 0 || aux.Target == nil {
			return nil
		}
		result := d.Result
		if d.RuntimeBound {
			result = d.RuntimeResult
		}
		operands := typeAssertOperands{typ: aux.Target, sourceType: aux.Source, assertedType: aux.AssertedType, commaOk: aux.CommaOk}
		return func(fr *frame) {
			fr.setReg(result, typeAssert(fr, fr.pfn.InstrForPC(int(fr.ipc)-1), operands, fr.reg(source)))
		}
	case opReturn:
		aux := d.Aux.(returnAux)
		args := d.Args
		if d.RuntimeBound && aux.RuntimeResults != nil {
			args = aux.RuntimeResults
		}
		return func(fr *frame) {
			for i, r := range args {
				if i >= aux.Results {
					break
				}
				fr.stack[i] = fr.reg(r)
			}
			fr.ipc = -1
		}
	case opPanic:
		aux := d.Aux.(panicAux)
		r := aux.RuntimeValue
		if !d.RuntimeBound && len(d.Args) > 0 {
			r = d.Args[0]
		}
		return func(fr *frame) {
			var err error = PanicError{stack: debugStack(fr), Value: fr.reg(r)}
			if interp.ctx.panicFunc != nil {
				err = interp.ctx.handlePanic(fr, fr.pfn.InstrForPC(int(fr.ipc)-1), err)
			}
			panic(err)
		}
	case opSend:
		aux := d.Aux.(sendAux)
		ic, iv := aux.RuntimeChan, aux.RuntimeValue
		if !d.RuntimeBound && len(d.Args) >= 2 {
			ic, iv = d.Args[0], d.Args[1]
		}
		return func(fr *frame) {
			c, x := fr.reg(ic), fr.reg(iv)
			ch := reflect.ValueOf(c)
			if interp.gcEnabled {
				fr.gcFrames()
			}
			if x == nil {
				ch.Send(reflect.New(ch.Type().Elem()).Elem())
			} else {
				ch.Send(reflect.ValueOf(x))
			}
		}
	case opStore:
		aux := d.Aux.(storeAux)
		ia, iv := aux.RuntimeAddress, aux.RuntimeValue
		if !d.RuntimeBound && len(d.Args) >= 2 {
			ia, iv = d.Args[0], d.Args[1]
		}
		return func(fr *frame) {
			x := reflect.ValueOf(fr.reg(ia))
			v := reflect.ValueOf(fr.reg(iv))
			if v.IsValid() {
				SetValue(x.Elem(), v)
			} else {
				SetValue(x.Elem(), reflect.Zero(x.Elem().Type()))
			}
		}
	case opPhi:
		aux := d.Aux.(phiAux)
		ir, ie, preds := aux.Results, aux.Edges, aux.Preds
		if d.RuntimeBound && aux.RuntimeResults != nil {
			ir, ie = aux.RuntimeResults, aux.RuntimeEdges
		}
		if len(ir) == 0 || len(ie) == 0 {
			return func(*frame) {}
		}
		if len(ir) == 1 {
			result, edges := ir[0], ie[0]
			return func(fr *frame) {
				for i, pred := range preds {
					if fr.pred == pred {
						fr.setReg(result, fr.reg(edges[i]))
						return
					}
				}
			}
		}
		if phiDirectAssignment(ir, ie) {
			return func(fr *frame) {
				for j, pred := range preds {
					if fr.pred == pred {
						for i, result := range ir {
							fr.setReg(result, fr.reg(ie[i][j]))
						}
						return
					}
				}
			}
		}
		return func(fr *frame) {
			predIndex := -1
			for i, pred := range preds {
				if fr.pred == pred {
					predIndex = i
					break
				}
			}
			if predIndex < 0 {
				return
			}
			if cap(fr.phiVals) < len(ir) {
				fr.phiVals = make([]value, len(ir))
			}
			// Read all edges before writing results to preserve cyclic assignments.
			values := fr.phiVals[:len(ir)]
			for i := range ir {
				values[i] = fr.reg(ie[i][predIndex])
			}
			for i := range ir {
				fr.setReg(ir[i], values[i])
			}
		}
	case opBinOp:
		return makeBinaryInstruction(d)
	case opUnOp:
		return makeUnaryInstruction(d)
	}
	return nil
}

// instructionGoCall builds the bytecode-only Go executor. Lowering keeps its
// own implementation in opblock.go.
func instructionGoCall(interp *Interp, desc callDescription) func(*frame) {
	return func(fr *frame) {
		fn, args := desc.prepare(interp, fr)
		root := rootFrame()
		run := interp.prepareGoCall(root, fn, args, nil)
		atomic.AddInt32(&interp.goroutines, 1)
		go func() {
			defer atomic.AddInt32(&interp.goroutines, -1)
			if interp.ctx.RunContext != nil {
				defer func() {
					if e := recover(); e != nil {
						interp.cherror <- PanicError{stack: debugStack(root), Value: e}
					}
				}()
			}
			run()
		}()
	}
}

// instructionDeferCall builds the bytecode-only Defer executor. Lowering keeps
// its own implementation in opblock.go.
func instructionDeferCall(interp *Interp, desc callDescription, id register) func(*frame) {
	if id < 0 {
		return func(fr *frame) {
			fn, args := desc.prepare(interp, fr)
			fr._defer = &_defer{fn: fn, args: args, tail: fr._defer}
		}
	}
	return func(fr *frame) {
		fn, args := desc.prepare(interp, fr)
		d := &_defer{fn: fn, args: args}
		head, ok := deferStackHead(fr.reg(id))
		if !ok {
			panic("range-over-func defer: missing deferstack token")
		}
		pushDeferAt(head, d)
	}
}

// makeBuiltinInstruction preserves the register-based builtin fast paths.
// Layout builtins use independently computed constants/errors; unknown names
// return nil to retain lowering.
func makeBuiltinInstruction(interp *Interp, d typedInstruction, aux builtinCallAux) func(*frame) {
	var fn builtinFunction
	if aux.Target != nil {
		fn = aux.Target
	} else {
		fn = &builtinDescription{name: aux.Name}
	}
	ir, ia := d.Result, aux.Arguments
	if aux.Runtime.Bound {
		ir, ia = aux.Runtime.Result, aux.Runtime.Arguments
	}
	switch fn.Name() {
	case "append":
		if len(ia) == 1 {
			return func(fr *frame) {
				fr.copyReg(ir, ia[0])
			}
		}
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			arg1 := fr.reg(ia[1])
			v0 := reflect.ValueOf(arg0)
			v1 := reflect.ValueOf(arg1)
			// append([]byte, ...string) []byte
			if v1.Kind() == reflect.String {
				v1 = reflect.ValueOf([]byte(v1.String()))
			}
			i0 := v0.Len()
			i1 := v1.Len()
			if i0+i1 < i0 {
				panic(fr.runtimeError(fn, errAppendOutOfRange))
			}
			fr.setReg(ir, reflect.AppendSlice(v0, v1).Interface())
		}
	case "copy": // copy([]T, []T) int or copy([]byte, string) int
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			arg1 := fr.reg(ia[1])
			fr.setReg(ir, reflect.Copy(reflect.ValueOf(arg0), reflect.ValueOf(arg1)))
		}

	case "close": // close(chan T)
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			reflect.ValueOf(arg0).Close()
		}

	case "delete": // delete(map[K]value, K)
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			arg1 := fr.reg(ia[1])
			m := reflect.ValueOf(arg0)
			m.SetMapIndex(mapKeyValue(m, arg1), reflect.Value{})
		}

	case "print", "println": // print(any, ...)
		ln := fn.Name() == "println"
		interfaceArgs := aux.InterfaceArgs
		return func(fr *frame) {
			var buf bytes.Buffer
			for i := 0; i < len(ia); i++ {
				arg := fr.reg(ia[i])
				if i > 0 && ln {
					buf.WriteRune(' ')
				}
				if i < len(interfaceArgs) && interfaceArgs[i] {
					writeinterface(&buf, arg)
					continue
				}
				writevalue(&buf, arg, interp.ctx.Mode&EnablePrintAny != 0)
			}
			if ln {
				buf.WriteRune('\n')
			}
			interp.ctx.writeOutput(buf.Bytes())
		}

	case "len":
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			fr.setReg(ir, reflect.ValueOf(arg0).Len())
		}

	case "cap":
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			fr.setReg(ir, reflect.ValueOf(arg0).Cap())
		}

	case "real":
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			c := reflect.ValueOf(arg0)
			switch c.Kind() {
			case reflect.Complex64:
				fr.setReg(ir, real(complex64(c.Complex())))
			case reflect.Complex128:
				fr.setReg(ir, real(c.Complex()))
			default:
				panic(fmt.Sprintf("real: illegal operand: %T", c))
			}
		}

	case "imag":
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			c := reflect.ValueOf(arg0)
			switch c.Kind() {
			case reflect.Complex64:
				fr.setReg(ir, imag(complex64(c.Complex())))
			case reflect.Complex128:
				fr.setReg(ir, imag(c.Complex()))
			default:
				panic(fmt.Sprintf("imag: illegal operand: %T", c))
			}
		}

	case "complex":
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			arg1 := fr.reg(ia[1])
			r := reflect.ValueOf(arg0)
			i := reflect.ValueOf(arg1)
			switch r.Kind() {
			case reflect.Float32:
				fr.setReg(ir, complex(float32(r.Float()), float32(i.Float())))
			case reflect.Float64:
				fr.setReg(ir, complex(r.Float(), i.Float()))
			default:
				panic(fmt.Sprintf("complex: illegal operand: %v", r.Kind()))
			}
		}

	case "panic":
		// Compiled panic calls use opPanic; only hand-built Call descriptions
		// reach this branch.
		if interp.ctx.panicFunc != nil {
			return func(fr *frame) {
				arg0 := fr.reg(ia[0])
				var err error = PanicError{stack: debugStack(fr), Value: arg0}
				panic(interp.ctx.handlePanic(fr, fn, err))
			}
		}
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			panic(PanicError{stack: debugStack(fr), Value: arg0})
		}

	case "recover":
		return func(fr *frame) {
			fr.setReg(ir, doRecover(fr))
		}

	case "ssa:wrapnilchk":
		return func(fr *frame) {
			recv := fr.reg(ia[0])
			if reflect.ValueOf(recv).IsNil() {
				recvType := fr.reg(ia[1])
				methodName := fr.reg(ia[2])
				var info value
				if s, ok := recvType.(string); ok && strings.HasPrefix(s, "main.") {
					info = s[5:]
				} else {
					info = recvType
				}
				panic(fr.plainError(fn, fmt.Sprintf("value method %s.%s called using nil *%s pointer",
					recvType, methodName, info)))
			}
			fr.setReg(ir, recv)
		}

	case "Add":
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			arg1 := fr.reg(ia[1])
			ptr := arg0.(unsafe.Pointer)
			length := asInt(arg1)
			fr.setReg(ir, unsafe.Add(ptr, length))
		}
	case "Slice":
		//func Slice(ptr *ArbitraryType, len IntegerType) []ArbitraryType
		//(*[len]ArbitraryType)(unsafe.Pointer(ptr))[:]
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			arg1 := fr.reg(ia[1])
			ptr := reflect.ValueOf(arg0)
			etyp := ptr.Type().Elem()
			length := asInt(arg1)
			if ptr.IsNil() {
				if length == 0 {
					fr.setReg(ir, reflect.New(reflect.SliceOf(etyp)).Elem().Interface())
					return
				}
				panic(fr.runtimeError(fn, "unsafe.Slice: ptr is nil and len is not zero"))
			}
			mem, overflow := mulUintptr(etyp.Size(), uintptr(length))
			if overflow || mem > -uintptr(ptr.Pointer()) {
				panic(fr.runtimeError(fn, "unsafe.Slice: len out of range"))
			}
			typ := reflect.ArrayOf(length, etyp)
			v := reflect.NewAt(typ, ptr.UnsafePointer())
			fr.setReg(ir, v.Elem().Slice(0, length).Interface())
		}
	case "SliceData":
		// SliceData returns a pointer to the underlying array of the argument
		// slice.
		//   - If cap(slice) > 0, SliceData returns &slice[:1][0].
		//   - If slice == nil, SliceData returns nil.
		//   - Otherwise, SliceData returns a non-nil pointer to an
		//     unspecified memory address.
		// func SliceData(slice []ArbitraryType) *ArbitraryType
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			v := reflect.ValueOf(arg0)
			if v.Cap() > 0 {
				fr.setReg(ir, v.Slice(0, 1).Index(0).Addr().Interface())
			} else if v.IsNil() {
				fr.setReg(ir, reflect.Zero(reflect.PtrTo(v.Type().Elem())).Interface())
			} else {
				fr.setReg(ir, reflect.New(v.Type().Elem()).Interface())
			}
		}
	case "String":
		// String returns a string value whose underlying bytes
		// start at ptr and whose length is len.
		//
		// The len argument must be of integer type or an untyped constant.
		// A constant len argument must be non-negative and representable by a value of type int;
		// if it is an untyped constant it is given type int.
		// At run time, if len is negative, or if ptr is nil and len is not zero,
		// a run-time panic occurs.
		//
		// Since Go strings are immutable, the bytes passed to String
		// must not be modified afterwards.
		// func String(ptr *byte, len IntegerType) string
		return func(fr *frame) {
			ptr := fr.reg(ia[0]).(*byte)
			length := asInt(fr.reg(ia[1]))
			if ptr == nil {
				if length == 0 {
					fr.setReg(ir, "")
					return
				}
				panic(fr.runtimeError(fn, "unsafe.String: ptr is nil and len is not zero"))
			}
			mem, overflow := mulUintptr(1, uintptr(length))
			if overflow || mem > -uintptr(unsafe.Pointer(ptr)) {
				panic(fr.runtimeError(fn, "unsafe.String: len out of range"))
			}
			fr.setReg(ir, unsafe.String(ptr, length))
		}
	case "StringData":
		// StringData returns a pointer to the underlying bytes of str.
		// For an empty string the return value is unspecified, and may be nil.
		//
		// Since Go strings are immutable, the bytes returned by StringData
		// must not be modified.
		// func StringData(str string) *byte
		return func(fr *frame) {
			fr.setReg(ir, unsafe.StringData(fr.string(ia[0])))
		}
	case "Sizeof", "Alignof", "Offsetof":
		if message := aux.Error; message != "" {
			return func(fr *frame) { panic(fr.plainError(fn, message)) }
		}
		result := aux.Constant
		return func(fr *frame) { fr.setReg(ir, result) }
	case "clear":
		return func(fr *frame) {
			arg0 := fr.reg(ia[0])
			valueClear(reflect.ValueOf(arg0))
		}
	case "max":
		return func(fr *frame) {
			v := reflect.ValueOf(fr.reg(ia[0]))
			for i, r := range ia {
				arg := reflect.ValueOf(fr.reg(r))
				if i > 0 {
					switch v.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						if v.Int() < arg.Int() {
							v = arg
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
						if v.Uint() < arg.Uint() {
							v = arg
						}
					case reflect.Float32, reflect.Float64:
						x, y := v.Float(), arg.Float()
						if x < y || math.IsNaN(y) || x == 0 && y == 0 && !math.Signbit(y) {
							v = arg
						}
					case reflect.String:
						if v.String() < arg.String() {
							v = arg
						}
					}
				}
			}
			fr.setReg(ir, v.Interface())
		}
	case "min":
		return func(fr *frame) {
			v := reflect.ValueOf(fr.reg(ia[0]))
			for i, r := range ia {
				arg := reflect.ValueOf(fr.reg(r))
				if i > 0 {
					switch v.Kind() {
					case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
						if v.Int() > arg.Int() {
							v = arg
						}
					case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
						if v.Uint() > arg.Uint() {
							v = arg
						}
					case reflect.Float32, reflect.Float64:
						x, y := v.Float(), arg.Float()
						if x > y || math.IsNaN(y) || x == 0 && y == 0 && math.Signbit(y) {
							v = arg
						}
					case reflect.String:
						if v.String() > arg.String() {
							v = arg
						}
					}
				}
			}
			fr.setReg(ir, v.Interface())
		}
	case "ssa:deferstack":
		return func(fr *frame) {
			fr.setReg(ir, &fr._defer)
		}
	default:
		return nil
	}
}

// makeInvokeInstruction caches a bytecode executor per receiver type.
func makeInvokeInstruction(interp *Interp, d typedInstruction, aux invokeCallAux) func(*frame) {
	mname := aux.Method
	ir, iv, ia, dests := d.Result, aux.Target, aux.Arguments, aux.ExtractResults
	if d.RuntimeBound {
		ir, iv, ia, dests = d.RuntimeResult, aux.runtime.Target, aux.runtime.Arguments, aux.runtime.ExtractResults
	}
	ia = append([]register{iv}, ia...)
	var methods sync.Map // map[reflect.Type]func(*frame)
	return func(fr *frame) {
		v := fr.reg(iv)
		if v == nil {
			panic(fr.runtimeError(fr.pfn.InstrForPC(int(fr.ipc)-1), "runtime error: invalid memory address or nil pointer dereference"))
		}
		rtype := reflect.TypeOf(v)
		if method, ok := methods.Load(rtype); ok {
			method.(func(*frame))(fr)
			return
		}

		var method func(*frame)
		// Find an interpreted method.
		if mset, ok := interp.msets[rtype]; ok {
			if fn, ok := mset[mname]; ok {
				method = func(fr *frame) {
					interp.callFunctionByStack(fr, fn, ir, ia, dests)
				}
			} else {
				ext, found := findUserMethod(rtype, mname)
				if !found {
					panic(fr.plainError(fr.pfn.InstrForPC(int(fr.ipc)-1), fmt.Sprintf("no code for method: %v.%v", rtype, mname)))
				}
				method = func(fr *frame) {
					interp.callExternalByStack(fr, ext, ir, ia, dests)
				}
			}
		} else {
			ext, found := findExternMethod(rtype, mname)
			if !found {
				panic(fr.plainError(fr.pfn.InstrForPC(int(fr.ipc)-1), fmt.Sprintf("no code for method: %v.%v", rtype, mname)))
			}
			if adapter, ok := resolveInvokeDirectCall(rtype, mname); ok {
				method = func(fr *frame) {
					interp.invokeDirectCall(fr, adapter, ir, ia)
				}
			} else {
				method = func(fr *frame) {
					interp.callExternalByStack(fr, ext, ir, ia, dests)
				}
			}
		}
		actual, _ := methods.LoadOrStore(rtype, method)
		actual.(func(*frame))(fr)
	}
}

// makeBodyCallInstruction constructs bytecode calls to known bodies.
func makeBodyCallInstruction(interp *Interp, body *function, result register, args, dests []register) func(*frame) {

	if body == nil {
		return nil
	}
	if body.Recover == nil {
		if run := makeCallSpecializerInstr(interp, body, args, dests); run != nil {
			return run
		}
		switch body.nres {
		case 0:
			return func(fr *frame) { interp.callFunctionByStackNoRecover0(fr, body, result, args) }
		case 1:
			return func(fr *frame) { interp.callFunctionByStackNoRecover1(fr, body, result, args) }
		default:
			return func(fr *frame) { interp.callFunctionByStackNoRecoverN(fr, body, result, args, dests) }
		}
	}
	switch body.nres {
	case 0:
		return func(fr *frame) { interp.callFunctionByStack0(fr, body, result, args) }
	case 1:
		return func(fr *frame) { interp.callFunctionByStack1(fr, body, result, args) }
	default:
		return func(fr *frame) { interp.callFunctionByStackN(fr, body, result, args, dests) }
	}
}
