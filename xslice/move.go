package xslice

// Move from i to j
func Move(length int, i int, j int, swap func(i int, j int)) {
	if i > j {
		for k := i; k > j; k-- {
			swap(k, k-1)
		}
	} else if i < j {
		for k := i; k < j; k++ {
			swap(k, k+1)
		}
	}
}

func MoveInts(arr []int, i int, j int) []int {
	Move(len(arr), i, j, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func MoveInt32s(arr []int32, i int, j int) []int32 {
	Move(len(arr), i, j, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func MoveInt64s(arr []int64, i int, j int) []int64 {
	Move(len(arr), i, j, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func MoveStrings(arr []string, i int, j int) []string {
	Move(len(arr), i, j, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}
