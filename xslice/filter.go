package xslice

func FilterStable(length int, f func(i int) bool, push func(i int)) {
	for i := 0; i < length; i++ {
		if f(i) {
			push(i)
		}
	}
}

func FilterUnstable(length int, f func(i int) bool, swap func(i int, j int)) int {
	var j = 0
	for i := 0; i < length; i++ {
		if f(i) {
			if i != j {
				swap(i, j)
			}
			j++
		}
	}
	return j
}

func Filter(length int, f func(i int) bool, push func(i int)) {
	FilterStable(length, f, push)
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

func FilterUnstableInts(arr []int, f func(i int) bool) []int {
	p := FilterUnstable(len(arr), f, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func FilterUnstableInt32s(arr []int32, f func(i int) bool) []int32 {
	p := FilterUnstable(len(arr), f, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func FilterUnstableInt64s(arr []int64, f func(i int) bool) []int64 {
	p := FilterUnstable(len(arr), f, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func FilterUnstableStrings(arr []string, f func(i int) bool) []string {
	p := FilterUnstable(len(arr), f, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}
