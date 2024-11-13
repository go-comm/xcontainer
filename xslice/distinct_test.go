package xslice

import "testing"

func TestDistinctStable(t *testing.T) {

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
		var dst []int
		DistinctStable(len(arr), func(i, j int) bool { return arr[i] == arr[j] }, func(i int) { dst = append(dst, arr[i]) })
		got := SortInts(dst)
		if !EqualInts(c.Want, SortInts(got)) {
			t.Fatalf("param: %v, want: %v, but got: %v", c.Param, c.Want, got)
		}
	}
}

func TestDistinctUnstable(t *testing.T) {

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
		{Param: []int{1, 2, 2, 3, 3, 4}, Want: []int{1, 2, 3, 4}},
		{Param: []int{1, 1, 1, 1, 1, 1, 1}, Want: []int{1}},
		{Param: []int{1, 1, 1, 1, 1, 1, 2}, Want: []int{1, 2}},
		{Param: []int{1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 5}, Want: []int{1, 2, 3, 5}},
		{Param: []int{1, 2, 3, 4, 6}, Want: []int{1, 2, 3, 4, 6}},
		{Param: []int{1, 1, 2, 2}, Want: []int{1, 2}},
	}

	for _, c := range cases {
		arr := c.Param
		p := DistinctSorted(len(arr), func(i, j int) { arr[i], arr[j] = arr[j], arr[i] }, func(i, j int) int { return CompareInt(arr[i], arr[j]) })
		got := SortInts(arr[:p])
		if !EqualInts(c.Want, SortInts(got)) {
			t.Fatalf("param: %v, want: %v, but got: %v", c.Param, c.Want, got)
		}
	}
}

func TestIsDistinct(t *testing.T) {
	var cases = []struct {
		Param []int
		Want1 int
		Want2 bool
	}{
		{Param: []int{1}, Want1: -1, Want2: true},
		{Param: []int{1, 2}, Want1: -1, Want2: true},
		{Param: []int{1, 2, 3, 4, 5, 6}, Want1: -1, Want2: true},
		{Param: []int{1, 1}, Want1: 1, Want2: false},
		{Param: []int{1, 2, 2, 3, 3, 4}, Want1: 2, Want2: false},
		{Param: []int{1, 1, 1, 1, 1, 1, 1}, Want1: 1, Want2: false},
	}

	for _, c := range cases {
		arr := c.Param
		got1, got2 := IsDistinctInts(arr)
		if c.Want1 != got1 || c.Want2 != got2 {
			t.Fatalf("param: %v, want: %v %v, but got: %v %v", c.Param, c.Want1, c.Want2, got1, got2)
		}
	}
}
