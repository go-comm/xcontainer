package xslice

func Group(length int, equal func(i int, j int) bool, f func(row ...int)) {
	var rows [][]int
	if length > 0 {
		rows = append(rows, []int{0})
	}
	var found int
	for i := 1; i < length; i++ {
		found = -1
		for j := len(rows) - 1; j >= 0; j-- {
			if len(rows[j]) > 0 && equal(i, rows[j][0]) {
				found = j
				break
			}
		}
		if found >= 0 {
			rows[found] = append(rows[found], i)
		} else {
			rows = append(rows, []int{i})
		}
	}

	for _, row := range rows {
		f(row...)
	}
}

func GroupInts(arr []int, f func(row ...int)) {
	Group(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, f)
}

func GroupInt32s(arr []int32, f func(row ...int)) {
	Group(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, f)
}

func GroupInt64s(arr []int64, f func(row ...int)) {
	Group(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, f)
}

func GroupStrings(arr []string, f func(row ...int)) {
	Group(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, f)
}
