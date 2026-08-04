package test

func sum(args ...int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

func sumC(args ...int) int { return func() int { return sum(args...) }() }

var sumD = func(args ...int) int { return sum(args...) }

var sumE = func() func(...int) int { return func(args ...int) int { return sum(args...) } }()

func sumF(args ...int) func() int { return func() int { return sum(args...) } }

func sumA(args []int) int {
	s := 0
	for _, v := range args {
		s += v
	}
	return s
}

func sumB(args []int) int { return sum(args...) }

func sum2(args ...int) int { return 2 * sum(args...) }

func sum3(args ...int) int { return 3 * sumA(args) }

func sum4(args ...int) int { return 4 * sumB(args) }

func target() int {
	if sum(1, 2, 3) != 6 {
		return 1
	}
	if sumC(4, 5, 6) != 15 || sumD(4, 5, 7) != 16 || sumE(4, 5, 8) != 17 {
		return 2
	}
	if sumF(4, 5, 9)() != 18 {
		return 3
	}
	if sum2(1, 2, 3) != 12 || sum3(1, 2, 3) != 18 || sum4(1, 2, 3) != 24 {
		return 4
	}
	return 0
}
