package xslice

import "sort"

var _ = sort.Search

func BinarySearch(length int, f func(int) bool) int {
	return sort.Search(length, f)
}

func BinarySearchN(length int, i int, j int, f func(int) bool) int {
	return i + BinarySearch(j-i, func(k int) bool { return f(i + k) })
}

func BinarySearchIndexInts(arr []int, x int) int {
	i := BinarySearch(len(arr), func(i int) bool { return arr[i] >= x })
	if i < len(arr) && arr[i] == x {
		return i
	}
	return -1
}

func BinarySearchIndexInt32s(arr []int32, x int32) int {
	i := BinarySearch(len(arr), func(i int) bool { return arr[i] >= x })
	if i < len(arr) && arr[i] == x {
		return i
	}
	return -1
}

func BinarySearchIndexInt64s(arr []int64, x int64) int {
	i := BinarySearch(len(arr), func(i int) bool { return arr[i] >= x })
	if i < len(arr) && arr[i] == x {
		return i
	}
	return -1
}

func BinarySearchIndexStrings(arr []string, x string) int {
	i := BinarySearch(len(arr), func(i int) bool { return arr[i] >= x })
	if i < len(arr) && arr[i] == x {
		return i
	}
	return -1
}
