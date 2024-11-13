package xslice

import "testing"

func TestUnion(t *testing.T) {

	var cases = []struct {
		Param1 []int
		Param2 []int
		Want   []int
	}{
		{Param1: []int{1}, Param2: []int{1}, Want: []int{1}},
		{Param1: []int{1, 2, 3}, Param2: []int{1, 2, 3}, Want: []int{1, 2, 3}},
		{Param1: []int{1, 2, 3}, Param2: []int{2}, Want: []int{1, 2, 3}},
		{Param1: []int{1, 2, 3}, Param2: []int{2, 3, 4}, Want: []int{1, 2, 3, 4}},
		{Param1: []int{1, 2}, Param2: []int{2, 3}, Want: []int{1, 2, 3}},
		{Param1: []int{1, 2, 3}, Param2: []int{4, 5, 6}, Want: []int{1, 2, 3, 4, 5, 6}},
	}

	for _, c := range cases {
		dst := UnionInts(c.Param1, c.Param2)
		got := SortInts(dst)
		if !EqualInts(c.Want, SortInts(got)) {
			t.Fatalf("param1: %v, param2: %v, want: %v, but got: %v", c.Param1, c.Param2, c.Want, got)
		}
	}
}
