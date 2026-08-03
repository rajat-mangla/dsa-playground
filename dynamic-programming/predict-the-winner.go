package dynamic_programming

// https://leetcode.com/problems/predict-the-winner/
// [Blog] https://medium.com/algorithms-digest/predict-the-winner-16668e9c1cb8
func predictTheWinner(nums []int) bool {
	n := len(nums)
	totalArr := make([]int, n)
	for i := range n {
		totalArr[i] = nums[i]
		if i > 0 {
			totalArr[i] += totalArr[i-1]
		}
	}

	dp := make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n)
	}

	for i := range n {
		for j := range n - i {

			start := j
			end := i + j

			l := end - start + 1
			if l == 1 {
				dp[start][end] = nums[start]
				continue
			}

			if l == 2 {
				dp[start][end] = max(nums[start], nums[end])
				continue
			}

			dp[start][end] = max(
				nums[start]+(totalArr[end]-totalArr[start])-dp[start+1][end],
				nums[end]+(totalArr[end-1]-totalArr[start]+nums[start])-dp[start][end-1])
		}
	}

	p1 := dp[0][n-1]
	return p1 >= (totalArr[n-1] - p1)
}
