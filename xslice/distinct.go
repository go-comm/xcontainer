package xslice

// DistinctUnsorted slice changed
func DistinctUnsorted(length int, swap func(i int, j int), equal func(i int, j int) bool) int {
	hi := length
	for i := length - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if equal(i, j) {
				if i != hi-1 {
					swap(i, hi-1)
				}
				hi--
				break
			}
		}
	}
	return hi
}

// DistinctSorted slice changed
func DistinctSorted(length int, swap func(i int, j int), equal func(i int, j int) bool) int {
	lo := 0
	if lo < length {
		hi := lo + 1
		for ; hi < length; hi++ {
			if !equal(lo, hi) {
				lo++
				swap(lo, hi)
			}
		}
		if hi >= length {
			lo++
		}
	}
	return lo
}

func DistinctInterfaces(arr []interface{}) int {
	return DistinctUnsorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
}

func DistinctInts(arr []int) []int {
	arr = SortInts(arr)
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctInt32s(arr []int32) []int32 {
	arr = SortInt32s(arr)
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctInt64s(arr []int64) []int64 {
	arr = SortInt64s(arr)
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}

func DistinctStrings(arr []string) []string {
	arr = SortStrings(arr)
	p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
	return arr[:p]
}
