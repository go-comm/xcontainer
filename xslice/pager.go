package xslice

func Pager(length int, pn int, ps int, fn func(i int, j int)) {
	from := (pn - 1) * ps
	to := pn * ps

	if from >= 0 && from < length {
		if to > length {
			to = length
		}
		fn(from, to)
	}
}

func PagerInts(arr []int, pn int, ps int, fn func(sub []int)) {
	Pager(len(arr), pn, ps, func(i, j int) { fn(arr[i:j]) })
}

func PagerInt32s(arr []int32, pn int, ps int, fn func(sub []int32)) {
	Pager(len(arr), pn, ps, func(i, j int) { fn(arr[i:j]) })
}

func PagerInt64s(arr []int64, pn int, ps int, fn func(sub []int64)) {
	Pager(len(arr), pn, ps, func(i, j int) { fn(arr[i:j]) })
}

func PagerStrings(arr []string, pn int, ps int, fn func(sub []string)) {
	Pager(len(arr), pn, ps, func(i, j int) { fn(arr[i:j]) })
}
