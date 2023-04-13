package xslice

func IsAllSame(length int, fn func(i int, j int) bool) bool {
	if length <= 0 {
		return false
	}
	for i := 1; i < length; i++ {
		if !fn(i, 0) {
			return false
		}
	}
	return true
}

func IsAllSameInts(arr []int) bool {
	return IsAllSame(len(arr), func(i int, j int) bool { return arr[i] == arr[j] })
}

func IsAllSameInt32s(arr []int32) bool {
	return IsAllSame(len(arr), func(i int, j int) bool { return arr[i] == arr[j] })
}

func IsAllSameInt64s(arr []int64) bool {
	return IsAllSame(len(arr), func(i int, j int) bool { return arr[i] == arr[j] })
}

func IsAllSameStrings(arr []string) bool {
	return IsAllSame(len(arr), func(i int, j int) bool { return arr[i] == arr[j] })
}
