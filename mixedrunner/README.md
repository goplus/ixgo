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
including Phi scratch space, and supports up to eight results. When enabled, the homogeneous
runner remains the first choice and retains its existing memory/call support.

## Registration and interpreter integration

`ixgo.RegisterMixedRunner(factory ixgo.TypedRunnerFactory)` replaces the mixed
factory; passing nil disables only mixed. The existing
`ixgo.RegisterTypedRunnerFactory` controls the homogeneous slot independently.
Register before compiling interpreters: registration is process-global, is not
safe concurrently with compilation, and does not alter already compiled functions.

Memo-selected functions and closed graphs containing memo-selected callees are
excluded before either factory is called. Function hooks, tracing, evaluation
callbacks and debug callbacks disable both runners through the same installation
path. See [memo](../memo/README.md) for memo safety rules.

`mixedrunner.MakeRunner` remains available for direct construction. It returns
nil for unsupported functions and does not install memoization or interpreter
hooks. Its existing implementation and limits are unchanged.
