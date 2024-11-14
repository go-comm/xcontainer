package xslice

// AppendUnionInts returns a∪b, a must be unique
func AppendUnionInts(a []int, b []int) []int {
	c := FilterInts(b, func(i int) bool { return FindInts(a, b[i]) < 0 })
	c = DistinctUnstableInts(c)
	return append(a, c...)
}

// AppendUnionInt32s returns a∪b, a must be unique
func AppendUnionInt32s(a []int32, b []int32) []int32 {
	c := FilterInt32s(b, func(i int) bool { return FindInt32s(a, b[i]) < 0 })
	c = DistinctUnstableInt32s(c)
	return append(a, c...)
}

// AppendUnionInt64s returns a∪b, a must be unique
func AppendUnionInt64s(a []int64, b []int64) []int64 {
	c := FilterInt64s(b, func(i int) bool { return FindInt64s(a, b[i]) < 0 })
	c = DistinctUnstableInt64s(c)
	return append(a, c...)
}

// AppendUnionUints returns a∪b, a must be unique
func AppendUnionUints(a []uint, b []uint) []uint {
	c := FilterUints(b, func(i int) bool { return FindUints(a, b[i]) < 0 })
	c = DistinctUnstableUints(c)
	return append(a, c...)
}

// AppendUnionUint32s returns a∪b, a must be unique
func AppendUnionUint32s(a []uint32, b []uint32) []uint32 {
	c := FilterUint32s(b, func(i int) bool { return FindUint32s(a, b[i]) < 0 })
	c = DistinctUnstableUint32s(c)
	return append(a, c...)
}

// AppendUnionUint64s returns a∪b, a must be unique
func AppendUnionUint64s(a []uint64, b []uint64) []uint64 {
	c := FilterUint64s(b, func(i int) bool { return FindUint64s(a, b[i]) < 0 })
	c = DistinctUnstableUint64s(c)
	return append(a, c...)
}

// AppendUnionStrings returns a∪b, a must be unique
func AppendUnionStrings(a []string, b []string) []string {
	c := FilterStrings(b, func(i int) bool { return FindStrings(a, b[i]) < 0 })
	c = DistinctUnstableStrings(c)
	return append(a, c...)
}

// DeepUnionInts returns a∪b∪c
func DeepUnionInts(a []int, b []int, c ...[]int) []int {
	var s = DistinctInts(a)
	s = AppendUnionInts(s, b)
	for _, d := range c {
		s = AppendUnionInts(s, d)
	}
	return s
}

// DeepUnionInt32s returns a∪b∪c
func DeepUnionInt32s(a []int32, b []int32, c ...[]int32) []int32 {
	var s = DistinctInt32s(a)
	s = AppendUnionInt32s(s, b)
	for _, d := range c {
		s = AppendUnionInt32s(s, d)
	}
	return s
}

// DeepUnionInt64s returns a∪b∪c
func DeepUnionInt64s(a []int64, b []int64, c ...[]int64) []int64 {
	var s = DistinctInt64s(a)
	s = AppendUnionInt64s(s, b)
	for _, d := range c {
		s = AppendUnionInt64s(s, d)
	}
	return s
}

// DeepUnionUints returns a∪b∪c
func DeepUnionUints(a []uint, b []uint, c ...[]uint) []uint {
	var s = DistinctUints(a)
	s = AppendUnionUints(s, b)
	for _, d := range c {
		s = AppendUnionUints(s, d)
	}
	return s
}

// DeepUnionUint32s returns a∪b∪c
func DeepUnionUint32s(a []uint32, b []uint32, c ...[]uint32) []uint32 {
	var s = DistinctUint32s(a)
	s = AppendUnionUint32s(s, b)
	for _, d := range c {
		s = AppendUnionUint32s(s, d)
	}
	return s
}

// DeepUnionUint64s returns a∪b∪c
func DeepUnionUint64s(a []uint64, b []uint64, c ...[]uint64) []uint64 {
	var s = DistinctUint64s(a)
	s = AppendUnionUint64s(s, b)
	for _, d := range c {
		s = AppendUnionUint64s(s, d)
	}
	return s
}

// DeepUnionStrings returns a∪b∪c
func DeepUnionStrings(a []string, b []string, c ...[]string) []string {
	var s = DistinctStrings(a)
	s = AppendUnionStrings(s, b)
	for _, d := range c {
		s = AppendUnionStrings(s, d)
	}
	return s
}
