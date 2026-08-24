//go:build go1.25

package main

import "reflect"

func main() {
	if reflect.TypeFor[int]().Kind() != reflect.Int {
		panic("TypeFor")
	}
	if got, ok := reflect.TypeAssert[int](reflect.ValueOf(42)); !ok || got != 42 {
		panic("TypeAssert success")
	}
	if _, ok := reflect.TypeAssert[string](reflect.ValueOf(42)); ok {
		panic("TypeAssert failure")
	}
}
