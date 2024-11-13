package xslice

func UnionInts(a []int, b []int) []int {
	a = DistinctInts(a)
	c := FilterInts(b, func(i int) bool { return FindInts(a, b[i]) < 0 })
	c = DistinctUnstableInts(c)
	return append(a, c...)
}

func UnionInt32s(a []int32, b []int32) []int32 {
	a = DistinctInt32s(a)
	c := FilterInt32s(b, func(i int) bool { return FindInt32s(a, b[i]) < 0 })
	c = DistinctUnstableInt32s(c)
	return append(a, c...)
}

func UnionInt64s(a []int64, b []int64) []int64 {
	a = DistinctInt64s(a)
	c := FilterInt64s(b, func(i int) bool { return FindInt64s(a, b[i]) < 0 })
	c = DistinctUnstableInt64s(c)
	return append(a, c...)
}

func UnionUints(a []uint, b []uint) []uint {
	a = DistinctUints(a)
	c := FilterUints(b, func(i int) bool { return FindUints(a, b[i]) < 0 })
	c = DistinctUnstableUints(c)
	return append(a, c...)
}

func UnionUint32s(a []uint32, b []uint32) []uint32 {
	a = DistinctUint32s(a)
	c := FilterUint32s(b, func(i int) bool { return FindUint32s(a, b[i]) < 0 })
	c = DistinctUnstableUint32s(c)
	return append(a, c...)
}

func UnionUint64s(a []uint64, b []uint64) []uint64 {
	a = DistinctUint64s(a)
	c := FilterUint64s(b, func(i int) bool { return FindUint64s(a, b[i]) < 0 })
	c = DistinctUnstableUint64s(c)
	return append(a, c...)
}

func UnionStrings(a []string, b []string) []string {
	a = DistinctStrings(a)
	c := FilterStrings(b, func(i int) bool { return FindStrings(a, b[i]) < 0 })
	c = DistinctUnstableStrings(c)
	return append(a, c...)
}
