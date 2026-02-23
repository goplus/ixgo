package main

import (
	"bytes"
	"fmt"
	"go/build"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	gorootTestSkips = make(map[string]string)
)

func init() {
	if runtime.GOARCH == "386" {
		gorootTestSkips["printbig.go"] = "load failed"
		gorootTestSkips["peano.go"] = "stack overflow"
	}
	gorootTestSkips["closure.go"] = "runtime.ReadMemStats"
	gorootTestSkips["divmod.go"] = "slow, 1m18s"
	gorootTestSkips["copy.go"] = "slow, 13s"
	gorootTestSkips["finprofiled.go"] = "slow, 21s"
	gorootTestSkips["gcgort.go"] = "slow, 2s"
	gorootTestSkips["nilptr.go"] = "skip drawin"
	gorootTestSkips["heapsampling.go"] = "runtime.MemProfileRecord"
	gorootTestSkips["makeslice.go"] = "TODO, panic info, allocation size out of range"
	// gorootTestSkips["stackobj.go"] = "skip gc"
	// gorootTestSkips["stackobj3.go"] = "skip gc"
	gorootTestSkips["nilptr_aix.go"] = "skip"
	// gorootTestSkips["init1.go"] = "skip gc"
	gorootTestSkips["ken/divconst.go"] = "slow, 3.5s"
	gorootTestSkips["ken/modconst.go"] = "slow, 3.3s"
	gorootTestSkips["fixedbugs/issue24491b.go"] = "timeout"
	gorootTestSkips["fixedbugs/issue16249.go"] = "slow, 4.5s"
	gorootTestSkips["fixedbugs/issue13169.go"] = "slow, 5.9s"
	gorootTestSkips["fixedbugs/issue11656.go"] = "ignore"
	// gorootTestSkips["fixedbugs/issue15281.go"] = "runtime.ReadMemStats"
	gorootTestSkips["fixedbugs/issue18149.go"] = "runtime.Caller macos //line not support c:/foo/bar.go:987"
	gorootTestSkips["fixedbugs/issue22662.go"] = "runtime.Caller got $goroot/test/fixedbugs/foo.go:1; want foo.go:1"
	// gorootTestSkips["fixedbugs/issue27518b.go"] = "BUG, runtime.SetFinalizer"
	// gorootTestSkips["fixedbugs/issue32477.go"] = "BUG, runtime.SetFinalizer"
	gorootTestSkips["fixedbugs/issue41239.go"] = "BUG, reflect.Append: different capacity on append"
	// gorootTestSkips["fixedbugs/issue32477.go"] = "BUG, runtime.SetFinalizer"
	gorootTestSkips["fixedbugs/issue45175.go"] = "BUG, ssa.Phi call order"
	gorootTestSkips["fixedbugs/issue4618.go"] = "testing.AllocsPerRun"
	gorootTestSkips["fixedbugs/issue4667.go"] = "testing.AllocsPerRun"
	gorootTestSkips["fixedbugs/issue8606b.go"] = "BUG, optimization check"
	gorootTestSkips["fixedbugs/issue30116u.go"] = "BUG, slice bound check"
	gorootTestSkips["chan/select5.go"] = "bug, select case expr call order"

	// fixedbugs/issue7740.go
	// const ulp = (1.0 + (2.0 / 3.0)) - (5.0 / 3.0)
	// Go 1.14 1.15 1.16 ulp = 1.4916681462400413e-154
	// Go 1.17 1.18 ulp = 0

	// go1.24.3 => 24
	ver, err := strconv.Atoi(runtime.Version()[4:6])
	if err != nil {
		panic("version error")
	}
	switch {
	case ver >= 17:
		// gorootTestSkips["fixedbugs/issue45045.go"] = "runtime.SetFinalizer"
		// gorootTestSkips["fixedbugs/issue46725.go"] = "runtime.SetFinalizer"
		gorootTestSkips["abi/fibish.go"] = "slow, 34s"
		gorootTestSkips["abi/fibish_closure.go"] = "slow, 35s"
		gorootTestSkips["abi/uglyfib.go"] = "5m48s"
		// gorootTestSkips["fixedbugs/issue23017.go"] = "BUG" //fixed https://github.com/golang/go/issues/55086

		gorootTestSkips["typeparam/chans.go"] = "runtime.SetFinalizer, maybe broken for go1.18 on linux workflows"
		// gorootTestSkips["typeparam/issue376214.go"] = "build SSA package error: variadic parameter must be of unnamed slice type"
		if ver != 20 {
			gorootTestSkips["typeparam/nested.go"] = "skip, run pass but output same as go1.20"
		}
		//go1.20
		gorootTestSkips["fixedbugs/bug514.go"] = "skip cgo"
		gorootTestSkips["fixedbugs/issue40954.go"] = "skip cgo"
		gorootTestSkips["fixedbugs/issue42032.go"] = "skip cgo"
		gorootTestSkips["fixedbugs/issue42076.go"] = "skip cgo"
		gorootTestSkips["fixedbugs/issue46903.go"] = "skip cgo"
		gorootTestSkips["fixedbugs/issue51733.go"] = "skip cgo"
		//go1.21
		gorootTestSkips["fixedbugs/issue19658.go"] = "skip command"
		// gorootTestSkips["fixedbugs/issue57823.go"] = "GC"
		if ver == 18 {
			gorootTestSkips["typeparam/cons.go"] = "skip golang.org/x/tools v0.7.0 on go1.18"
			gorootTestSkips["typeparam/list2.go"] = "skip golang.org/x/tools v0.7.0 on go1.18"
		}
		if ver >= 22 {
			gorootTestSkips["fixedbugs/bug369.go"] = "skip command"
			gorootTestSkips["fixedbugs/issue10607.go"] = "skip command"
			gorootTestSkips["fixedbugs/issue21317.go"] = "skip command"
			gorootTestSkips["fixedbugs/issue38093.go"] = "skip js"
			gorootTestSkips["fixedbugs/issue64565.go"] = "skip command"
			gorootTestSkips["fixedbugs/issue9355.go"] = "skip command"
			gorootTestSkips["fixedbugs/issue69110.go"] = "skip runtime link"
			gorootTestSkips["linkmain_run.go"] = "skip link"
			gorootTestSkips["linkobj.go"] = "skip link"
			gorootTestSkips["linkx_run.go"] = "skip link"
			gorootTestSkips["chanlinear.go"] = "skip -gc-exp"
		}
		if ver >= 25 {
			gorootTestSkips["fixedbugs/issue72844.go"] = "BUG, range for nil *op[N]"
			gorootTestSkips["fixedbugs/issue73476.go"] = "BUG, range for nil *op[N]"
		}
		if ver >= 26 {
			gorootTestSkips["range4.go"] = "BUG, range"
			gorootTestSkips["rangegen.go"] = "BUG, range"
		}
	case ver == 16:
		gorootTestSkips["fixedbugs/issue7740.go"] = "BUG, const float"
	case ver == 15:
		gorootTestSkips["fixedbugs/issue15039.go"] = "BUG, uint64 -> string"
		gorootTestSkips["fixedbugs/issue9355.go"] = "TODO, chdir"
		gorootTestSkips["fixedbugs/issue7740.go"] = "BUG, const float"
	case ver == 14:
		gorootTestSkips["fixedbugs/issue9355.go"] = "TODO, chdir"
		gorootTestSkips["fixedbugs/issue7740.go"] = "BUG, const float"
	}

	if runtime.GOOS == "windows" {
		gorootTestSkips["env.go"] = "skip GOARCH"
		gorootTestSkips["fixedbugs/issue15002.go"] = "skip windows"
		gorootTestSkips["fixedbugs/issue5493.go"] = "skip windows"
		gorootTestSkips["fixedbugs/issue5963.go"] = "skip windows"
		if ver >= 22 {
			gorootTestSkips["recover4.go"] = "skip windows"
			gorootTestSkips["sigchld.go"] = "skip windows"
		}

		skips := make(map[string]string)
		for k, v := range gorootTestSkips {
			skips[filepath.FromSlash(k)] = v
		}
		gorootTestSkips = skips
	} else if runtime.GOOS == "darwin" {
		gorootTestSkips["locklinear.go"] = "skip github"
		gorootTestSkips["env.go"] = "skip github"
	}
}

var (
	cmdGo   string
	cmdIxgo string
)

func init() {
	var err error
	cmdIxgo, err = exec.LookPath("ixgo")
	if err != nil {
		panic(fmt.Sprintf("not found ixgo: %v", err))
	}
	cmdGo, err = exec.LookPath("go")
	if err != nil {
		panic(fmt.Sprintf("not found go: %v", err))
	}
}

func runCommand(input string, chkout bool) bool {
	fmt.Println("Input:", input)
	start := time.Now()
	cmd := exec.Command(cmdIxgo, "run", "-exp-gc", input)
	data, err := cmd.CombinedOutput()
	if len(data) > 0 {
		fmt.Println(string(data))
	}
	if chkout {
		if chk, e := ioutil.ReadFile(input[:len(input)-3] + ".out"); e == nil {
			if bytes.Compare(data, chk) != 0 {
				err = fmt.Errorf("-- output check error --\n%v", string(chk))
			}
		}
	}

	sec := time.Since(start).Seconds()
	if err != nil || bytes.Contains(data, []byte("BUG")) {
		fmt.Println(err)
		fmt.Printf("FAIL %0.3fs\n", sec)
		return false
	}
	fmt.Printf("PASS %0.3fs\n", sec)
	return true
}

func printFailures(failures []string) {
	if failures != nil {
		fmt.Println("The following tests failed:")
		for _, f := range failures {
			fmt.Printf("\t%s\n", f)
		}
	}
}

type runfile struct {
	filePath string // go source file path
	checkOut bool   // check output file.out
}

// $GOROOT/test/*.go runs
func getGorootTestRuns() (dir string, run []runfile, runoutput []string) {
	dir = filepath.Join(build.Default.GOROOT, "test")
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			if path == dir {
				return nil
			}
			_, n := filepath.Split(path)
			switch n {
			case "bench", "dwarf", "codegen":
				return filepath.SkipDir
			case "typeparam":
				ver := runtime.Version()[:6]
				switch ver {
				case "go1.18", "go1.19", "go1.20":
					return nil
				default:
					return filepath.SkipDir
				}
			default:
				if strings.Contains(n, ".dir") {
					return filepath.SkipDir
				}
			}
			return nil
		}
		if filepath.Ext(path) != ".go" {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Printf("read %v error: %v\n", path, err)
			return nil
		}
		lines := strings.Split(string(data), "\n")
		if len(lines) > 0 {
			line := strings.TrimSpace(lines[0])
			if line == "// run" || line == "// run -gcflags=-G=3" {
				rf := runfile{filePath: path}
				if s, err := os.Stat(path[:len(path)-3] + ".out"); err == nil && !s.IsDir() {
					rf.checkOut = true
				}
				run = append(run, rf)
			} else if line == "// runoutput" {
				runoutput = append(runoutput, path)
			}
		}
		return nil
	})
	return
}

func execRunoutput(input string) (string, error) {
	cmd := exec.Command(cmdGo, "run", input)
	data, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	fileName := filepath.Join(os.TempDir(), "tmp.go")
	err = ioutil.WriteFile(fileName, data, 0666)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func main() {
	dir, runs, runoutput := getGorootTestRuns()
	var failures []string
	start := time.Now()
	for _, input := range runs {
		f := input.filePath[len(dir)+1:]
		if info, ok := gorootTestSkips[f]; ok {
			fmt.Println("Skip:", input.filePath, info)
			continue
		}
		if !runCommand(input.filePath, input.checkOut) {
			failures = append(failures, input.filePath)
		}
	}
	for _, input := range runoutput {
		f := input[len(dir)+1:]
		if info, ok := gorootTestSkips[f]; ok {
			fmt.Println("Skip:", input, info)
			continue
		}
		fmt.Println("runoutput:", input)
		file, err := execRunoutput(input)
		if err != nil {
			fmt.Printf("go run %v error, %v\n", input, err)
			continue
		}
		if !runCommand(file, false) {
			failures = append(failures, input)
		}
	}
	printFailures(failures)
	fmt.Printf("time: %.3fs\n", time.Since(start).Seconds())
	if failures != nil {
		os.Exit(2)
	}
}
