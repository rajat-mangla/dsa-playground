package bst

func lca(root *TreeNode, p, q int) *TreeNode {
	if root == nil {
		return root
	}

	currVal := root.Data
	if currVal > p && currVal > q {
		return lca(root.Left, p, q)
	}

	if currVal < p && currVal < q {
		return lca(root.Right, p, q)
	}

	return root
}
