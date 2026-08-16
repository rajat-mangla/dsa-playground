package segment_trees

// https://leetcode.com/problems/longest-substring-of-one-repeating-character
type node struct {
	totalLen  int
	prefixLen int
	suffixLen int
	maxLen    int

	// char
	leftChar  byte
	rightChar byte
}

type segmentTree struct {
	tree []node
	size int
}

func newSegmentTree(s string) segmentTree {
	byteArr := []byte(s)
	n := len(byteArr)
	sTree := segmentTree{
		tree: make([]node, 4*n),
		size: n,
	}

	sTree.build(byteArr, 0, 0, n-1)
	return sTree
}

func (s *segmentTree) build(arr []byte, pos, left, right int) {
	if left == right {
		s.tree[pos] = node{
			totalLen:  1,
			prefixLen: 1,
			suffixLen: 1,
			maxLen:    1,
			leftChar:  arr[left],
			rightChar: arr[right],
		}

		return
	}

	mid := (left + right) / 2
	leftChild := 2*pos + 1
	rightChild := 2*pos + 2
	s.build(arr, leftChild, left, mid)
	s.build(arr, rightChild, mid+1, right)

	s.merge(pos, s.tree[leftChild], s.tree[rightChild])
}

func (s *segmentTree) merge(pos int, leftChild, rightChild node) {
	currNode := node{
		totalLen:  leftChild.totalLen + rightChild.totalLen,
		prefixLen: leftChild.prefixLen,
		suffixLen: rightChild.suffixLen,
		maxLen:    max(leftChild.maxLen, rightChild.maxLen),
		//
		leftChar:  leftChild.leftChar,
		rightChar: rightChild.rightChar,
	}

	if leftChild.rightChar == rightChild.leftChar {
		// everything on left child is equal
		if leftChild.prefixLen == leftChild.totalLen {
			currNode.prefixLen = leftChild.prefixLen + rightChild.prefixLen
		}

		// everything on right child is equal
		if rightChild.suffixLen == rightChild.totalLen {
			currNode.suffixLen = rightChild.suffixLen + leftChild.suffixLen
		}

		mergeLen := leftChild.suffixLen + rightChild.prefixLen
		currNode.maxLen = max(currNode.maxLen, currNode.prefixLen, mergeLen, currNode.suffixLen)
	}

	s.tree[pos] = currNode
}

func (s *segmentTree) update(idx int, char byte) {
	s.updateTree(0, 0, s.size-1, idx, char)
}

func (s *segmentTree) updateTree(pos, left, right, idx int, char byte) {
	if left == right {
		s.tree[pos].leftChar = char
		s.tree[pos].rightChar = char
		return
	}

	mid := (left + right) / 2
	leftChild := 2*pos + 1
	rightChild := 2*pos + 2
	if idx <= mid {
		s.updateTree(leftChild, left, mid, idx, char)
	}

	if idx > mid {
		s.updateTree(rightChild, mid+1, right, idx, char)
	}

	s.merge(pos, s.tree[leftChild], s.tree[rightChild])
}

func longestRepeating(s string, queryCharacters string, queryIndices []int) []int {
	stree := newSegmentTree(s)

	k := len(queryIndices)
	ans := make([]int, k)
	for i := range k {
		idx := queryIndices[i]
		stree.update(idx, queryCharacters[i])
		ans[i] = stree.tree[0].maxLen
	}

	return ans
}
