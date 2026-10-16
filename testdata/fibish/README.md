# Ordinary-interpreter Fibonacci memoization

This directory contains six standalone Go programs, each with input 40. This
branch starts from upstream/main `72593742` and implements caching in ordinary
interpreter instruction entry/return handlers. It has no typed runner dependency.

## Execution and eligibility

Pure scalar self-recursion caches basic scalar results per top-level call tree.
The input key uses exact float32/float64 bit patterns, including signed zeros.
Additional parameters may only be forwarded unchanged. Cached values cannot
contain mutable objects. Unsupported instructions or observable effects fall back.

Integer accumulator graphs use a separate SSA proof of read-only input pointers,
local storage writes, and isolated load/add/store effects on one accumulator.
Parameter roles are inferred from SSA uses and call-argument mappings. Parameter
count and output position are unrestricted; call signatures and integer widths
may differ. Mutual recursion shares a private cache keyed by function and all
scalar input values and pointed-to input values in parameter order.
Cache hits add the saved delta. Nil inputs follow original control flow. Output
aliasing or partial overlap with input storage disables caching for that call tree.
When an additive integer proof is unavailable, numerical full-state caching
includes the output initial bits in the key and restores final bits on a hit.
This supports integer arithmetic, remainder, bit operations and shifts, and
float32/float64 arithmetic, comparisons and numeric conversions, without
regrouping operations. Only normal returns are cached. Multiple mutable outputs,
external memory effects and unsupported instructions retain ordinary execution.

Hooks, debug callbacks, tracing, and custom call interception disable memoization. Frame cleanup
releases cache references. The loop fixture retains normal loop execution.
Memoization changes repeated recursive work and therefore algorithmic cost.

## Reproduce

From the repository root:

```sh
GOWORK=off go build -o /tmp/ixgo-main-memo ./cmd/ixgo
GOWORK=off GOMAXPROCS=1 /usr/bin/time -p /tmp/ixgo-main-memo run testdata/fibish/fibish.go
GOWORK=off GOMAXPROCS=1 /usr/bin/time -p /tmp/ixgo-main-memo run testdata/fibish/uglyfib.go
GOWORK=off go test -run '^$' -bench 'BenchmarkInterpreterMemo|BenchmarkInterpreterAccumulator' -benchmem -count=3 .
```

Each source can be run independently using the same command. Expected outputs:

- fibish, closure, loop, float64: `f(40)=39088169,126491972`
- float32: `f(40)=39088168,126491968`
- uglyfib: `Fib(40)=102334155`

## Current full-process timings (numeric state update)

Measured 2026-10-16, Apple M5, macOS arm64, Go 1.27.1, GOWORK=off,
GOMAXPROCS=1. Three serial process runs per file, including startup and source
loading, excluding compilation. All 18 outputs checked. The first fibish process
had extra startup wall time; cause not isolated. The final fixture timing batch ran after builds and benchmarks completed. Small
process-time differences remain dominated by loading noise.

| File | Three rounds (s) | Median (s) |
| --- | --- | ---: |
| `fibish.go` | 1.061958, 0.052272, 0.052127 | 0.052272 |
| `fibish_closure.go` | 0.053467, 0.052436, 0.051326 | 0.052436 |
| `fibish_loop.go` | 0.052139, 0.051737, 0.050798 | 0.051737 |
| `fibish_float32.go` | 0.053521, 0.051449, 0.051082 | 0.051449 |
| `fibish_float64.go` | 0.050155, 0.050759, 0.052510 | 0.050759 |
| `uglyfib.go` | 0.051411, 0.053491, 0.052571 | 0.052571 |

Startup/loading dominates the cached examples. The unoptimized uglyfib(40)
program was deliberately not rerun. Execution-only baselines use input 20.

## Historical execution-only benchmark (before numerical full-state caching)

Three repetitions, medians. Both variants use ordinary interpreter instructions;
"baseline" disables cache lookup/storage on the same compiled functions.

| Example | Baseline | Memoized | Speedup | Memo B/op | Memo allocs/op |
| --- | ---: | ---: | ---: | ---: | ---: |
| fibish | 293.623 us | 2.267 us | 129.5x | 3,138 | 44 |
| fibish_float32 | 421.823 us | 2.356 us | 179.0x | 3,321 | 103 |
| fibish_float64 | 429.709 us | 2.325 us | 184.8x | 3,609 | 103 |
| uglyfib | 6.270 ms | 31.730 us | 197.6x | ~18,319 | 561 |

Baseline fibish uses ~779 B/op and 69 allocations; its memoized path increases
bytes while reducing execution. Float/uglyfib paths reduce repeated boxing and
allocation work. These gains include algorithmic savings from memoization.

## Tests

```sh
GOWORK=off go test -count=1 ./...
GOWORK=off go test -race -count=1 .
```

Tests also cover 2/3/5/6-parameter signatures, reordered outputs, multiple
read-only pointer inputs, mixed integer widths, and differing call arities.
Numeric conversions use full-state caching; multiple mutable outputs retain
ordinary execution. The proof accepts one mutable numerical output.

Tests cover all four accumulator entries, comparison with uncached execution,
integer overflow, unsigned narrow overflow, nil pointers/repeated panics,
unsafe partial overlap, alias fallback, rejected side effects, signed-zero
recursion, hook/debug/tracing fallback, and separate-interpreter concurrent calls.
Other OS/toolchain combinations and CI coverage remain for CI.

Local full tests and race tests passed. `go vet ./...` reports warnings in
existing export/unsafe/context-cancel code outside this change; no warnings
reference the new memoization files.

### Generalized accumulator analysis update

After removing fixed parameter-role assumptions, the input-20 accumulator
benchmark median is 34.615 us/op, approximately 19,887 B/op and 682 allocs/op
(three repetitions). The previous fixed-shape median was 31.730 us/op.
Variable-length keys include every scalar and pointed-to input value; this
adds allocation and lookup overhead. Full and race tests passed after this
update. Full-process timings above predate this generalization.

## Numerical full-state update

The state key includes output initial bits as well as all other numerical inputs.
Equal function/input/output-initial states can reuse a final value. Merely having
the same recursive input is insufficient; monotonically changing outputs may
have no cache hits and incur additional memory/time. Integer additive summaries
remain preferred where proved, since they can reuse results across initial values.

Synthetic repeated-state multiplication benchmark, recursion depth 16, three-run
medians, same ordinary interpreter with cache disabled/enabled:

| Output type | Cache disabled | Full-state cache | Cache B/op | allocs/op |
| --- | ---: | ---: | ---: | ---: |
| int64 | 7.325 ms | 3.347 us | 3,148 | 43 |
| float64 | 7.155 ms | 3.293 us | 3,156 | 44 |

The fixture deliberately resets the output before each subtree, creating repeated
states. These speedups establish functionality and are not representative of
all arithmetic workloads. Baseline timings had some run-to-run variation.

Current legacy-path medians: fibish 2.235 us, float32 2.305 us, float64 2.335 us,
uglyfib 36.912 us. Uglyfib allocation remains about 19,886 B/op and 682 allocs;
its earlier generalized median was 34.615 us, and this batch included a slow
57.815 us repetition. No speed improvement for uglyfib is claimed.

```sh
GOWORK=off go test -run '^$' -bench BenchmarkInterpreterState -benchmem -count=3 .
```

Additional tests compare numerical output bits for integer/float operations,
branches depending on output state, signed zero, infinities, NaNs, float rounding,
repeated divide/remainder-by-zero and negative-shift panics, and reject global
reads/multiple outputs. Full/race test suites passed for this extension.

State caches currently have no size bound. Memory can grow with the number
of unique visited states, including output initial states. This extension is
most useful when complete states repeat.
