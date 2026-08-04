package test

func interfaces() bool {
	var (
		nilN interface{}
		five = 5
	)
	ii := func(i1 interface{}, i2 interface{}) bool { return i1 == i2 }
	ni := func(n interface{}, i int) bool { return n == i }
	in := func(i int, n interface{}) bool { return i == n }

	return (interface{}(nil) == interface{}(nil)) == ii(nilN, nilN) &&
		(5 == interface{}(5)) == ni(five, five) &&
		(interface{}(5) == 5) == in(five, five)
}

func assertions() bool {
	value := interface{}(5)
	got, ok := value.(int)
	return ok && got == 5
}

func target() int {
	if interfaces() && assertions() {
		return 0
	}
	return 1
}
