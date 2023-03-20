package xslice

func Equal(na int, nb int, fn func(i int) bool) bool {
	if na != nb {
		return false
	}
	for i := 0; i < na; i++ {
		if !fn(i) {
			return false
		}
	}
	return true
}

func EqualInts(a []int, b []int) bool {
	return Equal(len(a), len(b), func(i int) bool { return a[i] == b[i] })
}

func EqualInt32s(a []int32, b []int32) bool {
	return Equal(len(a), len(b), func(i int) bool { return a[i] == b[i] })
}

func EqualInt64s(a []int64, b []int64) bool {
	return Equal(len(a), len(b), func(i int) bool { return a[i] == b[i] })
}

func EqualStrings(a []string, b []string) bool {
	return Equal(len(a), len(b), func(i int) bool { return a[i] == b[i] })
}
