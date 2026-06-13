package quite_hard

type TreeAncestor struct {
	sTable [][]int
}

func Constructor(n int, parent []int) TreeAncestor {
	col := 1
	for (1 << col) <= n {
		col++
	}

	sTable := make([][]int, n)
	for i := range n {
		sTable[i] = make([]int, col)
		sTable[i][0] = parent[i]
	}

	for j := 1; j < col; j++ {
		for i := range n {
			sTable[i][j] = -1
			if sTable[i][j-1] != -1 {
				sTable[i][j] = sTable[sTable[i][j-1]][j-1]
			}
		}
	}

	return TreeAncestor{
		sTable: sTable,
	}
}

func (this *TreeAncestor) GetKthAncestor(node int, k int) int {
	if k == 1 {
		return this.sTable[node][0]
	}

	for i := 0; 1<<i <= k; i++ {
		if k&(1<<i) != 0 {
			node = this.sTable[node][i]
			if node == -1 {
				return -1
			}
		}
	}

	return node
}

/**
 * Your TreeAncestor object will be instantiated and called as such:
 * obj := Constructor(n, parent);
 * param_1 := obj.GetKthAncestor(node,k);
 */
