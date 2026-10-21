# Mixed numeric runner

`github.com/goplus/ixgo/mixedrunner` enables the optional mixed numeric runner for
ixgo. It is useful for numeric workloads that combine types, such as integer
control variables with floating-point calculations, numeric conversions, and
numeric helper calls.

```go
import _ "github.com/goplus/ixgo/mixedrunner"
```

The `ixgo` command imports this package explicitly. Applications embedding ixgo
can enable it with the blank import above, or disable it for a specific context:

```go
optimized := ixgo.NewContext(ixgo.DisableMemo)
baseline := ixgo.NewContext(ixgo.DisableMemo | ixgo.DisableMixedRunner)
```

## Supported paths

The runner supports:

- Built-in integer, `float32`, `float64`, and `bool` values;
- Numeric conversions and independently typed shift counts;
- Branches, loops, parallel Phi assignments, and multiple results;
- Static calls, recursion, and mutual recursion;
- Pointer receivers with direct built-in numeric fields, including field loads and stores.

The analysis is conservative. Unsupported functions stay on the ordinary interpreter
path. This includes dynamic or interface calls, closures with observable environments,
arrays, slices, maps, nested structures, unsafe operations, unsupported memory
operations, non-numeric fields, and named numeric boundary types. A static call graph
is rejected when any required function is unsupported. The runner uses at most 64
registers and supports up to eight results.

Memo-selected functions and call graphs containing memo-selected callees are excluded.
Hooks, tracing, evaluation callbacks, and debug callbacks also keep functions on the
ordinary interpreter path. Import `memo` separately when memoization is needed.

## Direct construction

`mixedrunner.MakeRunner` compiles one SSA function and returns `nil` when the function
is unsupported. `RunArgs` requires the exact parameter count and exact built-in dynamic
types recorded by the compiled function; convertible values and named types are not
accepted.

```go
runner := mixedrunner.MakeRunner(fn)
if runner != nil {
	result := runner.RunArgs(args)
	_ = result
}
```

Register the plugin before compiling interpreters. Registration is process-global and
is not safe to change concurrently with compilation. Registering it does not alter
already compiled functions.

## Performance

These measurements compare ixgo's ordinary interpreter with `mixedrunner`. Memoization
is disabled in both configurations. Each value is the median of three independent
process runs with `GOMAXPROCS=1`; loading and compilation are excluded.

| Workload | Interpreter | mixedrunner | Speedup |
| --- | ---: | ---: | ---: |
| `test/abi/fibish.go` | 3.902 s | 2.433 s | 1.60x |
| `test/divmod.go` | 26.769 s | 10.473 s | 2.56x |

These are workload-specific measurements, not a general speedup guarantee. Short-lived
workloads can lose time to compilation. Unsupported parts of a program continue to use
the ordinary interpreter.

## Recursion and panics

Recursive calls use dynamically growing heap frames without a fixed depth limit. Very
deep or untrusted recursion can exhaust heap memory. Integer division or remainder by
zero and negative shifts use native Go operations so their recoverable panic values are
preserved. Functions using `defer` are not compiled by this runner.

## Tests

From the repository root:

```sh
go test ./mixedrunner
GOWORK=off go test -race ./mixedrunner ./mixedrunner/integration
```
