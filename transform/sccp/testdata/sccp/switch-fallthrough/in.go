package test

func switches() bool {
	i5 := 5
	i7 := 7

	switch true {
	case i5 < 5:
		return false
	case i5 == 5:
	case i5 > 5:
		return false
	}

	switch i5 {
	case 0, 1, 2, 3, 4:
		return false
	case 5:
	case 6, 7, 8, 9:
		return false
	default:
		return false
	}

	count := 0
	switch i5 {
	case 5:
		count++
		fallthrough
	case 6:
		count++
		fallthrough
	case 7:
		count++
		fallthrough
	case 8:
		count++
		fallthrough
	case 9:
		count++
		fallthrough
	default:
		if i5 != count {
			return false
		}
	}

	fired := false
	switch i := i5 + 2; i {
	case i7:
		fired = true
	default:
		return false
	}

	switch {
	default:
		count++
		fallthrough
	case false:
		count++
	}
	return fired && count == 7
}

func target() int {
	if switches() {
		return 0
	}
	return 1
}
