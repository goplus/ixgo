// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25
// +build go1.25

package iter

import (
	_ "iter"
	"reflect"
	_ "unsafe"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/github.com/goplus/ixgo/x/race"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "iter",
		Path: "iter",
		Deps: map[string]string{
			"github.com/goplus/ixgo/x/race": "race",
			"runtime":                       "runtime",
			"unsafe":                        "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"coroswitch": reflect.ValueOf(_coroswitch),
			"newcoro":    reflect.ValueOf(_newcoro),
		},
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
		Source:        source,
	})
}

type coro struct{}

//go:linkname _newcoro runtime.newcoro
func _newcoro(func(*coro)) *coro

//go:linkname _coroswitch runtime.coroswitch
func _coroswitch(*coro)

var source = "package iter\n\nimport (\n\t\"runtime\"\n\t\"unsafe\"\n\t\"github.com/goplus/ixgo/x/race\"\n)\n\ntype Seq[V any] func(yield func(V) bool)\n\ntype Seq2[K, V any] func(yield func(K, V) bool)\n\ntype coro struct{}\n\nfunc newcoro(func(*coro)) *coro\n\nfunc coroswitch(*coro)\n\nfunc Pull[V any](seq Seq[V]) (next func() (V, bool), stop func()) {\n\tvar pull struct {\n\t\tv\t\tV\n\t\tok\t\tbool\n\t\tdone\t\tbool\n\t\tyieldNext\tbool\n\t\tseqDone\t\tbool\n\t\tracer\t\tint\n\t\tpanicValue\tany\n\t}\n\tc := newcoro(func(c *coro) {\n\t\trace.Acquire(unsafe.Pointer(&pull.racer))\n\t\tif pull.done {\n\t\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\t\treturn\n\t\t}\n\t\tyield := func(v1 V) bool {\n\t\t\tif pull.done {\n\t\t\t\treturn false\n\t\t\t}\n\t\t\tif !pull.yieldNext {\n\t\t\t\tpanic(\"iter.Pull: yield called again before next\")\n\t\t\t}\n\t\t\tpull.yieldNext = false\n\t\t\tpull.v, pull.ok = v1, true\n\t\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\t\tcoroswitch(c)\n\t\t\trace.Acquire(unsafe.Pointer(&pull.racer))\n\t\t\treturn !pull.done\n\t\t}\n\n\t\tdefer func() {\n\t\t\tif p := recover(); p != nil {\n\t\t\t\tpull.panicValue = p\n\t\t\t} else if !pull.seqDone {\n\t\t\t\tpull.panicValue = goexitPanicValue\n\t\t\t}\n\t\t\tpull.done = true\n\t\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\t}()\n\t\tseq(yield)\n\t\tvar v0 V\n\t\tpull.v, pull.ok = v0, false\n\t\tpull.seqDone = true\n\t})\n\tnext = func() (v1 V, ok1 bool) {\n\t\trace.Write(unsafe.Pointer(&pull.racer))\n\n\t\tif pull.done {\n\t\t\treturn\n\t\t}\n\t\tif pull.yieldNext {\n\t\t\tpanic(\"iter.Pull: next called again before yield\")\n\t\t}\n\t\tpull.yieldNext = true\n\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\tcoroswitch(c)\n\t\trace.Acquire(unsafe.Pointer(&pull.racer))\n\n\t\tif pull.panicValue != nil {\n\t\t\tif pull.panicValue == goexitPanicValue {\n\n\t\t\t\truntime.Goexit()\n\t\t\t} else {\n\t\t\t\tpanic(pull.panicValue)\n\t\t\t}\n\t\t}\n\t\treturn pull.v, pull.ok\n\t}\n\tstop = func() {\n\t\trace.Write(unsafe.Pointer(&pull.racer))\n\n\t\tif !pull.done {\n\t\t\tpull.done = true\n\t\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\t\tcoroswitch(c)\n\t\t\trace.Acquire(unsafe.Pointer(&pull.racer))\n\n\t\t\tif pull.panicValue != nil {\n\t\t\t\tif pull.panicValue == goexitPanicValue {\n\n\t\t\t\t\truntime.Goexit()\n\t\t\t\t} else {\n\t\t\t\t\tpanic(pull.panicValue)\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\treturn next, stop\n}\n\nfunc Pull2[K, V any](seq Seq2[K, V]) (next func() (K, V, bool), stop func()) {\n\tvar pull struct {\n\t\tk\t\tK\n\t\tv\t\tV\n\t\tok\t\tbool\n\t\tdone\t\tbool\n\t\tyieldNext\tbool\n\t\tseqDone\t\tbool\n\t\tracer\t\tint\n\t\tpanicValue\tany\n\t}\n\tc := newcoro(func(c *coro) {\n\t\trace.Acquire(unsafe.Pointer(&pull.racer))\n\t\tif pull.done {\n\t\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\t\treturn\n\t\t}\n\t\tyield := func(k1 K, v1 V) bool {\n\t\t\tif pull.done {\n\t\t\t\treturn false\n\t\t\t}\n\t\t\tif !pull.yieldNext {\n\t\t\t\tpanic(\"iter.Pull2: yield called again before next\")\n\t\t\t}\n\t\t\tpull.yieldNext = false\n\t\t\tpull.k, pull.v, pull.ok = k1, v1, true\n\t\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\t\tcoroswitch(c)\n\t\t\trace.Acquire(unsafe.Pointer(&pull.racer))\n\t\t\treturn !pull.done\n\t\t}\n\n\t\tdefer func() {\n\t\t\tif p := recover(); p != nil {\n\t\t\t\tpull.panicValue = p\n\t\t\t} else if !pull.seqDone {\n\t\t\t\tpull.panicValue = goexitPanicValue\n\t\t\t}\n\t\t\tpull.done = true\n\t\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\t}()\n\t\tseq(yield)\n\t\tvar k0 K\n\t\tvar v0 V\n\t\tpull.k, pull.v, pull.ok = k0, v0, false\n\t\tpull.seqDone = true\n\t})\n\tnext = func() (k1 K, v1 V, ok1 bool) {\n\t\trace.Write(unsafe.Pointer(&pull.racer))\n\n\t\tif pull.done {\n\t\t\treturn\n\t\t}\n\t\tif pull.yieldNext {\n\t\t\tpanic(\"iter.Pull2: next called again before yield\")\n\t\t}\n\t\tpull.yieldNext = true\n\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\tcoroswitch(c)\n\t\trace.Acquire(unsafe.Pointer(&pull.racer))\n\n\t\tif pull.panicValue != nil {\n\t\t\tif pull.panicValue == goexitPanicValue {\n\n\t\t\t\truntime.Goexit()\n\t\t\t} else {\n\t\t\t\tpanic(pull.panicValue)\n\t\t\t}\n\t\t}\n\t\treturn pull.k, pull.v, pull.ok\n\t}\n\tstop = func() {\n\t\trace.Write(unsafe.Pointer(&pull.racer))\n\n\t\tif !pull.done {\n\t\t\tpull.done = true\n\t\t\trace.Release(unsafe.Pointer(&pull.racer))\n\t\t\tcoroswitch(c)\n\t\t\trace.Acquire(unsafe.Pointer(&pull.racer))\n\n\t\t\tif pull.panicValue != nil {\n\t\t\t\tif pull.panicValue == goexitPanicValue {\n\n\t\t\t\t\truntime.Goexit()\n\t\t\t\t} else {\n\t\t\t\t\tpanic(pull.panicValue)\n\t\t\t\t}\n\t\t\t}\n\t\t}\n\t}\n\treturn next, stop\n}\n\nvar goexitPanicValue any = new(int)\n"
