package reflect

import "reflect"

func TypeAssert[T any](v reflect.Value) (T, bool) {
	v2, ok := v.Interface().(T)
	return v2, ok
}
