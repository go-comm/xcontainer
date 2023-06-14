package xslice

func Min(length int, less func(i int, j int) bool) int {
	if length == 0 {
		return -1
	}
	if length == 1 {
		return 0
	}
	var k int = 0
	for i := k + 1; i < length; i++ {
		if less(i, k) {
			k = i
		}
	}
	return k
}

func MinInts(arr []int, def int) int {
	p := Min(len(arr), func(i, j int) bool { return arr[i] < arr[j] })
	if p < 0 {
		return def
	}
	return arr[p]
}

func MinInt32s(arr []int32, def int32) int32 {
	p := Min(len(arr), func(i, j int) bool { return arr[i] < arr[j] })
	if p < 0 {
		return def
	}
	return arr[p]
}

func MinInt64s(arr []int64, def int64) int64 {
	p := Min(len(arr), func(i, j int) bool { return arr[i] < arr[j] })
	if p < 0 {
		return def
	}
	return arr[p]
}

func MinStrings(arr []string, def string) string {
	p := Min(len(arr), func(i, j int) bool { return arr[i] < arr[j] })
	if p < 0 {
		return def
	}
	return arr[p]
}

func MinToInt(length int, index func(i int) int, def int) int {
	p := Min(length, func(i, j int) bool { return index(i) < index(j) })
	if p < 0 {
		return def
	}
	return index(p)
}

func MinToInt32(length int, index func(i int) int32, def int32) int32 {
	p := Min(length, func(i, j int) bool { return index(i) < index(j) })
	if p < 0 {
		return def
	}
	return index(p)
}

func MinToInt64(length int, index func(i int) int64, def int64) int64 {
	p := Min(length, func(i, j int) bool { return index(i) < index(j) })
	if p < 0 {
		return def
	}
	return index(p)
}
