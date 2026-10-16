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
Floating-point accumulator effects are excluded because regrouping changes rounding.

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

## Local full-process timings

Measured 2026-10-16, Apple M5, macOS arm64, Go 1.27.1, GOWORK=off,
GOMAXPROCS=1. Three serial process runs per file, including startup and source
loading, excluding compilation. All 18 outputs checked. The first fibish process
had extra startup wall time; cause not isolated. No benchmark or build ran concurrently in this final timing batch. Small
process-time differences remain dominated by loading noise.

| File | Three rounds (s) | Median (s) |
| --- | --- | ---: |
| `fibish.go` | 0.329815, 0.050441, 0.051668 | 0.051668 |
| `fibish_closure.go` | 0.052544, 0.051360, 0.052015 | 0.052015 |
| `fibish_loop.go` | 0.049367, 0.051262, 0.056108 | 0.051262 |
| `fibish_float32.go` | 0.052050, 0.051108, 0.051723 | 0.051723 |
| `fibish_float64.go` | 0.052843, 0.050789, 0.050134 | 0.050789 |
| `uglyfib.go` | 0.051207, 0.054727, 0.052299 | 0.052299 |

Startup/loading dominates the cached examples. The unoptimized uglyfib(40)
program was deliberately not rerun. Execution-only baselines use input 20.

## Historical execution-only benchmark (before generalized parameter analysis)

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
Unsupported integer conversions and multiple mutable outputs retain ordinary
execution. The proof currently accepts one additive integer output.

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
