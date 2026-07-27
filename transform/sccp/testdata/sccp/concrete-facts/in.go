package test

import "unsafe"

type record struct {
	Value int
}

type interfaceRecord struct {
	Value any
}

type namedRecordPointer *record

func consume(any) {}

func fieldValue(input *record) int {
	return input.Value
}

func recordResult() record {
	return record{}
}

func interfaceRecordResult() interfaceRecord {
	return interfaceRecord{}
}

func fieldDirect() int {
	return recordResult().Value
}

func interfaceField() any {
	return interfaceRecordResult().Value
}

func fieldPointer(input *record) *int {
	return &input.Value
}

func dereference(input *int) int {
	return *input
}

func changePointer(input *record) namedRecordPointer {
	return namedRecordPointer(input)
}

func convertPointer(input *record) unsafe.Pointer {
	return unsafe.Pointer(input)
}

func boxPointer(input *record) any {
	return input
}

func equalPointers(left, right *record) bool {
	return left == right
}

func notEqualPointers(left, right *record) bool {
	return left != right
}

func pointerIsNil(input *record) bool {
	return input == nil
}

func sliceIsNil(input []int) bool {
	return input == nil
}

func mapIsNil(input map[string]int) bool {
	return input == nil
}

func channelIsNil(input chan int) bool {
	return input == nil
}

func functionIsNil(input func()) bool {
	return input == nil
}

func equalAny(left, right any) bool {
	return left == right
}

func equalComplex(left, right complex128) bool {
	return left == right
}

func equalBool(left, right bool) bool     { return left == right }
func equalInt(left, right int) bool       { return left == right }
func equalUint(left, right uint) bool     { return left == right }
func equalFloat(left, right float64) bool { return left == right }
func equalString(left, right string) bool { return left == right }

func addComplex(left, right complex128) complex128 {
	return left + right
}

func identityBool(value bool) bool       { return value }
func identityInt8(value int8) int8       { return value }
func identityInt16(value int16) int16    { return value }
func identityInt32(value int32) int32    { return value }
func identityInt64(value int64) int64    { return value }
func identityInt(value int) int          { return value }
func identityUint8(value uint8) uint8    { return value }
func identityUint16(value uint16) uint16 { return value }
func identityUint32(value uint32) uint32 { return value }
func identityUint64(value uint64) uint64 { return value }
func identityUint(value uint) uint       { return value }
func identityUintptr(value uintptr) uintptr {
	return value
}
func identityFloat32(value float32) float32 { return value }
func identityFloat64(value float64) float64 { return value }
func identityString(value string) string    { return value }
func identityAny(value any) any             { return value }

func mutatingCall(input *record) int {
	consume(input)
	return input.Value
}

func derivedPaths(input *record, array *[1]int, condition bool) int {
	consume(&input.Value)
	consume(&array[0])
	consume(namedRecordPointer(input))
	consume(unsafe.Pointer(input))
	consume(any(input))
	var selected *record
	if condition {
		selected = input
	}
	consume(selected)
	return input.Value
}

func target(input *record) int {
	return fieldValue(input)
}
