//go:build instruction

package ixgo

import (
	"go/constant"
	"go/token"
	"go/types"
	"reflect"

	"golang.org/x/tools/go/ssa"
)

const instructionSupported = true

// instructionOpcode identifies an execution instruction independently of its
// SSA representation. The closure backend remains the executor while this
// description is introduced as the migration boundary toward bytecode.
type instructionOpcode uint16

const (
	opInvalid instructionOpcode = iota
	opAlloc
	opPhi
	opCall
	opBinOp
	opUnOp
	opChangeInterface
	opChangeType
	opConvert
	opMakeInterface
	opMakeClosure
	opMakeChan
	opMakeMap
	opMakeSlice
	opSlice
	opFieldAddr
	opField
	opIndexAddr
	opIndex
	opLookup
	opSelect
	opSliceToArrayPointer
	opRange
	opNext
	opTypeAssert
	opExtract
	opJump
	opIf
	opReturn
	opRunDefers
	opPanic
	opGo
	opDefer
	opSend
	opStore
	opMapUpdate
	opDebugRef
)

// instructionAux stores metadata whose shape depends on the opcode. Keeping
// it typed makes each instruction's execution inputs explicit before encoding.
type instructionAux interface {
	instructionAux()
}

// binaryOpAux records the operator, operand type and constant specialization.
type binaryOpAux struct {
	Op        token.Token
	Type      reflect.Type
	RightType reflect.Type
	Constant  [2]bool
	Values    [2]value
}

func (binaryOpAux) instructionAux() {}

// unaryOpAux records the unary operator; CommaOk requests the receive-ok result.
type unaryOpAux struct {
	Op       token.Token
	CommaOk  bool
	Type     reflect.Type
	Constant bool
	Global   bool
	Value    value
}

func (unaryOpAux) instructionAux() {}

// phiAux describes the parallel assignment performed by one grouped closure.
// Edges[i][j] supplies Results[i] when control arrives from Preds[j].
type phiAux struct {
	Preds   []int32
	Results []register
	Edges   [][]register
}

func (phiAux) instructionAux() {}

// jumpAux records the successor block ID.
type jumpAux struct{ Target int32 }

func (jumpAux) instructionAux() {}

// ifAux records true/false successor block IDs and condition specialization.
type ifAux struct {
	Targets   [2]int32
	Constant  bool
	Condition bool
	BasicBool bool
}

func (ifAux) instructionAux() {}

// returnAux records the number of returned values, whose registers are in Args.
type returnAux struct{ Results int }

func (returnAux) instructionAux() {}

// extractAux selects a zero-based element from a tuple operand.
type extractAux struct{ Index int }

func (extractAux) instructionAux() {}

// nextAux distinguishes string iteration from map iteration.
type nextAux struct{ IsString bool }

func (nextAux) instructionAux() {}

// fieldAux.Source is the operand type, a pointer for FieldAddr. Field and
// Offset describe the field within the struct, after dereferencing Source;
// Offset is relative to the start of that struct.
type fieldAux struct {
	Index  int
	Source reflect.Type
	Field  reflect.Type
	Offset uintptr
}

func (fieldAux) instructionAux() {}

// lookupAux records map key/element types. For a string Source, Key is the
// index operand type and Element is byte.
type lookupAux struct {
	Source  reflect.Type
	Key     reflect.Type
	Element reflect.Type
	CommaOk bool
}

func (lookupAux) instructionAux() {}

// indexAux.Source retains fixed array lengths; slice/string lengths and index
// bounds are checked against the runtime value. Target is the produced type.
type indexAux struct {
	Source reflect.Type
	Target reflect.Type
}

func (indexAux) instructionAux() {}

// sliceAux records source/target types and whether each boundary is present.
// Args are source, low, high, max, including -1 for omitted bounds.
// MakeSliceCheck preserves the checks used when slicing an allocation.
type sliceAux struct {
	Source         reflect.Type
	Target         reflect.Type
	Low, High, Max bool
	MakeSliceCheck bool
}

func (sliceAux) instructionAux() {}

// rangeAux records map iterator types. For a string Source, Key and Element
// are int and rune, respectively.
type rangeAux struct {
	Source  reflect.Type
	Key     reflect.Type
	Element reflect.Type
}

func (rangeAux) instructionAux() {}

// typeAssertAux records source/target types and the comma-ok result shape.
// Reflect types are the in-memory representation; a portable type table remains
// future work. These descriptions retain no SSA instruction or value.
type typeAssertAux struct {
	Source  reflect.Type
	Target  reflect.Type
	CommaOk bool
}

func (typeAssertAux) instructionAux() {}

// conversionAux records in-memory types, independently of the SSA values.
// Conversion behavior is selected by the instruction opcode and these types.
type conversionAux struct {
	Source reflect.Type
	Target reflect.Type
}

func (conversionAux) instructionAux() {}

// creationAux.Target is the produced type (a pointer for Alloc). Heap is only
// meaningful for Alloc; sizes and capacities retain their SSA order in Args.
type creationAux struct {
	Target reflect.Type
	Heap   bool
}

func (creationAux) instructionAux() {}

// selectStateAux describes one channel operation in a Select instruction.
type selectStateAux struct {
	Chan reflect.Type
	Send reflect.Type
	Dir  reflect.SelectDir
}

// selectAux preserves case order; Blocking is false when a default arm exists.
type selectAux struct {
	States   []selectStateAux
	Blocking bool
}

func (selectAux) instructionAux() {}

// sendAux records the channel and sent value types for Send.
type sendAux struct {
	Chan  reflect.Type
	Value reflect.Type
}

func (sendAux) instructionAux() {}

// storeAux records the destination pointer and stored value types for Store.
type storeAux struct {
	Address reflect.Type
	Value   reflect.Type
}

func (storeAux) instructionAux() {}

// mapUpdateAux records the map, key and stored value types for MapUpdate.
type mapUpdateAux struct {
	Map   reflect.Type
	Key   reflect.Type
	Value reflect.Type
}

func (mapUpdateAux) instructionAux() {}

// panicAux records the type of the interface operand passed to Panic.
type panicAux struct{ Value reflect.Type }

func (panicAux) instructionAux() {}

// closureAux records the runtime body, signature and capture order without SSA.
// A binding of -1 means that no register was allocated for that input.
type closureAux struct {
	Body      *function
	Signature reflect.Type
	Bindings  []register
}

func (closureAux) instructionAux() {}

// operandKind classifies the source of an input independently from its register.
// operandFunction and operandBuiltin may still have a valid Args register; only
// operandAbsent and operandUnallocated imply Args == -1.
type operandKind uint8

const (
	operandAbsent operandKind = iota
	operandRegister
	operandFunction
	operandBuiltin
	operandUnallocated
)

// callKind distinguishes dynamic, static, builtin and interface invocation targets.
type callKind uint8

const (
	callDynamic callKind = iota
	callStatic
	callBuiltin
	callInvoke
)

// callSiteAux records common call metadata. Target is
// the function-value register or invoke receiver register, or -1 if absent.
// Arguments exclude the invoke receiver; static method arguments and Signature
// both include the receiver as their first input.
// Builtins have no Signature because they have no representable Go function
// signature.
type callSiteAux struct {
	// DeferStack is the enclosing defer-stack register, or -1 for the current
	// frame or an unavailable token. An unavailable token retains lowering.
	DeferStack register
	Kind       callKind
	Signature  reflect.Type
	Target     register
	Arguments  []register
	// Name is the builtin name for callBuiltin or the method ID for callInvoke.
	// Method IDs include package identity for unexported names; execution method
	// sets currently dispatch on the plain method name. Empty for other kinds.
	Name string
}

func (callSiteAux) instructionAux() {}

// callAux adds result placement to common call metadata. A nil ExtractResults
// uses the ordinary result register; a -1 entry denotes an unused result.
// Unresolved targets and invoke calls retain this base auxiliary data.
// Dynamic calls with a target register use dynamicCallAux.
type callAux struct {
	callSiteAux
	ExtractResults []register
}

// directCallAux retains the adapter selected by lowering.
type directCallAux struct {
	callAux
	Adapter DirectCallAdapter
}

func (directCallAux) instructionAux() {}

// externalCallAux retains the callee register value allocated by lowering.
type externalCallAux struct {
	callAux
	External reflect.Value
}

func (externalCallAux) instructionAux() {}

// dynamicCallAux reads a function value from the target register at execution.
type dynamicCallAux struct {
	callAux
}

// builtinCallAux retains a builtin target and its interface argument layout.
type builtinCallAux struct {
	callAux
	InterfaceArgs []bool
	Constant      uintptr
	Error         string
}

func (builtinCallAux) instructionAux() {}

func (dynamicCallAux) instructionAux() {}

// bodyCallAux retains the interpreted body loaded by lowering.
type bodyCallAux struct {
	callAux
	Body *function
}

func (bodyCallAux) instructionAux() {}

// typedInstruction records partial execution metadata alongside the closure
// backend. Each entry corresponds to one generated closure; Phi auxiliary data
// covers the entire group; the common Result and Args fields still refer to
// the first Phi entry. Args preserve SSA operand order; ArgKinds distinguishes
// the source class independently from whether Args has a valid register.
// Opcode-specific lowering and static target resolution are incremental.
type typedInstruction struct {
	Op       instructionOpcode
	Result   register
	Args     []register
	ArgKinds []operandKind
	Aux      instructionAux
}

func descriptionRegister(pfn *function, v ssa.Value) register {
	if v != nil {
		if encoded, ok := pfn.index[v]; ok {
			return register(encoded & 0xffffff)
		}
	}
	return -1
}

// descriptionOperand reads the allocation without changing the register map.
func descriptionOperand(pfn *function, v ssa.Value) (kind, value) {
	encoded, ok := pfn.index[v]
	if !ok {
		return kindInvalid, nil
	}
	r := register(encoded & 0xffffff)
	if int(r) >= len(pfn.stack) {
		return kindInvalid, nil
	}
	return kind(encoded >> 30), pfn.stack[r]
}

func describeInstruction(pfn *function, instr ssa.Instruction, operands []*ssa.Value, elidedResults []register) typedInstruction {
	d := typedInstruction{Result: -1}
	for _, operand := range operands {
		if operand != nil {
			r := descriptionRegister(pfn, *operand)
			kind := operandRegister
			switch (*operand).(type) {
			case nil:
				kind = operandAbsent
			case *ssa.Function:
				kind = operandFunction
			case *ssa.Builtin:
				kind = operandBuiltin
			default:
				if r < 0 {
					kind = operandUnallocated
				}
			}
			d.Args = append(d.Args, r)
			d.ArgKinds = append(d.ArgKinds, kind)
		}
	}
	if v, ok := instr.(ssa.Value); ok {
		d.Result = descriptionRegister(pfn, v)
	}
	switch v := instr.(type) {
	case *ssa.BinOp:
		d.Op = opBinOp
		aux := binaryOpAux{Op: v.Op, Type: pfn.Interp.preToType(v.X.Type()), RightType: pfn.Interp.preToType(v.Y.Type())}
		for j, operand := range []ssa.Value{v.X, v.Y} {
			k, v := descriptionOperand(pfn, operand)
			if k == kindConst {
				aux.Constant[j] = true
				aux.Values[j] = v
			}
		}
		d.Aux = aux
	case *ssa.UnOp:
		d.Op = opUnOp
		aux := unaryOpAux{Op: v.Op, CommaOk: v.CommaOk, Type: pfn.Interp.preToType(v.X.Type())}
		k, value := descriptionOperand(pfn, v.X)
		aux.Constant, aux.Global = k == kindConst, k == kindGlobal
		if aux.Constant || aux.Global {
			aux.Value = value
		}
		d.Aux = aux
	case *ssa.Phi:
		d.Op = opPhi
		aux := phiAux{}
		if block := v.Block(); block != nil {
			aux.Preds = make([]int32, len(block.Preds))
			for i, pred := range block.Preds {
				aux.Preds[i] = int32(pred.Index)
			}
			for _, instr := range block.Instrs {
				phi, ok := instr.(*ssa.Phi)
				if !ok {
					break
				}
				aux.Results = append(aux.Results, descriptionRegister(pfn, phi))
				edges := make([]register, len(phi.Edges))
				for i, edge := range phi.Edges {
					edges[i] = descriptionRegister(pfn, edge)
				}
				aux.Edges = append(aux.Edges, edges)
			}
		}
		d.Aux = aux
	case *ssa.MakeClosure:
		d.Op = opMakeClosure
		fn := v.Fn.(*ssa.Function)
		signature := pfn.Interp.preToType(fn.Type())
		bindings := make([]register, len(v.Bindings))
		for i, binding := range v.Bindings {
			bindings[i] = descriptionRegister(pfn, binding)
		}
		d.Aux = closureAux{Body: pfn.Interp.funcs[fn], Signature: signature, Bindings: bindings}
	case *ssa.Call:
		d.Op = opCall
		if elidedResults != nil {
			d.Result = -1
		}
		d.Aux = describeCall(pfn, &v.Call, elidedResults, nil)
	case *ssa.Go:
		d.Op = opGo
		d.Aux = describeCall(pfn, &v.Call, nil, nil)
	case *ssa.Defer:
		d.Op = opDefer
		d.Aux = describeCall(pfn, &v.Call, nil, v.DeferStack)
	case *ssa.ChangeType:
		d.Op = opChangeType
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.Convert:
		d.Op = opConvert
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.ChangeInterface:
		d.Op = opChangeInterface
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.MakeInterface:
		d.Op = opMakeInterface
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.SliceToArrayPointer:
		d.Op = opSliceToArrayPointer
		d.Aux = describeConversion(pfn, v.X, v.Type())
	case *ssa.Alloc:
		d.Op = opAlloc
		d.Aux = creationAux{Target: pfn.Interp.preToType(v.Type()), Heap: v.Heap}
	case *ssa.MakeSlice:
		d.Op = opMakeSlice
		d.Aux = creationAux{Target: pfn.Interp.preToType(v.Type())}
	case *ssa.MakeMap:
		d.Op = opMakeMap
		d.Aux = creationAux{Target: pfn.Interp.preToType(v.Type())}
	case *ssa.MakeChan:
		d.Op = opMakeChan
		d.Aux = creationAux{Target: pfn.Interp.preToType(v.Type())}
	case *ssa.Field:
		d.Op = opField
		d.Aux = describeField(pfn, v.X, v.Field)
	case *ssa.FieldAddr:
		d.Op = opFieldAddr
		d.Aux = describeField(pfn, v.X, v.Field)
	case *ssa.Lookup:
		d.Op = opLookup
		source := pfn.Interp.preToType(v.X.Type())
		aux := lookupAux{Source: source, CommaOk: v.CommaOk}
		if source.Kind() == reflect.Map {
			aux.Key, aux.Element = source.Key(), source.Elem()
		} else {
			aux.Key, aux.Element = pfn.Interp.preToType(v.Index.Type()), reflect.TypeOf(byte(0))
		}
		d.Aux = aux
	case *ssa.Index:
		d.Op = opIndex
		d.Aux = indexAux{Source: pfn.Interp.preToType(v.X.Type()), Target: pfn.Interp.preToType(v.Type())}
	case *ssa.IndexAddr:
		d.Op = opIndexAddr
		d.Aux = indexAux{Source: pfn.Interp.preToType(v.X.Type()), Target: pfn.Interp.preToType(v.Type())}
	case *ssa.Slice:
		d.Op = opSlice
		_, allocation := v.X.(*ssa.Alloc)
		d.Aux = sliceAux{Source: pfn.Interp.preToType(v.X.Type()), Target: pfn.Interp.preToType(v.Type()), Low: v.Low != nil, High: v.High != nil, Max: v.Max != nil, MakeSliceCheck: allocation}
	case *ssa.Range:
		d.Op = opRange
		source := pfn.Interp.preToType(v.X.Type())
		aux := rangeAux{Source: source}
		if source.Kind() == reflect.Map {
			aux.Key, aux.Element = source.Key(), source.Elem()
		} else {
			aux.Key, aux.Element = reflect.TypeOf(int(0)), reflect.TypeOf(rune(0))
		}
		d.Aux = aux
	case *ssa.Select:
		d.Op = opSelect
		aux := selectAux{Blocking: v.Blocking, States: make([]selectStateAux, len(v.States))}
		for i, state := range v.States {
			chanType := pfn.Interp.preToType(state.Chan.Type())
			dir := reflect.SelectRecv
			if state.Dir == types.SendOnly {
				dir = reflect.SelectSend
			}
			var sendType reflect.Type
			if state.Send != nil {
				sendType = pfn.Interp.preToType(state.Send.Type())
			}
			aux.States[i] = selectStateAux{Chan: chanType, Send: sendType, Dir: dir}
		}
		d.Aux = aux
	case *ssa.Send:
		d.Op = opSend
		d.Aux = sendAux{Chan: pfn.Interp.preToType(v.Chan.Type()), Value: pfn.Interp.preToType(v.X.Type())}
	case *ssa.Store:
		d.Op = opStore
		d.Aux = storeAux{Address: pfn.Interp.preToType(v.Addr.Type()), Value: pfn.Interp.preToType(v.Val.Type())}
	case *ssa.MapUpdate:
		d.Op = opMapUpdate
		d.Aux = mapUpdateAux{Map: pfn.Interp.preToType(v.Map.Type()), Key: pfn.Interp.preToType(v.Key.Type()), Value: pfn.Interp.preToType(v.Value.Type())}
	case *ssa.Panic:
		d.Op = opPanic
		d.Aux = panicAux{Value: pfn.Interp.preToType(v.X.Type())}
	case *ssa.TypeAssert:
		d.Op = opTypeAssert
		d.Aux = typeAssertAux{
			Source:  pfn.Interp.preToType(v.X.Type()),
			Target:  pfn.Interp.preToType(v.AssertedType),
			CommaOk: v.CommaOk,
		}
	case *ssa.Extract:
		d.Op = opExtract
		d.Aux = extractAux{Index: v.Index}
	case *ssa.Next:
		d.Op = opNext
		d.Aux = nextAux{IsString: v.IsString}
	case *ssa.Return:
		d.Op = opReturn
		d.Aux = returnAux{Results: len(v.Results)}
	case *ssa.If:
		d.Op = opIf
		aux := ifAux{}
		if c, ok := v.Cond.(*ssa.Const); ok {
			aux.Constant, aux.Condition = true, constant.BoolVal(c.Value)
		}
		_, aux.BasicBool = v.Cond.Type().(*types.Basic)
		if block := v.Block(); block != nil && len(block.Succs) >= 2 {
			aux.Targets = [2]int32{int32(block.Succs[0].Index), int32(block.Succs[1].Index)}
		}
		d.Aux = aux
	case *ssa.Jump:
		d.Op = opJump
		aux := jumpAux{}
		if block := v.Block(); block != nil && len(block.Succs) >= 1 {
			aux.Target = int32(block.Succs[0].Index)
		}
		d.Aux = aux
	case *ssa.RunDefers:
		d.Op = opRunDefers
	case *ssa.DebugRef:
		d.Op = opDebugRef
	}
	return d
}

func describeCallSite(pfn *function, call *ssa.CallCommon) callSiteAux {
	aux := callSiteAux{DeferStack: -1, Kind: callDynamic, Target: descriptionRegister(pfn, call.Value), Arguments: make([]register, len(call.Args))}
	for i, arg := range call.Args {
		aux.Arguments[i] = descriptionRegister(pfn, arg)
	}
	if builtin, ok := call.Value.(*ssa.Builtin); ok {
		aux.Kind, aux.Name = callBuiltin, builtin.Name()
		return aux
	}
	sig := call.Signature()
	if !call.IsInvoke() && sig.Recv() != nil {
		sig = wrapMethodType(sig)
	}
	aux.Signature = pfn.Interp.preToType(sig)
	if call.IsInvoke() {
		aux.Kind, aux.Name = callInvoke, call.Method.Id()
	} else if _, ok := call.Value.(*ssa.Function); ok {
		aux.Kind = callStatic
	}
	return aux
}

// describeCall reuses lowering-selected targets, including adapters; dynamic
// function values use dynamicCallAux when their target register is available.
func describeCall(pfn *function, call *ssa.CallCommon, dests []register, deferStack ssa.Value) instructionAux {
	base := callAux{callSiteAux: describeCallSite(pfn, call), ExtractResults: dests}
	if deferStack != nil {
		base.DeferStack = descriptionRegister(pfn, deferStack)
		if base.DeferStack < 0 {
			// Keep lowering when the enclosing defer-stack token is unavailable.
			return base
		}
	}
	// Immediate closures have dedicated binding paths in lowering for Call,
	// Go and Defer; keep them until those environments are described directly.
	if _, immediate := call.Value.(*ssa.MakeClosure); immediate {
		return base
	}
	if base.Kind == callBuiltin {
		name := call.Value.(*ssa.Builtin).Name()
		aux := builtinCallAux{callAux: base, InterfaceArgs: builtinInterfaceArgs(pfn.Interp, name, call.Args)}
		switch name {
		case "Sizeof", "Alignof":
			typ := call.Args[0].Type()
			if len(call.Args) == 2 {
				typ = typ.Underlying().(*types.Pointer).Elem()
			}
			if name == "Sizeof" {
				aux.Constant = uintptr(pfn.Interp.ctx.sizes.Sizeof(typ))
			} else {
				aux.Constant = uintptr(pfn.Interp.ctx.sizes.Alignof(typ))
			}
		case "Offsetof":
			var offset int64
			if len(call.Args) > 1 {
				typ := call.Args[0].Type().Underlying().(*types.Pointer).Elem()
				for _, arg := range call.Args[1:] {
					index, _ := constant.Int64Val(arg.(*ssa.Const).Value)
					st := typ.Underlying().(*types.Struct)
					offset += structOffsetsof(pfn.Interp.ctx.sizes, st)[int(index)]
					typ = st.Field(int(index)).Type()
				}
			} else {
				var err error
				offset, err = builtinOffsetof(pfn.Interp.ctx, call.Pos(), call.Args)
				if err != nil {
					aux.Error = err.Error()
				}
			}
			aux.Constant = uintptr(offset)
		}
		return aux
	}
	if base.Kind == callDynamic && base.Target >= 0 {
		return dynamicCallAux{callAux: base}
	}
	if base.Kind != callStatic {
		return base
	}
	if pfn.directCall == call && pfn.directAdapter != nil {
		return directCallAux{callAux: base, Adapter: pfn.directAdapter}
	}
	fn := call.Value.(*ssa.Function)
	if fn.Blocks != nil {
		if body := pfn.Interp.funcs[fn]; body != nil {
			return bodyCallAux{callAux: base, Body: body}
		}
	} else if kind, value := descriptionOperand(pfn, fn); kind == kindFunction && value != nil {
		external := reflect.ValueOf(value)
		if external.Kind() == reflect.Func {
			return externalCallAux{callAux: base, External: external}
		}
	}
	return base
}

func describeField(pfn *function, source ssa.Value, index int) fieldAux {
	typ := pfn.Interp.preToType(source.Type())
	record := typ
	if record.Kind() == reflect.Pointer {
		record = record.Elem()
	}
	field := record.Field(index)
	return fieldAux{Index: index, Source: typ, Field: field.Type, Offset: field.Offset}
}

func describeConversion(pfn *function, source ssa.Value, target types.Type) conversionAux {
	return conversionAux{
		Source: pfn.Interp.preToType(source.Type()),
		Target: pfn.Interp.preToType(target),
	}
}
