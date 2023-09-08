package xslice

import "testing"

func TestIntersectionInts(t *testing.T) {

	var cases = []struct {
		Param1 []int
		Param2 []int
		Want   []int
	}{
		{Param1: []int{1, 2, 3, 4, 5}, Param2: []int{6, 7, 8}, Want: []int{}},
		{Param1: []int{1, 2, 3, 4, 5}, Param2: []int{1, 2, 3, 4, 5}, Want: []int{1, 2, 3, 4, 5}},
		{Param1: []int{1, 2, 3, 4, 5}, Param2: []int{2, 3, 4, 5}, Want: []int{2, 3, 4, 5}},
		{Param1: []int{1, 2, 3, 4, 5}, Param2: []int{4, 5, 6, 7}, Want: []int{4, 5}},
	}

	for _, c := range cases {
		got := IntersectionInts(c.Param1, c.Param2)
		got = SortInts(got)
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, param2: %v, want: %v, but got: %v", c.Param1, c.Param2, c.Want, got)
		}
	}

}
