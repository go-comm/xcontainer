package xslice

// ExceptStable returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptStable(lenA int, lenB int, equal func(ia int, ib int) bool, push func(ia int)) int {
	return FilterStable(lenA, func(i int) bool {
		return Find(lenB, func(j int) bool { return equal(i, j) }) < 0
	}, push)
}

// ExceptUnstable returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptUnstable(lenA int, lenB int, swapA func(i int, j int), equal func(ia int, ib int) bool) (ia int) {
	return FilterUnstable(lenA, func(i int) bool {
		return Find(lenB, func(j int) bool { return equal(i, j) }) < 0
	}, swapA)
}

// ExceptSorted returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptSorted(lenA int, lenSortedB int, swapA func(i int, j int), cmp func(ia int, ib int) int) int {
	hi := lenA
	if lenSortedB > 0 {
		for i := hi - 1; i >= 0; i-- {
			_, ok := BinaryFind(lenSortedB, func(j int) int { return cmp(i, j) })
			if ok {
				if i != hi-1 {
					swapA(i, hi-1)
				}
				hi--
			}
		}
	}
	return hi
}

// Except returns the elements in `a` that aren't in `b`. a-(a∩b)
func Except(lenA int, lenB int, equal func(ia int, ib int) bool, push func(ia int)) int {
	return ExceptStable(lenA, lenB, equal, push)
}

// ExceptInts returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptInts(a []int, b []int) []int {
	var c []int
	ExceptStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// ExceptInt32s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptInt32s(a []int32, b []int32) []int32 {
	var c []int32
	ExceptStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// ExceptInt64s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptInt64s(a []int64, b []int64) []int64 {
	var c []int64
	ExceptStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// ExceptStrings returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptStrings(a []string, b []string) []string {
	var c []string
	ExceptStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// ExceptUnstableInts returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptUnstableInts(a []int, b []int) []int {
	p := ExceptUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// ExceptUnstableInt32s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptUnstableInt32s(a []int32, b []int32) []int32 {
	p := ExceptUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// ExceptUnstableInt64s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptUnstableInt64s(a []int64, b []int64) []int64 {
	p := ExceptUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// ExceptUnstableStrings returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptUnstableStrings(a []string, b []string) []string {
	p := ExceptUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// ExceptSortInts returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptSortInts(a []int, b []int) []int {
	SortInts(b)
	p := ExceptSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareInt(a[ia], b[ib]) })
	return a[:p]
}

// ExceptSortInt32s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptSortInt32s(a []int32, b []int32) []int32 {
	SortInt32s(b)
	p := ExceptSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareInt32(a[ia], b[ib]) })
	return a[:p]
}

// ExceptSortInt64s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptSortInt64s(a []int64, b []int64) []int64 {
	SortInt64s(b)
	p := ExceptSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareInt64(a[ia], b[ib]) })
	return a[:p]
}

// ExceptSortStrings returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptSortStrings(a []string, b []string) []string {
	SortStrings(b)
	p := ExceptSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareString(a[ia], b[ib]) })
	return a[:p]
}
