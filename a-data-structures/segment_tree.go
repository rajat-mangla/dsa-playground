package a_data_structures

import (
	"math"
)

type Node struct {
	sum      int
	minValue int
	maxValue int
}

type SegmentTree struct {
	tree []Node
	size int
}

func NewSegmentTree(arr []int) SegmentTree {
	n := len(arr)
	st := SegmentTree{
		tree: make([]Node, 4*n),
		size: n,
	}

	st.build(arr, 0, 0, n-1)
	return st
}

func (s *SegmentTree) build(arr []int, pos, left, right int) {
	if left == right {
		s.tree[pos] = Node{arr[left], arr[left], arr[left]}
		return
	}

	mid := (left + right) / 2
	s.build(arr, 2*pos+1, left, mid)
	s.build(arr, 2*pos+2, mid+1, right)

	leftNode := s.tree[2*pos+1]
	rightNode := s.tree[2*pos+2]

	s.tree[pos] = Node{
		sum:      leftNode.sum + rightNode.sum,
		maxValue: max(leftNode.maxValue, rightNode.maxValue),
		minValue: min(leftNode.minValue, rightNode.minValue),
	}
}

func (s *SegmentTree) querySegmentTree(pos, left, right, qLeft, qRight int) *Node {
	if right < qLeft || qRight < left {
		return nil
	}

	if qLeft <= left && right <= qRight {
		return &s.tree[pos]
	}

	mid := (left + right) / 2
	leftNode := s.querySegmentTree(2*pos+1, left, mid, qLeft, qRight)
	rightNode := s.querySegmentTree(2*pos+2, mid+1, right, qLeft, qRight)

	n := &Node{minValue: math.MaxInt}
	if leftNode != nil {
		n.sum += leftNode.sum
		n.maxValue = max(n.maxValue, leftNode.maxValue)
		n.minValue = min(n.minValue, leftNode.minValue)
	}

	if rightNode != nil {
		n.sum += rightNode.sum
		n.maxValue = max(n.maxValue, rightNode.maxValue)
		n.minValue = min(n.minValue, rightNode.minValue)
	}

	return n
}

func (s *SegmentTree) UpdateSegmentTree(idx, val int) {
	s.updateSegmentTree(0, 0, s.size-1, idx, val)
}

func (s *SegmentTree) updateSegmentTree(pos, left, right, idx, val int) {
	if left == right {
		s.tree[pos].minValue = val
		s.tree[pos].maxValue = val
		s.tree[pos].sum = val
		return
	}

	mid := (left + right) / 2
	leftPos := 2*pos + 1
	rightPos := 2*pos + 2

	if idx <= mid {
		s.updateSegmentTree(leftPos, left, mid, idx, val)
	}
	if idx > mid {
		s.updateSegmentTree(rightPos, mid+1, right, idx, val)
	}

	leftNode := s.tree[leftPos]
	rightNode := s.tree[rightPos]
	// update
	s.tree[pos].sum = leftNode.sum + rightNode.sum
	s.tree[pos].minValue = min(leftNode.minValue, rightNode.minValue)
	s.tree[pos].maxValue = max(leftNode.maxValue, rightNode.maxValue)
}
