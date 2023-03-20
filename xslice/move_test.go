package xslice

import "testing"

func TestMove(t *testing.T) {

	var cases = []struct {
		Param1 []int
		Param2 int
		Param3 int

		Want []int
	}{
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 1, Param3: 3, Want: []int{0, 2, 3, 1, 4, 5, 6, 7}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 4, Param3: 1, Want: []int{0, 4, 1, 2, 3, 5, 6, 7}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 7, Param3: 1, Want: []int{0, 7, 1, 2, 3, 4, 5, 6}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 7, Param3: 0, Want: []int{7, 0, 1, 2, 3, 4, 5, 6}},
		{Param1: []int{0, 1, 2, 3, 4, 5, 6, 7}, Param2: 0, Param3: 7, Want: []int{1, 2, 3, 4, 5, 6, 7, 0}},
	}

	for _, c := range cases {
		got := MoveInts(c.Param1, c.Param2, c.Param3)
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, want: %v, but got: %v", c.Param1, c.Want, got)
		}
	}

}
