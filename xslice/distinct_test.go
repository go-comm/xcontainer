package xslice

import "testing"

func TestDistinctUnsorted(t *testing.T) {

	var cases = []struct {
		Param []int
		Want  []int
	}{
		{Param: []int{1}, Want: []int{1}},
		{Param: []int{1, 1}, Want: []int{1}},
		{Param: []int{1, 2, 2, 3, 3, 4}, Want: []int{1, 2, 3, 4}},
		{Param: []int{1, 1, 1, 1, 1, 1, 1}, Want: []int{1}},
		{Param: []int{1, 1, 1, 1, 1, 1, 2}, Want: []int{1, 2}},
		{Param: []int{1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5}, Want: []int{1, 2, 3, 5}},
		{Param: []int{1, 2, 3, 4, 6}, Want: []int{1, 2, 3, 4, 6}},
		{Param: []int{1, 1, 2, 2}, Want: []int{1, 2}},
	}

	for _, c := range cases {
		arr := c.Param
		p := DistinctUnsorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
		got := SortInts(arr[:p])
		if !EqualInts(c.Want, got) {
			t.Fatalf("param: %v, want: %v, but got: %v", c.Param, c.Want, got)
		}
	}

}

func TestDistinctSort(t *testing.T) {

	var cases = []struct {
		Param []int
		Want  []int
	}{
		{Param: []int{1}, Want: []int{1}},
		{Param: []int{1, 1}, Want: []int{1}},
		{Param: []int{1, 2, 2, 3, 3, 4}, Want: []int{1, 2, 3, 4}},
		{Param: []int{1, 1, 1, 1, 1, 1, 1}, Want: []int{1}},
		{Param: []int{1, 1, 1, 1, 1, 1, 2}, Want: []int{1, 2}},
		{Param: []int{1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5}, Want: []int{1, 2, 3, 5}},
		{Param: []int{1, 2, 3, 4, 6}, Want: []int{1, 2, 3, 4, 6}},
		{Param: []int{1, 1, 2, 2}, Want: []int{1, 2}},
	}

	for _, c := range cases {
		arr := c.Param
		p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return arr[i] - arr[j] })
		got := SortInts(arr[:p])
		if !EqualInts(c.Want, got) {
			t.Fatalf("param: %v, want: %v, but got: %v", c.Param, c.Want, got)
		}
	}
}
