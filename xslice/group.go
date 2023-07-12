package xslice

func Group(length int, cmp func(i int, j int) int, f func(indexes ...int)) {
	var rows [][]int
	if length > 0 {
		rows = append(rows, []int{0})
	}
	for i := 1; i < length; i++ {
		p, found := BinaryFind(len(rows), func(j int) int { return cmp(i, rows[j][0]) })
		if found {
			rows[p] = append(rows[p], i)
		} else {
			rows = append(rows, nil)
			copy(rows[p+1:], rows[p:])
			rows[p] = []int{i}
		}
	}

	for _, row := range rows {
		f(row...)
	}
}

func GroupInts(arr []int, f func(indexes ...int)) {
	Group(len(arr), func(i, j int) int { return arr[i] - arr[j] }, f)
}

func GroupInt32s(arr []int32, f func(indexes ...int)) {
	Group(len(arr), func(i, j int) int { return CompareInt32(arr[i], arr[j]) }, f)
}

func GroupInt64s(arr []int64, f func(indexes ...int)) {
	Group(len(arr), func(i, j int) int { return CompareInt64(arr[i], arr[j]) }, f)
}

func GroupStrings(arr []string, f func(indexes ...int)) {
	Group(len(arr), func(i, j int) int { return CompareString(arr[i], arr[j]) }, f)
}
