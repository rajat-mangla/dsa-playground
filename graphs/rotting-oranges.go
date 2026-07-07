package graphs

// Question:
// LC 994:Rotting Oranges
// https://leetcode.com/problems/rotting-oranges
func orangesRotting(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])

	var queue [][2]int
	freshOranges := 0
	for i := range rows {
		for j := range cols {
			if grid[i][j] == 1 {
				freshOranges++
			}

			if grid[i][j] == 2 {
				queue = append(queue, [2]int{i, j})
			}
		}
	}

	minutes := 0

	dirs := [][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for len(queue) > 0 && freshOranges > 0 {
		size := len(queue)

		for i := range size {
			curr := queue[i]

			for _, dir := range dirs {
				nRow, nCol := curr[0]+dir[0], curr[1]+dir[1]

				if nRow < 0 || nCol < 0 || nRow >= rows || nCol >= cols {
					continue
				}

				if grid[nRow][nCol] != 1 {
					continue
				}

				grid[nRow][nCol] = 2
				freshOranges--
				queue = append(queue, [2]int{nRow, nCol})
			}
		}

		queue = queue[size:]
		minutes++
	}

	if freshOranges > 0 {
		return -1
	}

	return minutes
}
