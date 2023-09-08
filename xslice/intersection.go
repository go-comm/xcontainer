package xslice

// IntersectionStable returns the elements in `a` and `b`. a∩b, a must be unique
func IntersectionStable(lenA int, lenB int, equal func(ia int, ib int) bool, push func(ia int)) int {
	size := 0
	var found bool
	for ia := 0; ia < lenA; ia++ {
		found = false
		for ib := 0; ib < lenB; ib++ {
			if equal(ia, ib) {
				found = true
				break
			}
		}
		if found {
			push(ia)
			size++
		}
	}
	return size
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
