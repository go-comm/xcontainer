package xslice

import "sort"

var _ = sort.Sort

func SortInterface(length int, swap func(i int, j int), less func(i int, j int) bool) sort.Interface {
	return &wrappedSorter{length: length, swap: swap, less: less}
}

func Sort(length int, swap func(i int, j int), less func(i int, j int) bool) {
	sort.Sort(SortInterface(length, swap, less))
}

func SortInts(arr []int) []int {
	Sort(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}

func SortInt32s(arr []int32) []int32 {
	Sort(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}

func SortInt64s(arr []int64) []int64 {
	Sort(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}

func SortStrings(arr []string) []string {
	Sort(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}

func insertSort(a int, b int, swap func(i int, j int), less func(i int, j int) bool) {
	for i := a + 1; i < b; i++ {
		for j := i; j > a && less(j, j-1); j-- {
			swap(j, j-1)
		}
	}
}

func siftDown(lo, hi, first int, swap func(i int, j int), less func(i int, j int) bool) {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(first+child, first+child+1) {
			child++
		}
		if !less(first+root, first+child) {
			return
		}
		swap(first+root, first+child)
		root = child
	}
}

func heapSort(a, b int, swap func(i int, j int), less func(i int, j int) bool) {
	first := a
	lo := 0
	hi := b - a
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(i, hi, first, swap, less)
	}
	for i := hi - 1; i >= 0; i-- {
		swap(first, first+i)
		siftDown(lo, i, first, swap, less)
	}
}

func IsSorted(length int, less func(i int, j int) bool) bool {
	for i := length - 1; i > 0; i-- {
		if less(i, i-1) {
			return false
		}
	}
	return true
}

func IsSortedInts(arr []int) bool {
	return IsSorted(len(arr), func(i, j int) bool { return arr[i] < arr[j] }) || IsSorted(len(arr), func(i, j int) bool { return arr[i] > arr[j] })
}

func IsSortedInt32s(arr []int32) bool {
	return IsSorted(len(arr), func(i, j int) bool { return arr[i] < arr[j] }) || IsSorted(len(arr), func(i, j int) bool { return arr[i] > arr[j] })
}

func IsSortedInt64s(arr []int64) bool {
	return IsSorted(len(arr), func(i, j int) bool { return arr[i] < arr[j] }) || IsSorted(len(arr), func(i, j int) bool { return arr[i] > arr[j] })
}

func IsSortedStrings(arr []string) bool {
	return IsSorted(len(arr), func(i, j int) bool { return arr[i] < arr[j] }) || IsSorted(len(arr), func(i, j int) bool { return arr[i] > arr[j] })
}

var _ sort.Interface = (*wrappedSorter)(nil)

type wrappedSorter struct {
	length int
	swap   func(i int, j int)
	less   func(i int, j int) bool
}

func (s *wrappedSorter) Len() int {
	return s.length
}

func (s *wrappedSorter) Less(i, j int) bool {
	return s.less(i, j)
}

func (s *wrappedSorter) Swap(i int, j int) {
	s.swap(i, j)
}
