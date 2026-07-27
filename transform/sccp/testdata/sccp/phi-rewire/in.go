package test

func target(input int) int {
	condition := true
	if input > 0 {
		condition = true
	}
	value := 0
	if condition {
		value = 1
	}
	return value
}

func parameter(input int, dynamic bool) int {
	condition := true
	if dynamic {
		condition = true
	}
	value := 0
	if condition {
		value = input
	}
	return value
}
