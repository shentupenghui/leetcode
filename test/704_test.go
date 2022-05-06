package test

import (
	"testing"
	"xabc.site/leetcode/core"
)

func TestSearch(t *testing.T) {
	var tests = []struct {
		in struct {
			nums   []int
			target int
		} // input
		expected int // expected result
	}{
		{
			in: struct {
				nums   []int
				target int
			}{
				nums:   []int{1, 2, 3, 4},
				target: 3,
			},
			expected: 2,
		},
		{
			in: struct {
				nums   []int
				target int
			}{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			expected: -1,
		},
	}
	for _, test := range tests {
		actual := core.Search(test.in.nums, test.in.target)
		if actual != test.expected {
			t.Errorf("search(%v, %v) = %d; expected %d", test.in.nums, test.in.target, actual, test.expected)
		}
	}
}
