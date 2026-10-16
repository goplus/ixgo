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
eligible integer recursive examples now memoize repeated states; floating-point
recursive examples retain their original recursive execution. The difference
also reflects the algorithms. Compare the same file between ixgo versions to measure
an interpreter optimization. The library requires a blank import of
`github.com/goplus/ixgo/typedrunner` to enable the optional typed runner; the CLI
already imports it.

## Typed runner implementation

The CLI imports the optional `typedrunner` plugin by default. Applications can
enable it with a blank import; `MakeTypedRunnerOf[T]` selects the scalar or
argument/memory implementation internally. Integer, float32, and float64 code
uses typed numeric registers with separate bool registers. Unsupported SSA or
callees fall back to the ordinary interpreter.

The scalar runner avoids boxed recursive frames. The argument/memory runner
supports closed direct call graphs, mutual recursion, scalar pointers, local
arrays, array copies, loads/stores, multiple arguments, and numeric/bool results.
Frame/slot handles preserve aliases when the frame arena grows; execution state
is private to each invocation. Each supported graph uses one numeric type plus
bool. Mixed numeric types and conversions, pointer results, external calls,
defer, and allocations inside loops are unsupported. Each frame is limited to
64 registers and 64 local scalar slots; recursion grows the arena as needed.

The argument/memory runner groups typed call argument moves by scalar, pointer, and bool kind, and enables dense numeric constant batching,
comparison/branch fusion, and eligible load/add/store fusion. These passes
process the complete call graph before publication and preserve branch targets,
intermediate registers, float behavior, and pointer aliases. They match generic
instruction patterns without hardcoded function names or Fibonacci call shapes.
The default factory prefers the scalar runner where supported.

Recursive integer scalar functions memoize results per invocation, using a
64-entry inline table and an overflow map. Floating-point recursion retains the
result-register fast path to preserve signed-zero and rounding behavior.
Non-recursive scalar functions keep the original scalar instruction path.

Eligible integer accumulator call graphs also memoize effect summaries. Their
signature is one scalar input, two read-only scalar pointers, and an accumulator
pointer, with no return values. The key contains the runner identity and three
input values; the cached value is the accumulator delta. SSA checks restrict
writes and accumulator reads, and runtime checks disable memoization when the
accumulator aliases an input. Nil inputs retain normal execution. The cache is
private to each invocation. Functions outside these conditions retain the
existing execution path. Memoization changes the amount of recursive work;
these results measure that algorithmic optimization as well as execution costs.

## Memoization update (2026-10-16)

Apple M5, macOS arm64, Go 1.27.1, `GOWORK=off`, `GOMAXPROCS=1`, input 40.
The following are three serial CLI runs of the memoized implementation, including
startup/loading and excluding compilation. All outputs were checked. The old
main/PR and other-language measurements below remain historical data and were
not rerun in this batch. These tables have different measurement dates and
should not be interpreted as a fresh controlled main/PR comparison.

| Example | Rounds (seconds) | Median |
| --- | --- | ---: |
| `fibish` | 0.902440, 0.051971, 0.053971 | 0.053971 s |
| `fibish_closure` | 0.051959, 0.051978, 0.053500 | 0.051978 s |
| `fibish_loop` | 0.053428, 0.052725, 0.051867 | 0.052725 s |
| `fibish_float32` | 1.758551, 1.793242, 1.763778 | 1.763778 s |
| `fibish_float64` | 1.933997, 2.045712, 2.034520 | 2.034520 s |
| `uglyfib` | 0.058626, 0.055114, 0.053765 | 0.055114 s |

The first `fibish` process had additional startup wall time. Its cause was not
isolated. Loading dominates the memoized integer examples at input 40.
Float examples retain ordinary recursive execution.

Reproduce the accumulator execution benchmark with:

```sh
GOWORK=off go test -run '^$' -bench BenchmarkTypedAccumulatorMemo -benchmem -count=3 .
```

The baseline and memo variants apply the same instruction fusion passes; only
memo enables effect-summary caching. Serial measurements at input 20 were 1.191 ms/op versus 10.142 us/op
(medians of three), with 57,344 B/op (1 allocation) versus 10,312 B/op
(9 allocations), approximately 117x faster. The scalar fibish(20) benchmark
was 747 ns/op, 200 B/op, 5 allocations. Re-run for machine-specific data.

Validation for this update passed `GOWORK=off go test -count=1 ./...` and
`GOWORK=off go test -race -count=1 . ./typedrunner`. New tests cover memo-table
overflow, floating-point fallback, all four accumulator entry points, integer
overflow, nil pointers, alias fallback, and rejected unsafe side effects.

## Historical comparison with main (before memoization)

Measured 2026-10-15 on Apple M5, macOS arm64, Go 1.27.1, `GOWORK=off`,
`GOMAXPROCS=1`, input 40.
Upstream main was verified as `5c99c936`; the final PR implementation is
`e9ee8a93`. Both CLIs were built with the same toolchain and execute the same
checked-in fixture sources from this PR.

Each entry is the median of three serial alternating main/PR process runs,
including process startup and interpreter loading, without a separate warmup.
The order is main/PR in rounds 1 and 3, PR/main in round 2. All 36 outputs were
checked against the documented expected values. No other benchmark or build
ran concurrently.

| Example | main `5c99c936` | PR `e9ee8a93` | main / PR |
| --- | ---: | ---: | ---: |
| `fibish` | 3.901 s | 1.995 s | 1.96x |
| `fibish_closure` | 6.367 s | 1.910 s | 3.33x |
| `fibish_loop` | 52.882 ms | 52.657 ms | 1.00x |
| `fibish_float32` | 8.868 s | 1.717 s | 5.17x |
| `fibish_float64` | 9.429 s | 2.019 s | 4.67x |
| `uglyfib` | 114.467 s | 19.415 s | 5.90x |

The recursive fixtures spend most of their time in execution. The loop at
input 40 performs little work, so loading dominates; its small timing difference
does not establish an execution regression or improvement. Recursive and loop
fixtures use different algorithms, so compare the same file across versions.
Three runs do not establish statistical significance for small differences.

<details><summary>All main/PR CLI rounds (seconds)</summary>

| Example | main rounds (s) | PR rounds (s) |
| --- | --- | --- |
| `fibish` | 4.067613, 3.901037, 3.876147 | 3.319925, 2.197261, 2.242073 |
| `fibish_closure` | 6.374504, 6.349313, 6.367240 | 2.084920, 1.987784, 1.909700 |
| `fibish_loop` | 0.054167, 0.052882, 0.052831 | 0.053662, 0.052202, 0.052657 |
| `fibish_float32` | 8.867665, 8.843980, 8.870807 | 1.875711, 1.729107, 1.716619 |
| `fibish_float64` | 9.428808, 9.422065, 9.441496 | 2.268138, 2.134130, 2.019422 |
| `uglyfib` | 114.467070, 114.270771, 115.469531 | 19.363909, 19.415228, 19.505525 |

</details>

## Lua, LuaJIT, Python and Ruby comparison

Apple M5, macOS arm64, input 40. All entries are medians of three serial
process runs, including startup and loading, without a separate warmup; all
outputs were checked. The Go rows use the latest main/PR measurements above,
with Go 1.27.1 and `GOMAXPROCS=1`. The language rows retain the previously
measured results from separate batches on the same machine with their default
runtime settings; they have not been rerun for this update or interleaved with
the latest Go measurements. The table is a descriptive cross-language
comparison, rather than a controlled test of interpreter implementation alone.

| Implementation | fibish | fibish_closure | uglyfib |
| --- | ---: | ---: | ---: |
| ixgo main `5c99c936` | 3.901 s | 6.367 s | 114.467 s |
| ixgo typed runner `e9ee8a93` | 1.995 s | 1.910 s | 19.415 s |
| Lua 5.5.1 | 2.412 s | 2.542 s | 26.266 s |
| LuaJIT 2.1.1791890110 | 0.197 s | 0.228 s | 8.987 s |
| Python 3.9.6 | 11.804 s | 12.290 s | 31.057 s |
| Ruby 2.6.10p210 | 6.335 s | 6.591 s | 21.004 s |

Standard Lua uses its bytecode interpreter. LuaJIT runs the same `.lua` files
with its default JIT enabled (`jit.status()` returned true), without tuning
flags. Python and Ruby use their default interpreter settings. The results
apply to the listed runtime versions.

The scripts preserve the recursive branches and normal multiple-result
mechanisms. `fibish_closure` forwards the function argument unchanged; it does
not invoke that argument or allocate a capturing closure per call. Uglyfib
preserves the four mutually recursive functions and local array sizes, using
shared mutable list/table/array slots for pointer aliases. Allocation and GC
behavior therefore differ from Go's pointer/array representation.

The recursive return fast path was also checked against uglyfib. Its `n=20` benchmark was 1.325 ms before and 1.337 ms after (about +0.9%, within noise), with identical allocations, so uglyfib receives no additional scalar benefit.

Scripts and repetition helpers are checked in under `testdata/fibish`.
Run `python3 bench_languages.py python ruby lua luajit` and
`python3 bench_uglyfib.py python ruby lua luajit` there to verify small inputs
and repeat input 40. Generated timing JSON files remain outside version control.

## Reproducing the comparison

Build each revision in its own checkout, then use the PR fixture paths for both
binaries. For example, build main in a separate worktree:

```sh
git worktree add --detach /tmp/ixgo-main 5c99c936
(cd /tmp/ixgo-main && GOWORK=off go build -o /tmp/ixgo-main-bin ./cmd/ixgo)
GOWORK=off go build -o /tmp/ixgo-pr-bin ./cmd/ixgo
GOMAXPROCS=1 /usr/bin/time -p /tmp/ixgo-main-bin run testdata/fibish/fibish.go
GOMAXPROCS=1 /usr/bin/time -p /tmp/ixgo-pr-bin run testdata/fibish/fibish.go
```

Repeat serially for all six examples, alternating binary order as described
above. Building is excluded from the timings. The measurements use wall-clock
time for the entire CLI process; they do not isolate runner execution from
loading. Compiler changes, linked code layout, and machine load can affect
results. Generated timing JSON files are not checked in.

The recursive return fast path was also checked against uglyfib. Its `n=20`
benchmark was 1.325 ms before and 1.337 ms after (about +0.9%, within noise),
with identical allocations, so uglyfib receives no additional scalar benefit.

## Validation

The final implementation passed:

```sh
GOWORK=off go test -count=1 ./...
GOWORK=off go test -race -count=1 . ./typedrunner
```

Tests cover actual execution through the default factory/plugin, all uglyfib
entry points, pointer aliases, intermediate values, nil panics, float bit
patterns, mixed bool/scalar constant initialization, recursive arena growth,
pass order/idempotence, and concurrent runner use. All main/PR CLI results were
checked, including native-Go-compatible float32 rounding.

The recursive fixtures come from Go's `test/abi` directory and retain their
copyright notices. See [Go's BSD license](https://go.dev/LICENSE).
