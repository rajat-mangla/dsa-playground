package heaps

import (
	"container/heap"
)

type minHeap []int

func (h *minHeap) Len() int { return len(*h) }

func (h *minHeap) Less(i, j int) bool { return (*h)[i] < (*h)[j] }

func (h *minHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *minHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func minOperations(nums []int, k int) int {
	h := &minHeap{}

	for _, num := range nums {
		heap.Push(h, num)
	}

	numOperations := 0
	for h.Len() >= 2 && (*h)[0] < k {
		x := heap.Pop(h).(int)
		y := heap.Pop(h).(int)
		newElement := 2*x + y
		heap.Push(h, newElement)
		numOperations++
	}

	if (*h)[0] < k {
		return -1
	}

	return numOperations
}
