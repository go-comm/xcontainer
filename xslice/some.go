package xslice

func Some(length int, index func(i int) bool) bool {
	return Find(length, index) >= 0
}

func SomeInts(arr []int, x int) bool {
	return Find(len(arr), func(i int) bool { return arr[i] == x }) >= 0
}

func SomeInt32s(arr []int32, x int32) bool {
	return Find(len(arr), func(i int) bool { return arr[i] == x }) >= 0
}

func SomeInt64s(arr []int64, x int64) bool {
	return Find(len(arr), func(i int) bool { return arr[i] == x }) >= 0
}

func SomeStrings(arr []string, x string) bool {
	return Find(len(arr), func(i int) bool { return arr[i] == x }) >= 0
}
