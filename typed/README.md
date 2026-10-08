# Typed callbacks

Register a signature by embedding `FuncVal` and binding a method. This example
makes interpreted `func(int) int` values ordinary Go functions:

```go
package myint

import "github.com/goplus/ixgo"

type callback struct{ ixgo.FuncVal }

func (c callback) Int(n int) int {
	return c.Call(n).(int)
}

func init() {
	ixgo.RegisterTypedCallbackFunc(func(c ixgo.FuncVal) func(int) int {
		return callback{c}.Int
	})
}
```

```go
src := `package main
func main() {
	add := func(x int) int { return x + 1 }
	println(add(41))
}
`
ixgo.RunFile("main.go", src, nil, 0)
```

`func()` is built into ixgo. Import this package for common `func() T` results.

```go
import _ "github.com/goplus/ixgo/typed"
```

Registered signatures:

- `func() bool`
- `func() int` / `int8` / `int16` / `int32` / `int64`
- `func() uint` / `uint8` / `uint16` / `uint32` / `uint64` / `uintptr`
- `func() float32` / `float64`
- `func() string`
- `func() error`

Named types (for example `type Action func()`) stay on `reflect.MakeFunc`
unless the host registers them.

A later registration for the same signature replaces an earlier one. Passing a
nil maker to `RegisterTypedCallback` or `RegisterTypedCallbackFunc` removes the
registration.
