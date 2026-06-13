package quite_hard

import (
	"container/heap"
	"math"
)

type node struct {
	min int
	max int
}

type segmentTree struct {
	tree []node
	size int
}

func NewSegmentTree(arr []int) segmentTree {
	n := len(arr)

	st := segmentTree{
		tree: make([]node, 4*n),
		size: n,
	}
	st.build(arr, 0, 0, n-1)
	return st
}

func (t *segmentTree) build(nums []int, pos, left, right int) {
	if left == right {
		t.tree[pos] = node{min: nums[left], max: nums[right]}
		return
	}

	mid := (left + right) / 2
	t.build(nums, 2*pos+1, left, mid)
	t.build(nums, 2*pos+2, mid+1, right)

	leftNode := t.tree[2*pos+1]
	rightNode := t.tree[2*pos+2]
	t.tree[pos] = node{
		min: min(leftNode.min, rightNode.min),
		max: max(leftNode.max, rightNode.max),
	}
}

func (t *segmentTree) Query(left, right int) int {
	n := t.query(0, 0, t.size-1, left, right)
	return n.max - n.min
}

func (t *segmentTree) query(pos, left, right, qLeft, qRight int) node {
	if qLeft > right || qRight < left {
		return node{min: math.MaxInt, max: math.MinInt}
	}

	if qLeft <= left && right <= qRight {
		return t.tree[pos]
	}

	mid := (left + right) / 2
	leftNode := t.query(2*pos+1, left, mid, qLeft, qRight)
	rightNode := t.query(2*pos+2, mid+1, right, qLeft, qRight)
	return node{
		min: min(leftNode.min, rightNode.min),
		max: max(leftNode.max, rightNode.max),
	}
}

type Item struct {
	val   int
	left  int
	right int
}
type maxHeap []Item

func (h maxHeap) Len() int { return len(h) }

func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h maxHeap) Less(i, j int) bool { return h[i].val > h[j].val }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(Item))
}

func (h *maxHeap) Pop() any {
	x := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return x
}

func maxTotalValue(nums []int, k int) int64 {
	st := NewSegmentTree(nums)
	maxH := &maxHeap{}

	size := len(nums)
	for i := range size {
		item := Item{
			val:   st.Query(i, size-1),
			left:  i,
			right: size - 1,
		}
		heap.Push(maxH, item)
	}

	var totalScore int64
	for maxH.Len() > 0 && k > 0 {
		item := heap.Pop(maxH).(Item)

		totalScore += int64(item.val)
		if item.right > item.left {
			heap.Push(maxH, Item{
				val:   st.Query(item.left, item.right-1),
				left:  item.left,
				right: item.right - 1,
			})
		}

		k--
	}

	return totalScore
}
