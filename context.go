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
	"context"
	"flag"
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"go/types"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/goplus/ixgo/load"
	"golang.org/x/tools/go/ssa"
)

// Mode is a bitmask of options affecting the interpreter.
type Mode uint

const (
	DisableRecover                 Mode                   = 1 << iota // Disable recover() in target programs; show interpreter crash instead.
	DisableCustomBuiltin                                              // Disable load custom builtin func
	EnableDumpImports                                                 // print import packages
	EnableDumpInstr                                                   // Print packages & SSA instruction code
	EnableTracing                                                     // Print a trace of all instructions as they are interpreted.
	EnablePrintAny                                                    // Enable builtin print for any type ( struct/array )
	EnableNoStrict                                                    // Enable no strict mode
	EnableLoadAllPackages                                             // Enable load all loaded packages
	ExperimentalSupportGC                                             // experimental support runtime.GC
	SupportMultipleInterp                                             // Support multiple interp, must manual release interp reflectx icall.
	CheckGopOverloadFunc                                              // Check and skip gop overload func
	DisableAutoLoadPatchs                                             // Disable automatic loading of package patches.
	DisableDynamicFuncCallAnalysis                                    // Disable dynamic func call analysis and optimization.
	OptionLoadRutimeImethod                                           // Option load runtime imethod for less imethod and memory space.
	OptionLoadAllImethod                                              // Option load all imethod.
	OptionLoadDefaultImethod       = 0                                // Option load default imethod.
	LastMode                       = OptionLoadAllImethod             // Last Mode
)

// Loader types loader interface
type Loader interface {
	Import(path string) (*types.Package, error)
	Installed(path string) (*Package, bool)
	Packages() []*types.Package
	LookupReflect(typ types.Type) (reflect.Type, bool)
	LookupTypes(typ reflect.Type) (types.Type, bool)
	SetImport(path string, pkg *types.Package, load func() error) error
}

// Context ssa context
type Context struct {
	Loader       Loader                                                   // types loader
	BuildContext build.Context                                            // build context, default build.Default
	RunContext   context.Context                                          // run context, default unset
	Importer     types.Importer                                           // importer
	output       io.Writer                                                // capture print/println output
	FileSet      *token.FileSet                                           // file set
	sizes        types.Sizes                                              // types unsafe sizes
	Lookup       func(root, path string) (dir string, found bool)         // lookup external import
	evalCallFn   func(interp *Interp, call *ssa.Call, res ...interface{}) // internal eval func for repl
	debugFunc    func(*DebugInfo)                                         // debug func
	panicFunc    func(*PanicInfo)                                         // panic func
	pkgs         map[string]*SourcePackage                                // imports
	override     map[string]reflect.Value                                 // override function
	evalInit     map[string]bool                                          // eval init check
	nestedMap    map[*types.Named]int                                     // nested named index
	mfnmap       map[mfnKey]mfnValue                                      // mfn map for global context
	mfnid        int                                                      // mfn id for global context
	root         string                                                   // project root
	callForPool  int                                                      // least call count for enable function pool
	Mode         Mode                                                     // mode
	BuilderMode  ssa.BuilderMode                                          // ssa builder mode
	Builder      *SSABuilder                                              // ssa builder
	evalMode     bool                                                     // eval mode
}

func (ctx *Context) setRoot(root string) {
	ctx.root = root
}

func (ctx *Context) lookupPath(path string) (dir string, found bool) {
	if ctx.Lookup != nil {
		dir, found = ctx.Lookup(ctx.root, path)
	}
	if !found {
		bp, err := ctx.BuildContext.Import(path, ctx.root, build.FindOnly)
		if err == nil && bp.ImportPath == path {
			return bp.Dir, true
		}
	}
	return
}

type SourcePackage struct {
	Context  *Context
	Package  *types.Package
	Info     *types.Info
	Importer types.Importer
	Files    []*ast.File
	Links    []*load.LinkSym
	Dir      string
	Register bool // register package
}

func (sp *SourcePackage) Loaded() bool {
	return sp.Info != nil
}

func (sp *SourcePackage) Load() (err error) {
	if sp.Info == nil {
		sp.Info = newTypesInfo()
		conf := &types.Config{
			Sizes:    sp.Context.sizes,
			Importer: sp.Importer,
		}
		if conf.Importer == nil {
			conf.Importer = sp.Context.Importer
		}
		if sp.Context.evalMode {
			conf.DisableUnusedImportCheck = true
		}
		if sp.Context.Mode&EnableNoStrict != 0 {
			conf.Error = func(e error) {
				if te, ok := e.(types.Error); ok {
					if hasTypesNotUsedError(te.Msg) {
						println(fmt.Sprintf("ixgo warning: %v", e))
						return
					}
				}
				if err == nil {
					err = e
				}
			}
		} else {
			conf.Error = func(e error) {
				if err == nil {
					err = e
				}
			}
		}
		types.NewChecker(conf, sp.Context.FileSet, sp.Package, sp.Info).Files(sp.Files)
		if err == nil {
			sp.Links, err = load.ParseLinkname(sp.Context.FileSet, sp.Package.Path(), sp.Files)
		}
	}
	return
}

// NewContext create a new Context
func NewContext(mode Mode) *Context {
	ctx := &Context{
		FileSet:      token.NewFileSet(),
		Mode:         mode,
		BuilderMode:  0, //ssa.SanityCheckFunctions,
		BuildContext: build.Default,
		pkgs:         make(map[string]*SourcePackage),
		override:     make(map[string]reflect.Value),
		nestedMap:    make(map[*types.Named]int),
		mfnmap:       make(map[mfnKey]mfnValue),
		callForPool:  64,
	}
	ctx.Loader = NewTypesLoader(ctx, mode)
	if mode&EnableDumpInstr != 0 {
		ctx.BuilderMode |= ssa.PrintFunctions
	}
	if mode&ExperimentalSupportGC != 0 {
		ctx.RegisterExternal("runtime.GC", runtimeGC)
	}
	ctx.sizes = types.SizesFor("gc", runtime.GOARCH)
	ctx.Lookup = new(load.ListDriver).Lookup
	ctx.Importer = NewImporter(ctx)
	ctx.Builder = NewSSABuilder(ctx)

	if mode&DisableAutoLoadPatchs == 0 {
		for path, srcs := range registerPatchs {
			for _, src := range srcs {
				err := ctx.RegisterPatch(path, src)
				if err != nil {
					log.Printf("import patch %v failed: %v\n", path, err)
				}
			}
		}
	}
	return ctx
}

func (ctx *Context) UnsafeRelease() {
	ctx.pkgs = nil
	ctx.Loader = nil
	ctx.override = nil
}

func (ctx *Context) IsEvalMode() bool {
	return ctx.evalMode
}

func (ctx *Context) SetEvalMode(b bool) {
	ctx.evalMode = b
}

// SetUnsafeSizes set the sizing functions for package unsafe.
func (ctx *Context) SetUnsafeSizes(sizes types.Sizes) {
	ctx.sizes = sizes
}

// SetLeastCallForEnablePool set least call count for enable function pool, default 64
func (ctx *Context) SetLeastCallForEnablePool(count int) {
	ctx.callForPool = count
}

func (ctx *Context) SetDebug(fn func(*DebugInfo)) {
	ctx.BuilderMode |= ssa.GlobalDebug
	ctx.debugFunc = fn
}

func (ctx *Context) SetPanic(fn func(*PanicInfo)) {
	ctx.panicFunc = fn
}

type Frame = frame

func (fr *Frame) CallerFrames() (frames []runtime.Frame) {
	rpc := make([]uintptr, 64)
	n := runtimeCallers(fr, 1, rpc)
	fs := runtime.CallersFrames(rpc[:n])
	for {
		f, more := runtimeFramesNext(fr, fs)
		frames = append(frames, f)
		if !more {
			break
		}
	}
	return
}

type PanicInfo struct {
	funcInstr
	*Frame
	fset  *token.FileSet
	Error error // PanicError / FatalError / PlainError / RuntimeError
}

func (i *PanicInfo) Position() token.Position {
	return i.fset.Position(i.Pos())
}

type funcInstr interface {
	Parent() *ssa.Function
	Pos() token.Pos
	String() string
}

func (ctx *Context) handlePanic(fr *frame, fn funcInstr, err error) error {
	info := &PanicInfo{funcInstr: fn, Frame: fr, fset: ctx.FileSet, Error: err}
	ctx.panicFunc(info)
	if info.Error != nil {
		return info.Error
	}
	return err
}

// RegisterExternal register external value must variable address or func.
func (ctx *Context) RegisterExternal(key string, i interface{}) {
	if i == nil {
		delete(ctx.override, key)
		return
	}
	v := reflect.ValueOf(i)
	switch v.Kind() {
	case reflect.Func, reflect.Ptr:
		ctx.override[key] = v
	default:
		log.Printf("register external must variable address or func. not %v\n", v.Kind())
	}
}

// SetPrintOutput is captured builtin print/println output
func (ctx *Context) SetPrintOutput(output *bytes.Buffer) {
	ctx.output = output
}

func (ctx *Context) writeOutput(data []byte) (n int, err error) {
	if ctx.output != nil {
		return ctx.output.Write(data)
	}
	return os.Stdout.Write(data)
}

// FileSystem represents a file system.
type FileSystem interface {
	ReadDir(dirname string) ([]fs.DirEntry, error)
	ReadFile(filename string) ([]byte, error)
	Join(elem ...string) string

	// Base returns the last element of path.
	// Trailing path separators are removed before extracting the last element.
	// If the path is empty, Base returns ".".
	// If the path consists entirely of separators, Base returns a single separator.
	Base(filename string) string

	// Abs returns an absolute representation of path.
	Abs(path string) (string, error)
}

func (ctx *Context) LoadFileSystem(fsys FileSystem, test bool, getImportPath func(pkgName string, dir string) (string, bool)) (pkg *ssa.Package, err error) {
	ctx.BuildContext.JoinPath = fsys.Join
	ctx.BuildContext.IsDir = func(path string) bool {
		return path == "."
	}
	ctx.BuildContext.ReadDir = func(dir string) (infos []fs.FileInfo, err error) {
		dirs, err := fsys.ReadDir(dir)
		if err != nil {
			return nil, err
		}
		for _, dir := range dirs {
			if dir.IsDir() {
				continue
			}
			info, err := dir.Info()
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}
		return
	}
	ctx.BuildContext.OpenFile = func(name string) (io.ReadCloser, error) {
		data, err := fsys.ReadFile(name)
		if err != nil {
			return nil, err
		}
		return io.NopCloser(bytes.NewReader(data)), nil
	}
	return ctx.LoadDirEx(".", test, getImportPath)
}

func (ctx *Context) LoadFS(fsys fs.FS, test bool, getImportPath func(pkgName string, dir string) (string, bool)) (pkg *ssa.Package, err error) {
	ctx.BuildContext.JoinPath = func(elem ...string) string {
		return filepath.Join(elem...)
	}
	ctx.BuildContext.IsDir = func(path string) bool {
		return path == "."
	}
	ctx.BuildContext.ReadDir = func(dir string) (infos []fs.FileInfo, err error) {
		dirs, err := fs.ReadDir(fsys, dir)
		if err != nil {
			return nil, err
		}
		for _, dir := range dirs {
			if dir.IsDir() {
				continue
			}
			info, err := dir.Info()
			if err != nil {
				continue
			}
			infos = append(infos, info)
		}
		return
	}
	ctx.BuildContext.OpenFile = func(name string) (io.ReadCloser, error) {
		return fsys.Open(name)
	}
	return ctx.LoadDirEx(".", test, getImportPath)
}

func (ctx *Context) LoadDir(dir string, test bool) (pkg *ssa.Package, err error) {
	return ctx.LoadDirEx(dir, test, func(pkgName string, dir string) (string, bool) {
		path, err := load.GetImportPath(pkgName, dir)
		if err != nil {
			log.Println("GetImportPath error:", err)
			return "", false
		}
		return path, true
	})
}

func (ctx *Context) LoadDirEx(dir string, test bool, getImportPath func(pkgName string, dir string) (string, bool)) (pkg *ssa.Package, err error) {
	bp, err := ctx.BuildContext.ImportDir(dir, 0)
	if err != nil {
		return nil, err
	}
	importPath := bp.Name
	if getImportPath != nil {
		if path, ok := getImportPath(bp.Name, dir); ok {
			importPath = path
		}
	}
	bp.ImportPath = importPath
	var sp *SourcePackage
	if test {
		sp, err = ctx.loadTestPackage(bp, importPath, dir)
	} else {
		sp, err = ctx.loadPackage(bp, importPath, dir)
	}
	if err != nil {
		return nil, err
	}
	if ctx.Mode&DisableCustomBuiltin == 0 {
		if f, err := ParseBuiltin(ctx.FileSet, sp.Package.Name()); err == nil {
			sp.Files = append([]*ast.File{f}, sp.Files...)
		}
	}
	ctx.setRoot(dir)
	if dir != "." {
		if wd, err := os.Getwd(); err == nil {
			os.Chdir(dir)
			defer os.Chdir(wd)
		}
	}
	err = sp.Load()
	if err != nil {
		return nil, err
	}
	return ctx.buildPackage(sp)
}

func RegisterFileProcess(ext string, fn SourceProcessFunc) {
	sourceProcessor[ext] = fn
}

type SourceProcessFunc func(ctx *Context, filename string, src interface{}) ([]byte, error)

var (
	sourceProcessor = make(map[string]SourceProcessFunc)
)

func (ctx *Context) AddImportFile(path string, filename string, src interface{}) (err error) {
	_, err = ctx.addImportFile(path, filename, src)
	return
}

func (ctx *Context) AddImport(path string, dir string) (err error) {
	_, err = ctx.addImport(path, dir)
	return
}

func (ctx *Context) SourcePackage(path string) *SourcePackage {
	return ctx.pkgs[path]
}

func (ctx *Context) RegisterPatch(pkg string, src interface{}) error {
	if sp, ok := ctx.pkgs[pkg+"@patch"]; ok {
		file, err := ctx.ParseFile(pkg+"_patch.go", src)
		if err != nil {
			return err
		}
		sp.Files = append(sp.Files, file)
		return nil
	}
	_, err := ctx.addImportFile(pkg+"@patch", pkg+"_patch.go", src)
	return err
}

func (ctx *Context) addImportFile(path string, filename string, src interface{}) (*SourcePackage, error) {
	tp, err := ctx.loadPackageFile(path, filename, src)
	if err != nil {
		return nil, err
	}
	ctx.Loader.SetImport(path, tp.Package, tp.Load)
	return tp, nil
}

func (ctx *Context) addImport(path string, dir string) (*SourcePackage, error) {
	bp, err := ctx.BuildContext.ImportDir(dir, 0)
	if err != nil {
		return nil, err
	}
	bp.ImportPath = path
	tp, err := ctx.loadPackage(bp, path, dir)
	if err != nil {
		return nil, err
	}
	ctx.Loader.SetImport(path, tp.Package, tp.Load)
	return tp, nil
}

func (ctx *Context) loadPackageFile(path string, filename string, src interface{}) (*SourcePackage, error) {
	file, err := ctx.ParseFile(filename, src)
	if err != nil {
		return nil, err
	}
	pkg := types.NewPackage(path, file.Name.Name)
	tp := &SourcePackage{
		Context: ctx,
		Package: pkg,
		Files:   []*ast.File{file},
	}
	ctx.pkgs[path] = tp
	return tp, nil
}

func (ctx *Context) loadPackage(bp *build.Package, path string, dir string) (*SourcePackage, error) {
	files, err := ctx.parseGoFiles(dir, append(bp.GoFiles, bp.CgoFiles...))
	if err != nil {
		return nil, err
	}
	embed, err := load.Embed(bp, ctx.FileSet, files, false, false)
	if err != nil {
		return nil, err
	}
	if embed != nil {
		files = append(files, embed)
	}
	if bp.Name == "main" {
		path = "main"
	}
	tp := &SourcePackage{
		Package: types.NewPackage(path, bp.Name),
		Files:   files,
		Dir:     dir,
		Context: ctx,
	}
	ctx.pkgs[path] = tp
	return tp, nil
}

func (ctx *Context) loadTestPackage(bp *build.Package, path string, dir string) (*SourcePackage, error) {
	if len(bp.TestGoFiles) == 0 && len(bp.XTestGoFiles) == 0 {
		return nil, ErrNoTestFiles
	}
	files, err := ctx.parseGoFiles(dir, append(append(bp.GoFiles, bp.CgoFiles...), bp.TestGoFiles...))
	if err != nil {
		return nil, err
	}
	embed, err := load.Embed(bp, ctx.FileSet, files, true, false)
	if err != nil {
		return nil, err
	}
	if embed != nil {
		files = append(files, embed)
	}
	// fix pkg name
	name := bp.Name
	if name == "main" {
		name = "main@test"
		for _, f := range files {
			f.Name.Name = name
		}
	}
	tp := &SourcePackage{
		Package: types.NewPackage(path, name),
		Files:   files,
		Dir:     dir,
		Context: ctx,
	}
	ctx.pkgs[path] = tp
	if len(bp.XTestGoFiles) > 0 {
		files, err := ctx.parseGoFiles(dir, bp.XTestGoFiles)
		if err != nil {
			return nil, err
		}
		embed, err := load.Embed(bp, ctx.FileSet, files, false, true)
		if err != nil {
			return nil, err
		}
		if embed != nil {
			files = append(files, embed)
		}
		tp := &SourcePackage{
			Package: types.NewPackage(path+"_test", bp.Name+"_test"),
			Files:   files,
			Dir:     dir,
			Context: ctx,
		}
		ctx.pkgs[path+"_test"] = tp
	}
	data, err := load.TestMain(&ctx.BuildContext, bp)
	if err != nil {
		return nil, err
	}
	f, err := parser.ParseFile(ctx.FileSet, "_testmain.go", data, parser.ParseComments)
	if err != nil {
		return nil, err
	}
	return &SourcePackage{
		Package: types.NewPackage(path+".test", "main"),
		Files:   []*ast.File{f},
		Dir:     dir,
		Context: ctx,
	}, nil
}

func (ctx *Context) openFile(path string) (io.ReadCloser, error) {
	if fn := ctx.BuildContext.OpenFile; fn != nil {
		return fn(path)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err // nil interface
	}
	return f, nil
}

func (ctx *Context) parseGoFiles(dir string, filenames []string) ([]*ast.File, error) {
	files := make([]*ast.File, len(filenames))
	errors := make([]error, len(filenames))

	var wg sync.WaitGroup
	wg.Add(len(filenames))
	for i, filename := range filenames {
		go func(i int, filepath string) {
			defer wg.Done()
			r, err := ctx.openFile(filepath)
			if err != nil {
				errors[i] = err
			} else {
				files[i], errors[i] = parser.ParseFile(ctx.FileSet, filepath, r, parser.ParseComments)
				r.Close()
			}
		}(i, filepath.Join(dir, filename))
	}
	wg.Wait()

	for _, err := range errors {
		if err != nil {
			return nil, err
		}
	}
	return files, nil
}

func (ctx *Context) LoadInterp(filename string, src interface{}) (*Interp, error) {
	pkg, err := ctx.LoadFile(filename, src)
	if err != nil {
		return nil, err
	}
	return ctx.NewInterp(pkg)
}

func (ctx *Context) LoadFile(filename string, src interface{}) (*ssa.Package, error) {
	file, err := ctx.ParseFile(filename, src)
	if err != nil {
		return nil, err
	}
	root, _ := filepath.Split(filename)
	ctx.setRoot(root)
	return ctx.LoadAstFile(file.Name.Name, file)
}

func (ctx *Context) ParseFile(filename string, src interface{}) (*ast.File, error) {
	if ext := filepath.Ext(filename); ext != "" {
		if fn, ok := sourceProcessor[ext]; ok {
			data, err := fn(ctx, filename, src)
			if err != nil {
				return nil, err
			}
			src = data
		}
	}
	return parser.ParseFile(ctx.FileSet, filename, src, parser.ParseComments)
}

func (ctx *Context) LoadAstFile(path string, file *ast.File) (*ssa.Package, error) {
	files := []*ast.File{file}
	if ctx.Mode&DisableCustomBuiltin == 0 {
		if f, err := ParseBuiltin(ctx.FileSet, file.Name.Name); err == nil {
			files = []*ast.File{f, file}
		}
	}
	dir, _ := filepath.Split(ctx.FileSet.Position(file.Package).Filename)
	if dir == "" {
		dir, _ = os.Getwd()
	}
	embed, err := load.EmbedFiles(file.Name.Name, filepath.Clean(dir), ctx.FileSet, files)
	if err != nil {
		return nil, err
	}
	if embed != nil {
		files = append(files, embed)
	}
	sp := &SourcePackage{
		Context: ctx,
		Package: types.NewPackage(path, file.Name.Name),
		Files:   files,
	}
	if err := sp.Load(); err != nil {
		return nil, err
	}
	ctx.pkgs[path] = sp
	return ctx.buildPackage(sp)
}

func (ctx *Context) LoadAstPackage(path string, apkg *ast.Package) (*ssa.Package, error) {
	var files []*ast.File
	for _, f := range apkg.Files {
		files = append(files, f)
	}
	if ctx.Mode&DisableCustomBuiltin == 0 {
		if f, err := ParseBuiltin(ctx.FileSet, apkg.Name); err == nil {
			files = append([]*ast.File{f}, files...)
		}
	}
	sp := &SourcePackage{
		Context: ctx,
		Package: types.NewPackage(path, apkg.Name),
		Files:   files,
	}
	err := sp.Load()
	if err != nil {
		return nil, err
	}
	return ctx.buildPackage(sp)
}

func (ctx *Context) RunPkg(mainPkg *ssa.Package, input string, args []string) (exitCode int, err error) {
	interp, err := ctx.NewInterp(mainPkg)
	if err != nil {
		return 2, err
	}
	return ctx.RunInterp(interp, input, args)
}

func (ctx *Context) RunInterp(interp *Interp, input string, args []string) (exitCode int, err error) {
	if ctx.RunContext != nil {
		return ctx.runInterpWithContext(interp, input, args, ctx.RunContext)
	}
	return ctx.runInterp(interp, input, args)
}

func (p *Context) runInterpWithContext(interp *Interp, input string, args []string, ctx context.Context) (int, error) {
	var exitCode int
	var err error
	ch := make(chan error, 1)
	interp.cherror = make(chan PanicError)
	go func() {
		exitCode, err = p.runInterp(interp, input, args)
		ch <- err
	}()
	select {
	case <-ctx.Done():
		interp.Abort()
		select {
		case <-time.After(1e9):
			exitCode = 2
			err = fmt.Errorf("interrupt timeout: all goroutines are asleep - deadlock!")
		case <-ch:
			err = ctx.Err()
		}
	case err = <-ch:
	case e := <-interp.cherror:
		switch v := e.Value.(type) {
		case exitPanic:
			exitCode = int(v)
		default:
			err = e
			interp.Abort()
			exitCode = 2
		}
	}
	return exitCode, err
}

func (ctx *Context) runInterp(interp *Interp, input string, args []string) (exitCode int, err error) {
	// reset os args and flag
	os.Args = []string{input}
	if args != nil {
		os.Args = append(os.Args, args...)
	}
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	if err = interp.RunInit(); err != nil {
		return 2, err
	}
	return interp.RunMain()
}

func (ctx *Context) RunFunc(mainPkg *ssa.Package, fnname string, args ...Value) (ret Value, err error) {
	interp, err := ctx.NewInterp(mainPkg)
	if err != nil {
		return nil, err
	}
	return interp.RunFunc(fnname, args...)
}

func (ctx *Context) NewInterp(mainPkg *ssa.Package) (*Interp, error) {
	return NewInterp(ctx, mainPkg)
}

var (
	testInit bool // fix testing second loading error "flag provided but not defined"
)

func (ctx *Context) TestPkg(pkg *ssa.Package, input string, args []string) error {
	var failed bool
	start := time.Now()
	defer func() {
		sec := time.Since(start).Seconds()
		if failed {
			fmt.Fprintf(os.Stdout, "FAIL\t%s %0.3fs\n", pkg.Pkg.Path(), sec)
		} else {
			fmt.Fprintf(os.Stdout, "ok\t%s %0.3fs\n", pkg.Pkg.Path(), sec)
		}
	}()
	os.Args = []string{input}
	if args != nil && !testInit {
		os.Args = append(os.Args, args...)
	}
	testInit = true
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	interp, err := NewInterp(ctx, pkg)
	if err != nil {
		failed = true
		fmt.Printf("create interp failed: %v\n", err)
	}
	if err = interp.RunInit(); err != nil {
		failed = true
		fmt.Printf("init error: %v\n", err)
	}
	exitCode, _ := interp.RunMain()
	if exitCode != 0 {
		failed = true
	}
	if failed {
		return ErrTestFailed
	}
	return nil
}

func (ctx *Context) RunFile(filename string, src interface{}, args []string) (exitCode int, err error) {
	pkg, err := ctx.LoadFile(filename, src)
	if err != nil {
		return 2, err
	}
	return ctx.RunPkg(pkg, filename, args)
}

func (ctx *Context) Run(path string, args []string) (exitCode int, err error) {
	if strings.HasSuffix(path, ".go") {
		return ctx.RunFile(path, nil, args)
	}
	pkg, err := ctx.LoadDir(path, false)
	if err != nil {
		return 2, err
	}
	if !isMainPkg(pkg) {
		return 2, ErrNotFoundMain
	}
	return ctx.RunPkg(pkg, path, args)
}

func isMainPkg(pkg *ssa.Package) bool {
	return pkg.Pkg.Name() == "main" && pkg.Func("main") != nil
}

func (ctx *Context) RunTest(dir string, args []string) error {
	pkg, err := ctx.LoadDir(dir, true)
	if err != nil {
		if err == ErrNoTestFiles {
			fmt.Println("?", err)
			return nil
		}
		return err
	}
	if filepath.IsAbs(dir) {
		os.Chdir(dir)
	}
	return ctx.TestPkg(pkg, dir, args)
}

// SSABuilder builds SSA (Static Single Assignment) form packages incrementally.
// It maintains a program and tracks which packages have already been created
// to avoid duplicate builds.
type SSABuilder struct {
	ctx     *Context
	prog    *ssa.Program
	created map[*types.Package]bool
}

func NewSSABuilder(ctx *Context) *SSABuilder {
	return &SSABuilder{
		ctx:     ctx,
		prog:    ssa.NewProgram(ctx.FileSet, ctx.BuilderMode|ssa.InstantiateGenerics),
		created: make(map[*types.Package]bool),
	}
}

func (b *SSABuilder) Reset() {
	ctx := b.ctx
	b.prog = ssa.NewProgram(ctx.FileSet, ctx.BuilderMode|ssa.InstantiateGenerics)
	b.created = make(map[*types.Package]bool)
}

func (b *SSABuilder) PrebuildSSA(packages ...string) error {
	ctx := b.ctx
	var pkgs []*types.Package
	for _, path := range packages {
		pkg, err := ctx.Importer.Import(path)
		if err != nil {
			return err
		}
		pkgs = append(pkgs, pkg)
	}
	b.BuildPackages(pkgs)
	return nil
}

func (b *SSABuilder) BuildImports(pkg *types.Package) {
	b.BuildPackages(pkg.Imports())
	if addin := b.checkPatch(pkg); addin != nil {
		b.BuildPackages(addin)
	}
}

func (b *SSABuilder) checkPatch(p *types.Package) (addin []*types.Package) {
	ctx := b.ctx
	for _, im := range p.Imports() {
		patch := im.Path() + "@patch"
		if p.Path() == patch {
			// skip p.path is pkg@patch
			continue
		}
		if pkg, ok := ctx.pkgs[patch]; ok && pkg.Loaded() {
			addin = append(addin, pkg.Package)
		}
	}
	return addin
}

func (b *SSABuilder) BuildPackages(pkgs []*types.Package) {
	ctx := b.ctx
	for _, p := range pkgs {
		if !b.created[p] {
			b.created[p] = true
			if pkg, ok := ctx.pkgs[p.Path()]; ok {
				if !pkg.Loaded() {
					continue
				}
				b.created[pkg.Package] = true
				b.BuildImports(pkg.Package)
				if ctx.Mode&EnableDumpImports != 0 {
					if pkg.Dir != "" {
						fmt.Println("# source", p.Path(), "<"+pkg.Dir+">")
					} else if pkg.Register {
						fmt.Println("# package", p.Path(), "<source>")
					} else if strings.HasSuffix(p.Path(), "@patch") {
						fmt.Println("# package", p.Path(), "<source>")
					} else {
						fmt.Println("# source", p.Path(), "<memory>")
					}
				}
				b.prog.CreatePackage(pkg.Package, pkg.Files, pkg.Info, true).Build()
				ctx.checkNested(pkg.Package, pkg.Info)
			} else {
				b.BuildImports(p)
				var indirect bool
				if !p.Complete() {
					indirect = true
					p.MarkComplete()
				}
				if ctx.Mode&EnableDumpImports != 0 {
					if indirect {
						fmt.Println("# virtual", p.Path())
					} else {
						fmt.Println("# package", p.Path())
					}
				}
				b.prog.CreatePackage(p, nil, nil, true).Build()
			}
		}
	}
}

func (b *SSABuilder) BuildMain(sp *SourcePackage) (pkg *ssa.Package, err error) {
	b.BuildImports(sp.Package)
	ctx := b.ctx
	if ctx.Mode&EnableDumpImports != 0 {
		if sp.Dir != "" {
			fmt.Println("# package", sp.Package.Path(), sp.Dir)
		} else {
			fmt.Println("# package", sp.Package.Path(), "<source>")
		}
	}
	// Create and build the primary package.
	pkg = b.prog.CreatePackage(sp.Package, sp.Files, sp.Info, false)
	pkg.Build()
	ctx.checkNested(sp.Package, sp.Info)
	return
}

func (ctx *Context) PrebuildSSA(pkgs ...string) (err error) {
	if ctx.Mode&DisableRecover == 0 {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("prebuild SSA error: %v", e)
			}
		}()
	}
	return ctx.Builder.PrebuildSSA(pkgs...)
}

func (ctx *Context) buildPackage(sp *SourcePackage) (pkg *ssa.Package, err error) {
	if ctx.Mode&DisableRecover == 0 {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("build SSA package error: %v", e)
			}
		}()
	}
	defer ctx.Builder.Reset()
	return ctx.Builder.BuildMain(sp)
}

func (ctx *Context) checkNested(pkg *types.Package, info *types.Info) {
	var nestedList []*types.Named
	for k, v := range info.Scopes {
		switch k.(type) {
		case *ast.BlockStmt, *ast.FuncType:
			for _, name := range v.Names() {
				obj := v.Lookup(name)
				if named, ok := obj.Type().(*types.Named); ok && named.Obj().Pkg() == pkg {
					nestedList = append(nestedList, named)
				}
			}
		}
	}
	if len(nestedList) == 0 {
		return
	}
	sort.Slice(nestedList, func(i, j int) bool {
		return nestedList[i].Obj().Pos() < nestedList[j].Obj().Pos()
	})
	for i, named := range nestedList {
		ctx.nestedMap[named] = i + 1
	}
}

func RunFile(filename string, src interface{}, args []string, mode Mode) (exitCode int, err error) {
	ctx := NewContext(mode)
	return ctx.RunFile(filename, src, args)
}

func Run(path string, args []string, mode Mode) (exitCode int, err error) {
	ctx := NewContext(mode)
	return ctx.Run(path, args)
}

func RunTest(path string, args []string, mode Mode) error {
	ctx := NewContext(mode)
	return ctx.RunTest(path, args)
}

var (
	builtinPkg = &Package{
		Name:          "builtin",
		Path:          "github.com/goplus/ixgo/builtin",
		Deps:          make(map[string]string),
		Interfaces:    map[string]reflect.Type{},
		NamedTypes:    map[string]reflect.Type{},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]TypedConst{},
		UntypedConsts: map[string]UntypedConst{},
	}
	builtinPrefix = "Builtin_"
)

func init() {
	RegisterPackage(builtinPkg)
}

func RegisterCustomBuiltin(key string, fn interface{}) error {
	v := reflect.ValueOf(fn)
	switch v.Kind() {
	case reflect.Func:
		if !strings.HasPrefix(key, builtinPrefix) {
			key = builtinPrefix + key
		}
		builtinPkg.Funcs[key] = v
		typ := v.Type()
		for i := 0; i < typ.NumIn(); i++ {
			checkBuiltinDeps(typ.In(i))
		}
		for i := 0; i < typ.NumOut(); i++ {
			checkBuiltinDeps(typ.Out(i))
		}
		return nil
	}
	return ErrNoFunction
}

func checkBuiltinDeps(typ reflect.Type) {
	if path := typ.PkgPath(); path != "" {
		v := strings.Split(path, "/")
		builtinPkg.Deps[path] = v[len(v)-1]
	}
}

func builtinFuncList() []string {
	var list []string
	for k, v := range builtinPkg.Funcs {
		if strings.HasPrefix(k, builtinPrefix) {
			name := k[len(builtinPrefix):]
			typ := v.Type()
			var ins []string
			var outs []string
			var call []string
			numIn := typ.NumIn()
			numOut := typ.NumOut()
			if typ.IsVariadic() {
				for i := 0; i < numIn-1; i++ {
					ins = append(ins, fmt.Sprintf("p%v %v", i, typ.In(i).String()))
					call = append(call, fmt.Sprintf("p%v", i))
				}
				ins = append(ins, fmt.Sprintf("p%v ...%v", numIn-1, typ.In(numIn-1).Elem().String()))
				call = append(call, fmt.Sprintf("p%v...", numIn-1))
			} else {
				for i := 0; i < numIn; i++ {
					ins = append(ins, fmt.Sprintf("p%v %v", i, typ.In(i).String()))
					call = append(call, fmt.Sprintf("p%v", i))
				}
			}
			for i := 0; i < numOut; i++ {
				outs = append(outs, typ.Out(i).String())
			}
			var str string
			if numOut > 0 {
				str = fmt.Sprintf("func %v(%v)(%v) { return builtin.%v(%v) }",
					name, strings.Join(ins, ","), strings.Join(outs, ","), k, strings.Join(call, ","))
			} else {
				str = fmt.Sprintf("func %v(%v) { builtin.%v(%v) }",
					name, strings.Join(ins, ","), k, strings.Join(call, ","))
			}
			list = append(list, str)
		}
	}
	return list
}

func ParseBuiltin(fset *token.FileSet, pkg string) (*ast.File, error) {
	list := builtinFuncList()
	if len(list) == 0 {
		return nil, os.ErrInvalid
	}
	var deps []string
	for k := range builtinPkg.Deps {
		deps = append(deps, strconv.Quote(k))
	}
	sort.Strings(deps)
	src := fmt.Sprintf(`package %v
import (
	"github.com/goplus/ixgo/builtin"
	%v
)
%v
`, pkg, strings.Join(deps, "\n"), strings.Join(list, "\n"))
	return parser.ParseFile(fset, "gossa_builtin.go", src, 0)
}
