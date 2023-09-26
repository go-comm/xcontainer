package xslice

func makeSortIndex(indexes []int, less func(i int, j int) bool) []int {
	for i := 0; i < len(indexes); i++ {
		indexes[i] = i
	}
	Sort(len(indexes), func(i, j int) { indexes[i], indexes[j] = indexes[j], indexes[i] }, func(i, j int) bool { return less(indexes[i], indexes[j]) })
	return indexes
}

func appendSortIndex(indexes []int, f func(i int) bool) []int {
	n := len(indexes)
	p := BinarySearch(len(indexes), func(i int) bool { return f(indexes[i]) })
	indexes = append(indexes, 0)
	copy(indexes[p+1:], indexes[p:])
	indexes[p] = n
	return indexes
}

func searchSortIndex(indexes []int, f func(int) bool) int {
	return BinarySearch(len(indexes), func(i int) bool { return f(indexes[i]) })
}

func findSortIndex(indexes []int, cmp func(int) int) (int, bool) {
	return BinaryFind(len(indexes), func(i int) int { return cmp(indexes[i]) })
}
