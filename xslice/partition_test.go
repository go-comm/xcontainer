package xslice

import (
	"context"
	"testing"
)

func TestPartition(t *testing.T) {

	var array []int
	for i := 1; i <= 101; i++ {
		array = append(array, i)
	}

	var cases = []struct {
		Param1 []int
		Param2 int

		Want []int
	}{
		{Param1: array, Param2: 10, Want: []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100, 101}},
		{Param1: array, Param2: 11, Want: []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 101}},
		{Param1: array, Param2: 100, Want: []int{100, 101}},
		{Param1: array, Param2: 101, Want: []int{101}},
		{Param1: array, Param2: 200, Want: []int{101}},
		{Param1: []int{1, 2, 3, 4, 5, 6}, Param2: 1, Want: []int{1, 2, 3, 4, 5, 6}},
		{Param1: []int{1, 2, 3, 4, 5, 6}, Param2: 3, Want: []int{3, 6}},
	}

	for _, c := range cases {

		var got []int
		PartitionInts(context.TODO(), c.Param1, c.Param2, func(sub []int) error {
			got = append(got, sub[len(sub)-1])
			return nil
		})

		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, want: %v, but got: %v", c.Param1, c.Want, got)
		}
	}

}
