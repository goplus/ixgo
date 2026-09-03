// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package os

import (
	q "os"

	"github.com/goplus/ixgo"
	fs "io/fs"
	time "time"
)

func init() {
	ixgo.RegisterDirectCalls("os", map[string]ixgo.DirectCallAdapter{
		"(*File).Chdir":              method_ptr_File_Chdir,
		"(*File).Chmod":              method_ptr_File_Chmod,
		"(*File).Chown":              method_ptr_File_Chown,
		"(*File).Close":              method_ptr_File_Close,
		"(*File).Fd":                 method_ptr_File_Fd,
		"(*File).Name":               method_ptr_File_Name,
		"(*File).SetDeadline":        method_ptr_File_SetDeadline,
		"(*File).SetReadDeadline":    method_ptr_File_SetReadDeadline,
		"(*File).SetWriteDeadline":   method_ptr_File_SetWriteDeadline,
		"(*File).Sync":               method_ptr_File_Sync,
		"(*File).Truncate":           method_ptr_File_Truncate,
		"(*LinkError).Error":         method_ptr_LinkError_Error,
		"(*LinkError).Unwrap":        method_ptr_LinkError_Unwrap,
		"(*Process).Kill":            method_ptr_Process_Kill,
		"(*Process).Release":         method_ptr_Process_Release,
		"(*Process).Signal":          method_ptr_Process_Signal,
		"(*ProcessState).ExitCode":   method_ptr_ProcessState_ExitCode,
		"(*ProcessState).Exited":     method_ptr_ProcessState_Exited,
		"(*ProcessState).Pid":        method_ptr_ProcessState_Pid,
		"(*ProcessState).String":     method_ptr_ProcessState_String,
		"(*ProcessState).Success":    method_ptr_ProcessState_Success,
		"(*ProcessState).Sys":        method_ptr_ProcessState_Sys,
		"(*ProcessState).SysUsage":   method_ptr_ProcessState_SysUsage,
		"(*ProcessState).SystemTime": method_ptr_ProcessState_SystemTime,
		"(*ProcessState).UserTime":   method_ptr_ProcessState_UserTime,
		"(*Root).Chmod":              method_ptr_Root_Chmod,
		"(*Root).Chown":              method_ptr_Root_Chown,
		"(*Root).Chtimes":            method_ptr_Root_Chtimes,
		"(*Root).Close":              method_ptr_Root_Close,
		"(*Root).FS":                 method_ptr_Root_FS,
		"(*Root).Lchown":             method_ptr_Root_Lchown,
		"(*Root).Link":               method_ptr_Root_Link,
		"(*Root).Mkdir":              method_ptr_Root_Mkdir,
		"(*Root).MkdirAll":           method_ptr_Root_MkdirAll,
		"(*Root).Name":               method_ptr_Root_Name,
		"(*Root).Remove":             method_ptr_Root_Remove,
		"(*Root).RemoveAll":          method_ptr_Root_RemoveAll,
		"(*Root).Rename":             method_ptr_Root_Rename,
		"(*Root).Symlink":            method_ptr_Root_Symlink,
		"(*Root).WriteFile":          method_ptr_Root_WriteFile,
		"(*SyscallError).Error":      method_ptr_SyscallError_Error,
		"(*SyscallError).Timeout":    method_ptr_SyscallError_Timeout,
		"(*SyscallError).Unwrap":     method_ptr_SyscallError_Unwrap,
		"Chdir":                      func_Chdir,
		"Chmod":                      func_Chmod,
		"Chown":                      func_Chown,
		"Chtimes":                    func_Chtimes,
		"Clearenv":                   func_Clearenv,
		"CopyFS":                     func_CopyFS,
		"DirFS":                      func_DirFS,
		"Environ":                    func_Environ,
		"Exit":                       func_Exit,
		"Expand":                     func_Expand,
		"ExpandEnv":                  func_ExpandEnv,
		"Getegid":                    func_Getegid,
		"Getenv":                     func_Getenv,
		"Geteuid":                    func_Geteuid,
		"Getgid":                     func_Getgid,
		"Getpagesize":                func_Getpagesize,
		"Getpid":                     func_Getpid,
		"Getppid":                    func_Getppid,
		"Getuid":                     func_Getuid,
		"IsExist":                    func_IsExist,
		"IsNotExist":                 func_IsNotExist,
		"IsPathSeparator":            func_IsPathSeparator,
		"IsPermission":               func_IsPermission,
		"IsTimeout":                  func_IsTimeout,
		"Lchown":                     func_Lchown,
		"Link":                       func_Link,
		"Mkdir":                      func_Mkdir,
		"MkdirAll":                   func_MkdirAll,
		"NewFile":                    func_NewFile,
		"NewSyscallError":            func_NewSyscallError,
		"Remove":                     func_Remove,
		"RemoveAll":                  func_RemoveAll,
		"Rename":                     func_Rename,
		"SameFile":                   func_SameFile,
		"Setenv":                     func_Setenv,
		"Symlink":                    func_Symlink,
		"TempDir":                    func_TempDir,
		"Truncate":                   func_Truncate,
		"Unsetenv":                   func_Unsetenv,
		"WriteFile":                  func_WriteFile,
	})
}

func func_Chdir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Chdir(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Chmod(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Chmod(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[fs.FileMode](ctx, 1)))
}

func func_Chown(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Chown(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_Chtimes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Chtimes(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1), ixgo.DirectCallArg[time.Time](ctx, 2)))
}

func func_Clearenv(ctx ixgo.DirectCallContext) {
	q.Clearenv()
}

func func_CopyFS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.CopyFS(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[fs.FS](ctx, 1)))
}

func func_DirFS(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.DirFS(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Environ(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Environ())
}

func func_Exit(ctx ixgo.DirectCallContext) {
	q.Exit(ixgo.DirectCallArg[int](ctx, 0))
}

func func_Expand(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Expand(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[func(string) string](ctx, 1)))
}

func func_ExpandEnv(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ExpandEnv(ixgo.DirectCallArg[string](ctx, 0)))
}

func method_ptr_File_Chdir(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Chdir(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Chmod(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Chmod(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[fs.FileMode](ctx, 1)))
}

func method_ptr_File_Chown(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Chown(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func method_ptr_File_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Close(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Fd(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Fd(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Name(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_SetDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).SetDeadline(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_File_SetReadDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).SetReadDeadline(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_File_SetWriteDeadline(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).SetWriteDeadline(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[time.Time](ctx, 1)))
}

func method_ptr_File_Sync(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Sync(ixgo.DirectCallArg[*q.File](ctx, 0)))
}

func method_ptr_File_Truncate(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.File).Truncate(ixgo.DirectCallArg[*q.File](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_Getegid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Getegid())
}

func func_Getenv(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Getenv(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Geteuid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Geteuid())
}

func func_Getgid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Getgid())
}

func func_Getpagesize(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Getpagesize())
}

func func_Getpid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Getpid())
}

func func_Getppid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Getppid())
}

func func_Getuid(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Getuid())
}

func func_IsExist(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsExist(ixgo.DirectCallArg[error](ctx, 0)))
}

func func_IsNotExist(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsNotExist(ixgo.DirectCallArg[error](ctx, 0)))
}

func func_IsPathSeparator(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsPathSeparator(ixgo.DirectCallArg[uint8](ctx, 0)))
}

func func_IsPermission(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsPermission(ixgo.DirectCallArg[error](ctx, 0)))
}

func func_IsTimeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.IsTimeout(ixgo.DirectCallArg[error](ctx, 0)))
}

func func_Lchown(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Lchown(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[int](ctx, 1), ixgo.DirectCallArg[int](ctx, 2)))
}

func func_Link(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Link(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_LinkError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LinkError).Error(ixgo.DirectCallArg[*q.LinkError](ctx, 0)))
}

func method_ptr_LinkError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.LinkError).Unwrap(ixgo.DirectCallArg[*q.LinkError](ctx, 0)))
}

func func_Mkdir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Mkdir(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[fs.FileMode](ctx, 1)))
}

func func_MkdirAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.MkdirAll(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[fs.FileMode](ctx, 1)))
}

func func_NewFile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewFile(ixgo.DirectCallArg[uintptr](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_NewSyscallError(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewSyscallError(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[error](ctx, 1)))
}

func method_ptr_Process_Kill(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Process).Kill(ixgo.DirectCallArg[*q.Process](ctx, 0)))
}

func method_ptr_Process_Release(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Process).Release(ixgo.DirectCallArg[*q.Process](ctx, 0)))
}

func method_ptr_Process_Signal(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Process).Signal(ixgo.DirectCallArg[*q.Process](ctx, 0), ixgo.DirectCallArg[q.Signal](ctx, 1)))
}

func method_ptr_ProcessState_ExitCode(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcessState).ExitCode(ixgo.DirectCallArg[*q.ProcessState](ctx, 0)))
}

func method_ptr_ProcessState_Exited(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcessState).Exited(ixgo.DirectCallArg[*q.ProcessState](ctx, 0)))
}

func method_ptr_ProcessState_Pid(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcessState).Pid(ixgo.DirectCallArg[*q.ProcessState](ctx, 0)))
}

func method_ptr_ProcessState_String(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcessState).String(ixgo.DirectCallArg[*q.ProcessState](ctx, 0)))
}

func method_ptr_ProcessState_Success(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcessState).Success(ixgo.DirectCallArg[*q.ProcessState](ctx, 0)))
}

func method_ptr_ProcessState_Sys(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcessState).Sys(ixgo.DirectCallArg[*q.ProcessState](ctx, 0)))
}

func method_ptr_ProcessState_SysUsage(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcessState).SysUsage(ixgo.DirectCallArg[*q.ProcessState](ctx, 0)))
}

func method_ptr_ProcessState_SystemTime(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcessState).SystemTime(ixgo.DirectCallArg[*q.ProcessState](ctx, 0)))
}

func method_ptr_ProcessState_UserTime(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.ProcessState).UserTime(ixgo.DirectCallArg[*q.ProcessState](ctx, 0)))
}

func func_Remove(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Remove(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_RemoveAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RemoveAll(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_Rename(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rename(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Root_Chmod(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Chmod(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[fs.FileMode](ctx, 2)))
}

func method_ptr_Root_Chown(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Chown(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func method_ptr_Root_Chtimes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Chtimes(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[time.Time](ctx, 2), ixgo.DirectCallArg[time.Time](ctx, 3)))
}

func method_ptr_Root_Close(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Close(ixgo.DirectCallArg[*q.Root](ctx, 0)))
}

func method_ptr_Root_FS(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).FS(ixgo.DirectCallArg[*q.Root](ctx, 0)))
}

func method_ptr_Root_Lchown(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Lchown(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[int](ctx, 2), ixgo.DirectCallArg[int](ctx, 3)))
}

func method_ptr_Root_Link(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Link(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Root_Mkdir(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Mkdir(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[fs.FileMode](ctx, 2)))
}

func method_ptr_Root_MkdirAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).MkdirAll(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[fs.FileMode](ctx, 2)))
}

func method_ptr_Root_Name(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Name(ixgo.DirectCallArg[*q.Root](ctx, 0)))
}

func method_ptr_Root_Remove(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Remove(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Root_RemoveAll(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).RemoveAll(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_Root_Rename(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Rename(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Root_Symlink(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).Symlink(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[string](ctx, 2)))
}

func method_ptr_Root_WriteFile(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Root).WriteFile(ixgo.DirectCallArg[*q.Root](ctx, 0), ixgo.DirectCallArg[string](ctx, 1), ixgo.DirectCallArg[[]byte](ctx, 2), ixgo.DirectCallArg[fs.FileMode](ctx, 3)))
}

func func_SameFile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.SameFile(ixgo.DirectCallArg[fs.FileInfo](ctx, 0), ixgo.DirectCallArg[fs.FileInfo](ctx, 1)))
}

func func_Setenv(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Setenv(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func func_Symlink(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Symlink(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[string](ctx, 1)))
}

func method_ptr_SyscallError_Error(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SyscallError).Error(ixgo.DirectCallArg[*q.SyscallError](ctx, 0)))
}

func method_ptr_SyscallError_Timeout(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SyscallError).Timeout(ixgo.DirectCallArg[*q.SyscallError](ctx, 0)))
}

func method_ptr_SyscallError_Unwrap(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.SyscallError).Unwrap(ixgo.DirectCallArg[*q.SyscallError](ctx, 0)))
}

func func_TempDir(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TempDir())
}

func func_Truncate(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Truncate(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[int64](ctx, 1)))
}

func func_Unsetenv(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Unsetenv(ixgo.DirectCallArg[string](ctx, 0)))
}

func func_WriteFile(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.WriteFile(ixgo.DirectCallArg[string](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[fs.FileMode](ctx, 2)))
}
