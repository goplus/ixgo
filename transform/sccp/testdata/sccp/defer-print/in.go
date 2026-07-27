package test

func target() int {
	defer println(42, true, false, true, 1.5, "world", (chan int)(nil), []int(nil), (map[string]int)(nil), (func())(nil), byte(255))
	defer println(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	defer print("printing: ")
	return 0
}
