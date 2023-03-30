package xslice

import (
	"testing"
)

func TestPager(t *testing.T) {

	var array []int
	for i := 1; i <= 101; i++ {
		array = append(array, i)
	}

	var cases = []struct {
		Param1 []int
		Param2 int
		Param3 int

		Want []int
	}{
		{Param1: array, Param2: 1, Param3: 5, Want: []int{1, 2, 3, 4, 5}},
		{Param1: array, Param2: 2, Param3: 10, Want: []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}},
		{Param1: array, Param2: 11, Param3: 10, Want: []int{101}},
	}

	for _, c := range cases {
		var got []int
		Pager(len(c.Param1), c.Param2, c.Param3, func(i, j int) {
			got = c.Param1[i:j]
		})

		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, want: %v, but got: %v", c.Param1, c.Want, got)
		}
	}

}
