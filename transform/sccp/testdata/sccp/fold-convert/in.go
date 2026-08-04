package test

var int8Sink int8
var int16Sink int16
var int32Sink int32
var uint8Sink uint8
var float32Sink float32
var float64Sink float64
var stringSink string

func conversions(condition bool) {
	signed, unsigned := int16(-7), uint16(255)
	float32Value, float64Value := float32(1.25), float64(2.9)
	if condition {
		signed, unsigned = -7, 255
		float32Value, float64Value = 1.25, 2.9
	}
	int8Sink = int8(signed)
	uint8Sink = uint8(signed)
	float32Sink = float32(signed)
	float64Sink = float64(signed)
	int16Sink = int16(unsigned)
	uint8Sink = uint8(unsigned)
	float32Sink = float32(unsigned)
	float64Sink = float64(unsigned)
	float64Sink = float64(float32Value)
	int32Sink = int32(float64Value)
	uint8Sink = uint8(float64Value)
	float32Sink = float32(float64Value)
	stringSink = string(signed)
}

func rejected(condition bool) {
	large, negative := float64(1e30), float64(-1)
	if condition {
		large, negative = 1e30, -1
	}
	int32Sink = int32(large)
	uint8Sink = uint8(negative)
}

func dynamic(value int16) {
	float64Sink = float64(value)
}

func target() int { return 0 }
