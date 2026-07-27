package test

func target() int {
	pv := []int{3, 4, 5}
	if pv[1] != 9 {
		pv = append(pv, 9)
	}
	tryit := func() bool {
		lpv := len(pv)
		if lpv == 101 {
			return false
		}
		if worst := pv[pv[1]&1]; worst != 101 {
			return true
		}
		return false
	}()
	if tryit {
		return pv[0]
	}
	return 0
}
