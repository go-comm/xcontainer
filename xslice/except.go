package xslice

// Except returns the elements in `a` that aren't in `b`. a-(a∩b)
func Except(lenA int, lenB int, swap func(i int, j int), equal func(ia int, ib int) bool) int {
	hi := lenA
	for i := hi - 1; i >= 0; i-- {
		for j := 0; j < lenB; j++ {
			if equal(i, j) {
				if i != hi-1 {
					swap(i, hi-1)
				}
				hi--
				break
			}
		}
	}
	return hi
}

// ExceptInts returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptInts(a []int, b []int) []int {
	p := Except(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// ExceptInt32s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptInt32s(a []int32, b []int32) []int32 {
	p := Except(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// ExceptInt64s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptInt64s(a []int64, b []int64) []int64 {
	p := Except(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}

// ExceptStrings returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptStrings(a []string, b []string) []string {
	p := Except(len(a), len(b), func(i, j int) { a[i], a[j] = a[j], a[i] }, func(ia, ib int) bool { return a[ia] == b[ib] })
	return a[:p]
}
