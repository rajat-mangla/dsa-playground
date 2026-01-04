package binary_search

import (
	"math"
)

// LeetCode #875. Koko Eating Bananas
// https://leetcode.com/problems/koko-eating-bananas/
func minEatingSpeed(nums []int, h int) int {
	low := 1
	high := 1

	for _, num := range nums {
		high = max(high, num)
	}
	ans := high

	for low <= high {
		mid := (low + high) / 2

		numHrs := findNumHours(nums, mid)
		if numHrs > h {
			low = mid + 1
			continue
		}

		high = mid - 1
		ans = mid
	}
	return ans
}

func findNumHours(nums []int, rate int) int {
	total := 0
	for _, num := range nums {
		timeTaken := math.Ceil(float64(num) / float64(rate))
		total += int(timeTaken)
	}
	return total
}
