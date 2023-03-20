package xslice

// ExceptInts returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptInts(a []int, b []int) []int {
	mb := make(map[int]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// ExceptInt32s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptInt32s(a []int32, b []int32) []int32 {
	mb := make(map[int32]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int32
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// ExceptInt64s returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptInt64s(a []int64, b []int64) []int64 {
	mb := make(map[int64]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []int64
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

// ExceptStrings returns the elements in `a` that aren't in `b`. a-(a∩b)
func ExceptStrings(a []string, b []string) []string {
	mb := make(map[string]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []string
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}
