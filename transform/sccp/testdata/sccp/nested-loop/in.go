package test

func target() bool {
	var n, a, b int64
	for i := int64(2); i < 10; i++ {
		for j := i; j < 10; j++ {
			if n%(i*j) == 0 && j > 1 && n/(i*j) == 1 {
				a, b = i, 0
				a = n / (i * j)
			}
		}
	}
	return a != b && a != n
}
