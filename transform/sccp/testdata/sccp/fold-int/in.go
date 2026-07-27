package test

var int8Sink int8
var uint8Sink uint8
var uint32Sink uint32
var uint64Sink uint64
var uintptrSink uintptr
var boolSink bool

func signed(condition bool) {
	left, right, shift := int8(7), int8(3), int(2)
	if condition {
		left, right, shift = 7, 3, 2
	}
	int8Sink = left + right
	int8Sink = left - right
	int8Sink = left * right
	int8Sink = left / right
	int8Sink = left % right
	int8Sink = left & right
	int8Sink = left | right
	int8Sink = left ^ right
	int8Sink = left &^ right
	int8Sink = -left
	int8Sink = ^left
	int8Sink = left << shift
	int8Sink = left >> shift
	boolSink = left == right
	boolSink = left != right
	boolSink = left < right
	boolSink = left <= right
	boolSink = left > right
	boolSink = left >= right
}

func widths(condition bool) {
	value32, value64, pointer := uint32(1), uint64(2), uintptr(3)
	if condition {
		value32, value64, pointer = 1, 2, 3
	}
	uint32Sink = value32 + 1
	uint64Sink = value64 + 1
	uintptrSink = pointer + 1
}

func unsigned(condition bool) {
	left, right, shift := uint8(7), uint8(3), uint(2)
	if condition {
		left, right, shift = 7, 3, 2
	}
	uint8Sink = left + right
	uint8Sink = left - right
	uint8Sink = left * right
	uint8Sink = left / right
	uint8Sink = left % right
	uint8Sink = left & right
	uint8Sink = left | right
	uint8Sink = left ^ right
	uint8Sink = left &^ right
	uint8Sink = -left
	uint8Sink = ^left
	uint8Sink = left << shift
	uint8Sink = left >> shift
	boolSink = left == right
	boolSink = left != right
	boolSink = left < right
	boolSink = left <= right
	boolSink = left > right
	boolSink = left >= right
}

func booleans(condition bool) {
	left, right := true, false
	if condition {
		left, right = true, false
	}
	boolSink = !left
	boolSink = left == right
	boolSink = left != right
}

func unsignedZero(condition bool) {
	zero := uint8(0)
	if condition {
		zero = 0
	}
	uint8Sink = 1 / zero
	uint8Sink = 1 % zero
}

func signedZero(condition bool) {
	zero := int8(0)
	if condition {
		zero = 0
	}
	int8Sink = 1 % zero
}

func negativeShift(condition bool) {
	shift := -1
	if condition {
		shift = -1
	}
	int8Sink = int8(1) << shift
}

func dynamic(input int8) {
	int8Sink = -input
	int8Sink = ^input
}

func capture(input int) func() int {
	return func() int {
		condition := true
		if input > 0 {
			condition = true
		}
		if condition {
			return input + 1
		}
		return input
	}
}

func external()

func target() int { return 0 }
