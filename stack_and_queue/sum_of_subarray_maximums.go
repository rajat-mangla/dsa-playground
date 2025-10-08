package stack_and_queue

func sumSubArrayMax(nums []int) int64 {
	nge := nge(nums)
	pgee := pgee(nums)
	var total int64
	for i := range nums {
		right := nge[i] - i
		left := i - pgee[i]
		total += int64(left * right * nums[i])
	}

	return total
}

func nge(nums []int) []int {
	l := len(nums)
	stack := []int{}

	nge := make([]int, l)
	for i := l - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		idx := l
		if len(stack) > 0 {
			idx = stack[len(stack)-1]
		}

		nge[i] = idx
		stack = append(stack, i)
	}

	return nge
}

func pgee(nums []int) []int {
	l := len(nums)
	stack := []int{}

	pgee := make([]int, l)
	for i := 0; i < l; i++ {
		for len(stack) > 0 && nums[i] > nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}

		idx := -1
		if len(stack) > 0 {
			idx = stack[len(stack)-1]
		}

		pgee[i] = idx
		stack = append(stack, i)
	}

	return pgee
}
