package dynamic_programming

// https://leetcode.com/problems/maximum-number-of-jumps-to-reach-the-last-index
func maximumJumps(nums []int, target int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range n {
		dp[i] = -1
	}
	dp[0] = 0

	for j := range n {
		for i := range j {
			if dp[i] != -1 && isWithin(target, nums[j]-nums[i]) {
				dp[j] = max(dp[j], dp[i]+1)
			}
		}
	}

	return dp[n-1]
}

func isWithin(bound, diff int) bool {
	return diff >= (-1*bound) && diff <= bound
}
