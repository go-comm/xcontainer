package xslice

// IntersectionStable returns a∩b, a must be unique
func IntersectionStable(lenA int, lenB int, equal func(ia int, ib int) bool, push func(ia int)) int {
	return FilterStable(lenA, func(i int) bool {
		return Find(lenB, func(j int) bool { return equal(i, j) }) >= 0
	}, push)
}

// IntersectionUnstable returns a∩b, a must be unique
func IntersectionUnstable(lenA int, lenB int, swapA func(i int, j int), equal func(ia int, ib int) bool) (ia int) {
	return FilterUnstable(lenA, func(i int) bool {
		return Find(lenB, func(j int) bool { return equal(i, j) }) >= 0
	}, swapA)
}

// IntersectionSort returns a∩b, a must be unique
func IntersectionSorted(lenA int, lenSortedB int, swapA func(i int, j int), cmp func(ia int, ib int) int) (ia int) {
	return FilterUnstable(lenA, func(i int) bool {
		_, ok := BinaryFind(lenSortedB, func(j int) int { return cmp(i, j) })
		return ok
	}, swapA)
}

// Intersection returns a∩b, a must be unique
func Intersection(lenA int, lenB int, equal func(ia int, ib int) bool, push func(ia int)) int {
	return IntersectionStable(lenA, lenB, equal, push)
}

// IntersectionInts returns a∩b, a must be unique
func IntersectionInts(a []int, b []int) []int {
	var c []int
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionInt32s returns a∩b, a must be unique
func IntersectionInt32s(a []int32, b []int32) []int32 {
	var c []int32
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionInt64s returns a∩b, a must be unique
func IntersectionInt64s(a []int64, b []int64) []int64 {
	var c []int64
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionUints returns a∩b, a must be unique
func IntersectionUints(a []uint, b []uint) []uint {
	var c []uint
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionUint32s returns a∩b, a must be unique
func IntersectionUint32s(a []uint32, b []uint32) []uint32 {
	var c []uint32
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionUint64s returns a∩b, a must be unique
func IntersectionUint64s(a []uint64, b []uint64) []uint64 {
	var c []uint64
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionStrings returns a∩b, a must be unique
func IntersectionStrings(a []string, b []string) []string {
	var c []string
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionUnstableInts returns a∩b, a must be unique
func IntersectionUnstableInts(a []int, b []int) []int {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionUnstableInt32s returns a∩b, a must be unique
func IntersectionUnstableInt32s(a []int32, b []int32) []int32 {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionUnstableInt64s returns a∩b, a must be unique
func IntersectionUnstableInt64s(a []int64, b []int64) []int64 {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionUnstableUints returns a∩b, a must be unique
func IntersectionUnstableUints(a []uint, b []uint) []uint {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionUnstableUint32s returns a∩b, a must be unique
func IntersectionUnstableUint32s(a []uint32, b []uint32) []uint32 {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionUnstableUint64s returns a∩b, a must be unique
func IntersectionUnstableUint64s(a []uint64, b []uint64) []uint64 {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionUnstableStrings returns a∩b, a must be unique
func IntersectionUnstableStrings(a []string, b []string) []string {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionSortInts returns a∩b, a must be unique
func IntersectionSortInts(a []int, b []int) []int {
	SortInts(b)
	p := IntersectionSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareInt(a[ia], b[ib]) })
	return a[:p]
}

// IntersectionSortInt32s returns a∩b, a must be unique
func IntersectionSortInt32s(a []int32, b []int32) []int32 {
	SortInt32s(b)
	p := IntersectionSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareInt32(a[ia], b[ib]) })
	return a[:p]
}

// IntersectionSortInt64s returns a∩b, a must be unique
func IntersectionSortInt64s(a []int64, b []int64) []int64 {
	SortInt64s(b)
	p := IntersectionSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareInt64(a[ia], b[ib]) })
	return a[:p]
}

// IntersectionSortUints returns a∩b, a must be unique
func IntersectionSortUints(a []uint, b []uint) []uint {
	SortUints(b)
	p := IntersectionSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareUint(a[ia], b[ib]) })
	return a[:p]
}

// IntersectionSortUint32s returns a∩b, a must be unique
func IntersectionSortUint32s(a []uint32, b []uint32) []uint32 {
	SortUint32s(b)
	p := IntersectionSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareUint32(a[ia], b[ib]) })
	return a[:p]
}

// IntersectionSortUint64s returns a∩b, a must be unique
func IntersectionSortUint64s(a []uint64, b []uint64) []uint64 {
	SortUint64s(b)
	p := IntersectionSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareUint64(a[ia], b[ib]) })
	return a[:p]
}

// IntersectionSortStrings returns a∩b, a must be unique
func IntersectionSortStrings(a []string, b []string) []string {
	SortStrings(b)
	p := IntersectionSorted(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) int { return CompareString(a[ia], b[ib]) })
	return a[:p]
}

// DeepIntersectionInts returns a∩b
func DeepIntersectionInts(a []int, b []int, c ...[]int) []int {
	var s = IntersectionInts(a, b)
	for _, d := range c {
		s = IntersectionInts(s, d)
	}
	return DistinctInts(s)
}

// DeepIntersectionInt32s returns a∩b
func DeepIntersectionInt32s(a []int32, b []int32, c ...[]int32) []int32 {
	var s = IntersectionInt32s(a, b)
	for _, d := range c {
		s = IntersectionInt32s(s, d)
	}
	return DistinctInt32s(s)
}

// DeepIntersectionInt64s returns a∩b
func DeepIntersectionInt64s(a []int64, b []int64, c ...[]int64) []int64 {
	var s = IntersectionInt64s(a, b)
	for _, d := range c {
		s = IntersectionInt64s(s, d)
	}
	return DistinctInt64s(s)
}

// DeepIntersectionUints returns a∩b
func DeepIntersectionUints(a []uint, b []uint, c ...[]uint) []uint {
	var s = IntersectionUints(a, b)
	for _, d := range c {
		s = IntersectionUints(s, d)
	}
	return DistinctUints(s)
}

// DeepIntersectionUint32s returns a∩b
func DeepIntersectionUint32s(a []uint32, b []uint32, c ...[]uint32) []uint32 {
	var s = IntersectionUint32s(a, b)
	for _, d := range c {
		s = IntersectionUint32s(s, d)
	}
	return DistinctUint32s(s)
}

// DeepIntersectionUint64s returns a∩b
func DeepIntersectionUint64s(a []uint64, b []uint64, c ...[]uint64) []uint64 {
	var s = IntersectionUint64s(a, b)
	for _, d := range c {
		s = IntersectionUint64s(s, d)
	}
	return DistinctUint64s(s)
}

// DeepIntersectionStrings returns a∩b
func DeepIntersectionStrings(a []string, b []string, c ...[]string) []string {
	var s = IntersectionStrings(a, b)
	for _, d := range c {
		s = IntersectionStrings(s, d)
	}
	return DistinctStrings(s)
}
