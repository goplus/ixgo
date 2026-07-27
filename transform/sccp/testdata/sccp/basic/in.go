package test

func target(condition bool) (int64, int64, int64, int64, int64, int64, int64, int64, int64, int64, int64, int64, int32, int64, bool, bool, bool, bool, bool, int64) {
	v1 := int64(20)
	v2 := int64(21)
	v3 := float64(21)
	v4 := true
	if condition {
		v1 = 20
		v2 = 21
		v3 = 21
		v4 = true
	}
	t1 := v1 + v2
	t2 := t1 / v1
	t3 := t1 + t2
	t4 := t3 - v2
	t5 := t4 * v2
	t6 := t5 % v2
	t7 := t6 & v2
	t8 := t7 | v2
	t9 := t8 ^ v2
	t10 := -t9
	t11 := ^t10
	t12 := -t11
	t18 := int32(t12)
	t19 := int64(v3)
	t23 := !v4
	t24 := v1 == v2
	t25 := v1 < v2
	t26 := v1 <= v2
	t27 := v4 == v4
	t28 := v2 << uint(v1)
	return t1, t2, t3, t4, t5, t6, t7, t8, t9, t10, t11, t12, t18, t19, t23, t24, t25, t26, t27, t28
}
