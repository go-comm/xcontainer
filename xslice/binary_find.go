package xslice

import (
	"strings"
)

func BinaryFind(n int, cmp func(int) int) (int, bool) {
	// The invariants here are similar to the ones in Search.
	// Define cmp(-1) > 0 and cmp(n) <= 0
	// Invariant: cmp(i-1) > 0, cmp(j) <= 0
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if cmp(h) > 0 {
			i = h + 1 // preserves cmp(i-1) > 0
		} else {
			j = h // preserves cmp(j) <= 0
		}
	}
	// i == j, cmp(i-1) > 0 and cmp(j) <= 0
	return i, i < n && cmp(i) == 0
}

func BinaryFindN(length int, i int, j int, cmp func(int) int) (int, bool) {
	p, ok := BinaryFind(j-i, func(k int) int { return cmp(i + k) })
	return p + i, ok
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
