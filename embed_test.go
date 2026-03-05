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

package ixgo_test

import (
	"testing"

	"github.com/goplus/ixgo"
	_ "github.com/goplus/ixgo/pkg/embed"
)

func TestEmbed(t *testing.T) {
	_, err := ixgo.Run("./testdata/embed", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
}

func TestEmbedErrorNoMatching(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata/embed/testdata/data.txt
var data string

func main() {
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}

func TestEmbedErrorMultipleVars(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata/embed/testdata/data1.txt
var data1, data2 string

func main() {
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}

func TestEmbedErrorMisplaced(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata/embed/testdata/data1
//var data1 string

func main() {
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}

func TestEmbedErrorCannotApply(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata/embed/testdata/data1.txt
var data1 [][]byte

func main() {
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}

func TestEmbedErrorVarWithInitializer(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata/embed/testdata/data1.txt
var data1 = "hello"

func main() {
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}

func TestEmbedErrorMultipleFiles(t *testing.T) {
	src := `package main

import (
	_ "embed"
)

//go:embed testdata/embed/testdata
var data1 []byte

func main() {
}
`
	_, err := ixgo.RunFile("main.go", src, nil, 0)
	if err == nil {
		t.Fatal("must panic")
	}
	t.Log(err)
}
