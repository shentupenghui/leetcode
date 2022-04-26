package test

import (
	"testing"
	"xabc.site/leetcode/core"
)

func TestPick(t *testing.T) {
	var tests = []struct {
		in       int // input
		expected int // expected result
	}{
		{
			in:       4,
			expected: 1,
		},
		{
			in:       11,
			expected: 1,
		},
	}
	nums := &core.Solution{1, 4, 2, 3, 45, 6}
	for _, test := range tests {
		actual := nums.Pick(test.in)
		if actual != test.expected {
			t.Errorf("Pick(%d) = %d; expected %d", test.in, actual, test.expected)
		}
	}
}
