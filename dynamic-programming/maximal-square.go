package dynamic_programming

// https://leetcode.com/problems/maximal-square/
// [Blog] https://algorithms-digest.medium.com/maximal-square-e0ba64fbebb6
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])

	maxSquare := make([][]int, m)
	for i := range m {
		maxSquare[i] = make([]int, n)
	}

	for i := range n {
		maxSquare[0][i] = int(matrix[0][i] - '0')
	}

	for i := range m {
		maxSquare[i][0] = int(matrix[i][0] - '0')
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j]-'0' == 0 {
				continue
			}

			maxSquare[i][j] = 1 + min(maxSquare[i-1][j], maxSquare[i][j-1], maxSquare[i-1][j-1])
		}
	}

	ans := 0
	for i := range m {
		for j := range n {
			ans = max(ans, maxSquare[i][j])
		}
	}
	return ans * ans
}
