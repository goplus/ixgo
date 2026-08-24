package json

import (
	qt "encoding/json/jsontext"
	q "encoding/json/v2"
)

func GetOption[T any](opts q.Options, setter func(T) q.Options) (T, bool) {
	var zero T
	if _, ok := any(zero).(bool); ok {
		v, found := getOptionBool(opts, func(v bool) q.Options {
			return setter(any(v).(T))
		})
		return any(v).(T), found
	}
	v, ok := getOptionAny(opts, func(v any) q.Options {
		return setter(v.(T))
	})
	if !ok || v == nil {
		return zero, false
	}
	return v.(T), true
}

func MarshalFunc[T any](fn func(T) ([]byte, error)) *q.Marshalers {
	return marshalFuncAny(func(v any) ([]byte, error) {
		return fn(v.(T))
	}, (*T)(nil))
}

func MarshalToFunc[T any](fn func(*qt.Encoder, T) error) *q.Marshalers {
	return marshalToFuncAny(func(enc *qt.Encoder, v any) error {
		return fn(enc, v.(T))
	}, (*T)(nil))
}

func UnmarshalFunc[T any](fn func([]byte, T) error) *q.Unmarshalers {
	return unmarshalFuncAny(func(data []byte, v any) error {
		return fn(data, v.(T))
	}, (*T)(nil))
}

func UnmarshalFromFunc[T any](fn func(*qt.Decoder, T) error) *q.Unmarshalers {
	return unmarshalFromFuncAny(func(dec *qt.Decoder, v any) error {
		return fn(dec, v.(T))
	}, (*T)(nil))
}

func getOptionAny(opts q.Options, setter func(any) q.Options) (any, bool)
func getOptionBool(opts q.Options, setter func(bool) q.Options) (bool, bool)
func marshalFuncAny(fn func(any) ([]byte, error), typ any) *q.Marshalers
func marshalToFuncAny(fn func(*qt.Encoder, any) error, typ any) *q.Marshalers
func unmarshalFuncAny(fn func([]byte, any) error, typ any) *q.Unmarshalers
func unmarshalFromFuncAny(fn func(*qt.Decoder, any) error, typ any) *q.Unmarshalers
