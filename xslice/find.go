package xslice

func Find(length int, index func(i int) bool) int {
	for i := 0; i < length; i++ {
		if index(i) {
			return i
		}
	}
	return -1
}

func FindN(length int, i int, j int, index func(i int) bool) int {
	for k := i; k < j; k++ {
		if index(k) {
			return k
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

func LastFind(length int, index func(i int) bool) int {
	for i := length - 1; i >= 0; i-- {
		if index(i) {
			return i
		}
	}
	return -1
}

func LastFindInts(arr []int, x int) int {
	return LastFind(len(arr), func(i int) bool { return arr[i] == x })
}

func LastFindInt32s(arr []int32, x int32) int {
	return LastFind(len(arr), func(i int) bool { return arr[i] == x })
}

func LastFindInt64s(arr []int64, x int64) int {
	return LastFind(len(arr), func(i int) bool { return arr[i] == x })
}

func LastFindStrings(arr []string, x string) int {
	return LastFind(len(arr), func(i int) bool { return arr[i] == x })
}
