package xslice

import (
	"reflect"
	"sort"
)

var _ = sort.Sort

func WrappedSorter(length int, swap func(i int, j int), less func(i int, j int) bool) sort.Interface {
	return &wrappedSorter{length: length, swap: swap, less: less}
}

func SortSlice(slice interface{}, less func(i int, j int) bool) {
	rv := reflect.ValueOf(slice)
	sort.Sort(WrappedSorter(rv.Len(), Swapper(slice), less))
}

func Sort(length int, swap func(i int, j int), less func(i int, j int) bool) {
	sort.Sort(WrappedSorter(length, swap, less))
}

func SortN(length int, i int, j int, swap func(int, int), less func(int, int) bool) {
	Sort(j-i, func(a, b int) { swap(a+i, b+i) }, func(a, b int) bool { return less(a+i, b+i) })
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

func IsSortedN(length int, i int, j int, less func(i int, j int) bool) bool {
	return IsSorted(j-i, func(a, b int) bool { return less(a+i, b+i) })
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

func FixSort(length int, swap func(i int, j int), less func(i int, j int) bool, x int) {
	p := BinarySearchN(length, 0, x, func(j int) bool { return less(x, j) })
	if p >= x {
		p = BinarySearchN(length, x+1, length, func(j int) bool { return less(x, j) })
	}
	if p == x {
		return
	}
	if p < x {
		for k := x; k > p; k-- {
			swap(k-1, k)
		}
	} else {
		for k := x + 1; k < p; k++ {
			swap(k-1, k)
		}
	}
}

func FixSortN(length int, i int, j int, swap func(int, int), less func(int, int) bool, x int) {
	FixSort(j-i, func(a, b int) { swap(a+i, b+i) }, func(a, b int) bool { return less(a+i, b+i) }, x)
}

func FixSortInterface(data sort.Interface, i int) {
	FixSort(data.Len(), data.Swap, data.Less, i)
}

func FixSortSlice(slice interface{}, less func(i int, j int) bool, i int) {
	rv := reflect.ValueOf(slice)
	FixSortInterface(WrappedSorter(rv.Len(), Swapper(slice), less), i)
}
