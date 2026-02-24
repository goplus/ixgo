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

package repl

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"

	"github.com/goplus/ixgo"
	"github.com/goplus/ixgo/cmd/internal/base"
	_ "github.com/goplus/ixgo/pkg"
	"github.com/goplus/ixgo/repl"
	"github.com/peterh/liner"
)

// Cmd - ixgo test
var Cmd = &base.Command{
	UsageLine: "ixgo repl",
	Short:     "ixgo repl mode",
}

var (
	flag          = &Cmd.Flag
	flagGoPlus    bool
	flagGoOnly    bool
	flagDumpInstr bool
	flagTrace     bool
)

func init() {
	Cmd.Run = runCmd
	flag.BoolVar(&flagGoPlus, "xgo", true, "support XGo mode")
	flag.BoolVar(&flagGoOnly, "go", false, "use Go mode only")
	flag.BoolVar(&flagDumpInstr, "dump", false, "dump SSA instruction code")
	flag.BoolVar(&flagTrace, "trace", false, "trace interpreter code")
}

// LinerUI implements repl.UI interface.
type LinerUI struct {
	state  *liner.State
	prompt string
}

// SetPrompt is required by repl.UI interface.
func (u *LinerUI) SetPrompt(prompt string) {
	u.prompt = prompt
}

// Printf is required by repl.UI interface.
func (u *LinerUI) Printf(format string, a ...interface{}) {
	fmt.Printf(format, a...)
}

var (
	welcomeGo string = fmt.Sprintf("iXGo v1.0.0 (build %v %v/%v)", runtime.Version(), runtime.GOOS, runtime.GOARCH)
)

var helpGo string = `Use ?expr to dump expr information
Use ?pkg.symbol to dump pkg symbol information
Use help() to show help information
Use exit() or Ctrl-D to exit`

var helpGop string = `Use ?expr to dump expr information
Use ?pkg.symbol to dump pkg symbol information
Use help to show help information
Use exit or Ctrl-D to exit`

var (
	gopVersion    string
	supportGoplus bool
)

func runCmd(cmd *base.Command, args []string) {
	err := flag.Parse(args)
	if err != nil {
		os.Exit(2)
	}
	if flagGoOnly {
		flagGoPlus = false
	}
	var help string
	var welcome string
	if supportGoplus && flagGoPlus {
		welcome = welcomeGo + " (XGo version " + gopVersion + ")"
		help = helpGop
	} else {
		welcome = welcomeGo
		help = helpGo
	}
	fmt.Println(welcome)
	fmt.Println(help)

	state := liner.NewLiner()
	defer state.Close()

	// state.SetCtrlCAborts(true)
	state.SetMultiLineMode(true)
	state.SetCompleter(func(line string) []string {
		if strings.TrimSpace(line) == "" {
			return []string{line + "    "}
		}
		return nil
	})
	ui := &LinerUI{state: state}
	var mode ixgo.Mode
	if flagDumpInstr {
		mode |= ixgo.EnableDumpInstr
	}
	if flagTrace {
		mode |= ixgo.EnableTracing
	}
	var r *repl.REPL
	ixgo.RegisterCustomBuiltin("exit", func() {
		r.Interp().Exit(0)
	})
	ixgo.RegisterCustomBuiltin("help", func() {
		fmt.Println(help)
	})
	r = repl.NewREPL(mode)
	r.SetUI(ui)
	if supportGoplus && flagGoPlus {
		r.SetFileName("main.xgo")
	}
	for {
		line, err := ui.state.Prompt(ui.prompt)
		if err != nil {
			if err == liner.ErrPromptAborted || err == io.EOF {
				fmt.Printf("exit\n")
				break
			}
			fmt.Printf("Problem reading line: %v\n", err)
			continue
		}
		if strings.TrimSpace(line) == "?" {
			fmt.Println(help)
			continue
		}
		if line != "" {
			state.AppendHistory(line)
		}
		err = r.Run(line)
		switch e := err.(type) {
		case nil:
			//
		case ixgo.ExitError:
			fmt.Printf("exit %v\n", int(e))
			return
		default:
			fmt.Printf("error: %v\n", err)
		}
	}
}
