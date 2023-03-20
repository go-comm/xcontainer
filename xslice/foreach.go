package xslice

func Foreach(length int, index func(i int) (exit bool)) {
	for i := 0; i < length; i++ {
		if index(i) {
			break
		}
	}
}

func ForeachInts(arr []int, index func(e int, i int) (exit bool)) {
	Foreach(len(arr), func(i int) (exit bool) { return index(arr[i], i) })
}

func ForeachInt32s(arr []int32, index func(e int32, i int) (exit bool)) {
	Foreach(len(arr), func(i int) (exit bool) { return index(arr[i], i) })
}

func ForeachInt64s(arr []int64, index func(e int64, i int) (exit bool)) {
	Foreach(len(arr), func(i int) (exit bool) { return index(arr[i], i) })
}

func ForeachStrings(arr []string, index func(e string, i int) (exit bool)) {
	Foreach(len(arr), func(i int) (exit bool) { return index(arr[i], i) })
}

func Forrange(length int, index func(i int)) {
	for i := 0; i < length; i++ {
		index(i)
	}
}

func ForrangeInts(arr []int, index func(e int, i int)) {
	Forrange(len(arr), func(i int) { index(arr[i], i) })
}

func ForrangeInt32s(arr []int32, index func(e int32, i int)) {
	Forrange(len(arr), func(i int) { index(arr[i], i) })
}

func ForrangeInt64s(arr []int64, index func(e int64, i int)) {
	Forrange(len(arr), func(i int) { index(arr[i], i) })
}

func ForrangeStrings(arr []string, index func(e string, i int)) {
	Forrange(len(arr), func(i int) { index(arr[i], i) })
}
