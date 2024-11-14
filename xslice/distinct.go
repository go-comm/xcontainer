package xslice

import "time"

// DistinctStable
func DistinctStable(length int, equal func(i int, j int) bool, push func(i int)) int {
	return FilterStable(length, func(i int) bool {
		return Find(i, func(j int) bool { return equal(i, j) }) < 0
	}, push)
}

// DistinctUnstable slice changed
func DistinctUnstable(length int, swap func(i int, j int), equal func(i int, j int) bool) int {
	return FilterUnstable(length, func(i int) bool {
		return Find(i, func(j int) bool { return equal(i, j) }) < 0
	}, swap)
}

// DistinctSorted slice changed
func DistinctSorted(length int, swap func(i int, j int), cmp func(i int, j int) int) int {
	if length <= 1 {
		return length
	}
	i := 1
	j := 0
	for ; i < length; i++ {
		p, found := BinaryFind(j+1, func(k int) int { return cmp(i, k) })
		if !found {
			j++
			if i != j {
				swap(i, j)
			}
			for k := j; k > p; k-- {
				swap(k, k-1)
			}
		}
	}
	return j + 1
}

func DistinctInts(arr []int) []int {
	var dst []int
	DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctInt32s(arr []int32) []int32 {
	var dst []int32
	DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctInt64s(arr []int64) []int64 {
	var dst []int64
	DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctUints(arr []uint) []uint {
	var dst []uint
	DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctUint32s(arr []uint32) []uint32 {
	var dst []uint32
	DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctUint64s(arr []uint64) []uint64 {
	var dst []uint64
	DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctFloat32s(arr []float32) []float32 {
	var dst []float32
	DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctFloat64s(arr []float64) []float64 {
	var dst []float64
	DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctStrings(arr []string) []string {
	var dst []string
	DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctTimes(arr []time.Time) []time.Time {
	var dst []time.Time
	DistinctStable(len(arr), func(i, j int) bool { return CompareTime(arr[i], arr[j]) == 0 }, func(i int) { dst = append(dst, arr[i]) })
	return dst
}

func DistinctUnstableInts(arr []int) []int {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctUnstableInt32s(arr []int32) []int32 {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctUnstableInt64s(arr []int64) []int64 {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctUnstableUints(arr []uint) []uint {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctUnstableUint32s(arr []uint32) []uint32 {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctUnstableUint64s(arr []uint64) []uint64 {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctUnstableFloat32s(arr []float32) []float32 {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctUnstableFloat64s(arr []float64) []float64 {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctUnstableStrings(arr []string) []string {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctUnstableTimes(arr []time.Time) []time.Time {
	p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return CompareTime(arr[i], arr[j]) == 0 })
	return arr[:p]
}

func DistinctSortInts(arr []int) []int {
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareInt(arr[i], arr[j]) })
	return arr[:p]
}

func DistinctSortInt32s(arr []int32) []int32 {
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareInt32(arr[i], arr[j]) })
	return arr[:p]
}

func DistinctSortInt64s(arr []int64) []int64 {
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareInt64(arr[i], arr[j]) })
	return arr[:p]
}

func DistinctSortUints(arr []uint) []uint {
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareUint(arr[i], arr[j]) })
	return arr[:p]
}

func DistinctSortUint32s(arr []uint32) []uint32 {
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareUint32(arr[i], arr[j]) })
	return arr[:p]
}

func DistinctSortUint64s(arr []uint64) []uint64 {
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareUint64(arr[i], arr[j]) })
	return arr[:p]
}

func DistinctSortFloat32s(arr []float32) []float32 {
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareFloat32(arr[i], arr[j]) })
	return arr[:p]
}

func DistinctSortFloat64s(arr []float64) []float64 {
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareFloat64(arr[i], arr[j]) })
	return arr[:p]
}

func DistinctSortStrings(arr []string) []string {
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareString(arr[i], arr[j]) })
	return arr[:p]
}

func IsDistinct(length int, equal func(i int, j int) bool) (int, bool) {
	p := Find(length, func(i int) bool {
		return Find(i, func(j int) bool { return equal(i, j) }) >= 0
	})
	return p, p < 0
}

func IsDistinctInts(arr []int) (int, bool) {
	return IsDistinct(len(arr), func(i, j int) bool { return arr[i] == arr[j] })
}

func IsDistinctInt32s(arr []int32) (int, bool) {
	return IsDistinct(len(arr), func(i, j int) bool { return arr[i] == arr[j] })
}

func IsDistinctInt64s(arr []int64) (int, bool) {
	return IsDistinct(len(arr), func(i, j int) bool { return arr[i] == arr[j] })
}

func IsDistinctUints(arr []uint) (int, bool) {
	return IsDistinct(len(arr), func(i, j int) bool { return arr[i] == arr[j] })
}

func IsDistinctUint32s(arr []uint32) (int, bool) {
	return IsDistinct(len(arr), func(i, j int) bool { return arr[i] == arr[j] })
}

func IsDistinctUint64s(arr []uint64) (int, bool) {
	return IsDistinct(len(arr), func(i, j int) bool { return arr[i] == arr[j] })
}

func IsDistinctFloat32s(arr []float32) (int, bool) {
	return IsDistinct(len(arr), func(i, j int) bool { return arr[i] == arr[j] })
}

func IsDistinctFloat64s(arr []float64) (int, bool) {
	return IsDistinct(len(arr), func(i, j int) bool { return arr[i] == arr[j] })
}

func IsDistinctStrings(arr []string) (int, bool) {
	return IsDistinct(len(arr), func(i, j int) bool { return arr[i] == arr[j] })
}
