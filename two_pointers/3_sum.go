package two_pointers

import (
	"slices"
)

// https://leetcode.com/problems/3sum/
func threeSum(nums []int) [][]int {
	// sorting is required for two sum walkthrough
	slices.Sort(nums)
	ans := [][]int{}

	l := len(nums)
	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		target := -1 * nums[i]
		j := i + 1
		k := l - 1

		// Two sum walkthrough
		for j < k {
			if nums[j]+nums[k] > target {
				k--
				continue
			}

			if nums[j]+nums[k] < target {
				j++
				continue
			}

			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				j++
				k--
				for j < k && nums[j] == nums[j-1] {
					j++
				}

				for j < k && nums[k] == nums[k+1] {
					k--
				}
			}
		}
	}
	return ans
}
