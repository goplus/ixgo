// export by github.com/goplus/ixgo/cmd/qexp

//go:build go1.27
// +build go1.27

package bits

import (
	q "math/bits"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterDirectCalls("math/bits", map[string]ixgo.DirectCallAdapter{
		"LeadingZeros":    func_LeadingZeros,
		"LeadingZeros16":  func_LeadingZeros16,
		"LeadingZeros32":  func_LeadingZeros32,
		"LeadingZeros64":  func_LeadingZeros64,
		"LeadingZeros8":   func_LeadingZeros8,
		"Len":             func_Len,
		"Len16":           func_Len16,
		"Len32":           func_Len32,
		"Len64":           func_Len64,
		"Len8":            func_Len8,
		"OnesCount":       func_OnesCount,
		"OnesCount16":     func_OnesCount16,
		"OnesCount32":     func_OnesCount32,
		"OnesCount64":     func_OnesCount64,
		"OnesCount8":      func_OnesCount8,
		"Rem":             func_Rem,
		"Rem32":           func_Rem32,
		"Rem64":           func_Rem64,
		"Reverse":         func_Reverse,
		"Reverse16":       func_Reverse16,
		"Reverse32":       func_Reverse32,
		"Reverse64":       func_Reverse64,
		"Reverse8":        func_Reverse8,
		"ReverseBytes":    func_ReverseBytes,
		"ReverseBytes16":  func_ReverseBytes16,
		"ReverseBytes32":  func_ReverseBytes32,
		"ReverseBytes64":  func_ReverseBytes64,
		"RotateLeft":      func_RotateLeft,
		"RotateLeft16":    func_RotateLeft16,
		"RotateLeft32":    func_RotateLeft32,
		"RotateLeft64":    func_RotateLeft64,
		"RotateLeft8":     func_RotateLeft8,
		"TrailingZeros":   func_TrailingZeros,
		"TrailingZeros16": func_TrailingZeros16,
		"TrailingZeros32": func_TrailingZeros32,
		"TrailingZeros64": func_TrailingZeros64,
		"TrailingZeros8":  func_TrailingZeros8,
	})
}

func func_LeadingZeros(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LeadingZeros(ixgo.DirectCallArg[uint](ctx, 0)))
}

func func_LeadingZeros16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LeadingZeros16(ixgo.DirectCallArg[uint16](ctx, 0)))
}

func func_LeadingZeros32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LeadingZeros32(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_LeadingZeros64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LeadingZeros64(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_LeadingZeros8(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.LeadingZeros8(ixgo.DirectCallArg[uint8](ctx, 0)))
}

func func_Len(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Len(ixgo.DirectCallArg[uint](ctx, 0)))
}

func func_Len16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Len16(ixgo.DirectCallArg[uint16](ctx, 0)))
}

func func_Len32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Len32(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_Len64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Len64(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_Len8(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Len8(ixgo.DirectCallArg[uint8](ctx, 0)))
}

func func_OnesCount(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OnesCount(ixgo.DirectCallArg[uint](ctx, 0)))
}

func func_OnesCount16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OnesCount16(ixgo.DirectCallArg[uint16](ctx, 0)))
}

func func_OnesCount32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OnesCount32(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_OnesCount64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OnesCount64(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_OnesCount8(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.OnesCount8(ixgo.DirectCallArg[uint8](ctx, 0)))
}

func func_Rem(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rem(ixgo.DirectCallArg[uint](ctx, 0), ixgo.DirectCallArg[uint](ctx, 1), ixgo.DirectCallArg[uint](ctx, 2)))
}

func func_Rem32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rem32(ixgo.DirectCallArg[uint32](ctx, 0), ixgo.DirectCallArg[uint32](ctx, 1), ixgo.DirectCallArg[uint32](ctx, 2)))
}

func func_Rem64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Rem64(ixgo.DirectCallArg[uint64](ctx, 0), ixgo.DirectCallArg[uint64](ctx, 1), ixgo.DirectCallArg[uint64](ctx, 2)))
}

func func_Reverse(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Reverse(ixgo.DirectCallArg[uint](ctx, 0)))
}

func func_Reverse16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Reverse16(ixgo.DirectCallArg[uint16](ctx, 0)))
}

func func_Reverse32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Reverse32(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_Reverse64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Reverse64(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_Reverse8(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.Reverse8(ixgo.DirectCallArg[uint8](ctx, 0)))
}

func func_ReverseBytes(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ReverseBytes(ixgo.DirectCallArg[uint](ctx, 0)))
}

func func_ReverseBytes16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ReverseBytes16(ixgo.DirectCallArg[uint16](ctx, 0)))
}

func func_ReverseBytes32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ReverseBytes32(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_ReverseBytes64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.ReverseBytes64(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_RotateLeft(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RotateLeft(ixgo.DirectCallArg[uint](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_RotateLeft16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RotateLeft16(ixgo.DirectCallArg[uint16](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_RotateLeft32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RotateLeft32(ixgo.DirectCallArg[uint32](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_RotateLeft64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RotateLeft64(ixgo.DirectCallArg[uint64](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_RotateLeft8(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.RotateLeft8(ixgo.DirectCallArg[uint8](ctx, 0), ixgo.DirectCallArg[int](ctx, 1)))
}

func func_TrailingZeros(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrailingZeros(ixgo.DirectCallArg[uint](ctx, 0)))
}

func func_TrailingZeros16(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrailingZeros16(ixgo.DirectCallArg[uint16](ctx, 0)))
}

func func_TrailingZeros32(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrailingZeros32(ixgo.DirectCallArg[uint32](ctx, 0)))
}

func func_TrailingZeros64(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrailingZeros64(ixgo.DirectCallArg[uint64](ctx, 0)))
}

func func_TrailingZeros8(ctx ixgo.DirectCallContext) {
	ctx.SetResult(q.TrailingZeros8(ixgo.DirectCallArg[uint8](ctx, 0)))
}
