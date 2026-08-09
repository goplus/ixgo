package pkg

import "fmt"

func Add(i, j int) int {
	return i + j
}

func Println(a ...interface{}) {
	fmt.Println(a...)
}
