package xslice

func Find(length int, index func(i int) bool) int {
	for i := 0; i < length; i++ {
		if index(i) {
			return i
		}
	}
	return -1
}

func FindInts(arr []int, x int) int {
	return Find(len(arr), func(i int) bool { return arr[i] == x })
}

func FindInt32s(arr []int32, x int32) int {
	return Find(len(arr), func(i int) bool { return arr[i] == x })
}

func FindInt64s(arr []int64, x int64) int {
	return Find(len(arr), func(i int) bool { return arr[i] == x })
}

func FindStrings(arr []string, x string) int {
	return Find(len(arr), func(i int) bool { return arr[i] == x })
}
