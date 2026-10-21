# mixedrunner

Package `mixedrunner` independently enables the optional mixed numeric runner in ixgo.

```go
import _ "github.com/goplus/ixgo/mixedrunner"
```

It does not import or enable any other runner. The CLI imports this plugin
explicitly.

The mixed runner supports built-in integer and floating-point parameters,
local values and results, booleans, numeric conversions, independently typed
shift counts, branches, loops and parallel Phi copies. It stores numeric bits
in 64-bit registers and executes each operation with its compiled operand type.
Static calls compile as a closed graph, including recursion and mutual recursion.
Arguments and multiple results (including mixed numeric/bool tuples) pass between
invocation-local frames as raw bits, without internal boxing. Any unsupported
callee rejects the entire graph; dynamic/interface calls, builtins, pointers,
arrays, closures, other memory operations and named boundary types are rejected. It uses at most 64 registers,
including Phi scratch space, and supports up to eight results. Unsupported graphs use the ordinary interpreter.

## Registration and interpreter integration

`ixgo.RegisterMixedRunner(factory ixgo.TypedRunnerFactory)` replaces the mixed
factory; passing nil disables mixed. There is only one optional runner slot.
Register before compiling interpreters: registration is process-global, is not
safe concurrently with compilation, and does not alter already compiled functions.

Memo-selected functions and closed graphs containing memo-selected callees are
excluded before the factory is called. Function hooks, tracing, evaluation
callbacks and debug callbacks disable the runner through the same installation
path. See [memo](../memo/README.md) for memo safety rules.

`mixedrunner.MakeRunner` remains available for direct construction. It returns
nil for unsupported functions and does not install memoization or interpreter
hooks. `RunArgs` requires the exact parameter count and exact built-in dynamic types;
convertible values and named types are not accepted.

## Recursion and runtime panics

Recursive calls use dynamically growing heap frames with no fixed depth limit.
This does not reproduce the native Go stack limit: runaway recursion can exhaust
heap memory rather than trigger a fatal stack overflow. A recoverable depth-limit
panic would not preserve Go runtime semantics, so no artificial limit is imposed.
Do not use this experimental runner for unbounded or untrusted recursion.
Integer division/remainder by zero and negative shifts use native Go operations
to preserve their recoverable runtime panic values; deferred recovery remains in
the interpreter because functions with defer are not compiled by this runner.
