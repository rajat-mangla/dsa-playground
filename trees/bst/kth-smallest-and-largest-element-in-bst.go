package bst

type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

func kLargesSmall(root *TreeNode, k int) []int {
	return []int{kSmall(root, k), kLarge(root, k)}
}

func kSmall(root *TreeNode, k int) int {
	cnt := 0
	result := 0
	inOrder(root, k, &cnt, &result)
	return result
}

func inOrder(node *TreeNode, k int, cnt, result *int) {
	if node == nil || *cnt >= k {
		return
	}

	inOrder(node.Left, k, cnt, result)
	*cnt += 1
	if *cnt == k {
		*result = node.Data
	}
	inOrder(node.Right, k, cnt, result)
}

func kLarge(root *TreeNode, k int) int {
	cnt := 0
	result := 0
	inOrderL(root, k, &cnt, &result)
	return result
}

func inOrderL(node *TreeNode, k int, cnt, result *int) {
	if node == nil || *cnt >= k {
		return
	}

	inOrderL(node.Right, k, cnt, result)
	*cnt += 1
	if *cnt == k {
		*result = node.Data
	}
	inOrderL(node.Left, k, cnt, result)
}
