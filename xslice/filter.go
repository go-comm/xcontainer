package xslice

func Filter(length int, f func(i int) bool, push func(i int)) {
	for i := 0; i < length; i++ {
		if f(i) {
			push(i)
		}
	}
}

func FilterInts(arr []int, f func(i int) bool) []int {
	var dst []int
	Filter(len(arr), f, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func FilterInt32s(arr []int32, f func(i int) bool) []int32 {
	var dst []int32
	Filter(len(arr), f, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func FilterInt64s(arr []int64, f func(i int) bool) []int64 {
	var dst []int64
	Filter(len(arr), f, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func FilterStrings(arr []string, f func(i int) bool) []string {
	var dst []string
	Filter(len(arr), f, func(i int) { dst = append(dst, arr[i]) })
	return dst
}
