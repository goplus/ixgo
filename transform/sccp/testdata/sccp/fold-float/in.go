package test

var float32Sink float32
var float64Sink float64
var boolSink bool

func fold32(condition bool) {
	left, right := float32(7.5), float32(2.5)
	if condition {
		left, right = 7.5, 2.5
	}
	float32Sink = left + right
	float32Sink = left - right
	float32Sink = left * right
	float32Sink = left / right
	float32Sink = -left
	boolSink = left == right
	boolSink = left != right
	boolSink = left < right
	boolSink = left <= right
	boolSink = left > right
	boolSink = left >= right
}

func fold64(condition bool) {
	left, right := float64(7.5), float64(2.5)
	if condition {
		left, right = 7.5, 2.5
	}
	float64Sink = left + right
	float64Sink = left - right
	float64Sink = left * right
	float64Sink = left / right
	float64Sink = -left
	boolSink = left == right
	boolSink = left != right
	boolSink = left < right
	boolSink = left <= right
	boolSink = left > right
	boolSink = left >= right
}

func exceptional(condition bool) {
	negative, zero := float64(-1), float64(0)
	if condition {
		negative, zero = -1, 0
	}
	float64Sink = negative / zero
	float64Sink = negative * zero
	float64Sink = -zero
}

func target() int { return 0 }
