package graphs

// https://leetcode.com/problems/detect-cycles-in-2d-grid
func containsCycle(grid [][]byte) bool {
	m := len(grid)
	n := len(grid[0])

	dir := [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	visited := make([][]bool, m)
	for i := range m {
		visited[i] = make([]bool, n)
	}

	var dfs func(r, c, pR, pC int) bool
	dfs = func(r, c, pR, pC int) bool {
		if visited[r][c] {
			return true
		}

		visited[r][c] = true
		for i := range dir {
			newR := r + dir[i][0]
			newC := c + dir[i][1]

			if newR < 0 || newR >= m || newC < 0 || newC >= n || grid[newR][newC] != grid[pR][pC] {
				continue
			}

			if newR == pR && newC == pC {
				continue
			}

			if dfs(newR, newC, r, c) {
				return true
			}
		}

		return false
	}

	for i := range grid {
		for j := range grid[i] {
			if !visited[i][j] && dfs(i, j, i, j) {
				return true
			}
		}
	}

	return false
}
