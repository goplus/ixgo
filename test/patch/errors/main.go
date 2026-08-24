//go:build go1.26

package main

import (
	"errors"
	"fmt"
)

func main() {
	e := fmt.Errorf("wrapped: %w", errors.New("x"))
	if _, ok := errors.AsType[error](e); !ok {
		panic("AsType")
	}
}
