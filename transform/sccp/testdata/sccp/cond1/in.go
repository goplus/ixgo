package test

func target(condition bool) int64 {
	v1 := int64(0)
	if condition {
		v1 = 0
	}
	value := int64(3)
	if v1 < 1 {
		value = 3
	} else {
		value = 4
	}
	return value
}
