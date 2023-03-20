package xslice

func Map(n int, mapper func(i int)) {
	for i := 0; i < n; i++ {
		mapper(i)
	}
}

func MapInts(a []int, mapper func(int) int) []int {
	b := make([]int, len(a))
	Map(len(a), func(i int) { b[i] = mapper(a[i]) })
	return b
}

func MapInt32s(a []int32, mapper func(int32) int32) []int32 {
	b := make([]int32, len(a))
	Map(len(a), func(i int) { b[i] = mapper(a[i]) })
	return b
}

func MapInt64s(a []int64, mapper func(int64) int64) []int64 {
	b := make([]int64, len(a))
	Map(len(a), func(i int) { b[i] = mapper(a[i]) })
	return b
}

func MapStrings(a []string, mapper func(string) string) []string {
	b := make([]string, len(a))
	Map(len(a), func(i int) { b[i] = mapper(a[i]) })
	return b
}

func MapToInts(length int, mapper func(int) int) []int {
	b := make([]int, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}

func MapToInt32s(length int, mapper func(int) int32) []int32 {
	b := make([]int32, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}

func MapToInt64s(length int, mapper func(int) int64) []int64 {
	b := make([]int64, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}

func MapToStrings(length int, mapper func(int) string) []string {
	b := make([]string, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}
