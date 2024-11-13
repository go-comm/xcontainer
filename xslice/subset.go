package xslice

// IsSubset b⊆a
func IsSubset(lenA int, lenB int, equal func(ia int, ib int) bool) (ib int, ok bool) {
	ib = Find(lenB, func(j int) bool {
		return Find(lenA, func(i int) bool { return equal(i, j) }) < 0
	})
	return ib, ib < 0
}

// IsSubsetInts b⊆a
func IsSubsetInts(a []int, b []int) (ib int, ok bool) {
	return IsSubset(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] })
}

// IsSubsetInt32s b⊆a
func IsSubsetInt32s(a []int32, b []int32) (ib int, ok bool) {
	return IsSubset(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] })
}

// IsSubsetInt64s b⊆a
func IsSubsetInt64s(a []int64, b []int64) (ib int, ok bool) {
	return IsSubset(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] })
}

// IsSubsetStrings b⊆a
func IsSubsetStrings(a []string, b []string) (ib int, ok bool) {
	return IsSubset(len(a), len(b), func(ia, ib int) bool { return a[ia] == b[ib] })
}
