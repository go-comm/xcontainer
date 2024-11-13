package xslice

import "time"

func ConcatInts(a []int, b []int) []int {
	c := append([]int(nil), a...)
	return append(c, b...)
}

func ConcatInt32s(a []int32, b []int32) []int32 {
	c := append([]int32(nil), a...)
	return append(c, b...)
}

func ConcatInt64s(a []int64, b []int64) []int64 {
	c := append([]int64(nil), a...)
	return append(c, b...)
}

func ConcatUints(a []uint, b []uint) []uint {
	c := append([]uint(nil), a...)
	return append(c, b...)
}

func ConcatUint32s(a []uint32, b []uint32) []uint32 {
	c := append([]uint32(nil), a...)
	return append(c, b...)
}

func ConcatUint64s(a []uint64, b []uint64) []uint64 {
	c := append([]uint64(nil), a...)
	return append(c, b...)
}

func ConcatStrings(a []string, b []string) []string {
	c := append([]string(nil), a...)
	return append(c, b...)
}

func ConcatTimes(a []time.Time, b []time.Time) []time.Time {
	c := append([]time.Time(nil), a...)
	return append(c, b...)
}
