package xslice

import "time"

func SetToInts(length int, mapper func(int) int) []int {
	return DistinctInts(MapToInts(length, mapper))
}

func SetToInt32s(length int, mapper func(int) int32) []int32 {
	return DistinctInt32s(MapToInt32s(length, mapper))
}

func SetToInt64s(length int, mapper func(int) int64) []int64 {
	return DistinctInt64s(MapToInt64s(length, mapper))
}

func SetToUints(length int, mapper func(int) uint) []uint {
	return DistinctUints(MapToUints(length, mapper))
}

func SetToUint32s(length int, mapper func(int) uint32) []uint32 {
	return DistinctUint32s(MapToUint32s(length, mapper))
}

func SetToUint64s(length int, mapper func(int) uint64) []uint64 {
	return DistinctUint64s(MapToUint64s(length, mapper))
}

func SetToFloat32s(length int, mapper func(int) float32) []float32 {
	return DistinctFloat32s(MapToFloat32s(length, mapper))
}

func SetToFloat64s(length int, mapper func(int) float64) []float64 {
	return DistinctFloat64s(MapToFloat64s(length, mapper))
}

func SetToStrings(length int, mapper func(int) string) []string {
	return DistinctStrings(MapToStrings(length, mapper))
}

func SetToTimes(length int, mapper func(int) time.Time) []time.Time {
	return DistinctTimes(MapToTimes(length, mapper))
}
