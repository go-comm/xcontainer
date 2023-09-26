package xslice

import "testing"

func TestDistinct(t *testing.T) {

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
		p := DistinctUnstable(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) bool { return arr[i] == arr[j] })
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
		{Param: []int{1, 5, 2, 3, 3, 4}, Want: []int{1, 2, 3, 4, 5}},
		{Param: []int{1, 1, 1, 1, 1, 1, 1}, Want: []int{1}},
		{Param: []int{1, 1, 1, 1, 1, 1, 2}, Want: []int{1, 2}},
		{Param: []int{1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5}, Want: []int{1, 2, 3, 5}},
		{Param: []int{5, 3, 3, 3, 5, 1, 1, 1, 3, 3, 3, 1, 1, 2, 3, 3, 3, 3}, Want: []int{1, 2, 3, 5}},
		{Param: []int{1, 2, 3, 4, 6}, Want: []int{1, 2, 3, 4, 6}},
		{Param: []int{1, 1, 2, 2}, Want: []int{1, 2}},
	}

	for _, c := range cases {
		arr := c.Param
		got := DistinctInts(arr)
		if !EqualInts(c.Want, SortInts(got)) {
			t.Fatalf("param: %v, want: %v, but got: %v", c.Param, c.Want, got)
		}
	}
}

func TestDistinctStable(t *testing.T) {

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
		got := DistinctUnstableInts(arr)
		if !EqualInts(c.Want, SortInts(got)) {
			t.Fatalf("param: %v, want: %v, but got: %v", c.Param, c.Want, got)
		}
	}
}
