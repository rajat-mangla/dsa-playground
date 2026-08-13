package heaps

import (
	"container/heap"
)

type maxHeap []int

func (h *maxHeap) Len() int { return len(*h) }

func (h *maxHeap) Less(i, j int) bool { return (*h)[i] > (*h)[j] }

func (h *maxHeap) Swap(i, j int) { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }

func (h *maxHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	// Initialize a max-heap to store the energy differences where we used energy drinks
	h := &maxHeap{}
	n := len(heights)

	for i := 0; i < n-1; i++ {
		diff := heights[i+1] - heights[i]

		if diff <= 0 {
			// No resources needed for descent or flat terrain
			continue
		}

		// Use energy drinks for this climb
		heap.Push(h, diff)
		bricks -= diff

		if bricks < 0 && ladders == 0 {
			// No resources left to climb further
			return i
		}

		// If bricks run out, try to use ladder
		if bricks < 0 && ladders > 0 {
			// We need to free up energy by replacing an energy-funded climb with a rope.
			// To make bricks non-negative by replacing the *max* expensive
			// energy climb with a rope, we pop the largest element from the heap
			// and add it back to bricks.
			if h.Len() > 0 {
				maxEnergyUsed := heap.Pop(h).(int)
				bricks += maxEnergyUsed // Recover bricks
				ladders--
			} else {
				// This case should not be hit if logic is sound, but as a safeguard
				// if bricks becomes negative and heap is empty.
				return i
			}
		}
	}

	// Reached the end of the heights array
	return n - 1
}
