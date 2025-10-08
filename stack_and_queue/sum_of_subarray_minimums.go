package stack_and_queue

func sumSubarrayMins(nums []int) int64 {
	var total int64
	pse := psee(nums)
	nse := nse(nums)
	for i := range nums {
		left := i - pse[i]
		right := nse[i] - i
		total += int64(right * left * nums[i])
	}
	return total
}

func nse(nums []int) []int {
	l := len(nums)
	stack := []int{}
	ans := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}

		idx := l
		if len(stack) > 0 {
			idx = stack[len(stack)-1]
		}

		stack = append(stack, i)
		ans[i] = idx
	}

	return ans
}

func psee(nums []int) []int {
	l := len(nums)
	stack := []int{}
	ans := make([]int, l)
	for i := range nums {
		for len(stack) > 0 && nums[stack[len(stack)-1]] > nums[i] {
			stack = stack[:len(stack)-1]
		}

		idx := -1
		if len(stack) > 0 {
			idx = stack[len(stack)-1]
		}

		stack = append(stack, i)
		ans[i] = idx
	}

	return ans
}
