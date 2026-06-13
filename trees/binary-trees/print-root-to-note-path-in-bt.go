package binary_trees

func allRootToLeaf(root *TreeNode) [][]int {
	//your code goes here
	var result [][]int
	if root == nil {
		return result
	}

	dfsPath(root, &result, []int{})
	return result
}

func dfsPath(node *TreeNode, result *[][]int, currPath []int) {
	currPath = append(currPath, node.Data)
	if isLeaf(node) {
		temp := make([]int, len(currPath))
		copy(temp, currPath)
		*result = append(*result, temp)
		return
	}

	if node.Left != nil {
		dfsPath(node.Left, result, currPath)
	}

	if node.Right != nil {
		dfsPath(node.Right, result, currPath)
	}

	currPath = currPath[:len(currPath)-1]
}

//func isLeaf(node *TreeNode) bool {
//	return node.Left == nil && node.Right == nil
//}
