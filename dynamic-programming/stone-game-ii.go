package dynamic_programming

// https://leetcode.com/problems/stone-game-ii
func stoneGameII(piles []int) int {
	n := len(piles)

	suffixSum := make([]int, n)
	suffixSum[n-1] = piles[n-1]

	i := n - 2
	for i >= 0 {
		suffixSum[i] = piles[i] + suffixSum[i+1]
		i--
	}

	memoTable := make([][]int, n)
	for i := range n {
		memoTable[i] = make([]int, n+1)
	}

	var dfs func(i, m int) int
	dfs = func(i, m int) int {
		if i+2*m >= n {
			return suffixSum[i]
		}

		if memoTable[i][m] != 0 {
			return memoTable[i][m]
		}

		maxStones := 0
		for x := 1; x <= 2*m; x++ {
			newM := max(x, m)

			currStones := suffixSum[i] - dfs(i+x, newM)
			maxStones = max(maxStones, currStones)
		}

		memoTable[i][m] = maxStones
		return maxStones
	}

	return dfs(0, 1)
}
