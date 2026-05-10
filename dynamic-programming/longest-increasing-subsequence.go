package dynamic_programming

import (
	"slices"
)

// https://leetcode.com/problems/longest-increasing-subsequence

// dp solution
// Time complexity: O(n^2)
// space complexity: O(n)
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range n {
		dp[i] = 1
	}

	for j := range n {
		for i := range j {
			if nums[j] > nums[i] {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}

	result := 0
	for i := range n {
		result = max(result, dp[i])
	}

	return result
}

// binary seach solution
// Time complexity: O(n logn)
// space complexity: O(n)
func lengthOfLISV2(nums []int) int {
	n := len(nums)
	var tails []int

	for i := range n {

		idx, _ := slices.BinarySearch(tails, nums[i])
		if idx == len(tails) {
			tails = append(tails, nums[i])
			continue
		}

		tails[idx] = nums[i]
	}

	return len(tails)
}
