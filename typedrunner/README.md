# typedrunner

Package `typedrunner` enables the optional scalar typed runner in ixgo.

```go
import _ "github.com/goplus/ixgo/typedrunner"
```

Without this import, ixgo uses its normal boxed interpreter path.

The plugin selects scalar functions containing direct recursion or SSA Phi
instructions, including ordinary integer loops and branch merges. Phi inputs
are copied on incoming control-flow edges with parallel assignment semantics.
Integer and boolean Phi values are supported. Unsupported instructions and
functions that exceed the typed register limit fall back to the normal runner.
