package xslice

import "time"

func Map(n int, mapper func(i int)) {
	for i := 0; i < n; i++ {
		mapper(i)
	}
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

func MapToUints(length int, mapper func(int) uint) []uint {
	b := make([]uint, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}

func MapToUint32s(length int, mapper func(int) uint32) []uint32 {
	b := make([]uint32, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}

func MapToUint64s(length int, mapper func(int) uint64) []uint64 {
	b := make([]uint64, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}

func MapToFloat32s(length int, mapper func(int) float32) []float32 {
	b := make([]float32, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}

func MapToFloat64s(length int, mapper func(int) float64) []float64 {
	b := make([]float64, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}

func MapToStrings(length int, mapper func(int) string) []string {
	b := make([]string, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}

func MapToTimes(length int, mapper func(int) time.Time) []time.Time {
	b := make([]time.Time, length)
	Map(length, func(i int) { b[i] = mapper(i) })
	return b
}
