package binary_trees

import (
	"sort"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Data int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func verticalTraversal(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}

	type NodeInfo struct {
		node *TreeNode
		row  int
		col  int
	}

	type Value struct {
		row int
		val int
	}

	columnTable := map[int][]Value{}
	minCol, maxCol := 0, 0
	queue := []NodeInfo{{node: root, row: 0, col: 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		arr := columnTable[curr.col]
		columnTable[curr.col] = append(arr, Value{
			row: curr.row,
			val: curr.node.Data,
		})

		if curr.node.Left != nil {
			queue = append(queue, NodeInfo{node: curr.node.Left, row: curr.row + 1, col: curr.col - 1})
			if curr.col-1 < minCol {
				minCol = curr.col - 1
			}
		}
		if curr.node.Right != nil {
			queue = append(queue, NodeInfo{node: curr.node.Right, row: curr.row + 1, col: curr.col + 1})
			if curr.col+1 > maxCol {
				maxCol = curr.col + 1
			}
		}
	}

	for col := minCol; col <= maxCol; col++ {
		var colNodes []int
		if rows, exists := columnTable[col]; exists {
			// Sort by row, then by value
			sort.Slice(rows, func(i, j int) bool {
				if rows[i].row == rows[j].row {
					return rows[i].val < rows[j].val
				}
				return rows[i].row < rows[j].row
			})

			for _, v := range rows {
				colNodes = append(colNodes, v.val)
			}
		}
		result = append(result, colNodes)
	}

	return result
}
