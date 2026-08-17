package dynamic_programming

// https://leetcode.com/problems/stone-game-v
func stoneGameV(stoneValue []int) int {
	n := len(stoneValue)
	memo := make([][]int, n)

	prefixArr := make([]int, n)
	for i := range n {
		memo[i] = make([]int, n)

		prefixArr[i] = stoneValue[i]
		if i > 0 {
			prefixArr[i] += prefixArr[i-1]
		}
	}

	compSum := func(l, r int) int {
		sum := prefixArr[r]
		if l > 0 {
			sum -= prefixArr[l-1]
		}

		return sum
	}

	var dfs func(l, r int) int
	dfs = func(l, r int) int {
		if l == r {
			return 0
		}

		if memo[l][r] != 0 {
			return memo[l][r]
		}

		maxScore := 0
		for k := l; k < r; k++ {
			leftSum := compSum(l, k)
			rightSum := compSum(k+1, r)

			currScore := 0
			if leftSum < rightSum {
				currScore = leftSum + dfs(l, k)
			}

			if leftSum > rightSum {
				currScore = rightSum + dfs(k+1, r)
			}

			if leftSum == rightSum {
				currScore = max(dfs(l, k), dfs(k+1, r)) + leftSum
			}

			maxScore = max(maxScore, currScore)
		}

		memo[l][r] = maxScore
		return maxScore
	}

	return dfs(0, n-1)
}
