package host

// Counter is a host type used by direct-call interpreter tests.
type Counter int

// Value returns the current counter value.
func (c *Counter) Value() int {
	return int(*c)
}

// FallbackCounter is a host type without a direct-call adapter.
type FallbackCounter int

// Value returns the current fallback counter value.
func (c *FallbackCounter) Value() int {
	return int(*c)
}

// ValueCounter verifies pointer calls to a value-receiver method.
type ValueCounter int

// Value returns the current counter value.
func (c ValueCounter) Value() int {
	return int(c)
}

// Recorder is a host type used to verify Go and defer interface invocation.
type Recorder struct {
	Values chan int
}

// Record publishes value to Values.
func (r *Recorder) Record(value int) {
	r.Values <- value
}
