package graphs

import (
	"container/heap"
)

// https://leetcode.com/problems/find-the-safest-path-in-a-grid
type maxHeap [][3]int

func (h maxHeap) Len() int { return len(h) }

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h maxHeap) Less(i, j int) bool {
	return h[i][2] > h[j][2]
}

func (h *maxHeap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.([3]int))
}

func maximumSafenessFactor(grid [][]int) int {
	n := len(grid)

	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return 0
	}

	dist := make([][]int, n)
	for i := range n {
		dist[i] = make([]int, n)
		for j := range n {
			dist[i][j] = -1
			if grid[i][j] == 1 {
				dist[i][j] = 0
			}
		}
	}

	dirs := [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	var bfs func(queue [][3]int)
	bfs = func(queue [][3]int) {
		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			row, col, currDist := curr[0], curr[1], curr[2]

			for _, dir := range dirs {
				nRow, nCol := row+dir[0], col+dir[1]

				if nRow < 0 || nRow >= n || nCol < 0 || nCol >= n {
					continue
				}

				newDist := dist[nRow][nCol]
				if newDist != -1 && newDist <= currDist+1 {
					continue
				}

				dist[nRow][nCol] = currDist + 1
				queue = append(queue, [3]int{nRow, nCol, currDist + 1})
			}
		}
	}

	for i := range n {
		for j := range n {
			if grid[i][j] == 1 {
				bfs([][3]int{{i, j, 0}})
			}
		}
	}

	visited := make([][]bool, n)
	for i := range n {
		visited[i] = make([]bool, n)
		for j := range n {
			visited[i][j] = false
		}
	}

	ans := 2 * n
	h := &maxHeap{{0, 0, dist[0][0]}}
	visited[0][0] = true
	for h.Len() > 0 {
		curr := heap.Pop(h).([3]int)

		row, col, currDist := curr[0], curr[1], curr[2]
		ans = min(ans, currDist)
		if row == n-1 && col == n-1 {
			break
		}

		for _, dir := range dirs {
			nRow, nCol := row+dir[0], col+dir[1]

			if nRow < 0 || nRow >= n || nCol < 0 || nCol >= n {
				continue
			}

			if visited[nRow][nCol] {
				continue
			}

			visited[nRow][nCol] = true
			heap.Push(h, [3]int{nRow, nCol, dist[nRow][nCol]})
		}
	}

	return ans
}
