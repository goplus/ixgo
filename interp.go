// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package ixgo defines an interpreter for the SSA
// representation of Go programs.
//
// This interpreter is provided as an adjunct for testing the SSA
// construction algorithm.  Its purpose is to provide a minimal
// metacircular implementation of the dynamic semantics of each SSA
// instruction.  It is not, and will never be, a production-quality Go
// interpreter.
//
// The following is a partial list of Go features that are currently
// unsupported or incomplete in the interpreter.
//
// * Unsafe operations, including all uses of unsafe.Pointer, are
// impossible to support given the "boxed" value representation we
// have chosen.
//
// * The reflect package is only partially implemented.
//
// * The "testing" package is no longer supported because it
// depends on low-level details that change too often.
//
// * "sync/atomic" operations are not atomic due to the "boxed" value
// representation: it is not possible to read, modify and write an
// interface value atomically. As a consequence, Mutexes are currently
// broken.
//
// * recover is only partially implemented.  Also, the interpreter
// makes no attempt to distinguish target panics from interpreter
// crashes.
//
// * the sizes of the int, uint and uintptr types in the target
// program are assumed to be the same as those of the interpreter
// itself.
//
// * all values occupy space, even those of types defined by the spec
// to have zero size, e.g. struct{}.  This can cause asymptotic
// performance degradation.
//
// * os.Exit is implemented using panic, causing deferred functions to
// run.

package ixgo

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/token"
	"go/types"
	"reflect"
	"strings"
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/goplus/ixgo/internal/typesalias"
	"github.com/goplus/ixgo/load"
	"github.com/goplus/reflectx"
	"github.com/timandy/routine"
	"github.com/visualfc/xtype"
	"golang.org/x/tools/go/ssa"
)

var (
	maxMemLen int
)

const intSize = 32 << (^uint(0) >> 63)

func init() {
	if intSize == 32 {
		maxMemLen = 1<<31 - 1
	} else {
		v := int64(1) << 59
		maxMemLen = int(v)
	}
}

type Interp struct {
	ctx          *Context
	rctx         *reflectx.Context
	mainpkg      *ssa.Package                                // the SSA main package
	record       *TypesRecord                                // lookup type and ToType
	globals      map[string]value                            // addresses of global variables (immutable)
	chkinit      map[string]bool                             // init vars
	preloadTypes map[types.Type]reflect.Type                 // preload types.Type -> reflect.Type
	funcs        map[*ssa.Function]*function                 // ssa.Function -> *function
	msets        map[reflect.Type](map[string]*ssa.Function) // user defined type method sets
	chexit       chan int                                    // call os.Exit code by chan for runtime.Goexit
	cherror      chan PanicError                             // call by go func error for context
	deferMap     sync.Map                                    // defer goroutine id -> call frame
	rfuncMap     sync.Map                                    // reflect.Value(fn).Pointer -> *function
	typesMutex   sync.RWMutex                                // findType/toType mutex
	mainid       uint64                                      // main goroutine id
	exitCode     int                                         // call os.Exit code
	goroutines   int32                                       // atomically updated
	deferCount   int32                                       // fast has defer check
	goexited     int32                                       // is call runtime.Goexit
	exited       int32                                       // is call os.Exit
}

func (i *Interp) MainPkg() *ssa.Package {
	return i.mainpkg
}

func (i *Interp) installed(path string) (pkg *Package, ok bool) {
	pkg, ok = i.ctx.Loader.Installed(path)
	return
}

func (i *Interp) loadFunction(fn *ssa.Function) *function {
	if pfn, ok := i.funcs[fn]; ok {
		return pfn
	}
	pfn := &function{
		Interp:     i,
		Fn:         fn,
		index:      make(map[ssa.Value]uint32),
		instrIndex: make(map[ssa.Instruction][]uint32),
		narg:       len(fn.Params),
		nenv:       len(fn.FreeVars),
	}
	if len(fn.Blocks) > 0 {
		pfn.Main = fn.Blocks[0]
	}
	if res := fn.Signature.Results(); res != nil {
		pfn.nres = res.Len()
		pfn.stack = make([]value, pfn.nres)
	}
	i.funcs[fn] = pfn
	return pfn
}

func (i *Interp) findType(rt reflect.Type, local bool) (types.Type, bool) {
	i.typesMutex.Lock()
	defer i.typesMutex.Unlock()
	if local {
		return i.record.LookupLocalTypes(rt)
	} else {
		return i.record.LookupTypes(rt)
	}
}

func (i *Interp) tryDeferFrame() *frame {
	if i != nil && atomic.LoadInt32(&i.deferCount) != 0 {
		if f, ok := i.deferMap.Load(goroutineID()); ok {
			return f.(*frame)
		}
	}
	return &frame{}
}

func (pfn *function) callFunctionByReflect(mtyp reflect.Type, args []reflect.Value, env []interface{}) []reflect.Value {
	return pfn.Interp.callFunctionByReflect(pfn.Interp.tryDeferFrame(), pfn, mtyp, args, env)
}

func (i *Interp) FindMethod(mtyp reflect.Type, fn *types.Func) func([]reflect.Value) []reflect.Value {
	typ := fn.Type().(*types.Signature).Recv().Type()
	if f := i.mainpkg.Prog.LookupMethod(typ, fn.Pkg(), fn.Name()); f != nil {
		pfn := i.loadFunction(f)
		return func(args []reflect.Value) []reflect.Value {
			return pfn.callFunctionByReflect(mtyp, args, nil)
		}
	}
	name := fn.FullName()
	if v, ok := externValues[name]; ok && v.Kind() == reflect.Func {
		if v.Type().IsVariadic() {
			return func(args []reflect.Value) []reflect.Value {
				return v.CallSlice(args)
			}
		}
		return func(args []reflect.Value) []reflect.Value {
			return v.Call(args)
		}
	}
	panic(fmt.Sprintf("Not found method %v", fn))
}

func (pfn *function) makeFunction(typ reflect.Type, env []value) reflect.Value {
	interp := pfn.Interp
	return reflect.MakeFunc(typ, func(args []reflect.Value) []reflect.Value {
		return interp.callFunctionByReflect(interp.tryDeferFrame(), pfn, typ, args, env)
	})
}

type _defer struct {
	fn      value
	tail    *_defer
	args    []value
	ssaArgs []ssa.Value
}

type frame struct {
	interp  *Interp
	caller  *frame
	callee  *frame
	pfn     *function
	_defer  *_defer
	_panic  *_panic
	block   *ssa.BasicBlock
	stack   []value // result args env datas
	ipc     int
	pred    int
	deferid uint64
}

func dumpBlock(block *ssa.BasicBlock, level int, pc ssa.Instruction) {
	if level == 0 {
		fmt.Printf("--- %v ---\n", block.Parent())
	}
	fmt.Printf("%v.%v %v    Jump:%v Idom:%v\n", strings.Repeat("  ", level), block.Index, block.Comment, block.Succs, block.Idom())
	for _, instr := range block.Instrs {
		var head string
		if instr == pc {
			head = " " + strings.Repeat("    ", level) + "=>"
		} else {
			head = "   " + strings.Repeat("    ", level)
		}
		if value, ok := instr.(ssa.Value); ok {
			fmt.Printf("%v %-20T %-4v = %v %v\n", head, instr, value.Name(), instr, value.Type())
		} else {
			fmt.Printf("%v %-20T   %v\n", head, instr, instr)
		}
	}
	for _, v := range block.Dominees() {
		dumpBlock(v, level+1, pc)
	}
}

type tasks struct {
	jumps map[*ssa.BasicBlock]bool
}

func checkJumps(block *ssa.BasicBlock, jumps map[*ssa.BasicBlock]bool, succs map[*ssa.BasicBlock]*ssa.BasicBlock) {
	if s, ok := succs[block]; ok {
		if jumps[s] {
			return
		}
		jumps[s] = true
		checkJumps(s, jumps, succs)
		return
	}
	for _, s := range block.Succs {
		if jumps[s] {
			continue
		}
		jumps[s] = true
		checkJumps(s, jumps, succs)
	}
}

func checkRuns(block *ssa.BasicBlock, jumps map[*ssa.BasicBlock]bool, succs map[*ssa.BasicBlock]*ssa.BasicBlock) {
	if jumps[block] {
		return
	}
	jumps[block] = true
	if s, ok := succs[block]; ok {
		checkRuns(s, jumps, succs)
		return
	}
	for _, s := range block.Dominees() {
		checkRuns(s, jumps, succs)
	}
}

func (fr *frame) gc() {
	alloc := make(map[int]bool)
	checkAlloc := func(instr ssa.Instruction) {
		for _, v := range fr.pfn.instrIndex[instr] {
			vk := kind(v >> 30)
			if vk.isStatic() {
				continue
			}
			rk := reflect.Kind(v >> 24 & 0x3f)
			switch rk {
			case reflect.String, reflect.Func, reflect.Ptr, reflect.Array, reflect.Slice, reflect.Map, reflect.Struct, reflect.Interface:
			default:
				continue
			}
			alloc[int(v&0xffffff)] = true
		}
	}
	// check params and freevar
	checkAlloc(nil)
	// check alloc
	cur := fr.pfn.ssaInstrs[fr.ipc-1]
	var remain int
	for i, instr := range fr.block.Instrs {
		checkAlloc(instr)
		if cur == instr {
			remain = i
			break
		}
	}

	// check
	seen := make(map[*ssa.BasicBlock]bool)
	var checkChild func(block *ssa.BasicBlock)
	checkChild = func(block *ssa.BasicBlock) {
		for _, child := range block.Dominees() {
			if seen[child] {
				continue
			}
			seen[child] = true
			checkChild(child)
		}
	}
	// check block child
	checkChild(fr.block)

	block := fr.block
	// check seen
	for block != nil {
		idom := block.Idom()
		if idom == nil {
			break
		}
		var find bool
		switch idom.Comment {
		case "for.loop":
			find = true
		case "rangeindex.loop":
			find = true
		case "rangechan.loop":
		case "rangeiter.loop":
		}
		for _, v := range idom.Succs {
			if find {
				seen[v] = true
				checkChild(v)
			}
			if block == v {
				find = true
			}
		}
		for _, instr := range idom.Instrs {
			checkAlloc(instr)
		}
		block = idom
	}
	if fr.block.Comment == "for.done" {
		delete(seen, fr.block)
	}
	// check used in block
	var ops []*ssa.Value
	for _, instr := range fr.block.Instrs[remain+1:] {
		for _, op := range instr.Operands(ops[:0]) {
			if *op == nil {
				continue
			}
			reg := fr.pfn.regIndex(*op)
			alloc[int(reg)] = false
		}
	}
	// check unused in seen
	for block := range seen {
		var ops []*ssa.Value
		for _, instr := range block.Instrs {
			for _, op := range instr.Operands(ops[:0]) {
				if *op == nil {
					continue
				}
				reg := fr.pfn.regIndex(*op)
				alloc[int(reg)] = false
			}
		}
	}
	// remove unused
	for i, b := range alloc {
		if !b {
			continue
		}
		fr.stack[i] = nil
	}
}

func (fr *frame) valid() bool {
	return fr != nil && fr.pfn != nil && fr.block != nil
}

func (fr *frame) pc() uintptr {
	return uintptr(fr.pfn.base + fr.ipc)
}

func (fr *frame) aborted() bool {
	return fr != nil && fr.ipc != -1
}

func (fr *frame) setReg(ir register, v value) {
	fr.stack[ir] = v
}

func (fr *frame) reg(ir register) value {
	return fr.stack[ir]
}

func (fr *frame) bytes(ir register) []byte {
	return xtype.Bytes(fr.stack[ir])
}

func (fr *frame) runes(ir register) []rune {
	return xtype.Runes(fr.stack[ir])
}

func (fr *frame) bool(ir register) bool {
	return xtype.Bool(fr.stack[ir])
}

func (fr *frame) int(ir register) int {
	return xtype.Int(fr.stack[ir])
}

func (fr *frame) int8(ir register) int8 {
	return xtype.Int8(fr.stack[ir])
}

func (fr *frame) int16(ir register) int16 {
	return xtype.Int16(fr.stack[ir])
}

func (fr *frame) int32(ir register) int32 {
	return xtype.Int32(fr.stack[ir])
}

func (fr *frame) int64(ir register) int64 {
	return xtype.Int64(fr.stack[ir])
}

func (fr *frame) uint(ir register) uint {
	return xtype.Uint(fr.stack[ir])
}

func (fr *frame) uint8(ir register) uint8 {
	return xtype.Uint8(fr.stack[ir])
}

func (fr *frame) uint16(ir register) uint16 {
	return xtype.Uint16(fr.stack[ir])
}

func (fr *frame) uint32(ir register) uint32 {
	return xtype.Uint32(fr.stack[ir])
}

func (fr *frame) uint64(ir register) uint64 {
	return xtype.Uint64(fr.stack[ir])
}

func (fr *frame) uintptr(ir register) uintptr {
	return xtype.Uintptr(fr.stack[ir])
}

func (fr *frame) float32(ir register) float32 {
	return xtype.Float32(fr.stack[ir])
}

func (fr *frame) float64(ir register) float64 {
	return xtype.Float64(fr.stack[ir])
}

func (fr *frame) complex64(ir register) complex64 {
	return xtype.Complex64(fr.stack[ir])
}

func (fr *frame) complex128(ir register) complex128 {
	return xtype.Complex128(fr.stack[ir])
}

func (fr *frame) string(ir register) string {
	return xtype.String(fr.stack[ir])
}

func (fr *frame) pointer(ir register) unsafe.Pointer {
	return xtype.Pointer(fr.stack[ir])
}

func (fr *frame) copyReg(dst register, src register) {
	fr.stack[dst] = fr.stack[src]
}

func (fr *frame) runtimeError(instr funcInstr, text string) error {
	if fr.interp.ctx.panicFunc != nil {
		return fr.interp.ctx.handlePanic(fr, instr, RuntimeError(text))
	}
	return RuntimeError(text)
}

func (fr *frame) plainError(instr funcInstr, text string) error {
	if fr.interp.ctx.panicFunc != nil {
		return fr.interp.ctx.handlePanic(fr, instr, PlainError(text))
	}
	return PlainError(text)
}

type _panic struct {
	arg       interface{}
	link      *_panic
	pcs       []uintptr
	aborted   bool
	recovered bool
}

func (p *_panic) isNil() bool {
	return p == nil || p.recovered
}

// runDefer runs a deferred call d.
// It always returns normally, but may set or clear fr.panic.
func (fr *frame) runDefer(d *_defer) (ok bool) {
	defer func() {
		if !ok {
			// Deferred call created a new state of panic.
			if fr._panic != nil {
				fr._panic.aborted = true
			}
			fr._panic = &_panic{arg: recover(), link: fr._panic}
			// no tail add callee.pc
			if d.tail != nil {
				callee := fr.callee
				for callee.aborted() {
					fr._panic.pcs = append([]uintptr{callee.pc()}, fr._panic.pcs...)
					callee = callee.callee
				}
			}
		}
	}()
	fr.interp.callDiscardsResult(fr, d.fn, d.args, d.ssaArgs)
	ok = true
	return
}

// runDefers executes fr's deferred function calls in LIFO order.
//
// On entry, fr.panicking indicates a state of panic; if
// true, fr.panic contains the panic value.
//
// On completion, if a deferred call started a panic, or if no
// deferred call recovered from a previous state of panic, then
// runDefers itself panics after the last deferred call has run.
//
// If there was no initial state of panic, or it was recovered from,
// runDefers returns normally.
func (fr *frame) runDefers() {
	interp := fr.interp
	atomic.AddInt32(&interp.deferCount, 1)
	fr.deferid = goroutineID()
	interp.deferMap.Store(fr.deferid, fr)
	for d := fr._defer; d != nil; d = d.tail {
		fr.runDefer(d)
	}
	interp.deferMap.Delete(fr.deferid)
	atomic.AddInt32(&interp.deferCount, -1)
	fr.deferid = 0
	fr._defer = nil
	// runtime.Goexit() fr.panic == nil
	if !fr._panic.isNil() {
		panic(fr._panic.arg) // new panic, or still panicking
	}
}

// lookupMethod returns the method set for type typ, which may be one
// of the interpreter's fake types.
func lookupMethod(i *Interp, typ types.Type, meth *types.Func) *ssa.Function {
	return i.mainpkg.Prog.LookupMethod(typ, meth.Pkg(), meth.Name())
}

func SetValue(v reflect.Value, x reflect.Value) {
	switch v.Kind() {
	case reflect.Bool:
		v.SetBool(x.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(x.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v.SetUint(x.Uint())
	case reflect.Uintptr:
		v.SetUint(x.Uint())
	case reflect.Float32, reflect.Float64:
		v.SetFloat(x.Float())
	case reflect.Complex64, reflect.Complex128:
		v.SetComplex(x.Complex())
	case reflect.String:
		v.SetString(x.String())
	case reflect.UnsafePointer:
		v.SetPointer(unsafe.Pointer(x.Pointer()))
	case reflect.Struct:
		v.Set(x)
	default:
		v.Set(x)
	}
}

type DebugInfo struct {
	*ssa.DebugRef
	fset    *token.FileSet
	toValue func() (*types.Var, interface{}, bool) // var object value
}

func (i *DebugInfo) Position() token.Position {
	return i.fset.Position(i.Pos())
}

func (i *DebugInfo) AsVar() (*types.Var, interface{}, bool) {
	return i.toValue()
}

func (i *DebugInfo) AsFunc() (*types.Func, bool) {
	v, ok := i.Object().(*types.Func)
	return v, ok
}

// prepareCall determines the function value and argument values for a
// function call in a Call, Go or Defer instruction, performing
// interface method lookup if needed.
func (i *Interp) prepareCall(fr *frame, call *ssa.CallCommon, iv register, ia []register, ib []register) (fv value, args []value) {
	if call.Method == nil {
		switch f := call.Value.(type) {
		case *ssa.Builtin:
			fv = f
		case *ssa.Function:
			if f.Blocks == nil {
				ext, ok := findExternFunc(i, f)
				if !ok {
					// skip pkg.init
					if f.Pkg != nil && f.Name() == "init" {
						fv = func() {}
					} else {
						panic(fmt.Errorf("no code for function: %v", f))
					}
				} else {
					fv = ext
				}
			} else {
				fv = f
			}
		case *ssa.MakeClosure:
			var bindings []value
			for i := range f.Bindings {
				bindings = append(bindings, fr.reg(ib[i]))
			}
			fv = &closure{i.funcs[f.Fn.(*ssa.Function)], bindings}
		default:
			fv = fr.reg(iv)
		}
	} else {
		v := fr.reg(iv)
		rtype := reflect.TypeOf(v)
		mname := call.Method.Name()
		if mset, ok := i.msets[rtype]; ok {
			if f, ok := mset[mname]; ok {
				fv = f
			} else {
				ext, ok := findUserMethod(rtype, mname)
				if !ok {
					panic(fmt.Errorf("no code for method: %v.%v", rtype, mname))
				}
				fv = ext
			}
		} else {
			ext, ok := findExternMethod(rtype, mname)
			if !ok {
				panic(fmt.Errorf("no code for method: %v.%v", rtype, mname))
			}
			fv = ext
		}
		args = append(args, v)
	}
	for i := range call.Args {
		v := fr.reg(ia[i])
		args = append(args, v)
	}
	return
}

// call interprets a call to a function (function, builtin or closure)
// fn with arguments args, returning its result.
// callpos is the position of the callsite.
func (i *Interp) call(caller *frame, fn value, args []value, ssaArgs []ssa.Value) value {
	switch fn := fn.(type) {
	case *ssa.Function:
		return i.callFunction(caller, i.funcs[fn], args, nil)
	case *closure:
		return i.callFunction(caller, fn.pfn, args, fn.env)
	case *ssa.Builtin:
		return i.callBuiltin(caller, fn, args, ssaArgs)
	case reflect.Value:
		return i.callExternal(caller, fn, args, nil)
	default:
		return i.callExternal(caller, reflect.ValueOf(fn), args, nil)
	}
}

// call interprets a call to a function (function, builtin or closure)
// fn with arguments args, returning its result.
// callpos is the position of the callsite.
func (i *Interp) callDiscardsResult(caller *frame, fn value, args []value, ssaArgs []ssa.Value) {
	switch fn := fn.(type) {
	case *ssa.Function:
		i.callFunctionDiscardsResult(caller, i.funcs[fn], args, nil)
	case *closure:
		i.callFunctionDiscardsResult(caller, fn.pfn, args, fn.env)
	case *ssa.Builtin:
		i.callBuiltinDiscardsResult(caller, fn, args, ssaArgs)
	case reflect.Value:
		i.callExternalDiscardsResult(caller, fn, args, nil)
	default:
		i.callExternalDiscardsResult(caller, reflect.ValueOf(fn), args, nil)
	}
}

func (i *Interp) callFunction(caller *frame, pfn *function, args []value, env []value) (result value) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < pfn.narg; i++ {
		fr.stack[i+pfn.nres] = args[i]
	}
	for i := 0; i < pfn.nenv; i++ {
		fr.stack[pfn.narg+i+pfn.nres] = env[i]
	}
	fr.run()
	if pfn.nres == 1 {
		result = fr.stack[0]
	} else if pfn.nres > 1 {
		result = tuple(fr.stack[0:pfn.nres])
	}
	pfn.deleteFrame(caller, fr)
	return
}

func (i *Interp) callFunctionByReflect(caller *frame, pfn *function, typ reflect.Type, args []reflect.Value, env []value) (results []reflect.Value) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < pfn.narg; i++ {
		fr.stack[i+pfn.nres] = args[i].Interface()
	}
	for i := 0; i < pfn.nenv; i++ {
		fr.stack[pfn.narg+i+pfn.nres] = env[i]
	}
	fr.run()
	if pfn.nres > 0 {
		results = make([]reflect.Value, pfn.nres)
		for i := 0; i < pfn.nres; i++ {
			v := fr.stack[i]
			if v == nil {
				results[i] = reflect.New(typ.Out(i)).Elem()
			} else {
				results[i] = reflect.ValueOf(v)
			}
		}
	}
	pfn.deleteFrame(caller, fr)
	return
}

func (i *Interp) callFunctionDiscardsResult(caller *frame, pfn *function, args []value, env []value) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < pfn.narg; i++ {
		fr.stack[i+pfn.nres] = args[i]
	}
	for i := 0; i < pfn.nenv; i++ {
		fr.stack[pfn.narg+i+pfn.nres] = env[i]
	}
	fr.run()
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callFunctionByStack0(caller *frame, pfn *function, ir register, ia []register) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < len(ia); i++ {
		fr.stack[i] = caller.reg(ia[i])
	}
	fr.run()
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callFunctionByStack1(caller *frame, pfn *function, ir register, ia []register) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < len(ia); i++ {
		fr.stack[i+1] = caller.reg(ia[i])
	}
	fr.run()
	caller.setReg(ir, fr.stack[0])
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callFunctionByStackN(caller *frame, pfn *function, ir register, ia []register) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < len(ia); i++ {
		fr.stack[i+pfn.nres] = caller.reg(ia[i])
	}
	fr.run()
	caller.setReg(ir, tuple(fr.stack[0:pfn.nres]))
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callFunctionByStack(caller *frame, pfn *function, ir register, ia []register) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < len(ia); i++ {
		fr.stack[i+pfn.nres] = caller.reg(ia[i])
	}
	fr.run()
	if pfn.nres == 1 {
		caller.setReg(ir, fr.stack[0])
	} else if pfn.nres > 1 {
		caller.setReg(ir, tuple(fr.stack[0:pfn.nres]))
	}
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callFunctionByStackNoRecover0(caller *frame, pfn *function, ir register, ia []register) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < len(ia); i++ {
		fr.stack[i] = caller.reg(ia[i])
	}
	for fr.ipc != -1 {
		fn := fr.pfn.Instrs[fr.ipc]
		fr.ipc++
		fn(fr)
	}
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callFunctionByStackNoRecover1(caller *frame, pfn *function, ir register, ia []register) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < len(ia); i++ {
		fr.stack[i+1] = caller.reg(ia[i])
	}
	for fr.ipc != -1 {
		fn := fr.pfn.Instrs[fr.ipc]
		fr.ipc++
		fn(fr)
	}
	caller.setReg(ir, fr.stack[0])
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callFunctionByStackNoRecoverN(caller *frame, pfn *function, ir register, ia []register) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < len(ia); i++ {
		fr.stack[i+pfn.nres] = caller.reg(ia[i])
	}
	for fr.ipc != -1 {
		fn := fr.pfn.Instrs[fr.ipc]
		fr.ipc++
		fn(fr)
	}
	caller.setReg(ir, tuple(fr.stack[0:pfn.nres]))
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callFunctionByStackWithEnv(caller *frame, pfn *function, ir register, ia []register, env []value) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < pfn.narg; i++ {
		fr.stack[i+pfn.nres] = caller.reg(ia[i])
	}
	for i := 0; i < pfn.nenv; i++ {
		fr.stack[pfn.narg+i+pfn.nres] = env[i]
	}
	fr.run()
	if pfn.nres == 1 {
		caller.setReg(ir, fr.stack[0])
	} else if pfn.nres > 1 {
		caller.setReg(ir, tuple(fr.stack[0:pfn.nres]))
	}
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callFunctionByStackNoRecoverWithEnv(caller *frame, pfn *function, ir register, ia []register, env []value) {
	fr := pfn.allocFrame(caller)
	for i := 0; i < pfn.narg; i++ {
		fr.stack[i+pfn.nres] = caller.reg(ia[i])
	}
	for i := 0; i < pfn.nenv; i++ {
		fr.stack[pfn.narg+i+pfn.nres] = env[i]
	}
	for fr.ipc != -1 {
		fn := fr.pfn.Instrs[fr.ipc]
		fr.ipc++
		fn(fr)
	}
	if pfn.nres == 1 {
		caller.setReg(ir, fr.stack[0])
	} else if pfn.nres > 1 {
		caller.setReg(ir, tuple(fr.stack[0:pfn.nres]))
	}
	pfn.deleteFrame(caller, fr)
}

func (i *Interp) callExternal(caller *frame, fn reflect.Value, args []value, env []value) value {
	if caller != nil && caller.deferid != 0 {
		i.deferMap.Store(caller.deferid, caller)
	}
	var ins []reflect.Value
	typ := fn.Type()
	isVariadic := fn.Type().IsVariadic()
	if isVariadic {
		for i := 0; i < len(args)-1; i++ {
			if args[i] == nil {
				ins = append(ins, reflect.New(typ.In(i)).Elem())
			} else {
				ins = append(ins, reflect.ValueOf(args[i]))
			}
		}
		ins = append(ins, reflect.ValueOf(args[len(args)-1]))
	} else {
		ins = make([]reflect.Value, len(args))
		for i := 0; i < len(args); i++ {
			if args[i] == nil {
				ins[i] = reflect.New(typ.In(i)).Elem()
			} else {
				ins[i] = reflect.ValueOf(args[i])
			}
		}
	}
	var results []reflect.Value
	if isVariadic {
		results = fn.CallSlice(ins)
	} else {
		results = fn.Call(ins)
	}
	switch len(results) {
	case 0:
		return nil
	case 1:
		return results[0].Interface()
	default:
		var res []value
		for _, r := range results {
			res = append(res, r.Interface())
		}
		return tuple(res)
	}
}
func (i *Interp) callExternalDiscardsResult(caller *frame, fn reflect.Value, args []value, env []value) {
	if caller != nil && caller.deferid != 0 {
		i.deferMap.Store(caller.deferid, caller)
	}
	var ins []reflect.Value
	typ := fn.Type()
	isVariadic := fn.Type().IsVariadic()
	if isVariadic {
		for i := 0; i < len(args)-1; i++ {
			if args[i] == nil {
				ins = append(ins, reflect.New(typ.In(i)).Elem())
			} else {
				ins = append(ins, reflect.ValueOf(args[i]))
			}
		}
		ins = append(ins, reflect.ValueOf(args[len(args)-1]))
		fn.CallSlice(ins)
	} else {
		ins = make([]reflect.Value, len(args))
		for i := 0; i < len(args); i++ {
			if args[i] == nil {
				ins[i] = reflect.New(typ.In(i)).Elem()
			} else {
				ins[i] = reflect.ValueOf(args[i])
			}
		}
		fn.Call(ins)
	}
}

func (i *Interp) callExternalByStack(caller *frame, fn reflect.Value, ir register, ia []register) {
	if caller.deferid != 0 {
		i.deferMap.Store(caller.deferid, caller)
	}
	var ins []reflect.Value
	typ := fn.Type()
	isVariadic := fn.Type().IsVariadic()
	if isVariadic {
		var i int
		for n := len(ia) - 1; i < n; i++ {
			arg := caller.reg(ia[i])
			if arg == nil {
				ins = append(ins, reflect.New(typ.In(i)).Elem())
			} else {
				ins = append(ins, reflect.ValueOf(arg))
			}
		}
		ins = append(ins, reflect.ValueOf(caller.reg(ia[i])))
	} else {
		n := len(ia)
		ins = make([]reflect.Value, n)
		for i := 0; i < n; i++ {
			arg := caller.reg(ia[i])
			if arg == nil {
				ins[i] = reflect.New(typ.In(i)).Elem()
			} else {
				ins[i] = reflect.ValueOf(arg)
			}
		}
	}
	var results []reflect.Value
	if isVariadic {
		results = fn.CallSlice(ins)
	} else {
		results = fn.Call(ins)
	}
	switch len(results) {
	case 0:
	case 1:
		caller.setReg(ir, results[0].Interface())
	default:
		var res []value
		for _, r := range results {
			res = append(res, r.Interface())
		}
		caller.setReg(ir, tuple(res))
	}
}

func (i *Interp) callExternalWithFrameByStack(caller *frame, fn reflect.Value, ir register, ia []register) {
	if caller.deferid != 0 {
		i.deferMap.Store(caller.deferid, caller)
	}
	var ins []reflect.Value
	typ := fn.Type()
	isVariadic := fn.Type().IsVariadic()
	if isVariadic {
		ins = append(ins, reflect.ValueOf(caller))
		var i int
		for n := len(ia) - 1; i < n; i++ {
			arg := caller.reg(ia[i])
			if arg == nil {
				ins = append(ins, reflect.New(typ.In(i)).Elem())
			} else {
				ins = append(ins, reflect.ValueOf(arg))
			}
		}
		ins = append(ins, reflect.ValueOf(caller.reg(ia[i])))
	} else {
		n := len(ia)
		ins = make([]reflect.Value, n+1)
		ins[0] = reflect.ValueOf(caller)
		for i := 0; i < n; i++ {
			arg := caller.reg(ia[i])
			if arg == nil {
				ins[i+1] = reflect.New(typ.In(i)).Elem()
			} else {
				ins[i+1] = reflect.ValueOf(arg)
			}
		}
	}
	var results []reflect.Value
	if isVariadic {
		results = fn.CallSlice(ins)
	} else {
		results = fn.Call(ins)
	}
	switch len(results) {
	case 0:
	case 1:
		caller.setReg(ir, results[0].Interface())
	default:
		var res []value
		for _, r := range results {
			res = append(res, r.Interface())
		}
		caller.setReg(ir, tuple(res))
	}
}

// runFrame executes SSA instructions starting at fr.block and
// continuing until a return, a panic, or a recovered panic.
//
// After a panic, runFrame panics.
//
// After a normal return, fr.result contains the result of the call
// and fr.block is nil.
//
// A recovered panic in a function without named return parameters
// (NRPs) becomes a normal return of the zero value of the function's
// result type.
//
// After a recovered panic in a function with NRPs, fr.result is
// undefined and fr.block contains the block at which to resume
// control.
func (fr *frame) run() {
	if fr.pfn.Recover != nil {
		defer func() {
			if fr.ipc == -1 || fr._defer == nil {
				return // normal return
			}
			fr._panic = &_panic{arg: recover()}
			callee := fr.callee
			for callee.aborted() {
				if !callee._panic.isNil() {
					if !callee._panic.link.isNil() {
						fr._panic.link = callee._panic.link
						// check panic link
						link := callee._panic.link
						for link.link != nil {
							link = link.link
						}
						link.pcs = append(link.pcs, callee.pc())
					} else {
						fr._panic.pcs = append([]uintptr{callee.pc()}, fr._panic.pcs...)
					}
					fr._panic.pcs = append(append([]uintptr{}, callee._panic.pcs...), fr._panic.pcs...)
				} else {
					fr._panic.pcs = append([]uintptr{callee.pc()}, fr._panic.pcs...)
				}
				callee = callee.callee
			}
			fr.runDefers()
			for _, fn := range fr.pfn.Recover {
				fn(fr)
			}
		}()
	}

	for fr.ipc != -1 && atomic.LoadInt32(&fr.interp.exited) == 0 {
		fn := fr.pfn.Instrs[fr.ipc]
		fr.ipc++
		fn(fr)
	}
}

// doRecover implements the recover() built-in.
func doRecover(caller *frame) value {
	// recover() must be exactly one level beneath the deferred
	// function (two levels beneath the panicking function) to
	// have any effect.  Thus we ignore both "defer recover()" and
	// "defer f() -> g() -> recover()".
	if caller.interp.ctx.Mode&DisableRecover == 0 &&
		caller._panic.isNil() &&
		caller.caller != nil && !caller.caller._panic.isNil() {
		p := caller.caller._panic.arg
		caller.caller._panic.recovered = true
		switch p := p.(type) {
		case PanicError:
			// The target program explicitly called panic().
			return p.Value
		default:
			return p
		}
	}
	return nil //iface{}
}

// Interpret interprets the Go program whose main package is mainpkg.
// mode specifies various interpreter options.  filename and args are
// the initial values of os.Args for the target program.  sizes is the
// effective type-sizing function for this program.
//
// Interpret returns the exit code of the program: 2 for panic (like
// gc does), or the argument to os.Exit for normal termination.
//
// The SSA program must include the "runtime" package.
//

func NewInterp(ctx *Context, mainpkg *ssa.Package) (*Interp, error) {
	return newInterp(ctx, mainpkg, nil)
}

func newInterp(ctx *Context, mainpkg *ssa.Package, globals map[string]interface{}) (*Interp, error) {
	var rctx *reflectx.Context
	if ctx.Mode&SupportMultipleInterp == 0 {
		reflectx.ResetAll()
		rctx = reflectx.Default
		rctx.Reset()
	} else {
		rctx = reflectx.NewContext()
	}
	i := &Interp{
		ctx:          ctx,
		rctx:         rctx,
		mainpkg:      mainpkg,
		globals:      make(map[string]value),
		chkinit:      make(map[string]bool),
		goroutines:   1,
		preloadTypes: make(map[types.Type]reflect.Type),
		funcs:        make(map[*ssa.Function]*function),
		msets:        make(map[reflect.Type](map[string]*ssa.Function)),
		chexit:       make(chan int),
		mainid:       goroutineID(),
	}
	prog := mainpkg.Prog
	if ctx.Mode&OptionLoadAllImethod != 0 {
		rctx.SetHasImethod(func(typ reflect.Type, method reflectx.Method) bool {
			return true
		})
	} else if ctx.Mode&OptionLoadRutimeImethod != 0 {
		rtyps := make(map[string]bool)
		for _, typ := range prog.RuntimeTypes() {
		retry:
			switch t := typ.(type) {
			case *types.Named:
				rtyps[t.String()] = true
			case *typesalias.Alias:
				typ = typesalias.Unalias(t)
				goto retry
			}
		}
		rctx.SetHasImethod(func(typ reflect.Type, method reflectx.Method) bool {
			if method.Type == tyUnusedFunc {
				return false
			}
			if _, ok := rtyps[typ.String()]; ok {
				if ast.IsExported(method.Name) {
					return true
				}
				if typ.PkgPath() != method.PkgPath {
					return true
				}
			}
			return false
		})
	} else {
		rctx.SetHasImethod(func(typ reflect.Type, method reflectx.Method) bool {
			if method.Type == tyUnusedFunc {
				return false
			}
			if ast.IsExported(method.Name) {
				return true
			}
			if typ.PkgPath() != method.PkgPath {
				return true
			}
			return false
		})
	}

	i.record = NewTypesRecord(ctx, rctx, ctx.Loader, i, ctx.nestedMap)
	i.record.Load(mainpkg)

	var pkgs []*ssa.Package

	for _, pkg := range prog.AllPackages() {
		// skip external pkg
		if pkg.Func("init").Blocks == nil {
			continue
		}
		pkgs = append(pkgs, pkg)
		// Initialize global storage.
		for _, m := range pkg.Members {
			switch v := m.(type) {
			case *ssa.Global:
				typ := i.preToType(deref(v.Type()))
				key := v.String()
				if ext, ok := findExternValue(i, key); ok && ext.Kind() == reflect.Ptr && ext.Elem().Type() == typ {
					i.globals[key] = ext.Interface()
					i.chkinit[key] = true
				} else {
					i.globals[key] = reflect.New(typ).Interface()
				}
			}
		}
	}
	// check linkname var
	var links []*load.LinkSym
	for _, sp := range i.ctx.pkgs {
		for _, link := range sp.Links {
			if link.Kind == ast.Var {
				localName, targetName := link.PkgPath+"."+link.Name, link.Linkname.PkgPath+"."+link.Linkname.Name
				if _, ok := i.globals[localName]; ok {
					if ext, ok := findExternVar(i, link.Linkname.PkgPath, link.Linkname.Name); ok && ext.Kind() == reflect.Ptr {
						i.globals[localName] = ext.Interface()
						i.chkinit[targetName] = true
						links = append(links, link)
					} else if v, ok := i.globals[targetName]; ok {
						i.globals[localName] = v
						links = append(links, link)
					}
				}
			}
		}
	}
	// check globals for repl
	if globals != nil {
		for k := range i.globals {
			if fv, ok := globals[k]; ok {
				i.globals[k] = fv
			}
		}
	}
	// static types check
	err := checkPackages(i, pkgs)
	if err != nil {
		return i, err
	}
	// check linkname duplicated definition
	for _, link := range links {
		localName, targetName := link.PkgPath+"."+link.Name, link.Linkname.PkgPath+"."+link.Linkname.Name
		if i.chkinit[localName] && i.chkinit[targetName] {
			return i, fmt.Errorf("duplicated definition of symbol %v, from %v and %v", targetName, link.PkgPath, link.Linkname.PkgPath)
		}
	}
	return i, err
}

func (i *Interp) loadType(typ types.Type) {
	if _, ok := i.preloadTypes[typ]; !ok {
		switch typ.String() {
		case "func() *deferStack":
			if sig, ok := typ.(*types.Signature); ok && sig.Params() == nil && sig.Results().Len() == 1 {
				if ptr, ok := sig.Results().At(0).Type().(*types.Pointer); ok {
					if elem := ptr.Elem(); elem.String() == "deferStack" {
						pt := reflect.TypeOf((*_defer)(nil))
						i.preloadTypes[elem] = pt
						i.preloadTypes[ptr] = reflect.PtrTo(pt)
						rt := reflect.FuncOf(nil, []reflect.Type{reflect.PtrTo(pt)}, false)
						i.preloadTypes[typ] = rt
						return
					}
				}
			}
		case "**deferStack":
			if pptr, ok := typ.(*types.Pointer); ok {
				if ptr, ok := pptr.Elem().(*types.Pointer); ok {
					if elem := ptr.Elem(); elem.String() == "deferStack" {
						pt := reflect.TypeOf((*_defer)(nil))
						i.preloadTypes[elem] = pt
						i.preloadTypes[ptr] = reflect.PointerTo(pt)
						i.preloadTypes[typ] = reflect.PtrTo(reflect.PointerTo(pt))
						return
					}
				}
			}
		}
		rt, nested := i.record.ToType(typ)
		if nested {
			return
		}
		i.preloadTypes[typ] = rt
	}
}

func (i *Interp) preToType(typ types.Type) reflect.Type {
	if t, ok := i.preloadTypes[typ]; ok {
		return t
	}
	rt, nested := i.record.ToType(typ)
	if !nested {
		i.preloadTypes[typ] = rt
	}
	return rt
}

func (i *Interp) toType(typ types.Type) reflect.Type {
	if t, ok := i.preloadTypes[typ]; ok {
		return t
	}
	// log.Panicf("toType %v %p\n", typ, typ)
	i.typesMutex.Lock()
	defer i.typesMutex.Unlock()
	rt, _ := i.record.ToType(typ)
	return rt
}

func (i *Interp) RunFunc(name string, args ...Value) (r Value, err error) {
	fr := &frame{interp: i}
	defer func() {
		if i.ctx.Mode&DisableRecover != 0 {
			return
		}
		switch p := recover().(type) {
		case nil:
			// nothing
		case exitPanic:
			i.exitCode = int(p)
			atomic.StoreInt32(&i.exited, 1)
		case goexitPanic:
			// check goroutines
			if atomic.LoadInt32(&i.goroutines) == 1 {
				err = ErrGoexitDeadlock
			} else {
				i.exitCode = <-i.chexit
				atomic.StoreInt32(&i.exited, 1)
			}
		case PanicError:
			err = p
		default:
			// runtimeError / plainError ...
			pfr := fr
			for pfr.callee != nil {
				pfr = pfr.callee
			}
			err = FatalError{stack: debugStack(pfr), Value: p}
			if i.ctx.panicFunc != nil {
				i.ctx.handlePanic(fr, i.mainpkg.Func(name), err)
			}
		}
	}()
	if fn := i.mainpkg.Func(name); fn != nil {
		r = i.call(fr, fn, args, nil)
	} else {
		err = fmt.Errorf("no function %v", name)
	}
	return
}

func (i *Interp) ExitCode() int {
	return i.exitCode
}

func (i *Interp) RunInit() (err error) {
	i.goexited = 0
	i.exitCode = 0
	atomic.StoreInt32(&i.exited, 0)
	_, err = i.RunFunc("init")
	return
}

// ResetAllIcall is reset all reflectx icall, all interp methods invalid.
func ResetAllIcall() {
	reflectx.ResetAll()
}

// IcallStat return reflectx icall allocate stat
func IcallStat() (capacity int, allocate int, aviable int) {
	return reflectx.IcallStat()
}

// IcallCached return reflectx icall cached
func IcallCached() int {
	return reflectx.IcallCached()
}

// icall allocate
func (i *Interp) IcallAlloc() int {
	return i.rctx.IcallAlloc()
}

// ResetIcall is reset reflectx icall, all methods invalid.
func (i *Interp) ResetIcall() {
	i.rctx.Reset()
}

// UnsafeRelease is unsafe release interp. interp all invalid.
func (i *Interp) UnsafeRelease() {
	if i.ctx.Mode&SupportMultipleInterp != 0 {
		i.rctx.Reset()
	}
	i.record.Release()
	for _, v := range i.funcs {
		v.UnsafeRelease()
	}
	i.funcs = nil
	i.msets = nil
	i.globals = nil
	i.preloadTypes = nil
	i.record = nil
}

func (i *Interp) Abort() {
	atomic.StoreInt32(&i.exited, 1)
}

func (i *Interp) RunMain() (exitCode int, err error) {
	if atomic.LoadInt32(&i.exited) == 1 {
		return i.exitCode, nil
	}
	_, err = i.RunFunc("main")
	if err != nil {
		exitCode = 2
	}
	if atomic.LoadInt32(&i.exited) == 1 {
		exitCode = i.exitCode
	}
	return
}

func (i *Interp) GetFunc(key string) (interface{}, bool) {
	m, ok := i.mainpkg.Members[key]
	if !ok {
		return nil, false
	}
	fn, ok := m.(*ssa.Function)
	if !ok {
		return nil, false
	}
	return i.funcs[fn].makeFunction(i.toType(fn.Type()), nil).Interface(), true
}

func (i *Interp) GetVarAddr(key string) (interface{}, bool) {
	m, ok := i.mainpkg.Members[key]
	if !ok {
		return nil, false
	}
	v, ok := m.(*ssa.Global)
	if !ok {
		return nil, false
	}
	p, ok := i.globals[v.String()]
	return p, ok
}

func (i *Interp) GetConst(key string) (constant.Value, bool) {
	m, ok := i.mainpkg.Members[key]
	if !ok {
		return nil, false
	}
	v, ok := m.(*ssa.NamedConst)
	if !ok {
		return nil, false
	}
	return v.Value.Value, true
}

func (i *Interp) GetType(key string) (reflect.Type, bool) {
	m, ok := i.mainpkg.Members[key]
	if !ok {
		return nil, false
	}
	t, ok := m.(*ssa.Type)
	if !ok {
		return nil, false
	}
	return i.toType(t.Type()), true
}

func (i *Interp) GetSymbol(key string) (m ssa.Member, v interface{}, ok bool) {
	defer func() {
		if v := recover(); v != nil {
			ok = false
		}
	}()
	ar := strings.Split(key, ".")
	var pkg *ssa.Package
	switch len(ar) {
	case 1:
		pkg = i.mainpkg
	case 2:
		pkgs := i.mainpkg.Prog.AllPackages()
		for _, p := range pkgs {
			if p.Pkg.Path() == ar[0] || p.Pkg.Name() == ar[0] {
				pkg = p
				break
			}
		}
		if pkg == nil {
			return
		}
		key = ar[1]
	default:
		return
	}
	m, ok = pkg.Members[key]
	if !ok {
		return
	}
	switch p := m.(type) {
	case *ssa.NamedConst:
		v = p.Value.Value
	case *ssa.Global:
		v, ok = globalToValue(i, p)
	case *ssa.Function:
		typ := i.toType(p.Type())
		v = i.funcs[p].makeFunction(typ, nil)
	case *ssa.Type:
		v = i.toType(p.Type())
	}
	return
}

func (i *Interp) Exit(code int) {
	if i != nil && atomic.LoadInt32(&i.goexited) == 1 {
		i.chexit <- code
	} else {
		panic(exitPanic(code))
	}
}

// deref returns a pointer's element type; otherwise it returns typ.
// TODO(adonovan): Import from ssa?
func deref(typ types.Type) types.Type {
	if p, ok := typ.Underlying().(*types.Pointer); ok {
		return p.Elem()
	}
	return typ
}

func goroutineID() uint64 {
	return routine.Goid()
}
