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

func MaxToInt(length int, index func(i int) int, def int) int {
	p := Min(length, func(i, j int) bool { return index(j) < index(i) })
	if p < 0 {
		return def
	}
	return index(p)
}

func MaxToInt32(length int, index func(i int) int32, def int32) int32 {
	p := Min(length, func(i, j int) bool { return index(j) < index(i) })
	if p < 0 {
		return def
	}
	return index(p)
}

func MaxToInt64(length int, index func(i int) int64, def int64) int64 {
	p := Min(length, func(i, j int) bool { return index(j) < index(i) })
	if p < 0 {
		return def
	}
	return index(p)
}
