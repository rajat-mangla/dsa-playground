package binary_trees

//Question: https://takeuforward.org/plus/dsa/problems/boundary-traversal

/* Definition for a binary tree node.
type TreeNode struct {
    Data  int
    Left  *TreeNode
    Right *TreeNode
}
*/

func Boundary(root *TreeNode) []int {
	//your code goes here
	if root == nil {
		return nil
	}

	if isLeaf(root) {
		return []int{root.Data}
	}
	result := []int{root.Data}
	leftBoundary(root, &result)
	leaves(root, &result)
	rightBoundary(root, &result)

	return result
}

func leftBoundary(node *TreeNode, result *[]int) {
	curr := node.Left

	for curr != nil {
		if !isLeaf(curr) {
			*result = append(*result, curr.Data)
		}

		if curr.Left != nil {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
}

func rightBoundary(node *TreeNode, result *[]int) {
	stack := []int{}
	curr := node.Right

	for curr != nil {
		if !isLeaf(curr) {
			stack = append(stack, curr.Data)
		}

		if curr.Right != nil {
			curr = curr.Right
		} else {
			curr = curr.Left
		}
	}

	size := len(stack)
	for i := size - 1; i >= 0; i-- {
		*result = append(*result, stack[i])
	}
}

func leaves(node *TreeNode, result *[]int) {
	if isLeaf(node) {
		*result = append(*result, node.Data)
		return
	}

	if node.Left != nil {
		leaves(node.Left, result)
	}

	if node.Right != nil {
		leaves(node.Right, result)
	}
}

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}
