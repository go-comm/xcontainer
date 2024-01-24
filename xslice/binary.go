package xslice

func BinaryWhere(length int, condition func(int) bool, cb func(int) (exit bool)) int {
	p := BinarySearch(length, condition)
	var cnt = 0
	if p < length {
		for ; p < length && condition(p); p++ {
			cnt++
			if cb(p) {
				break
			}
		}
	}
	return cnt
}

func BinaryCount(length int, condition func(int) bool) int {
	return BinaryWhere(length, condition, func(i int) (exit bool) { return false })
}

func BinaryGroup(length int, cmp func(i int, j int) int, f func(indexes ...int)) {
	var rows [][]int
	if length > 0 {
		rows = append(rows, []int{0})
	}
	j := len(rows) - 1
	for i := 1; i < length; i++ {
		if cmp(rows[j][0], i) == 0 {
			rows[j] = append(rows[j], i)
		} else {
			rows = append(rows, nil)
			j = len(rows) - 1
			rows[j] = []int{i}
		}
	}

	for _, row := range rows {
		f(row...)
	}
}
