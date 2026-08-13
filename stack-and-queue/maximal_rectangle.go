package stack_and_queue

func maximalRectangle(matrix [][]byte) int {
	rowLen := len(matrix[0])
	colLen := len(matrix)

	heights := make([][]int, colLen)
	for i := range heights {
		heights[i] = make([]int, rowLen)
	}

	for i := range colLen {
		for j := range rowLen {
			if matrix[i][j] == '1' {
				heights[i][j] = 1
			}
		}
	}

	total := 0
	for i := range colLen {
		if i > 0 {
			for j := range rowLen {
				if heights[i][j] == 0 {
					continue
				}

				heights[i][j] += heights[i-1][j]
			}
		}

		total = max(total, largestRectangleArea(heights[i]))
	}
	return total
}
