package xslice

import (
	"sort"
	"strings"
)

var _ = sort.Find

func BinaryFind(length int, cmp func(int) int) (int, bool) {
	return sort.Find(length, cmp)
}

func BinaryFindIndexInts(arr []int, x int) (int, bool) {
	return BinaryFind(len(arr), func(i int) int { return x - arr[i] })
}

func BinaryFindIndexInt32s(arr []int32, x int32) (int, bool) {
	return BinaryFind(len(arr), func(i int) int { return CompareInt32(x, arr[i]) })
}

func BinaryFindIndexInt64s(arr []int64, x int64) (int, bool) {
	return BinaryFind(len(arr), func(i int) int { return CompareInt64(x, arr[i]) })
}

func BinaryFindIndexStrings(arr []string, x string) (int, bool) {
	return BinaryFind(len(arr), func(i int) int { return strings.Compare(x, arr[i]) })
}
