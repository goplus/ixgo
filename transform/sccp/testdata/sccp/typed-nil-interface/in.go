package test

type ET struct{}

func (*ET) Error() string { return "err" }

func target() int {
	if (*ET)(nil) == error(nil) {
		return 1
	}
	if (*ET)(nil) != error(nil) {
		return 0
	}
	return 2
}
