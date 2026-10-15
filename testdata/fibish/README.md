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

Scalar functions use the original scalar instruction path. Fusion is limited
to the argument/memory runner.

## Comparison with main

Measured 2026-10-15 on Apple M5, macOS arm64, Go 1.27.1, `GOWORK=off`,
`GOMAXPROCS=1`, input 40.
Upstream main was verified as `5c99c936`; the final PR implementation is
`f3b1d418`. Both CLIs were built with the same toolchain and execute the same
checked-in fixture sources from this PR.

Each entry is the median of three serial alternating main/PR process runs,
including process startup and interpreter loading, without a separate warmup.
The order is main/PR in rounds 1 and 3, PR/main in round 2. All 36 outputs were
checked against the documented expected values. No other benchmark or build
ran concurrently.

| Example | main `5c99c936` | PR `f3b1d418` | main / PR |
| --- | ---: | ---: | ---: |
| `fibish` | 3.901 s | 2.242 s | 1.74x |
| `fibish_closure` | 6.367 s | 2.241 s | 2.84x |
| `fibish_loop` | 52.882 ms | 52.657 ms | 1.00x |
| `fibish_float32` | 8.868 s | 1.892 s | 4.69x |
| `fibish_float64` | 9.429 s | 2.380 s | 3.96x |
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
| `fibish_closure` | 6.374504, 6.349313, 6.367240 | 2.159013, 2.240693, 2.265239 |
| `fibish_loop` | 0.054167, 0.052882, 0.052831 | 0.053662, 0.052202, 0.052657 |
| `fibish_float32` | 8.867665, 8.843980, 8.870807 | 1.883144, 1.900261, 1.891633 |
| `fibish_float64` | 9.428808, 9.422065, 9.441496 | 2.296484, 2.379577, 2.416122 |
| `uglyfib` | 114.467070, 114.270771, 115.469531 | 19.363909, 19.415228, 19.505525 |

</details>

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
