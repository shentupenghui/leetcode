package test

import (
	"testing"
	"xabc.site/leetcode/core"
)

func TestRemoveCoveredIntervals(t *testing.T) {
	var tests = []struct {
		input  [][]int
		expect int
	}{
		{
			input:  [][]int{{1, 4}, {3, 6}, {2, 8}},
			expect: 2,
		},
		{
			input:  [][]int{{34335, 39239}, {15875, 91969}, {29673, 66453}, {53548, 69161}, {40618, 93111}},
			expect: 2,
		},
	}
	for _, test := range tests {
		actual := core.RemoveCoveredIntervals(test.input)
		if actual != test.expect {
			t.Errorf("RemoveCoveredIntervals(%v) = %d; expected %d", test.input, actual, test.expect)
		}
	}

}
