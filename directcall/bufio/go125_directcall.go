// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.25 && !go1.26
// +build go1.25,!go1.26

package bufio

import (
	q "bufio"

	"github.com/goplus/ixgo"
	io "io"
)

func init() {
	ixgo.RegisterDirectCalls("bufio", map[string]ixgo.DirectCallAdapter{
		"(*Reader).Buffered":        method_ptr_Reader_Buffered,
		"(*Reader).Reset":           method_ptr_Reader_Reset,
		"(*Reader).Size":            method_ptr_Reader_Size,
		"(*Reader).UnreadByte":      method_ptr_Reader_UnreadByte,
		"(*Reader).UnreadRune":      method_ptr_Reader_UnreadRune,
		"(*Scanner).Buffer":         method_ptr_Scanner_Buffer,
		"(*Scanner).Bytes":          method_ptr_Scanner_Bytes,
		"(*Scanner).Err":            method_ptr_Scanner_Err,
		"(*Scanner).Scan":           method_ptr_Scanner_Scan,
		"(*Scanner).Split":          method_ptr_Scanner_Split,
		"(*Scanner).Text":           method_ptr_Scanner_Text,
		"(*Writer).Available":       method_ptr_Writer_Available,
		"(*Writer).AvailableBuffer": method_ptr_Writer_AvailableBuffer,
		"(*Writer).Buffered":        method_ptr_Writer_Buffered,
		"(*Writer).Flush":           method_ptr_Writer_Flush,
		"(*Writer).Reset":           method_ptr_Writer_Reset,
		"(*Writer).Size":            method_ptr_Writer_Size,
		"(*Writer).WriteByte":       method_ptr_Writer_WriteByte,
		"NewReadWriter":             func_NewReadWriter,
		"NewReader":                 func_NewReader,
		"NewReaderSize":             func_NewReaderSize,
		"NewScanner":                func_NewScanner,
		"NewWriter":                 func_NewWriter,
		"NewWriterSize":             func_NewWriterSize,
	})
}

func func_NewReadWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReadWriter(ixgo.DirectCallArg[*q.Reader](ctx, 0), ixgo.DirectCallArg[*q.Writer](ctx, 1)))
}

func func_NewReader(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReader(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewReaderSize(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewReaderSize(ixgo.DirectCallArg[io.Reader](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_NewScanner(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewScanner(ixgo.DirectCallArg[io.Reader](ctx, 0)))
}

func func_NewWriter(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriter(ixgo.DirectCallArg[io.Writer](ctx, 0)))
}

func func_NewWriterSize(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.NewWriterSize(ixgo.DirectCallArg[io.Writer](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func method_ptr_Reader_Buffered(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).Buffered(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_Reset(ctx ixgo.DirectCallContext) {
	(*q.Reader).Reset(ixgo.DirectCallArg[*q.Reader](ctx, 0), ixgo.DirectCallArg[io.Reader](ctx, 1))
}

func method_ptr_Reader_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).Size(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_UnreadByte(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).UnreadByte(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Reader_UnreadRune(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Reader).UnreadRune(ixgo.DirectCallArg[*q.Reader](ctx, 0)))
}

func method_ptr_Scanner_Buffer(ctx ixgo.DirectCallContext) {
	(*q.Scanner).Buffer(ixgo.DirectCallArg[*q.Scanner](ctx, 0), ixgo.DirectCallArg[[]byte](ctx, 1), ixgo.DirectCallArg[int](ctx, 2))
}

func method_ptr_Scanner_Bytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).Bytes(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func method_ptr_Scanner_Err(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).Err(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func method_ptr_Scanner_Scan(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).Scan(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func method_ptr_Scanner_Split(ctx ixgo.DirectCallContext) {
	(*q.Scanner).Split(ixgo.DirectCallArg[*q.Scanner](ctx, 0), ixgo.DirectCallArg[q.SplitFunc](ctx, 1))
}

func method_ptr_Scanner_Text(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Scanner).Text(ixgo.DirectCallArg[*q.Scanner](ctx, 0)))
}

func method_ptr_Writer_Available(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Available(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_AvailableBuffer(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).AvailableBuffer(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Buffered(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Buffered(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Flush(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Flush(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_Reset(ctx ixgo.DirectCallContext) {
	(*q.Writer).Reset(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[io.Writer](ctx, 1))
}

func method_ptr_Writer_Size(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).Size(ixgo.DirectCallArg[*q.Writer](ctx, 0)))
}

func method_ptr_Writer_WriteByte(ctx ixgo.DirectCallContext) {
	ctx.SetResult((*q.Writer).WriteByte(ixgo.DirectCallArg[*q.Writer](ctx, 0), ixgo.DirectCallArg[byte](ctx, 1)))
}
