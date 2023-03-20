package xslice

func Splice(length int, i int, howmany int, swap func(i int, j int)) int {
	if howmany <= 0 {
		return length
	}
	if i+howmany < length {
		for k := i; k < length && k+howmany < length; k++ {
			swap(k, k+howmany)
		}
	}
	return length - howmany
}

func SpliceInts(arr []int, i int, howmany int) []int {
	p := Splice(len(arr), i, howmany, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func SpliceInt32s(arr []int32, i int, howmany int) []int32 {
	p := Splice(len(arr), i, howmany, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func SpliceInt64s(arr []int64, i int, howmany int) []int64 {
	p := Splice(len(arr), i, howmany, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}

func SpliceStrings(arr []string, i int, howmany int) []string {
	p := Splice(len(arr), i, howmany, func(i, j int) { arr[i], arr[j] = arr[j], arr[i] })
	return arr[:p]
}
