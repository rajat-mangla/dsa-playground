package binary_trees

import (
	"math"
)

// TreeNode https://leetcode.com/problems/binary-tree-maximum-path-sum
// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	dfs(root, &maxSum)
	return maxSum
}

func dfs(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}

	lSum := dfs(node.Left, maxSum)
	rSum := dfs(node.Right, maxSum)

	currSum := max(node.Data, node.Data+max(lSum, rSum))
	*maxSum = max(*maxSum, max(currSum, lSum+rSum+node.Data))
	return currSum
}
