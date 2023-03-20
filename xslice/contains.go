package xslice

func Contains(length int, index func(i int) bool) bool {
	return Find(length, index) >= 0
}

func ContainsInts(arr []int, x int) bool {
	return Contains(len(arr), func(i int) bool { return arr[i] == x })
}

func ContainsInt32s(arr []int32, x int32) bool {
	return Contains(len(arr), func(i int) bool { return arr[i] == x })
}

func ContainsInt64s(arr []int64, x int64) bool {
	return Contains(len(arr), func(i int) bool { return arr[i] == x })
}

func ContainsStrings(arr []string, x string) bool {
	return Contains(len(arr), func(i int) bool { return arr[i] == x })
}
