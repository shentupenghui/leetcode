package core

import "fmt"

func RemoveCoveredIntervals(intervals [][]int) int {
	count := len(intervals)
	for k, v := range intervals {
		for i, v2 := range intervals {
			if i == k {
				continue
			}
			if v[0] >= v2[0] && v[1] <= v2[1] {
				fmt.Printf("[%d, %d]\n", v[0], v[1])
				count = count - 1
				break
			}
		}
	}
	return count
}
