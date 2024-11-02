package xslice

func Every(length int, index func(i int) bool) bool {
	if length <= 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if !index(i) {
			return false
		}
	}
	return true
}

func EveryInts(arr []int, x int) bool {
	return Every(len(arr), func(i int) bool { return arr[i] == x })
}

func EveryInt32s(arr []int32, x int32) bool {
	return Every(len(arr), func(i int) bool { return arr[i] == x })
}

func EveryInt64s(arr []int64, x int64) bool {
	return Every(len(arr), func(i int) bool { return arr[i] == x })
}

func EveryStrings(arr []string, x string) bool {
	return Every(len(arr), func(i int) bool { return arr[i] == x })
}
