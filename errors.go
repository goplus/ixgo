/*
 * Copyright (c) 2022 The GoPlus Authors (goplus.org). All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ixgo

import (
	"bytes"
	"errors"
	"fmt"
)

var (
	ErrNotFoundMain    = errors.New("not found main package")
	ErrTestFailed      = errors.New("test failed")
	ErrNotFoundPackage = errors.New("not found package")
	ErrGoexitDeadlock  = errors.New("fatal error: no goroutines (main called runtime.Goexit) - deadlock!")
	ErrNoFunction      = errors.New("no function")
	ErrNoTestFiles     = errors.New("[no test files]")
	ErrNoTestProcessor = errors.New("test processor not registered: import _ \"github.com/goplus/ixgo/pkg/testing\" or call ixgo.RegisterTestProcessor")
)

type ExitError int

func (r ExitError) Error() string {
	return fmt.Sprintf("exit %v", int(r))
}

type PlainError string

func (e PlainError) RuntimeError() {}

func (e PlainError) Error() string {
	return string(e)
}

type RuntimeError string

func (e RuntimeError) RuntimeError() {}

func (e RuntimeError) Error() string {
	return "runtime error: " + string(e)
}

// If the target program panics, the interpreter panics with this type.
type PanicError struct {
	stack []byte
	Value value
}

func (p PanicError) Error() string {
	var buf bytes.Buffer
	writeany(&buf, p.Value)
	return buf.String()
}

func (p PanicError) Stack() []byte {
	return p.stack
}

// run func fatal error
type FatalError struct {
	stack []byte
	Value value
}

func (p FatalError) Error() string {
	var buf bytes.Buffer
	writeany(&buf, p.Value)
	return buf.String()
}

func (p FatalError) Stack() []byte {
	return p.stack
}

// If the target program calls exit, the interpreter panics with this type.
type exitPanic int

type goexitPanic int
