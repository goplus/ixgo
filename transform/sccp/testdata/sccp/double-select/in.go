package test

func target(n int, c1, c2, c3, c4 chan<- int) {
	defer close(c1)
	defer close(c2)
	defer close(c3)
	defer close(c4)

	for i := 0; i < n; i++ {
		select {
		case c1 <- i:
		case c2 <- i:
		case c3 <- i:
		case c4 <- i:
		}
	}
}
