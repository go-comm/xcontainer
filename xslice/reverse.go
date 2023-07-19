package xslice

import "reflect"

func Reverse(length int, swap func(i int, j int)) {
	for i := 0; i < length/2; i++ {
		swap(i, length-1-i)
	}
}

func ReverseSlice(slice interface{}) {
	rv := reflect.ValueOf(slice)
	Reverse(rv.Len(), Swapper(slice))
}

func ReverseInts(arr []int) []int {
	Reverse(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func ReverseInt32s(arr []int32) []int32 {
	Reverse(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func ReverseInt64s(arr []int64) []int64 {
	Reverse(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}

func ReverseStrings(arr []string) []string {
	Reverse(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr
}
