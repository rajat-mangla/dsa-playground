package dynamic_programming

// https://leetcode.com/problems/burst-balloons/
// [Blog] https://medium.com/algorithms-digest/bursting-balloons-1820664a4ffa
func maxCoins(nums []int) int {
	n := len(nums)
	m := n + 2

	dpNums := make([]int, m)
	dpNums[0] = 1
	dpNums[m-1] = 1
	for i := range n {
		dpNums[i+1] = nums[i]
	}

	dpState := make([][]int, m)
	for i := range m {
		dpState[i] = make([]int, m)
	}

	for l := 2; l < m; l++ {
		for i := 0; i < m-l; i++ {
			j := i + l

			for k := i + 1; k < j; k++ {
				currState := dpState[i][k] + dpState[k][j] + dpNums[i]*dpNums[k]*dpNums[j]
				dpState[i][j] = max(dpState[i][j], currState)
			}
		}
	}

	return dpState[0][m-1]
}
