/*
 Copyright 2021 The GoPlus Authors (goplus.org)

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

// Package build implements the “ixgo version” command.
package version

import (
	"fmt"
	"runtime"

	"github.com/goplus/ixgo/cmd/internal/base"
)

// -----------------------------------------------------------------------------

// Cmd - ixgo build
var Cmd = &base.Command{
	UsageLine: "ixgo version",
	Short:     "print version",
}

var (
	flag = &Cmd.Flag
)

func init() {
	Cmd.Run = versionCmd
}

func versionCmd(cmd *base.Command, args []string) {
	var ext string
	if base.LLGo {
		ext = "llgo "
	}
	fmt.Printf("ixgo build %v%v %v/%v\n", ext, runtime.Version(), runtime.GOOS, runtime.GOARCH)
}
