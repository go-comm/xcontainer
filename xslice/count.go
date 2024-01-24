package xslice

func Count(length int, condition func(i int) bool) int {
	return CountN(length, 0, length, condition)
}

func CountN(length int, i int, j int, condition func(i int) bool) int {
	var c int
	for k := i; k < j; k++ {
		if condition(k) {
			c++
		}
	}
	return c
}

func CountToInt(length int, condition func(i int) bool) int {
	return Count(length, condition)
}

func CountToInt32(length int, condition func(i int) bool) int32 {
	return int32(Count(length, condition))
}

func CountToUint32(length int, condition func(i int) bool) uint32 {
	return uint32(Count(length, condition))
}

func CountToInt64(length int, condition func(i int) bool) int64 {
	return int64(Count(length, condition))
}

func CountToUint64(length int, condition func(i int) bool) uint64 {
	return uint64(Count(length, condition))
}
