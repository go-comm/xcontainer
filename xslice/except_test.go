package xslice

import "testing"

func TestExceptInts(t *testing.T) {

	var cases = []struct {
		Param1 []int
		Param2 []int
		Want   []int
	}{
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{1, 2, 3, 4, 5, 6, 8}, Want: []int{}},
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{1, 2, 3, 4, 5, 6, 6}, Want: []int{8, 8}},
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{1, 2, 8}, Want: []int{3, 4, 5, 6, 6}},
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{5, 6, 8}, Want: []int{1, 2, 3, 4}},
		{Param1: []int{1, 2, 3, 3}, Param2: []int{5, 6, 8}, Want: []int{1, 2, 3, 3}},
		{Param1: []int{1, 2, 3, 3}, Param2: []int{3, 3, 4, 5}, Want: []int{1, 2}},
		{Param1: []int{1, 2, 3, 3}, Param2: []int{1, 2, 3, 5}, Want: []int{}},
	}

	for _, c := range cases {
		got := ExceptInts(c.Param1, c.Param2)
		got = SortInts(got)
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, param2: %v, want: %v, but got: %v", c.Param1, c.Param2, c.Want, got)
		}
	}

}

func TestExceptUnstableInts(t *testing.T) {

	var cases = []struct {
		Param1 []int
		Param2 []int
		Want   []int
	}{
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{1, 2, 3, 4, 5, 6, 8}, Want: []int{}},
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{1, 2, 3, 4, 5, 6, 6}, Want: []int{8, 8}},
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{1, 2, 8}, Want: []int{3, 4, 5, 6, 6}},
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{5, 6, 8}, Want: []int{1, 2, 3, 4}},
		{Param1: []int{1, 2, 3, 3}, Param2: []int{5, 6, 8}, Want: []int{1, 2, 3, 3}},
		{Param1: []int{1, 2, 3, 3}, Param2: []int{3, 3, 4, 5}, Want: []int{1, 2}},
		{Param1: []int{1, 2, 3, 3}, Param2: []int{1, 2, 3, 5}, Want: []int{}},
	}

	for _, c := range cases {
		got := ExceptUnstableInts(c.Param1, c.Param2)
		got = SortInts(got)
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, param2: %v, want: %v, but got: %v", c.Param1, c.Param2, c.Want, got)
		}
	}

}

func TestExceptSortedInts(t *testing.T) {

	var cases = []struct {
		Param1 []int
		Param2 []int
		Want   []int
	}{
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{1, 2, 3, 4, 5, 6, 8}, Want: []int{}},
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{1, 2, 3, 4, 5, 6, 6}, Want: []int{8, 8}},
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{1, 2, 8}, Want: []int{3, 4, 5, 6, 6}},
		{Param1: []int{1, 2, 3, 4, 5, 6, 6, 8, 8}, Param2: []int{5, 6, 8}, Want: []int{1, 2, 3, 4}},
		{Param1: []int{1, 2, 3, 3}, Param2: []int{5, 6, 8}, Want: []int{1, 2, 3, 3}},
		{Param1: []int{1, 2, 3, 3}, Param2: []int{3, 3, 4, 5}, Want: []int{1, 2}},
		{Param1: []int{1, 2, 3, 3}, Param2: []int{1, 2, 3, 5}, Want: []int{}},
	}

	for _, c := range cases {
		got := ExceptSortInts(c.Param1, c.Param2)
		got = SortInts(got)
		if !EqualInts(c.Want, got) {
			t.Fatalf("param1: %v, param2: %v, want: %v, but got: %v", c.Param1, c.Param2, c.Want, got)
		}
	}

}
