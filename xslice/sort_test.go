package xslice

import "testing"

func TestSort(t *testing.T) {

	var cases = []struct {
		Param1 []int

		Want []int
	}{
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Want: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7, 7}, Want: []int{0, 1, 2, 3, 4, 5, 6, 7, 7}},
		{Param1: []int{7, 6, 5, 4, 3, 2, 1, 0}, Want: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{Param1: []int{1, 1, 1, 1, 1, 1, 1}, Want: []int{1, 1, 1, 1, 1, 1, 1}},
		{Param1: []int{1}, Want: []int{1}},
		{Param1: []int{1, 3, 2}, Want: []int{1, 2, 3}},
		{Param1: []int{1, 3, 2, 8, 2}, Want: []int{1, 2, 2, 3, 8}},
		{Param1: []int{1, 1, 3, 3, 2, 2}, Want: []int{1, 1, 2, 2, 3, 3}},
		{Param1: []int{0, 1, 1, 3, 3, 2, 2, 4, 2}, Want: []int{0, 1, 1, 2, 2, 2, 3, 3, 4}},
	}

	for _, c := range cases {
		arr := CloneInts(c.Param1)
		got := SortInts(arr)
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, want: %v, but got: %v", c.Param1, c.Want, got)
		}
	}

}

func TestInsertSort(t *testing.T) {

	var cases = []struct {
		Param1 []int

		Want []int
	}{
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Want: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7, 7}, Want: []int{0, 1, 2, 3, 4, 5, 6, 7, 7}},
		{Param1: []int{7, 6, 5, 4, 3, 2, 1, 0}, Want: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{Param1: []int{1, 1, 1, 1, 1, 1, 1}, Want: []int{1, 1, 1, 1, 1, 1, 1}},
		{Param1: []int{1}, Want: []int{1}},
		{Param1: []int{1, 3, 2}, Want: []int{1, 2, 3}},
		{Param1: []int{1, 3, 2, 8, 2}, Want: []int{1, 2, 2, 3, 8}},
		{Param1: []int{1, 1, 3, 3, 2, 2}, Want: []int{1, 1, 2, 2, 3, 3}},
		{Param1: []int{0, 1, 1, 3, 3, 2, 2, 4, 2}, Want: []int{0, 1, 1, 2, 2, 2, 3, 3, 4}},
	}

	for _, c := range cases {
		arr := CloneInts(c.Param1)
		insertSort(0, len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] < arr[j] })
		got := arr
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, want: %v, but got: %v", c.Param1, c.Want, got)
		}
	}

}

func TestHeapSort(t *testing.T) {

	var cases = []struct {
		Param1 []int

		Want []int
	}{
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Want: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7, 7}, Want: []int{0, 1, 2, 3, 4, 5, 6, 7, 7}},
		{Param1: []int{7, 6, 5, 4, 3, 2, 1, 0}, Want: []int{0, 1, 2, 3, 4, 5, 6, 7}},
		{Param1: []int{1, 1, 1, 1, 1, 1, 1}, Want: []int{1, 1, 1, 1, 1, 1, 1}},
		{Param1: []int{1}, Want: []int{1}},
		{Param1: []int{1, 3, 2}, Want: []int{1, 2, 3}},
		{Param1: []int{1, 3, 2, 8, 2}, Want: []int{1, 2, 2, 3, 8}},
		{Param1: []int{1, 1, 3, 3, 2, 2}, Want: []int{1, 1, 2, 2, 3, 3}},
		{Param1: []int{0, 1, 1, 3, 3, 2, 2, 4, 2}, Want: []int{0, 1, 1, 2, 2, 2, 3, 3, 4}},
	}

	for _, c := range cases {
		arr := CloneInts(c.Param1)
		heapSort(0, len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] < arr[j] })
		got := arr
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, want: %v, but got: %v", c.Param1, c.Want, got)
		}
	}

}

func TestFixSort(t *testing.T) {

	var cases = []struct {
		Param1 []int
		Param2 int

		Want []int
	}{
		{Param1: []int{0, 1, 2, 3, 4, 5}, Param2: 0, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{0, 1, 2, 3, 4, 5}, Param2: 1, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{0, 1, 2, 3, 4, 5}, Param2: 3, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{0, 1, 2, 3, 4, 5}, Param2: 4, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{0, 1, 2, 3, 4, 5}, Param2: 5, Want: []int{0, 1, 2, 3, 4, 5}},

		{Param1: []int{0, 1, 3, 4, 5, 2}, Param2: 5, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{1, 2, 3, 4, 5, 0}, Param2: 5, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{0, 1, 3, 4, 2, 5}, Param2: 4, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{0, 1, 2, 3, 5, 4}, Param2: 4, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{0, 5, 1, 2, 3, 4}, Param2: 1, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{1, 0, 2, 3, 4, 5}, Param2: 1, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{5, 0, 1, 2, 3, 4}, Param2: 0, Want: []int{0, 1, 2, 3, 4, 5}},
		{Param1: []int{4, 0, 1, 2, 3, 5}, Param2: 0, Want: []int{0, 1, 2, 3, 4, 5}},
	}

	for _, c := range cases {
		arr := CloneInts(c.Param1)
		FixSort(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] < arr[j] }, c.Param2)
		got := arr
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, want: %v, but got: %v", c.Param1, c.Want, got)
		}
	}

}
