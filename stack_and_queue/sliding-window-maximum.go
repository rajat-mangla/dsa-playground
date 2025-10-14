package stack_and_queue

// The idea is to maintain a deque (double-ended queue) that stores indices of array elements.
// The deque is maintained in such a way that the elements corresponding to the indices in the deque are in decreasing order.
// For each new element, we remove elements from the back of the deque while they are less than or equal to the new element.
// This ensures that the front of the deque always contains the index of the maximum element for the current window.
// We also remove indices from the front of the deque if they are out of the bounds of the current window.
func maxSlidingWindow(nums []int, k int) []int {
	l := len(nums)
	var queueIdx []int
	ans := make([]int, l-k+1)
	for i := 0; i < l; i++ {
		if len(queueIdx) > 0 && queueIdx[0] <= i-k {
			queueIdx = queueIdx[1:]
		}

		for len(queueIdx) > 0 && nums[i] >= nums[queueIdx[len(queueIdx)-1]] {
			queueIdx = queueIdx[:len(queueIdx)-1]
		}

		queueIdx = append(queueIdx, i)
		if i >= k-1 {
			ans[i-k+1] = nums[queueIdx[0]]
		}
	}

	return ans
}
