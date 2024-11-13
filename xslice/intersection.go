package xslice

// IntersectionStable returns the elements in `a` and `b`. a∩b, a must be unique
func IntersectionStable(lenA int, lenB int, equal func(ia int, ib int) bool, push func(ia int)) int {
	return FilterStable(lenA, func(i int) bool {
		return Find(lenB, func(j int) bool { return equal(i, j) }) >= 0
	}, push)
}

// IntersectionUnstable returns the elements in `a` and `b`. a∩b, a must be unique
func IntersectionUnstable(lenA int, lenB int, swapA func(i int, j int), equal func(ia int, ib int) bool) (ia int) {
	return FilterUnstable(lenA, func(i int) bool {
		return Find(lenB, func(j int) bool { return equal(i, j) }) >= 0
	}, swapA)
}

// Intersection returns the elements in `a` and `b`. a∩b, a must be unique
func Intersection(lenA int, lenB int, equal func(ia int, ib int) bool, push func(ia int)) int {
	return IntersectionStable(lenA, lenB, equal, push)
}

// IntersectionInts returns the elements in `a` and `b`. a∩b, a must be unique
func IntersectionInts(a []int, b []int) []int {
	var c []int
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionInt32s returns the elements in `a` and `b`. a∩b, a must be unique
func IntersectionInt32s(a []int32, b []int32) []int32 {
	var c []int32
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionInt64s returns the elements in `a` and `b`. a∩b, a must be unique
func IntersectionInt64s(a []int64, b []int64) []int64 {
	var c []int64
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionStrings returns the elements in `a` and `b`. a∩b, a must be unique
func IntersectionStrings(a []string, b []string) []string {
	var c []string
	IntersectionStable(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] }, func(ia int) { c = append(c, a[ia]) })
	return c
}

// IntersectionUnstableInts returns the elements in `a` that aren't in `b`. a-(a∩b)
func IntersectionUnstableInts(a []int, b []int) []int {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionUnstableInt32s returns the elements in `a` that aren't in `b`. a-(a∩b)
func IntersectionUnstableInt32s(a []int32, b []int32) []int32 {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionUnstableInt64s returns the elements in `a` that aren't in `b`. a-(a∩b)
func IntersectionUnstableInt64s(a []int64, b []int64) []int64 {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// IntersectionUnstableStrings returns the elements in `a` that aren't in `b`. a-(a∩b)
func IntersectionUnstableStrings(a []string, b []string) []string {
	p := IntersectionUnstable(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}
