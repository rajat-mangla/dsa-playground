package dynamic_programming

// https://leetcode.com/problems/stone-game-iii
func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	dp := make([]int, n+1)

	for i := n - 1; i >= 0; i-- {
		runningSum := 0
		maxAdvantage := -1 << 31

		for k := 0; k < 3 && i+k < n; k++ {
			runningSum += stoneValue[i+k]
			maxAdvantage = max(maxAdvantage, runningSum-dp[i+k+1])
		}

		dp[i] = maxAdvantage
	}

	if dp[0] > 0 {
		return "Alice"
	}

	if dp[0] < 0 {
		return "Bob"
	}

	return "Tie"
}
