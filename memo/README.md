# Optional memoization

`github.com/goplus/ixgo/memo` enables memoization for the ordinary ixgo interpreter.
The package is optional for applications embedding ixgo. The `ixgo` command imports
it explicitly, so CLI runs enable memoization by default.

## Enable memoization

Add a blank import to the program that creates or runs the interpreter:

```go
import _ "github.com/goplus/ixgo/memo"
```

The package registers the memo plugin during `init`. Without this import, ixgo keeps
memoization disabled. The `ixgo` command includes this blank import, so command-line
runs enable memoization automatically. Applications embedding ixgo choose whether to
include the import.

## Supported paths

The plugin analyzes each function with SSA (Static Single Assignment) information
before installing a cache. It supports these paths:

- Pure scalar recursion with integer, `float32`, and `float64` values.
- Integer accumulator recursion with one isolated pointer output.
- Receiver methods with cached receiver fields.
- Integer additive, affine, and supported bitwise receiver transforms.
- Floating-point receiver state caching using exact IEEE-754 bit patterns.
- Mutual recursion when the complete call graph satisfies the safety checks.

The analysis is conservative. Unsupported or observable behavior stays on the ordinary
interpreter path. This includes multiple mutable outputs, unsafe pointer aliasing,
external side effects, hooks, tracing, debug callbacks, custom call interception,
and unsupported instructions.

## Floating-point behavior

Float keys and cached state use exact bit patterns. The implementation preserves:

- `+0` and `-0`;
- NaN payloads;
- positive and negative infinity;
- `float32` precision and step-by-step rounding.

The plugin does not regroup floating-point operations. The `uglyfib` floating-point
files are examples only and are not memoized because preserving their operation order
does not provide a safe reusable state for this optimizer.

## Package layout

The SSA eligibility checks and receiver analysis live in this package. Shared runtime
bit-pattern helpers used by the interpreter live in `github.com/goplus/ixgo/internal/memo`.

## Tests

From the repository root:

```sh
go test ./memo
GOWORK=off go test -race ./memo
```

The integration tests for the ordinary interpreter are in `interp_memo_test.go`. The
SSA eligibility tests are in this package because the analysis is package-private.

## fibish performance comparison

The following measurements compare the ordinary interpreter with memoization enabled
against the `main` baseline. Both sides use the same `testdata/fibish` source,
`GOWORK=off`, `GOMAXPROCS=1`, macOS arm64, and include process startup and source
loading. Times are shown to 0.001 seconds. Each value is the median of five
complete-process runs. The table lists fixtures with a recorded baseline, correct
output, and a useful memoization result.

| Fixture | `main` | Memoized | Speedup |
| --- | ---: | ---: | ---: |
| `fibish.go` | 3.997 s | 0.053 s | 75.4x |
| `fibish_closure.go` | 6.371 s | 0.053 s | 120.2x |
| `fibish_struct.go` | 4.105 s | 0.058 s | 70.8x |
| `fibish_struct2.go` | 14.104 s | 0.055 s | 256.4x |
| `uglyfib.go` | 114.526 s | 0.056 s | 2,045.1x |
| `uglyfib_struct.go` | 124.912 s | 0.056 s | 2,230.6x |
| `fibish_float32.go` | 8.831 s | 0.053 s | 166.6x |
| `fibish_float64.go` | 9.420 s | 0.056 s | 168.2x |
| `fibish_struct_float32.go` | 7.081 s | 0.055 s | 128.7x |
| `fibish_struct_float64.go` | 8.243 s | 0.056 s | 147.2x |
| `fibish_struct2_float32.go` | 18.845 s | 0.053 s | 358.5x |
| `fibish_struct2_float64.go` | 22.719 s | 0.052 s | 433.9x |

These files are examples only. They are not memoized:

- `uglyfib_float32.go`
- `uglyfib_float64.go`
- `uglyfib_struct_float32.go`
- `uglyfib_struct_float64.go`

