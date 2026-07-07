package graphs

import (
	"fmt"
	"strconv"
	"strings"
)

func pacificAtlantic(matrix [][]int) [][]int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}

	rows, cols := len(matrix), len(matrix[0])
	pacific := make(map[string]bool)
	atlantic := make(map[string]bool)

	var dfs func(startR, startC, r, c int, visited map[string]bool)
	dfs = func(startR, startC, r, c int, visited map[string]bool) {
		key := fmt.Sprintf("%d,%d", r, c)
		if visited[key] {
			return
		}

		visited[key] = true

		if r == 0 || c == 0 {
			pacific[fmt.Sprintf("%d,%d", startR, startC)] = true
		}
		if r == rows-1 || c == cols-1 {
			atlantic[fmt.Sprintf("%d,%d", startR, startC)] = true
		}

		directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && matrix[nr][nc] <= matrix[r][c] {
				dfs(startR, startC, nr, nc, visited)
				delete(visited, fmt.Sprintf("%d,%d", nr, nc))
			}
		}
	}

	// Perform full DFS from each cell.
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			visited := make(map[string]bool)
			dfs(r, c, r, c, visited)
		}
	}

	var result [][]int
	for cell := range pacific {
		if atlantic[cell] {
			parts := strings.Split(cell, ",")
			r, _ := strconv.Atoi(parts[0])
			c, _ := strconv.Atoi(parts[1])
			result = append(result, []int{r, c})
		}
	}
	return result
}
