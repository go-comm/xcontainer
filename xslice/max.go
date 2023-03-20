package xslice

func Max(length int, less func(i int, j int) bool) int {
	return Min(length, func(i, j int) bool { return less(j, i) })
}

func MaxInts(arr []int, def int) int {
	p := Min(len(arr), func(i, j int) bool { return arr[j] < arr[i] })
	if p < 0 {
		return def
	}
	return arr[p]
}

func MaxInt32s(arr []int32, def int32) int32 {
	p := Min(len(arr), func(i, j int) bool { return arr[j] < arr[i] })
	if p < 0 {
		return def
	}
	return arr[p]
}

func MaxInt64s(arr []int64, def int64) int64 {
	p := Min(len(arr), func(i, j int) bool { return arr[j] < arr[i] })
	if p < 0 {
		return def
	}
	return arr[p]
}

func MaxStrings(arr []string, def string) string {
	p := Min(len(arr), func(i, j int) bool { return arr[j] < arr[i] })
	if p < 0 {
		return def
	}
	return arr[p]
}
