# typedrunner

Package `typedrunner` enables the optional scalar typed runner in ixgo.

```go
import _ "github.com/goplus/ixgo/typedrunner"
```

Without this import, ixgo uses its normal boxed interpreter path.

The plugin selects integer, float32, float64, and boolean functions supported
by the core compiler, including recursive functions, loops, and branch merges.
Numeric registers use one scalar type per function and call graph; boolean
registers are separate. Mixed numeric types and numeric conversions fall back
to the normal interpreter.

Floating-point runners support arithmetic (`+`, `-`, `*`, `/`), comparisons,
unary negation, and scalar pointers and local arrays. Integer-only operators
remain available for integer runners. Phi inputs use parallel assignment on
incoming control-flow edges. Unsupported instructions and functions that
exceed the typed register limit fall back to the normal runner.
