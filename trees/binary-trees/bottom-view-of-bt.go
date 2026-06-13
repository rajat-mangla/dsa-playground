package binary_trees

func bottomView(root *TreeNode) []int {
	//your code goes here
	var result []int
	if root == nil {
		return result
	}

	type NodeInfo struct {
		node *TreeNode
		col  int
	}

	columnTable := map[int]int{}
	minCol, maxCol := 0, 0
	queue := []NodeInfo{{node: root, col: 0}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		columnTable[curr.col] = curr.node.Data
		minCol = min(minCol, curr.col)
		maxCol = max(maxCol, curr.col)

		if curr.node.Left != nil {
			queue = append(queue, NodeInfo{node: curr.node.Left, col: curr.col - 1})
		}

		if curr.node.Right != nil {
			queue = append(queue, NodeInfo{node: curr.node.Right, col: curr.col + 1})
		}
	}

	for i := minCol; i <= maxCol; i++ {
		val, ok := columnTable[i]
		if ok {
			result = append(result, val)
		}
	}

	return result
}
