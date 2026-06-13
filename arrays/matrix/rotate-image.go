package matrix

// https://leetcode.com/problems/rotate-image
func rotate(matrix [][]int) {
	n := len(matrix)

	for i := range n {
		for j := range n {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for i := range n {
		for j := range n {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}
