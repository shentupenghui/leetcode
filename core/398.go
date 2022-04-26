package core

import "math/rand"

type Solution []int

func Constructor(nums []int) Solution {
	return nums
}

func (s Solution) Pick(target int) (ans int) {
	count := 0
	for i, num := range s {
		if num == target {
			count++
			if rand.Intn(count) == 0 {
				ans = i
			}
		}
	}
	return
}
