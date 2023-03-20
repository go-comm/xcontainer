package xslice

func CloneInts(arr []int) []int {
	b := make([]int, 0, len(arr))
	b = append(b, arr...)
	return b
}

func CloneInt32s(arr []int32) []int32 {
	b := make([]int32, 0, len(arr))
	b = append(b, arr...)
	return b
}

func CloneInt64s(arr []int64) []int64 {
	b := make([]int64, 0, len(arr))
	b = append(b, arr...)
	return b
}

func CloneStrings(arr []string) []string {
	b := make([]string, 0, len(arr))
	b = append(b, arr...)
	return b
}
