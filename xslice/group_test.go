package xslice

import "testing"

func TestGroup(t *testing.T) {

	var cases = []struct {
		Param []int
		Want  [][]int
	}{
		{Param: []int{0, 1, 2, 3, 4}, Want: [][]int{{0}, {1}, {2}, {3}, {4}}},
		{Param: []int{0, 1, 2, 3, 4, 4}, Want: [][]int{{0}, {1}, {2}, {3}, {4, 4}}},
		{Param: []int{4, 2, 0, 1, 4, 3}, Want: [][]int{{0}, {1}, {2}, {3}, {4, 4}}},
		{Param: []int{4, 2, 5, 5, 0, 1, 4, 3}, Want: [][]int{{0}, {1}, {2}, {3}, {4, 4}, {5, 5}}},
		{Param: []int{4, 2}, Want: [][]int{{2}, {4}}},
	}

	for _, c := range cases {
		var got [][]int
		GroupInts(c.Param, func(indexes ...int) {
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
