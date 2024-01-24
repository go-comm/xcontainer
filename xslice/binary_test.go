package xslice

import "testing"

func TestBinaryGroup(t *testing.T) {

	var cases = []struct {
		Param []int
		Want  [][]int
	}{
		{Param: []int{0, 1, 2, 3, 4}, Want: [][]int{{0}, {1}, {2}, {3}, {4}}},
		{Param: []int{0, 1, 2, 3, 4, 4}, Want: [][]int{{0}, {1}, {2}, {3}, {4, 4}}},
		{Param: []int{0, 1, 2, 3, 4, 4}, Want: [][]int{{0}, {1}, {2}, {3}, {4, 4}}},
		{Param: []int{0, 1, 2, 3, 4, 4, 5, 5}, Want: [][]int{{0}, {1}, {2}, {3}, {4, 4}, {5, 5}}},
		{Param: []int{2, 4}, Want: [][]int{{2}, {4}}},
	}

	for _, c := range cases {
		var got [][]int
		arr := c.Param
		BinaryGroup(len(arr), func(i, j int) int { return arr[i] - arr[j] },
			func(indexes ...int) {
				got = append(got, MapToInts(len(indexes), func(j int) int { return c.Param[indexes[j]] }))
			})
		if len(got) != len(c.Want) {
			t.Fatal("no match length")
		}
		for i := 0; i < len(got); i++ {
			if !EqualInts(got[i], c.Want[i]) {
				t.Fatal("no match row")
			}
		}
	}

}
