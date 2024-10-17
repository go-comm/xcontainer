package xslice

import (
	"reflect"
	"time"
)

var Swapper func(slice interface{}) func(int, int) = reflect.Swapper

func SwapInt8s(slice []int8) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapInt16s(slice []int16) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapInts(slice []int) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapInt32s(slice []int32) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapInt64s(slice []int64) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapUint8s(slice []int8) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapUint16s(slice []uint16) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapUints(slice []uint) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapUint32s(slice []uint32) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapUint64s(slice []uint64) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapFloat32s(slice []float32) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapFloat64s(slice []float64) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapBytes(slice []byte) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapStrings(slice []string) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapTimes(slice []time.Time) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}

func SwapDurations(slice []time.Duration) func(int, int) {
	return func(i, j int) { slice[i], slice[j] = slice[j], slice[i] }
}
