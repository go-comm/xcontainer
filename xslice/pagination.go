package xslice

var (
	Pager        = Pagination
	PagerInts    = PaginationInts
	PagerInt32s  = PaginationInt32s
	PagerInt64s  = PaginationInt64s
	PagerStrings = PaginationStrings
)

func Pagination(length int, pn int, ps int, fn func(i int, j int)) {
	from := (pn - 1) * ps
	to := pn * ps

	if from >= 0 && from < length {
		if to > length {
			to = length
		}
		fn(from, to)
	}
}

func PaginationInts(arr []int, pn int, ps int, fn func(sub []int)) {
	Pagination(len(arr), pn, ps, func(i, j int) { fn(arr[i:j]) })
}

func PaginationInt32s(arr []int32, pn int, ps int, fn func(sub []int32)) {
	Pagination(len(arr), pn, ps, func(i, j int) { fn(arr[i:j]) })
}

func PaginationInt64s(arr []int64, pn int, ps int, fn func(sub []int64)) {
	Pagination(len(arr), pn, ps, func(i, j int) { fn(arr[i:j]) })
}

func PaginationStrings(arr []string, pn int, ps int, fn func(sub []string)) {
	Pagination(len(arr), pn, ps, func(i, j int) { fn(arr[i:j]) })
}
