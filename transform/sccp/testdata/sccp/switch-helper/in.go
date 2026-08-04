package test

func target(x int) int {
	y := 0
	switch {
	case x > 0:
		y += 5
		return 0 + y
	case x < 1:
		y += 6
		fallthrough
	default:
		y += 7
		return 2 + y
	}
}
