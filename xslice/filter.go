package xslice

import "time"

func FilterStable(length int, f func(i int) bool, push func(i int)) int {
	size := 0
	for i := 0; i < length; i++ {
		if f(i) {
			push(i)
			size++
		}
	}
	return size
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

func FilterUints(arr []uint, f func(i int) bool) []uint {
	var dst []uint
	Filter(len(arr), f, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func FilterUint32s(arr []uint32, f func(i int) bool) []uint32 {
	var dst []uint32
	Filter(len(arr), f, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func FilterUint64s(arr []uint64, f func(i int) bool) []uint64 {
	var dst []uint64
	Filter(len(arr), f, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func FilterStrings(arr []string, f func(i int) bool) []string {
	var dst []string
	Filter(len(arr), f, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func FilterTimes(arr []time.Time, f func(i int) bool) []time.Time {
	var dst []time.Time
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

func FilterUnstableUints(arr []uint, f func(i int) bool) []uint {
	p := FilterUnstable(len(arr), f, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func FilterUnstableUint32s(arr []uint32, f func(i int) bool) []uint32 {
	p := FilterUnstable(len(arr), f, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func FilterUnstableUint64s(arr []uint64, f func(i int) bool) []uint64 {
	p := FilterUnstable(len(arr), f, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func FilterUnstableStrings(arr []string, f func(i int) bool) []string {
	p := FilterUnstable(len(arr), f, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func FilterUnstableTimes(arr []time.Time, f func(i int) bool) []time.Time {
	p := FilterUnstable(len(arr), f, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}
