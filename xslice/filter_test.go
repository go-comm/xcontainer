package xslice

import "testing"

func TestFilter(t *testing.T) {

	var cases = []struct {
		Param1 []int
		Param2 int

		Want []int
	}{
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 1, Want: []int{2, 3, 4, 5, 6, 7}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 2, Want: []int{3, 4, 5, 6, 7}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 3, Want: []int{4, 5, 6, 7}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 4, Want: []int{5, 6, 7}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 5, Want: []int{6, 7}},
	}

	for _, c := range cases {
		got := FilterInts(c.Param1, func(i int) bool { return c.Param1[i] > c.Param2 })
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, want: %v, but got: %v", c.Param1, c.Want, got)
		}
	}

}
