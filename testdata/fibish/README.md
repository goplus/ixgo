# fibish comparisons

Run each Go file separately:

- `fibish.go`: the recursive Go ABI example.
- `fibish_float32.go` and `fibish_float64.go`: the same recursive computation using floating-point parameters and results.
- `fibish_closure.go`: the Go ABI example that forwards a function argument.
- `fibish_loop.go`: the same recurrence evaluated with a loop and four scalar states.
- `uglyfib.go`: the slower Go ABI example with pointers, arrays, and mutual recursion.

All Go examples use `x = 40`. The integer fibish examples and the float64
version print:

```text
f(40)=39088169,126491972
```

The float32 version prints `f(40)=39088168,126491968`, matching native Go.
The difference comes from float32 rounding once the results exceed its exact
integer range. `uglyfib.go` prints `Fib(40)=102334155`.

Build the CLI once, then measure each program:

```sh
go build -o /tmp/ixgo-fibish ./cmd/ixgo
/usr/bin/time -p /tmp/ixgo-fibish run testdata/fibish/fibish.go
/usr/bin/time -p /tmp/ixgo-fibish run testdata/fibish/fibish_float32.go
/usr/bin/time -p /tmp/ixgo-fibish run testdata/fibish/fibish_float64.go
/usr/bin/time -p /tmp/ixgo-fibish run testdata/fibish/fibish_closure.go
/usr/bin/time -p /tmp/ixgo-fibish run testdata/fibish/fibish_loop.go
/usr/bin/time -p /tmp/ixgo-fibish run testdata/fibish/uglyfib.go
```

These commands include interpreter loading time. The loop performs O(x) work;
the recursive examples perform exponentially more work, so the difference also
reflects the algorithms. Compare the same file between ixgo versions to measure
an interpreter optimization. The library requires a blank import of
`github.com/goplus/ixgo/typedrunner` to enable the optional typed runner; the CLI
already imports it.

## Floating-point runners

Measured on the same Apple M5 with Go 1.27.1, `GOMAXPROCS=1`, and input 40.
Each value is the median of three serial, alternating process runs, including
startup and loading. The baseline is `0a45aa05`, whose typedrunner plugin falls
back for floats; the float implementation is the current working tree based on
that commit. Both binaries execute the same source. Every output was checked
against the corresponding native Go binary.

| Example | Before float support | With float support | Speedup |
| --- | ---: | ---: | ---: |
| `fibish_float32.go` | 8.881 s | 2.548 s | 3.49x |
| `fibish_float64.go` | 9.428 s | 2.820 s | 3.34x |

The generic core supports one numeric type per function and call graph, with
separate bool registers. Float32 and float64 runners support arithmetic,
comparisons, negation, Phi, recursion, pointers, and local arrays. Mixed numeric
types and numeric conversions currently fall back to the ordinary interpreter.

Integer regression checks use five alternating runs at `GOMAXPROCS=1` with a
250 ms benchmark duration. Integer fibish(40) remains effectively unchanged
(2.497 s before, 2.481 s after), as do the multiargument shallow/mixed/deep
benchmarks (within 1.5%). Loop results vary: sum, branch, and nested improve;
fibloop is unchanged; the Phi swap case regresses from 8.352 to 9.019 us
(8.0%). These are measurements of the two default binary layouts. The layout
investigation below limits how they can be attributed to the float changes.
Raw timing JSON files are not checked in.

Integer-only operator checks use three alternating runs, a 300 ms duration,
and loops of 999 iterations. Each entry is the median, with zero allocations.

| Type / operator | Before | After | Time change |
| --- | ---: | ---: | ---: |
| `int/rem` | 9.048 us | 9.363 us | +3.5% |
| `int/bits` | 10.758 us | 9.279 us | -13.7% |
| `int/shl` | 8.500 us | 8.549 us | +0.6% |
| `int/shr` | 9.666 us | 10.391 us | +7.5% |
| `uint64/rem` | 8.878 us | 8.691 us | -2.1% |
| `uint64/bits` | 9.904 us | 10.734 us | +8.4% |
| `uint64/shl` | 7.311 us | 7.911 us | +8.2% |
| `uint64/shr` | 9.061 us | 8.933 us | -1.4% |

The default binaries show mixed effects, including up to 8.4% slower cases.
Further investigation identifies a substantial sensitivity to linked code and
data layout rather than extra arithmetic instructions for int/uint64:

- The int scalar execution loop remains 2752 bytes; uint64 remains 2720 bytes.
  The remainder and shift helpers are inlined. After normalizing relocations,
  the int loop changes only the generic dictionary offset used by recursive
  calls. Its nonrecursive operator paths have the same instructions. The
  uint64 remainder uses a different temporary register, with the same instruction
  count; the other nonrecursive paths are unchanged.
- A temporary diagnostic binary runs both the original integer loop and the
  float-capable loop against the exact same compiled instruction slice. Three
  alternating process runs at `GOMAXPROCS=1`, 200 ms per benchmark, show the
  following median results when only linker alignment is changed:

| Same-process swap diagnostic | Original loop | Float-capable loop | Time change |
| --- | ---: | ---: | ---: |
| Default function layout | 7.568 us | 8.579 us | +13.4% |
| `-ldflags=-funcalign=64` | 8.524 us | 7.141 us | -16.2% |

The separate original/new binaries with the same 64-byte function alignment
also change the swap result to 8.216 vs 8.122 us (-1.1%). Other benchmarks can
move in the opposite direction under this alignment, so it is a diagnostic
control rather than a general performance fix. These tests establish layout
sensitivity; they do not isolate instruction-cache effects from branch prediction
or other hardware effects.

Narrow integer code does have a separate implementation difference: uint8
remainder changes from 32-bit to 64-bit division, and its shift-count comparisons
also widen. The int/uint64 benchmarks above do not measure that cost. Supporting
floats does not add a floating-point type check or an operation callback to each
integer instruction.

## CLI and language comparison

Measured on Apple M5, macOS arm64, with input 40. Each entry is the median
of three serial process runs, including startup and loading, without a separate
warmup. All outputs were checked. The Go comparison alternates main `5c99c936`
and typed runner `59f19132`, both built with Go 1.27.1. The other interpreters
were measured in separate serial batches on the same machine; they were not
interleaved with the Go measurements.

| Implementation | fibish | fibish_closure | uglyfib |
| --- | ---: | ---: | ---: |
| ixgo main `5c99c936` | 3.898 s | 6.333 s | 91.501 s |
| ixgo typed runner `59f19132` | 2.632 s | 2.587 s | 25.617 s |
| Lua 5.5.1 | 2.412 s | 2.542 s | 26.266 s |
| LuaJIT 2.1.1791890110 | 0.197 s | 0.228 s | 8.987 s |
| Python 3.9.6 | 11.804 s | 12.290 s | 31.057 s |
| Ruby 2.6.10p210 | 6.335 s | 6.591 s | 21.004 s |

Standard Lua uses its bytecode interpreter. LuaJIT uses the same `.lua` files
with its default JIT enabled (`jit.status()` returned true); no JIT tuning flags
were passed. Python and Ruby use their default interpreter settings. These
results describe the listed versions, not newer Python or Ruby releases.

The `.py`, `.lua`, and `.rb` versions preserve the recursive branches and return
two values using each language's normal multiple-result mechanism. As in the
Go source, `fibish_closure` passes the function itself as an additional argument
and forwards that argument unchanged; it does not invoke the argument or create
a capturing closure on each call. Uglyfib preserves the four mutually recursive
functions and local array sizes, using a shared mutable list/table/array's first
slot to represent `*int` and preserve aliases. Its object allocation and garbage
collection behavior differs from Go's pointer/array representation, so these are
equivalent computations rather than identical memory implementations.

Run the scripts separately (an optional argument overrides the default 40):

```sh
python3 fibish.py
ruby fibish_closure.rb
lua uglyfib.lua
luajit uglyfib.lua
```

From this directory, the helper scripts verify small inputs and repeat the full
input three times per selected implementation:

```sh
python3 bench_languages.py python ruby lua luajit
python3 bench_uglyfib.py python ruby lua luajit
```

They write `language_timings.json` and `uglyfib_timings.json` locally. These
generated timing files are not included in the repository.

The latest Go loop comparison measured 50.250 ms on main and 50.588 ms with
the typed runner. At only 40 iterations, loading dominates and the small
difference does not establish a regression.

For a loop workload dominated by execution, changing only the temporary Go
fixture's input to 100,000,000 measured 3.165 s on main `5c99c936` versus
1.147 s on typed runner `59f19132` (2.76x). These are medians of three
alternating serial CLI runs, including loading, with default GOMAXPROCS.
All runs printed the same pair. This large input intentionally overflows int64
and measures throughput; it is not a mathematical Fibonacci result. The
checked-in fixture retains input 40.

## Typed argument and memory runner

The extended generic runner now compiles uglyfib's `f/g/h/k` call graph,
including pointer parameters, local arrays, array value copies, loads and
stores. Calls execute in an arena of reusable frames. Local pointers are
frame/slot handles, preserving aliases when the arena grows; external scalar
pointers are read and written directly. The original scalar runner remains
the first choice for eligible fibish functions.

The supported closed call graph uses one integer type T, bool parameters,
scalar *T pointers and local [N]T arrays. It supports multiple actual arguments,
up to eight integer/bool results (including mixed tuples), direct calls including
mutual recursion, unary and
binary arithmetic, and scalar/bool/pointer Phi. Unsupported instructions or
callees reject the whole graph. Pointer results, external calls, defer, and
allocations inside loops fall back to the ordinary interpreter. Pure bool functions
use the argument runner with a default integer register type; bool values stay
in separate bool registers. Each frame is
limited to 64 registers and 64 local scalar slots; recursive depth grows the
arena as needed. The runner does not use unsafe pointers or hardcoded function
names or call shapes.

For a repeatable execution-only comparison (including RunFunc entry overhead),
run:

```sh
GOWORK=off go test -run '^$' -bench BenchmarkTypedArgumentsUglyFib -benchmem -benchtime=300ms -count=3 .
```

Correctness checks cover multiple arguments and results, pointer aliases,
array copies, pointer Phi, unary operations, mutual recursion, arena growth,
concurrent runner use, nil dereferences, array bounds, and unsupported-callee
fallback. Run the focused tests and race checks from the repository root:

```sh
GOWORK=off go test -run 'TestTypedArguments|TestArgumentsSelection' -count=1 . ./typedrunner
GOWORK=off go test -race -run 'TestTypedArguments|TestArgumentsSelection' -count=1 . ./typedrunner
```

The recursive files come from Go's `test/abi` directory and retain their copyright
notices. See [Go's BSD license](https://go.dev/LICENSE).

## Current execution benchmarks

Remeasured on the script/documentation revision `846b215e` (runner implementation
unchanged from `59f19132`), Apple M5, Go 1.27.1, GOMAXPROCS=1. Values are medians
of three 300 ms benchmark runs. These measurements exclude source loading;
loop and uglyfib benchmarks include `RunFunc` entry overhead. Both modes use
the current branch: boxed disables the runner for the target function (all
functions for uglyfib), so this table is not a comparison against main.
Loop inputs are 1,000 except nested (50); uglyfib uses 20.

| Benchmark | Boxed | Typed | Boxed / typed | Boxed / typed allocs/op |
| --- | ---: | ---: | ---: | ---: |
| sum loop | 33.511 us | 8.303 us | 4.04x | 1,725 / 4 |
| swap loop | 27.816 us | 8.504 us | 3.27x | 750 / 5 |
| fibloop | 36.865 us | 8.306 us | 4.44x | 1,738 / 7 |
| branch loop | 54.643 us | 12.699 us | 4.30x | 1,718 / 4 |
| nested loop | 41.731 us | 13.852 us | 3.01x | 1,153 / 3 |
| uglyfib(20) | 8.009 ms | 1.750 ms | 4.58x | 196,642 / 6 |

Uglyfib allocates 2,557,888 B/op boxed and 84,824 B/op typed. Every iteration
checks `Fib(20)=6765`. The scalar fibish(20) benchmark, compiled from SSA and
called directly without `RunFunc`, measured 187.455 us/op with zero allocations.

| Direct RunArgs benchmark | Time | B/op | allocs/op |
| --- | ---: | ---: | ---: |
| Pointer leaf | 39.09 ns | 0 | 0 |
| Pointer and mixed results | 177.30 ns | 56 | 2 |
| Pointer and depth-100 recursion | 39.345 us | 428,672 | 4 |

The direct argument benchmarks exclude interpreter entry and loading. Scalar
dispatch reads immutable instructions through pointers. Argument runners reserve
one initial frame for leaf functions and eight for functions containing calls;
this avoids initializing unused frames in shallow calls while retaining
recursion capacity. The mixed result tuple is boxed at the outer return.
Each invocation owns its arena, preserving concurrent runner safety.

Run these benchmarks from the repository root:

```sh
GOWORK=off GOMAXPROCS=1 go test -run '^$' -bench 'BenchmarkTypedRunnerFibish/n20|BenchmarkTypedRunnerLoops|BenchmarkTypedArguments' -benchmem -benchtime=300ms -count=3 .
```

Timing varies with machine load, compiler and interpreter versions. The earlier
scalar-only, bool-change, and pre-review tables have been replaced by the
current implementation measurements above.
